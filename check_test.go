package domain

import (
	"strings"
	"testing"
)

func TestCheck(t *testing.T) {
	if err := Check(""); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("example.com"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := Check("EXAMPLE.com"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := Check("foo-bar.com"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := Check("www1.foo-bar.com"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := Check("192.168.1.1.example.com"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := Check(strings.Repeat("a", 300)); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check(strings.Repeat("a", 70) + ".com"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("example.com" + strings.Repeat("a", 70)); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("?"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("\t"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("exàmple.com"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("www.\xbd\xb2.com"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("-example.com"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("example-.com"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("example.-com"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("example.com-"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("example.1com"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check(".example.com"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("example..com"); err == nil {
		t.Errorf("unexpected nil error")
	}
	if err := Check("example.com."); err == nil {
		t.Errorf("unexpected nil error")
	}
}
