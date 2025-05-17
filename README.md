# Crowdfunding-App
/** Deskripsi:
Aplikasi ini memungkinkan pengguna untuk melihat dan mencatat kontribusi
terhadap proyek crowdfunding. Data utama yang digunakan adalah daftar proyek,
target dana, dan jumlah kontribusi. Pengguna aplikasi adalah individu yang ingin
berpartisipasi dalam proyek crowdfunding atau pemilik proyek yang mencari
pendanaan.
Spesifikasi:
a. b. c. d. e. Pengguna dapat menambahkan, mengubah, dan menghapus proyek
crowdfunding yang mereka buat.
Sistem mencatat jumlah dana yang terkumpul dan jumlah donatur.
Pengguna dapat mencari proyek berdasarkan nama atau kategori menggunakan
Sequential dan Binary Search.
Pengguna dapat mengureutkan proyek berdasarkan total dana terkumpul atau
jumlah donatur menggunakan Selection dan Insertion Sort.
Sistem menampilkan daftar proyek yang sudah mencapai target pendanaan.
**/

package main

import "fmt"

const NMAX int = 10

type event struct {
	kategori       int
	nama           string
	jumlahDonasi   float64
	targetDonasi   float64
	deskripsiEvent string
}

type tabEvent [NMAX]event

func main() {
	var x, nData int
	var event tabEvent
	menu()
	fmt.Scan(&x)
	switch x {
	case 1:
		buatEvent(&event, &nData)
	case 2:
		cekEvent()
	case 3:
		donasi()
	case 4:
		eventSelesai()
	}
	if x == 4 {
		fmt.Print("Sampai Jumpa Kembali ^_^")
		break

	}
}

func menu() {
	fmt.Println("=============================================")
	fmt.Println("		   	CROWDFUNDING PLUS PLUS				")
	fmt.Println("=============================================")
	fmt.Println()
	fmt.Println("=============================================")
	fmt.Println("					MENU				  	")
	fmt.Println("=============================================")
	fmt.Println("1. Buat Event")
	fmt.Println("2. Cek Event ")
	fmt.Println("3. Donasi	  ")
	fmt.Println("4. Event Selesai")
	fmt.Println("5. Exit")
	fmt.Println("Pilih (1/2/3/4/5) ?")
}

func buatEvent(A *tabEvent, n *int) {
	var x string
	fmt.Println("====================================================================================================")
	fmt.Println(" Helo Selamat datang, silahkan isi data dibawah ini untuk membuat proyek baru yang akan di publish ")
	fmt.Println("====================================================================================================")
	fmt.Print("Silahkan Masukan Judul Event : ")
	fmt.Scan(&A[*n].nama)
	fmt.Print("\nMasukan Target Dana : ")
	fmt.Scan(&A[*n].targetDonasi)
	fmt.Println("Deskripsikan Proyek anda : ")
	fmt.Scan(&A[*n].deskripsiEvent)
	fmt.Println("====================================================================================================")
	fmt.Print("Projek siap di publish ! (ya / tidak)")
	fmt.Scan(&x)
	unPub(A, &x, n)
}

func unPub(A *tabEvent, X *string, n *int) {
	if *X == "yes" {
		*n++
		fmt.Print("Data Berhasil di Publish ^_^")
		return
	} else {
		A[*n].nama = 

	}

	fmt.Println("====================================================================================================")
}
# TUBES-ALPRO
