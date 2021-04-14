package main

import (
	"testing"
	"testing/quick"
)

// We are building Roman Numeral Kata
// http://codingdojo.org/kata/RomanNumerals/

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func assertInt(t testing.TB, got, want uint16) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d', want '%d'", got, want)
	}
}

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		ArabicDigit uint16
		Roman       string
	}{
		{ArabicDigit: 1, Roman: "I"},
		{ArabicDigit: 2, Roman: "II"},
		{ArabicDigit: 4, Roman: "IV"},
		{ArabicDigit: 5, Roman: "V"},
		{ArabicDigit: 6, Roman: "VI"},
		{ArabicDigit: 8, Roman: "VIII"},
		{ArabicDigit: 9, Roman: "IX"},
		{ArabicDigit: 10, Roman: "X"},
		{ArabicDigit: 14, Roman: "XIV"},
		{ArabicDigit: 19, Roman: "XIX"},
		{ArabicDigit: 34, Roman: "XXXIV"},
		{ArabicDigit: 40, Roman: "XL"},
		{ArabicDigit: 47, Roman: "XLVII"},
		{ArabicDigit: 49, Roman: "XLIX"},
		{ArabicDigit: 50, Roman: "L"},
		{ArabicDigit: 90, Roman: "XC"},
		{ArabicDigit: 400, Roman: "CD"},
		{ArabicDigit: 500, Roman: "D"},
		{ArabicDigit: 793, Roman: "DCCXCIII"},
		{ArabicDigit: 900, Roman: "CM"},
		{ArabicDigit: 1000, Roman: "M"},
		{ArabicDigit: 1337, Roman: "MCCCXXXVII"},
		{ArabicDigit: 1400, Roman: "MCD"},
		{ArabicDigit: 1992, Roman: "MCMXCII"},
		{ArabicDigit: 3999, Roman: "MMMCMXCIX"},
	}
	for _, test := range cases {
		arabic, err := NewArabic(test.ArabicDigit)

		if err != nil {
			t.Errorf("arabic digit constraints broken")
		}

		got := ConvertToRoman(*arabic)
		assertString(t, got, test.Roman)
	}
}

func TestArabicDigitNumerals(t *testing.T) {
	cases := []struct {
		ArabicDigit uint16
		Roman       string
	}{
		{ArabicDigit: 1, Roman: "I"},
		{ArabicDigit: 2, Roman: "II"},
		{ArabicDigit: 4, Roman: "IV"},
		{ArabicDigit: 5, Roman: "V"},
		{ArabicDigit: 6, Roman: "VI"},
		{ArabicDigit: 8, Roman: "VIII"},
		{ArabicDigit: 9, Roman: "IX"},
		{ArabicDigit: 10, Roman: "X"},
		{ArabicDigit: 14, Roman: "XIV"},
		{ArabicDigit: 19, Roman: "XIX"},
		{ArabicDigit: 34, Roman: "XXXIV"},
		{ArabicDigit: 40, Roman: "XL"},
		{ArabicDigit: 47, Roman: "XLVII"},
		{ArabicDigit: 49, Roman: "XLIX"},
		{ArabicDigit: 50, Roman: "L"},
		{ArabicDigit: 90, Roman: "XC"},
		{ArabicDigit: 400, Roman: "CD"},
		{ArabicDigit: 500, Roman: "D"},
		{ArabicDigit: 793, Roman: "DCCXCIII"},
		{ArabicDigit: 900, Roman: "CM"},
		{ArabicDigit: 1000, Roman: "M"},
		{ArabicDigit: 1337, Roman: "MCCCXXXVII"},
		{ArabicDigit: 1400, Roman: "MCD"},
		{ArabicDigit: 1992, Roman: "MCMXCII"},
		{ArabicDigit: 3999, Roman: "MMMCMXCIX"},
	}
	for _, test := range cases {
		got := ConvertToArabic(test.Roman)
		assertInt(t, got.Val(), test.ArabicDigit)
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(seed uint16) bool {
		arabic, err := NewArabic(seed)
		if err == nil {
			roman := ConvertToRoman(*arabic)
			fromRoman := ConvertToArabic(roman)
			return fromRoman == *arabic
		}
		return true
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
