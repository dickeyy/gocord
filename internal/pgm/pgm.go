package pgm

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func InstallRequiredPackages(projectDir string) error {
	reqPackages := []string{
		"github.com/bwmarrin/discordgo",
		"github.com/joho/godotenv",
		"github.com/rs/zerolog",
	}

	// Change to project directory
	absPath, err := filepath.Abs(projectDir)
	if err != nil {
		return err
	}

	// Install each package
	for _, pkg := range reqPackages {
		cmd := exec.Command("go", "get", pkg)
		cmd.Dir = absPath
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install %s: %w", pkg, err)
		}
	}

	return nil
}
