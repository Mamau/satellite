package docker

import "testing"

func TestIsAllowed(t *testing.T) {
	ex := Exec("wrong-type")
	r := ex.IsAllowed()
	if r != false {
		t.Errorf("expected exec is %q, got %q", "false", "true")
	}

	for _, v := range List {
		if v.IsAllowed() != true {
			t.Errorf("expected exec is %q, got %q", "true", "false")
		}
	}
}
