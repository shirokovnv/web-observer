package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/shirokovnv/web-observer/src/client"
	"gopkg.in/yaml.v3"
)

var (
	// YML configuration must be injected during the build process
	YmlSiteConfig  string
	YmlSlackConfig string

	// Main script configuration
	SiteConfig  client.SiteConfig
	SlackConfig client.SlackConfig
)

func parseArgs() string {

	helpPtr := flag.Bool("help", false, "")
	flag.Parse()

	if *helpPtr {
		return "help"
	}

	return "start"
}

func parseConfigs() {
	configs := map[string]interface{}{
		YmlSiteConfig:  &SiteConfig,
		YmlSlackConfig: &SlackConfig,
	}

	for yml, conf := range configs {
		err := yaml.Unmarshal([]byte(yml), conf)
		if err != nil {
			fmt.Printf("Error parsing YAML file: %s\n", err)
			return
		}
	}
}

func startScheduler() {
	scheduler := gocron.NewScheduler(time.UTC)

	fmt.Println("New scheduler created.")

	for _, element := range SiteConfig.Sites {

		task := func(site client.Site) {
			defer func() {
				if err := recover(); err != nil {
					log.Println("panic occurred:", err)
				}
			}()

			statusCode, err := site.CheckUrlStatus()

			if err != nil {
				log.Fatal(err)
			}

			if int(site.Code) != statusCode {
				sr := client.SlackRequest{
					Text:      fmt.Sprintf("%s %s , %s: %d", site.URL, "is unreachable", "code", statusCode),
					IconEmoji: ":alert:",
				}

				err := SlackConfig.SendSlackNotification(sr)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		_, err := scheduler.Every(element.Interval).Do(task, element)
		if err != nil {
			log.Fatal("error scheduling task")
		}
	}

	scheduler.StartBlocking()

	log.Println("Finished.")
}

func printHelp() {
	log.Printf("%s \n %s",
		"Website Observer.",
		"Pings web resources in a pre-defined schedule.",
	)
}

func main() {
	if parseArgs() == "help" {
		printHelp()
		return
	}

	parseConfigs()
	startScheduler()
}
