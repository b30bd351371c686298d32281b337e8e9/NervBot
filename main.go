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
	// RNG command group
	router.Group(func(r *exrouter.Route) {
		r.Cat(randomCat)

		r.On("coin", func(ctx *exrouter.Context) {
			rand.Seed(time.Now().Unix())
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
