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

// // InitializeDB sets up the database connection
// func InitializeDB() (*gorm.DB, error) {
//     // Example DB connection string for PostgreSQL
//     db, err := gorm.Open("postgres", "host=localhost user=youruser dbname=yourdb password=yourpassword sslmode=disable")
//     if err != nil {
//         return nil, err
//     }

//     db.AutoMigrate(&ElectedOfficial{}) // Auto migrate to create the table
//     return db, nil
// }
