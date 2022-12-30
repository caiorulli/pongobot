package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func run_metrics_server() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

var (
	messagesReceived = promauto.NewCounter(prometheus.CounterOpts{
		Name: "telegram_messages_received_total",
		Help: "The total number of received messages from telegram polling",
	})
	messagesSent = promauto.NewCounter(prometheus.CounterOpts{
		Name: "telegram_messages_sent_total",
		Help: "The total number of messages sent from bots",
	})
)

func main() {

	go run_metrics_server()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	c := cron.New()
	c.AddFunc("@daily", func() {
		fmt.Println("Testing...")
	})
	c.Start()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			messagesReceived.Inc()

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "miau")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
			messagesSent.Inc()
		}
	}
}
