package modparser

import (
	"errors"
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

// ParseModInfo returns the parsed modinfo from lua script
func ParseModInfo(luaScript string) (ModInfo, error) {
	return ParseModInfoWithEnv("", "", luaScript)
}

// ParseModInfoWithEnv parse mod info from lua script with mod environment variables.
func ParseModInfoWithEnv(folderName, locale, luaScript string) (ModInfo, error) {
	l := lua.NewState()
	defer l.Close()

	// prepare mod pre environment
	// see https://forums.kleientertainment.com/forums/topic/150829-game-update-571392/
	l.SetGlobal("locale", lua.LString(locale))
	// dir name like "workshop-1274919201"
	l.SetGlobal("folder_name", lua.LString(folderName))
	// ChooseTranslationTable function will be called in the script,
	// if is needed to translate configuration_options by specific language
	// egs. ChooseTranslationTable(table,[key])
	l.SetGlobal("ChooseTranslationTable", ChooseTranslationTable(l, locale))

	// parse script
	if err := l.DoString(luaScript); err != nil {
		return ModInfo{}, err
	}

	// parse simple info
	modInfo, err := parseModSimpleInfo(l.G.Global)
	if err != nil {
		return ModInfo{}, err
	}

	// parse options
	modOptions, err := parseModOptions(LTable(l.G.Global).GetTable("configuration_options").T())
	if err != nil {
		return ModInfo{}, err
	}
	modInfo.ConfigurationOptions = modOptions

	return modInfo, nil
}

// ChooseTranslationTable returns *lua.LFunction, this function used to choose the translation table in lua state.
func ChooseTranslationTable(l *lua.LState, locale string) *lua.LFunction {
	return l.NewFunction(func(fnl *lua.LState) int {

		// check first table param
		translationTable := fnl.ToTable(1)
		if translationTable == lua.LNil || translationTable.Len() == 0 {
			return 1
		}

		// Get specific locale table
		target := translationTable.RawGetString(locale)
		if target != lua.LNil {
			fnl.Push(target)
		} else { // or use the default
			fnl.Push(translationTable.RawGetInt(1))
		}

		return 1
	})
}

// parse simple info
func parseModSimpleInfo(table *lua.LTable) (ModInfo, error) {
	var modinfo ModInfo
	if table == lua.LNil {
		return modinfo, errors.New("nil lua global")
	}
	g := LTable(table)

	// basic info
	modinfo.Id = g.GetString("id")
	modinfo.Name = g.GetString("name")
	modinfo.Description = g.GetString("description")
	modinfo.Author = g.GetString("author")
	modinfo.Version = g.GetString("version")

	// ds
	modinfo.ApiVersion = int(g.GetInt64("api_version"))
	modinfo.DontStarveCompatible = g.GetBool("dont_starve_compatible")
	modinfo.ReignOfGiantsCompatible = g.GetBool("reign_of_giants_compatible")
	modinfo.ShipWreckedCompatible = g.GetBool("shipwrecked_compatible")
	modinfo.HamletCompatible = g.GetBool("hamlet_compatible")

	// dst
	modinfo.ApiVersionDst = int(g.GetInt64("api_version_dst"))
	modinfo.DstCompatible = g.GetBool("dst_compatible")
	modinfo.AllClientRequired = g.GetBool("all_client_required")
	modinfo.ClientOnly = g.GetBool("client_only_mod")
	modinfo.ServerOnly = g.GetBool("server_only_mod")
	modinfo.ForgeCompatible = g.GetBool("forge_compatible")

	// meta info
	if g.GetTable("server_filter_tags") != nil {
		g.GetTable("server_filter_tags").T().ForEach(func(key lua.LValue, value lua.LValue) {
			modinfo.FilterTags = append(modinfo.FilterTags, value.String())
		})
	}
	modinfo.Priority = g.GetFloat64("priority")
	modinfo.Icon = g.GetString("icon")
	modinfo.IconAtlas = g.GetString("icon_atlas")

	return modinfo, nil
}

// parse configuration_options from lua script
func parseModOptions(options *lua.LTable) ([]ModOption, error) {
	if options == nil || options == lua.LNil {
		return nil, errors.New("nil configuration_options table")
	}

	var modOptions []ModOption

	// iterate configuration_options
	options.ForEach(func(index lua.LValue, option lua.LValue) {
		if option.Type() != lua.LTTable {
			return
		}
		var modOption ModOption

		optTable := option.(*lua.LTable)
		loptTable := LTable(optTable)

		if t := loptTable.GetTable("options").T(); t != nil || t != lua.LNil {
			modOption.Options = parseModOptionItems(t)
		}
		modOption.Name = loptTable.GetString("name")
		modOption.Label = loptTable.GetString("label")
		modOption.Hover = loptTable.GetString("hover")
		modOption.Client = loptTable.GetBool("client")

		// default value
		defaultValue := LTable(optTable).Get("default")
		switch defaultValue.Type() {
		case lua.LTString:
			modOption.Default = lua.LVAsString(defaultValue)
		case lua.LTNumber:
			modOption.Default = float64(lua.LVAsNumber(defaultValue))
		case lua.LTBool:
			modOption.Default = lua.LVAsBool(defaultValue)
		}

		if loptTable.GetTable("tags") != nil {
			loptTable.GetTable("tags").T().ForEach(func(key lua.LValue, value lua.LValue) {
				modOption.Tags = append(modOption.Tags, value.String())
			})
		}

		modOptions = append(modOptions, modOption)
	})

	return modOptions, nil
}

// parse mod option items
func parseModOptionItems(optTable *lua.LTable) []ModOptionItem {
	if optTable == nil || optTable == lua.LNil {
		return nil
	}

	var items []ModOptionItem
	// iterate items
	optTable.ForEach(func(index lua.LValue, item lua.LValue) {
		if item.Type() != lua.LTTable {
			return
		}

		// build item
		var modItem ModOptionItem

		itemTable := LTable(item.(*lua.LTable))
		modItem.Description = itemTable.GetString("description")

		dataValue := itemTable.Get("data")
		switch dataValue.Type() {
		case lua.LTBool:
			modItem.Data = lua.LVAsBool(dataValue)
		case lua.LTNumber:
			modItem.Data = float64(lua.LVAsNumber(dataValue))
		case lua.LTString:
			modItem.Data = lua.LVAsString(dataValue)
		}

		// if it has no description, use the string of data
		if len(modItem.Description) == 0 {
			modItem.Description = fmt.Sprintf("%+v", modItem.Data)
		}

		items = append(items, modItem)
	})

	return items
}
