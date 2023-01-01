
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