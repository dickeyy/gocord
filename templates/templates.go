package templates

import (
	"embed"
	"encoding/json"
	"fmt"
	"text/template"
)

//go:embed files/* templates.json
var templateFS embed.FS

// TemplateData holds all variables that can be used in templates
type TemplateData struct {
	ProjectName string
	ModulePath  string
}

// Template represents a single template file
type Template struct {
	Path        string `json:"path"`        // Path within the template filesystem
	Destination string `json:"destination"` // Relative path in generated project
}

// GetTemplates returns all available templates
func GetTemplates() ([]Template, error) {
	// Read the templates.json file
	data, err := templateFS.ReadFile("templates.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read templates.json: %w", err)
	}

	var templates []Template
	if err := json.Unmarshal(data, &templates); err != nil {
		return nil, fmt.Errorf("failed to parse templates.json: %w", err)
	}

	return templates, nil
}

// GetFS returns the embedded filesystem
func GetFS() embed.FS {
	return templateFS
}

// ParseTemplate parses a single template file and returns the template object
func ParseTemplate(path string) (*template.Template, error) {
	content, err := templateFS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read template %s: %w", path, err)
	}

	tmpl, err := template.New(path).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template %s: %w", path, err)
	}

	return tmpl, nil
}

// ValidateTemplateData ensures all required template data is present
func ValidateTemplateData(data TemplateData) error {
	if data.ProjectName == "" {
		return fmt.Errorf("project name is required")
	}
	if data.ModulePath == "" {
		return fmt.Errorf("module path is required")
	}
	return nil
}
