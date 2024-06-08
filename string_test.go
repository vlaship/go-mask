package mask

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMaskShortString(t *testing.T) {
	result := String("123")
	expected := maskedPart
	assert.Equal(t, expected, result, "String should return ***** for short input")
}

func TestMaskLongString(t *testing.T) {
	result := String("12345abcde")
	expected := "12" + maskedPart + "de"
	assert.Equal(t, expected, result, "String should mask the middle part of a long input")
}

func TestMaskEdge5Case(t *testing.T) {
	result := String("12345")
	expected := maskedPart
	assert.Equal(t, expected, result, "String should return ***** for input of length 5")
}

func TestMaskEdge6Case(t *testing.T) {
	result := String("123456")
	expected := "12" + maskedPart + "56"
	assert.Equal(t, expected, result, "String should return ***** for input of length 6")
}

func TestMaskEmptyString(t *testing.T) {
	result := String("")
	expected := maskedPart
	assert.Equal(t, expected, result, "String should return ***** for an empty input")
}

func TestMaskShortMaskedString(t *testing.T) {
	result := String("*****")
	expected := maskedPart
	assert.Equal(t, expected, result, "String should return ***** for an input already containing *****")
}

func TestMaskShortStringWithSpecialCharacters(t *testing.T) {
	result := String("$%#")
	expected := maskedPart
	assert.Equal(t, expected, result, "String should return ***** for short input with special characters")
}

func TestMaskLongStringWithSpecialCharacters(t *testing.T) {
	input := "ab" + strings.Repeat("$%#", 50) + "xz"
	result := String(input)
	expected := "ab" + maskedPart + "xz"
	assert.Equal(t, expected, result, "String should mask the middle part of a long input with special characters")
}

func TestMaskStrings(t *testing.T) {
	input := []string{"123", "12345abcde", "12345", "123456", "", "*****", "$%#", "ab" + strings.Repeat("$%#", 50) + "xz"}
	result := Strings(input)
	expected := []string{maskedPart, "12" + maskedPart + "de", maskedPart, "12" + maskedPart + "56", maskedPart, maskedPart, maskedPart, "ab" + maskedPart + "xz"}
	assert.Equal(t, expected, result, "Strings should mask the middle part of each string in the slice")
}

func TestMaskRunes(t *testing.T) {
	input := []rune("1234567890")
	result := Runes(input)
	expected := []rune("12" + maskedPart + "90")
	assert.Equal(t, expected, result, "Runes should mask the middle part of a rune slice")
}

func TestMaskBytes(t *testing.T) {
	input := []byte("1234567890")
	result := Bytes(input)
	expected := []byte("12" + maskedPart + "90")
	assert.Equal(t, expected, result, "Bytes should mask the middle part of a byte slice")
}
