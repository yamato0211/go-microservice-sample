package loader

import (
	"context"
	"go-micro-sample/bff/client"
	"net/http"

	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// Loaders 各DataLoaderを取りまとめるstruct
type Loaders struct {
	UserLoader *dataloader.Loader
}

// NewLoaders Loadersの初期化メソッド
func NewLoaders(client client.IUserService) *Loaders {
	// define the data loader
	userLoader := &UserLoader{
		userService: client,
	}
	loaders := &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userLoader.BatchGetUsers),
	}
	return loaders
}

// Middleware LoadersをcontextにインジェクトするHTTPミドルウェア
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	loaders.UserLoader.ClearAll()
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// GetLoaders ContextからLoadersを取得する
func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
