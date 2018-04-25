package actions

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/B0go/octopus/external/format"
	"github.com/B0go/octopus/external/git"
	"github.com/B0go/octopus/external/system"

	"github.com/B0go/octopus/config"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

func init() {
	log.SetHandler(cli.Default)
}

var (
	normalPadding    = cli.Default.Padding
	increasedPadding = normalPadding * 2
)

//Database represents database configuration
type Database struct {
	Name string `yaml:"name"`
}

//EnvVariable represents an environment variable
type EnvVariable struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

//Descriptor describes a service
type Descriptor struct {
	Project      string        `yaml:"project"`
	EnvVars      []EnvVariable `yaml:"envVars"`
	CustomScript *string       `yaml:"customScript"`
	Database     *Database     `yaml:"database"`
}

//InstallAllProjects installs all projects configured in the config.yml
func InstallAllProjects() error {
	log.Warn("\033[1mALL PROJECTS INSTALLATION\033[0m")

	config, err := loadConfig()
	if err != nil {
		return nil
	}

	for _, project := range config.Projects {
		err = install(project)
		if err != nil {
			log.WithError(err).
				Error("Failed to install project")
			log.Warn("Skipping")
		}
	}
	return nil
}

//InstallTeamProjects install the projects related to a specific team
func InstallTeamProjects(team string) error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	for _, project := range config.Projects {
		if project.Team == team {
			err = install(project)
			if err != nil {
				log.WithError(err).
					Error("Failed to install project")
				log.Warn("Skipping")
			}
		}
	}
	return nil
}

//InstallProject installs the provided project
func InstallProject(name string) error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	for _, project := range config.Projects {
		if project.Name == name {
			return install(project)
		}
	}
	return errors.New("project not configured")
}

func install(project config.Project) error {
	log.WithField("repository", project.Repository).
		WithField("team", project.Team).
		WithField("project", project.Name).
		Infof("\033[1mPROJECT INSTALLATION\033[0m")

	cli.Default.Padding = increasedPadding

	usr, err := user.Current()
	if err != nil {
		return err
	}

	projectsPath := fmt.Sprintf("%s/Code", usr.HomeDir)

	os.MkdirAll(projectsPath, 0755)

	projectPath := fmt.Sprintf("%s/%s", projectsPath, project.Name)

	cloneProject(project, projectPath)

	descriptorPath := fmt.Sprintf("%s/octopus.yml", projectPath)

	if _, err := os.Stat(descriptorPath); os.IsNotExist(err) {
		return err
	}

	mgrConfigPath := fmt.Sprintf("%s/.octopus/projects/%s", usr.HomeDir, project.Name)

	os.MkdirAll(mgrConfigPath, 0755)

	descriptor := Descriptor{}

	ymlManipuler := format.DefaultYmlManipuler{}

	err = ymlManipuler.ReadYml(&descriptor, descriptorPath)
	if err != nil {
		return err
	}

	execSpecificInstallaton(descriptor, mgrConfigPath)

	log.Info("Successfully installed!")
	cli.Default.Padding = normalPadding
	return nil
}

func execSpecificInstallaton(descriptor Descriptor, mgrConfigPath string) {

	installEnvironmentVariables(mgrConfigPath, descriptor)

	if descriptor.CustomScript != nil {
		executeCustomScript(descriptor.CustomScript)
	}

	if descriptor.Database != nil {
		createRequestedDatabase(descriptor.Database)
	}
}

func cloneProject(project config.Project, projectPath string) {
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		git.CloneGitRepository(project.Repository, projectPath)
	} else {
		log.Info("Project already cloned!")
	}
}

func installEnvironmentVariables(mgrConfigPath string, descriptor Descriptor) {
	log.Info("Installing environment variables...")
	envVarsPath := fmt.Sprintf("%s/environment_variables", mgrConfigPath)

	var buffer bytes.Buffer

	for _, variable := range descriptor.EnvVars {
		line := fmt.Sprintf("%s=%s\n", variable.Name, variable.Value)
		buffer.WriteString(line)
	}

	ioutil.WriteFile(envVarsPath, []byte(buffer.String()), 0644)
}

func executeCustomScript(script *string) {
	log.Info("Executing custom script...")
	scriptSlice := strings.Split(*script, " ")

	cmd := exec.Command(scriptSlice[0], scriptSlice[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func createRequestedDatabase(database *Database) {
	log.Info("Creating requested database...")

	cmd := exec.Command("createdb", "-Upostgres", database.Name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func loadConfig() (*config.Config, error) {
	usrRetriever := system.OSUserRetriever{}
	fsReader := system.OSFileSystemReader{}
	ymlManipuler := format.DefaultYmlManipuler{}
	return config.Load(usrRetriever, fsReader, ymlManipuler)
}
