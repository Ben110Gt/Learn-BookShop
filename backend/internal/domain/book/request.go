package book

type CreateBookRequest struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	CategoryID  string  `json:"category_id"` // FK อยู่ที่ Book
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
type BookResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
