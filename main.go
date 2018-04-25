package main

import (
	"errors"
	"os"

	"github.com/B0go/octopus/actions"

	lcli "github.com/apex/log/handlers/cli"
	"github.com/urfave/cli"

	"github.com/apex/log"
)

func init() {
	log.SetHandler(lcli.Default)
}

func main() {
	app := cli.NewApp()
	app.Name = "octopus"
	app.Usage = "Manages git projects locally"
	app.Version = "1.0.3"
	app.Commands = []cli.Command{
		{
			Name:        "install",
			Aliases:     []string{"i"},
			Usage:       "octopus install myproject, octopus install --team=myteam, octopus install --all",
			Description: "Install projects. Defaults to all projects installation",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "team",
					Value: "all-teams",
					Usage: "Inform a specific team to get the projects installed: octopus install --team=myteam",
				},
				cli.BoolFlag{
					Name:  "all",
					Usage: "Install all configured projects",
				},
			},
			Action: func(c *cli.Context) error {
				return install(c)
			},
		},
		{
			Name:        "get",
			Aliases:     []string{"g"},
			Usage:       "octopus get SUBCOMMAND",
			Description: "List configured teams and projects",
			Subcommands: []cli.Command{
				{
					Name:        "projects",
					Usage:       "octopus get projects",
					Description: "List configured projects",
					Action: func(c *cli.Context) error {
						return printConfiguredProjects()
					},
				},
				{
					Name:        "teams",
					Usage:       "octopus get teams",
					Description: "List configured teams",
					Action: func(c *cli.Context) error {
						return printConfiguredTeams()
					},
				},
			},
		},
		{
			Name:        "run",
			Aliases:     []string{"r"},
			Usage:       "octopus run myproject",
			Description: "Run a specific project locally",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "branch",
					Value: "master",
					Usage: "Inform a specific branch to run the project",
				},
			},
			Action: func(c *cli.Context) error {
				return run(c)
			},
		},
	}
	app.Run(os.Args)
}

func install(context *cli.Context) error {
	var err error
	if context.Bool("all") {
		err = actions.InstallAllProjects()
	} else if context.String("team") != "all-teams" {
		err = actions.InstallTeamProjects(context.String("team"))
	} else if len(context.Args()) > 0 {
		err = actions.InstallProject(context.Args().Get(0))
	} else {
		log.WithError(errors.New("project to be installed not provided")).
			Error("failed to install project")
		return cli.NewExitError("\n", 1)
	}
	if err != nil {
		log.WithError(err).Error("Failed to install the requested projects")
		return cli.NewExitError("\n", 1)
	}
	return nil
}

func run(context *cli.Context) error {
	if len(context.Args()) > 0 {
		project := context.Args().Get(0)
		branch := context.String("branch")
		if err := actions.RunProject(project, branch); err != nil {
			log.WithError(err).
				Error("Failed to run project")
			return cli.NewExitError("\n", 1)
		}
		return nil
	}
	log.WithError(errors.New("project to be run not provided")).
		Error("Failed to run project")
	return cli.NewExitError("\n", 1)
}

func printConfiguredProjects() error {
	if err := actions.PrintConfiguredProjects(); err != nil {
		log.WithError(err).Error("Failed to get the configured projects!")
		return cli.NewExitError("\n", 1)
	}
	return nil
}

func printConfiguredTeams() error {
	if err := actions.PrintConfiguredTeams(); err != nil {
		log.WithError(err).Error("Failed to get the configured teams!")
		return cli.NewExitError("\n", 1)
	}
	return nil
}
