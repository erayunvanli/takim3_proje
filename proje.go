package main

import (
	"fmt"
	"math/rand"
)

type food struct {
	name     string
	foodType int
	price    int
}
type menu struct {
	name  string
	corba string
	yemek string
	tatlı string
	price int
}

func main() {

	//08:00 - 20:00
	//random 8-19 10-19
	//09:00 -19
	//17 16
	//Çorba 1 Yemek 2 Tatlı 3
	var foods map[int]food = map[int]food{
		0: {"Tirsik", 2, 50},
		1: {"Kuru fasulye", 2, 30},
		2: {"pilav", 2, 25},
		3: {"Lokum Et", 2, 75},
		4: {"Steak Burger", 2, 75},
		5: {"Mercimek Çorbası", 1, 25},
		6: {"Ezogelin Çorbası", 1, 25},
		7: {"Adana", 2, 100},
		8: {"Magnolya", 3, 40},
		9: {"Sütlaç", 3, 25},
	}
	var menuler map[int]menu = map[int]menu{
		0: {"Sultan Menü", foods[5].name, foods[3].name, foods[8].name, 300},
		1: {"Esnaf Menü", foods[6].name, foods[1].name, foods[9].name, 100},
		2: {"Özel Menü", foods[5].name, foods[0].name, foods[8].name, 300},
	}
	saat := 8
	hesap := 0
	for {
		musteri := rand.Intn(10)
		arttir := 1
	x:

		if saat < 19 {
			fmt.Println("Yemekleri Görmek İçin 1")
			fmt.Println("Menüleri Görmek İçin 2")
		} else {
			fmt.Println("Yeni Yemek Eklemek İçin 3")
			fmt.Println("Yeni Menü Eklemek İçin 4")
			fmt.Println("Yemek Düzenlemek İçin 5")
			fmt.Println("Menü Düzenlemek İçin 6")
			fmt.Println("Günü kapatmak için 7")
		}

		fmt.Print("Seçiniz : ")
		foodOrMenu := 0
		//go format
		fmt.Scanf("%d", &foodOrMenu)
		if foodOrMenu == 99 {
			break
		}
		if saat < 19 {

			if foodOrMenu == 1 || foodOrMenu == 2 {
				for arttir <= musteri {
					fmt.Println(arttir, ". Müsteri")

					if foodOrMenu == 1 {
						fmt.Println(foods)
						fmt.Println("Sipariş Girmek İçin  1, Ana ekrana dönmek için 0")
						deger := 0
						yemekSecimi := 0
						fmt.Scanf("%d", &deger)
						if deger == 1 {
							fmt.Println(foods)
							fmt.Print("Seçim Yap")
							fmt.Scanf("%d", &yemekSecimi)
							hesap += foods[yemekSecimi].price
							fmt.Println(hesap)
							saat++
						} else if deger == 0 {

							goto x

						} else {
							fmt.Println("Yanlış Seçim")
							arttir--

						}
					} else if foodOrMenu == 2 {
						deger := 0
						yemekSecimi := 0
						fmt.Println(menuler)
						fmt.Println("Sipariş girmek için 1")
						fmt.Println("Ana menüye dönmek için 1 hariç başka bir tuşa basın")
						fmt.Scanf("%d", &deger)
						if deger == 1 {
							fmt.Println(menuler)
							fmt.Print("secim yap")
							fmt.Scanf("%d", &yemekSecimi)
							hesap += menuler[yemekSecimi].price
							fmt.Println(hesap)
							saat++
						} else {
							goto x
						}

					}
					fmt.Println("Saat: ", saat)
					arttir++
				}
			} else {
				fmt.Println("Yanlış Seçim")
			}

		} else {

			if foodOrMenu == 3 {
				foodsToplamAdet := len(foods)
				thisFood := foods[foodsToplamAdet]
				fmt.Printf("Yemek Adı : ")
				fmt.Scanf("%s", &thisFood.name)
				fmt.Printf("Çorba için 1 Yemek için 2 Tatlı için 3 : ")
				fmt.Scanf("%d", &thisFood.foodType)
				fmt.Printf("Fiyat : ")
				fmt.Scanf("%d", &thisFood.price)
				foods[foodsToplamAdet] = thisFood
			} else if foodOrMenu == 4 {
				menuToplamAdet := len(menuler)
				thisMenu := menuler[menuToplamAdet]
				fmt.Printf("Menü Adı : ")
				fmt.Scanf("%s", &thisMenu.name)
				for i := 0; i < len(foods); {
					if foods[i].foodType == 1 {
						fmt.Println(foods[i].name, "için", i, "basımız")
					}
					i++
				}
				fmt.Printf("Çorba : ")
				test := 0
				fmt.Scanf("%d", &test)
				if foods[test].foodType == 1 {
					for i := 0; i < len(foods); {
						if foods[i].foodType == 2 {
							fmt.Println(foods[i].name, "için", i, "basımız")
						}
						i++
					}
					thisMenu.corba = foods[test].name
					fmt.Printf("Yemek : ")
					test2 := 0
					fmt.Scanf("%d", &test2)
					if foods[test2].foodType == 2 {
						for i := 0; i < len(foods); {
							if foods[i].foodType == 3 {
								fmt.Println(foods[i].name, "için", i, "basımız")
							}
							i++
						}
						thisMenu.yemek = foods[test2].name
						fmt.Printf("Tatlı : ")
						test3 := 0
						fmt.Scanf("%d", &test3)
						if foods[test3].foodType == 3 {
							thisMenu.tatlı = foods[test3].name
							fmt.Printf("Fiyat : ")
							fmt.Scanf("%d", &thisMenu.price)
							menuler[menuToplamAdet] = thisMenu
						}
					}
				}
				fmt.Println(menuler)
			} else if foodOrMenu == 5 {
				duzenlecekVeri := 0
				fmt.Println(foods)
				fmt.Println("Düzenlemek İstediğiniz Veri: ")
				fmt.Scanf("%d", &duzenlecekVeri)
				thisFood := foods[duzenlecekVeri]
				fmt.Println("Yemek Adı : ", foods[duzenlecekVeri].name)
				fmt.Printf("Yeni Veri :")
				fmt.Scanf("%s", &thisFood.name)
				fmt.Println("Türü  : ", foods[duzenlecekVeri].foodType)
				fmt.Printf("Çorba için 1 Yemek için 2 Tatlı için 3 : ")
				fmt.Print("Yeni Veri :")
				fmt.Scanf("%d", &thisFood.foodType)
				fmt.Println("Fiyat : ", foods[duzenlecekVeri].price)
				fmt.Print("Yeni Fiyat :")
				fmt.Scanf("%d", &thisFood.price)
				foods[duzenlecekVeri] = thisFood
				fmt.Println(foods)
			} else if foodOrMenu == 6 {
				duzenlecekVeri := 0
				fmt.Println(menuler)
				fmt.Print("Düzenlemek İstediğiniz Veri: ")
				fmt.Scanf("%d", &duzenlecekVeri)
				thisFood := menuler[duzenlecekVeri]
				fmt.Println("Menü Adı : ", menuler[duzenlecekVeri].name)
				fmt.Printf("Yeni Veri :")
				fmt.Scanf("%s", &thisFood.name)
				fmt.Println("Çorba  : ", menuler[duzenlecekVeri].corba)
				for i := 0; i < len(foods); {
					if foods[i].foodType == 1 {
						fmt.Println(foods[i].name, "için", i, "basımız")
					}
					i++
				}
				fmt.Printf("Çorba Yeni Veri: ")
				corba := 0
				fmt.Scanf("%d", &corba)
				thisFood.corba = foods[corba].name

				fmt.Println("Yemek : ", menuler[duzenlecekVeri].yemek)
				for i := 0; i < len(foods); i++ {
					if foods[i].foodType == 2 {
						fmt.Println(foods[i].name, "için", i, "basımız")
					}

				}
				fmt.Printf("Yemek Yeni Veri: ")
				test2 := 0
				fmt.Scanf("%d", &test2)
				thisFood.yemek = foods[test2].name

				fmt.Println("Tatlı : ", menuler[duzenlecekVeri].tatlı)
				for i := 0; i < len(foods); i++ {
					if foods[i].foodType == 3 {
						fmt.Println(foods[i].name, "için", i, "basımız")
					}
				}
				fmt.Print("Tatlı Yeni Veri: ")
				test3 := 0
				fmt.Scanf("%d", &test3)
				thisFood.tatlı = foods[test3].name

				fmt.Println("Fiyat : ", menuler[duzenlecekVeri].price)
				fmt.Print("Yeni Fiyat : ")
				fmt.Scanf("%d", &thisFood.price)
				menuler[duzenlecekVeri] = thisFood
				fmt.Println(menuler)
			} else if foodOrMenu == 7 {
				fmt.Println(hesap)
				saat = 8
			}
		}
	}

}
