package main

import (
	"ethos/syscall"
	"ethos/altEthos"
    // "ethos/kernelTypes"
    // "ethos/defined"
	"log"
	// "math"
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

// var fd syscall.Fd
// var fa kernelTypes.AccountServer
// var readData

// func writeToFile(data string, path string) {

// 	fd, status := altEthos.DirectoryOpen(path)
// 	if status != syscall.StatusOk {
// 		logger.Fatalf("Open Directory Failed %v: %v\n", path, status)
// 		status = altEthos.DirectoryCreate(path, &fa, "label")
// 		if status != syscall.StatusOk {
// 			logger.Fatalf("Create Directory Failed %v: %v\n", path, status)
// 		}
// 	}

// 	status = altEthos.writeStream(fd, &data)
// 	if status != syscall.StatusOk {
// 		logger.Fatalf("Error Writing to %v: %v\n", path, status)
// }

// func readFile(data string, path string) {

// 	fd, status := altEthos.DirectoryOpen(path)
// 	if status != syscall.StatusOk {
// 		logger.Fatalf("Open Directory Failed %v: %v\n", path, status)
// 	}
// 	else {
// 		status = altEthos.readStream(fd, &readData)
// 		if status != syscall.StatusOk {
// 			logger.Fatalf("Error Reading file %v: %v\n", path, status)
// 		}
// 	}
// 	log.Printf("Value is %v \n", readData)
// }

func main() {
	
	altEthos.LogToDirectory("")
	// userName := altEthos.GetUser()
	// path := "/user/" + userName + "AccountServer"
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