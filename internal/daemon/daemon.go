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
	if s.TransitionStep != nil {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-step=%d", *s.TransitionStep))
	}
	if s.TransitionDuration != nil {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-duration=%d", *s.TransitionDuration))
	}
	if s.TransitionFPS != nil {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-fps=%d", *s.TransitionFPS))
	}
	if s.TransitionAngle != nil {
		awwwFlags = append(awwwFlags, fmt.Sprintf("--transition-angle=%d", *s.TransitionAngle))
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
