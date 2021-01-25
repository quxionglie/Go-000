package main

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net"
	"os"
	"os/signal"
	"syscall"
	"week09/comm"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetLevel(log.InfoLevel)
}

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	gw, _ := errgroup.WithContext(ctx)

	gw.Go(func() error {
		return tcpServer(ctx)
	})

	gw.Go(func() error {
		return signalServer(ctx, cancelFn)
	})

	if err := gw.Wait(); err != nil {
		fmt.Println("出错:", err)
	}

}

func tcpServer(ctx context.Context) error {
	address := ":8018"
	log.Println("run tcp server in", address)
	listen, err := net.Listen("tcp", address)
	checkError(err)
	//defer listen.Close()

	go func() {
		<-ctx.Done()
		fmt.Println("tcp shutdown ...")
		listen.Close()
	}()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		log.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go handConn(ctx, conn)
	}
}

func handConn(ctx context.Context, conn net.Conn) {
	ch := comm.NewChannnel(conn)
	go ch.Write()
	for {
		ch.Read()
	}
}

func checkError(err error) {
	if err != nil {
		log.Error("Fatal error ", err.Error())
		os.Exit(1)
	}
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
