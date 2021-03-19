package main

import (
	"context"
	"encoding/json"
	"errors"
	pb "finaljob/api/helloworld"
	"finaljob/biz"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	username := vars["u"]
	user_srv, err := biz.InitService()
	if err != nil {
		fmt.Fprintf(w, "{\"code\":\"9999\",\"msg\":\"系统处理异常，请稍后再试\"}")
		return
	}

	res := user_srv.GetUser(username[0])
	json_bytes, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, "{\"code\":\"9999\",\"msg\":\"系统处理异常，请稍后再试\"}")
		return
	}
	w.Write(json_bytes)
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
		return grpcServer(ctx)
	})

	gw.Go(func() error {
		return signalServer(ctx, cancelFn)
	})

	if err := gw.Wait(); err != nil {
		fmt.Println("出错:", err)
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func grpcServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	go func() {
		<-ctx.Done()
		fmt.Println("grpc shutdown ...")
		s.Stop()
	}()

	return s.Serve(lis)
}

func httpServer(ctx context.Context) error {
	defer fmt.Println("exiting http")
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/user", GetUserHandler)
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
