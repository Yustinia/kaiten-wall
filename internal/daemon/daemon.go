package daemon

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Yustinia/kaiten-wall/internal/config"
	"github.com/Yustinia/kaiten-wall/internal/fs"
)

func buildAwwwFlags(s *config.AwwwModel) []string {
	var awwwFlags []string

	if s.Filter != "" {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--filter=%s", s.Filter))
	}
	if s.TransitionType != "" {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-type=%s", s.TransitionType))
	}
	if s.TransitionStep != "" {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-step=%s", s.TransitionStep))
	}
	if s.TransitionDuration != "" {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-duration=%s", s.TransitionDuration))
	}
	if s.TransitionFPS != "" {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-fps=%s", s.TransitionFPS))
	}
	if s.TransitionAngle != "" {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-angle=%s", s.TransitionAngle))
	}
	if s.TransitionPOS != "" {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-pos=%s", s.TransitionPOS))
	}
	if s.TransitionBezier != "" {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-bezier=%s", s.TransitionBezier))
	}
	if s.TransitionWave != "" {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-wave=%s", s.TransitionWave))
	}

	return awwwFlags
}

func RunAwww(daemonInUse string, wallPath string, daemonSet *config.AwwwModel) error {
	if err := fs.IsCommandExist(daemonInUse); err != nil {
		return err
	}

	awwwFlags := []string{
		"img",
		wallPath,
	}
	awwwFlags = append(awwwFlags, buildAwwwFlags(daemonSet)...)

	cmd := exec.Command(daemonInUse, awwwFlags...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run: %w", err)
	}

	return nil
}
