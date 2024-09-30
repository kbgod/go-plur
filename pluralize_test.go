package plur

import "testing"

func TestTextUkrainianNoun(t *testing.T) {
	one := "гривня"
	two := "гривні"
	many := "гривень"

	t.Run("1 гривня", func(t *testing.T) {
		v := 1
		want := one
		got := Text(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("2 гривні", func(t *testing.T) {
		v := 2
		want := two
		got := Text(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("5 гривень", func(t *testing.T) {
		v := 5
		want := many
		got := Text(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestTextUkrainianVerb(t *testing.T) {
	one := "гравець переміг"
	two := "гравці перемогли"
	many := "гравців перемогли"

	t.Run("1 гравець переміг", func(t *testing.T) {
		v := 1
		want := one
		got := Text(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("2 гравці перемогли", func(t *testing.T) {
		v := 2
		want := two
		got := Text(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("5 гравців перемогли", func(t *testing.T) {
		v := 5
		want := many
		got := Text(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("11 гравців перемогли", func(t *testing.T) {
		v := 11
		want := many
		got := Text(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("112 гравців перемогли", func(t *testing.T) {
		v := 112
		want := many
		got := Text(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

func TestNumberText(t *testing.T) {
	one := "гривня"
	two := "гривні"
	many := "гривень"

	t.Run("1 гривня", func(t *testing.T) {
		v := 1
		want := "1 гривня"
		got := NumberText(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("2 гривні", func(t *testing.T) {
		v := 2
		want := "2 гривні"
		got := NumberText(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("12 гривень", func(t *testing.T) {
		v := 12
		want := "12 гривень"
		got := NumberText(v, one, two, many)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
