package actions

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

func init() {
	log.SetHandler(cli.Default)
}

//RunProject runs the provided project using its specific script
func RunProject(name string, branch string) error {
	log.WithField("branch", branch).
		WithField("project", name).
		Info("Starting to run")

	usr, err := user.Current()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/Code/%s", usr.HomeDir, name)

	runScript := fmt.Sprintf("%s/scripts/run.sh", path)
	if _, err := os.Stat(runScript); os.IsNotExist(err) {
		return err
	}

	var cmd *exec.Cmd
	if branch != "master" {
		cmd = exec.Command("/bin/sh", runScript, branch)
	} else {
		cmd = exec.Command("/bin/sh", runScript)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return nil
}
