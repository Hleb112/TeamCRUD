package models

type Guitar struct {
	Model       string        `json:"model"`
	Brand       string        `json:"brand"`
	Color       string        `json:"color"`
	Year        int           `json:"year"`
	IsAvailable bool          `json:"is_available"`
	Strings     GuitarStrings `json:"strings"`
}

type GuitarStrings struct {
	Material    string `json:"material"`
	Number      int    `json:"number"`
	IsAvailable bool   `json:"is_available"`
}

type Customer struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	IsNew   bool   `json:"is_new"` // Is it first time in our shop
	Card    Card   `json:"card"`   // Card of loyalty

}

type Card struct {
	ID         int    `json:"id"`          // Id number of card
	LvlCard    string `json:"lvl_card"`    // Different lvl of loyalty (bronze,silver,gold)
	DateStart  int    `json:"date_start"`  // date of loyalty card issuance
	DateFinish int    `json:"date_finish"` // loyalty card expiration date
}
