package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	defer fmt.Println("main end")

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	gw, _ := errgroup.WithContext(ctx)

	gw.Go(func() error {
		return httpServer(ctx)
	})

	gw.Go(func() error {
		return signalServer(ctx, cancelFn)
	})

	if err := gw.Wait(); err != nil {
		fmt.Println("出错:", err)
	}
}

func httpServer(ctx context.Context) error {
	defer fmt.Println("exiting http")
	http.HandleFunc("/", HelloHandler)
	srv := &http.Server{
		Addr: ":8000",
	}

	go func() {
		<-ctx.Done()
		fmt.Println("http shutdown ...")
		srv.Shutdown(context.TODO())
	}()

	return srv.ListenAndServe()
}

func signalServer(ctx context.Context, cancelFn func()) error {
	defer fmt.Println("exit signalServer")

	existSigns := []os.Signal{os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}

	sigs := make(chan os.Signal, len(existSigns))
	signal.Notify(sigs, existSigns...)
	for {
		fmt.Println("awaiting signal")
		select {
		case <-ctx.Done():
			fmt.Printf("signal ctx done")
			return ctx.Err()
		case sig := <-sigs:
			fmt.Println("退出信号：", sig)
			cancelFn()
			return errors.New("os.Signal quit")
		}
	}
}
