package main

import "fmt"

func belanja(){
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
		fmt.Println("3. Cancel")
		fmt.Print("pilih: ")
		fmt.Scan(&opsi_sub)
		if opsi_sub == 2 {
			bayar(keranjang)
		} else if opsi_sub == 3 {
			index()
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
	
}

func isi_keranjang(keranjang *[NMAX]data, x, n int) {
	var i int
	var kosong data
	for i = 0; keranjang[i] != kosong; i++ {
	}

	keranjang[i] = gudang[x]
	keranjang[i].stok = n

	fmt.Printf("%-11s %-35s %-16s %-15s %s \n", "ID", "Nama Barang", "Kategori", "Harga", "Banyak")
	for k := 0; keranjang[k] != kosong; k++{
		fmt.Printf("%-11d %-35s %-16s %-15d %d \n", keranjang[i].index, keranjang[i].nama, keranjang[i].kategori, keranjang[i].harga, keranjang[i].stok)
	}

}


func cari_barang(barang string) bool{
	var bit_barang int = biner_barang(barang)
	var nData int = panjang_database()
	var i int = 0
	var bit_data int
	var tampil bool= false

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

func biner_barang(x string)int {
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

func panjang_database() int{
	var kosong data
	var i int = 0
	for gudang[i] != kosong{
		i++
	}
	return i
}

func urut_id_naik(){
	var kosong data
	var max int
	for i := 1; gudang[i] != kosong; i++ {
		max = i-1
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

func urut_id_turun(){
	var kosong data
	var min int 
	for i := 1; gudang[i] != kosong; i++ {
		min = i-1
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

func urut_produk_naik(){
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

func urut_produk_turun(){
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

func urut_kategori_naik(){
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

func urut_kategori_turun(){
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

}

func update_data() {

}

func delete_data(){

}