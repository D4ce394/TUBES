package main

import (
	"fmt"
)

func index() {
	var pilih int = -1
	for pilih != 0 {
		fmt.Println("----------------------------------------")
		fmt.Println("                 DR.Mart                ")
		fmt.Println("")
		fmt.Println("Pilih:")
		fmt.Println("1. Belanja")
		fmt.Println("2. List Data Barang")
		fmt.Println("3. Log")
		fmt.Println("0. Exit")
		fmt.Println("")
		fmt.Println("----------------------------------------")
		fmt.Scan(&pilih)
		if pilih == 1 {
			belanja()
		} else if pilih == 2 {
			list_barang()
		} else if pilih == 3 {
			logHeader()
		}
	}
}

func view_belanja() {
	fmt.Println("----------------------------------------")
	fmt.Println("              View Belanja              ")
	fmt.Println("----------------------------------------")
	fmt.Print("Nama Barang : ") // search barang dari data

	// nampilin hasil search barang

}

func view_bayar(keranjang [NMAX]data) {
	var kosong data
	var tBarang [NMAX]int
	var tBelanja, uang int
	fmt.Printf("%-35s %-15s %-18s %s \n", "Nama Barang", "Banyak", "Harga", "Total")
	for i := 0; keranjang[i] != kosong; i++ {
		tBarang[i] = keranjang[i].harga * keranjang[i].stok
		tBelanja += tBarang[i]
		fmt.Printf("%-35s %-15d Rp.%-15d Rp.%d \n", keranjang[i].nama, keranjang[i].stok, keranjang[i].harga, tBarang[i])
	}
	tBelanja += tBelanja / 10
	fmt.Println()
	fmt.Printf("%-35s %-15s Rp.%-15d %s \n", "PPN 10%", "", tBelanja/10, "")
	fmt.Printf("%-35s %-15s %-18s Rp.%d \n", "Total Belanja", "", "", tBelanja)
	fmt.Print("Uang yang dibayarkan: ")
	fmt.Scan(&uang)
	for uang < tBelanja {
		fmt.Println("Uang kurang")
		fmt.Print("Uang yang dibayarkan: ")
		fmt.Scan(&uang)
	}
	fmt.Println("Uang kembalian:", uang-tBelanja)
}

func list_barang() {
	var opsi int = -1
	for opsi != 0 {
		fmt.Println("               List Barang              ")
		fmt.Println("----------------------------------------")
		header_tabel_database()
		show_database()
		fmt.Println("Opsi :")
		fmt.Println("1. Urutkan tampilan")
		fmt.Println("2. Tambah data")
		fmt.Println("3. Update data")
		fmt.Println("4. Delete data")
		fmt.Println("0. Menu")
		fmt.Scan(&opsi)
		if opsi == 1 {
			view_urutan()
		} else if opsi == 2 {
			create_data()
		} else if opsi == 3 {
			update_data()
		} else if opsi == 4 {
			delete_data()
		}
		fmt.Println("----------------------------------------")
	}
}

func logHeader() {
	var omset int
	fmt.Println("               Log Barang               ") // nampilin data log pembelanjaan-------------
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Printf("%-11s %-35s %-16s %-15s %s \n", "No", "Nama Barang", "Kategori", "Banyak", "Harga")
	fmt.Println("-------------------------------------------------------------------------------------------")
	for i:=0; i < panjang_log(); i++{
		fmt.Printf("%-11d %-35s %-16s %-15d %d \n", log[i].index, log[i].nama, log[i].kategory, log[i].jumlah, log[i].total)
			omset = omset + log[i].total
	}
	log_footer(omset)
	menu_log()

}

func log_footer(omset int){
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Printf("%23s %62d \n", "Total Omset", omset)
	fmt.Printf("%33s %52d \n", "Total Omset setelah PPH 10%", omset-(omset*10)/100)
}

func menu_log(){
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("Menu Log")	
	fmt.Println()
	fmt.Println("Pilih: ")
	fmt.Println("1. Urut Log Menaik (ASCENDING)")
	fmt.Println("2. Urut Log Menurun (DESCENDING)")
	fmt.Println("0. Kembali")

	var pilihLog int
	fmt.Scan(&pilihLog)
	for pilihLog != 0{
		if pilihLog == 1 {
			logMenaik()
		}else if pilihLog == 2 {
			logMenurun()
		}
	}
	index()
}



func header_tabel_database() {
	// fmt.Print(" NO ")
	// fmt.Print("        Nama Barang         ")
	// fmt.Print("  Kategori  ")
	// fmt.Print("       Harga        ")
	// fmt.Print("  stok  ")
	// fmt.Println()

	fmt.Printf("%-11s %-35s %-16s %-15s %s \n", "ID", "Nama Barang", "Kategori", "Harga", "Stok")
}

func view_urutan() {
	var opsi int
	fmt.Println("----------------------------------------")
	fmt.Println("Pilih opsi urutan data")
	fmt.Println("1. ID Ascending")
	fmt.Println("2. ID Decending")
	fmt.Println("3. Nama produk ascending")
	fmt.Println("4. Nama produk decending")
	fmt.Println("5. Kategori ascending")
	fmt.Println("6. Kategori decending")
	fmt.Println("0. Kembali")
	fmt.Println("----------------------------------------")
	fmt.Scan(&opsi)
	if opsi == 1 {
		urut_id_naik()
	} else if opsi == 2 {
		urut_id_turun()
	} else if opsi == 3 {
		urut_produk_naik()
	} else if opsi == 4 {
		urut_produk_turun()
	} else if opsi == 5 {
		urut_kategori_naik()
	} else if opsi == 6 {
		urut_kategori_turun()
	}
}
