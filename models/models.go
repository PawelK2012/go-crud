package models

type Menu struct {
	Breakfast      string `json:"breakfast"`
	LargeBreakfast string `json:"large-reakfast"`
	Lunch          string `json:"lunch"`
	LargeLunch     string `json:"large-lunch"`
	Dinner         string `json:"dinner"`
	LargeDinner    string `json:"large-dinner"`
	KidsMenu       string `json:"kids-menu"`
	Desert         string `json:"desert"`
	Drink          string `json:"drink"`
	Sides          string `json:"sides"`
}
