package gtfs

type FareRule struct {
	FareID        string `csv:"fare_id"`
	RouteID       string `csv:"route_id"`
	OriginID      string `csv:"origin_id"`
	DestinationID string `csv:"destination_id"`
	ContainsID    string `csv:"contains_id"`
}
