// Code generated by 'malatd gen' command.
// DO NOT EDIT!

package api

import (
	td "github.com/swxctx/malatd"
	"github.com/swxctx/malatd/binding"

	"github.com/txxzx/goAi/args"
	"github.com/txxzx/goAi/logic"
)

// Ping handler
func PingHandle(ctx *td.Context) {
	// bind arg
	arg := new(args.PingArgsV1)
	if err := binding.Binder(ctx, arg); err != nil {
		ctx.RenderRerr(td.RerrInternalServer.SetReason(err.Error()))
		return
	}

	// api logic
	result, rerr := logic.V1_Test_Ping(ctx, arg)
	if rerr != nil {
		ctx.RenderRerr(rerr)
		return
	}
	ctx.Render(result)
}