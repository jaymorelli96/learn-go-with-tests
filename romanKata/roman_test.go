package romankata

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Decimal uint16
	Roman   string
}{
	{Decimal: 1, Roman: "I"},
	{Decimal: 2, Roman: "II"},
	{Decimal: 3, Roman: "III"},
	{Decimal: 4, Roman: "IV"},
	{Decimal: 5, Roman: "V"},
	{Decimal: 6, Roman: "VI"},
	{Decimal: 7, Roman: "VII"},
	{Decimal: 9, Roman: "IX"},
	{Decimal: 10, Roman: "X"},
	{Decimal: 12, Roman: "XII"},
	{Decimal: 14, Roman: "XIV"},
	{Decimal: 18, Roman: "XVIII"},
	{Decimal: 19, Roman: "XIX"},
	{Decimal: 23, Roman: "XXIII"},
	{Decimal: 39, Roman: "XXXIX"},
	{Decimal: 40, Roman: "XL"},
	{Decimal: 47, Roman: "XLVII"},
	{Decimal: 50, Roman: "L"},
	{Decimal: 90, Roman: "XC"},
	{Decimal: 97, Roman: "XCVII"},
	{Decimal: 100, Roman: "C"},
	{Decimal: 500, Roman: "D"},
	{Decimal: 400, Roman: "CD"},
	{Decimal: 900, Roman: "CM"},
	{Decimal: 1000, Roman: "M"},
	{Decimal: 1984, Roman: "MCMLXXXIV"},
	{Decimal: 3999, Roman: "MMMCMXCIX"},
	{Decimal: 2014, Roman: "MMXIV"},
	{Decimal: 1006, Roman: "MVI"},
	{Decimal: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Decimal, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Decimal)

			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToDecimal(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Decimal), func(t *testing.T) {
			got := ConvertToDecimal(test.Roman)

			if got != test.Decimal {
				t.Errorf("got %d, want %d", got, test.Decimal)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(decimal uint16) bool {
		if decimal > 3999 {
			return true
		}
		t.Log("testing", decimal)
		roman := ConvertToRoman(decimal)
		fromRoman := ConvertToDecimal(roman)
		return fromRoman == decimal
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
