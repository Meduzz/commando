package commando

import (
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

func init() {
	RegisterVisitor(flags.NewVisitor(model.FlagIntKind, func(f *model.Flag, c *cobra.Command) {
		c.Flags().Int(f.Name, f.Default.(int), f.Description)
	}, func(name string, c *cobra.Command) (any, error) {
		return c.Flags().GetInt(name)
	}))

	RegisterVisitor(flags.NewVisitor(model.FlagStringKind, func(f *model.Flag, c *cobra.Command) {
		c.Flags().String(f.Name, f.Default.(string), f.Description)
	}, func(name string, c *cobra.Command) (any, error) {
		return c.Flags().GetString(name)
	}))

	RegisterVisitor(flags.NewVisitor(model.FlagInt64Kind, func(f *model.Flag, c *cobra.Command) {
		c.Flags().Int64(f.Name, f.Default.(int64), f.Description)
	}, func(name string, c *cobra.Command) (any, error) {
		return c.Flags().GetInt64(name)
	}))

	RegisterVisitor(flags.NewVisitor(model.FlagBoolKind, func(f *model.Flag, c *cobra.Command) {
		c.Flags().Bool(f.Name, f.Default.(bool), f.Description)
	}, func(name string, c *cobra.Command) (any, error) {
		return c.Flags().GetBool(name)
	}))
}
