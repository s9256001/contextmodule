package contextmodule

import "context"

type TData struct {
	Name string
}

func SetData(ctx *context.Context, data *TData) {
	*ctx = context.WithValue(*ctx, "data", data)
}

func GetData(ctx *context.Context) *TData {
	if data := (*ctx).Value("data"); data != nil {
		return data.(*TData)
	}
	return nil
}
