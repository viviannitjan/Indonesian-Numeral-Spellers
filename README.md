# Indonesian Numeral Spellers

## Latar Belakang
Mengeja angka merupakan salah satu kegiatan dasar yang dilakukan setiap harinya. Contoh kegiatan tersebut yaitu mengeja harga barang, nilai data, tanggal dan tahun, serta masih banyak lagi. Meskipun terkesan hal sepele, berdasarkan penelitian dari para dokter di Indonesia, seorang anak baru bisa membaca dan mengeja angka pada umur 4-6 tahun. Rentang usia tersebut tentunya terasa kurang cepat. Padahal, semakin cepat seorang bisa membaca dan mengeja angka, maka semakin cepat pula anak tersebut dapat belajar berhitung dan mempelajari hal-hal lainnya, bahkan termasuk belajar pemrograman.

Dari permasalah di atas, maka diperlukanlah suatu sarana pembelajaran yang dapat membantu anak-anak balita di Indonesia untuk membaca dan mengeja angka. Dengan adanya solusi tersebut, diharapkan anak-anak dapat membaca dan mengeja angka lebih cepat sehingga mampu segera mempelajari hal-hal lebih besar lainnya dan tentunya meningkatkan tingkat pendidikan di Indonesia.

## Stack yang digunakan
1. Go dengan library mux
2. ReactJS dengan library axios

## Langkah penggunaan
1. Pastikan library mux dan axios sudah terinstall
2. Clone repository ini kedalam GOPATH yang sudah dipastikan dapat berjalan dengan baik
3. Jalankan API dengan cara
```
go run "api.go" "speller.go"
```
4. Untuk menjalankan webapp masuk ke dir speller-app kemudian
```
npm start
```

## Contoh Kasus Uji
### Contoh Kasus Uji 1 : Pengejaan
Request :
```
GET '/spell?number=123456'
```
Response :
```JSON
STATUS CODE 200
{
    "status" : "OK",
    "text" : "seratus dua puluh tiga ribu empat ratus lima puluh enam"
}
```
### Contoh Kasus Uji 2 : Pembacaan
Request:
```JSON
POST '/read'
{
    "text" : "seribu sembilan ratus sembilan puluh tujuh"
}
```
Response :
```JSON
STATUS CODE 200
{
    "status" : "OK",
    "number" : 1997
}
```