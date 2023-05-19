package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:  "install",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		u, err := user.Current()
		if err != nil {
			return err
		}

		// Create jumppad data directory
		err = os.MkdirAll(dataDir, os.ModePerm)
		if err != nil {
			return err
		}

		homedir := u.HomeDir

		// if bash is present
		if _, err := exec.LookPath("bash"); err == nil {
			snippetPath := path.Join(dataDir, "bash")
			err = os.WriteFile(snippetPath, []byte(bashSnippet), 0644)
			if err != nil {
				return err
			}

			err = injectPrompt(path.Join(homedir, ".bashrc"), snippetPath)
			if err != nil {
				return err
			}

			err = injectPrompt(path.Join(homedir, ".bash_profile"), snippetPath)
			if err != nil {
				return err
			}
		}

		// if zsh is present
		if _, err := exec.LookPath("zsh"); err == nil {
			snippetPath := path.Join(dataDir, "zsh")
			err = os.WriteFile(snippetPath, []byte(zshSnippet), 0644)
			if err != nil {
				return err
			}

			err = injectPrompt(path.Join(homedir, ".zshrc"), snippetPath)
			if err != nil {
				return err
			}
		}

		// if fish is present
		if _, err := exec.LookPath("fish"); err == nil {
			fmt.Println("fish support is not implemented")
		}

		return nil
	},
}

func injectPrompt(filePath string, snippetPath string) error {
	if _, err := os.Stat(filePath); err == nil {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0)
		if err != nil {
			return err
		}
		defer file.Close()

		file.WriteString("source " + snippetPath)
	}

	return nil
}
