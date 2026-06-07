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
