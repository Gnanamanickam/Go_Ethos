package main

import (
        "ethos/altEthos"
        "ethos/syscall"
        "log"
        // "strconv"
        // "ethos/kernelTypes"
        // "ethos/defined"
		// "math"
)

func init() {
	// SetupAccountcreateAccountReply(createAccountReply)
	SetupAccountgetBalanceReply(getBalanceReply)
	SetupAccounttransferReply(transferReply)
	SetupAccountgetStatusReply(getStatusReply)
}

// func createAccountReply(Name string, ID uint8, Amount float64, Status string) (AccountProcedure) {
// 	log.Printf("Account Client : Received Account Reply: %v, %v, %v, %v \n", Name, ID, Amount, Status)
// 	return nil
// }

func getBalanceReply(Amount float64) (AccountProcedure) {
	log.Printf("Account Client, Balance: %v \n", Amount)
	return nil
}

func transferReply(Amount1 float64, Amount2 float64) (AccountProcedure) {
	log.Printf("Account Client : Amount Transferred \n")
	return nil
}

func getStatusReply(Status string) (AccountProcedure) {
	log.Printf("Account Client : Status: %v \n, Status")
	return nil
}

func main() {

	account1 := AccountStruct{AccountID: 1, Name: "Gnani", Balance: 1250.0, Status: "Active"}
	account2 := AccountStruct{AccountID: 2, Name: "Prem", Balance: 575.0, Status: "Active"}

	altEthos.LogToDirectory("assignment/accountClient")
	log.Printf("Account Client before_call \n")

	fd, status := altEthos.IpcRepeat("Account", "", nil)
	if status != syscall.StatusOk { 
		log.Printf("Ipc failed: %v \n", status)
		altEthos.Exit(status)
	}

	call := AccountgetBalance{account1}
	status = altEthos.ClientCall(fd, &call)
		
	if status != syscall.StatusOk {
			log.Printf("Client Call Failed: %v \n", status)
			altEthos.Exit(status)
	}

	fd1, status1 := altEthos.IpcRepeat("Account", "", nil)
	if status != syscall.StatusOk { 
		log.Printf("Ipc failed: %v \n", status)
		altEthos.Exit(status)

	}

		call1 := Accounttransfer{account1, account2, float64(100.0)}
		status1 = altEthos.ClientCall(fd1, &call1)
		
		if status != syscall.StatusOk {
			log.Printf("Client Call Failed: %v \n", status1)
			altEthos.Exit(status1)
	}

	fd2, status2 := altEthos.IpcRepeat("Account", "", nil)
	if status2 != syscall.StatusOk { 
		log.Printf("Ipc failed: %v \n", status2)
		altEthos.Exit(status2)

	}

		call2 := AccountgetStatus{account1}
		status2 = altEthos.ClientCall(fd2, &call2)
		
		if status != syscall.StatusOk {
			log.Printf("Client Call Failed: %v \n", status2)
			altEthos.Exit(status2)
	}

	log.Printf("Account Client done \n")

}