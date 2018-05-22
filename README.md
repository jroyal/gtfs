# gtfs
Parse a [GTFS transit feed](https://developers.google.com/transit/gtfs/reference/)

## Usage

```golang
import (
	"github.com/jroyal/gtfs"
	"fmt"
)

func main() {
	feed := gtfs.LoadFromZip("sample-feed.zip")

	fmt.Printf("There are %d stops on %d routes for %s", len(feed.Stops), len(feed.Routes), feed.Agency.Name)
}
```

## Limitations
* Doesn't validate the incoming feed. It just parses it out.

