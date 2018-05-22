package gtfs

type StopTimes struct {
	TripID            string `csv:"trip_id"`
	ArrivalTime       string `csv:"arrival_time"`
	DepartureTime     string `csv:"departure_time"`
	StopID            string `csv:"stop_id"`
	StopSequence      string `csv:"stop_sequence"`
	StopHeadsign      string `csv:"stop_headsign"`
	PickupType        int64  `csv:"pickup_type"`
	DropOffType       int64  `csv:"drop_off_type"`
	ShapeDistTraveled string `csv:"shape_dist_traveled"`
	Timepoint         int64  `csv:"timepoint"`
}
