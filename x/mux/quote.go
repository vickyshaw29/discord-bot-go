package mux

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type QuoteResponse struct {
	Content string `json:"content"`
}

func (m *Mux) Quote(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.quotable.io/random", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Test Stream")
	res, _ := client.Do(req)

	data, _ := ioutil.ReadAll(res.Body)
	var obj QuoteResponse
	err := json.Unmarshal(data, &obj)
	if err != nil {
		log.Panic(err)
	}

	_, err = ds.ChannelMessageSend(dm.ChannelID, obj.Content)
	if err != nil {
		fmt.Print(err)
	}
	return
}
