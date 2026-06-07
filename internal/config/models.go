package config

type ConfigModel struct {
	General      GeneralModel   `toml:"general"`
	ClientParams Client         `toml:"client"`
	Wallhaven    WallhavenModel `toml:"wallhaven"`
	Awww         AwwwModel      `toml:"awww"`
	Matugen      MatugenModel   `toml:"matugen"`
	Wallust      WallustModel   `toml:"wallust"`
}

type GeneralModel struct {
	UseDaemon   string `toml:"use_daemon"`
	UseThemer   string `toml:"use_themer"`
	DefaultPath string `toml:"default_path"`
}

type AwwwModel struct {
	Filter             string `toml:"filter"`
	TransitionType     string `toml:"transition_type"`
	TransitionStep     string `toml:"transition_step"`
	TransitionDuration string `toml:"transition_duration"`
	TransitionFPS      string `toml:"transition_fps"`
	TransitionAngle    string `toml:"transition_angle"`
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
	TopRange   string `toml:"toprange"`
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

type WallustModel struct {
	Backend          string `toml:"backend"`
	Colorspace       string `toml:"colorspace"`
	Fallback         string `toml:"fallback"`
	Palette          string `toml:"palette"`
	Saturation       string `toml:"saturation"`
	Threshold        string `toml:"threshold"`
	DynamicThreshold bool   `toml:"dynamic_threshold"`
}
