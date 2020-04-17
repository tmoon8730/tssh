package main

import (
	"fmt"
	"log"
	"os"
	"sshtemplate/sshtemplate"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		// tssh <name>
		Action: func(c *cli.Context) error {
			argName := c.Args().First()

			a := sshtemplate.ReadFromFile("tempData.json")
			template := a[argName]
			sshtemplate.ExecuteCommand(template)
			fmt.Println(template.Command)
			return nil
		},
		Commands: []cli.Command{
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []cli.Command{
					// tssh template add <name> <command>
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(c *cli.Context) error {
							name := c.Args().First()
							command := c.Args().Get(1)

							sshtemplate.AddTemplate(name, command, "tempData.json")

							return nil
						},
					},
					// tssh template remove <name> <command>
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) error {
							name := c.Args().First()

							sshtemplate.RemoveTemplate(name, "tempData.json")
							return nil
						},
					},
				},
			},
			{
				Name:    "help",
				Aliases: []string{"h"},
				Usage:   "display available commands",
				Action: func(c *cli.Context) error {
					fmt.Println("tssh Available Commands:")
					fmt.Println("Launch saved ssh template: tssh <name>")
					fmt.Println("Add template: tssh template add <name> <command>")
					fmt.Println("Remove template: tssh template remove <name> <command>")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
