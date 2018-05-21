package main

import (
	"fmt"

	"github.com/jroyal/gtfs/gtfs"
)

func main() {
	feed := gtfs.LoadDirectory("/users/james/downloads/capmetro")
	for _, r := range feed.Routes {
		fmt.Println(r)
	}
}
