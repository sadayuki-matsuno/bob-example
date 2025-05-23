package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/stephenafamo/bob/gen"
	helpers "github.com/stephenafamo/bob/gen/bobgen-helpers"
	"github.com/stephenafamo/bob/gen/bobgen-psql/driver"
	"github.com/urfave/cli/v2"
)

//go:embed custom_templates
var templates embed.FS

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	app := &cli.App{
		Name:      "bobgen-psql",
		Usage:     "Generate models and factories from your PostgreSQL database",
		UsageText: "bobgen-psql [-c FILE]",
		Version:   helpers.Version(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   helpers.DefaultConfigPath,
				Usage:   "Load configuration from `FILE`",
			},
		},
		Action: run,
	}

	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	config, driverConfig, err := helpers.GetConfigFromFile[any, driver.Config](c.String("config"), "psql")
	if err != nil {
		return err
	}
	config.Generator = "sadayuki-matsuno bob-ex"

	d := driver.New(driverConfig)
	modelTemplates, err := fs.Sub(templates, "custom_templates/models")
	if err != nil {
		return err
	}
	outputs := helpers.DefaultOutputs(
		driverConfig.Output, driverConfig.Pkgname, config.NoFactory,
		&helpers.Templates{
			Models: []fs.FS{
				modelTemplates,
				gen.PSQLModelTemplates,
			},
		},
	)

	state := &gen.State[any]{
		Config:  config,
		Outputs: outputs,
	}

	return gen.Run(c.Context, state, d)
}
