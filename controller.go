package main

import "fmt"

func belanja(){
	var nama_barang string
	view_belanja()
	fmt.Scan(&nama_barang)
	header_tabel_database()
	cari_barang(nama_barang)
}

func cari_barang(barang string) {
	
}