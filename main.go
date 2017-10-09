package main

import (
	"log"

	"github.com/shomali11/slacker"
)

func main() {
	log.Println("Starting SlackBot")
	bot := slacker.NewClient("YOUR-TOKEN-HERE")

	//bot commands
	bot.Command("status", "Display status information about the bot", handleStatus)
	bot.Command("info <coin>", "Get information about coin", handleCoinInfo)
	bot.Command("price <coin>", "Get price of a coin", handleCoinPrice)

	err := bot.Listen()
	if err != nil {
		log.Fatal(err)
	}
}

//status command to see if bot is online
func handleStatus(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply("I'm ok. Thanks for asking")
}

//info <coin> command
func handleCoinInfo(request *slacker.Request, response slacker.ResponseWriter) {
	coin := request.Param("info")
	coinInfo, err := GetCoinInfo(coin)
	if err != nil {
		log.Println("GetCoinInfo: ", err)
	}
	response.Reply(coinInfo)
}

//price <coin> info
func handleCoinPrice(request *slacker.Request, response slacker.ResponseWriter) {
	coin := request.Param("price")
	coinPrice, err := GetCoinPrice(coin)
	if err != nil {
		log.Println("GetCoinInfo: ", err)
	}
	response.Reply(coinPrice)
}
