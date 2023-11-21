package goinital

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"os"
	"os/exec"
)

func (i *Interactor) InitGomod(packageName string) error {
	cmd := exec.Command("go", "mod", "init", packageName)
	cmd.Stderr = os.Stderr
	_, err := cmd.Output()
	if err != nil {
		return errors.Wrapf(err, "failed to execute command: %s", cmd.String())
	}
	fmt.Println("Create go.mod")
	return nil
}
