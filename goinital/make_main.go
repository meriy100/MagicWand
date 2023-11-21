package goinital

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/meriy100/magicwand/entities"
	"os"
	"path/filepath"
)

func (i *Interactor) CreateMain(applicationCmdName string, applicationType entities.ApplicationType) error {
	fmt.Printf("Create cmd/%s/main.go\n", applicationCmdName)
	fmt.Printf("%v\n", applicationType)
	if err := os.Mkdir("cmd", 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	if err := os.Mkdir(filepath.Join("cmd", applicationCmdName), 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	// TODO : create main.go

	return nil
}
