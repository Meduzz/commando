package model

type (
	Flag struct {
		Name        string   `json:"name"`
		Kind        FlagKind `json:"kind"`
		Default     any      `json:"default,omitempty"`
		Description string   `json:"description,omitempty"`
	}

	FlagKind string
)

const (
	FlagStringKind = FlagKind("string")
	FlagIntKind    = FlagKind("int")
	FlagInt64Kind  = FlagKind("int64")
	FlagBoolKind   = FlagKind("bool")
)
