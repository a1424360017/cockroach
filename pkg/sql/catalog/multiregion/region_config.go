// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Package multiregion provides functions and structs for interacting with the
// static multi-region state configured by users on their databases.
package multiregion

import (
	"sort"

	"github.com/cockroachdb/cockroach/pkg/config/zonepb"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/catpb"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/descpb"
	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgcode"
	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgerror"
	"github.com/cockroachdb/errors"
)

// minNumRegionsForSurviveRegionGoal is the minimum number of regions that a
// a database must have to survive a REGION failure.
const minNumRegionsForSurviveRegionGoal = 3

// RegionConfig represents the user configured state of a multi-region database.
// RegionConfig is intended to be a READ-ONLY struct and as such all members
// are private. Any modifications to the underlying type desc / db desc that
// inform be a RegionConfig must be made directly on those structs and a new
// RegionConfig must be synthesized to pick up those changes.
type RegionConfig struct {
	survivalGoal         descpb.SurvivalGoal
	regions              catpb.RegionNames
	transitioningRegions catpb.RegionNames
	primaryRegion        catpb.RegionName
	regionEnumID         descpb.ID
	placement            descpb.DataPlacement
	superRegions         []descpb.SuperRegion
	zoneCfgExtensions    descpb.ZoneConfigExtensions
	secondaryRegion      catpb.RegionName
}

// SurvivalGoal returns the survival goal configured on the RegionConfig.
func (r *RegionConfig) SurvivalGoal() descpb.SurvivalGoal {
	return r.survivalGoal
}

// PrimaryRegion returns the primary region configured on the RegionConfig.
func (r *RegionConfig) PrimaryRegion() catpb.RegionName {
	return r.primaryRegion
}

// Regions returns the list of regions added to the RegionConfig.
func (r *RegionConfig) Regions() catpb.RegionNames {
	return r.regions
}

// WithRegions returns a copy of the RegionConfig with only the provided
// regions.
func (r *RegionConfig) WithRegions(regions catpb.RegionNames) RegionConfig {
	cpy := *r
	cpy.regions = regions
	return cpy
}

// IsMemberOfExplicitSuperRegion returns whether t the region is an explicit
// member of a super region.
func (r *RegionConfig) IsMemberOfExplicitSuperRegion(region catpb.RegionName) bool {
	for _, superRegion := range r.SuperRegions() {
		for _, regionOfSuperRegion := range superRegion.Regions {
			if region == regionOfSuperRegion {
				return true
			}
		}
	}
	return false
}

// GetSuperRegionRegionsForRegion returns the members of the super region the
// specified region is part of.
// If the region is not a member of any super regions, the function returns an
// error.
func (r *RegionConfig) GetSuperRegionRegionsForRegion(
	region catpb.RegionName,
) (catpb.RegionNames, bool) {
	for _, superRegion := range r.SuperRegions() {
		for _, regionOfSuperRegion := range superRegion.Regions {
			if region == regionOfSuperRegion {
				return superRegion.Regions, true
			}
		}
	}
	return nil, false
}

// IsValidRegionNameString implements the tree.DatabaseRegionConfig interface.
func (r RegionConfig) IsValidRegionNameString(regionStr string) bool {
	for _, region := range r.Regions() {
		if string(region) == regionStr {
			return true
		}
	}
	return false
}

// PrimaryRegionString implements the tree.DatabaseRegionConfig interface.
func (r RegionConfig) PrimaryRegionString() string {
	return string(r.PrimaryRegion())
}

// TransitioningRegions returns all the regions which are currently transitioning
// from or to being PUBLIC.
func (r RegionConfig) TransitioningRegions() catpb.RegionNames {
	return r.transitioningRegions
}

// RegionEnumID returns the multi-region enum ID.
func (r *RegionConfig) RegionEnumID() descpb.ID {
	return r.regionEnumID
}

// Placement returns the data placement strategy for the region config.
func (r *RegionConfig) Placement() descpb.DataPlacement {
	return r.placement
}

// IsPlacementRestricted returns true if the database is in restricted
// placement, false otherwise.
func (r *RegionConfig) IsPlacementRestricted() bool {
	return r.placement == descpb.DataPlacement_RESTRICTED
}

