
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