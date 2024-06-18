package main

import (
	"fmt"
)

func belanja() {
	var nama_barang string
	var id_barang, opsi_sub, banyak_barang int
	var keranjang [NMAX]data
	var tampil bool
	opsi_sub = 1

	for opsi_sub == 1 {
		view_belanja()
		fmt.Scan(&nama_barang)
		header_tabel_database()
		tampil = cari_barang(nama_barang)

		if tampil {
			fmt.Print("Masukkan id barang yang ingin dibeli: ")
			fmt.Scan(&id_barang)
			if id_barang != -1 {
				cari_id_barang(&id_barang)
				fmt.Print("Masukkan banyak barang yang ingin dibeli: ")
				fmt.Scan(&banyak_barang)
				isi_keranjang(&keranjang, id_barang, banyak_barang)

			}
		}

		fmt.Println("Pilih opsi:")
		fmt.Println("1. Tambah barang")
		fmt.Println("2. Bayar")
		fmt.Println("0. Cancel")
		fmt.Print("pilih: ")
		fmt.Scan(&opsi_sub)
		if opsi_sub == 2 {
			bayar(keranjang)
			catat_Log(id_barang, banyak_barang)
		}
	}

}
func cari_id_barang(idx *int) {
	var kosong data
	for i := 0; gudang[i] != kosong; i++ {
		if *idx == gudang[i].index {
			*idx = i
		}
	}
}

func bayar(keranjang [NMAX]data) {
	view_bayar(keranjang)
	var kosong data
	for i := 0; keranjang[i] != kosong; i++ {
		for j := 0; gudang[j] != kosong; j++ {
			if keranjang[i].index == gudang[j].index {
				gudang[j].stok -= keranjang[i].stok
			}

		}
	}
}

func isi_keranjang(keranjang *[NMAX]data, x, n int) {
	var i int
	var kosong data
	for i = 0; keranjang[i] != kosong; i++ {
	}

	keranjang[i] = gudang[x]
	keranjang[i].stok = n

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Printf("%-11s %-35s %-16s %-15s %s \n", "ID", "Nama Barang", "Kategori", "Harga", "Jumlah")
	fmt.Println("----------------------------------------------------------------------------------------")
	for k := 0; keranjang[k] != kosong; k++ {
		fmt.Printf("%-11d %-35s %-16s %-15d %d \n", keranjang[i].index, keranjang[i].nama, keranjang[i].kategori, keranjang[i].harga, keranjang[i].stok)
	}

}

func cari_barang(barang string) bool {
	var bit_barang int = biner_barang(barang)
	var nData int = panjang_database()
	var i int = 0
	var bit_data int
	var tampil bool = false

	for i < nData {
		bit_data = biner_barang(gudang[i].nama)
		if bit_barang == bit_data {
			fmt.Printf("%-11d %-35s %-16s %-15d %d \n", gudang[i].index, gudang[i].nama, gudang[i].kategori, gudang[i].harga, gudang[i].stok)
			tampil = true
		}
		i++
	}
	if !tampil {
		fmt.Printf("%-11d %-35s %-16s %-15d %d \n", -1, "Tidak Ditemukan", "", 0, 0)
	}
	return tampil

}

func biner_barang(x string) int {
	var bit []byte = []byte(x)
	var n int = len(x)
	var jmlh int
	uppercase(bit, n)
	for i := 0; i < n && bit[i] != ' '; i++ {
		jmlh += int(bit[i])
	}

	return jmlh
}

func uppercase(bit []byte, n int) {
	for i := 0; i < n; i++ {
		if bit[i] >= 'a' && bit[i] <= 'z' {
			bit[i] -= 32
		}
	}
}

func panjang_database() int {
	var kosong data
	var i int = 0
	for gudang[i] != kosong {
		i++
	}
	return i
}

func urut_id_naik() {
	var kosong data
	var max int
	for i := 1; gudang[i] != kosong; i++ {
		max = i - 1
		j := i
		for gudang[j] != kosong {
			if gudang[j].index < gudang[max].index {
				max = j
			}
			j++
		}
		gudang[max], gudang[i-1] = gudang[i-1], gudang[max]
	}
	header_tabel_database()
	show_database()
}

