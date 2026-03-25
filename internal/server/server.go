package server

import (
	"context"
	"fmt"

	"github.com/rayzhao/grpc-kv-store/internal/store"
	kvstorev1 "github.com/rayzhao/grpc-kv-store/pkg/kvstore/v1"
)

type Server struct {
	kvstorev1.UnimplementedKVStoreServer
	store *store.Store
}

func New(st *store.Store) *Server {
	return &Server{store: st}
}

func (s *Server) Get(ctx context.Context, req *kvstorev1.GetRequest) (*kvstorev1.GetResponse, error) {
	val, ok := s.store.Get(req.Id)
	if !ok {
		return nil, fmt.Errorf("key %s not found", req.Id)
	}
	return &kvstorev1.GetResponse{Value: val}, nil
}

func (s *Server) Put(ctx context.Context, req *kvstorev1.PutRequest) (*kvstorev1.PutResponse, error) {
	val := s.store.Put(req.Id, req.Value)
	return &kvstorev1.PutResponse{Value: val}, nil
}

func (s *Server) Delete(ctx context.Context, req *kvstorev1.DeleteRequest) (*kvstorev1.DeleteResponse, error) {
	ok := s.store.Delete(req.Id)
	if !ok {
		return nil, fmt.Errorf("key %s not found", req.Id)
	}
	return &kvstorev1.DeleteResponse{Deleted: true}, nil
}

// Watch 实现 server streaming — 客户端订阅，服务端持续推送
// TODO: 你来实现
//   1. 调用 s.store.Subscribe(req.Id) 拿到 event channel
//   2. defer s.store.Unsubscribe(req.Id, ch) 清理
//   3. for 循环从 ch 读取 event
//   4. 每次读到 event，用 stream.Send() 发给客户端
//   5. 如果 stream.Context().Done() 了（客户端断开），退出循环
//
// 提示：用 select 同时监听 ch 和 ctx.Done()
func (s *Server) Watch(req *kvstorev1.WatchRequest, stream kvstorev1.KVStore_WatchServer) error {
	// TODO: 实现
	return nil
}

func (s *Server) List(ctx context.Context, req *kvstorev1.ListRequest) (*kvstorev1.ListResponse, error) {
	data := s.store.List()
	items := make([]*kvstorev1.Pack, 0, len(data))
	for k, v := range data {
		items = append(items, &kvstorev1.Pack{Id: k, Value: v})
	}
	return &kvstorev1.ListResponse{Items: items}, nil
}
