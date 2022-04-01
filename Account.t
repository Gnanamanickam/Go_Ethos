Account interface { 
		
        getBalance(accountName string) (float64)
		transfer(From_ID string, To_ID string, amount float64) (float64, float64)
        getStatus(accountName string) (string)

}

AccountStruct struct {

	AccountID uint8
	Name string
	Balance float64
	Status string

}
