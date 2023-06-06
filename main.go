package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func Minimaxsum(){
	var a, b, c, d, e int
	fmt.Print("Masukan angka ke 1 : ")
	fmt.Scan(&a)
	fmt.Print("Masukan angka ke 2 : ")
	fmt.Scan(&b)
	fmt.Print("Masukan angka ke 3 : ")
	fmt.Scan(&c)
	fmt.Print("Masukan angka ke 4 : ")
	fmt.Scan(&d)
	fmt.Print("Masukan angka ke 5 : ")
	fmt.Scan(&e)

	var array = []int{}
	array1 := append(array, a, b, c, d, e)
	fmt.Println(array1)

	//mengurutkan bilangan dari yang paling kecil
	sort.Sort(sort.IntSlice(array1))

	fmt.Println(array1)
	//menghitung jumalah dari bilangan terkecil dalam array
	minimalsum := array1[0:4]
	var summinimal int = 0
	for _, i := range minimalsum {
		summinimal += i
	}
	//menghitung bilangan terbesar dalam array
	maximal_array := array1[1:5]
	var maximalsum int = 0
	for _, i := range maximal_array {
		maximalsum += i
	}
	//Cetak hasil 
	fmt.Println(summinimal, "  ", maximalsum)
	Ulang()
	

}

//Mencari rasio dengan panjang array sebanyak N

func Ratios() {
	//masukan Input untuk menjumlahkan Jumnlah array
	var n int
	fmt.Print("Masukkan jumlah variabel input (n): ")
	fmt.Scanln(&n)

	inputs := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan input ke-%d: ", i+1)
		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())
		inputs[i] = input
	}

	fmt.Println("Variabel input:")
	fmt.Println(inputs)
	var angkanegatif float64 = 0
	var angkapositif float64 = 0
	var angka_nol float64 = 0
	for _, num := range inputs {
		if num < 0 {
			angkanegatif++
		}
		if num > 0 {
			angkapositif++
		}
		if num == 0 {
			angka_nol++
		}

	}
	//rumus perbandingan untuk nilai angka nol
	var nilai_sekarang_nol float64 = angka_nol / float64(n)
	var nilai_sekarang_positif float64 = angkapositif / float64(n)
	var nilai_sekarang_negatif float64 = angkanegatif / float64(n)
	fmt.Println("Berikut merupakan Perbandingannya")
	fmt.Printf("Perbandingan dari Bilangan nol Sebesar "+"%.6f\n", nilai_sekarang_nol)
	fmt.Printf("Perbandingan dari Bilangan Positif Sebesar "+"%.6f\n", nilai_sekarang_positif)
	fmt.Printf("Perbandingan dari Bilangan Negatif Sebesar"+"%.6f\n", nilai_sekarang_negatif)
	Ulang()

}


func ConversiWaktu() (string, error) {
	var input string
	fmt.Print("Masukan waktu yang akan di konversi : ")
	fmt.Scan(&input)

	format := "03:04:05PM"

	t, err := time.Parse(format, input)
	if err != nil {
		return "Gagal Mengkonversi", err
	}
	var waktu = t.Format("15:04:05")
	return waktu, err
	
}

func Golang(pilihan int){

	if pilihan == 1 {
		Minimaxsum()
	}
	if pilihan == 2 {
		Ratios()
	}
	if pilihan == 3 {
		waktu,_ := ConversiWaktu()
		fmt.Println(waktu)
		fmt.Println("terima kasih telah mencoba program ini :)")
	}
	if pilihan > 3 {
		fmt.Println("maaf Pilihan  tidak tersedia")
	}
	

	
}

func Ulang() {
	var n int
	fmt.Print("Masukkan Pilihan Program yang anda inginkan (n): ")
	fmt.Scanln(&n)
	Golang(n)
}

func main() {
	Ulang()

}
