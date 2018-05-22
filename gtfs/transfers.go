package gtfs

type Transfer struct {
	FromStopID      string `csv:"from_stop_id"`
	ToStopID        string `csv:"to_stop_id"`
	TransferType    int64  `csv:"transfer_type"`
	MinTransferTime int64  `csv:"min_transfer_time"`
}
