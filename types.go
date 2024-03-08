package modparser

type ModInfo struct {
	// basic information
	Id          string `mapstructure:"id"`
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Author      string `mapstructure:"author"`
	Version     string `mapstructure:"version"`
	ForumThread string `mapstructure:"forum_thread"`

	// dont starve
	ApiVersion              int  `mapstructure:"api_version"`
	DontStarveCompatible    bool `mapstructure:"dont_starve_compatible"`
	ReignOfGiantsCompatible bool `mapstructure:"reign_of_giants_compatible"`
	ShipWreckedCompatible   bool `mapstructure:"shipwrecked_compatible"`
	HamletCompatible        bool `mapstructure:"hamlet_compatible"`

	// dont starve together
	ApiVersionDst     int  `mapstructure:"api_version_dst"`
	DstCompatible     bool `mapstructure:"dst_compatible"`
	AllClientRequired bool `mapstructure:"all_client_required"`
	ClientOnly        bool `mapstructure:"client_only_mod"`
	ServerOnly        bool `mapstructure:"server_only_mod"`
	ForgeCompatible   bool `mapstructure:"forge_compatible"`

	// meta info
	FilterTags []string `mapstructure:"server_filter_tags"`
	Priority   float64  `mapstructure:"priority"`
	Icon       string   `mapstructure:"icon"`
	IconAtlas  string   `mapstructure:"icon_atlas"`

	// configuration
	ConfigurationOptions []ModOption `mapstructure:"configuration_options"`
}

// ModOption represents a mod option in 'configuration_options'
type ModOption struct {
	// option name, maybe empty
	Name string `mapstructure:"name"`
	// option label, maybe empty
	Label string `mapstructure:"label"`
	// hover tooltip, maybe empty
	Hover string `mapstructure:"hover"`
	// default value of this option
	Default any      `mapstructure:"default"`
	Client  bool     `mapstructure:"client"`
	Tags    []string `mapstructure:"tags"`
	// options available
	Options []ModOptionItem `mapstructure:"options"`
}

type ModOptionItem struct {
	Description string `mapstructure:"description"`
	Data        any    `mapstructure:"data"`
}
