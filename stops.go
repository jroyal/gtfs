package gtfs

// Stop is an individual location where vehicles pick up or drop off passengers
type Stop struct {
	ID           string `csv:"stop_id"`
	Code         string `csv:"stop_code"`
	Name         string `csv:"stop_name"`
	Desc         string `csv:"stop_desc"`
	Latitude     string `csv:"stop_lat"`
	Longitude    string `csv:"stop_lon"`
	ZoneID       string `csv:"zone_id"`
	URL          string `csv:"stop_url"`
	LocationType string `csv:"location_type"`
}
