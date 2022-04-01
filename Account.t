Account interface { 
		
        getBalance(account AccountStruct) (float64)
		transfer(From_ID AccountStruct, To_ID AccountStruct, amount float64) (float64, float64)
        getStatus(account AccountStruct) (string)

}

AccountStruct struct {

	AccountID uint8
	Name string
	Balance float64
	Status string

}
