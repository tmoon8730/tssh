package main

import (
	"fmt"
	"log"
	"os"
	"sshtemplate/sshtemplate"
	"sshtemplate/utilities"

	"github.com/urfave/cli"
)

func main() {
	home, _ := os.UserHomeDir()
	saveDir := home + "/.tssh/"
	saveFilePath := saveDir + "data.json"
	if _, err := os.Stat(saveFilePath); os.IsNotExist(err) {
		err := os.Mkdir(saveDir,0755)
		utilities.CreateEmptyFile(saveFilePath)
		utilities.Check(err)
	}

	app := &cli.App{
		// tssh <name>
		Action: func(c *cli.Context) error {
			argName := c.Args().First()
			template := sshtemplate.ReadFromFile(saveFilePath)[argName]
			sshtemplate.ExecuteCommand(template)
			fmt.Println(template.Arguments)
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
							arguments := c.Args()[1:]
							sshtemplate.AddTemplate(name, saveFilePath, arguments)
							return nil
						},
					},
					// tssh template remove <name> <command>
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) error {
							name := c.Args().First()
							sshtemplate.RemoveTemplate(name, saveFilePath)
							return nil
						},
					},
					// tssh template list
					{
						Name:  "list",
						Usage: "list existing templates",
						Action: func(c *cli.Context) error {
							sshtemplate.ListTemplates(saveFilePath)
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
