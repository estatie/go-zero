package bug

import (
	"github.com/estatie/go-zero/tools/goctl/internal/cobrax"
	"github.com/spf13/cobra"
)

// Cmd describes a bug command.
var Cmd = cobrax.NewCommand("bug", cobrax.WithRunE(cobra.NoArgs), cobrax.WithArgs(cobra.NoArgs))
