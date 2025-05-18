package main

import "fmt"

const MaxAssets = 100
const MaxLaporan = 100

type cryptoAsset struct {
	asetid    int
	nama      string
	diffculty float64
	reward    float64
	algoritma string
}

type HasilMining struct {
	idLaporan       int
	idAset          int
	namaAset        string
	rewardDidapat   float64
	waktuDibutuhkan float64 // dalam detik
	dayaDigunakan   float64 // dalam kWh
}
type dataPengguna struct {
	Dayakomputasipengguna float64
	Dayaperangkat         float64
}

var datap dataPengguna

type Tabasset [MaxAssets]cryptoAsset
type TabLaporan [MaxLaporan]HasilMining

var daftarAset Tabasset
var nAsset int = 0
var AssetId int = 1

var laporanM TabLaporan
var nLaporan int = 0
var LaporanId int = 1

func main() {
	var pilihMenu int
	menuAsset(&pilihMenu)
}
func menuAsset(pilihMenu *int) {
	*pilihMenu = 0
	for *pilihMenu != 6 {
		fmt.Println("\n-----------------------CRYPTO MINING SIMULATION------------------------")
		fmt.Println("Harap Atur spesifikasi pengguna terlebih dahulu.")
		fmt.Println("1. Kelola Asset Crypto")
		fmt.Println("2. Search CryptoAsset")
		fmt.Println("3. Atur Spesifikasi Pengguna")
		fmt.Println("4. Simulasi Mining")
		fmt.Println("5. Laporan Hasil Mining")
		fmt.Println("6. Quit")
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Print("Silahkan Pilih Menu: ")
		fmt.Scan(pilihMenu)

		switch *pilihMenu {
		case 1:
			fmt.Println("-----------------------------------------------------------------------")
			MenuKelolaAset(pilihMenu)
		case 2:
			fmt.Println("-----------------------------------------------------------------------")
			pencarianCrypto(pilihMenu)
		case 3:
			AturSpesifikasiPengguna()
		case 4:
			MenuMining()
		case 5:
			laporanMining()
		case 6:
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Println("      ♡Terimakasih telah menggunakan Crypto Mining Simulation♡         ")
			fmt.Println("-----------------------------------------------------------------------")
		default:
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Print("Pilihan tidak valid")
			fmt.Println("-----------------------------------------------------------------------")
		}
	}
}

func MenuKelolaAset(pilihMenu *int) {
	var pilihKelola int
	pilihKelola = 0
	for pilihKelola != 5 {
		fmt.Println("\n-----------------------CRYPTO MINING SIMULATION------------------------")
		fmt.Println("1. Tambah Asset Crypto")
		fmt.Println("2. Ganti Asset Crypto")
		fmt.Println("3. Hapus Asset Crypto")
		fmt.Println("4. Tampilkan Semua Asset")
		fmt.Println("5. Back")
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Print("Silahkan Pilih Menu: ")
		fmt.Scan(&pilihKelola)

		switch pilihKelola {
		case 1:
			fmt.Println("-----------------------------------------------------------------------")
			tambahAset(&daftarAset, &nAsset)
		case 2:
			fmt.Println("-----------------------------------------------------------------------")
			gantiAsset(&daftarAset, nAsset)
		case 3:
			fmt.Println("-----------------------------------------------------------------------")
			hapusAsset(&daftarAset, nAsset)
		case 4:
			fmt.Println("-----------------------------------------------------------------------")
			listAsset(daftarAset, nAsset)
		case 5:
			fmt.Println("-----------------------------------------------------------------------")
			*pilihMenu = 0
		default:
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Print("Pilihan tidak valid")

		}
	}
}
func pencarianCrypto(pilihMenu *int) {
	var pilihsearch int = 0
	for pilihsearch != 4 {
		fmt.Println("\n-----------------------CRYPTO MINING SIMULATION------------------------")
		fmt.Println("1. Search Nama Asset Crypto")
		fmt.Println("2. Urutkan Asset Crypto berdasarkan difficulty")
		fmt.Println("3. Urutkan Asset Crypto berdasarkan Reward")
		fmt.Println("4. Back")
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Print("Silahkan Pilih Menu: ")
		fmt.Scan(&pilihsearch)

		switch pilihsearch {
		case 1:
			fmt.Println("-----------------------------------------------------------------------")
			Searchmetod(pilihsearch)
		case 2:
			fmt.Println("-----------------------------------------------------------------------")
			selectionSortKesulitan(&daftarAset, nAsset)
		case 3:
			fmt.Println("-----------------------------------------------------------------------")
			insertionSortReward(&daftarAset, nAsset)
		case 4:
			fmt.Println("-----------------------------------------------------------------------")
			*pilihMenu = 0
		default:
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Print("Pilihan tidak valid")

		}
	}
}

