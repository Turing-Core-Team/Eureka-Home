package contract


type OpportunitiesResponse struct {
	Tags            string `json:"tags"`
	Link            string `json:"link"`
	Title           string `json:"title"`
	Requirements    string `json:"requirements"`
	Awards          string `json:"awards"`
	Description     string `json:"description"`
	PublicationDate string `json:"publication_date"`
	UpdateDate      string `json:"update_date"`
	DueDate         string `json:"due_date"`
}