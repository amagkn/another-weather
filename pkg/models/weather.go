package models

type Weather struct {
	City      string  `json:"city"`
	Int       float64 `json:"temperature"`
	UpdatedAt float64 `json:"updated_at"`
}