// metode pencarian yang akan digunakan
func Searchmetod(pilihsearch int) {
	var pilihmetode int
	for pilihmetode != 3 {
		fmt.Println("\n-----------------------CRYPTO MINING SIMULATION------------------------")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Println("3. Back")
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Print("Silahkan Pilih Menu: ")
		fmt.Scan(&pilihmetode)

		switch pilihmetode {
		case 1:
			fmt.Println("-----------------------------------------------------------------------")
			var nama string
			fmt.Print("Masukkan nama Asset yang ingin ditemukan: ")
			fmt.Scan(&nama)
			idx := sequentialSearchByNama(*&daftarAset, nAsset, nama)
			if idx != -1 {
				fmt.Println("Aset ditemukan: ")
				fmt.Println("-----------------------------------------------------------------------")
				fmt.Printf("ID: %d, Nama: %s, Kesulitan: %.2f, Reward: %.2f, Algoritma: %s\n",
					daftarAset[idx].asetid, daftarAset[idx].nama, daftarAset[idx].diffculty, daftarAset[idx].reward, daftarAset[idx].algoritma)
			} else {
				fmt.Println("-----------------------------------------------------------------------")
				fmt.Println("Aset tidak ditemukan,silahkan periksa kembali.")
				fmt.Println("-----------------------------------------------------------------------")
			}

		case 2:
			fmt.Println("-----------------------------------------------------------------------")
			var nama string
			fmt.Print("Masukkan nama Asset yang ingin ditemukan: ")
			fmt.Scan(&nama)
			idx := binarySearchByNama(*&daftarAset, nAsset, nama)
			if idx != -1 {
				fmt.Println("Aset ditemukan: ")
				fmt.Printf("ID: %d, Nama: %s, Kesulitan: %d, Reward: %d, Algoritma: %s\n",
					daftarAset[idx].asetid, daftarAset[idx].nama, daftarAset[idx].diffculty, daftarAset[idx].reward, daftarAset[idx].algoritma)
			} else {
				fmt.Println("-----------------------------------------------------------------------")
				fmt.Println("Aset tidak ditemukan, silahkan periksa kembali.")
				fmt.Println("-----------------------------------------------------------------------")
			}
		case 3:
			fmt.Println("-----------------------------------------------------------------------")
			pilihsearch = 0
		default:
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Print("Pilihan tidak valid")

		}
	}
}
func AturSpesifikasiPengguna() {
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("                     Atur Spesifikasi Pengguna                         ")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Print("Masukkan Daya Komputasi (hash/s): ")
	fmt.Scan(&datap.Dayakomputasipengguna)
	fmt.Print("Masukkan Daya perangkat (watt): ")
	fmt.Scan(&datap.Dayaperangkat)
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("                    Spesifikasi berhasil diatur.                       ")
	fmt.Println("-----------------------------------------------------------------------")
}

