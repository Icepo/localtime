package localtime

import "time"

const (
	defaultLocStr = "Asia/Chongqing"

	Virgule = "2006/01/02"
	Minus   = "2006-01-02"
	Dot     = "2006.01.02"

	VirguleTime = Virgule + " 15:04:05"
	MinusTime   = Minus + " 15:04:05"
	DotTime     = Dot + " 15:04:05"
)

var defaultLoc *time.Location

func init() {
	var err error
	defaultLoc, err = time.LoadLocation(defaultLocStr)
	if err != nil {
		panic(err)
	}
}

type LocalTime struct {
	location *time.Location
}

func NewLocalTime() *LocalTime {
	return &LocalTime{location: defaultLoc}
}

func (t *LocalTime) Location(loc string) error {
	location, err := time.LoadLocation(loc)
	if err != nil {
		return err
	}
	t.location = location
	return nil
}

func (t *LocalTime) Now() time.Time {
	return time.Now().In(t.location)
}

func (t *LocalTime) NowNs() int64 {
	return t.Now().UnixNano()
}

func (t *LocalTime) NowMs() int64 {
	return t.Now().UnixNano() / 1e6
}

func (t *LocalTime) NowS() int64 {
	return t.Now().Unix()
}

func (t *LocalTime) NowFormat(layout string) string {
	return t.Now().Format(layout)
}
