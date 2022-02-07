package main

import (
	"context"
	"fmt"
	"github.com/nightlord189/uptime-pinger/internal/app"
	"github.com/nightlord189/uptime-pinger/internal/config"
	"github.com/nightlord189/uptime-pinger/internal/pinger"
	"github.com/nightlord189/uptime-pinger/internal/tg"
)

func main() {
	fmt.Println("start")

	ctx := context.Background()
	configInst, err := config.Load("configs/config.json")
	if err != nil {
		panic(fmt.Sprintf("load config err: %v", err))
	}
	pingerInstance := pinger.Pinger{}
	tgInstance, err := tg.NewTgAdapter(configInst.TgToken)
	if err != nil {
		panic(fmt.Sprintf("init tg err: %v", err))
	}
	go tgInstance.Listen()

	ctx = context.WithValue(ctx, "config", configInst)
	ctx = context.WithValue(ctx, "tg", tgInstance)
	ctx = context.WithValue(ctx, "pinger", pingerInstance)

	app.Run(ctx)
}
