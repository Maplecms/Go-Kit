package mock

import (
	"math/rand"
	"strings"
	"time"
	"fmt"
)

type Mocker interface {
	Mock()
}

func IntBetween(a, b int) int {
	size := b - a
	return rand.Int() % size + a
}

func Float32Between(a, b float32) float32 {
	size := b - a
	return rand.Float32() * size + a
}

func Float64Between(a, b float64) float64 {
	size := b - a
	return rand.Float64() * size + a
}

func TimeBetween(a, b time.Time) time.Time {
	durationNanoSeconds := int64(rand.Int())  % b.Sub(a).Nanoseconds()
	durationString := fmt.Sprintf("%dns", durationNanoSeconds)
	duration, _ := time.ParseDuration(durationString)
	return a.Add(duration)
}

func Name() string {
	return Names[rand.Int() % len(Names)]
}

func Domain() string {
	return Domains[rand.Int() % len(Domains)]
}

func Company() string {
	return Companies[rand.Int() % len(Companies)]
}

func CompanySanitized() string {
	return strings.ToLower(strings.Replace(strings.Replace(Company(), ".", "", -1), " ", "", -1))
}

func Email() string {
	return Name() + Name() + "@" + CompanySanitized() + "." + Domain()
}

func OneOf(values ...interface{}) interface{} {
	return values[rand.Int() % len(values)]
}

func CharacterIn(characters string) byte {
	return characters[rand.Int() % len(characters)]
}

func LowerAlpha() byte {
	return CharacterIn("abcdefghijklmnopqrstuvwxyz")
}

func UpperAlpha() byte {
	return CharacterIn("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func Numeric() byte {
	return CharacterIn("0123456789")
}

func Alpha() byte {
	if rand.Int() % 2 == 0 {
		return LowerAlpha()
	} else {
		return UpperAlpha()
	}
}

func AlphaNumeric() byte {
	if rand.Int() % 2 == 0 {
		return Alpha()
	} else {
		return Numeric()
	}
}

func StringWithGenerator(length int, generator func() byte) string {
	result := make([]byte, length)
	for i := 0; i < length; i ++ {
		result[i] = generator()
	}
	return string(result)
}

func LowerString(length int) string {
	return StringWithGenerator(length, LowerAlpha)
}

func UpperString(length int) string {
	return StringWithGenerator(length, UpperAlpha)
}

func AlphaString(length int) string {
	return StringWithGenerator(length, Alpha)
}

func DigitString(length int) string {
	return StringWithGenerator(length, Numeric)
}

func String(length int) string {
	return StringWithGenerator(length, AlphaNumeric)
}

func VersionString() string {
	major := DigitString(2)
	minor := DigitString(2)
	build := DigitString(2)
	return major + "." + minor + "." + build
}

