package backend

import "errors"

type Order struct {
	Id       int    `json:"id" db:"id"`
	Category string `json:"category" db:"category" binding:"required"`
	Status   string `json:"status" db:"status"`
	Date     string `json:"date" db:"date" binding:"required"`
	Comment  string `json:"comment" db:"comment"`
	UserId   string `json:"user_id"`
}

type UpdateOrderInput struct {
	Category *string `json:"category"`
	Status   *string `json:"status"`
	Date     *string `json:"date"`
	Comment  *string `json:"comment"`
}

func (i UpdateOrderInput) Validate() error {
	if i.Category == nil && i.Status == nil && i.Date == nil && i.Comment == nil {
		return errors.New("update structure has no values")
	}

	return nil
}