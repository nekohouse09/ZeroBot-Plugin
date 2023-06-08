// Package alipayvoice 支付宝到账语音
package alipayvoice

import (
	"fmt"
	"strings"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	alipayvoiceURL = "https://mm.cqu.cc/share/zhifubaodaozhang/mp3/%v.mp3"
)

func init() { // 插件主体
	engine := control.Register("alipayvoice", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault:  false,
		Brief:             "支付宝到账语音",
		Help:              "- 支付宝到账 1",
		PrivateDataFolder: "alipayvoice",
	})

	// 开启
	engine.OnRegex(`^支付宝到账([0-9]+)`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			args := ctx.State["regex_matched"].([]string)[1]
			ctx.SendChain(message.Record(fmt.Sprintf(alipayvoiceURL, strings.TrimSpace(args))))
		})
}
