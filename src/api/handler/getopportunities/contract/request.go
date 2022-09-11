package contract

type URLParams struct {
	WhoFilter   string `json:"who" binding:"required"`
	TypeFilter  string `json:"type" binding:"required"`
	AreaFilter  string `json:"area" binding:"required"`
	ExtraFilter string `json:"extra"`
}
