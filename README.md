# Lokanta İşletmeciliği

Lokantada sunuclacak en az 3 adet günlük menü olmalı ve Lokantanın en az 10 adet pişirebildiği yemek olmalıdır. Lokantaya günlük olarak müşteriler gelmeli Random olarak menü yada yemek seçmeli ve yemeklerini yemelidir. Ya bir adet menü seçmeli yada gerçeğe uygun şekilde yemek yanına içecek gibi seçimler yapmalı bunların hepsi Random olarak gerçekleşmeşli fakat gerçeğe uygun olmalıdır ( yemek kategorilendirlirse bu sorun çözülecektir). Her müşteri bir saat de yemeğini yemelidir bu sistem için 1 sn olabilir.
Başlangıçta sistem tarafında tanımlı en az 10 adet yemek bulunmalı ve kullanıcı tarfından yeni menüler konsoldan eklenebilmelidir. Her müşteriden sonra kasa durumu adım adım ekranda gösterilmeli her aşama için kullanıcıya bilgi verilmelidir. Gün sonunda günlük rapor çıkartılmalıdır.

## Yapılacaklar

- Var sayılan olarak en az 10 adet Yemek ve 3 adet menü sisteme yazılım aşamasında tanımlanacak
- Yazılım başladığında kullanıcıya her menü ve yemeğin tüm özellikleri tek tek rapor edilecek
- Kullanıcı isterse konsoldan (ftm.Scan...) gireceği komutlarla çalışma anında default menülerin bilgilerini düzenleyebilecektir.
- Kullanıcı isterse konsoldan kendi menülerini tanımlayabilecektir.
- Kullanıcı başlat komutu verdiğinde oluşan lokanta çalışmaya başlayacak ve mesai bitimine kadar müşterilerine yemek servisi yapacaktır.
- Her müşteriden sonra kasa durumu tek tek bilgilendirilecek . Her müşteri geldiğinde sipraişi alınacak ve ekrana yazılacak gün sonunda tüm rapor ekrana dökülecektir.


## Çalışma Prensibi

- Lokanta açıldığında Random olarak müşteriler gelicek her saate farklı sayıda Random müşteri gelecek.
- Gelen müşteriler Random olarak yemek seçecek ve yiyecekler
- Gün sonunda yapılan işlemlerin ve kasa durumunun sonucu rapor edilecektir.


## Notlar

- Yemek ve Menü tanımlamalarında Struct kullanılacak.
- Yemeklerin ve Menülerin özellerikleri tanımlanırken "methotlar" kullanılacak (foknsiyonlar değil).
- Yemekler ve Menüler in-memory map olarak hafızada tutlacak ( map kullanılacak ).
- Log akışları sırasında time.Sleep ile log akışı yavaşlatılırsa daha gerçekci olacaktır.
- Yazılım ilk başladığında nasıl çalıştığını kullanıcın neler yapabileceğini anlatan bir çıktıyı ekranda göstermelidir.
