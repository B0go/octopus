package git

import (
	"os"
	"os/exec"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

func init() {
	log.SetHandler(cli.Default)
}

//CloneGitRepository Clones the provided github url in the provided path
func CloneGitRepository(projectURL string, clonePath string) {
	cmdargs := []string{"clone", projectURL, clonePath}
	cmd := exec.Command("git", cmdargs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
