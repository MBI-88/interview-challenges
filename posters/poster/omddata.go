package poster

var (
	key = "43517a81"
	api = "http://www.omdbapi.com/" //http://www.omdbapi.com/?apikey=43517a81&t=Blade
)

// OmdAPI struct. Contain old data for the omdapi result.
type OmdAPI struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Ratings    []rating
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"ImdbRating"`
	ImdbVotes  string `json:"ImdbVotes"`
	ImdbID     string `json:"ImdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

type rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
