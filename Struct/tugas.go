package main

import (
	"fmt"
)

type Mhs struct {
	NIM uint8
	Nama string
	Gender string
	Alamat string
}

type Matkul struct {
	NamaMatkul string
	NilaiTugas uint8
	NilaiUTS uint8
	NilaiUAS uint8
}

type MhsMatkul struct {
	Mahasiswa Mhs
	MataKuliah Matkul
}

func main() {
	var MhsMatkul MhsMatkul
	MhsMatkul.Mahasiswa.NIM = 1
	MhsMatkul.Mahasiswa.Nama = "Ninokoni"
	MhsMatkul.Mahasiswa.Alamat = "Jl. Kebon Jeruk"
	MhsMatkul.Mahasiswa.Gender = "Perempuan"
	MhsMatkul.MataKuliah.NamaMatkul = "Pemrograman Berorientasi Objek"
	MhsMatkul.MataKuliah.NilaiTugas = 80
	MhsMatkul.MataKuliah.NilaiUTS = 90
	MhsMatkul.MataKuliah.NilaiUAS = 100

	fmt.Println(MhsMatkul)
}