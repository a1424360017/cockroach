// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Code generated by generate-logictest, DO NOT EDIT.

package testlocal

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/build/bazel"
	"github.com/cockroachdb/cockroach/pkg/security/securityassets"
	"github.com/cockroachdb/cockroach/pkg/security/securitytest"
	"github.com/cockroachdb/cockroach/pkg/server"
	"github.com/cockroachdb/cockroach/pkg/sql/logictest"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlitelogictest"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/cockroachdb/cockroach/pkg/testutils/testcluster"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/randutil"
)

const configIdx = 0

var sqliteLogicTestDir string

func init() {
}

func TestMain(m *testing.M) {
	flag.Parse()
	if *logictest.Bigtest {
		if bazel.BuiltWithBazel() {
			var err error
			sqliteLogicTestDir, err = bazel.Runfile("external/com_github_cockroachdb_sqllogictest")
			if err != nil {
				panic(err)
			}
		} else {
			var err error
			sqliteLogicTestDir, err = sqlitelogictest.FindLocalLogicTestClone()
			if err != nil {
				panic(err)
			}
		}
	}
	securityassets.SetLoader(securitytest.EmbeddedAssets)
	randutil.SeedForTests()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	serverutils.InitTestClusterFactory(testcluster.TestClusterFactory)
	os.Exit(m.Run())
}

func runSqliteLogicTest(t *testing.T, file string) {
	skip.UnderDeadlock(t, "times out and/or hangs")
	if !*logictest.Bigtest {
		skip.IgnoreLint(t, "-bigtest flag must be specified to run this test")
	}
	// SQLLite logic tests can be very memory intensive, so we give them larger
	// limit than other logic tests get.
	serverArgs := logictest.TestServerArgs{
		MaxSQLMemoryLimit: 512 << 20, // 512 MiB
		// TODO(yuzefovich): remove this once the flake in #84022 is fixed.
		DisableWorkmemRandomization: true,
	}
	logictest.RunLogicTest(t, serverArgs, configIdx, filepath.Join(sqliteLogicTestDir, file))
}

func TestSqlLiteLogic_testindexbetween1slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/1/slt_good_0.test")
}

func TestSqlLiteLogic_testindexbetween10slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/10/slt_good_0.test")
}

func TestSqlLiteLogic_testindexbetween10slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/10/slt_good_1.test")
}

func TestSqlLiteLogic_testindexbetween10slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/10/slt_good_2.test")
}

func TestSqlLiteLogic_testindexbetween10slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/10/slt_good_3.test")
}

func TestSqlLiteLogic_testindexbetween10slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/10/slt_good_4.test")
}

func TestSqlLiteLogic_testindexbetween10slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/10/slt_good_5.test")
}

func TestSqlLiteLogic_testindexbetween100slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/100/slt_good_0.test")
}

func TestSqlLiteLogic_testindexbetween100slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/100/slt_good_1.test")
}

func TestSqlLiteLogic_testindexbetween100slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/100/slt_good_2.test")
}

func TestSqlLiteLogic_testindexbetween100slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/100/slt_good_3.test")
}

func TestSqlLiteLogic_testindexbetween100slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/100/slt_good_4.test")
}

func TestSqlLiteLogic_testindexbetween1000slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/between/1000/slt_good_0.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_0.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_1.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_10_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_10.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_11_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_11.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_12_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_12.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_13_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_13.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_14_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_14.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_15_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_15.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_16_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_16.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_17_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_17.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_18_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_18.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_19_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_19.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_2.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_20_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_20.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_21_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_21.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_22_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_22.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_23_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_23.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_24_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_24.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_25_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_25.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_26_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_26.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_27_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_27.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_28_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_28.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_29_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_29.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_3.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_30_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_30.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_31_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_31.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_32_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_32.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_33_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_33.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_34_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_34.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_4.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_5.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_6_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_6.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_7_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_7.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_8_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_8.test")
}

