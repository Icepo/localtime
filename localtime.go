package localtime

import "time"

const (
	defaultLocation = "Asia/Chongqing"

	Virgule = "2006/01/02"
	Minus   = "2006-01-02"
	Dot     = "2006.01.02"

	VirguleTime = Virgule + " 15:04:05"
	MinusTime   = Minus + " 15:04:05"
	DotTime     = Dot + " 15:04:05"
)

var location *time.Location

func init() {
	var err error
	location, err = time.LoadLocation(defaultLocation)
	if err != nil {
		panic(err)
	}
}

func LocalNow() time.Time {
	return time.Now().In(location)
}

func LocalNowNs() int64 {
	return LocalNow().UnixNano()
}

func LocalNowMs() int64 {
	return LocalNow().UnixNano() / 1e6
}

func LocalNowS() int64 {
	return LocalNow().Unix()
}

func LocalNowFormat(layout string) string {
	return LocalNow().Format(layout)
}
