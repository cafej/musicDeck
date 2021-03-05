package structs

// Genre is a category of music that represent the DB structure
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Song is a artistic sound composition that represents the DB structure
type Song struct {
	ID     int    `json:"id"`
	Artist string `json:"artist"`
	Song   string `json:"song"`
	Genre  Genre  `json:"genre"`
	Length int    `json:"length"`
}
