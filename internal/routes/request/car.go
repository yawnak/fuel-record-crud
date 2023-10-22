package request

type CreateCar struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int32  `json:"year"`
}
