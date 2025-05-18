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
Pengguna dapat mengurutkan proyek berdasarkan total dana terkumpul atau
jumlah donatur menggunakan Selection dan Insertion Sort.
Sistem menampilkan daftar proyek yang sudah mencapai target pendanaan.
**/

package main

import "fmt"

const NMAX int = 10

type event struct {
	kategori       string
	nama           string
	jumlahDonasi   int
	targetDonasi   int
	deskripsiEvent string
	jumlahOrang    int
}

type tabEvent [NMAX]event

func main() {
	var x, nData int
	var event tabEvent
	for {
		menu()
		fmt.Scan(&x)
		switch x {
		case 1:
			buatEvent(&event, &nData)
		case 2:
			cekProyek(&event, &nData)
		case 3:
			donasi(&event, nData)
		case 4:
			//eventSelesai()
		}
		if x == 5 {
			fmt.Print("Sampai Jumpa Kembali ^_^")
			break

		}
	}
}

func menu() {
	fmt.Println("=============================================")
	fmt.Println()
	fmt.Println("           CROWDFUNDING PLUS PLUS            ")
	fmt.Println()
	fmt.Println("=============================================")
	fmt.Println("                     MENU                    ")
	fmt.Println("=============================================")
	fmt.Println("1. Buat Event")
	fmt.Println("2. Cek Event ")
	fmt.Println("3. Donasi	  ")
	fmt.Println("4. Event Selesai")
	fmt.Println("5. Exit")
	fmt.Print("Pilih (1/2/3/4/5) ? ")
}

func buatEvent(A *tabEvent, n *int) {
	var z int

	if *n >= NMAX {
		*n = NMAX
		fmt.Println("Kuota Proyek Sudah Mencapai Batas Maksimum")
	}
	fmt.Println("====================================================================================================")
	fmt.Println(" Helo Selamat datang, silahkan isi data dibawah ini untuk membuat proyek baru yang akan di publish ")
	fmt.Println("====================================================================================================")
	fmt.Println("Silahkan pilih kategori projek anda :")
	fmt.Println("1. Teknologi")
	fmt.Println("2. Kesehatan")
	fmt.Println("3. Startup  ")
	fmt.Println("4. Kreatif  ")
	fmt.Println("(1/2/3/4) : ")
	fmt.Scan(&z)
	switch z {
	case 1:
		A[*n].kategori = "Teknologi"
	case 2:
		A[*n].kategori = "Kesehatan"
	case 3:
		A[*n].kategori = "Startup"
	case 4:
		A[*n].kategori = "Kreatif"
	}
	fmt.Print("Silahkan Masukan Judul Event : ")
	fmt.Scan(&A[*n].nama)
	fmt.Print("Masukan Target Dana : ")
	fmt.Scan(&A[*n].targetDonasi)
	fmt.Println("Deskripsikan Proyek anda : ")
	fmt.Scan(&A[*n].deskripsiEvent)

	fmt.Println("=============================")
	fmt.Println("Projek berhasil dipublish !!")
	fmt.Println("=============================")
	*n = *n + 1
}

func cekProyek(event *tabEvent, n *int) {
	var cari int
	var namaDicari string
	fmt.Println("=======================================================================================")
	fmt.Println("Halo, kami menyediakan beberapa menu jika ingin mengubah data yang sudah di publish :")
	fmt.Println("=======================================================================================")
	fmt.Println("1. Edit Proyek")
	fmt.Println("2. Hapus Proyek")
	fmt.Println("3. Mencari Proyek")
	fmt.Print("opsi yang dipilih (1/2/3) ? ")
	fmt.Scan(&cari)

	switch cari {
	case 1:
		fmt.Print("Masukkan nama/kategori proyek yang ingin diedit: ")
		fmt.Scan(&namaDicari)
		editProyek(event, *n, namaDicari)
	case 2:
		fmt.Print("Masukkan nama/kategori proyek yang ingin dihapus: ")
		fmt.Scan(&namaDicari)
		hapusProyek(event, n, namaDicari)
	case 3:
		fmt.Print("Masukkan nama/kategori proyek yang ingin dicari: ")
		fmt.Scan(&namaDicari)
		mencariProyek(*event, *n, namaDicari)
	}
}

func sequentialSearchIndeks(data tabEvent, jumlah int, x string) int {
	var idx, i int

	idx = -1
	for i < jumlah && idx == -1 {
		if data[i].nama == x || data[i].kategori == x {
			idx = i
		}
		i++
	}
	return idx
}

func mencariProyek(data tabEvent, jumlah int, x string) {
	var idx int

	idx = sequentialSearchIndeks(data, jumlah, x)
	if idx == -1 {
		fmt.Println("Proyek Tidak Ditemukan")
	} else {
		fmt.Printf("Proyek Ditemukan: %s | Kategori: %s | Donasi: %d | Donatur: %d\n",
			data[idx].nama, data[idx].kategori, data[idx].targetDonasi, data[idx].jumlahDonasi)
	}

}

func hapusProyek(data *tabEvent, jumlah *int, x string) {
	var idx, i int

	idx = sequentialSearchIndeks(*data, *jumlah, x)
	if idx == -1 {
		fmt.Println("Proyek Tidak ditemukan")
	} else {
		i = idx
		for i <= *jumlah-2 {
			data[i] = data[i+1]
			i++
		}
		*jumlah = *jumlah - 1
	}
	fmt.Println("Proyek berhasil Dihapus")
}

func editProyek(data *tabEvent, jumlah int, x string) {
	var idx int

	idx = sequentialSearchIndeks(*data, jumlah, x)
	if idx == -1 {
		fmt.Println("Proyek Tidak Ditemukan")
	} else {
		fmt.Print("Masukkan Nama Proyek Baru: ")
		fmt.Scan(&data[idx].nama)
		fmt.Print("Masukkan Target Dana Baru: ")
		fmt.Scan(&data[idx].targetDonasi)
		fmt.Print("Masukkan Deskripsi Baru: ")
		fmt.Scan(&data[idx].deskripsiEvent)
	}
	fmt.Println("Proyek berhasil Diganti")

}

func donasi(A *tabEvent, n int) {
	var x string
	var idx, donasi int

	fmt.Print("=================================================================")
	fmt.Print("Haloo, silahkan masukan nama projek yang ingin di donasikan : ")
	fmt.Scan(&x)
	idx = sequentialSearchIndeks(*A, n, x)

	if idx == -1 {
		fmt.Print("Proyek Tidak ditemukan")
	} else {
		fmt.Printf("Silahkan Masukan Nominal yang akan di donasikan ke proyek %s  : Rp ", A[idx].nama)
		fmt.Scan(&donasi)
		A[idx].jumlahDonasi += donasi
		A[idx].jumlahOrang += 1
		fmt.Printf("Terimakasih sudah melakukan donasi kepada Proyek %s sebesar %d", A[idx].nama, donasi)
	}
}
