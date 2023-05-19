package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:  "uninstall",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		u, err := user.Current()
		if err != nil {
			return err
		}

		homedir := u.HomeDir

		// if bash is present
		if _, err := exec.LookPath("bash"); err == nil {
			snippetPath := path.Join(dataDir, "bash")

			err = removePrompt(path.Join(homedir, ".bashrc"), snippetPath)
			if err != nil {
				return err
			}

			err = removePrompt(path.Join(homedir, ".bash_profile"), snippetPath)
			if err != nil {
				return err
			}
		}

		// if zsh is present
		if _, err := exec.LookPath("zsh"); err == nil {
			snippetPath := path.Join(dataDir, "zsh")

			err = removePrompt(path.Join(homedir, ".zshrc"), snippetPath)
			if err != nil {
				return err
			}
		}

		if _, err := exec.LookPath("fish"); err == nil {
			fmt.Println("fish support is not implemented")
		}

		return nil
	},
}

func removePrompt(file string, snippetPath string) error {
	content := []byte{}
	if _, err := os.Stat(file); err == nil {
		content, err = os.ReadFile(file)
		if err != nil {
			return err
		}
	}

	// find the snippet and replace it
	content = bytes.Replace(content, []byte("source "+snippetPath), []byte(""), 1)
	err := os.WriteFile(file, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
