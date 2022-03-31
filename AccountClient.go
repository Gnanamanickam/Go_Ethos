package src

import (
        "ethos/altEthos"
        "ethos/syscall"
        "log"
        "strconv"
        "ethos/kernelTypes"
        "ethos/defined"
		"math"
)

func init() {
	SetupCreateAccountReply(createAccountReply)
	SetupGetBalanceReply(getBalanceReply)
	SetupAccountTransferReply(accountTransferReply)
	SetupGetStatusReply(getStatusReply)
}

func createAccountReply(Name string, ID uint8, Amount float64, Status string) (AccountProcedure) {
	log.Printf("Account Client : Received Account Reply: %v, %v, %v, %v \n", Name, ID, Amount, Status)
	return nil
}

func getBalanceReply(Amount float64) (AccountProcedure) {
	log.Printf("Account Client : Received Balance: %v \n", Amount)
	return nil
}

func accountTransferReply(From_ID uint8, To_ID uint8) (AccountProcedure) {
	log.Printf("Account Client : Amount Transferred \n")
	return nil
}

func getStatusReply(Status string) (AccountProcedure) {
	log.Printf("Account Client : Status: %v \n, Status")
	return nil
}

func main() {

	altEthos.LogToDirectory("AccountClient")
	log.Printf("Account Client before_call \n")
	fd, status := altEthos.IpcRepeat("AccountType", "", nil)
	if status != syscall.StatusOk { 
		log.Printf("Ipc failed: %v \n", status)
		altEthos.Exit(status)

	}

		call := FillItLater{}
		status = altEthos.ClientCall(fd, &call)
		
		if status != syscall.StatusOk {
			log.Printf("Client Call Failed: %v \n", status)
			altEthos.Exit(status)
	}

	log.Printf("Account Client done \n")

}