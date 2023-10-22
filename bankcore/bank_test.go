package bankcore

import (
	"testing"
)

// 测试 客户 和 账户  结构
func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "xiao yu",
			Address: "xi'an Chian",
			Phone: "15055667788",
		},
		Number: 1001,
		Balance: 0,
	}

	if account.Name == " " {
		t.Error("can't create an Account object")
	}
}

// 测试 存款
func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "xiao yu",
			Address: "xi'an Chian",
			Phone: "15055667788",
		},
		Number: 1001,
		Balance: 0,
	}	

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("存款后，余额更新失败")
	}
}

// 测试 存入金额为 负数 的情况
func TestDepositInvalid(t *testing.T) {
    account := Account{
		Customer: Customer{
			Name: "xiao yu",
			Address: "xi'an Chian",
			Phone: "15055667788",
		},
		Number: 1001,
		Balance: 0,
	}

    if err := account.Deposit(-10); err == nil {
        t.Error("存款金额不能为负数")
    }
}

// 测试 取款
func TestWithdraw(t *testing.T) {
    account := Account{
		Customer: Customer{
			Name: "xiao yu",
			Address: "xi'an Chian",
			Phone: "15055667788",
		},
		Number: 1001,
		Balance: 0,
	}

    account.Deposit(10)
    account.Withdraw(10)

    if account.Balance != 0 {
        t.Error("balance is not being updated after withdraw")
    }
}

// 测试 对账单
func TestStatement(t *testing.T) {
    account := Account{
		Customer: Customer{
			Name: "xiao yu",
			Address: "xi'an Chian",
			Phone: "15055667788",
		},
		Number: 1001,
		Balance: 0,
	}

    account.Deposit(100)
    statement := account.Statement()
    if statement != "1001 - xiao yu - 100" {
        t.Error("账户信息错误")
    }
}

// 测试 转账
func TestTransfer(t *testing.T) {
    accountA := Account{
		Customer: Customer{
			Name: "zhang san",
			Address: "xi'an Chian",
			Phone: "15055667788",
		},
		Number: 1001,
		Balance: 0,
	}

	accountB := Account{
		Customer: Customer{
			Name: "li si",
			Address: "xi'an Chian",
			Phone: "13044556677",
		},
		Number: 1002,
		Balance: 0,
	}

	accountA.Deposit(100)
	err := accountA.Transfer(50, &accountB)

	if accountA.Balance != 50 && accountB.Balance != 50 {
		t.Error("转账功能异常", err)
	}
}