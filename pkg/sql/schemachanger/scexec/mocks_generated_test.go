// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scexec (interfaces: Catalog,Dependencies,Backfiller,Merger,BackfillerTracker,IndexSpanSplitter,PeriodicProgressFlusher)

// Package scexec_test is a generated GoMock package.
package scexec_test

import (
	context "context"
	reflect "reflect"

	username "github.com/cockroachdb/cockroach/pkg/security/username"
	catalog "github.com/cockroachdb/cockroach/pkg/sql/catalog"
	scexec "github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scexec"
	scmutationexec "github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scexec/scmutationexec"
	catid "github.com/cockroachdb/cockroach/pkg/sql/sem/catid"
	gomock "github.com/golang/mock/gomock"
)

// MockCatalog is a mock of Catalog interface.
type MockCatalog struct {
	ctrl     *gomock.Controller
	recorder *MockCatalogMockRecorder
}

// MockCatalogMockRecorder is the mock recorder for MockCatalog.
type MockCatalogMockRecorder struct {
	mock *MockCatalog
}

// NewMockCatalog creates a new mock instance.
func NewMockCatalog(ctrl *gomock.Controller) *MockCatalog {
	mock := &MockCatalog{ctrl: ctrl}
	mock.recorder = &MockCatalogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCatalog) EXPECT() *MockCatalogMockRecorder {
	return m.recorder
}

// GetFullyQualifiedName mocks base method.
func (m *MockCatalog) GetFullyQualifiedName(arg0 context.Context, arg1 catid.DescID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullyQualifiedName", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullyQualifiedName indicates an expected call of GetFullyQualifiedName.
func (mr *MockCatalogMockRecorder) GetFullyQualifiedName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullyQualifiedName", reflect.TypeOf((*MockCatalog)(nil).GetFullyQualifiedName), arg0, arg1)
}

// MustReadImmutableDescriptors mocks base method.
func (m *MockCatalog) MustReadImmutableDescriptors(arg0 context.Context, arg1 ...catid.DescID) ([]catalog.Descriptor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MustReadImmutableDescriptors", varargs...)
	ret0, _ := ret[0].([]catalog.Descriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MustReadImmutableDescriptors indicates an expected call of MustReadImmutableDescriptors.
func (mr *MockCatalogMockRecorder) MustReadImmutableDescriptors(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustReadImmutableDescriptors", reflect.TypeOf((*MockCatalog)(nil).MustReadImmutableDescriptors), varargs...)
}

// MustReadMutableDescriptor mocks base method.
func (m *MockCatalog) MustReadMutableDescriptor(arg0 context.Context, arg1 catid.DescID) (catalog.MutableDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustReadMutableDescriptor", arg0, arg1)
	ret0, _ := ret[0].(catalog.MutableDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MustReadMutableDescriptor indicates an expected call of MustReadMutableDescriptor.
func (mr *MockCatalogMockRecorder) MustReadMutableDescriptor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustReadMutableDescriptor", reflect.TypeOf((*MockCatalog)(nil).MustReadMutableDescriptor), arg0, arg1)
}

// NewCatalogChangeBatcher mocks base method.
func (m *MockCatalog) NewCatalogChangeBatcher() scexec.CatalogChangeBatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCatalogChangeBatcher")
	ret0, _ := ret[0].(scexec.CatalogChangeBatcher)
	return ret0
}

// NewCatalogChangeBatcher indicates an expected call of NewCatalogChangeBatcher.
func (mr *MockCatalogMockRecorder) NewCatalogChangeBatcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCatalogChangeBatcher", reflect.TypeOf((*MockCatalog)(nil).NewCatalogChangeBatcher))
}

// MockDependencies is a mock of Dependencies interface.
type MockDependencies struct {
	ctrl     *gomock.Controller
	recorder *MockDependenciesMockRecorder
}

// MockDependenciesMockRecorder is the mock recorder for MockDependencies.
type MockDependenciesMockRecorder struct {
	mock *MockDependencies
}

// NewMockDependencies creates a new mock instance.
func NewMockDependencies(ctrl *gomock.Controller) *MockDependencies {
	mock := &MockDependencies{ctrl: ctrl}
	mock.recorder = &MockDependenciesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDependencies) EXPECT() *MockDependenciesMockRecorder {
	return m.recorder
}

// BackfillProgressTracker mocks base method.
func (m *MockDependencies) BackfillProgressTracker() scexec.BackfillerTracker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackfillProgressTracker")
	ret0, _ := ret[0].(scexec.BackfillerTracker)
	return ret0
}

