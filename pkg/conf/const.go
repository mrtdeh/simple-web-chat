package conf

const (
	UPLOAD_DIR    = "/var/www/uploads/"
	UPLOAD_SERVER = "http://localhost:8005/"
)

const (
	DirectionPrev int32 = iota
	DirectionNext
	DirectionLast
	DirectionNone
)