// WithPlacementDefault returns a copy of the RegionConfig with the data
// placement strategy configured as placement default.
func (r *RegionConfig) WithPlacementDefault() RegionConfig {
	cpy := *r
	cpy.placement = descpb.DataPlacement_DEFAULT
	return cpy
}

// SuperRegions returns the list of super regions in the database.
func (r *RegionConfig) SuperRegions() []descpb.SuperRegion {
	return r.superRegions
}

// ZoneConfigExtensions returns the zone configuration extensions applied to the
// database.
func (r *RegionConfig) ZoneConfigExtensions() descpb.ZoneConfigExtensions {
	return r.zoneCfgExtensions
}

// ApplyZoneConfigExtensionForGlobal applies the global table zone configuration
// extensions to the provided zone configuration, returning the updated config.
func (r *RegionConfig) ApplyZoneConfigExtensionForGlobal(zc zonepb.ZoneConfig) zonepb.ZoneConfig {
	if ext := r.zoneCfgExtensions.Global; ext != nil {
		zc = extendZoneCfg(zc, *ext)
	}
	return zc
}

// ApplyZoneConfigExtensionForRegionalIn applies the regional table zone
// configuration extensions for the provided region to the provided zone
// configuration, returning the updated config.
func (r *RegionConfig) ApplyZoneConfigExtensionForRegionalIn(
	zc zonepb.ZoneConfig, region catpb.RegionName,
) zonepb.ZoneConfig {
	if ext := r.zoneCfgExtensions.Regional; ext != nil {
		zc = extendZoneCfg(zc, *ext)
	}
	if ext, ok := r.zoneCfgExtensions.RegionalIn[region]; ok {
		zc = extendZoneCfg(zc, ext)
	}
	return zc
}

// "extending" a zone config means having the extension inherit any missing
// fields from the zone config while replacing any set fields.
func extendZoneCfg(zc, ext zonepb.ZoneConfig) zonepb.ZoneConfig {
	ext.InheritFromParent(&zc)
	return ext
}

// GlobalTablesInheritDatabaseConstraints returns whether GLOBAL tables can
// inherit replica constraints from their database zone configuration, or
// whether they must set these constraints themselves.
func (r *RegionConfig) GlobalTablesInheritDatabaseConstraints() bool {
	if r.placement == descpb.DataPlacement_RESTRICTED {
		// Placement restricted does not apply to GLOBAL tables.
		return false
	}
	if r.zoneCfgExtensions.Global != nil {
		// Global tables have a zone config extension that will not be set at
		// the database level.
		return false
	}
	if r.zoneCfgExtensions.Regional != nil {
		// Regional tables have a zone config extension that will be set at the
		// database level but which should not apply to GLOBAL tables.
		return false
	}
	if _, ok := r.zoneCfgExtensions.RegionalIn[r.primaryRegion]; ok {
		// Regional tables in the primary region have a zone config extension that
		// will be set at the database level but which should not apply to GLOBAL
		// tables.
		return false
	}
	return true
}

// RegionalInTablesInheritDatabaseConstraints returns whether REGIONAL
// tables/partitions with affinity to the specified region can inherit replica
// constraints from their database zone configuration, or whether they must set
// these constraints themselves.
func (r *RegionConfig) RegionalInTablesInheritDatabaseConstraints(region catpb.RegionName) bool {
	if _, ok := r.zoneCfgExtensions.RegionalIn[r.primaryRegion]; ok {
		// Regional tables in the primary region have a zone config extension that
		// will be set at the database level but which should not apply to regional
		// tables in any other region.
		return r.primaryRegion == region
	}
	return true
}

// SecondaryRegion returns the secondary region configured on the RegionConfig.
func (r *RegionConfig) SecondaryRegion() catpb.RegionName {
	return r.secondaryRegion
}

// HasSecondaryRegion returns whether the RegionConfig has a secondary
// region set.
func (r *RegionConfig) HasSecondaryRegion() bool {
	return r.secondaryRegion != ""
}

// MakeRegionConfigOption is an option for MakeRegionConfig
type MakeRegionConfigOption func(r *RegionConfig)

