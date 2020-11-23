package main

import (
	"flag"
	"strings"
	"time"

	"github.com/calanco/worldclock/internal/utils"
	"github.com/gosuri/uilive"
)

var capitals string

func init() {
	flag.StringVar(&capitals, "capitals", "", "Insert the capitals to get time from [Continent/Capital1, Continent/Capital2, etc]")
	flag.Parse()
}

func main() {
	// Split capitals parameter to get a list of the capitals
	capitalCities := strings.Split(capitals, ",")

	// Create a writer to refresh the printed output
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	ch := make(chan utils.PrintFields, len(capitalCities))

	// Mapping capitals with their last revealed time
	out := make(map[string]string)

	for {
		// Triggering concurrent goroutines to get the times of all requested capitals
		for _, capital := range capitalCities {
			go utils.GetTime(capital, ch)
		}

		// Waiting for the channel to be empty and checking if the news are valid times
		for i := 0; i < len(capitalCities); i++ {
			tempPf := <-ch

			if tempPf.DateTime != "" {
				out[tempPf.Capital] = tempPf.DateTime
			}
		}

		utils.PrintOutput(out, writer)
		time.Sleep(2 * time.Second)
	}

}
