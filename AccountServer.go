package main

import (
	"ethos/syscall"
	"ethos/altEthos"
    // "ethos/kernelTypes"
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

func getBalance(accountName string) (AccountProcedure) {
	account := readFile("/user/" + altEthos.GetUser() + "/account", accountName)
	log.Printf("Account Server: get Balance \n")
	return &AccountgetBalanceReply{account.Balance}
}


func transfer(From_ID string, To_ID string, amount float64) (AccountProcedure) {
	account1 := readFile("/user/" + altEthos.GetUser() + "/account", From_ID)
	account2 := readFile("/user/" + altEthos.GetUser() + "/account", To_ID)
	if account1.Balance <= amount {
		log.Printf("Amount given is lesser than the balance \n")
	} else {
		account1.Balance = account1.Balance - amount
		account2.Balance = account2.Balance + amount
		log.Printf("Amount Transferred : %v \n", amount)
		log.Printf("Current balance of From_ID: %v \n", account1.Balance)
		log.Printf("Current balance of To_ID: %v \n", account2.Balance)
		log.Printf("Account Server: account Transfer \n")
	}
	return &AccounttransferReply{account1.Balance, account2.Balance}
}

func getStatus(accountName string) (AccountProcedure) {
	account := readFile("/user/" + altEthos.GetUser() + "/account", accountName)
	log.Printf("Account Server: get Status \n")
	return &AccountgetStatusReply{account.Status}
}


func writeToFile(path string, value string, account AccountStruct) {

	// var data = kernelTypes.String(value)
	// _, status := altEthos.DirectoryOpen(path)
	// if status != syscall.StatusOk {
	// 	log.Printf("Open Directory Failed %v: %v\n", path, status)
	// 	status = altEthos.DirectoryCreate(path, &data, "variable")
	// 	if status != syscall.StatusOk {
	// 		log.Printf("Create Directory Failed %v: %v\n", path, status)
	// 	}
	// 	altEthos.DirectoryOpen(path)	
	// }
	altEthos.DirectoryCreate(path, &account, "variable")
	status := altEthos.Write(path + "/account_" + strconv.Itoa(int(account.AccountID)), &account)
	log.Printf("Status is %v \n", status)
	if status != syscall.StatusOk {
		log.Printf("Error Writing to %v: %v\n", path, status)
	}
}

func readFile(path string, ID string) (AccountStruct) {
	
	var data AccountStruct
	// _, status := altEthos.DirectoryOpen(path)
	// if status != syscall.StatusOk {
	// 	log.Printf("Open Directory Failed %v: %v\n", path, status)
	// 	return "failed"
	// }

	status := altEthos.Read(path + "/account_" + ID, &data)
	if status != syscall.StatusOk {
		log.Printf("Error Reading file %v: %v\n", path, status)	
		altEthos.Exit(status)
	}

	return data
}

func main() {

	userName := altEthos.GetUser()
	path := "/user/" + userName + "/account"

	account1 := AccountStruct{AccountID: 1, Name: "Gnani", Balance: 1250.0, Status: "Active"}
	account2 := AccountStruct{AccountID: 2, Name: "Prem", Balance: 575.0, Status: "Active"}
	account3 := AccountStruct{AccountID: 3, Name: "Virat", Balance: 677.0, Status: "Closed"}
	account4 := AccountStruct{AccountID: 4, Name: "Dhoni", Balance: 890.0, Status: "Active"}

	log.Printf("Writing to file \n")
	writeToFile(path, "account1", account1)
	writeToFile(path, "account2", account2)
	writeToFile(path, "account3", account3)
	writeToFile(path, "account4", account4)
	log.Printf("Writing to file done \n")

	
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