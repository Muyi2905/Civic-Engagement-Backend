package models

// ElectedOfficial represents an elected official in the United States
type ElectedOfficial struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Office   string `json:"office"`
	Party    string `json:"party"`
	State    string `json:"state"`
	District string `json:"district,omitempty"` // Optional, used for House Representatives
}
