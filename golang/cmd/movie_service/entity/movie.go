package entity

type SearchRequest struct {
	Query string
	Page  int64
}

type SearchResponse struct {
	Movies       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"ImdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
