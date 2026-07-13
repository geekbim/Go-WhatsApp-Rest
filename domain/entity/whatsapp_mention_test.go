package entity

import (
	"reflect"
	"testing"
)

func TestNormalizeMentionJIDs(t *testing.T) {
	got := NormalizeMentionJIDs([]string{"all", "+62 812-345", "0812345,628999@s.whatsapp.net", "@628999@s.whatsapp.net"})
	want := []string{"62812345@s.whatsapp.net", "628999@s.whatsapp.net"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("NormalizeMentionJIDs() = %#v, want %#v", got, want)
	}

	if !HasMentionAll([]string{"62812345", "@all"}) {
		t.Fatal("HasMentionAll() = false, want true")
	}
}