func TestSqlLiteLogic_testindexcommute10slt_good_9_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/10/slt_good_9.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_0.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_1.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_10_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_10.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_11_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_11.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_12_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_12.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_2.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_3.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_4.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_5.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_6_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_6.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_7_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_7.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_8_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_8.test")
}

func TestSqlLiteLogic_testindexcommute100slt_good_9_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/100/slt_good_9.test")
}

func TestSqlLiteLogic_testindexcommute1000slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/1000/slt_good_0.test")
}

func TestSqlLiteLogic_testindexcommute1000slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/1000/slt_good_1.test")
}

func TestSqlLiteLogic_testindexcommute1000slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/1000/slt_good_2.test")
}

func TestSqlLiteLogic_testindexcommute1000slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/commute/1000/slt_good_3.test")
}

func TestSqlLiteLogic_testindexdelete1slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/1/slt_good_0.test")
}

func TestSqlLiteLogic_testindexdelete10slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/10/slt_good_0.test")
}

func TestSqlLiteLogic_testindexdelete10slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/10/slt_good_1.test")
}

func TestSqlLiteLogic_testindexdelete10slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/10/slt_good_2.test")
}

func TestSqlLiteLogic_testindexdelete10slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/10/slt_good_3.test")
}

func TestSqlLiteLogic_testindexdelete10slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/10/slt_good_4.test")
}

func TestSqlLiteLogic_testindexdelete10slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/10/slt_good_5.test")
}

func TestSqlLiteLogic_testindexdelete100slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/100/slt_good_0.test")
}

func TestSqlLiteLogic_testindexdelete100slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/100/slt_good_1.test")
}

func TestSqlLiteLogic_testindexdelete100slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/100/slt_good_2.test")
}

func TestSqlLiteLogic_testindexdelete100slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/100/slt_good_3.test")
}

func TestSqlLiteLogic_testindexdelete1000slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/1000/slt_good_0.test")
}

func TestSqlLiteLogic_testindexdelete1000slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/1000/slt_good_1.test")
}

func TestSqlLiteLogic_testindexdelete10000slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/delete/10000/slt_good_0.test")
}

func TestSqlLiteLogic_testindexin10slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/10/slt_good_0.test")
}

func TestSqlLiteLogic_testindexin10slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/10/slt_good_1.test")
}

func TestSqlLiteLogic_testindexin10slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/10/slt_good_2.test")
}

func TestSqlLiteLogic_testindexin10slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/10/slt_good_3.test")
}

func TestSqlLiteLogic_testindexin10slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/10/slt_good_4.test")
}

func TestSqlLiteLogic_testindexin10slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/10/slt_good_5.test")
}

func TestSqlLiteLogic_testindexin100slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/100/slt_good_0.test")
}

func TestSqlLiteLogic_testindexin100slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/100/slt_good_1.test")
}

func TestSqlLiteLogic_testindexin100slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/100/slt_good_2.test")
}

func TestSqlLiteLogic_testindexin100slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/100/slt_good_3.test")
}

func TestSqlLiteLogic_testindexin100slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/100/slt_good_4.test")
}

func TestSqlLiteLogic_testindexin1000slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/1000/slt_good_0.test")
}

func TestSqlLiteLogic_testindexin1000slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/in/1000/slt_good_1.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_0.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_1.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_10_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_10.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_11_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_11.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_12_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_12.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_13_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_13.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_14_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_14.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_15_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_15.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_16_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_16.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_17_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_17.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_18_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_18.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_19_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_19.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_2.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_20_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_20.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_21_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_21.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_22_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_22.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_23_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_23.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_24_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_24.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_25_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_25.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_3.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_4.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_5.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_6_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_6.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_7_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_7.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_8_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_8.test")
}

func TestSqlLiteLogic_testindexorderby10slt_good_9_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/10/slt_good_9.test")
}

func TestSqlLiteLogic_testindexorderby100slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/100/slt_good_0.test")
}

func TestSqlLiteLogic_testindexorderby100slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/100/slt_good_1.test")
}

