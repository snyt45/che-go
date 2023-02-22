package main

import (
	"fmt"
	"log"
	"os"

	"github.com/snyt45/che-go/cmd"
	"github.com/snyt45/che-go/internal/yaml"
	"github.com/urfave/cli/v2"
)

func main() {
	path, err := yaml.YamlPath()
	if err != nil {
		log.Fatal("Failed:", err)
	}
	err = yaml.CreateYaml(path)
	if err != nil {
		log.Fatal("Failed:", err)
	}

	app := &cli.App{
		Name:    "che-go",
		Usage:   "A cheat sheet manager",
		Version: "0.1.0",
		Commands: []*cli.Command{
			makeCommand("add", "Add a command", cmd.AddCommand),
			makeCommand("edit", "Edit a command", cmd.EditCommand),
			makeCommand("list", "List all commands", cmd.ListCommand),
			makeCommand("remove", "Remove a command", cmd.RemoveCommand),
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makeCommand(name string, usage string, fn func(*cli.Context) error) *cli.Command {
	cmd := &cli.Command{
		Name:   name,
		Usage:  usage,
		Action: fn,
	}
	return cmd
}