func urut_id_turun() {
	var kosong data
	var min int
	for i := 1; gudang[i] != kosong; i++ {
		min = i - 1
		j := i
		for gudang[j] != kosong {
			if gudang[j].index > gudang[min].index {
				min = j
			}
			j++
		}
		gudang[min], gudang[i-1] = gudang[i-1], gudang[min]
	}
	header_tabel_database()
	show_database()
}

func urut_produk_naik() {
	var kosong data

	for i := 1; gudang[i] != kosong; i++ {
		j := i
		for j > 0 {
			// if biner_barang(gudang[j].nama) < biner_barang(gudang[j-1].nama) {
			// 	gudang[j], gudang[j-1] = gudang[j-1], gudang[j]
			// }
			if gudang[j].nama < gudang[j-1].nama {
				gudang[j], gudang[j-1] = gudang[j-1], gudang[j]
			}
			j--
		}
	}
	header_tabel_database()
	show_database()
}

func urut_produk_turun() {
	var kosong data

	for i := 1; gudang[i] != kosong; i++ {
		j := i
		for j > 0 {
			// if biner_barang(gudang[j].nama) > biner_barang(gudang[j-1].nama) {
			// 	gudang[j], gudang[j-1] = gudang[j-1], gudang[j]
			// }
			if gudang[j].nama > gudang[j-1].nama {
				gudang[j], gudang[j-1] = gudang[j-1], gudang[j]
			}
			j--
		}
	}
	header_tabel_database()
	show_database()
}

func urut_kategori_naik() {
	var kosong data

	for i := 1; gudang[i] != kosong; i++ {
		j := i
		for j > 0 {
			if gudang[j].kategori < gudang[j-1].kategori {
				gudang[j], gudang[j-1] = gudang[j-1], gudang[j]
			} else if gudang[j].kategori == gudang[j-1].kategori && biner_barang(gudang[j].nama) < biner_barang(gudang[j-1].nama) {
				gudang[j], gudang[j-1] = gudang[j-1], gudang[j]
			}
			j--
		}
	}
	header_tabel_database()
	show_database()
}

func urut_kategori_turun() {
	var kosong data

	for i := 1; gudang[i] != kosong; i++ {
		j := i
		for j > 0 {
			if gudang[j].kategori > gudang[j-1].kategori {
				gudang[j], gudang[j-1] = gudang[j-1], gudang[j]
			} else if gudang[j].kategori == gudang[j-1].kategori && biner_barang(gudang[j].nama) > biner_barang(gudang[j-1].nama) {
				gudang[j], gudang[j-1] = gudang[j-1], gudang[j]
			}
			j--
		}
	}
	header_tabel_database()
	show_database()
}

func create_data() {
	var kosong data
	// var nilInt int = 0
	// var nilStr string = ""
	var n int

	for n = 0; gudang[n] != kosong; n++ {
		//cari limit index paling belakangnya
	}
	fmt.Println("Masukkan data dengan urutan sebagai berikut:")
	fmt.Println("Masukkan ID barang:")
	fmt.Scan(&gudang[n].index)
	fmt.Println("Masukkan NAMA barang:")
	fmt.Println("Note* Spasi diganti dengan '_'")
	fmt.Scan(&gudang[n].nama)
	UnderscoreToSpace(&gudang[n].nama)
	fmt.Println("Masukkan KATEGORI barang:")
	fmt.Scan(&gudang[n].kategori)
	fmt.Println("Masukkan HARGA barang:")
	fmt.Scan(&gudang[n].harga)
	fmt.Println("Masukkan STOK barang:")
	fmt.Scan(&gudang[n].stok)

}

