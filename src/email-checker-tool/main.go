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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("domain, hasMX, hasSPF, sprRecord, ")
	fmt.Println("")

	for scanner.Scan() {
		fmt.Printf("Domain to check :%s", scanner.Text())
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Fatal Error : could not read from input:%v", err)
	}
}

func checkDomain(domain string) {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Return DNS txt records for given domain name
	txtRecords, err := net.LookupTXT(domain)

	fmt.Printf("Records : %v", txtRecords)
	fmt.Println(" ---------------")
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	if len(txtRecords) > 0 {
		hasMX = true
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	txtRecords2, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("Eror %v", err)
	}

	for _, record := range txtRecords2 {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v , %v , %v , %v , %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
