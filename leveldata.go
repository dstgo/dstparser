package modparser

import lua "github.com/yuin/gopher-lua"

type LevelOverrideItem struct {
	Name  string `mapstructure:"name"`
	Value any    `mapstructure:"value"`
}

// LevelDataOverrides represents level data overrides information
type LevelDataOverrides struct {
	Id                    string  `mapstructure:"id"`
	Name                  string  `mapstructure:"name"`
	Version               float64 `mapstructure:"version"`
	Desc                  string  `mapstructure:"desc"`
	Location              string  `mapstructure:"location"`
	PlayStyle             string  `mapstructure:"playstyle"`
	HideMiniMap           bool    `mapstructure:"hide_minimap"`
	MaxPlayerListPosition float64 `mapstructure:"max_playerlist_position"`
	MinPlayerListPosition float64 `mapstructure:"min_playerlist_position"`
	NumRandomSetPieces    int     `mapstructure:"numrandom_set_pieces"`
	OverrideLevelString   bool    `mapstructure:"override_level_string"`

	// setting
	SettingId   string `mapstructure:"setting_id"`
	SettingName string `mapstructure:"setting_name"`
	SettingDesc string `mapstructure:"setting_desc"`

	// worldgen
	worldGenId   string `mapstructure:"worldgen_id"`
	worldGenName string `mapstructure:"worldgen_name"`
	worldGenDesc string `mapstructure:"worldgen_desc"`

	// meta info
	Overrides           []LevelOverrideItem `mapstructure:"overrides"`
	RandomSetPieces     []string            `mapstructure:"random_set_pieces"`
	RequiredPrefabs     []string            `mapstructure:"required_prefabs"`
	RequiredSetPieces   []string            `mapstructure:"required_setpieces"`
	Substitutes         []string            `mapstructure:"substitutes"`
	BackGroundNodeRange []float64           `mapstructure:"background_node_range"`
}

// ParseLevelDataOverrides parses the leveldataoverrides.lua, returns LevelDataOverrides information
func ParseLevelDataOverrides(luaScript string) (LevelDataOverrides, error) {
	l := lua.NewState()
	defer l.Close()
	if err := l.DoString(luaScript); err != nil {
		return LevelDataOverrides{}, err
	}

	overrideTable := l.ToTable(-1)
	overrideTableL := LTable(overrideTable)

	var levelDataOverrides LevelDataOverrides

	// basic level data
	levelDataOverrides.Id = overrideTableL.GetString("id")
	levelDataOverrides.Desc = overrideTableL.GetString("desc")
	levelDataOverrides.Version = overrideTableL.GetFloat64("version")
	levelDataOverrides.Name = overrideTableL.GetString("name")
	levelDataOverrides.HideMiniMap = overrideTableL.GetBool("hideminimap")
	levelDataOverrides.Location = overrideTableL.GetString("location")
	levelDataOverrides.MaxPlayerListPosition = overrideTableL.GetFloat64("max_playlist_position")
	levelDataOverrides.MinPlayerListPosition = overrideTableL.GetFloat64("min_playlist_position")
	levelDataOverrides.NumRandomSetPieces = int(overrideTableL.GetInt64("numrandom_set_pieces"))
	levelDataOverrides.OverrideLevelString = overrideTableL.GetBool("override_level_string")
	levelDataOverrides.PlayStyle = overrideTableL.GetString("playstyle")

	// setting
	levelDataOverrides.SettingId = overrideTableL.GetString("settings_id")
	levelDataOverrides.SettingDesc = overrideTableL.GetString("settings_desc")
	levelDataOverrides.SettingName = overrideTableL.GetString("settings_name")

	// world gen
	levelDataOverrides.worldGenId = overrideTableL.GetString("worldgen_id")
	levelDataOverrides.worldGenDesc = overrideTableL.GetString("worldgen_desc")
	levelDataOverrides.worldGenName = overrideTableL.GetString("worldgen_name")

	// world override options
	if overrideTableL.GetTable("overrides") != nil {
		overrideTableL.GetTable("overrides").T().ForEach(func(name lua.LValue, value lua.LValue) {
			levelDataOverrides.Overrides = append(levelDataOverrides.Overrides, LevelOverrideItem{
				Name:  name.String(),
				Value: judgeOptionValue(value),
			})
		})
	}

	// random_set_pieces
	if overrideTableL.GetTable("random_set_pieces") != nil {
		overrideTableL.GetTable("random_set_pieces").T().ForEach(func(index lua.LValue, value lua.LValue) {
			levelDataOverrides.RandomSetPieces = append(levelDataOverrides.RandomSetPieces, value.String())
		})
	}

	// random_set_pieces
	if overrideTableL.GetTable("required_prefabs") != nil {
		overrideTableL.GetTable("required_prefabs").T().ForEach(func(index lua.LValue, value lua.LValue) {
			levelDataOverrides.RequiredPrefabs = append(levelDataOverrides.RequiredPrefabs, value.String())
		})
	}

	// random_set_pieces
	if overrideTableL.GetTable("required_setpieces") != nil {
		overrideTableL.GetTable("required_setpieces").T().ForEach(func(index lua.LValue, value lua.LValue) {
			levelDataOverrides.RequiredSetPieces = append(levelDataOverrides.RequiredSetPieces, value.String())
		})
	}

	// substitutes
	if overrideTableL.GetTable("substitutes") != nil {
		overrideTableL.GetTable("substitutes").T().ForEach(func(index lua.LValue, value lua.LValue) {
			levelDataOverrides.Substitutes = append(levelDataOverrides.Substitutes, value.String())
		})
	}

	// background_node_range
	if overrideTableL.GetTable("background_node_range") != nil {
		overrideTableL.GetTable("background_node_range").T().ForEach(func(index lua.LValue, value lua.LValue) {
			levelDataOverrides.BackGroundNodeRange = append(levelDataOverrides.BackGroundNodeRange, float64(lua.LVAsNumber(value)))
		})
	}

	return levelDataOverrides, nil
}
