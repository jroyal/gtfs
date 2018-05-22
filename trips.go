package gtfs

type Trip struct {
	ID                   string `csv:"trip_id"`
	ServiceID            string `csv:"service_id"`
	RouteID              string `csv:"route_id"`
	HeadSign             string `csv:"trip_headsign"`
	ShortName            string `csv:"trip_short_name"`
	DirectionID          int64  `csv:"direction_id"`
	BlockID              string `csv:"block_id"`
	ShapeID              string `csv:"shape_id"`
	WheelChairAccessible int64  `csv:"wheelchair_accessible"`
	BikesAllowed         int64  `csv:"bikes_allowed"`
}
