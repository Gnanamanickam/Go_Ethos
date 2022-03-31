Account interface { 
		
        balance(amount float64) (bal float64)
        status(id uint8) (status string)

}

type AccountStruct struct {

	AccountID uint8
	Name string
	Balance float64
	Status string

}
