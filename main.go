package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/yyh-gl/goa-tutorial/usecase"

	svruser "github.com/yyh-gl/goa-tutorial/gen/http/user/server"
	"github.com/yyh-gl/goa-tutorial/gen/user"
	"github.com/yyh-gl/goa-tutorial/presentation/rest"
	goahttp "goa.design/goa/v3/http"
)

func main() {
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	handleServer(ctx, &wg, errc)

	err := <-errc
	fmt.Println(err)
	cancel()
	wg.Wait()
}

func handleServer(ctx context.Context, wg *sync.WaitGroup, errc chan error) {
	// goa専用のリクエストデコーダとレスポンスエンコーダ
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// サービスエンドポイントのまとめ役（マルチプレクサ）
	mux := goahttp.NewMuxer()

	// 各サーバ（サービス）をmuxにマウント
	uu := usecase.NewUser()
	ru := rest.NewUser(uu)
	eu := user.NewEndpoints(ru)
	us := svruser.New(eu, mux, dec, enc, nil, nil)
	svruser.Mount(mux, us)

	var handler http.Handler = mux

	// HTTPサーバを定義
	host := ":8088"
	server := &http.Server{
		Addr:    host,
		Handler: handler,
	}

	// サーバ起動・シャットダウン処理
	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		go func() {
			fmt.Println("Server start " + host)
			errc <- server.ListenAndServe()
		}()

		<-ctx.Done()
		fmt.Println("Server shutdown " + host)

		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			fmt.Println(err)
		}
	}()
}
