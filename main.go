package main

import (
	"bufio"
	"evermos-test-cmd/handler"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=== Rizky Mochammad Soleh ===")
		fmt.Println("Ini adalah aplikasi simulasi untuk mengetest aplikasi yang telah saya buat")
		fmt.Println("Adapun langkah - langkah yang sudah saya siapkan disini adalah")
		fmt.Println("ketik '1' untuk membuat customer baru")
		fmt.Println("ketik '2' untuk membuat product baru")
		fmt.Println("ketik '3' untuk membuat keranjang belanja baru")
		fmt.Println("ketik '4' untuk checkout dari keranjang belanja baru")
		fmt.Println("ketik '5' untuk melakukan pembayaran")
		fmt.Println("ketik '6' untuk melihat inventory adjustment")
		fmt.Println("ketik '7' untuk menambah inventory adjustment product")
		fmt.Println("ketik '8' untuk mengurangi inventory adjustment product")
		fmt.Println("")
		inputText, _ := reader.ReadString('\n')
		// convert CRLF to LF
		inputText = strings.TrimSpace(strings.ToLower(inputText))
		if inputText == "1" {
			handler.CreateCustomer()
			time.Sleep(5 * time.Second)
		} else if inputText == "2" {
			handler.CreateProduct()
			time.Sleep(5 * time.Second)
		} else if inputText == "3" {
			handler.CreateCart()
			time.Sleep(5 * time.Second)
		} else if inputText == "4" {
			handler.CheckoutCreateCart()
			time.Sleep(5 * time.Second)
		} else if inputText == "5" {
			handler.CreatePayment()
			time.Sleep(5 * time.Second)
		} else if inputText == "6" {
			handler.FindProduct()
			time.Sleep(30 * time.Second)
		} else if inputText == "7" {
			handler.CreateInventoryAdjustment("in")
			time.Sleep(5 * time.Second)
		} else if inputText == "8" {
			handler.CreateInventoryAdjustment("out")
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Fungsi tidak ditemukan")
			time.Sleep(5 * time.Second)
		}


		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

