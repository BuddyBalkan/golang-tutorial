package main

import(
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!123456", "654321!"},
	}

	for _, tc := range testcases {
		rev, errInTC := Reverse(tc.in)
		if errInTC != nil{
			return
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}


func FuzzReverse(f *testing.F) {
	testcases := []string {"Hello, world!", "  ", "1234567!"}

	for _, tc := range testcases {
		// use f.Add to provide a seed corpus
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, revErr := Reverse(orig)
		// nil 判断
		if revErr != nil { return }
		
		doubleRev, doubleRevErr := Reverse(rev)
		// nil 判断
		if doubleRevErr != nil { return }
		// log
		t.Logf("Number of runes: orig:%d, rev:%d, doubleRev:%d.", 
			utf8.RuneCountInString(orig),
			utf8.RuneCountInString(rev),
			utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}