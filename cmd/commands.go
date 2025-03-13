package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func createLaravelCmd() *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "create [project_name]",
		Short: "Create a new Laravel project from template",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]

			fmt.Printf("Creating Laravel project: %s\n", projectName)

			// Clone the template
			if err := cloneTemplate(projectName); err != nil {
				return err
			}

			// Setup the project
			if err := setupProject(projectName); err != nil {
				return err
			}

			fmt.Println("Setup complete! Your Laravel project is ready.")
			return nil
		},
	}

	return createCmd
}
