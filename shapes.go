package gtfs

type Shape struct {
	ID           string `csv:"shape_id"`
	Latitude     string `csv:"shape_pt_lat"`
	Longitude    string `csv:"shape_pt_lon"`
	Sequence     string `csv:"shape_pt_sequence"`
	DistTraveled string `csv:"shape_dist_traveled"`
}
