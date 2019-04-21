package df

type Province  struct {
	Name   string `json:"name_en"`
	NameCN string `json:"name"`
	City   []City   `json:"city"`
}
