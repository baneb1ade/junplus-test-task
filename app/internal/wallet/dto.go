package wallet

type UpdateWalletRequest struct {
	UUID          string  `json:"uuid" validate:"required,uuid"`
	OperationType string  `json:"operation_type" validate:"required,oneof=deposit withdraw"`
	Amount        float32 `json:"amount" validate:"required"`
}
