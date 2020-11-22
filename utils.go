package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/gosuri/uilive"
	"github.com/kataras/tablewriter"
)

// CityClock defines the time of a specific city. It's used as the json target
type CityClock struct {
	DateTime string `json:"datetime"`
}

// PrintFields is a struct used as a channel to model the output to print
type PrintFields struct {
	capital  string
	dateTime string
}

// getTime writes the current time of capital to the channel
// worldtimeapi.org is queried
func getTime(capital string, ch chan<- PrintFields) {

	reply, err := http.Get(fmt.Sprintf("http://worldtimeapi.org/api/timezone/%s", capital))
	if err != nil {
		ch <- PrintFields{}
		reply.Body.Close()
		return
	}
	defer reply.Body.Close()

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

	ch <- PrintFields{capital: capital, dateTime: dateTime}
}

// printOutput generates an ASCII table on the fly
func printOutput(out map[string]string, writer *uilive.Writer) {
	keys := make([]string, 0, len(out))

	for k := range out {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Capital City", "Time"})

	for _, k := range keys {
		table.Append([]string{k, out[k]})
	}
	table.Render()
}
