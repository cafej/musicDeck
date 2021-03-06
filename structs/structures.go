package structs

// Genre is a category of music that represent the DB structure
type Genre struct {
	ID   int    `json:"id" db:"genreID"`
	Name string `json:"name" db:"name"`
}

// GenreTotal is a category of music that represent the DB structure
type GenreTotal struct {
	TotalTime  int    `json:"totalTime" db:"totalTime"`
	TotalSongs int    `json:"totalSongs" db:"totalSongs"`
	Name       string `json:"name" db:"name"`
}

// Song is a artistic sound composition that represents the DB structure
type Song struct {
	ID     int    `json:"id" db:"ID"`
	Artist string `json:"artist" db:"artist"`
	Song   string `json:"song" db:"song"`
	Length int    `json:"length" db:"length"`
	Genre  Genre  `json:"genre" db:"genre"`
}
