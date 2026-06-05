package fs

import (
	"fmt"
	"os"
	"os/exec"
)

func IsFileExist(file string) error {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %w", err)
	} else if err != nil {
		return fmt.Errorf("file inaccessible: %w", err)
	}

	return nil
}

func IsCommandExist(cmd string) error {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return fmt.Errorf("%q daemon missing: %w", cmd, err)
	}

	return nil
}