// BackfillProgressTracker indicates an expected call of BackfillProgressTracker.
func (mr *MockDependenciesMockRecorder) BackfillProgressTracker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackfillProgressTracker", reflect.TypeOf((*MockDependencies)(nil).BackfillProgressTracker))
}

// Catalog mocks base method.
func (m *MockDependencies) Catalog() scexec.Catalog {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Catalog")
	ret0, _ := ret[0].(scexec.Catalog)
	return ret0
}

// Catalog indicates an expected call of Catalog.
func (mr *MockDependenciesMockRecorder) Catalog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Catalog", reflect.TypeOf((*MockDependencies)(nil).Catalog))
}

// Clock mocks base method.
func (m *MockDependencies) Clock() scmutationexec.Clock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clock")
	ret0, _ := ret[0].(scmutationexec.Clock)
	return ret0
}

// Clock indicates an expected call of Clock.
func (mr *MockDependenciesMockRecorder) Clock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clock", reflect.TypeOf((*MockDependencies)(nil).Clock))
}

// DescriptorMetadataUpdater mocks base method.
func (m *MockDependencies) DescriptorMetadataUpdater(arg0 context.Context) scexec.DescriptorMetadataUpdater {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescriptorMetadataUpdater", arg0)
	ret0, _ := ret[0].(scexec.DescriptorMetadataUpdater)
	return ret0
}

// DescriptorMetadataUpdater indicates an expected call of DescriptorMetadataUpdater.
func (mr *MockDependenciesMockRecorder) DescriptorMetadataUpdater(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorMetadataUpdater", reflect.TypeOf((*MockDependencies)(nil).DescriptorMetadataUpdater), arg0)
}

// EventLogger mocks base method.
func (m *MockDependencies) EventLogger() scexec.EventLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventLogger")
	ret0, _ := ret[0].(scexec.EventLogger)
	return ret0
}

// EventLogger indicates an expected call of EventLogger.
func (mr *MockDependenciesMockRecorder) EventLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventLogger", reflect.TypeOf((*MockDependencies)(nil).EventLogger))
}

// GetTestingKnobs mocks base method.
func (m *MockDependencies) GetTestingKnobs() *scexec.TestingKnobs {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestingKnobs")
	ret0, _ := ret[0].(*scexec.TestingKnobs)
	return ret0
}

// GetTestingKnobs indicates an expected call of GetTestingKnobs.
func (mr *MockDependenciesMockRecorder) GetTestingKnobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestingKnobs", reflect.TypeOf((*MockDependencies)(nil).GetTestingKnobs))
}

// IndexBackfiller mocks base method.
func (m *MockDependencies) IndexBackfiller() scexec.Backfiller {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexBackfiller")
	ret0, _ := ret[0].(scexec.Backfiller)
	return ret0
}

// IndexBackfiller indicates an expected call of IndexBackfiller.
func (mr *MockDependenciesMockRecorder) IndexBackfiller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexBackfiller", reflect.TypeOf((*MockDependencies)(nil).IndexBackfiller))
}

// IndexMerger mocks base method.
func (m *MockDependencies) IndexMerger() scexec.Merger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexMerger")
	ret0, _ := ret[0].(scexec.Merger)
	return ret0
}

// IndexMerger indicates an expected call of IndexMerger.
func (mr *MockDependenciesMockRecorder) IndexMerger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexMerger", reflect.TypeOf((*MockDependencies)(nil).IndexMerger))
}

// IndexSpanSplitter mocks base method.
func (m *MockDependencies) IndexSpanSplitter() scexec.IndexSpanSplitter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexSpanSplitter")
	ret0, _ := ret[0].(scexec.IndexSpanSplitter)
	return ret0
}

// IndexSpanSplitter indicates an expected call of IndexSpanSplitter.
func (mr *MockDependenciesMockRecorder) IndexSpanSplitter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexSpanSplitter", reflect.TypeOf((*MockDependencies)(nil).IndexSpanSplitter))
}

