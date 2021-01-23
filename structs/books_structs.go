package structs

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required,min=2,max=200"`
	Description string `json:"description"`
	ISBN        string `json:"isbn" binding:"required,min=8,max=10"`
	Authors     string `json:"authors" binding:"required,min=2,max=200"`
	Available   bool
}
