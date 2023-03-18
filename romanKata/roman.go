package romankata

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
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

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func ConvertToRoman(decimal uint16) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for decimal >= numeral.Value {
			result.WriteString(numeral.Symbol)
			decimal -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToDecimal(roman string) uint16 {

	var result uint16

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		//look ahead to the next symbol
		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			// get the value of the two character string
			value := allRomanNumerals.ValueOf(symbol, nextSymbol)

			if value != 0 {
				result += value
				i++ //moved past two numbers
			} else {
				result += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			result += allRomanNumerals.ValueOf(symbol)
		}
	}

	return result
}

func couldBeSubtractive(index int, symbol uint8, roman string) bool {
	return index+1 < len(roman) && (symbol == 'I' || symbol == 'X' || symbol == 'C')
}
