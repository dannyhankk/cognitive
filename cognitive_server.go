package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dannyhankk/cognitive/db"
	. "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type MyServer struct {
	UnimplementedText2SpeechServer
}

var (
	ConfigFile = flag.String("c", "./config.json", "you must supply a config file")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := RegisterText2SpeechHandlerFromEndpoint(
		ctx, mux, fmt.Sprintf("localhost:%d", util.RootConfig.RpcPort), opts)
	if err != nil {
		return fmt.Errorf("register rpc server handler failed, %s", err.Error())
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", util.RootConfig.HttpPort), mux)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("panic: %s", err)
		}
	}()
	if err := initAll(); err != nil {
		log.Fatalf("init failed, %s", err.Error())
		return
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", util.RootConfig.RpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}
	s := grpc.NewServer()
	RegisterText2SpeechServer(s, &MyServer{})
	log.Printf("server lstening at: %s", listen.Addr())
	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to server: %s", err)
		}
	}()
	if err := run(); err != nil {
		log.Fatalf("failed to start gate server: %s", err.Error())
	}

}

func initAll() error {
	flag.Parse()
	if err := util.LogInit("./chat.log"); err != nil {
		return fmt.Errorf("init loggger failed, %s", err.Error())
	}
	if err := util.InitConfig(*ConfigFile); err != nil {
		return fmt.Errorf("init config failed, %s", err.Error())
	}
	if err := db.InitRedis(); err != nil {
		return fmt.Errorf("init redis failed, %s", err.Error())
	}
	if err := util.InitTicker(); err != nil {
		return fmt.Errorf("init ticker error, %s", err.Error())
	}
	return nil
}
