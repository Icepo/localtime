package localtime

import "testing"

func TestLocalNow(t *testing.T) {
	t.Logf("%v", LocalNow())
}

func TestLocalNowS(t *testing.T) {
	t.Logf("%v", LocalNowS())
}

func TestLocalNowMs(t *testing.T) {
	t.Logf("%v", LocalNowMs())
}

func TestLocalNowNs(t *testing.T) {
	t.Logf("%v", LocalNowNs())
}

func TestLocalNowFormat(t *testing.T) {
	layouts := []string{Virgule, Minus, Dot, VirguleTime, MinusTime, DotTime}
	for _, layout := range layouts {
		t.Logf("%s\t%v", layout, LocalNowFormat(layout))
	}
}
