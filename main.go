package main

import (
	"fmt"

	"github.com/jroyal/gtfs/gtfs"
)

func main() {
	feed := gtfs.LoadFromZip("/users/james/downloads/capmetro.zip")
	for _, r := range feed.Routes {
		fmt.Println(r)
	}
}
