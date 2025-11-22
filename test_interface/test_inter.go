package test_interface

import "fmt"

type Account interface {
	getBalance() int
	pay(amount int) bool // pay( int) bool 也可以，入参不要求有名称
}

type Payment interface { // 声明一个接口，接口包含了一系列方法， 不包含具体的实现，只需要在结构体中实现这些方法即认为实现了接口
	Account
}

type Object interface{} // 空接口，可以接收任意类型 类似于java中的object 可以复制任意类型的变量

type CreditCard struct {
	balance int
	loan    int
}

type DebitCard struct {
	balance int
}

func (card *CreditCard) getLoan() int {
	return card.loan
}

func (card *CreditCard) getBalance() int {
	return card.balance
}

func (card *CreditCard) pay(amount int) bool {
	card.balance -= amount //当欠款
	if card.balance < 0 {
		card.loan -= card.balance
		card.balance = 0
		fmt.Println("欠款为", card.loan, "余额为", card.balance)
	}
	return true
}

func (card *DebitCard) getBalance() int {
	return card.balance
}

func (card *DebitCard) pay(amount int) bool {
	if amount > card.balance {
		return false
	}
	card.balance -= amount
	return true
}

func purchase(p Payment, amount int) { // 接口只能接收指针，所以不用强调 *Payment
	if p.pay(amount) {
		fmt.Println("购买成功, 余额为", p.getBalance())
	} else {
		fmt.Println("购买失败")
	}
}

func Run() {
	fmt.Println("test_interface")
	cc := CreditCard{
		balance: 100,
		loan:    0,
	}
	dc := DebitCard{
		balance: 100,
	}

	purchase(&cc, 500) // 结构体实现的方法的入参都是指针，所以传参必须用指针， 如果方法的入参不是指针，那么传参可以是值也可以是指针, 不是很好懂
	purchase(&dc, 50)

	var account Account = &dc // Account方法的实现 使用dc实例中的实现逻辑
	account.pay(50)
	fmt.Println("account 余额为,", account.getBalance())
}
