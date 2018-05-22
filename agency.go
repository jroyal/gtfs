package gtfs

type Agency struct {
	ID       string `csv:"agency_id"`
	Name     string `csv:"agency_name"`
	URL      string `csv:"agency_url"`
	TimeZone string `csv:"agency_timezone"`
	Lang     string `csv:"agency_lang"`
	Phone    string `csv:"agency_phone"`
	FareURL  string `csv:"agency_fare_url"`
	Email    string `csv:"agency_email"`
}