// simulasi mining
func MenuMining() {
	stop := false
	if datap.Dayakomputasipengguna <= 0 || datap.Dayaperangkat <= 0 {
		fmt.Println("Harap atur spesifikasi pengguna terlebih dahulu melalui menu 3.")
		fmt.Println("-----------------------------------------------------------------------")
		stop = true
	}
	if nAsset == 0 {
		fmt.Println("Belum ada aset kripto untuk ditambang. Silakan tambahkan aset terlebih dahulu.")
		fmt.Println("-----------------------------------------------------------------------")
		stop = true
	}
	if nLaporan >= MaxLaporan {
		fmt.Println("Kapasitas laporan mining sudah penuh.")
		fmt.Println("-----------------------------------------------------------------------")
		stop = true
	}
	if !stop {
		fmt.Println()
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("              Silahkan pilih Aset untuk Simulasi Mining                ")
		fmt.Println("-----------------------------------------------------------------------")
		listAsset(daftarAset, nAsset)
		fmt.Print("Masukkan ID Aset yang ingin disimulasikan: ")
		var idAsetPilihan int
		fmt.Scan(&idAsetPilihan)

		idxAset := findAsetByID(idAsetPilihan, daftarAset, nAsset)
		if idxAset == -1 {
			fmt.Println("ID Aset tidak valid.")
			fmt.Println("-----------------------------------------------------------------------")
			stop = true
		}

		asetDipilih := daftarAset[idxAset]
		estimasiWaktuDetik := asetDipilih.diffculty * 1000 / datap.Dayakomputasipengguna
		estimasiDayaKWH := estimasiWaktuDetik * datap.Dayaperangkat

		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("                     Hasil Estimasi Mining                             ")
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Printf("Aset: %s)\n", asetDipilih.nama)
		fmt.Printf("Estimasi Waktu Mining: %.2f \n", estimasiWaktuDetik)
		fmt.Printf("Estimasi Daya Listrik Digunakan: %.2f kWh\n", estimasiDayaKWH)
		fmt.Printf("Estimasi Reward yang Akan Didapat: %.2f \n", asetDipilih.reward)
		fmt.Println("-----------------------------------------------------------------------")

		fmt.Print("Simulasikan mining untuk aset ini dan catat hasilnya? (ya/tidak): ")
		var konfirmasiSimulasi string
		fmt.Scan(&konfirmasiSimulasi)

		if konfirmasiSimulasi == "ya" {
			laporanM[nLaporan].idLaporan = LaporanId
			laporanM[nLaporan].idAset = asetDipilih.asetid
			laporanM[nLaporan].namaAset = asetDipilih.nama
			laporanM[nLaporan].rewardDidapat = asetDipilih.reward
			laporanM[nLaporan].waktuDibutuhkan = estimasiWaktuDetik
			laporanM[nLaporan].dayaDigunakan = estimasiDayaKWH
			nLaporan++
			LaporanId++
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Println("           Simulasi mining berhasil dicatat dalam laporan.             ")
			fmt.Println("-----------------------------------------------------------------------")
		} else {
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Println("                   Simulasi mining dibatalkan.                         ")
			fmt.Println("-----------------------------------------------------------------------")
		}
	}
}

// Menu laporan Mining
func laporanMining() {
	if nLaporan == 0 {
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("                 Belum ada data laporan mining.                        ")
	} else {
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("                   Laporan Hasil Mining Pengguna")
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Printf("%-15s %-15s %-15s %-15s %-15s \n", "ID Laporan", "Nama Aset", "Reward ", "Waktu", "Daya")
		fmt.Println("-----------------------------------------------------------------------")
		var totalReward float64 = 0
		var totalWaktu float64 = 0
		var totalDaya float64 = 0
		for i := 0; i < nLaporan; i++ {
			lpr := laporanM[i]
			fmt.Printf("%-15d %-15s %-15.2f %-15.2f %-15.2f\n",
				lpr.idLaporan, lpr.namaAset, lpr.rewardDidapat, lpr.waktuDibutuhkan, lpr.dayaDigunakan)
			totalReward += lpr.rewardDidapat
			totalWaktu += lpr.waktuDibutuhkan
			totalDaya += lpr.dayaDigunakan
		}
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("Total Keseluruhan:")
		fmt.Printf("Total Reward: %.2f\n", totalReward)
		fmt.Printf("Total Waktu Mining: %.2f detik (%.2f menit / %.2f jam)\n", totalWaktu, totalWaktu/60.0, totalWaktu/3600.0)
		fmt.Printf("Total Daya Listrik Digunakan: %.2f kWh\n", totalDaya)
	}
	fmt.Println("-----------------------------------------------------------------------")
}

