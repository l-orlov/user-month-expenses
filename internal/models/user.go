package models

type (
	UserToCreate struct {
		Gender bool  `json:"gender"`
		Age    int16 `json:"age"`
	}
	User struct {
		ID     uint64 `json:"id" binding:"required" db:"id"`
		Gender *bool  `json:"gender" db:"gender"`
		Age    int16  `json:"age" db:"age"`
	}
	UserParams struct {
		ID     *uint64 `json:"id"`
		Gender *bool   `json:"gender"`
		Age    *int16  `json:"age"`
	}
)
