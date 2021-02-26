package uniq

import (
	"reflect"
	"testing"
)

func newRunOptions(c, d, u, i bool, f, s int) RunOptions {
	return RunOptions{
		Count:      c,
		Duplicates: d,
		Unique:     u,
		SkipFields: f,
		SkipChars:  s,
		IgnoreCase: i,
	}
}

func TestUniq(t *testing.T) {
	t.Run("Duplicates", func(t *testing.T) {
		lines := []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, false, false, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("All unique", func(t *testing.T) {
		lines := []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, false, false, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Skip symbols", func(t *testing.T) {
		lines := []string{
			"Ia love music.",
			"As love music.",
			"Cd love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"Ia love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, false, false, 0, 2))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Skip fields", func(t *testing.T) {
		lines := []string{
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"We love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, false, false, 1, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Ignore case", func(t *testing.T) {
		lines := []string{
			"I lOvE mUsic.",
			"I loVe musIc.",
			"i love muSic.",
			"",
			"I love MUSIC of Kartik.",
			"I love music OF Kartik.",
			"Thanks.",
		}
		expected := []string{
			"I lOvE mUsic.",
			"",
			"I love MUSIC of Kartik.",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, false, true, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})
}

func TestUniqDuplicates(t *testing.T) {
	t.Run("Duplicates", func(t *testing.T) {
		lines := []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"I love music.",
			"I love music of Kartik.",
		}
		got := Uniq(lines, newRunOptions(false, true, false, false, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("All unique", func(t *testing.T) {
		lines := []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{}
		got := Uniq(lines, newRunOptions(false, true, false, false, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Skip symbols", func(t *testing.T) {
		lines := []string{
			"Ia love music.",
			"As love music.",
			"Cd love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"Ia love music.",
		}
		got := Uniq(lines, newRunOptions(false, true, false, false, 0, 2))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Skip fields", func(t *testing.T) {
		lines := []string{
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"We love music.",
			"I love music of Kartik.",
		}
		got := Uniq(lines, newRunOptions(false, true, false, false, 1, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Ignore case", func(t *testing.T) {
		lines := []string{
			"I lOvE mUsic.",
			"I loVe musIc.",
			"i love muSic.",
			"",
			"I love MUSIC of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"I lOvE mUsic.",
		}
		got := Uniq(lines, newRunOptions(false, true, false, true, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})
}

func TestUniqOnlyUnique(t *testing.T) {
	t.Run("Duplicates", func(t *testing.T) {
		lines := []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, true, false, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("All unique", func(t *testing.T) {
		lines := []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, true, false, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Skip symbols", func(t *testing.T) {
		lines := []string{
			"Ia love music.",
			"As love music.",
			"Cd love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, true, false, 0, 2))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Skip fields", func(t *testing.T) {
		lines := []string{
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, true, false, 1, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Ignore case", func(t *testing.T) {
		lines := []string{
			"I lOvE mUsic.",
			"I loVe musIc.",
			"i love muSic.",
			"",
			"I love MUSIC of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"",
			"I love MUSIC of Kartik.",
			"Thanks.",
		}
		got := Uniq(lines, newRunOptions(false, false, true, true, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})
}

func TestUniqCount(t *testing.T) {
	t.Run("Duplicates", func(t *testing.T) {
		lines := []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"3 I love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
		}
		got := Uniq(lines, newRunOptions(true, false, false, false, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("All unique", func(t *testing.T) {
		lines := []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"1 I love music.",
			"1 ",
			"1 I love music of Kartik.",
			"1 Thanks.",
		}
		got := Uniq(lines, newRunOptions(true, false, false, false, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Skip symbols", func(t *testing.T) {
		lines := []string{
			"Ia love music.",
			"As love music.",
			"Cd love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"3 Ia love music.",
			"1 ",
			"1 I love music of Kartik.",
			"1 Thanks.",
		}
		got := Uniq(lines, newRunOptions(true, false, false, false, 0, 2))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Skip fields", func(t *testing.T) {
		lines := []string{
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"3 We love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
		}
		got := Uniq(lines, newRunOptions(true, false, false, false, 1, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Ignore case", func(t *testing.T) {
		lines := []string{
			"I lOvE mUsic.",
			"I loVe musIc.",
			"i love muSic.",
			"",
			"I love MUSIC of Kartik.",
			"Thanks.",
		}
		expected := []string{
			"3 I lOvE mUsic.",
			"1 ",
			"1 I love MUSIC of Kartik.",
			"1 Thanks.",
		}
		got := Uniq(lines, newRunOptions(true, false, false, true, 0, 0))

		if !reflect.DeepEqual(got, expected) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})
}
