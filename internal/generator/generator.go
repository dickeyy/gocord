// internal/generator/generator.go
package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dickeyy/gocord/templates"
)

type Generator struct {
	config templates.TemplateData
}

func New(config templates.TemplateData) *Generator {
	return &Generator{config: config}
}

func (g *Generator) Generate() error {
	if err := templates.ValidateTemplateData(g.config); err != nil {
		return err
	}

	tmpls, err := templates.GetTemplates()
	if err != nil {
		return fmt.Errorf("failed to load templates: %w", err)
	}

	for _, tmpl := range tmpls {
		if err := g.generateFile(tmpl); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generateFile(tmpl templates.Template) error {
	// Parse template
	parsed, err := templates.ParseTemplate(tmpl.Path)
	if err != nil {
		return err
	}

	// Create output directory
	outPath := filepath.Join(g.config.ProjectName, tmpl.Destination)
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}

	// Create output file
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Execute template
	return parsed.Execute(f, g.config)
}
