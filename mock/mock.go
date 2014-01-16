package mock

import (
	"math/rand"
)

type Mocker interface {
	Mock()
}

func Name() string {
	return Names[rand.Int() % len(Names)]
}





