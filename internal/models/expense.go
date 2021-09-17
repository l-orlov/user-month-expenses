package models

type (
	UserExpense struct {
		UserID   uint64  `json:"id" binding:"required" db:"user_id"`
		Category string  `json:"category" binding:"required" db:"category"`
		Amount   float64 `json:"amount" db:"amount"`
	}
	UserExpenseParams struct {
		UserID   *uint64  `json:"id"`
		Category *string  `json:"category"`
		Amount   *float64 `json:"amount"`
	}
	UserExpenseByCategory struct {
		Category string  `json:"category" db:"category"`
		Amount   float64 `json:"amount" db:"amount"`
	}
	UserExpenseByCategoryToShow struct {
		Category string `json:"category" db:"category"`
		Amount   uint64 `json:"amount" db:"amount"`
	}
)

func (u UserExpenseByCategory) ToUserExpenseByCategoryToShow() UserExpenseByCategoryToShow {
	return UserExpenseByCategoryToShow{
		Category: u.Category,
		Amount:   uint64(u.Amount),
	}
}
