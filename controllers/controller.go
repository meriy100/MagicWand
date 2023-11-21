package controllers

import (
	"github.com/cockroachdb/errors"
	"github.com/manifoldco/promptui"
	"github.com/meriy100/magicwand/entities"
	"github.com/meriy100/magicwand/goinital"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Run() error {
	interactor := goinital.NewInteractor()
	cs := entities.ConfigSet{}
	var err error

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

	cs.PackageName, err = prompt.Run()

	if err != nil {
		return errors.Wrapf(err, "Prompt failed %v")
	}

	if err := interactor.InitGomod(cs.PackageName); err != nil {
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

	cs.ApplicationCommandName, err = prompt.Run()

	if err != nil {
		return errors.Wrapf(err, "Prompt failed %v")
	}

	options := []struct {
		Name  string
		Value entities.ApplicationType
	}{
		{
			Name:  "Rest",
			Value: entities.Rest,
		},
		{
			Name:  "GraphQL",
			Value: entities.GraphQL,
		},
	}
	templates := &promptui.SelectTemplates{
		Active:   "{{ .Name | green }} ",
		Inactive: "{{ .Name }}",
	}

	promptS := promptui.Select{
		Label:        "application type",
		Items:        options,
		Templates:    templates,
		HideSelected: true,
	}

	idx, _, err := promptS.Run()
	cs.ApplicationType = options[idx].Value

	if err != nil {
		return errors.Wrapf(err, "Prompt failed %v")
	}

	if err := interactor.CreateMain(cs.ApplicationCommandName, cs.ApplicationType); err != nil {
		return err
	}

	return nil
}
