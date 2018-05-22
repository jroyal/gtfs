package gtfs

type Frequency struct {
	TripID      string `csv:"trip_id"`
	StartTime   string `csv:"start_time"`
	EndTime     string `csv:"end_time"`
	HeadwaySecs string `csv:"headway_secs"`
	ExactTimes  string `csv:"exact_times"`
}