// menambahkan Asset
func tambahAset(Asset *Tabasset, nAsset *int) {
	full := false
	if *nAsset > MaxAssets {
		fmt.Print("Kapasitas Penuh")
		fmt.Println("-----------------------------------------------------------------------")
		full = true
	}
	if !full {
		var aset cryptoAsset
		aset.asetid = AssetId
		fmt.Print("Masukkan Nama Asset:")
		fmt.Scan(&aset.nama)
		fmt.Print("Masukkan Tingkat Kesulitan: ")
		fmt.Scan(&aset.diffculty)
		fmt.Print("Masukkan Estimasi Reward: ")
		fmt.Scan(&aset.reward)
		fmt.Print("Masukkan Algotitma yang digunakan: ")
		fmt.Scan(&aset.algoritma)

		Asset[*nAsset] = aset
		*nAsset++
		AssetId++
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("Asset Berhasil Ditambahkan.")
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Print("id Asset: ", aset.asetid)
		fmt.Print("\n-----------------------------------------------------------------------")
	}
}

// mencari asset dengan ID
func findAsetByID(id int, data Tabasset, n int) int {
	for i := 0; i < n; i++ {
		if data[i].asetid == id {
			return i
		}
	}
	return -1
}

// Hapus Asset
func hapusAsset(Asset *Tabasset, nAsset int) {
	stop := false
	if nAsset == 0 {
		fmt.Println("Tidak ada Asset Crypto yang bisa di hapus")
		fmt.Println("-----------------------------------------------------------------------")
		stop = true
	}
	if !stop {
		var idAset int
		fmt.Print("Masukkan ID aset kripto yang ingin dihapus: ")
		fmt.Scan(&idAset)

		idx := findAsetByID(idAset, *Asset, nAsset)

		if idx == -1 {
			fmt.Print("Aset Crypto dengan ID tersebut tidak ditemukan.")
			fmt.Println("\n-----------------------------------------------------------------------")
		} else {

			fmt.Printf("Anda yakin ingin menghapus aset: ID %d, Nama %s? (ya/tidak): ", Asset[idx].asetid, Asset[idx].nama)
			var konfirmasi string
			fmt.Scan(&konfirmasi)

			if konfirmasi == "ya" {
				for i := idx; i < nAsset-1; i++ {
					Asset[i] = Asset[i+1]
				}
				nAsset--
				fmt.Println("-----------------------------------------------------------------------")
				fmt.Println("Aset kripto berhasil dihapus.")
				fmt.Println("\n-----------------------------------------------------------------------")
			} else {
				fmt.Println("\n-----------------------------------------------------------------------")
				fmt.Println("Penghapusan aset dibatalkan.")
				fmt.Println("\n-----------------------------------------------------------------------")
			}
		}
	}
}

// gantiAsset
func gantiAsset(Asset *Tabasset, nAsset int) {
	stop := false
	if nAsset == 0 {
		fmt.Print("Tidak ada Asset Crypto yang bisa di ganti.")
		fmt.Println("\n-----------------------------------------------------------------------")
		stop = true
	}
	if !stop {
		var idAset int
		fmt.Print("Masukkan ID aset Crypto yang ingin diubah: ")
		fmt.Scan(&idAset)

		idx := findAsetByID(idAset, *Asset, nAsset)

		if idx == -1 {
			fmt.Println("\n-----------------------------------------------------------------------")
			fmt.Print("Aset Crypto dengan ID tersebut tidak ditemukan.")
			fmt.Println("\n-----------------------------------------------------------------------")
		} else {

			fmt.Println("Data aset saat ini:")
			fmt.Printf("%-5s %-10s %-20s %-20s %-15s \n", "ID", "Nama", "Tingkat Kesulitan", "Estimasi Reward", "Algoritma")
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Printf("%-5d %-10s %-20d %-20d %-15s\n", Asset[idx].asetid, Asset[idx].nama, Asset[idx].diffculty, Asset[idx].reward, Asset[idx].algoritma)
			fmt.Println("Masukkan data baru (kosongkan jika tidak ingin diubah, masukkan '-' untuk string, 0 untuk angka):")

			var tempNama, tempAlgoritma string
			var tempKesulitan float64
			var tempReward float64

			fmt.Print("Nama Aset baru: ")
			fmt.Scan(&tempNama)
			fmt.Print("Tingkat Kesulitan baru: ")
			fmt.Scan(&tempKesulitan)
			if tempKesulitan != 0 {
				Asset[idx].diffculty = tempKesulitan
			}

			fmt.Print("Estimasi Reward baru: ")
			fmt.Scan(&tempReward)
			if tempReward != 0 {
				Asset[idx].reward = tempReward
			}
			fmt.Print("Algoritma baru: ")
			fmt.Scan(&tempAlgoritma)
			if tempAlgoritma != "-" && tempAlgoritma != "" {
				Asset[idx].algoritma = tempAlgoritma
			}
			fmt.Println("\n-----------------------------------------------------------------------")
			fmt.Println("                  Data aset kripto berhasil diubah.                      ")
			fmt.Println("\n-----------------------------------------------------------------------")
		}
	}

}

