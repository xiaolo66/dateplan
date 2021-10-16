package dto

type Bubble struct {
	ID int  `json:"id"`
	Title string  `json:"title",form:"title",binding:"required"`
	Status bool	 `json:"status",form:"status"`
}