package app

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nightlord189/uptime-pinger/internal/config"
	"github.com/nightlord189/uptime-pinger/internal/pinger"
	"github.com/nightlord189/uptime-pinger/internal/tg"
	"strings"
	"sync"
)

func Run(ctx context.Context) {
	configInst := ctx.Value("config").(*config.Config)
	adapter := ctx.Value("tg").(*tg.Adapter)
	wg := sync.WaitGroup{}
	for i := 0; i < configInst.WorkersCount; i++ {
		wg.Add(1)
		go tgWorker(ctx, adapter.ChanMsg, &wg)
	}
	wg.Wait()
}

func tgWorker(ctx context.Context, ch chan *tgbotapi.Message, wg *sync.WaitGroup) {
	for msg := range ch {
		processTgMessage(ctx, msg)
	}
	wg.Done()
}

func processTgMessage(ctx context.Context, msg *tgbotapi.Message) {
	text := msg.Text
	fmt.Printf("tg msg: [%s] %s\n", msg.From.UserName, msg.Text)
	adapter := ctx.Value("tg").(*tg.Adapter)
	pingerInst := ctx.Value("pinger").(*pinger.Pinger)
	if strings.HasPrefix(text, "http") {
		resp := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("trying to ping %s", text))
		resp.ReplyToMessageID = msg.MessageID
		adapter.Send(resp)

		//pinging...
		checkResult := pingerInst.CheckUrl(text)
		fmt.Printf("ping %v\n", checkResult)
		resp2 := tgbotapi.NewMessage(msg.Chat.ID, pingResultToText(checkResult))
		resp2.ReplyToMessageID = msg.MessageID
		adapter.Send(resp2)
	} else {
		resp := tgbotapi.NewMessage(msg.Chat.ID, "Hello! Please enter a site url and i will try to ping it. Example: https://google.com")
		resp.ReplyToMessageID = msg.MessageID
		adapter.Send(resp)
	}
}

func pingResultToText(resp pinger.CheckUrlResult) string {
	if resp.Result == "success" {
		return fmt.Sprintf("ping of %s is successfull, time: %v", resp.URL, resp.Time)
	} else {
		return fmt.Sprintf("ping of %s failed, error: %v", resp.URL, resp.Err)
	}
}
