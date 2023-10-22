package bankcore

import (
	"errors"
	"fmt"
)

func Hello() string {
	return "Hey! Wellcome to Simple Online Bank!"
}

// 用户
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// 账户
type Account struct {
	Customer
	Number  int32
	Balance float64
}

// 存款
func (a *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("存入金额不能小于0")
	}

	a.Balance += amount

	return nil
}

// 取款
func (a *Account) Withdraw(amunt float64) error {
	if amunt <= 0 {
		return errors.New("取款金额不能为负数")
	}

	if a.Balance < amunt {
		return errors.New("余额不足")
	}

	a.Balance -= amunt
	return nil
}

// Bank ...
type Bank interface {
	Statement() string
}

func Statement(b Bank) string {
	return b.Statement()
}

// 对账单
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}


// 转账
func (a *Account) Transfer(amount float64, dest *Account) error {
	if amount <= 0 {
		return errors.New("转账金额需要大于0")
	}

	if a.Balance < amount {
		return errors.New("当前账户余额不足，操作失败")
	}

	a.Withdraw(amount)
	dest.Deposit(amount)
	return nil
}