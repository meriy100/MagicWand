package controllers

import (
	"github.com/cockroachdb/errors"
	"github.com/manifoldco/promptui"
	"github.com/meriy100/magicwand/goinital"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Run() error {
	interactor := goinital.NewInteractor()
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("Invalid. package name is required")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "package name",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		return errors.Wrapf(err, "Prompt failed %v")
	}

	if err := interactor.InitGomod(result); err != nil {
		return err
	}

	validate = func(input string) error {
		if len(input) < 1 {
			return errors.New("Invalid. package name is required")
		}
		return nil
	}

	prompt = promptui.Prompt{
		Label:    "application command name",
		Validate: validate,
	}

	result, err = prompt.Run()

	if err != nil {
		return errors.Wrapf(err, "Prompt failed %v")
	}

	if err := interactor.CreateMain(result); err != nil {
		return err
	}

	return nil
}
