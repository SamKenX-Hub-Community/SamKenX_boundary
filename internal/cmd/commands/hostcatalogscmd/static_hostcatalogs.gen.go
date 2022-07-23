// Code generated by "make cli"; DO NOT EDIT.
package hostcatalogscmd

import (
	"errors"
	"fmt"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/hostcatalogs"
	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/hashicorp/boundary/internal/cmd/common"
	"github.com/hashicorp/go-secure-stdlib/strutil"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

func initStaticFlags() {
	flagsOnce.Do(func() {
		extraFlags := extraStaticActionsFlagsMapFunc()
		for k, v := range extraFlags {
			flagsStaticMap[k] = append(flagsStaticMap[k], v...)
		}
	})
}

var (
	_ cli.Command             = (*StaticCommand)(nil)
	_ cli.CommandAutocomplete = (*StaticCommand)(nil)
)

type StaticCommand struct {
	*base.Command

	Func string

	plural string
}

func (c *StaticCommand) AutocompleteArgs() complete.Predictor {
	initStaticFlags()
	return complete.PredictAnything
}

func (c *StaticCommand) AutocompleteFlags() complete.Flags {
	initStaticFlags()
	return c.Flags().Completions()
}

func (c *StaticCommand) Synopsis() string {
	if extra := extraStaticSynopsisFunc(c); extra != "" {
		return extra
	}

	synopsisStr := "host catalog"

	synopsisStr = fmt.Sprintf("%s %s", "static-type", synopsisStr)

	return common.SynopsisFunc(c.Func, synopsisStr)
}

func (c *StaticCommand) Help() string {
	initStaticFlags()

	var helpStr string
	helpMap := common.HelpMap("host catalog")

	switch c.Func {

	default:

		helpStr = c.extraStaticHelpFunc(helpMap)

	}

	// Keep linter from complaining if we don't actually generate code using it
	_ = helpMap
	return helpStr
}

var flagsStaticMap = map[string][]string{

	"create": {"scope-id", "name", "description"},

	"update": {"id", "name", "description", "version"},
}

func (c *StaticCommand) Flags() *base.FlagSets {
	if len(flagsStaticMap[c.Func]) == 0 {
		return c.FlagSet(base.FlagSetNone)
	}

	set := c.FlagSet(base.FlagSetHTTP | base.FlagSetClient | base.FlagSetOutputFormat)
	f := set.NewFlagSet("Command Options")
	common.PopulateCommonFlags(c.Command, f, "static-type host catalog", flagsStaticMap, c.Func)

	extraStaticFlagsFunc(c, set, f)

	return set
}

func (c *StaticCommand) Run(args []string) int {
	initStaticFlags()

	switch c.Func {
	case "":
		return cli.RunResultHelp

	}

	c.plural = "static-type host catalog"
	switch c.Func {
	case "list":
		c.plural = "static-type host catalogs"
	}

	f := c.Flags()

	if err := f.Parse(args); err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}

	if strutil.StrListContains(flagsStaticMap[c.Func], "id") && c.FlagId == "" {
		c.PrintCliError(errors.New("ID is required but not passed in via -id"))
		return base.CommandUserError
	}

	var opts []hostcatalogs.Option

	if strutil.StrListContains(flagsStaticMap[c.Func], "scope-id") {
		switch c.Func {

		case "create":
			if c.FlagScopeId == "" {
				c.PrintCliError(errors.New("Scope ID must be passed in via -scope-id or BOUNDARY_SCOPE_ID"))
				return base.CommandUserError
			}

		}
	}

	client, err := c.Client()
	if c.WrapperCleanupFunc != nil {
		defer func() {
			if err := c.WrapperCleanupFunc(); err != nil {
				c.PrintCliError(fmt.Errorf("Error cleaning kms wrapper: %w", err))
			}
		}()
	}
	if err != nil {
		c.PrintCliError(fmt.Errorf("Error creating API client: %w", err))
		return base.CommandCliError
	}
	hostcatalogsClient := hostcatalogs.NewClient(client)

	switch c.FlagName {
	case "":
	case "null":
		opts = append(opts, hostcatalogs.DefaultName())
	default:
		opts = append(opts, hostcatalogs.WithName(c.FlagName))
	}

	switch c.FlagDescription {
	case "":
	case "null":
		opts = append(opts, hostcatalogs.DefaultDescription())
	default:
		opts = append(opts, hostcatalogs.WithDescription(c.FlagDescription))
	}

	switch c.FlagRecursive {
	case true:
		opts = append(opts, hostcatalogs.WithRecursive(true))
	}

	if c.FlagFilter != "" {
		opts = append(opts, hostcatalogs.WithFilter(c.FlagFilter))
	}

	var version uint32

	switch c.Func {

	case "update":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, hostcatalogs.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	}

	if ok := extraStaticFlagsHandlingFunc(c, f, &opts); !ok {
		return base.CommandUserError
	}

	var resp *api.Response
	var item *hostcatalogs.HostCatalog

	var createResult *hostcatalogs.HostCatalogCreateResult

	var updateResult *hostcatalogs.HostCatalogUpdateResult

	switch c.Func {

	case "create":
		createResult, err = hostcatalogsClient.Create(c.Context, "static", c.FlagScopeId, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = createResult.GetResponse()
		item = createResult.GetItem()

	case "update":
		updateResult, err = hostcatalogsClient.Update(c.Context, c.FlagId, version, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = updateResult.GetResponse()
		item = updateResult.GetItem()

	}

	resp, item, err = executeExtraStaticActions(c, resp, item, err, hostcatalogsClient, version, opts)
	if exitCode := c.checkFuncError(err); exitCode > 0 {
		return exitCode
	}

	output, err := printCustomStaticActionOutput(c)
	if err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}
	if output {
		return base.CommandSuccess
	}

	switch c.Func {

	}

	switch base.Format(c.UI) {
	case "table":
		c.UI.Output(printItemTable(item, resp))

	case "json":
		if ok := c.PrintJsonItem(resp); !ok {
			return base.CommandCliError
		}
	}

	return base.CommandSuccess
}

func (c *StaticCommand) checkFuncError(err error) int {
	if err == nil {
		return 0
	}
	if apiErr := api.AsServerError(err); apiErr != nil {
		c.PrintApiError(apiErr, fmt.Sprintf("Error from controller when performing %s on %s", c.Func, c.plural))
		return base.CommandApiError
	}
	c.PrintCliError(fmt.Errorf("Error trying to %s %s: %s", c.Func, c.plural, err.Error()))
	return base.CommandCliError
}

var (
	extraStaticActionsFlagsMapFunc = func() map[string][]string { return nil }
	extraStaticSynopsisFunc        = func(*StaticCommand) string { return "" }
	extraStaticFlagsFunc           = func(*StaticCommand, *base.FlagSets, *base.FlagSet) {}
	extraStaticFlagsHandlingFunc   = func(*StaticCommand, *base.FlagSets, *[]hostcatalogs.Option) bool { return true }
	executeExtraStaticActions      = func(_ *StaticCommand, inResp *api.Response, inItem *hostcatalogs.HostCatalog, inErr error, _ *hostcatalogs.Client, _ uint32, _ []hostcatalogs.Option) (*api.Response, *hostcatalogs.HostCatalog, error) {
		return inResp, inItem, inErr
	}
	printCustomStaticActionOutput = func(*StaticCommand) (bool, error) { return false, nil }
)
