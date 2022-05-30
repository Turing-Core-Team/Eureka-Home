package contract

type URLParams struct {
	FirstFilter  string `json:"first" binding:"required"`
	SecondFilter string `json:"second" binding:"required"`
	ThirdFilter  string `json:"third" binding:"required"`
	FourthFilter string `json:"fourth"`
}
