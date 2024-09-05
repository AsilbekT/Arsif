package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// startprojectCmd represents the startproject command
var startprojectCmd = &cobra.Command{
	Use:   "startproject [name]",
	Short: "Creates a new project directory structure",
	Long: `StartProject initializes a new project directory with a predefined structure. For example:

Usage:
  arsif startproject mynewapp

This command creates directories for various parts of the application structure.`,
	Args: cobra.ExactArgs(1), // Ensuring exactly one argument is provided, the project name
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		createProjectStructure(projectName)
		fmt.Printf("Project '%s' created successfully\n", projectName)
	},
}

func createProjectStructure(projectName string) {
	baseDir := filepath.Join(".", projectName)
	dirs := []string{
		"cmd/" + projectName,
		"internal/apps",
		"internal/core",
		"pkg",
		"configs",
		"deployments",
		"docs",
		"tests",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(baseDir, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", fullPath, err)
			return
		}
	}
}

func init() {
	rootCmd.AddCommand(startprojectCmd)

	// Additional flags can be added here if needed
}
