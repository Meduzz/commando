package flags

import "github.com/Meduzz/commando/model"

func StringFlag(name string, defaultValue string, description string) *model.Flag {
	return &model.Flag{Name: name, Kind: model.FlagStringKind, Default: defaultValue, Description: description}
}

func IntFlag(name string, defaultValue int, description string) *model.Flag {
	return &model.Flag{Name: name, Kind: model.FlagIntKind, Default: defaultValue, Description: description}
}

func Int64Flag(name string, defaultValue int64, description string) *model.Flag {
	return &model.Flag{Name: name, Kind: model.FlagInt64Kind, Default: defaultValue, Description: description}
}

func BoolFlag(name string, defaultValue bool, description string) *model.Flag {
	return &model.Flag{Name: name, Kind: model.FlagBoolKind, Default: defaultValue, Description: description}
}
