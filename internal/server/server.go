package server

import (
	"context"
	"fmt"

	"github.com/rayzhao/grpc-kv-store/internal/store"
	kvstorev1 "github.com/rayzhao/grpc-kv-store/pkg/kvstore/v1"
)

// TODO: 实现 gRPC 服务
//
// Server 已经帮你定义好了，你需要实现四个方法：
//   - Get: 调用 s.store.Get()，找不到时返回 gRPC codes.NotFound 错误
//   - Put: 调用 s.store.Put()，返回 PutResponse
//   - Delete: 调用 s.store.Delete()，返回是否删除成功
//   - List: 调用 s.store.List()，把 map 转成 []*kvstorev1.Pack
//
// 需要的包：
//   - "google.golang.org/grpc/codes"
//   - "google.golang.org/grpc/status"
//
// 示例（Get 的大致结构）：
//   func (s *Server) Get(ctx context.Context, req *kvstorev1.GetRequest) (*kvstorev1.GetResponse, error) {
//       val, ok := s.store.Get(req.GetId())
//       if !ok {
//           return nil, status.Errorf(codes.NotFound, "key %s not found", req.GetId())
//       }
//       return &kvstorev1.GetResponse{Value: val}, nil
//   }

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
		return nil, fmt.Errorf("Get(id=%s)", req.Id)
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
		return nil, fmt.Errorf("Delete(id=%s)", req.Id)
	}
	return &kvstorev1.DeleteResponse{Deleted: true}, nil
}

func (s *Server) List(ctx context.Context) (*kvstorev1.ListResponse, error) {
	if s == nil {
		return nil, fmt.Errorf("List()")
	}
	data := s.store.List()
	items := make([]*kvstorev1.Pack, 0, len(data))
	for k, v := range data {
		items = append(items, &kvstorev1.Pack{Id: k, Value: v})
	}
	return &kvstorev1.ListResponse{Items: items}, nil
}
