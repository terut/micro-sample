package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/terut/micro-sample/proto/gen"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("health check")
	})
	http.HandleFunc("/word", wordHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("listen started")
}

func wordHandler(w http.ResponseWriter, r *http.Request) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create logger :%v", err)
	}
	var servName string
	if s := os.Getenv("BACKEND_SERVICE_NAME"); s != "" {
		servName = s
	} else {
		servName = "127.0.0.1"
	}
	logger.Debug("go", zap.String("servName", servName))
	conn, err := grpc.Dial(servName+":8000", grpc.WithInsecure(), grpc.WithUnaryInterceptor(
		grpc_zap.UnaryClientInterceptor(logger),
	))
	grpc_zap.ReplaceGrpcLogger(logger)
	if err != nil {
		log.Fatalf("failed to connect :%v", err)
	}
	defer conn.Close()

	val := r.URL.Query().Get("val")
	//if err != nil {
	//		logger.Error("got value error", zap.Error(err))
	//	}
	client := pb.NewWordClient(conn)
	ctx := context.Background()
	res, err := client.Build(ctx, &pb.WordRequest{Val: val})
	if err != nil {
		logger.Error("got error from server", zap.Error(err))
	}
	logger.Info("got response", zap.String("value", res.Val))
	b, err := json.Marshal(res)
	if err != nil {
		logger.Error("json parse error", zap.Error(err))
	}
	w.Write(b)
}
