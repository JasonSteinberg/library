package structs

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Authors     string `json:"authors"`
}
