package main

import(
	"fmt"
	"os"
)

func index(){
	var pilih int

	fmt.Println("----------------------------------------")
	fmt.Println("                 DR.Mart                ")
	fmt.Println("")
	fmt.Println("Pilih:")
	fmt.Println("1. Belanja")
	fmt.Println("2. List Data Barang")
	fmt.Println("3. Log")
	fmt.Println("4. Exit")
	fmt.Println("")
	fmt.Println("----------------------------------------")
	fmt.Scanln(&pilih)
	if pilih == 1 {
		belanja()
	} else if pilih == 2 {
		list_barang()
	} else if pilih == 3 {
		log()
	} else {
		os.Exit(0)
	}
}

func view_belanja(){
	fmt.Println("----------------------------------------")
	fmt.Println("              View Belanja              ")
	fmt.Println("----------------------------------------")
	fmt.Print("nama Barang : ") // search barang dari data

	// nampilin hasil search barang
	
}

func list_barang(){
	fmt.Println("               List Barang              ") // nampilin database
	fmt.Println("----------------------------------------")
	header_tabel_database()
	show_database()
	fmt.Println("----------------------------------------")
}

func log(){
	fmt.Println("               Log Barang               ") // nampilin data log pembelanjaan
	fmt.Println("----------------------------------------")
	fmt.Printf("%-11s %-35s %-16s %-15s %s \n", "ID", "Nama Barang", "Kategori", "Harga", "Banyak")
	fmt.Println("----------------------------------------")
}

func header_tabel_database(){
	// fmt.Print(" NO ")
	// fmt.Print("        Nama Barang         ")
	// fmt.Print("  Kategori  ")
	// fmt.Print("       Harga        ")
	// fmt.Print("  stok  ")
	// fmt.Println()

	fmt.Printf("%-11s %-35s %-16s %-15s %s \n", "ID", "Nama Barang", "Kategori", "Harga", "Stok")
}