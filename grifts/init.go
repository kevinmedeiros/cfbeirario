package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kevinmedeiros/cfbeirario/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
