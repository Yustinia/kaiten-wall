package theming

import (
	"fmt"

	"github.com/Yustinia/kaiten-wall/internal/config"
)

func buildWallustFlags(s *config.WallustModel) []string {
	var wallustFlags []string

	if s.Backend != "" {
		wallustFlags = append(wallustFlags, fmt.Sprintf("--backend=%s", s.Backend))
	}
	if s.Colorspace != "" {
		wallustFlags = append(wallustFlags, fmt.Sprintf("--colorspace=%s", s.Colorspace))
	}
	if s.Fallback != "" {
		wallustFlags = append(wallustFlags, fmt.Sprintf("--fallback-generator=%s", s.Fallback))
	}
	if s.Palette != "" {
		wallustFlags = append(wallustFlags, fmt.Sprintf("--palette=%s", s.Palette))
	}
	if s.Saturation != "" {
		wallustFlags = append(wallustFlags, fmt.Sprintf("--saturation=%s", s.Saturation))
	}
	if s.DynamicThreshold {
		wallustFlags = append(wallustFlags, "--dynamic-threshold")
	} else if s.Threshold != "" {
		wallustFlags = append(wallustFlags, fmt.Sprintf("--threshold=%s", s.Threshold))
	}

	return wallustFlags
}

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

func ApplyWallust(wallPath string, cfg *config.WallustModel) error {
	wallustFlags := []string{
		"run",
		wallPath,
	}

	if err := runCommand("wallust", wallustFlags, buildWallustFlags(cfg)); err != nil {
		return err
	}

	return nil
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
