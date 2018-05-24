package gtfs

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

// All the files that make up the GTFS spec
const (
	agencyFile         = "agency.txt"
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
	FeedInfo       *FeedInfo
	Stops          []*Stop
	Routes         []*Route
	Shapes         []*Shape
	Agencies       []*Agency
	Trips          []*Trip
	StopTimes      []*StopTimes
	Calendar       []*Calendar
	CalendarDates  []*CalendarDate
	FareAttributes []*FareAttribute
	FareRules      []*FareRule
	Frequencies    []*Frequency
	Transfers      []*Transfer
}

func NewFeed() Feed {
	return Feed{
		// All the individual fields if you need to get all
		Agencies:       []*Agency{},
		FeedInfo:       &FeedInfo{},
		Stops:          []*Stop{},
		Routes:         []*Route{},
		Trips:          []*Trip{},
		StopTimes:      []*StopTimes{},
		Calendar:       []*Calendar{},
		CalendarDates:  []*CalendarDate{},
		FareAttributes: []*FareAttribute{},
		FareRules:      []*FareRule{},
		Shapes:         []*Shape{},
		Frequencies:    []*Frequency{},
		Transfers:      []*Transfer{},
	}
}

// LoadFromDirectory takes a directory that should contain all the required GTFS files
func LoadFromDirectory(dir string) Feed {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	feed := NewFeed()
	for _, file := range files {
		f, err := os.Open(filepath.Join(dir, file.Name()))
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		parse(file.Name(), f, &feed)
	}
	return feed
}

// LoadFromZip takes a zip file that should contain all the required GTFS files
func LoadFromZip(archive string) Feed {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		log.Fatal(err)
	}
	feed := NewFeed()
	for _, file := range reader.File {
		f, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		parse(file.Name, f, &feed)
	}
	return feed
}

func parse(fName string, f io.Reader, feed *Feed) {
	var err error
	switch fName {
	case feedInfoFile:
		feedInfo := []*FeedInfo{}
		err = gocsv.Unmarshal(f, &feedInfo)
		if len(feedInfo) != 0 {
			feed.FeedInfo = feedInfo[0]
		}
	case agencyFile:
		err = gocsv.Unmarshal(f, &feed.Agencies)
	case stopsFile:
		err = gocsv.Unmarshal(f, &feed.Stops)
	case routesFile:
		err = gocsv.Unmarshal(f, &feed.Routes)
	case tripsFile:
		err = gocsv.Unmarshal(f, &feed.Trips)
	case stopTimesFile:
		err = gocsv.Unmarshal(f, &feed.StopTimes)
	case calendarFile:
		err = gocsv.Unmarshal(f, &feed.Calendar)
	case calendarDatesFile:
		err = gocsv.Unmarshal(f, &feed.CalendarDates)
	case fareAttributesFile:
		err = gocsv.Unmarshal(f, &feed.FareAttributes)
	case fareRulesFile:
		err = gocsv.Unmarshal(f, &feed.FareRules)
	case shapesFile:
		err = gocsv.Unmarshal(f, &feed.Shapes)
	case frequenciesFile:
		err = gocsv.Unmarshal(f, &feed.Frequencies)
	case transfersFile:
		err = gocsv.Unmarshal(f, &feed.Transfers)
	}
	if err != nil {
		log.Fatal(err)
	}
}
