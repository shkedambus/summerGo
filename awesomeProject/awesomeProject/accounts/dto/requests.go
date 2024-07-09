package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountNameRequest struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}

type ChangeAccountBalanceRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

//type DeleteAccountRequest struct {
//	Name string `json:"name"`
//}

//type PatchAccountRequest struct {
//	Name string `json:"name"`
//}

//type ChangeAccountRequest struct {
//	Name   string `json:"name"`
//	Amount int    `json:"amount"`
//}
