package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"github.com/vikpe/automapshot/internal/pkg/mapshot"
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	app := NewApp()

	if 1 == len(args) {
		args = append(args, "--help")
	}

	err := app.Run(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return 1
	}

	return 0
}

func NewApp() *cli.App {
	cli.AppHelpTemplate = `{{.Name}} [{{.Version}}]
{{.Description}}

  Usage:   {{.UsageText}}
Example:   {{.Name}} dm2 dm4 dm6
`

	return &cli.App{
		Name:        "automapshot",
		Description: "Automate screenshots of QuakeWorld maps.",
		UsageText:   "automapshot [<maps> ...]",
		Version:     "__VERSION__", // updated during build workflow
		Action: func(c *cli.Context) error {
			// validate setup
			mapSettings, err := mapshot.NewMapSettingsFromJsonFile("map_settings.json")

			if err != nil {
				return err
			}

			err = godotenv.Load()

			if err != nil {
				return err
			}

			// create mapshots
			client := mapshot.NewClient(
				os.Getenv("EZQUAKE_PROCESS_USERNAME"),
				os.Getenv("EZQUAKE_BIN_PATH"),
			)
			mapNames := c.Args().Slice()
			return client.Mapshots(mapNames, mapSettings)
		},
	}
}
