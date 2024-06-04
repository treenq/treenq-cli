package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/treenq/treenq-cli/src/store"
	"github.com/treenq/treenq-cli/src/usecase"
	"github.com/urfave/cli/v2"
)

func main() {
	configFolder := filepath.Join(xdg.Home, ".tq")
	os.MkdirAll(configFolder, os.ModePerm)

	configPath := filepath.Join(configFolder, "config.json")
	f, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		err = fmt.Errorf("failed to open config file %s: %w", configPath, err)
		log.Fatalln(err)
	}

	store, close, err := store.NewStore(f)
	if err != nil {
		err = fmt.Errorf("failed to create store: %w", err)
		log.Fatalln(err)
	}
	defer close()
	contextUsecase := usecase.NewContextUsecase(store)

	app := &cli.App{
		Name:  "tq",
		Usage: "treenq CLI",
		Commands: []*cli.Command{
			{
				Name:  "ctx",
				Usage: "Manage connected treenq instances",
				Subcommands: []*cli.Command{
					{
						Name:  "new",
						Usage: "Create new context",
						Action: func(c *cli.Context) error {
							name, url := c.Args().Get(0), c.Args().Get(1)
							if name == "" {
								return fmt.Errorf("name is required")
							}
							if url == "" {
								return fmt.Errorf("url is required")
							}

							return contextUsecase.NewContext(c.Context, name, url)
						},
					},
					{
						Name:  "set",
						Usage: "Set current context",
						Action: func(c *cli.Context) error {
							name := c.Args().First()
							return contextUsecase.SetContext(c.Context, name)
						},
					},
					{
						Name:  "list",
						Usage: "List contexts",
						Action: func(c *cli.Context) error {
							list, _ := contextUsecase.ListContexts(c.Context)
							for i := range list {
								tpl := "%s: %s\n"
								if list[i].Active {
									tpl = "* " + tpl
								}
								fmt.Printf(tpl, list[i].Name, list[i].Url)
							}

							return nil
						},
					},
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
