package caller

import (
	"1d1Go/cmd/sandbox/mocks/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)


func TestCar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	car := mocks.NewMockCarInterface(ctrl)
	h := Highway{
		[]CarInterface{
			car,
		},
	}

	car.EXPECT().SpeedUp().Return(4)
	if !h.DoJob() {
		t.Fail()
	}
	car.EXPECT().SpeedUp().Return(5)
	if !h.ValidateCar() {
		t.Fail()
	}

	car.EXPECT().SpeedUp().Return(5)
	if !h.DoJob() {
		t.Fail()
	}
	car.EXPECT().SpeedUp().Return(6)
	if h.ValidateCar() {
		t.Fail()
	}

	car.EXPECT().SpeedUp().Return(6)
	if h.DoJob() {
		t.Fail()
	}

	car.EXPECT().SpeedUp().Times(11).Return(5)
	if !RealWithDependencyInjection([]CarInterface{car}) {
		t.Fail()
	}
	car.EXPECT().SpeedUp().Times(11).Return(6)
	if RealWithDependencyInjection([]CarInterface{car}) {
		t.Fail()
	}
}