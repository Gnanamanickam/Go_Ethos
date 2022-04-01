package main

import (
	"ethos/syscall"
	"ethos/altEthos"
    "ethos/kernelTypes"
    // "ethos/defined"
	"log"
	// "math"
	"strconv"
)

func init() {
	// SetupAccountcreateAccount(createAccount)
	SetupAccountgetBalance(getBalance)
	SetupAccounttransfer(transfer)
	SetupAccountgetStatus(getStatus)
}

// func createAccount(Name string, ID uint8, Amount float64, Status string) (AccountProcedure) {
// 	log.Printf("Account Server: create Account \n")
// }

func getBalance(account AccountStruct) (AccountProcedure) {
	log.Printf("Account Server: get Balance \n")
	return &AccountgetBalanceReply{account.Balance}
}

func transfer(From_ID AccountStruct, To_ID AccountStruct, amount float64) (AccountProcedure) {
	if From_ID.Balance <= amount {
		log.Printf("Amount given is lesser than the balance \n")
	} else {
		From_ID.Balance = From_ID.Balance - amount
		To_ID.Balance = To_ID.Balance + amount
		log.Printf("Amount Transferred : %v \n", amount)
		log.Printf("Current balance of From_ID: %v \n", From_ID.Balance)
		log.Printf("Current balance of To_ID: %v \n", To_ID.Balance)
		log.Printf("Account Server: account Transfer \n")
	}
	return &AccounttransferReply{From_ID.Balance, To_ID.Balance}
}

func getStatus(account AccountStruct) (AccountProcedure) {
	log.Printf("Account Server: get Status \n")
	return &AccountgetStatusReply{account.Status}
}


func writeToFile(path string, value string, account AccountStruct) {

	var data = kernelTypes.String(value)
	_, status := altEthos.DirectoryOpen(path)
	if status != syscall.StatusOk {
		log.Printf("Open Directory Failed %v: %v\n", path, status)
		status = altEthos.DirectoryCreate(path, &data, "variable")
		if status != syscall.StatusOk {
			log.Printf("Create Directory Failed %v: %v\n", path, status)
		}
		altEthos.DirectoryOpen(path)	
	}

	status = altEthos.Write(path + "/account_" + strconv.Itoa(int(account.AccountID)), &account)
	if status != syscall.StatusOk {
		log.Printf("Error Writing to %v: %v\n", path, status)
	}
}

func readFile(path string, ID uint8) (AccountStruct) {
	
	var data AccountStruct
	// _, status := altEthos.DirectoryOpen(path)
	// if status != syscall.StatusOk {
	// 	log.Printf("Open Directory Failed %v: %v\n", path, status)
	// 	return "failed"
	// }

	status := altEthos.Read(path + "/account_" + strconv.Itoa(int(ID)), &data)
	if status != syscall.StatusOk {
		log.Printf("Error Reading file %v: %v\n", path, status)	
		altEthos.Exit(status)
	}

	return data
}

func main() {

	userName := altEthos.GetUser()
	path := "/home/" + userName + "assignment"

	account1 := AccountStruct{AccountID: 1, Name: "Gnani", Balance: 1250.0, Status: "Active"}
	account2 := AccountStruct{AccountID: 2, Name: "Prem", Balance: 575.0, Status: "Active"}
	account3 := AccountStruct{AccountID: 3, Name: "Virat", Balance: 677.0, Status: "Closed"}

	writeToFile(path, "account1", account1)
	writeToFile(path, "account2", account2)
	writeToFile(path, "account3", account3)

	
	altEthos.LogToDirectory("assignment/accountServer")

	log.Printf("Inside main function \n")
	listeningFd, status := altEthos.Advertise("Account")
	log.Printf("Status is %v \n", status)
	log.Printf("ListeningFd is %v \n", listeningFd)
	if status != syscall.StatusOk {
		log.Printf("Advertising service failed: %s\n", status)
		altEthos.Exit(status)
	}

	for {
		log.Printf("Inside for loop before import\n")
		_, fd, status := altEthos.Import(listeningFd)
		log.Printf("Inside for loop after import\n")
		if status != syscall.StatusOk {
			log.Printf(" Error calling Import %v \n", status)
			altEthos.Exit(status)
		}

		log.Printf("Connection accepted \n")

		t:= Account{}
		altEthos.Handle(fd, &t)

	}
}