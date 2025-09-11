package dto

type GetCityCoordinates struct {
	Lot  string `json:"lon"`
	Lat  string `json:"lat"`
	Name string `json:"name"`
}
