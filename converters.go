package commando

import (
	"fmt"

	"github.com/Meduzz/commando/model"
	"github.com/spf13/pflag"
)

func init() {
	NewConverter(model.FlagStringKind, func(flag *model.Flag, flagSet *pflag.FlagSet) error {
		defaultValue, ok := flag.Default.(string)

		if !ok {
			return fmt.Errorf("could not turn %v into string", flag.Default)
		}

		flagSet.String(flag.Name, defaultValue, flag.Description)

		return nil
	})

	NewConverter(model.FlagIntKind, func(flag *model.Flag, flagSet *pflag.FlagSet) error {
		defaultValue, ok := flag.Default.(int)

		if !ok {
			return fmt.Errorf("could not turn %v into int", flag.Default)
		}

		flagSet.Int(flag.Name, defaultValue, flag.Description)

		return nil
	})

	NewConverter(model.FlagInt64Kind, func(flag *model.Flag, flagSet *pflag.FlagSet) error {
		defaultValue, ok := flag.Default.(int64)

		if !ok {
			return fmt.Errorf("could not turn %v into int64", flag.Default)
		}

		flagSet.Int64(flag.Name, defaultValue, flag.Description)

		return nil
	})

	NewConverter(model.FlagBoolKind, func(flag *model.Flag, flagSet *pflag.FlagSet) error {
		defaultValue, ok := flag.Default.(bool)

		if !ok {
			return fmt.Errorf("could not turn %v into bool", flag.Default)
		}

		flagSet.Bool(flag.Name, defaultValue, flag.Description)

		return nil
	})
}
