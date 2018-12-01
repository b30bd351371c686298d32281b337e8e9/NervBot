package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/asaskevich/govalidator"
	"github.com/bwmarrin/discordgo"
	"github.com/swoldemi/nervbot/storage/datastore"
)

const (
	botPrefix = "^"
	otherCat  = "Other/Misc."
	randomCat = "Random"
	project   = "nervbot-222113"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	storageClient, err := datastore.New(context.Background(), project)
	check(err)
	token, err := storageClient.Get("Secret", "token")
	check(err)

	s, err := discordgo.New("Bot " + token)
	check(err)

	router := exrouter.New()

	// Other/Misc. command group
	router.Group(func(r *exrouter.Route) {
		r.Cat(otherCat)

		router.On("ping", func(ctx *exrouter.Context) {
			ctx.Reply("pong")
		}).Desc("Responds with pong.").Cat(otherCat)
	})

	// RNG command group
	router.Group(func(r *exrouter.Route) {
		r.Cat(randomCat)

		r.On("coin", func(ctx *exrouter.Context) {
			rand.Seed(time.Now().Unix())
			coin := make([]string, 0)
			coin = append(coin, "head", "tail")
			message := fmt.Sprint(coin[rand.Intn(len(coin))])
			ctx.Reply(message)
		}).Desc("Flips a coin.").Cat(randomCat)

		r.On("number", func(ctx *exrouter.Context) {
			if length := len(ctx.Args); length == 1 {
				ctx.Reply("Please provide a number for for the RNG.")
			} else if govalidator.IsInt(ctx.Args[1]) {
				input, err := strconv.Atoi(ctx.Args[1])
				if err != nil {
					message := fmt.Sprintf("Unable to convert %v to a number.", input)
					ctx.Reply(message)
				}
				ctx.Reply(rand.Intn(input))
			} else {
				ctx.Reply(ctx.Args[1], " is not a number.")
			}

		}).Desc("Generates a random number.").Cat(randomCat)
	})

	// Help command
	router.Default = router.On("help", func(ctx *exrouter.Context) {
		var text = ""
		for _, v := range router.Routes {
			text += v.Name + " : \t" + v.Description + " (" + v.Category + ")\n"
		}
		ctx.Reply("```" + text + "```")
	}).Desc("Prints this help menu.").Cat(otherCat)

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
