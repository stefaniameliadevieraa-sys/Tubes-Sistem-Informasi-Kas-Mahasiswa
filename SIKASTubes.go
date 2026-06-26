package main

import "fmt"

const IURAN = 5000
const JUMLAH_BULAN = 24

type Transaksi struct {
	NIM      string
	Tanggal  string
	Nominal  int
}

var riwayat [1000]Transaksi
var nTransaksi int

type Mahasiswa struct {
	NIM           string
	Nama          string
	NominalBayar  int
	TanggalBayar  string
	Tunggakan     int
	Status        string
}

var data [100]Mahasiswa
var n int

func tambahMahasiswa() {
	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&data[n].NIM)

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&data[n].Nama)

	data[n].NominalBayar = 0
	data[n].TanggalBayar = "-"
	data[n].Tunggakan = IURAN * JUMLAH_BULAN
	data[n].Status = "Belum Lunas"

	n++

	fmt.Println("Data berhasil ditambah")
}

func ubahMahasiswa() {
	var nimCari string

	fmt.Print("Masukkan NIM yang akan diubah: ")
	fmt.Scan(&nimCari)

	idx := -1

	for i := 0; i < n; i++ {
		if data[i].NIM == nimCari {
			idx = i
			break
		}
	}

	if idx != -1 {
		fmt.Print("Nama baru: ")
		fmt.Scan(&data[idx].Nama)

		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusMahasiswa() {
	var nimCari string
	idx := -1

	fmt.Print("Masukkan NIM yang akan dihapus: ")
	fmt.Scan(&nimCari)

	for i := 0; i < n; i++ {
		if data[i].NIM == nimCari {
			idx = i
			break
		}
	}

	if idx != -1 {
		for i := idx; i < n-1; i++ {
			data[i] = data[i+1]
		}

		n--

		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func bayarKas() {
	var nimCari string
	var bayar int

	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nimCari)

	idx := -1

	for i := 0; i < n; i++ {
		if data[i].NIM == nimCari {
			idx = i
			break
		}
	}

	if idx != -1 {

		fmt.Print("Nominal Pembayaran: ")
		fmt.Scan(&bayar)

		fmt.Print("Tanggal Pembayaran: ")
		fmt.Scan(&data[idx].TanggalBayar)

		data[idx].NominalBayar += bayar

		totalTagihan := IURAN * JUMLAH_BULAN
		data[idx].Tunggakan = totalTagihan - data[idx].NominalBayar

		if data[idx].Tunggakan <= 0 {
			data[idx].Tunggakan = 0
			data[idx].Status = "Lunas"
		} else {
			data[idx].Status = "Belum Lunas"
		}

		riwayat[nTransaksi].NIM = data[idx].NIM
		riwayat[nTransaksi].Tanggal = data[idx].TanggalBayar
		riwayat[nTransaksi].Nominal = bayar
		nTransaksi++

		fmt.Println("Pembayaran berhasil")

	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func tampilkanData() {
	if n == 0 {
		fmt.Println("Belum ada data mahasiswa")
		return
	}

	fmt.Println("\n=== Data Mahasiswa ===")

	for i := 0; i < n; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("NIM             :", data[i].NIM)
		fmt.Println("Nama            :", data[i].Nama)
		fmt.Println("Nominal Bayar   :", data[i].NominalBayar)
		fmt.Println("Tanggal Bayar   :", data[i].TanggalBayar)
		fmt.Println("Tunggakan       :", data[i].Tunggakan)
		fmt.Println("Status          :", data[i].Status)
		fmt.Println("-----------------------------------------")
	}
}

func tampilRiwayat() {
	if nTransaksi == 0 {
		fmt.Println("Belum ada transaksi")
		return
	}

	fmt.Println("\n=== Riwayat Transaksi ===")

	for i := 0; i < nTransaksi; i++ {
		fmt.Println("Transaksi ke-", i+1)
		fmt.Println("NIM      :", riwayat[i].NIM)
		fmt.Println("Tanggal  :", riwayat[i].Tanggal)
		fmt.Println("Nominal  :", riwayat[i].Nominal)
		fmt.Println("------------------------")
	}
}

func selectionSortNamaAsc() {
	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if data[j].Nama < data[min].Nama {
				min = j
			}
		}

		data[i], data[min] = data[min], data[i]
	}

	fmt.Println("Data sudah diurutkan")
}

func insertionSortTunggakanDesc() {
	for i := 1; i < n; i++ {
		temp := data[i]
		j := i - 1

		for j >= 0 && data[j].Tunggakan < temp.Tunggakan {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}

	fmt.Println("Data sudah diurutkan")
}

func sequentialSearchBelumLunas() {
	found := false

	fmt.Println("\n=== Mahasiswa Belum Lunas ===")

	for i := 0; i < n; i++ {
		if data[i].Status == "Belum Lunas" {
			fmt.Println(data[i].NIM, "-", data[i].Nama,
				"| Tunggakan:", data[i].Tunggakan)
			found = true
		}
	}

	if !found {
		fmt.Println("Semua mahasiswa sudah lunas")
	}
}

func binarySearchNIM() {
	var nimCari string

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j].NIM > data[j+1].NIM {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	fmt.Print("Masukkan NIM yang dicari: ")
	fmt.Scan(&nimCari)

	low := 0
	high := n - 1
	found := false

	for low <= high {
		mid := (low + high) / 2

		if data[mid].NIM == nimCari {
			fmt.Println("\nData Ditemukan")
			fmt.Println("NIM :", data[mid].NIM)
			fmt.Println("Nama :", data[mid].Nama)
			fmt.Println("Nominal Bayar :", data[mid].NominalBayar)
			fmt.Println("Tanggal Bayar :", data[mid].TanggalBayar)
			fmt.Println("Tunggakan :", data[mid].Tunggakan)
			fmt.Println("Status :", data[mid].Status)

			found = true
			break

		} else if nimCari < data[mid].NIM {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func statistik() {
	totalKas := 0
	jumlahLunas := 0

	for i := 0; i < n; i++ {
		totalKas += data[i].NominalBayar

		if data[i].Status == "Lunas" {
			jumlahLunas++
		}
	}

	fmt.Println("\n=== Statistik Kas ===")
	fmt.Println("Total Saldo Kas :", totalKas)
	fmt.Println("Jumlah Mahasiswa Lunas :", jumlahLunas)
}

func main() {
	var pilih int

	for {
		fmt.Println("\n==== Informasi Kas Mahasiswa ===")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Ubah Data Mahasiswa")
		fmt.Println("3. Hapus Data Mahasiswa")
		fmt.Println("4. Pembayaran Kas")
		fmt.Println("5. Tampilkan Data")
		fmt.Println("6. Tampilkan Riwayat Transaksi")
		fmt.Println("7. Sequential Search (Belum Lunas)")
		fmt.Println("8. Binary Search (NIM)")
		fmt.Println("9. Selection Sort Nama")
		fmt.Println("10. Insertion Sort Tunggakan")
		fmt.Println("11. Statistik")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahMahasiswa()
		} else if pilih == 2 {
			ubahMahasiswa()
		} else if pilih == 3 {
			hapusMahasiswa()
		} else if pilih == 4 {
			bayarKas()
		} else if pilih == 5 {
			tampilkanData()
		} else if pilih == 6 {
			tampilRiwayat()
		} else if pilih == 7 {
			sequentialSearchBelumLunas()
		} else if pilih == 8 {
			binarySearchNIM()
		} else if pilih == 9 {
			selectionSortNamaAsc()
		} else if pilih == 10 {
			insertionSortTunggakanDesc()
		} else if pilih == 11 {
			statistik()
		} else if pilih == 0 {
			break
		} else {
			fmt.Println("Menu tidak tersedia")
		}
	}
}