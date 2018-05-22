package gtfs

type FareAttribute struct {
	FareID           string `csv:"fare_id"`
	Price            string `csv:"price"`
	CurrencyType     string `csv:"currency_type"`
	PaymentMethod    int64  `csv:"payment_method"`
	Transfers        int64  `csv:"transfers"`
	AgencyID         string `csv:"agency_id"`
	TransferDuration string `csv:"transfer_duration"`
}
