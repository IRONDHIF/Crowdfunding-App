package main

import "fmt"

const NMAX int = 5

var idProyek int = 101

type proyek struct {
	id              int
	kategori        string
	nama            string
	jumlahDonasi    int
	targetDonasi    int
	deskripsiproyek [100]string
	jumlahOrang     int
	status          bool
}

type tabProyek [NMAX]proyek

func main() {
	var x, nData int
	var proyek tabProyek
	for {
		menu()
		fmt.Scan(&x)
		switch x {
		case 1:
			buatproyek(&proyek, &nData)
		case 2:
			listproyek(proyek, nData)
		case 3:
			cekProyek(&proyek, &nData)
		case 4:
			donasi(&proyek, nData)
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
	fmt.Println("           CROWDFUNDING PLUS-PLUS            ")
	fmt.Println()
	fmt.Println("=============================================")
	fmt.Println("                   MENU UTAMA                ")
	fmt.Println("=============================================")
	fmt.Println("1. Buat Proyek")
	fmt.Println("2. List Proyek")
	fmt.Println("3. Cek Proyek")
	fmt.Println("4. Donasi")
	fmt.Println("5. Exit")
	fmt.Print("Silakan pilih menu (1/2/3/4/5): ")
}

func buatproyek(A *tabProyek, n *int) {
	var z, i int
	var word string
	if *n > NMAX-1 {
		*n = NMAX
		fmt.Println("Kuota Proyek Sudah Mencapai Batas Maksimum")
		return
	}
	fmt.Println("====================================================================================================")
	fmt.Println("Selamat datang! Silakan isi informasi berikut untuk membuat dan mempublikasikan proyek baru.")
	fmt.Println("====================================================================================================")

	A[*n].id = idProyek
	idProyek++
	fmt.Println("Pilih kategori proyek Anda:")
	fmt.Println("1. Teknologi")
	fmt.Println("2. Kesehatan")
	fmt.Println("3. Startup  ")
	fmt.Println("4. Kreatif  ")
	fmt.Print("Pilih Kategori (1/2/3/4) : ")
	fmt.Scan(&z)
	fmt.Println()
	if z == 1 {
		A[*n].kategori = "Teknologi"
	} else if z == 2 {
		A[*n].kategori = "Kesehatan"
	} else if z == 3 {
		A[*n].kategori = "Startup"
	} else if z == 4 {
		A[*n].kategori = "Kreatif"
	}

	fmt.Print("Masukkan judul proyek (Gunakan '_' untuk spasi): ")
	fmt.Scan(&A[*n].nama)
	fmt.Print("Masukkan target donasi (Rp): ")
	fmt.Scan(&A[*n].targetDonasi)
	fmt.Println("Tuliskan deskripsi proyek (akhiri dengan titik '.'): ")
	i = 0
	for {
		fmt.Scan(&word)
		if word == "." {
			break
		}
		if i < 100 {
			A[*n].deskripsiproyek[i] = word
			i++
		} else {
			fmt.Println("Deskripsi terlalu panjang. Pengisian dihentikan.")
			break
		}
	}
	A[*n].jumlahDonasi = 0
	A[*n].jumlahOrang = 0
	fmt.Println("================================")
	fmt.Println("Proyek berhasil dipublikasikan.")
	fmt.Println("================================")
	fmt.Println()
	*n = *n + 1
}

func listproyek(A tabProyek, n int) {
	var answer int
	if n < 1 {
		fmt.Println("Belum ada proyek yang tersedia.")
	} else {
		fmt.Printf("Daftar proyek:\n")
		cetakData(A, n)
		fmt.Print()
		fmt.Println("1. Lihat proyek yang sudah mencapai target")
		fmt.Println("2. Urutkan list proyek berdasarkan jumlah donatur (menaik)")
		fmt.Println("3. Urutkan list proyek berdasarkan jumlah donatur (menurun)")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih (1/2/3/4) : ")
		fmt.Scan(&answer)
		fmt.Println()
		if answer == 1 {
			sequentialSearchStatus(A, n)
		} else if answer == 2 {
			insertionSort(&A, n)
			fmt.Println("Proyek telah diurutkan berdasarkan jumlah donatur (menaik):")
			cetakData(A, n)
		} else if answer == 3 {
			selectionSort(&A, n)
			fmt.Println("Proyek telah diurutkan berdasarkan jumlah donatur (menurun):")
			cetakData(A, n)
		} else {
			return
		}
	}
}

func cekProyek(proyek *tabProyek, n *int) {
	var cari, pilih, idxMax, idxMin int
	var j int
	var namaDicari string
	var cariId int
	if *n < 1 {
		fmt.Println("Belum ada proyek yang tersedia.")
	} else {
		fmt.Println("=======================================================================================")
		fmt.Println("Silakan pilih menu untuk mengelola proyek yang telah dipublikasikan:")
		fmt.Println("=======================================================================================")
		fmt.Println("1. Edit Proyek")
		fmt.Println("2. Hapus Proyek")
		fmt.Println("3. Cari Proyek")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih (1/2/3/4) : ")
		fmt.Scan(&cari)
		fmt.Println("=======================================================================================")
		fmt.Println()

		if cari == 1 {
			fmt.Print("Masukkan nama proyek yang ingin diedit (Gunakan _ untuk spasi): ")
			fmt.Scan(&namaDicari)
			editProyek(proyek, *n, namaDicari)
		} else if cari == 2 {
			fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
			fmt.Scan(&namaDicari)
			hapusProyek(proyek, n, namaDicari)
		} else if cari == 3 {
			fmt.Println("1. Cari berdasarkan nama atau kategori proyek")
			fmt.Println("2. Cari berdasarkan Id proyek")
			fmt.Println("3. Cari berdasarkan target dana terbesar")
			fmt.Println("4. Cari berdasarkan target dana terkecil")
			fmt.Print("Pilih (1/2/3/4) : ")
			fmt.Scan(&pilih)
			if pilih == 1 {
				fmt.Print("Masukkan nama atau kategori proyek yang ingin dicari: ")
				fmt.Scan(&namaDicari)
				mencariSemuaProyek(*proyek, *n, namaDicari)
			} else if pilih == 2 {
				fmt.Print("Masukkan Id proyek yang ingin dicari: ")
				fmt.Scan(&cariId)
				cariIdProyek(*proyek, *n, cariId)
			} else if pilih == 3 {
				idxMax = mencariMax(*proyek, *n)
				fmt.Println("=========================================================================================================")
				fmt.Printf("ID: %d || Judul: %s || Kategori: %s || Target Donasi: Rp %d || Total Donasi Terkumpul: Rp %d || Jumlah Donatur: %d\n", proyek[idxMax].id, proyek[idxMax].nama, proyek[idxMax].kategori, proyek[idxMax].targetDonasi, proyek[idxMax].jumlahDonasi, proyek[idxMax].jumlahOrang)
				fmt.Print("Deskripsi Proyek: ")
				j = 0
				for j < 100 && proyek[idxMax].deskripsiproyek[j] != "." {
					fmt.Print(proyek[idxMax].deskripsiproyek[j], " ")
					j++
				}
				fmt.Println()
				fmt.Println("=========================================================================================================")
			} else if pilih == 4 {
				idxMin = mencariMin(*proyek, *n)
				fmt.Println("=========================================================================================================")
				fmt.Printf("ID: %d || Judul: %s || Kategori: %s || Target Donasi: Rp %d || Total Donasi Terkumpul: Rp %d || Jumlah Donatur: %d\n", proyek[idxMin].id, proyek[idxMin].nama, proyek[idxMin].kategori, proyek[idxMin].targetDonasi, proyek[idxMin].jumlahDonasi, proyek[idxMin].jumlahOrang)
				fmt.Print("Deskripsi Proyek: ")
				j = 0
				for j < 100 && proyek[idxMax].deskripsiproyek[j] != "." {
					fmt.Print(proyek[idxMin].deskripsiproyek[j], " ")
					j++
				}
				fmt.Println()
				fmt.Println("=========================================================================================================")
			}
		}
	}
}

func donasi(A *tabProyek, n int) {
	var x string
	var idx, donasi int
	if n < 1 {
		fmt.Println("Belum ada proyek yang tersedia.")
	} else {
		fmt.Println("=================================================================")
		fmt.Print("Masukkan nama proyek yang ingin Anda berikan donasi: ")
		fmt.Scan(&x)
		idx = sequentialSearchIndeks(*A, n, x)

		if idx == -1 {
			fmt.Println("Proyek tidak ditemukan.")
		} else {
			if A[idx].jumlahDonasi >= A[idx].targetDonasi {
				fmt.Println("Proyek ini telah mencapai target donasi.")
			} else {
				fmt.Printf("Masukkan nominal donasi untuk proyek '%s': Rp ", A[idx].nama)
				fmt.Scan(&donasi)
				A[idx].jumlahDonasi += donasi
				if A[idx].jumlahDonasi >= A[idx].targetDonasi {
					fmt.Printf("Terima kasih atas donasi sebesar Rp %d untuk proyek '%s'. Target donasi telah tercapai.\n", donasi, A[idx].nama)
					A[idx].jumlahOrang += 1
					A[idx].status = true
					return
				}
				fmt.Printf("Terima kasih atas donasi sebesar Rp %d untuk proyek '%s'.\n", donasi, A[idx].nama)
				A[idx].jumlahOrang += 1

			}
		}
	}
}

func mencariSemuaProyek(data tabProyek, n int, x string) {
	var ketemu bool
	var i, j int

	ketemu = false
	i = 0

	for i < n {
		if data[i].nama == x || data[i].kategori == x {
			fmt.Println("============================")
			fmt.Printf("Proyek Ditemukan:\n")
			fmt.Printf("Id Proyek : %d\n", data[i].id)
			fmt.Printf("Judul: %s\n", data[i].nama)
			fmt.Printf("Kategori: %s\n", data[i].kategori)
			fmt.Printf("Target Donasi: Rp %d\n", data[i].targetDonasi)
			fmt.Printf("Total Donasi Terkumpul: Rp %d\n", data[i].jumlahDonasi)
			fmt.Printf("Jumlah Donatur: %d\n", data[i].jumlahOrang)
			fmt.Println("Deskripsi Proyek: ")
			j = 0

			for j < 100 && data[i].deskripsiproyek[j] != "." {
				fmt.Print(data[i].deskripsiproyek[j], " ")
				j++
			}
			fmt.Println()
			ketemu = true
		}
		i++
	}
	if !ketemu {
		fmt.Println("Proyek tidak ditemukan")
		fmt.Println()
	}
}

func cariIdProyek(data tabProyek, n int, x int) {
	var j, idx int

	idx = BinSearchIndeks(data, n, x)

	if idx == -1 {
		fmt.Println("Proyek tidak ditemukan")
		return
	} else {
		fmt.Println("============================")
		fmt.Printf("Proyek Ditemukan:\n")
		fmt.Printf("Id Proyek : %d\n", data[idx].id)
		fmt.Printf("Judul: %s\n", data[idx].nama)
		fmt.Printf("Kategori: %s\n", data[idx].kategori)
		fmt.Printf("Target Donasi: Rp %d\n", data[idx].targetDonasi)
		fmt.Printf("Total Donasi Terkumpul: Rp %d\n", data[idx].jumlahDonasi)
		fmt.Printf("Jumlah Donatur: %d\n", data[idx].jumlahOrang)
		fmt.Println("Deskripsi Proyek: ")
		j = 0

		for j < 100 && data[idx].deskripsiproyek[j] != "." {
			fmt.Print(data[idx].deskripsiproyek[j], " ")
			j++
		}
	}
	fmt.Println()
}

// PROSEDUR HAPUS, EDIT
func hapusProyek(data *tabProyek, jumlah *int, x string) {
	var idx, i int

	idx = sequentialSearchIndeks(*data, *jumlah, x)
	if idx == -1 {
		fmt.Println("Proyek tidak ditemukan")
		return
	} else {
		i = idx
		for i <= *jumlah-2 {
			data[i] = data[i+1]
			i++
		}
		*jumlah = *jumlah - 1
	}
	fmt.Println("Proyek berhasil dihapus.")
}

func editProyek(data *tabProyek, jumlah int, x string) {
	var idx, i int
	var word string
	idx = sequentialSearchIndeks(*data, jumlah, x)
	if idx == -1 {
		fmt.Println("Proyek tidak ditemukan")
		return
	} else {
		fmt.Print("Masukkan Nama Proyek Baru: ")
		fmt.Scan(&data[idx].nama)
		fmt.Print("Masukkan Target Dana Baru: ")
		fmt.Scan(&data[idx].targetDonasi)
		fmt.Println("Deskripsikan Proyek anda (masukan '.' untuk mengakhiri):")
		i = 0
		for {
			fmt.Scan(&word)
			if word == "." {
				break
			}
			if i < 100 {
				data[*&idx].deskripsiproyek[i] = word
				i++
			} else {
				fmt.Println("Deskripsi terlalu panjang, dihentikan.")
				break
			}
		}
		fmt.Println("Proyek berhasil diperbaharui.")
	}

}

// PROSEDUR CETAK
func cetakData(A tabProyek, n int) {
	var i, j int
	for i = 0; i < n; i++ {
		fmt.Println("=========================================================================================================")
		fmt.Printf("ID: %d || Judul: %s || Kategori: %s || Target Donasi: Rp %d || Total Donasi Terkumpul: Rp %d || Jumlah Donatur: %d\n", A[i].id, A[i].nama, A[i].kategori, A[i].targetDonasi, A[i].jumlahDonasi, A[i].jumlahOrang)
		fmt.Print("Deskripsi Proyek: ")
		j = 0
		for j < 100 && A[i].deskripsiproyek[j] != "." {
			fmt.Print(A[i].deskripsiproyek[j], " ")
			j++
		}
		fmt.Println()
		fmt.Println("=========================================================================================================")
	}
}

// FUNC PENCARIAN MIN MAX
func mencariMax(A tabProyek, n int) int {
	var max, i int
	max = 0
	for i = 1; i < n; i++ {
		if A[i].targetDonasi > A[max].targetDonasi {
			max = i
		}
	}
	return max
}

func mencariMin(A tabProyek, n int) int {
	var min, i int
	min = 0
	for i = 1; i < n; i++ {
		if A[i].targetDonasi < A[min].targetDonasi {
			min = i
		}
	}
	return min
}

// FUNCTION SEARCHING
func sequentialSearchIndeks(data tabProyek, n int, x string) int {
	var idx, i int

	idx = -1
	for i < n && idx == -1 {
		if data[i].nama == x || data[i].kategori == x {
			idx = i
		}
		i++
	}
	return idx
}

func sequentialSearchStatus(data tabProyek, n int) {
	var found bool
	var i, j int

	found = false
	for i = 0; i < n; i++ {
		if data[i].status {
			fmt.Println("=========================================================================================================")
			fmt.Printf("ID: %d || Judul: %s || Kategori: %s || Target Donasi: Rp %d || Total Donasi Terkumpul: Rp %d || Jumlah Donatur: %d\n", data[i].id, data[i].nama, data[i].kategori, data[i].targetDonasi, data[i].jumlahDonasi, data[i].jumlahOrang)
			fmt.Print("Deskripsi Proyek: ")
			for j = 0; j < 100 && data[i].deskripsiproyek[j] != "."; j++ {
				fmt.Print(data[i].deskripsiproyek[j], " ")
			}
			fmt.Println()
			found = true
		}
	}
	if !found {
		fmt.Println("Belum ada proyek yang mencapai target.")
	}
}

func BinSearchIndeks(data tabProyek, n int, x int) int {
	var left, right, mid int
	var idx int

	left = 0
	right = n - 1
	idx = -1

	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < data[mid].id {
			right = mid - 1
		} else if x > data[mid].id {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

// PROSEDUR SORTING
func insertionSort(A *tabProyek, n int) {
	var i, pass int
	var temp proyek
	if n < 1 {
		fmt.Println("Belum ada proyek yang tersedia.")
		return
	} else {
		pass = 1
		for pass <= n-1 {
			i = pass
			temp = A[pass]
			for i > 0 && temp.jumlahOrang < A[i-1].jumlahOrang {
				A[i] = A[i-1]
				i = i - 1
			}
			A[i] = temp
			pass = pass + 1
		}
	}
}

func selectionSort(A *tabProyek, n int) {
	var i, idx, pass int
	var temp proyek
	if n < 1 {
		fmt.Println("Belum ada proyek yang tersedia.")
		return
	} else {
		for pass = 0; pass < n-1; pass++ {
			idx = pass
			for i = pass + 1; i < n; i++ {
				if A[i].jumlahOrang > A[idx].jumlahOrang {
					idx = i
				}
			}
			temp = A[pass]
			A[pass] = A[idx]
			A[idx] = temp
		}
	}
}
