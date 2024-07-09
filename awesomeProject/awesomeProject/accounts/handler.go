package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[name]; !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	delete(h.accounts, name)
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) ChangeAccountName(c echo.Context) error {
	var request dto.ChangeAccountNameRequest // {"old_name": "alice", "new_name": "bob"}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[request.OldName]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	if _, exists := h.accounts[request.NewName]; exists {
		return c.String(http.StatusForbidden, "new account name already exists")
	}

	delete(h.accounts, request.OldName)
	account.Name = request.NewName
	h.accounts[request.NewName] = account

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) ChangeAccountBalance(c echo.Context) error {
	var request dto.ChangeAccountBalanceRequest // {"name": "alice", "amount": 30}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[request.Name]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	account.Amount = request.Amount
	return c.NoContent(http.StatusNoContent)
}

// Написать клиент консольный, который делает запросы
