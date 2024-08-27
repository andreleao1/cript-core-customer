package in

type ReserveBalanceInDTO struct {
	WalletId string  `json:"walletId" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
}
