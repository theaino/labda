package analysis

import (
	"testing"
)

func TestLexTrivial(t *testing.T) {
	body := "$x a b"
	want := []Token{Lambda, Word("x"), Word("a"), Word("b")}
	got := Lex(body)
	for idx, wantedToken := range want {
		gotToken := got[idx]
		if wantedToken != gotToken {
			t.Fatalf("Wanted: %v, got: %v", want, got)
			return
		}
	}
}

func TestLexComplex(t *testing.T) {
	body := "($abc abc 1) he11o_1337!"
	want := []Token{LParen, Lambda, Word("abc"), Word("abc"), Number("1"), RParen, Word("he11o_1337!")}
	got := Lex(body)
	for idx, wantedToken := range want {
		gotToken := got[idx]
		if wantedToken != gotToken {
			t.Fatalf("Wanted: %v, got: %v", want, got)
			return
		}
	}
}

func TestLexString(t *testing.T) {
	body := "\"hello\" 1 .$'a'"
	want := []Token{String("hello"), Number("1"), Dot, Lambda, String("a")}
	got := Lex(body)
	for idx, wantedToken := range want {
		gotToken := got[idx]
		if wantedToken != gotToken {
			t.Fatalf("Wanted: %v, got: %v", want, got)
			return
		}
	}
}

func TestLexStringEscaped(t *testing.T) {
	body := "\"hel\\\"lo\" 1 2 3"
	want := []Token{String("hel\"lo"), Number("1"), Number("2"), Number("3")}
	got := Lex(body)
	for idx, wantedToken := range want {
		gotToken := got[idx]
		if wantedToken != gotToken {
			t.Fatalf("Wanted: %v, got: %v", want, got)
			return
		}
	}
}

func TestLexNumber(t *testing.T) {
	body := "12343basd 53"
	want := []Token{Number("12343"), Word("basd"), Number("53")}
	got := Lex(body)
	for idx, wantedToken := range want {
		gotToken := got[idx]
		if wantedToken != gotToken {
			t.Fatalf("Wanted: %v, got: %v", want, got)
			return
		}
	}
}

