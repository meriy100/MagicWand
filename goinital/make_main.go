package goinital

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"os"
	"path/filepath"
)

func (i *Interactor) CreateMain(applicationCmdName string) error {
	fmt.Printf("Create cmd/%s/main.go\n", applicationCmdName)
	if err := os.Mkdir("cmd", 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	if err := os.Mkdir(filepath.Join("cmd", applicationCmdName), 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	// TODO : create main.go

	return nil
}