func update_data() {
	var id int
	var kosong data
	fmt.Println("Masukkan ID barang yang ingin di update:")
	fmt.Scan(&id)

	for i := 0; gudang[i] != kosong; i++ {
		if gudang[i].index == id {
			fmt.Println("Masukkan data yang ingin di ubah dengan urutan sebagai berikut:")
			fmt.Println("Masukkan ID barang baru:")
			fmt.Scan(&gudang[i].index)
			fmt.Println("Masukkan NAMA barang baru:")
			fmt.Println("Note* Spasi diganti dengan '_'")
			fmt.Scan(&gudang[i].nama)
			UnderscoreToSpace(&gudang[i].nama)
			fmt.Println("Masukkan KATEGORI barang baru:")
			fmt.Scan(&gudang[i].kategori)
			fmt.Println("Masukkan HARGA barang baru:")
			fmt.Scan(&gudang[i].harga)
			fmt.Println("Masukkan STOK barang baru:")
			fmt.Scan(&gudang[i].stok)
		}
	}
}

func delete_data() {
	var id, n int
	var yakin string
	var kosong data
	for n = 0; gudang[n] != kosong; n++ {
		// cari banyak isi tabelnya
	}
	fmt.Println("Masukkan ID barang yang ingin di hapus:")
	fmt.Scan(&id)
	fmt.Println("Apakah anda yakin untuk menghapus barang?")
	fmt.Println("YES untuk melanjutkan")
	fmt.Println("NO untuk kembali")
	fmt.Scan(&yakin)
	if yakin == "YES" || yakin == "yes" {
		for i := 0; i < n; i++ {
			if gudang[i].index == id {
				for i < n {
					gudang[i] = gudang[i+1]
					i++
				}
			}
		}
	} else {

	}

}

func UnderscoreToSpace(x *string) {
	var n int = len(*x)
	var kata []byte
	var output string
	kata = []byte(*x)

	for i := 0; i < n; i++ {
		if kata[i] == '_' {
			kata[i] = ' '
		}
		output += string(kata[i])
	}
	kata = []byte(output)
	uppercase(kata, n)
	*x = string(kata)
}

func catat_Log(id, jumlah int){
	i := panjang_log()
	harga := gudang[id].harga * jumlah
	total := harga + (harga * 10) / 100
	log[i].index = i+1
	log[i].nama = gudang[id].nama
	log[i].kategory = gudang[id].kategori
	log[i].jumlah = jumlah
	log[i].total = total
}

func panjang_log()int{
	var kosong logData
	s := 0
	for i:= 0;log[i] != kosong;i++{
		s++
	}
	return s
}

func logMenaik(){
	n := panjang_log()
	pass := 1
	for pass <= n-1{
		max := pass-1
		i := pass
		for i < n{
			if log[max].total < log[i].total {
				max = i
			}
			i++
		}
		temp := log[pass-1]
		log[pass-1] = log[max]
		log[max] = temp
		pass++
	}
	fmt.Println("Sorting Log Menaik (ASCENDING)")
	logHeader()
}

func logMenurun(){
	n := panjang_log()
	pass := 1
	for pass <= n-1{
		i := pass
		temp := log[pass]
		for i > 0 && temp.total < log[i-1].total{
			log[i] = log[i-1]
			i--
		}
		log[i] = temp
		pass++
	}
	fmt.Println("Sorting Log Menurun (DESCENDING)")
	logHeader()
}

func validasi(username, password string) bool {
	i := 0
	var isLogin bool
	for i < jumlah_akun(){
		if username == tabLogin[i].username && password == tabLogin[i].password{
			isLogin = true
		}else{
			isLogin = false
		}
		i++
	}
	return isLogin
}

func jumlah_akun()int{
	var kosong loginData
	a := 0
	for i := 0; tabLogin[i] != kosong; i++ {
		a++
	}
	return a
}

func register(username, password string){
	i := jumlah_akun()
	tabLogin[i].username = username
	tabLogin[i].password = password
	fmt.Println("Akun Telah Dibuat Harap login")
	login()
}

func cekDuplikat(username, password string)bool{
	var duplikat bool
	i := 0
	for i < jumlah_akun(){
		if tabLogin[i].username == username{
			duplikat = true
		}else{
			duplikat = false
		}
		i++
	}
	return duplikat
}