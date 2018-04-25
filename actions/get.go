package actions

import (
	"os"

	"github.com/B0go/octopus/external/format"
)

//PrintConfiguredProjects prints the configured projects in the current config.yml
func PrintConfiguredProjects() error {
	config, err := loadConfig()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	ymlManipuler := format.DefaultYmlManipuler{}

	if len(config.Projects) > 0 {
		return ymlManipuler.PrintAsYml(config)
	}

	return nil
}

//PrintConfiguredTeams prints the configured teams in the current config.yml
func PrintConfiguredTeams() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	teams := []string{}

	for _, project := range config.Projects {
		teams = append(teams, project.Team)
	}

	teamsStruct := struct {
		Teams []string
	}{eliminateDuplicates(teams)}

	ymlManipuler := format.DefaultYmlManipuler{}

	return ymlManipuler.PrintAsYml(teamsStruct)
}

func eliminateDuplicates(slice []string) []string {
	m := map[string]bool{}

	for _, value := range slice {
		if _, seen := m[value]; !seen {
			slice[len(m)] = value
			m[value] = true
		}
	}

	return slice[:len(m)]
}
