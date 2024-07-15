package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "awesomeProject/proto"
)

type server struct {
	pb.UnimplementedAccountServiceServer
	accounts map[string]*pb.GetAccountResponse
	guard    sync.RWMutex
}

func (s *server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if _, exists := s.accounts[req.Name]; exists {
		return nil, fmt.Errorf("account already exists")
	}

	s.accounts[req.Name] = &pb.GetAccountResponse{
		Name:   req.Name,
		Amount: req.Amount,
	}

	return &pb.CreateAccountResponse{}, nil
}

func (s *server) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	s.guard.RLock()
	defer s.guard.RUnlock()

	account, exists := s.accounts[req.Name]
	if !exists {
		return nil, fmt.Errorf("account not found")
	}

	return account, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if _, exists := s.accounts[req.Name]; !exists {
		return nil, fmt.Errorf("account not found")
	}

	delete(s.accounts, req.Name)

	return &pb.DeleteAccountResponse{}, nil
}

func (s *server) ChangeName(ctx context.Context, req *pb.ChangeNameRequest) (*pb.ChangeNameResponse, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	account, exists := s.accounts[req.OldName]
	if !exists {
		return nil, fmt.Errorf("account not found")
	}

	if _, exists := s.accounts[req.NewName]; exists {
		return nil, fmt.Errorf("new account name already exists")
	}

	delete(s.accounts, req.OldName)
	account.Name = req.NewName
	s.accounts[req.NewName] = account

	return &pb.ChangeNameResponse{}, nil
}

func (s *server) ChangeBalance(ctx context.Context, req *pb.ChangeBalanceRequest) (*pb.ChangeBalanceResponse, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	account, exists := s.accounts[req.Name]
	if !exists {
		return nil, fmt.Errorf("account not found")
	}

	account.Amount = req.Amount

	return &pb.ChangeBalanceResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":1323")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &server{accounts: make(map[string]*pb.GetAccountResponse)})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
