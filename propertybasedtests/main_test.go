package main

import "testing"

// We are building Roman Numeral Kata
// http://codingdojo.org/kata/RomanNumerals/

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d', want '%d'", got, want)
	}
}

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Arabic int
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 19, Roman: "XIX"},
		{Arabic: 34, Roman: "XXXIV"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 793, Roman: "DCCXCIII"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1337, Roman: "MCCCXXXVII"},
		{Arabic: 1400, Roman: "MCD"},
		{Arabic: 1992, Roman: "MCMXCII"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
	}
	for _, test := range cases {
		got := ConvertToRoman(test.Arabic)
		assertString(t, got, test.Roman)
	}
}

func TestArabicNumerals(t *testing.T) {
	cases := []struct {
		Arabic int
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 19, Roman: "XIX"},
		{Arabic: 34, Roman: "XXXIV"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 793, Roman: "DCCXCIII"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1337, Roman: "MCCCXXXVII"},
		{Arabic: 1400, Roman: "MCD"},
		{Arabic: 1992, Roman: "MCMXCII"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
	}
	for _, test := range cases {
		got := ConvertToArabic(test.Roman)
		assertInt(t, got, test.Arabic)
	}
}
