package gtfs

type FeedInfo struct {
	PublisherName string `csv:"feed_publisher_name"`
	PublisherURL  string `csv:"feed_publisher_url"`
	Lang          string `csv:"feed_lang"`
	StartDate     string `csv:"feed_start_date"`
	EndDate       string `csv:"feed_end_date"`
	Version       string `csv:"feed_version"`
}
