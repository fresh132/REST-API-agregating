package api

type CreateSubscriptionRequest struct {
	ServiceName string `json:"service_name" binding:"required" example:"Netflix"`
	Price       int    `json:"price" binding:"required" example:"500"`
	UserID      string `json:"user_id" binding:"required" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string `json:"start_date" binding:"required" example:"01-2025"`
	EndDate     string `json:"end_date,omitempty" example:"12-2025"`
}

type UpdateSubscriptionRequest struct {
	ServiceName string `json:"service_name" binding:"required" example:"Netflix Premium"`
	Price       int    `json:"price" binding:"required" example:"600"`
	UserID      string `json:"user_id" binding:"required" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string `json:"start_date" binding:"required" example:"01-2025"`
	EndDate     string `json:"end_date,omitempty" example:"12-2025"`
}

type CreateSubscriptionResponse struct {
	ID       int    `json:"id" example:"1"`
	Memssage string `json:"essage" example:"subscription created successfully"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"operation completed successfully"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"error description"`
}

type TotalCostResponse struct {
	TotalCost int `json:"total_cost" example:"1500"`
}
