package main

import "fmt"

const NMAX int = 1000000
type data struct {
	index int
	nama string
	kategori string
	harga int
	stok int
}

type logData struct {
	index, jumlah, total int
	nama, kategory string
}
var log [NMAX]logData

var gudang [NMAX]data
func data_gudang(){

	gudang[0].index = 0000
	gudang[0].nama = "ABC JUICE GUAVA 1LTR"
	gudang[0].kategori = "minuman"
	gudang[0].harga = 15000
	gudang[0].stok = 64

	gudang[1].index = 0001
	gudang[1].nama = "ABC JUICE JAMBU 250ML"
	gudang[1].kategori = "minuman"
	gudang[1].harga = 4500
	gudang[1].stok = 64

	gudang[2].index = 0002
	gudang[2].nama = "ABC KACANG HIJAU 250ML"
	gudang[2].kategori = "minuman"
	gudang[2].harga = 4500
	gudang[2].stok = 64

	gudang[3].index = 0003
	gudang[3].nama = "AIR MINERAL ADES 1500ML"
	gudang[3].kategori = "minuman"
	gudang[3].harga = 6000
	gudang[3].stok = 64

	gudang[4].index = 0004
	gudang[4].nama = "AIR MINERAL ADES 600ML"
	gudang[4].kategori = "minuman"
	gudang[4].harga = 3500
	gudang[4].stok = 64

	gudang[5].index = 5
	gudang[5].nama = "AISO ISOTONIC 200ML"
	gudang[5].kategori = "minuman"
	gudang[5].harga = 3000
	gudang[5].stok = 64

	gudang[6].index = 6
	gudang[6].nama = "AQUA AIR MINERAL 1500ML"
	gudang[6].kategori = "minuman"
	gudang[6].harga = 7000
	gudang[6].stok = 64

	gudang[7].index = 7
	gudang[7].nama = "AQUA AIR MINERAL 600ML"
	gudang[7].kategori = "minuman"
	gudang[7].harga = 4000
	gudang[7].stok = 64

	gudang[8].index = 8
	gudang[8].nama = "BEAR BRAND 140ML"
	gudang[8].kategori = "minuman"
	gudang[8].harga = 8000
	gudang[8].stok = 64

	gudang[9].index = 9
	gudang[9].nama = "BENG2 SHARE PACK 5SX23G"
	gudang[9].kategori = "makanan"
	gudang[9].harga = 13000
	gudang[9].stok = 64

	gudang[10].index = 10
	gudang[10].nama = "BUAVITA 200ML JAMBU"
	gudang[10].kategori = "minuman"
	gudang[10].harga = 8000
	gudang[10].stok = 64

	gudang[11].index = 11
	gudang[11].nama = "BUAVITA 200ML JUS APEL"
	gudang[11].kategori = "minuman"
	gudang[11].harga = 8000
	gudang[11].stok = 64

	gudang[12].index = 12
	gudang[12].nama = "BUAVITA 200ML JUS JERUK"
	gudang[12].kategori = "minuman"
	gudang[12].harga = 8000
	gudang[12].stok = 64

	gudang[13].index = 13
	gudang[13].nama = "BUAVITA 200ML JUS APEL"
	gudang[13].kategori = "minuman"
	gudang[13].harga = 8000
	gudang[13].stok = 64

	gudang[14].index = 14
	gudang[14].nama = "BUAVITA 200ML LYCHEE"
	gudang[14].kategori = "minuman"
	gudang[14].harga = 8000
	gudang[14].stok = 64

	gudang[15].index = 15
	gudang[15].nama = "BUAVITA 200ML MANGO"
	gudang[15].kategori = "minuman"
	gudang[15].harga = 8000
	gudang[15].stok = 64

	gudang[16].index = 16
	gudang[16].nama = "CADBURY 200G DAIRY MILK"
	gudang[16].kategori = "makanan"
	gudang[16].harga = 20000
	gudang[16].stok = 64

	gudang[17].index = 17
	gudang[17].nama = "CHOKI-CHOKI COKLAT PASTA 368G"
	gudang[17].kategori = "makanan"
	gudang[17].harga = 21000
	gudang[17].stok = 64

}

func show_database(){
	var kosong data
	for i:=0; gudang[i] != kosong; i++ {
		fmt.Printf("%-11d %-35s %-16s %-15d %d \n",gudang[i].index, gudang[i].nama, gudang[i].kategori, gudang[i].harga, gudang[i].stok)
	}
	fmt.Println()
}