// WithTransitioningRegions is an option to include transitioning
// regions into MakeRegionConfig.
func WithTransitioningRegions(transitioningRegions catpb.RegionNames) MakeRegionConfigOption {
	return func(r *RegionConfig) {
		r.transitioningRegions = transitioningRegions
	}
}

// WithSecondaryRegion is an option to include a secondary region.
func WithSecondaryRegion(secondaryRegion catpb.RegionName) MakeRegionConfigOption {
	return func(r *RegionConfig) {
		r.secondaryRegion = secondaryRegion
	}
}

// MakeRegionConfig constructs a RegionConfig.
func MakeRegionConfig(
	regions catpb.RegionNames,
	primaryRegion catpb.RegionName,
	survivalGoal descpb.SurvivalGoal,
	regionEnumID descpb.ID,
	placement descpb.DataPlacement,
	superRegions []descpb.SuperRegion,
	zoneCfgExtensions descpb.ZoneConfigExtensions,
	opts ...MakeRegionConfigOption,
) RegionConfig {
	ret := RegionConfig{
		regions:           regions,
		primaryRegion:     primaryRegion,
		survivalGoal:      survivalGoal,
		regionEnumID:      regionEnumID,
		placement:         placement,
		superRegions:      superRegions,
		zoneCfgExtensions: zoneCfgExtensions,
	}
	for _, opt := range opts {
		opt(&ret)
	}
	return ret
}

// CanSatisfySurvivalGoal returns true if the survival goal is satisfiable by
// the given region config.
func CanSatisfySurvivalGoal(survivalGoal descpb.SurvivalGoal, numRegions int) error {
	if survivalGoal == descpb.SurvivalGoal_REGION_FAILURE {
		if numRegions < minNumRegionsForSurviveRegionGoal {
			return errors.WithHintf(
				pgerror.Newf(
					pgcode.InvalidParameterValue,
					"at least %d regions are required for surviving a region failure",
					minNumRegionsForSurviveRegionGoal,
				),
				"you must add additional regions to the database or "+
					"change the survivability goal",
			)
		}
	}
	return nil
}

// ValidateRegionConfig validates that the given RegionConfig is valid.
func ValidateRegionConfig(config RegionConfig) error {
	if config.regionEnumID == descpb.InvalidID {
		return errors.AssertionFailedf("expected a valid multi-region enum ID to be initialized")
	}
	if len(config.regions) == 0 {
		return errors.AssertionFailedf("expected > 0 number of regions in the region config")
	}
	if config.placement == descpb.DataPlacement_RESTRICTED &&
		config.survivalGoal == descpb.SurvivalGoal_REGION_FAILURE {
		return errors.AssertionFailedf(
			"cannot have a database with restricted placement that is also region survivable")
	}

	var err error
	ValidateSuperRegions(config.SuperRegions(), config.SurvivalGoal(), config.Regions(), func(validateErr error) {
		if err == nil {
			err = validateErr
		}
	})
	if err != nil {
		return err
	}

	ValidateZoneConfigExtensions(config.Regions(), config.ZoneConfigExtensions(), func(validateErr error) {
		if err == nil {
			err = validateErr
		}
	})
	if err != nil {
		return err
	}

	return CanSatisfySurvivalGoal(config.survivalGoal, len(config.regions))
}

