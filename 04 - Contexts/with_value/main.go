package withvalue

import "context"

type contextKey string

func Run() {
	ctx := context.WithValue(context.Background(), contextKey("user"), "jelps")
	auth(ctx)
}

func auth(ctx context.Context) {
	// get user from context
	user := ctx.Value(contextKey("user"))
	if user == nil {
		println("unauthorized")
		return
	}

	println("authorized")
}
