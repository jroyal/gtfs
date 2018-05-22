package gtfs

type Calendar struct {
	ServiceID string `csv:"service_id"`
	Monday    int64  `csv:"monday"`
	Tuesday   int64  `csv:"tuesday"`
	Wednesday int64  `csv:"wednesday"`
	Thursday  int64  `csv:"thursday"`
	Friday    int64  `csv:"friday"`
	Saturday  int64  `csv:"saturday"`
	Sunday    int64  `csv:"sunday"`
	StartDate string `csv:"start_date"`
	EndDate   string `csv:"end_date"`
}

type CalendarDate struct {
	ServiceID     string `csv:"service_id"`
	Date          string `csv:"date"`
	ExceptionType int64  `csv:"exception_type"`
}
