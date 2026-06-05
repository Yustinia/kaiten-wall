package config

type ConfigModel struct {
	General GeneralModel `toml:"general"`
	Awww    AwwwModel    `toml:"awww"`
}

type GeneralModel struct {
	UseDaemon   string `toml:"use_daemon"`
	DefaultPath string `toml:"default_path"`
}

type AwwwModel struct {
	Filter             string `toml:"filter"`
	TransitionType     string `toml:"transition_type"`
	TransitionStep     int    `toml:"transition_step"`
	TransitionDuration int    `toml:"transition_duration"`
	TransitionFPS      int    `toml:"transition_fps"`
	TransitionAngle    int    `toml:"transition_angle"`
	TransitionPOS      string `toml:"transition_pos"`
	TransitionBezier   string `toml:"transition_bezier"`
	TransitionWave     string `toml:"transition_wave"`
}
