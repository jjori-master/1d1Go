package caller

import "math/rand"

type CarInterface interface {
	SpeedUp() int
}

type BMW struct {
	speed int
}

func (b *BMW) SpeedUp() int {
	if rand.Int()%4 == 0 {
		return 0
	}
	b.speed++
	return b.speed
}

type Highway struct {
	Cars []CarInterface
}

func (h *Highway) DoJob() bool {
	for _, car := range h.Cars {
		if car.SpeedUp() > 5 {
			return false
		}
	}
	return true
}

func (h *Highway) ValidateCar() bool {
	for _, car := range h.Cars {
		if car.SpeedUp() > 5 {
			return false
		}
	}
	return true
}

func Real() bool {
	h := Highway{
		Cars: []CarInterface{
			&BMW{},
			&BMW{},
			&BMW{},
			&BMW{},
			&BMW{},
		},
	}

	for i := 0; i < 10; i++ {
		h.DoJob()
	}

	return h.ValidateCar()
}

func RealWithDependencyInjection(cars []CarInterface) bool {
	h := Highway{
		Cars: cars,
	}

	for i := 0; i < 10; i++ {
		h.DoJob()
	}

	return h.ValidateCar()
}
