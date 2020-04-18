
package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	scraper "github.com/NickBrisebois/DiscordServerMessageScraper/scraper"
	"log"
    "fmt"
)

var version string
var gitCommit string
var buildTime string

func main() {
	configPath := flag.String("config", "./config.toml", "Path to config.toml")
    versionFlag := flag.Bool("v", false, "Show version")
	flag.Parse()