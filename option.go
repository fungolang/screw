package screw

import (
	"fmt"

	"github.com/antlabs/strsim"
)

func (c *Screw) maybeOpt(optionName string) string {
	opts := make([]string, len(c.shortAndLong))
	index := 0
	for k := range c.shortAndLong {
		opts[index] = k
		index++
	}

	//Direct return without long and short commands
	if len(opts) == 0 {
		return ""
	}

	m := strsim.FindBestMatchOne(optionName, opts)
	if m.Score > 0.0 {
		return m.S
	}

	return ""
}

func (c *Screw) genMaybeHelpMsg(optionName string) string {
	if s := c.maybeOpt(optionName); len(s) > 0 {
		return fmt.Sprintf("\n	Did you mean --%s?\n", s)
	}

	if _, ok := c.subcommand[optionName]; ok {
		return fmt.Sprintf("\n	Did you mean '%s' subcommand?\n", optionName)
	}
	return ""
}
