package logic

import (
	td "github.com/swxctx/malatd"
	"github.com/txxzx/goAi/args"
)

/**
    @date: 2024/4/16
**/

// Do handler
func V1_Chat_Do(ctx *td.Context, arg *args.ChatDoArgsV1) (*args.ChatDoResultV1, *td.Rerror) {
	switch arg.Platform {
	case args.PlatformBaidu:

	case args.PlatformZhiPu:

	case args.PlatformXunFei:

	}
	return new(args.ChatDoResultV1), nil
}

//
// ChatZhiPu
//  @Description: 智普
//  @param ctx
//  @param arg
//  @return *args.ChatDoResultV1
//  @return *td.Rerror
//
//func ChatZhiPu(ctx *td.Context, arg *args.ChatDoArgsV1) (*args.ChatDoResultV1, *td.Rerror) {
//	// 流式输出
//
//}
