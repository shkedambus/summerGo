package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "awesomeProject/proto"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	Amount  int
	NewName string
}

func main() {
	portVal := flag.Int("port", 1323, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newNameVal := flag.String("new_name", "", "new name of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *amountVal,
		NewName: *newNameVal,
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cmd.Host, cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAccountServiceClient(conn)

	if err := cmd.do(client); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (cmd *Command) do(client pb.AccountServiceClient) error {
	switch cmd.Cmd {
	case "create":
		if err := cmd.create(client); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}
		return nil
	case "get":
		if err := cmd.get(client); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}
		return nil
	case "delete":
		if err := cmd.deleteAccount(client); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}
		return nil
	case "change_name":
		if err := cmd.changeName(client); err != nil {
			return fmt.Errorf("change name failed: %w", err)
		}
		return nil
	case "change_balance":
		if err := cmd.changeBalance(client); err != nil {
			return fmt.Errorf("change balance failed: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func (cmd *Command) create(client pb.AccountServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.CreateAccount(ctx, &pb.CreateAccountRequest{Name: cmd.Name, Amount: int32(cmd.Amount)})
	if err != nil {
		return fmt.Errorf("could not create account: %v", err)
	}

	fmt.Println("Account created successfully")
	return nil
}

func (cmd *Command) get(client pb.AccountServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetAccount(ctx, &pb.GetAccountRequest{Name: cmd.Name})
	if err != nil {
		return fmt.Errorf("could not get account: %v", err)
	}

	fmt.Printf("Account details - Name: %s, Amount: %d\n", res.Name, res.Amount)
	return nil
}

func (cmd *Command) deleteAccount(client pb.AccountServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.DeleteAccount(ctx, &pb.DeleteAccountRequest{Name: cmd.Name})
	if err != nil {
		return fmt.Errorf("could not delete account: %v", err)
	}

	fmt.Println("Account deleted successfully")
	return nil
}

func (cmd *Command) changeName(client pb.AccountServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.ChangeName(ctx, &pb.ChangeNameRequest{OldName: cmd.Name, NewName: cmd.NewName})
	if err != nil {
		return fmt.Errorf("could not change name: %v", err)
	}

	fmt.Println("Account name changed successfully")
	return nil
}

func (cmd *Command) changeBalance(client pb.AccountServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.ChangeBalance(ctx, &pb.ChangeBalanceRequest{Name: cmd.Name, Amount: int32(cmd.Amount)})
	if err != nil {
		return fmt.Errorf("could not change balance: %v", err)
	}

	fmt.Println("Account balance changed successfully")
	return nil
}
