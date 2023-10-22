package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	bank "github.com/zyu0211/bankcore"
)

var accounts = map[float64]*bank.Account{}

func main() {
	fmt.Println(bank.Hello())

	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name: "zhang san",
			Address: "xi'an Chian",
			Phone: "15055667788",
		},
		Number: 1001,
		Balance: 0,
	}

    accounts[1002] = &bank.Account{
		Customer: bank.Customer{
			Name: "li si",
			Address: "xi'an Chian",
			Phone: "13055667788",
		},
		Number: 1002,
		Balance: 0,
	}

    http.HandleFunc("/statement", statement)
    http.HandleFunc("/deposit", deposit)
    http.HandleFunc("/withdraw", withdraw)
    http.HandleFunc("/transfer", transfer)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 公开对账单方法
// func statement(w http.ResponseWriter, req *http.Request) {

//     numberqs := req.URL.Query().Get("number")

//     if numberqs == "" {
//         fmt.Fprintf(w, "账号不能为空")
//         return
//     }

//     if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
//         fmt.Fprintf(w, "非法账号")
//     } else {
//         account, ok := accounts[number]
//         if !ok {
//             fmt.Fprintf(w, "账号: %v , 不存在", number)
//         } else {
//             fmt.Fprintf(w, account.Statement())
//         }
//     }
// }
func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "账号不能为空")
		return
	}

	number, err := strconv.ParseFloat(numberqs, 64)
	if err != nil {
		fmt.Fprintf(w, "非法账号")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "账号: %v , 不存在", number)
		} else {
			json.NewEncoder(w).Encode(bank.Statement(account))
		}
	}
}


// 公开存款方法
func deposit(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "账号不能为空")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "非法账号")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "非法金额")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "账号: %v , 不存在", number)
        } else {
            err := account.Deposit(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account.Statement())
            }
        }
    }
}

//公开取款方法
func withdraw(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "账号不能为空")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "非法账号")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "非法金额")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "账号: %v , 不存在", number)
        } else {
            err := account.Withdraw(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account.Statement())
            }
        }
    }
}

// 公开转账方法
func transfer(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")
	destqs := req.URL.Query().Get("dest")

	if numberqs == "" {
		fmt.Fprintf(w, "账号不能为空")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "非法账号")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "非法金额")
	} else if dest, err := strconv.ParseFloat(destqs, 64); err != nil {
		fmt.Fprintf(w, "目标账目异常")
	} else {
		if accountA, ok := accounts[number]; !ok {
			fmt.Fprintf(w, "账号: %v , 不存在", number)
		} else if accountB, ok := accounts[dest]; !ok {
			fmt.Fprintf(w, "账号: %v , 不存在", dest)
		} else {
			err := accountA.Transfer(amount, accountB)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, accountA.Statement())
			}
		}
	}
}