// ValidateSuperRegions validates that:
//   1. Region names are unique within a super region and are sorted.
//   2. All region within a super region map to a region on the RegionConfig.
//   3. Super region names are unique.
//   4. Each region can only belong to one super region.
func ValidateSuperRegions(
	superRegions []descpb.SuperRegion,
	survivalGoal descpb.SurvivalGoal,
	regionNames catpb.RegionNames,
	errorHandler func(error),
) {
	seenRegions := make(map[catpb.RegionName]struct{})
	superRegionNames := make(map[string]struct{})

	// Ensure that the super region names are in sorted order.
	if !sort.SliceIsSorted(superRegions, func(i, j int) bool {
		return superRegions[i].SuperRegionName < superRegions[j].SuperRegionName
	}) {
		err := errors.AssertionFailedf("super regions are not in sorted order based on the super region name %v", superRegions)
		errorHandler(err)
	}

	for _, superRegion := range superRegions {
		if len(superRegion.Regions) == 0 {
			err := errors.AssertionFailedf("no regions found within super region %s", superRegion.SuperRegionName)
			errorHandler(err)
		}

		if err := CanSatisfySurvivalGoal(survivalGoal, len(superRegion.Regions)); err != nil {
			err := errors.HandleAsAssertionFailure(errors.Wrapf(err, "super region %s only has %d regions", superRegion.SuperRegionName, len(superRegion.Regions)))
			errorHandler(err)
		}

		_, found := superRegionNames[superRegion.SuperRegionName]
		if found {
			err := errors.AssertionFailedf("duplicate super regions with name %s found", superRegion.SuperRegionName)
			errorHandler(err)
		}
		superRegionNames[superRegion.SuperRegionName] = struct{}{}

		// Ensure that regions within a super region are sorted.
		if !sort.SliceIsSorted(superRegion.Regions, func(i, j int) bool {
			return superRegion.Regions[i] < superRegion.Regions[j]
		}) {
			err := errors.AssertionFailedf("the regions within super region %s were not in a sorted order", superRegion.SuperRegionName)
			errorHandler(err)
		}

		seenRegionsInCurrentSuperRegion := make(map[catpb.RegionName]struct{})
		for _, region := range superRegion.Regions {
			_, found := seenRegionsInCurrentSuperRegion[region]
			if found {
				err := errors.AssertionFailedf("duplicate region %s found in super region %s", region, superRegion.SuperRegionName)
				errorHandler(err)
				continue
			}
			seenRegionsInCurrentSuperRegion[region] = struct{}{}
			_, found = seenRegions[region]
			if found {
				err := errors.AssertionFailedf("region %s found in multiple super regions", region)
				errorHandler(err)
				continue
			}
			seenRegions[region] = struct{}{}

			// Ensure that the region actually maps to a region on the regionConfig.
			found = false
			for _, regionName := range regionNames {
				if region == regionName {
					found = true
				}
			}
			if !found {
				err := errors.Newf("region %s not part of database", region)
				errorHandler(err)
			}
		}
	}
}

// ValidateZoneConfigExtensions validates that zone configuration extensions are
// coherent with the rest of the multi-region configuration. It validates that:
//   1. All per-region extensions map to a region on the RegionConfig.
//   2. TODO(nvanbenschoten): add more zone config extension validation in the
//      future to ensure zone config extensions do not subvert other portions
//      of the multi-region config (e.g. like breaking REGION survivability).
func ValidateZoneConfigExtensions(
	regionNames catpb.RegionNames,
	zoneCfgExtensions descpb.ZoneConfigExtensions,
	errorHandler func(error),
) {
	// Ensure that all per-region extensions map to a region on the RegionConfig.
	for regionExt := range zoneCfgExtensions.RegionalIn {
		found := false
		for _, regionInDB := range regionNames {
			if regionExt == regionInDB {
				found = true
				break
			}
		}
		if !found {
			errorHandler(errors.AssertionFailedf("region %s has REGIONAL IN "+
				"zone config extension, but is not a region in the database", regionExt))
		}
	}
}

// IsMemberOfSuperRegion returns a boolean representing if the region is part
// of a super region and the name of the super region.
func IsMemberOfSuperRegion(name catpb.RegionName, config RegionConfig) (bool, string) {
	for _, superRegion := range config.SuperRegions() {
		for _, region := range superRegion.Regions {
			if region == name {
				return true, superRegion.SuperRegionName
			}
		}
	}

	return false, ""
}

// CanDropRegion returns an error if the survival goal doesn't allow for
// removing regions or if the region is part of a super region.
func CanDropRegion(name catpb.RegionName, config RegionConfig) error {
	isMember, superRegion := IsMemberOfSuperRegion(name, config)
	if isMember {
		return errors.WithHintf(
			pgerror.Newf(pgcode.DependentObjectsStillExist, "region %s is part of super region %s", name, superRegion),
			"you must first drop super region %s before you can drop the region %s", superRegion, name,
		)
	}
	return CanSatisfySurvivalGoal(config.survivalGoal, len(config.regions)-1)
}
