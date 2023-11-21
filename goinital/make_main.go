package goinital

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/meriy100/magicwand/entities"
	"os"
	"path/filepath"
)

func (i *Interactor) CreateMain(appName string, appType entities.AppType) error {
	fmt.Printf("Create cmd/%s/main.go\n", appName)
	fmt.Printf("%v\n", appType)
	if err := os.Mkdir("cmd", 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	if err := os.Mkdir(filepath.Join("cmd", appName), 0755); err != nil {
		return errors.Wrapf(err, "make directory failed")
	}

	// TODO : create main.go

	return nil
}
