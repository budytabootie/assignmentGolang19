## Penjelasan Singkat
Program ini berfungsi untuk membalik urutan karakter dalam sebuah string. Program menerima string sebagai input dan mengembalikan versi terbaliknya sebagai output.

## Contoh Penggunaan
Jika kita menggunakan fungsi `reverseString` dengan input `"hello"`, maka hasil yang dikembalikan adalah `"olleh"`.

### Contoh Kode
```go
package main

import (
    "fmt"
)

func reverseString(s string) string {
    reversed := ""
    for _, char := range s {
        reversed = string(char) + reversed
    }
    return reversed
}

func main() {
    input := "hello"
    fmt.Println("Original:", input)
    fmt.Println("Reversed:", reverseString(input))
}
```

## Output yang dihasilkan:
Original: hello
Reversed: olleh

## Cara Menjalankan Program
Pastikan Go telah terpasang di komputer Anda.
Simpan kode ke dalam file bernama main.go.
Buka terminal, arahkan ke direktori tempat file main.go disimpan.

## Jalankan perintah berikut untuk menjalankan program:
```
go run main.go
```

Program akan mencetak string dalam urutan terbalik.


