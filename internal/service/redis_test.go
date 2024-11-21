package service

import (
	"regexp"
	"testing"
)

var s = RedisService{}

func TestHelloName(t *testing.T) {
	res := s.Test()

	want := regexp.MustCompile("foobar")
	if !want.MatchString(res) {
		t.Fatalf(`Test error`)
	}
}
