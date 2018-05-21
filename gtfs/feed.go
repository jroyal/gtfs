package gtfs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

// All the files that make up the GTFS spec
const (
	agencyFile         = "agenct.txt"
	stopsFile          = "stops.txt"
	routesFile         = "routes.txt"
	tripsFile          = "trips.txt"
	stopTimesFile      = "stop_times.txt"
	calendarFile       = "calendar.txt"
	calendarDatesFile  = "calender_dates.txt"
	fareAttributesFile = "fare_attributes.txt"
	fareRulesFile      = "fare_rules.txt"
	shapesFile         = "shapes.txt"
	frequenciesFile    = "frequencies.txt"
	transfersFile      = "transfers.txt"
	feedInfoFile       = "feed_info.txt"
)

// Feed contains all the individual parts of a GTFS Feed
type Feed struct {
	Agency   []*Agency
	FeedInfo []*FeedInfo
	Stops    []*Stop
	Routes   []*Route
}

func parse(fName string, out interface{}) {
	file, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	err = gocsv.UnmarshalFile(file, out)
	if err != nil {
		log.Fatal(err)
	}
}

// LoadDirectory takes a directory that should contain all the required GTFS files
func LoadDirectory(dir string) Feed {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	feed := Feed{
		Agency:   []*Agency{},
		FeedInfo: []*FeedInfo{},
		Stops:    []*Stop{},
		Routes:   []*Route{},
	}

	for _, f := range files {
		fmt.Println(f.Name())
		path := filepath.Join(dir, f.Name())
		switch f.Name() {
		case agencyFile:
			parse(path, &feed.Agency)
		case feedInfoFile:
			parse(path, &feed.FeedInfo)
		case stopsFile:
			parse(path, &feed.Stops)
		case routesFile:
			parse(path, &feed.Routes)
		}
	}
	return feed
}
