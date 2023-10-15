package order

type ListOrdersResponse struct {
	Success bool          `json:"success"`
	Result  ListOrdersDTO `json:"result"`
}

type ListOrdersDTO struct {
	Orders []GetOrderDTO `json:"orders"`
}

type GetOrderDTO struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
