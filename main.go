package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gosuri/uilive"
)

// CityClock defines the time of a specific city
type CityClock struct {
	DateTime string `json:"datetime"`
}

// PrintFields is a struct used as a channel to model the output to print
type PrintFields struct {
	capital  string
	dateTime string
}

var capitals string

func init() {
	flag.StringVar(&capitals, "capitals", "", "Insert the capitals to get time from [Continent/Capital1, Continent/Capital2, etc]")
	flag.Parse()
}

func main() {
	capitalCities := strings.Split(capitals, ",")

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	ch := make(chan PrintFields, len(capitalCities))

	for {
		for _, capital := range capitalCities {
			go printTime(capital, ch)
		}

		pfs := make([]PrintFields, len(capitalCities))

		for k := range pfs {
			pfs[k] = <-ch
		}

		fmt.Fprintln(writer, pfs)
	}

}

// printTime prints the current time of the capital to writer
func printTime(capital string, ch chan<- PrintFields) {

	reply, err := http.Get(fmt.Sprintf("http://worldtimeapi.org/api/timezone/%s", capital))
	if err != nil {
		ch <- PrintFields{}
		return
	}

	body, err := ioutil.ReadAll(reply.Body)
	if err != nil {
		ch <- PrintFields{}
		return
	}

	cc := CityClock{}
	err = json.Unmarshal(body, &cc)
	if err != nil {
		ch <- PrintFields{}
		return
	}

	if cc.DateTime == "" {
		ch <- PrintFields{}
		return
	}

	dateTime := string(cc.DateTime[strings.Index(cc.DateTime, "T")+1 : strings.Index(cc.DateTime, "T")+6])

	pf := PrintFields{capital: capital, dateTime: dateTime}
	ch <- pf
}
