package docker

import "testing"

func TestIsAllowedLogLevel(t *testing.T) {
	ll := LogLevel("wrong-type")
	r := ll.IsAllowed()
	if r != false {
		t.Errorf("expected exec is %q, got %q", "false", "true")
	}

	for _, v := range LogLevelList {
		if v.IsAllowed() != true {
			t.Errorf("expected exec is %q, got %q", "true", "false")
		}
	}
}
