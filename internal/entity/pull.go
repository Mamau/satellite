package entity

import (
	"fmt"

	"github.com/mamau/satellite/pkg"
)

// Pull documentation for service params
// https://docs.docker.com/engine/reference/commandline/pull
type Pull struct {
	Name                string `yaml:"name" validate:"required,min=1"`
	Image               string `yaml:"image" validate:"required,min=1"`
	Version             string `yaml:"version"`
	DisableContentTrust string `yaml:"disable-content-trust"`
	Description         string `yaml:"description"`
	AllTags             bool   `yaml:"all-tags"`
	Quiet               bool   `yaml:"quiet"`
}

func (p *Pull) GetExecCommand() string {
	return string(DOCKER)
}

func (p *Pull) GetDescription() string {
	return p.Description
}

func (p *Pull) GetQuiet() string {
	if p.Quiet {
		return "--quiet"
	}
	return ""
}

func (p *Pull) GetDisableContentTrust() string {
	if p.DisableContentTrust != "" {
		return fmt.Sprintf("--disable-content-trust=%s", p.DisableContentTrust)
	}

	return ""
}
func (p *Pull) GetImage() string {
	if p.Version != "" {
		return fmt.Sprintf("%s:%s", p.Image, p.Version)
	}
	return p.Image
}

func (p *Pull) GetAllTags() string {
	if p.AllTags {
		return "--all-tags"
	}
	return ""
}

func (p *Pull) GetName() string {
	return p.Name
}

func (p *Pull) ToCommand(args []string) []string {
	bc := pkg.MergeSliceOfString([]string{
		"pull",
		p.GetDisableContentTrust(),
		p.GetAllTags(),
		p.GetQuiet(),
		p.GetImage(),
	})
	configurator := newPureConfigConfigurator(bc, []string{})
	return append(bc, configurator.getClientCommand()...)
}
