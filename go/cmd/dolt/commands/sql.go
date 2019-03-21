package commands

import (
	"github.com/liquidata-inc/ld/dolt/go/cmd/dolt/cli"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/env"
	"github.com/liquidata-inc/ld/dolt/go/libraries/utils/argparser"
)

var sqlShortDesc = "EXPERIMENTAL: Runs a SQL query"
var sqlLongDesc = "EXPERIMENTAL: Runs a SQL query you specify. By default, begins an interactive session to run " +
	"queries and view the results. With the -q option, runs the given query and prints any results, then exits."
var sqlSynopsis = []string{
	"[options] -q query_string",
	"[options]",
}

const (
	queryFlag = "query"
)


func Sql(commandStr string, args []string, dEnv *env.DoltEnv) int {
	ap := argparser.NewArgParser()
	ap.SupportsFlag(queryFlag, "q", "run a single query and exit")
	help, usage := cli.HelpAndUsagePrinters(commandStr, sqlShortDesc, sqlLongDesc, sqlSynopsis, ap)

	apr := cli.ParseArgs(ap, args, help)
	args = apr.Args()

	if len(args) == 0 {
		usage()
		return 1
	}

	return 1
}