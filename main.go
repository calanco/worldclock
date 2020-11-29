package main

import (
	"flag"
	"strings"
	"time"

	"github.com/calanco/worldclock/internal/utils"
	"github.com/gosuri/uilive"
)

var cities string

func init() {
	flag.StringVar(&cities, "cities", "", "Insert the cities to get time from [Continent1/City1, Continent2/City2, etc]")
	flag.Parse()
}

func main() {
	// Split cities parameter to get a list of the cities
	cs := strings.Split(cities, ",")

	// Create a writer to refresh the printed output
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	ch := make(chan utils.PrintFields, len(cs))

	// Mapping cities with their last revealed time
	out := make(map[string]string)

	for {
		// Triggering concurrent goroutines to get the times of all requested cities
		for _, c := range cs {
			go utils.GetTime(c, ch)
		}

		// Waiting for the channel to be empty and checking if the news are valid times
		for i := 0; i < len(cs); i++ {
			tempPf := <-ch

			if tempPf.DateTime != "" {
				out[tempPf.City] = tempPf.DateTime
			}
		}

		utils.PrintOutput(out, writer)
		time.Sleep(2 * time.Second)
	}

}
