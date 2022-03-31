package src

import (
	"ethos/syscall"
	"ethos/altEthos"
    "strconv"
    "ethos/kernelTypes"
    "ethos/defined"
	"log"
	"math"
)

type Account struct {

	AccountID uint8
	Name string
	Balance float64
	Status string

}

func init() {
	SetupCreateAccount(createAccount)
	SetupGetBalance(getBalance)
	SetupAccountTransfer(accountTransfer)
	SetupGetStatus(getStatus)
}

func createAccount(Name string, ID uint8, Amount float64, Status string) (AccountProcedure) {
	log.Printf("Account Server: create Account \n")
}

func getBalance(ID uint8) (AccountProcedure) {
	log.Printf("Account Server: get Balance \n")
	return &getBalanceReply{Balance}
}

func accountTransfer(From_ID uint8, To_ID uint8) (AccountProcedure) {
	log.Printf("Account Server: account Transfer \n")
	return &accountTransferReply{"Success"}
}

func getStatus(ID uint8) (AccountProcedure) {
	log.Printf("Account Server: get Status \n")
	return &getStatusReply{Status}
}

var logger = log.Initialize("AccountServer")
var fd syscall.Fd
var fa kernelTypes.AccountServer
var readData

func writeToFile(data string, path string) {

	fd, status := altEthos.DirectoryOpen(path)
	if status != syscall.StatusOk {
		logger.Fatalf("Open Directory Failed %v: %v\n", path, status)
		status = altEthos.DirectoryCreate(path, &fa, "label")
		if status != syscall.StatusOk {
			logger.Fatalf("Create Directory Failed %v: %v\n", path, status)
		}
	}

	status := altEthos.writeStream(fd, &data)
	if status != syscall.StatusOk {
		logger.Fatalf("Error Writing to %v: %v\n", path, status)
}

func readFile(data string, path string) {

	fd, status := altEthos.DirectoryOpen(path)
	if status != syscall.StatusOk {
		logger.Fatalf("Open Directory Failed %v: %v\n", path, status)
	}
	else {
		status := altEthos.readStream(fd, &readData)
		if status != syscall.StatusOk {
			logger.Fatalf("Error Reading file %v: %v\n", path, status)
		}
	}
	log.Printf("Value is %v \n", readData)
}

func main() {
	
	altEthos.LogToDirectory("application/AccountServer")
	userName := altEthos.GetUser()
	path := "/user/" + userName + "AccountServer"

	listeningFd, status := altEthos.Advertise("AccountType")
	if status != syscall.StatusOk {
		log.Printf("Advertising service failed: %s\n", status)
		altEthos.Exit(status)
	}
	else {
		log.Printf("Advertising service initialized \n")
	}

	for {
		name, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Println(" Error calling Import %v \n", status)
			altEthos.Exit(status)
		}

		log.Printf("Connection accepted \n")

		t:= AccountType{}
		altEthos.Handle(fd, &t)

}