// list asset
func listAsset(Asset Tabasset, nAsset int) {

	if nAsset == 0 {
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("Tidak ada data aset kripto.")
		fmt.Println("-----------------------------------------------------------------------")
	} else {
		fmt.Printf("%-5s %-10s %-20s %-20s %-15s \n", "ID", "Nama", "Tingkat Kesulitan", "Estimasi Reward", "Algoritma")
		fmt.Println("-----------------------------------------------------------------------")
		for i := 0; i < nAsset; i++ {
			fmt.Printf("%-5d %-10s %-20.2f %-20.2f %-15s\n", Asset[i].asetid, Asset[i].nama, Asset[i].diffculty, Asset[i].reward, Asset[i].algoritma)
		}
	}
}

// cari nama sequential & binary
func sequentialSearchByNama(Asset Tabasset, n int, nama string) int {
	for i := 0; i < n; i++ {
		if Asset[i].nama == nama {
			return i
		}
	}
	return -1
}

func binarySearchByNama(Asset Tabasset, n int, nama string) int {
	var left, right, mid int
	var idx int

	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if nama < Asset[mid].nama {
			right = mid - 1
		} else if nama > Asset[mid].nama {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

// urutkan data
func selectionSortKesulitan(Asset *Tabasset, n int) { //ascending
	stop := false
	if n == 0 {
		fmt.Println("          Asset belum ada, silahkan tambahkan pada menu utama          ")
		fmt.Println("-----------------------------------------------------------------------")
		stop = true
	}
	if !stop {
		for i := 0; i < n-1; i++ {
			maxIdx := i
			for j := i + 1; j < n; j++ {
				if (*Asset)[j].diffculty > (*Asset)[maxIdx].diffculty {
					maxIdx = j
				}
			}
			(*Asset)[i], (*Asset)[maxIdx] = (*Asset)[maxIdx], (*Asset)[i]
		}
		fmt.Printf("%-5s %-10s %-20s %-20s %-15s \n", "ID", "Nama", "Tingkat Kesulitan", "Estimasi Reward", "Algoritma")
		fmt.Println("-----------------------------------------------------------------------")
		for i := 0; i < n; i++ {
			fmt.Printf("%-5d %-10s %-20d %-20.2f %-15s\n", Asset[i].asetid, Asset[i].nama, Asset[i].diffculty, Asset[i].reward, Asset[i].algoritma)
		}
		fmt.Println("-----------------------------------------------------------------------")
	}
}
func insertionSortReward(Asset *Tabasset, n int) { //descanding
	stop := false
	if n == 0 {
		fmt.Println("          Asset belum ada, silahkan tambahkan pada menu utama          ")
		fmt.Println("-----------------------------------------------------------------------")
		stop = true
	}
	if !stop {
		for i := 1; i < n; i++ {
			temp := (*Asset)[i]
			j := i - 1
			for j >= 0 && (*Asset)[j].reward < temp.reward {
				(*Asset)[j+1] = (*Asset)[j]
				j--
			}
			(*Asset)[j+1] = temp
		}
		fmt.Printf("%-5s %-10s %-20s %-20s %-15s \n", "ID", "Nama", "Tingkat Kesulitan", "Estimasi Reward", "Algoritma")
		fmt.Println("-----------------------------------------------------------------------")
		for i := 0; i < n; i++ {
			fmt.Printf("%-5d %-10s %-20d %-20.2f %-15s\n", Asset[i].asetid, Asset[i].nama, Asset[i].diffculty, Asset[i].reward, Asset[i].algoritma)
		}
		fmt.Println("-----------------------------------------------------------------------")
	}
}
