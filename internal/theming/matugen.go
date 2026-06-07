package theming

import (
	"fmt"

	"github.com/Yustinia/kaiten-wall/internal/config"
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

func ApplyMatugen(wallPath string, cfg *config.MatugenModel) error {
	matugenFlags := []string{
		"image",
		wallPath,
	}

	if err := runCommand("matugen", matugenFlags, buildMatugenFlags(cfg)); err != nil {
		return err
	}

	return nil
}
