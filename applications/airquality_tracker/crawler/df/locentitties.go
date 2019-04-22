package df

type Province struct {
	Name   string `json:"name_en"`
	NameCN string `json:"name"`
	City   []City `json:"city"`
}

type City struct {
	Name   string   `json:"name"`
	County []County `json:"county"`
}

type County struct {
	Name   string `json:"name_en"`
	Code   string `json:"code"`
	NameCN string `json:"name"`
}
