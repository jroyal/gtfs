package gtfs

// RouteType describes the type of transportation used on a route
type RouteType int64

var (
	// LightRail Used for Tram, Streetcar, Light rail. Any light rail or street level system within a metropolitan area.
	LightRail RouteType = 0
	//Subway Used for subway, metro. Any underground rail system within a metropolitan area.
	Subway RouteType = 1
	// Rail Used for intercity or long-distance travel
	Rail RouteType = 2
	// Bus Used for short- and long-distance bus routes
	Bus RouteType = 3
	// Ferry Used for short- and long-distance boat service
	Ferry RouteType = 4
	//CableCar Used for street-level cable cars where the cable runs beneath the car
	CableCar RouteType = 5
	//Gondola Used for Gondola, Suspended cable car. Typically used for aerial cable cars where the car is suspended from the cable
	Gondola RouteType = 6
	//Funicular Used for any rail system designed for steep inclines
	Funicular RouteType = 7
)

//Route is a group of trips that are displayed to riders as a single service
type Route struct {
	ID        string    `csv:"route_id"`
	AgencyID  string    `csv:"agency_id"`
	ShortName string    `csv:"route_short_name"`
	LongName  string    `csv:"route_long_name"`
	Desc      string    `csv:"route_desc"`
	Type      RouteType `csv:"route_type"`
	URL       string    `csv:"route_url"`
	Color     string    `csv:"route_color"`
	TextColor string    `csv:"route_text_color"`
	SortOrder int64     `csv:"route_sort_order"`
}
