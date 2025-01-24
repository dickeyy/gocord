package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dickeyy/gocord/internal/generator"
	"github.com/dickeyy/gocord/internal/pgm"
	"github.com/dickeyy/gocord/templates"
)

func main() {
	// Define command line flags
	projectName := flag.String("name", "", "Name of the Discord bot project")
	modulePath := flag.String("module", "", "Go module path (e.g., github.com/username/botname)")

	// Parse flags
	flag.Parse()

	// Validate required flags
	if *projectName == "" {
		fmt.Println("Error: Project name is required. Use -name flag")
		flag.Usage()
		os.Exit(1)
	}

	// If module path not provided, construct default from project name
	if *modulePath == "" {
		*modulePath = fmt.Sprintf("github.com/dickeyy/%s", *projectName)
	}

	// Create template data
	config := templates.TemplateData{
		ProjectName: *projectName,
		ModulePath:  *modulePath,
	}

	// Create and run generator
	gen := generator.New(config)
	if err := gen.Generate(); err != nil {
		fmt.Printf("Error generating project: %v\n", err)
		os.Exit(1)
	}

	// Install required packages
	if err := pgm.InstallRequiredPackages(*projectName); err != nil {
		fmt.Printf("Error installing required packages: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created Discord bot '%s'\n", *projectName)
	fmt.Printf("To get started:\n")
	fmt.Printf("  cd %s\n", *projectName)
	fmt.Printf("  go mod tidy\n")
}