// IndexValidator mocks base method.
func (m *MockDependencies) IndexValidator() scexec.IndexValidator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexValidator")
	ret0, _ := ret[0].(scexec.IndexValidator)
	return ret0
}

// IndexValidator indicates an expected call of IndexValidator.
func (mr *MockDependenciesMockRecorder) IndexValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexValidator", reflect.TypeOf((*MockDependencies)(nil).IndexValidator))
}

// PeriodicProgressFlusher mocks base method.
func (m *MockDependencies) PeriodicProgressFlusher() scexec.PeriodicProgressFlusher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeriodicProgressFlusher")
	ret0, _ := ret[0].(scexec.PeriodicProgressFlusher)
	return ret0
}

// PeriodicProgressFlusher indicates an expected call of PeriodicProgressFlusher.
func (mr *MockDependenciesMockRecorder) PeriodicProgressFlusher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeriodicProgressFlusher", reflect.TypeOf((*MockDependencies)(nil).PeriodicProgressFlusher))
}

// Statements mocks base method.
func (m *MockDependencies) Statements() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Statements")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Statements indicates an expected call of Statements.
func (mr *MockDependenciesMockRecorder) Statements() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Statements", reflect.TypeOf((*MockDependencies)(nil).Statements))
}

// StatsRefresher mocks base method.
func (m *MockDependencies) StatsRefresher() scexec.StatsRefreshQueue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatsRefresher")
	ret0, _ := ret[0].(scexec.StatsRefreshQueue)
	return ret0
}

// StatsRefresher indicates an expected call of StatsRefresher.
func (mr *MockDependenciesMockRecorder) StatsRefresher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatsRefresher", reflect.TypeOf((*MockDependencies)(nil).StatsRefresher))
}

// Telemetry mocks base method.
func (m *MockDependencies) Telemetry() scexec.Telemetry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Telemetry")
	ret0, _ := ret[0].(scexec.Telemetry)
	return ret0
}

// Telemetry indicates an expected call of Telemetry.
func (mr *MockDependenciesMockRecorder) Telemetry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Telemetry", reflect.TypeOf((*MockDependencies)(nil).Telemetry))
}

// TransactionalJobRegistry mocks base method.
func (m *MockDependencies) TransactionalJobRegistry() scexec.TransactionalJobRegistry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionalJobRegistry")
	ret0, _ := ret[0].(scexec.TransactionalJobRegistry)
	return ret0
}

// TransactionalJobRegistry indicates an expected call of TransactionalJobRegistry.
func (mr *MockDependenciesMockRecorder) TransactionalJobRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionalJobRegistry", reflect.TypeOf((*MockDependencies)(nil).TransactionalJobRegistry))
}

// User mocks base method.
func (m *MockDependencies) User() username.SQLUsername {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(username.SQLUsername)
	return ret0
}

// User indicates an expected call of User.
func (mr *MockDependenciesMockRecorder) User() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockDependencies)(nil).User))
}

// MockBackfiller is a mock of Backfiller interface.
type MockBackfiller struct {
	ctrl     *gomock.Controller
	recorder *MockBackfillerMockRecorder
}

// MockBackfillerMockRecorder is the mock recorder for MockBackfiller.
type MockBackfillerMockRecorder struct {
	mock *MockBackfiller
}

// NewMockBackfiller creates a new mock instance.
func NewMockBackfiller(ctrl *gomock.Controller) *MockBackfiller {
	mock := &MockBackfiller{ctrl: ctrl}
	mock.recorder = &MockBackfillerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackfiller) EXPECT() *MockBackfillerMockRecorder {
	return m.recorder
}

