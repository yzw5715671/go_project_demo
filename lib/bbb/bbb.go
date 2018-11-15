package libbbb

type SetInput struct {
	A int `json:"a"`
	B int `json:"b"`
}

func Set(in SetInput) (ret int, err error) {
	ret = in.A + in.B

	return
}

