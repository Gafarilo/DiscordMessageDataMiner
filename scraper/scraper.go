
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

	// Get an array of text channels
	var textChannels []*discordgo.Channel
	for _, guild := range sc.sesh.State.Guilds {
		channels, _ := sc.sesh.GuildChannels(guild.ID)
		for _, c := range channels {
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}else {
				textChannels = append(textChannels, c)
			}

		}
	}

	dumpPath := "./dump"
	os.Mkdir(dumpPath, os.ModePerm)

	var wg sync.WaitGroup
	for _, channel := range textChannels {
		log.Printf("Starting dump for %s\n", channel.Name)
		wg.Add(1)
		go sc.BulkDownloadMessages(&wg, channel, dumpPath)
	}
	wg.Wait()

	sc.sesh.Close()
	return nil
}

func (sc *ServerScraper) BulkDownloadMessages(wg *sync.WaitGroup, channel *discordgo.Channel, dumpPath string) {
	defer wg.Done()
	var messages []*discordgo.Message
	var err error

	// Create dump file to write to
	var dumpFile *os.File
	dumpFile, err = os.Create(dumpPath + "/dump-" + channel.Name + ".txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dumpFile.Close()
	dumpWriter := bufio.NewWriter(dumpFile)

	doneDumping := false
	beforeID := ""
	for !doneDumping {
		// Get all the messages we can (max is a limit per API call)
		messages, err = sc.sesh.ChannelMessages(channel.ID, 100, beforeID, "", "")
		if err != nil {
			log.Fatal(err.Error())
		}

		if len(messages) == 0 {
			doneDumping = true
			break
		}

		// Loop through all the messages we fetched
		for _, msg := range messages {
			// Grab the last ID to get more messages from before
			beforeID = msg.ID
			dumpWriter.WriteString(msg.Content + "\n")
		}
	}

	log.Printf("Done dump for %s\n", channel.Name)
	dumpWriter.Flush()
}

