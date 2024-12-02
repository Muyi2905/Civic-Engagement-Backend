package models


type ElectedOfficial struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`                        // Name of the official
	Office   string `json:"office"`                      // Office title (e.g., President, Governor, Mayor)
	Party    string `json:"party"`                       // Political party
	Level    string `json:"level"`                       // Level of government (e.g., National, State, Local)
	State    string `json:"state,omitempty"`             // State associated with the official (if applicable)
	District string `json:"district,omitempty"`          // District for State or Local level (if applicable)
	City     string `json:"city,omitempty"`              // City for Local level (if applicable)
}
