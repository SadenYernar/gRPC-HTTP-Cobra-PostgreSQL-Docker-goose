package entity

type Product struct {
	Name     string `json:"name"`
	Price    int32  `json:"price"`
	Category string `json:"category"`
	Action   bool   `json:"action"`
	Color    string `json:"color"`
	Width    uint32 `json:"width"`
	Height   uint32 `json:"height"`
	Depth    uint32 `json:"deoth"`
}
