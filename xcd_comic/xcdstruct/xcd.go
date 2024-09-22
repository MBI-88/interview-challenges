package xcdstruct

// XcdComic struct
type XcdComic struct {
	Month      string `json:"month"`
	NumID      int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Image      string `json:"image"`
	Alt        string `json:"alt"`
	Day        string `json:"day"`
}

// Comics struct
type Comics []*XcdComic
