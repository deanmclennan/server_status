package main

import (
	"fmt"
	"os"
	"time"

	"os/exec"
)

// cleasr the screen
func clear_screen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Bar back and forth
func bar() {
	for {
		for _, r := range `/-\` {
			fmt.Printf("\r%c", r)
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// Pong back and forth
func pong() {
	for {
		for _, r := range `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏` {
			fmt.Printf("\r%c", r)
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// Green text for output
func green(text string) string {
	return "\033[32m" + text + "\033[39m"
}

// Red text for output
func red(text string) string {
	return "\033[31m" + text + "\033[39m"
}

// Green tick with brackets for output
func green_tick() string {
	return green("[✓]")
}

// Red cross with brackets for output
func red_cross() string {
	return red("[✗]")
}

// Ping using ICMP protocol (ping) to check if a host is alive or not

func ping_ip(ip string) bool {
	_, err := exec.Command("ping", "-c", "1", ip).Output()

	if err != nil {
		return false
	}
	return true

}

// Get system time and date ubuntu
func get_time() string {
	cmd := exec.Command("date")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return " "
}

func server_status() {
	ip := "192.168.20.1"
	// space the output

	go pong()
	clear_screen()
	fmt.Println(green("SERVER STATUS CHECKER"))

	println("V1.0 - 2023")
	println("Built by Dean McLennan - MIT License")
	println(" ")
	println("Exit by pressing CTRL + C")

	println("IP Address: " + ip)
	fmt.Println("Checking if server is reachable... ")
	print(" ")

	if ping_ip(ip) == true {
		fmt.Println(" ")
		fmt.Println(green_tick() + " Server is reachable")
		fmt.Println(get_time())
	} else {
		fmt.Println(" ")
		fmt.Println(red_cross() + " Server is not reachable")
		fmt.Println(get_time())

	}

}

func main() {
	server_status()
}
