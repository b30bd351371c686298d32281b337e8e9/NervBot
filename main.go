package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	botPrefix = "^"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	s, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	router := exrouter.New()

	router.On("ping", func(ctx *exrouter.Context) {
		ctx.Reply("pong")
	}).Desc("Responds with pong.")

	router.Default = router.On("help", func(ctx *exrouter.Context) {
		var text = ""
		for _, v := range router.Routes {
			text += v.Name + " : \t" + v.Description + "\n"
		}
		ctx.Reply("```" + text + "```")
	}).Desc("Prints this help menu.")

	s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		router.FindAndExecute(s, botPrefix, s.State.User.ID, m.Message)
	})

	err = s.Open()
	defer s.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
