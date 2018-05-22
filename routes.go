package gtfs

type Route struct {
	ID        string `csv:"route_id"`
	AgencyID  string `csv:"agency_id"`
	ShortName string `csv:"route_short_name"`
	LongName  string `csv:"route_long_name"`
	Desc      string `csv:"route_desc"`
	Type      string `csv:"route_type"`
	URL       string `csv:"route_url"`
	Color     string `csv:"route_color"`
	TextColor string `csv:"route_text_color"`
	SortOrder int64  `csv:"route_sort_order"`
}
