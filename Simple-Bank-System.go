type Bank struct {
	accountsBalance   []int64
	numberOfCustomers int
}

func Constructor(balance []int64) Bank {
	return Bank{accountsBalance: balance, numberOfCustomers: len(balance)}
}
func (this *Bank) IsAccountIDValid(id int) bool {
	return id <= this.numberOfCustomers && id >= 1
}
func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if this.IsAccountIDValid(account1) && this.IsAccountIDValid(account2) {
		if this.accountsBalance[account1-1] >= money {
			this.accountsBalance[account1-1] -= money
			this.accountsBalance[account2-1] += money
			return true
		}
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if this.IsAccountIDValid(account) {
		this.accountsBalance[account-1] += money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if this.IsAccountIDValid(account) && this.accountsBalance[account-1] >= money {
		this.accountsBalance[account-1] -= money
		return true
	}
	return false
}