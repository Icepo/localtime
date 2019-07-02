package localtime

import "testing"

func TestLocalTime_Now(t *testing.T) {
	t.Logf("%v", NewLocalTime().Now())
}

func TestLocalTime_Location(t *testing.T) {
	loc := NewLocalTime()
	t.Logf("%v", loc.Now())
	err := loc.Location("UTC")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", loc.Now())
}

func TestLocalTime_LocalNowS(t *testing.T) {
	t.Logf("%v", NewLocalTime().NowS())
}

func TestLocalTime_NowMs(t *testing.T) {
	t.Logf("%v", NewLocalTime().NowMs())
}

func TestLocalTime_NowNs(t *testing.T) {
	t.Logf("%v", NewLocalTime().NowNs())
}

func TestLocalTime_NowFormat(t *testing.T) {
	layouts := []string{Virgule, Minus, Dot, VirguleTime, MinusTime, DotTime}
	for _, layout := range layouts {
		t.Logf("%s\t%v", layout, NewLocalTime().NowFormat(layout))
	}
}
