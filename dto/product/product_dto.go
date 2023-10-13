package product

type GetProductResponse struct {
	Success bool          `json:"success"`
	Result  GetProductDTO `json:"result"`
}

type ListProductsResponse struct {
	Success bool            `json:"success"`
	Result  ListProductsDTO `json:"result"`
}

type ListProductsDTO struct {
	Products []GetProductDTO `json:"products"`
}

type GetProductDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DailyQuota int    `json:"dailyQuota"`
	Status     string `json:"status"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}
