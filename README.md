# gtfs
Parse a [GTFS transit feed](https://developers.google.com/transit/gtfs/reference/)

## Usage

```golang
package main

import (
	"github.com/jroyal/gtfs"
	"fmt"
)

func main() {
	feed := gtfs.LoadFromZip("sample-feed.zip")

	fmt.Printf("There are %d stops on %d routes for %s\n", len(feed.Stops), len(feed.Routes), feed.Agencies[0].Name)
}
```

## Limitations
* Doesn't validate the incoming feed. It just parses it out.

