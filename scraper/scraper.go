
package discord_scraper

import (
	"bufio"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"sync"
)

type ServerScraper struct {
	botConf *Config
	sesh *discordgo.Session
}

func NewServerScraper(config *Config) *ServerScraper {
	return &ServerScraper{
		botConf: config,
	}
}

func (sc *ServerScraper) InitScraper() error {
	log.Println("Initializing Discord Server Scraper")

	var err error
	sc.sesh, err = discordgo.New("Bot " + sc.botConf.DiscordToken)

	if err != nil {
		return err
	}

	err = sc.sesh.Open()
	if err != nil {
		return err
	}