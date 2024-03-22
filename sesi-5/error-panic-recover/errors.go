// package main

// import (
// 	"fmt"
// 	"errors"
// )

// // func main() {
// // 	var number int
// // 	var err error

// // 	number, err = strconv.Atoi("123GH")

// // 	if err == nil {
// // 		fmt.Println(number)
// // 	} else {
// // 		fmt.Println(err.Error())
// // 	}

// // 	number, err = strconv.Atoi("123")

// // 	if err == nil {
// // 		fmt.Println(number)
// // 	} else {
// // 		fmt.Println(err.Error())
// // 	}
// // }

// func main() {
// 	var password string

// 	fmt.Scanln(&password)

// 	if valid, err := validPassword(password); err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(valid)
// 	}
// }

// func validPassword(password string) (string, error) {
// 	pl := len(password)

// 	if pl < 5 {
// 		return "", errors.New("password has to have more than 4 characters")
// 	}

// 	return "Valid password", nil
// }

package main

import (
	"errors"
	"fmt"
)

func main() {
	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "Valid password", nil
}