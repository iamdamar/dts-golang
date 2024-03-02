package main

import (
	"fmt"
	"os"
)

type Student struct {
	name	string
	address	string
	occupation	string
	reason	string
}

func main() {
	students := []Student{
		{
			name : "Arya",
			address : "Jl. Ya'm Sabran, Gg. Bersama, Pontianak Timur",
			occupation : "Fresh Graduate",
			reason : "Ingin belajar ilmu baru",
		},
		{
			name : "Rahmat",
			address : "Jl. Adisucipto, Komplek Bersama, Pontianak Tenggara",
			occupation : "Mahasiswa",
			reason : "Ingin mendalami go language",
		},
		{
			name : "Yani",
			address : "Jl. Johar No 2, Pontianak Kota",
			occupation : "Mahasiswa",
			reason : "Ingin mendalami web dev dengan golang",
		},
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run biodata.go <nomor absen>")
		return
	}

	absen := args[1]

	for i, student := range students {
		if absen == fmt.Sprintf("%d", i+1) {
			fmt.Printf("Nama: %s\n", student.name)
			fmt.Printf("Alamat: %s\n", student.address)
			fmt.Printf("Pekerjaan: %s\n", student.occupation)
			fmt.Printf("Alasan mengikuti kelas Golang: %s\n", student.reason)
			return
		}
	}

	fmt.Println("Tidak ada siswa dengan nomor absen tersebut")

}


// 	var student1 = Student{
// 		name : "Arya",
// 		address : "Jl. Ya'm Sabran, Gg. Bersama, Pontianak Timur",
// 		occupation : "Fresh Graduate",
// 		reason : "Ingin belajar ilmu baru",
// 	}

// 	var student2 = Student{
// 		name : "Rahmat",
// 		address : "Jl. Adisucipto, Komplek Bersama, Pontianak Tenggara",
// 		occupation : "Mahasiswa",
// 		reason : "Ingin mendalami go language",
// 	}

// 	var student3 = Student{
// 		name : "Yani",
// 		address : "Jl. Johar No 2, Pontianak Kota",
// 		occupation : "Mahasiswa",
// 		reason : "Ingin mendalami web dev dengan golang",
// 	}

// 	fmt.Printf("Student1: %+v\n", student1)
// 	fmt.Printf("Student2: %+v\n", student2)
// 	fmt.Printf("Student3: %+v\n", student3)
// }