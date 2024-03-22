// decoding json to struct

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Employee struct {
// 	FullName	string	`json:"full_name"`
// 	Email		string	`json:"email"`
// 	Age			int		`json:"age"`
// }

// func main() {
// 	var jsonString = `
// 		{
// 			"full_name": "Airell Jordan",
// 			"email": "airell@mail.com",
// 			"age": 23
// 		}
// 	`

// 	var result Employee

// 	var err = json.Unmarshal([]byte(jsonString), &result)
	
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println("full_name", result.FullName)
// 	fmt.Println("email", result.Email)
// 	fmt.Println("age", result.Age)
// }

// decoding json to map

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	var jsonString = `
// 		{
// 			"full_name": "Airell Jordan",
// 			"email": "airell@mail.com",
// 			"age": 23
// 		}
// 	`

// 	var result map[string]interface{}

// 	var err = json.Unmarshal([]byte(jsonString), &result)
	
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println("full_name", result["full_name"])
// 	fmt.Println("email", result["email"])
// 	fmt.Println("age", result["age"])
// }

// decoding json to empty interface

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	var jsonString = `
// 		{
// 			"full_name": "Airell Jordan",
// 			"email": "airell@mail.com",
// 			"age": 23
// 		}
// 	`

// 	var temp interface{}

// 	var err = json.Unmarshal([]byte(jsonString), &temp)
	
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	var result = temp.(map[string]interface{})

// 	fmt.Println("full_name", result["full_name"])
// 	fmt.Println("email", result["email"])
// 	fmt.Println("age", result["age"])
// }

// decoding json array to slice of struct

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName	string	`json:"full_name"`
	Email		string	`json:"email"`
	Age			int		`json:"age"`
}

func main() {
	var jsonString = `[
			{
				"full_name": "Airell Jordan",
				"email": "airell@mail.com",
				"age": 23
			},
			{
				"full_name": "Ananda RHP",
				"email": "ananda@mail.com",
				"age": 23
			}
		]
	`

	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}