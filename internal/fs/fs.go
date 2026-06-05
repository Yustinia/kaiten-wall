package fs

import (
	"fmt"
	"os"
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
