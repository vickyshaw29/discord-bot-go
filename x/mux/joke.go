package mux

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

func (m *Mux) Joke(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Test Stream")
	res, _ := client.Do(req)

	data, _ := ioutil.ReadAll(res.Body)
	var obj JokeResponse
	err := json.Unmarshal(data, &obj)
	if err != nil {
		log.Panic(err)
	}

	_, err = ds.ChannelMessageSend(dm.ChannelID, obj.Joke)
	if err != nil {
		fmt.Print(err)
	}
	return
}
