package main

import (
	"flag"
	"log"
	"math/rand"

	"github.com/gobridge-kr/bot-sample/internal"
	"github.com/sbstjn/hanu"
)

const version = "0.0.1"
var tokenContainer = bot.NewContainer("")

func init() {
	var token string
	flag.StringVar(&token, "token", "", "Slack bot API token")
	flag.Parse()

	tokenContainer.Set(token)
	tokenContainer.Freeze()
}

func main() {
	token := tokenContainer.Get().(string)
	if (token == "") {
		log.Fatal("API token not provided.\nPlease try with option -h for usage")
	}

	slack, err := hanu.New(token)
	if err != nil {
		log.Fatal(err)
	}

	slack.Command("^안(녕|뇽|냥).*", func(conv hanu.ConversationInterface) {
		user := conv.Message().User()
		conv.Reply("안녕하세요, %s님!", user)
	})

	slack.Command("[야\\s]*<something:string>\\s+(말해줘?|해봐)", func(conv hanu.ConversationInterface) {
		something, _ := conv.String("something")
		conv.Reply(something)
	})

	slack.Command("^.+?$", func(conv hanu.ConversationInterface) {
		yes := rand.Float64() < 0.5
		if yes {
			conv.Reply("네!")
		} else {
			conv.Reply("아니요!")
		}
	})

	slack.Listen()
}
