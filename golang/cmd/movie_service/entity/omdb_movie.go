package entity

type SearchRequest struct {
	Query string
	Page  int64
}

type WatchlistRequest struct {
	UserID string
	OmdbID string
	Fav    bool
}

type SearchResponse struct {
	Movies       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
	Page         int64   `json:"page"`
	Result       int64   `json:"result"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"ImdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type DetailResponse struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	ImdbID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	Response   string   `json:"Response"`
	Error      string   `json:"error,omitempty"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type Watchlist struct {
	ID         string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID     string `json:"user_id" bson:"user_id"`
	OmdbID     string `json:"omdb_id" bson:"omdb_id"`
	MovieTitle string `json:"movie_title" bson:"movie_title"`
}
