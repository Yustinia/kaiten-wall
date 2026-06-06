package theming

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Yustinia/kaiten-wall/internal/config"
	"github.com/Yustinia/kaiten-wall/internal/fs"
)

func buildMatugenFlags(s *config.MatugenModel) []string {
	var matugenFlags []string

	if s.Scheme != "" {
		matugenFlags = append(matugenFlags, fmt.Sprintf("--type=%s", s.Scheme))
	}
	if s.Contrast != "" {
		matugenFlags = append(matugenFlags, fmt.Sprintf("--contrast=%s", s.Contrast))
	}
	if s.Mode != "" {
		matugenFlags = append(matugenFlags, fmt.Sprintf("--mode=%s", s.Mode))
	}
	if s.SourceIndex != "" {
		matugenFlags = append(matugenFlags, fmt.Sprintf("--source-color-index=%s", s.SourceIndex))
	}

	return matugenFlags
}

func ApplyMatugen(wallPath string, daemonSet *config.MatugenModel) error {
	if err := fs.IsCommandExist("matugen"); err != nil {
		return err
	}

	matugenFlags := []string{
		"image",
		wallPath,
	}
	matugenFlags = append(matugenFlags, buildMatugenFlags(daemonSet)...)

	cmd := exec.Command("matugen", matugenFlags...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run: %w", err)
	}

	return nil
}
