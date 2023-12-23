package models

type Menu struct {
	ID             string `json:"id"`
	MenuName       string `json:"menu_name"`
	Breakfast      string `json:"breakfast"`
	LargeBreakfast string `json:"large_breakfast"`
	Lunch          string `json:"lunch"`
	LargeLunch     string `json:"large_lunch"`
	Dinner         string `json:"dinner"`
	LargeDinner    string `json:"large_dinner"`
	KidsMenu       string `json:"kids_menu"`
	Desert         string `json:"desert"`
	Drink          string `json:"drink"`
	Sides          string `json:"sides"`
}
