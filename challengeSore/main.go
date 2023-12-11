package main

import "fmt"

func main() {
	peoples := []map[string]string{
		{
			"name": "Hank",
			"Age":  "50",
			"Job":  "Polisi",
		},
		{
			"name": "Heisenberg",
			"Age":  "52",
			"Job":  "Ilmuwan",
		},
		{
			"name": "Skyler",
			"Age":  "48",
			"Job":  "Akuntan",
		},
	}
	for _, person := range peoples {
		fmt.Printf("Hi Perkenalkan, Nama saya %v, umur saya %v, dan saya bekerja sebagai %v\n", person["name"], person["Age"], person["Job"])
	}

	soal2()
	soal3()
}

func soal2() {
	var rows1 int = 5

	for i := 0; i < rows1; i++ {
		fmt.Println("*")
	}
}

func soal3() {
	var rows3 int = 5
	for i := 0; i <= rows3; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
