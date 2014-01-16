package mock

import (
	"math/rand"
	"strings"
)

type Mocker interface {
	Mock()
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



















