package main

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
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

func ConvertToRoman(digit int) string {
	var roman strings.Builder

	if digit == 4 {
		return "IV"
	}

	for _, numeral := range romanNumerals {
		for digit >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			digit -= numeral.Value
		}
	}
	return roman.String()
}

func isSubtractiveSymbol(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

// func ConvertToArabic(roman string) (arabic int) {
// 	for i := 0; i < len(roman); i++ {
// 		// When we index strings, we get bytes.
// 		symbol := roman[i]
//
// 		if couldBeSubtractive(i, symbol, roman) {
// 			nextSymbol := roman[i+1]
// 			value, found := romanNumerals.ValueOf(symbol, nextSymbol)
//
// 			if found {
// 				arabic += value
// 				i++
// 			} else {
// 				// it's a regular and lonely I, or X or C
// 				value, _ := romanNumerals.ValueOf(symbol)
// 				arabic += value
// 			}
// 		} else {
// 			value, _ := romanNumerals.ValueOf(symbol)
// 			arabic += value
// 		}
// 	}
// 	return
// }
func ConvertToArabic(roman string) (arabic int) {
	for _, symbol := range windowedRoman(roman) {
		arabic += romanNumerals.ValueOf(symbol...)
		// TODO: Understand how these ... work
	}
	return
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
			// TODO: study how these instantiations work
			// raw bytes datatypes still a mistery
		}
	}
	return
}

// func OldConvert(digit int) string {
// 	for digit > 0 {
// 		switch {
// 		case digit > 9:
// 			roman.WriteString("X")
// 			digit -= 10
// 		case digit > 8:
// 			roman.WriteString("IX")
// 			digit -= 9
// 		case digit > 4:
// 			roman.WriteString("V")
// 			digit -= 5
// 		case digit > 3:
// 			roman.WriteString("IV")
// 			digit -= 4
// 		default:
// 			roman.WriteString("I")
// 			digit--
// 		}
// 	}
// 	return roman.String()
// }
