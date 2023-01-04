package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the email: ")
	s.Scan()
	validateEmail(s.Text())
	// for scanner.Scan() {
	// 	validateEmail(scanner.Text())
	// }

	if err := s.Err(); err != nil {
		log.Fatalf("Error: could not read from input %v \n", err)
	}
}

func validateEmail(email string) {
	// var hasMX, hasSPF, hasDMARC bool
	// var SPFRecord, DMARCRecord string
	var isValid string
	domain := strings.Split(email, "@")[1]
	hasMX := checkMXRecord(domain)
	txtRecords := checkTXTRecords(domain)
	hasSPF, SPFRecord := checkSPFRecords(txtRecords)
	hasDMARC, DMARCRecord := checkDMARCRecord(txtRecords)

	if hasMX {
		isValid = "valid"
	} else {
		isValid = "invalid"
	}

	fmt.Printf("\nThe email \033[32m%s\033[0m is %s \nHas MX Record: %v \nHas SPF Record: %v \nHas DMARC Record: %v \nSPF Record: %v \nDMARC Record: %v \n", email, isValid, hasMX, hasSPF, hasDMARC, SPFRecord, DMARCRecord)
}

func checkDMARCRecord(txtRecords []string) (bool, string) {
	for _, d := range txtRecords {
		if strings.HasPrefix(d, "v=DMARC") {
			return true, d
		}
	}

	return false, ""
}

func checkTXTRecords(domain string) []string {
	tr, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Println("could not found TXT records")
	}

	return tr
}

func checkSPFRecords(txtRecords []string) (bool, string) {
	for _, r := range txtRecords {
		if strings.HasPrefix(r, "v=spf") {
			return true, r
		}
	}

	return false, ""
}

func checkMXRecord(domain string) bool {
	mx, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("could not found MX records")
	}

	return len(mx) > 0
}