func TestSqlLiteLogic_testindexorderby100slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/100/slt_good_2.test")
}

func TestSqlLiteLogic_testindexorderby100slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/100/slt_good_3.test")
}

func TestSqlLiteLogic_testindexorderby1000slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby/1000/slt_good_0.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_0.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_1.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_10_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_10.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_11_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_11.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_12_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_12.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_13_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_13.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_14_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_14.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_15_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_15.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_16_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_16.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_17_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_17.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_18_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_18.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_19_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_19.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_2.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_20_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_20.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_21_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_21.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_22_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_22.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_23_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_23.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_24_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_24.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_25_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_25.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_26_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_26.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_27_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_27.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_28_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_28.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_29_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_29.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_3.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_30_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_30.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_31_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_31.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_32_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_32.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_33_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_33.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_34_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_34.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_35_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_35.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_36_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_36.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_37_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_37.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_38_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_38.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_39_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_39.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_4.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_5.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_6_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_6.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_7_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_7.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_8_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_8.test")
}

func TestSqlLiteLogic_testindexorderby_nosort10slt_good_9_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/10/slt_good_9.test")
}

func TestSqlLiteLogic_testindexorderby_nosort100slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/100/slt_good_0.test")
}

func TestSqlLiteLogic_testindexorderby_nosort100slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/100/slt_good_1.test")
}

func TestSqlLiteLogic_testindexorderby_nosort100slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/100/slt_good_2.test")
}

func TestSqlLiteLogic_testindexorderby_nosort100slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/100/slt_good_3.test")
}

func TestSqlLiteLogic_testindexorderby_nosort100slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/100/slt_good_4.test")
}

func TestSqlLiteLogic_testindexorderby_nosort100slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/100/slt_good_5.test")
}

func TestSqlLiteLogic_testindexorderby_nosort100slt_good_6_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/100/slt_good_6.test")
}

func TestSqlLiteLogic_testindexorderby_nosort1000slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/1000/slt_good_0.test")
}

func TestSqlLiteLogic_testindexorderby_nosort1000slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/orderby_nosort/1000/slt_good_1.test")
}

func TestSqlLiteLogic_testindexview10slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/10/slt_good_0.test")
}

func TestSqlLiteLogic_testindexview10slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/10/slt_good_1.test")
}

func TestSqlLiteLogic_testindexview10slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/10/slt_good_2.test")
}

func TestSqlLiteLogic_testindexview10slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/10/slt_good_3.test")
}

func TestSqlLiteLogic_testindexview10slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/10/slt_good_4.test")
}

func TestSqlLiteLogic_testindexview10slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/10/slt_good_5.test")
}

func TestSqlLiteLogic_testindexview10slt_good_6_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/10/slt_good_6.test")
}

func TestSqlLiteLogic_testindexview10slt_good_7_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/10/slt_good_7.test")
}

func TestSqlLiteLogic_testindexview100slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/100/slt_good_0.test")
}

func TestSqlLiteLogic_testindexview100slt_good_1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/100/slt_good_1.test")
}

func TestSqlLiteLogic_testindexview100slt_good_2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/100/slt_good_2.test")
}

func TestSqlLiteLogic_testindexview100slt_good_3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/100/slt_good_3.test")
}

func TestSqlLiteLogic_testindexview100slt_good_4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/100/slt_good_4.test")
}

func TestSqlLiteLogic_testindexview100slt_good_5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/100/slt_good_5.test")
}

func TestSqlLiteLogic_testindexview1000slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/1000/slt_good_0.test")
}

func TestSqlLiteLogic_testindexview10000slt_good_0_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/index/view/10000/slt_good_0.test")
}

func TestSqlLiteLogic_testselect1_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/select1.test")
}

func TestSqlLiteLogic_testselect2_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/select2.test")
}

func TestSqlLiteLogic_testselect3_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/select3.test")
}

func TestSqlLiteLogic_testselect4_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/select4.test")
}

func TestSqlLiteLogic_testselect5_test(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runSqliteLogicTest(t, "/test/select5.test")
}
