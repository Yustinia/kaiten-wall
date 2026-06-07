package theming

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Yustinia/kaiten-wall/internal/fs"
)

func runCommand(executable string, baseArgs []string, extraArgs []string) error {
	if err := fs.IsCommandExist(executable); err != nil {
		return err
	}

	args := append(baseArgs, extraArgs...)

	cmd := exec.Command(executable, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run %q: %w", executable, err)
	}

	return nil
}