// BackfillIndexes mocks base method.
func (m *MockBackfiller) BackfillIndexes(arg0 context.Context, arg1 scexec.BackfillProgress, arg2 scexec.BackfillerProgressWriter, arg3 catalog.TableDescriptor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackfillIndexes", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// BackfillIndexes indicates an expected call of BackfillIndexes.
func (mr *MockBackfillerMockRecorder) BackfillIndexes(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackfillIndexes", reflect.TypeOf((*MockBackfiller)(nil).BackfillIndexes), arg0, arg1, arg2, arg3)
}

// MaybePrepareDestIndexesForBackfill mocks base method.
func (m *MockBackfiller) MaybePrepareDestIndexesForBackfill(arg0 context.Context, arg1 scexec.BackfillProgress, arg2 catalog.TableDescriptor) (scexec.BackfillProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybePrepareDestIndexesForBackfill", arg0, arg1, arg2)
	ret0, _ := ret[0].(scexec.BackfillProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybePrepareDestIndexesForBackfill indicates an expected call of MaybePrepareDestIndexesForBackfill.
func (mr *MockBackfillerMockRecorder) MaybePrepareDestIndexesForBackfill(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybePrepareDestIndexesForBackfill", reflect.TypeOf((*MockBackfiller)(nil).MaybePrepareDestIndexesForBackfill), arg0, arg1, arg2)
}

// MockMerger is a mock of Merger interface.
type MockMerger struct {
	ctrl     *gomock.Controller
	recorder *MockMergerMockRecorder
}

// MockMergerMockRecorder is the mock recorder for MockMerger.
type MockMergerMockRecorder struct {
	mock *MockMerger
}

// NewMockMerger creates a new mock instance.
func NewMockMerger(ctrl *gomock.Controller) *MockMerger {
	mock := &MockMerger{ctrl: ctrl}
	mock.recorder = &MockMergerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMerger) EXPECT() *MockMergerMockRecorder {
	return m.recorder
}

// MergeIndexes mocks base method.
func (m *MockMerger) MergeIndexes(arg0 context.Context, arg1 scexec.MergeProgress, arg2 scexec.BackfillerProgressWriter, arg3 catalog.TableDescriptor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeIndexes", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// MergeIndexes indicates an expected call of MergeIndexes.
func (mr *MockMergerMockRecorder) MergeIndexes(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeIndexes", reflect.TypeOf((*MockMerger)(nil).MergeIndexes), arg0, arg1, arg2, arg3)
}

// MockBackfillerTracker is a mock of BackfillerTracker interface.
type MockBackfillerTracker struct {
	ctrl     *gomock.Controller
	recorder *MockBackfillerTrackerMockRecorder
}

// MockBackfillerTrackerMockRecorder is the mock recorder for MockBackfillerTracker.
type MockBackfillerTrackerMockRecorder struct {
	mock *MockBackfillerTracker
}

// NewMockBackfillerTracker creates a new mock instance.
func NewMockBackfillerTracker(ctrl *gomock.Controller) *MockBackfillerTracker {
	mock := &MockBackfillerTracker{ctrl: ctrl}
	mock.recorder = &MockBackfillerTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackfillerTracker) EXPECT() *MockBackfillerTrackerMockRecorder {
	return m.recorder
}

// FlushCheckpoint mocks base method.
func (m *MockBackfillerTracker) FlushCheckpoint(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushCheckpoint", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushCheckpoint indicates an expected call of FlushCheckpoint.
func (mr *MockBackfillerTrackerMockRecorder) FlushCheckpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushCheckpoint", reflect.TypeOf((*MockBackfillerTracker)(nil).FlushCheckpoint), arg0)
}

// FlushFractionCompleted mocks base method.
func (m *MockBackfillerTracker) FlushFractionCompleted(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushFractionCompleted", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushFractionCompleted indicates an expected call of FlushFractionCompleted.
func (mr *MockBackfillerTrackerMockRecorder) FlushFractionCompleted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushFractionCompleted", reflect.TypeOf((*MockBackfillerTracker)(nil).FlushFractionCompleted), arg0)
}

// GetBackfillProgress mocks base method.
func (m *MockBackfillerTracker) GetBackfillProgress(arg0 context.Context, arg1 scexec.Backfill) (scexec.BackfillProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackfillProgress", arg0, arg1)
	ret0, _ := ret[0].(scexec.BackfillProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackfillProgress indicates an expected call of GetBackfillProgress.
func (mr *MockBackfillerTrackerMockRecorder) GetBackfillProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackfillProgress", reflect.TypeOf((*MockBackfillerTracker)(nil).GetBackfillProgress), arg0, arg1)
}

// GetMergeProgress mocks base method.
func (m *MockBackfillerTracker) GetMergeProgress(arg0 context.Context, arg1 scexec.Merge) (scexec.MergeProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMergeProgress", arg0, arg1)
	ret0, _ := ret[0].(scexec.MergeProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMergeProgress indicates an expected call of GetMergeProgress.
func (mr *MockBackfillerTrackerMockRecorder) GetMergeProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeProgress", reflect.TypeOf((*MockBackfillerTracker)(nil).GetMergeProgress), arg0, arg1)
}

// SetBackfillProgress mocks base method.
func (m *MockBackfillerTracker) SetBackfillProgress(arg0 context.Context, arg1 scexec.BackfillProgress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBackfillProgress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBackfillProgress indicates an expected call of SetBackfillProgress.
func (mr *MockBackfillerTrackerMockRecorder) SetBackfillProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBackfillProgress", reflect.TypeOf((*MockBackfillerTracker)(nil).SetBackfillProgress), arg0, arg1)
}

// SetMergeProgress mocks base method.
func (m *MockBackfillerTracker) SetMergeProgress(arg0 context.Context, arg1 scexec.MergeProgress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMergeProgress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMergeProgress indicates an expected call of SetMergeProgress.
func (mr *MockBackfillerTrackerMockRecorder) SetMergeProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMergeProgress", reflect.TypeOf((*MockBackfillerTracker)(nil).SetMergeProgress), arg0, arg1)
}

// MockIndexSpanSplitter is a mock of IndexSpanSplitter interface.
type MockIndexSpanSplitter struct {
	ctrl     *gomock.Controller
	recorder *MockIndexSpanSplitterMockRecorder
}

// MockIndexSpanSplitterMockRecorder is the mock recorder for MockIndexSpanSplitter.
type MockIndexSpanSplitterMockRecorder struct {
	mock *MockIndexSpanSplitter
}

// NewMockIndexSpanSplitter creates a new mock instance.
func NewMockIndexSpanSplitter(ctrl *gomock.Controller) *MockIndexSpanSplitter {
	mock := &MockIndexSpanSplitter{ctrl: ctrl}
	mock.recorder = &MockIndexSpanSplitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexSpanSplitter) EXPECT() *MockIndexSpanSplitterMockRecorder {
	return m.recorder
}

// MaybeSplitIndexSpans mocks base method.
func (m *MockIndexSpanSplitter) MaybeSplitIndexSpans(arg0 context.Context, arg1 catalog.TableDescriptor, arg2 catalog.Index) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybeSplitIndexSpans", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MaybeSplitIndexSpans indicates an expected call of MaybeSplitIndexSpans.
func (mr *MockIndexSpanSplitterMockRecorder) MaybeSplitIndexSpans(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeSplitIndexSpans", reflect.TypeOf((*MockIndexSpanSplitter)(nil).MaybeSplitIndexSpans), arg0, arg1, arg2)
}

// MockPeriodicProgressFlusher is a mock of PeriodicProgressFlusher interface.
type MockPeriodicProgressFlusher struct {
	ctrl     *gomock.Controller
	recorder *MockPeriodicProgressFlusherMockRecorder
}

// MockPeriodicProgressFlusherMockRecorder is the mock recorder for MockPeriodicProgressFlusher.
type MockPeriodicProgressFlusherMockRecorder struct {
	mock *MockPeriodicProgressFlusher
}

// NewMockPeriodicProgressFlusher creates a new mock instance.
func NewMockPeriodicProgressFlusher(ctrl *gomock.Controller) *MockPeriodicProgressFlusher {
	mock := &MockPeriodicProgressFlusher{ctrl: ctrl}
	mock.recorder = &MockPeriodicProgressFlusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeriodicProgressFlusher) EXPECT() *MockPeriodicProgressFlusherMockRecorder {
	return m.recorder
}

// StartPeriodicUpdates mocks base method.
func (m *MockPeriodicProgressFlusher) StartPeriodicUpdates(arg0 context.Context, arg1 scexec.BackfillerProgressFlusher) func() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartPeriodicUpdates", arg0, arg1)
	ret0, _ := ret[0].(func() error)
	return ret0
}

// StartPeriodicUpdates indicates an expected call of StartPeriodicUpdates.
func (mr *MockPeriodicProgressFlusherMockRecorder) StartPeriodicUpdates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPeriodicUpdates", reflect.TypeOf((*MockPeriodicProgressFlusher)(nil).StartPeriodicUpdates), arg0, arg1)
}