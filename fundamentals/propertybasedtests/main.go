package main

import (
	"errors"
	"strings"
)

type Arabic struct {
	value uint16
}

func (a *Arabic) Val() uint16 {
	return a.value
}

func NewArabic(value uint16) (*Arabic, error) {
	if value > 3999 {
		return nil, errors.New("Invalid value. Valid range: 1 <= v < 3999.")
	}
	return &Arabic{value: value}, nil
}

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

var romanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(digit Arabic) string {
	var roman strings.Builder
	total := digit.Val()

	for _, numeral := range romanNumerals {
		for total >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			total -= numeral.Value
		}
	}
	return roman.String()
}

func isSubtractiveSymbol(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

func ConvertToArabic(roman string) Arabic {
	var arabic uint16

	for _, symbol := range windowedRoman(roman) {
		arabic += romanNumerals.ValueOf(symbol...)
		// TODO: Understand how these ... work
	}
	nArabic, _ := NewArabic(arabic)
	return *nArabic
}

// Iterates over a string of romans and neatly deals with subtractive symbols:
//	1. finds and
//	2. properly ties them as ONE symbol and
//  3. and move the iteration forward
func windowedRoman(roman string) (symbols [][]byte) {
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		notAtEnd := i+1 < len(roman)
		if notAtEnd && isSubtractiveSymbol(symbol) && romanNumerals.Exists(symbol, roman[i+1]) {
			nextSymbol := roman[i+1]
			symbols = append(symbols, []byte{byte(symbol), byte(nextSymbol)})
			i++
		} else {
			symbols = append(symbols, []byte{byte(symbol)})
		}
	}
	return
}
