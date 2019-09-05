package mocks

// mockgen: https://github.com/golang/mock
// $ go get github.com/golang/mock/gomock
// $ go install github.com/golang/mock/mockgen

//go:generate mockgen -package mocks -source=../caller/car.go -destination=./car.go
