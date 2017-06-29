package money

type CTrans struct {
	From, To string
	//Fam From amout, Tam, To amount
	Fam, Tam M
}

type Account struct {
	Id  string
	Cur string
	Val M
}

type TransactionList struct {
}
