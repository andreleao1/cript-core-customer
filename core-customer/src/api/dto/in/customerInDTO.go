package in

type CustomerInDTO struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	TypedPassword   string `json:"typedPassword" binding:"required"`
	ReTypedPassword string `json:"reTypedPassword" binding:"required"`
}
