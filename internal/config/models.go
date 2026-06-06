package config

type ConfigModel struct {
	General      GeneralModel   `toml:"general"`
	ClientParams Client         `toml:"client"`
	Wallhaven    WallhavenModel `toml:"wallhaven"`
	Awww         AwwwModel      `toml:"awww"`
	Matugen      MatugenModel   `toml:"matugen"`
}

type GeneralModel struct {
	UseDaemon   string `toml:"use_daemon"`
	DefaultPath string `toml:"default_path"`
}

type AwwwModel struct {
	Filter             string `toml:"filter"`
	TransitionType     string `toml:"transition_type"`
	TransitionStep     *int   `toml:"transition_step"`
	TransitionDuration *int   `toml:"transition_duration"`
	TransitionFPS      *int   `toml:"transition_fps"`
	TransitionAngle    *int   `toml:"transition_angle"`
	TransitionPOS      string `toml:"transition_pos"`
	TransitionBezier   string `toml:"transition_bezier"`
	TransitionWave     string `toml:"transition_wave"`
}

type Client struct {
	APIKey string `toml:"api_key"`
}

type WallhavenModel struct {
	Query      string `toml:"query"`
	Categories string `toml:"categories"`
	Purity     string `toml:"purity"`
	Sorting    string `toml:"sorting"`
	Order      string `toml:"order"`
	AtLeast    string `toml:"at_least"`
	Resolution string `toml:"resolution"`
	Ratios     string `toml:"ratios"`
	Seed       string `toml:"seed"`
}

type MatugenModel struct {
	Scheme      string `toml:"scheme"`
	Contrast    string `toml:"contrast"`
	Mode        string `toml:"mode"`
	SourceIndex string `toml:"source-color-index"`
}
