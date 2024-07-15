package accounts

import (
	"context"
	"fmt"
	"sync"

	pb "awesomeProject/proto"
)

type AccountHandler struct {
	pb.UnimplementedAccountServiceServer
	accounts map[string]*pb.GetAccountResponse
	guard    sync.RWMutex
}

func NewHandler() *AccountHandler {
	return &AccountHandler{
		accounts: make(map[string]*pb.GetAccountResponse),
	}
}

func (h *AccountHandler) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	h.guard.Lock()
	defer h.guard.Unlock()

	if _, exists := h.accounts[req.Name]; exists {
		return nil, fmt.Errorf("account already exists")
	}

	h.accounts[req.Name] = &pb.GetAccountResponse{
		Name:   req.Name,
		Amount: req.Amount,
	}

	return &pb.CreateAccountResponse{}, nil
}

func (h *AccountHandler) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	h.guard.RLock()
	defer h.guard.RUnlock()

	account, exists := h.accounts[req.Name]
	if !exists {
		return nil, fmt.Errorf("account not found")
	}

	return account, nil
}

func (h *AccountHandler) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	h.guard.Lock()
	defer h.guard.Unlock()

	if _, exists := h.accounts[req.Name]; !exists {
		return nil, fmt.Errorf("account not found")
	}

	delete(h.accounts, req.Name)

	return &pb.DeleteAccountResponse{}, nil
}

func (h *AccountHandler) ChangeName(ctx context.Context, req *pb.ChangeNameRequest) (*pb.ChangeNameResponse, error) {
	h.guard.Lock()
	defer h.guard.Unlock()

	account, exists := h.accounts[req.OldName]
	if !exists {
		return nil, fmt.Errorf("account not found")
	}

	if _, exists := h.accounts[req.NewName]; exists {
		return nil, fmt.Errorf("new account name already exists")
	}

	delete(h.accounts, req.OldName)
	account.Name = req.NewName
	h.accounts[req.NewName] = account

	return &pb.ChangeNameResponse{}, nil
}

func (h *AccountHandler) ChangeBalance(ctx context.Context, req *pb.ChangeBalanceRequest) (*pb.ChangeBalanceResponse, error) {
	h.guard.Lock()
	defer h.guard.Unlock()

	account, exists := h.accounts[req.Name]
	if !exists {
		return nil, fmt.Errorf("account not found")
	}

	account.Amount = req.Amount

	return &pb.ChangeBalanceResponse{}, nil
}
