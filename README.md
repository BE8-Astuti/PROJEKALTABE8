# Projek_tim_gorm

### Deskripsi Projek
Untuk deskripsi projek bisa dibaca di link ini --> [Deskripsi Project](https://docs.google.com/document/d/1lNPeN9i2zWyYxSG4kHKVGN7nxxHksFQea1RWVm3y2vc/edit).


### Cara menjalankan projek di device Anda.
* Pastikan golang dan mySQL sudah terinstall di device anda.
* Clone repository ini dengan menjalankan `git clone github.com:BE8-Astuti/PROJEKALTABE8.git` di terminal. Kemudian jalan `cd projekrentbook`
* Hapus folder .git yang terdapat di repository cloning dengan menjalankan `rm -rf .git`.
* Buka repository cloning menggunakan code editor yang ada di device anda atau dengan menjalankan `code .` di terminal (jika menggunakan Visual Studio Code).
* Buat buah sebuah database di mySQL dengan menjalankan `CREATE DATABASE <nama database>`
* Buat sebuah file baru di dalam folder utama dengan nama .env.
* Isi file .env dengan `export CONNECTION_DB="root:pass@tcp(127.0.0.1:3306)/DBname?charset=utf8mb4&parseTime=True&loc=Local"` dengan `root` adalah nama username server SQL, `pass` adalah password ketika ingin membuka SQL, dan `DBname` nama database yang anda buat di mySQL.
* Buka terminal (untuk windows user disarankan menggunakan Git Bash) lagi anda dan jalankan kode kode berikut. `go get -u gorm.io/gorm` kemudian `go get -u gorm.io/driver/mysql` dan yang terakhir `source .env`.
* Untuk menikmati programm kami, jalan lupa untuk menjalankan `go run main.go`. SELAMAT MENIKMATI PROGAM!!!

