package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func Rastgale(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {

	var oyunMod int
	var input string
	var err error
	var TahminSayi int
	var tahmin int
	var hak int

	fmt.Println("Sayi tahmin Oyununa Hoş geldiniz.")
	fmt.Println("1 ile 100 arasında bir sayı tahmin ettim. Bakalım bulabilecek misin?")
	fmt.Println("Seçtiğin Zorluğa göre tahmin hakkın değişecektir.")
	fmt.Println("--------------------------------------------------------")
	for {
		fmt.Println("1-Kolay (10 tahmin hakkı)")
		fmt.Println("2-Orta (5 tahmin hakkı)")
		fmt.Println("3-Zor (3 tahmin hakkı)")
		fmt.Print("Oynamak istediğiniz zorluğu seçiniz:")
		fmt.Scanln(&input)
		oyunMod, err = strconv.Atoi(input)
		if err != nil || oyunMod <= 0 || oyunMod > 3 {
			fmt.Println("<----- Hatalı işlem yaptınız ----->")
			continue
		}
		break
	}

	if oyunMod == 1 {
		fmt.Println("Kolay Modu Seçtiniz.\n10 tahmin Hakkınız var.")
		fmt.Println("Aklımdaki sayiyi bil bakalım...")
		TahminSayi = Rastgale(1, 10)
		hak = 10
		for hak > 0 {
			fmt.Print("Tahmin:")
			_, err1 := fmt.Scanln(&tahmin)
			if err1 != nil || tahmin < 1 || tahmin > 100 {
				fmt.Println("Lütfen 1 ile 100 arasında bir sayı giriniz!!")
				continue
			}
			if tahmin < TahminSayi {
				fmt.Println("Sayıyı arttırın!")
			} else if tahmin > TahminSayi {
				fmt.Println("Sayıyı azalt!")
			} else {
				fmt.Printf("%d sayısı doğru tahmindi.\n Tebrikler Kazandınız.\n", tahmin)
				fmt.Printf("%d. tahmin hakkında bildin.", 11-hak)
				break
			}

			hak--
		}

		if hak == 0 {
			fmt.Printf("Hakkınız Bitti. Tahmin ettiğim sayı:%d", TahminSayi)
		}

	}

	if oyunMod == 2 {
		fmt.Println("Orta Modu Seçtiniz.\n10 tahmin Hakkınız var.")
		fmt.Println("Aklımdaki sayiyi bil bakalım...")
		TahminSayi = Rastgale(1, 10)
		hak = 5
		for hak > 0 {
			fmt.Print("Tahmin:")
			_, err1 := fmt.Scanln(&tahmin)
			if err1 != nil || tahmin < 1 || tahmin > 100 {
				fmt.Println("Lütfen 1 ile 100 arasında bir sayı giriniz!!")
				continue
			}
			if tahmin < TahminSayi {
				fmt.Println("Sayıyı arttırın!")
			} else if tahmin > TahminSayi {
				fmt.Println("Sayıyı azalt!")
			} else {
				fmt.Printf("%d sayısı doğru tahmindi.\n Tebrikler Kazandınız.\n", tahmin)
				fmt.Printf("%d. tahmin hakkında bildin.", 11-hak)
				break
			}

			hak--
		}

		if hak == 0 {
			fmt.Printf("Hakkınız Bitti. Tahmin ettiğim sayı:%d", TahminSayi)
		}

	}

	if oyunMod == 3 {
		fmt.Println("Zor Modu Seçtiniz.\n10 tahmin Hakkınız var.")
		fmt.Println("Aklımdaki sayiyi bil bakalım...")
		TahminSayi = Rastgale(1, 10)
		hak = 3
		for hak > 0 {
			fmt.Print("Tahmin:")
			_, err1 := fmt.Scanln(&tahmin)
			if err1 != nil || tahmin < 1 || tahmin > 100 {
				fmt.Println("Lütfen 1 ile 100 arasında bir sayı giriniz!!")
				continue
			}
			if tahmin < TahminSayi {
				fmt.Println("Sayıyı arttırın!")
			} else if tahmin > TahminSayi {
				fmt.Println("Sayıyı azalt!")
			} else {
				fmt.Printf("%d sayısı doğru tahmindi.\n Tebrikler Kazandınız.\n", tahmin)
				fmt.Printf("%d. tahmin hakkında bildin.", 11-hak)
				break
			}

			hak--
		}

		if hak == 0 {
			fmt.Printf("Hakkınız Bitti. Tahmin ettiğim sayı:%d", TahminSayi)
		}

	}

}
