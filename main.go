package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yudapc/go-rupiah"
)

/**
Golongan :: Upah Tetap :: Tunjangan Lembur Per Jam
1			6.000.000		20.000
2			5.400.000		18.000
3			5.000.000		15.000
4			4.800.000		12.000
5			4.000.000		10.000
**/

func main() {
	var (
		nama string
		golongan,  jamLembur int
		upahKotor, upahLembur, potongan, upahBersih float64
		scanner = bufio.NewScanner(os.Stdin)
	)

	fmt.Println("PROGRAM GAJI DAN TUNJANGAN")
	fmt.Print("Nama Karyawan : ")
	scanner.Scan()
	nama = scanner.Text()
	fmt.Print("Golongan Karyawan : ")
	fmt.Scanf("%d", &golongan)
	fmt.Print("Jam Lembur : ")
	fmt.Scanf("%d", &jamLembur)

	// Perhitungan upah 
	switch golongan {
	case 1 :
		upahLembur = float64(20000*jamLembur)
		upahKotor = float64(6000000 + upahLembur)
		break
	case 2 :
		upahLembur = float64(18000*jamLembur)
		upahKotor = float64(5400000 + upahLembur)
		break
	case 3 :
		upahLembur = float64(15000*jamLembur)
		upahKotor = float64(5000000 + upahLembur)
		break
	case 4 :
		upahLembur = float64(12000*jamLembur)
		upahKotor = float64(4800000 + upahLembur)
		break
	case 5 :
		upahLembur = float64(10000*jamLembur)
		upahKotor = float64(4000000 + upahLembur)
		break
	default :
		fmt.Println("Golongan Salah!")
	} 
		
	// Potohan
	if(upahKotor > 5000000) {
		potongan = 0.00025*upahKotor
	}else if(upahKotor >= 3000000) {
		potongan = 0.00001*upahKotor
	}else{
		potongan = 0
	}

	upahBersih = upahKotor - potongan

	fmt.Println("_________________________________________________")
	fmt.Println("STRUK GAJI KARYAWAN")
	fmt.Println("_________________________________________________")
	fmt.Println("Nama Karyawan \t:", nama)
	fmt.Println("Golongan \t:", golongan)
	fmt.Println("Jam Lembur \t:", jamLembur, "jam")

	fmt.Println("Upah Lembur \t:", rupiah.FormatRupiah(upahLembur))
	fmt.Println("Upah Kerja \t:", rupiah.FormatRupiah(upahKotor))
	fmt.Println("Potongan \t:", rupiah.FormatRupiah(potongan))
	fmt.Println("Upah Bersih \t:", rupiah.FormatRupiah(upahBersih))

}