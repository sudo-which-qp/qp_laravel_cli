package main

import (
	"fmt"
	"os/exec"

	"github.com/go-git/go-git/v5"
)

var templateRepo = "https://github.com/sudo-which-qp/qp_laravel_api_template_v2.git"

func cloneTemplate(projectName string) error {
	fmt.Printf("Cloning Laravel template into %s...\n", projectName)
	_, err := git.PlainClone(projectName, false, &git.CloneOptions{
		URL: templateRepo,
	})
	if err != nil {
		return err
	}
	fmt.Println("Template cloned successfully!")
	return nil
}

func setupProject(projectName string) error {
	fmt.Println("Installing dependencies...")
	composerCmd := exec.Command("composer", "install")
	composerCmd.Dir = projectName
	if err := composerCmd.Run(); err != nil {
		return err
	}

	fmt.Println("Setting up .env file...")
	if err := setupEnvFile(projectName); err != nil {
		return err
	}

	fmt.Println("Generating application key...")
	artisanCmd := exec.Command("php", "artisan", "key:generate")
	artisanCmd.Dir = projectName
	if err := artisanCmd.Run(); err != nil {
		return err
	}

	return nil
}

func setupEnvFile(projectDir string) error {
	// Copy .env.example to .env
	copyCmd := exec.Command("cp", ".env.example", ".env")
	copyCmd.Dir = projectDir
	return copyCmd.Run()
}
