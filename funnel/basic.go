package funnel

type Basic struct {
	Like  string `json:"like,omitempty" query:"like"`
	IdIn  []uint `json:"id_in,omitempty" query:"id_in"`
	order string
}

func (x *Basic) Order(by string) *Basic {
	x.order = by
	return x
}

func (x *Basic) OrderBy() string {
	return x.order
}
