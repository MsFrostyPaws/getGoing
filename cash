/*The code declares a BankAccount struct with a balance field. 
You need to add a withdraw() method for the BankAccount struct. It should take an integer argument and decrease the balance of the Bank Account by the given amount.
In case there is not enough money in the account, the method should output "Insufficient Funds".
*/

package main
import "fmt"

type BankAccount struct {
  holder string
  balance int
}

func (ptr *BankAccount) withdraw(val int){
	if ptr.balance >= val{
		ptr.balance -= val
	} else {
		fmt.Println("Insufficient Funds")
	}
}
func main() {
  acc := BankAccount{"James Smith", 100000}
  
  var amount int
  fmt.Scanln(&amount)
  
  acc.withdraw(amount)
  fmt.Println(acc)
}
