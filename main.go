package main

import (
	"fmt"
	"os"
	"strconv"
)

type Students struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var student = []Students{
		{nama: "putu",
			alamat:    "bali",
			pekerjaan: "mahasiswa",
			alasan:    "ingin menjadi programmer"},

		{nama: "adi",
			alamat:    "jawa",
			pekerjaan: "siswa",
			alasan:    "ingin menjadi orang hebat"},

		{nama: "putra",
			alamat:    "aceh",
			pekerjaan: "PNS",
			alasan:    "ingin menjadi orang kaya"},
	}

	input := os.Args[1]

	absent, _ := strconv.ParseInt(input, 6, 12)

	showStudent(student, int(absent-1))
}

func showStudent(student []Students, absent int) {
	if len(student) < (absent) {
		fmt.Printf("tidak ada siswa yang memiliki no absen itu")

	} else if 1 > (absent) {
		fmt.Printf("tidak ada siswa yang memiliki no absen itu")

	} else {
		fmt.Printf("nama %+v\n", student[absent].nama)
		fmt.Printf("alamat %+v\n", student[absent].alamat)
		fmt.Printf("pekerjaan %+v\n", student[absent].pekerjaan)
		fmt.Printf("alasan %+v\n", student[absent].alasan)
	}
}
