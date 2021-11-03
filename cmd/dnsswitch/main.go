package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var debug *bool

func init() {
	debug = flag.Bool("d", false, "debug flag")
	flag.Parse()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) != 2 {
			continue
		}
		ip := parts[1]
		parts = strings.SplitN(parts[0], ".", 2)
		if len(parts) != 2 {
			continue
		}
		rr := parts[0]
		domain := parts[1]

		err := EnsureDomainRecordExist(domain, rr, ip)
		if err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

type GetResponse struct {
	DomainRecords struct {
		Record []DomainRecord `json:"Record"`
	} `json:"DomainRecords"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	RequestId  string `json:"RequestId"`
	TotalCount int    `json:"TotalCount"`
}
type DomainRecord struct {
	DomainName string `json:"DomainName"`
	Line       string `json:"Line"`
	Locked     bool   `json:"Locked"`
	RR         string `json:"RR"`
	RecordId   string `json:"RecordId"`
	Status     string `json:"Status"`
	TTL        int    `json:"TTL"`
	Type       string `json:"Type"`
	Value      string `json:"Value"`
	Weight     int    `json:"Weight"`
}

var getCmdFmt = `aliyun alidns DescribeDomainRecords --DomainName %s --RRKeyWord %s`
var updateCmdFmt = `aliyun alidns UpdateDomainRecord --RecordId %s --RR %s --Type A --Value %s`
var createCmdFmt = `aliyun alidns AddDomainRecord --DomainName %s --RR %s --Type A --Value %s --Line default`

func runCmd(cmdStr string) ([]byte, error) {
	var stdout, stderr bytes.Buffer
	slice := strings.Split(cmdStr, " ")
	cmd := exec.Command(slice[0], slice[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	if stderr.String() != "" {
		return nil, errors.New(stderr.String())
	}

	if *debug {
		fmt.Println(string(stdout.Bytes()))
	}

	return stdout.Bytes(), nil
}

func EnsureDomainRecordExist(domain, rr, ip string) error {
	record, err := GetDomainRecord(domain, rr)
	if err != nil {
		return err
	}
	if record == nil {
		// new one
		return CreateDomainRecord(domain, rr, ip)
	}
	// keep if no change
	if record.Value == ip {
		fmt.Printf("keeping: %s.%s => %s\n", rr, domain, ip)
		return nil
	}
	// update existing
	return UpdateDomainRecord(record.RecordId, record.RR, ip, domain)
}

func GetDomainRecord(domain, rr string) (*DomainRecord, error) {
	getCmd := fmt.Sprintf(getCmdFmt, domain, rr)
	output, err := runCmd(getCmd)
	if err != nil {
		return nil, err
	}

	var resp GetResponse
	err = json.Unmarshal(output, &resp)
	if err != nil {
		return nil, err
	}

	for _, record := range resp.DomainRecords.Record {
		// need exactly match rr
		if record.RR == rr {
			return &record, nil
		}
	}
	return nil, nil
}

func UpdateDomainRecord(recordId, rr, ip, domain string) error {
	fmt.Printf("updating %s: %s.%s => %s\n", recordId, rr, domain, ip)
	updateCmd := fmt.Sprintf(updateCmdFmt, recordId, rr, ip)
	_, err := runCmd(updateCmd)
	return err
}

func CreateDomainRecord(domain, rr, ip string) error {
	fmt.Printf("creating: %s.%s => %s\n", rr, domain, ip)
	createCmd := fmt.Sprintf(createCmdFmt, domain, rr, ip)
	_, err := runCmd(createCmd)
	return err
}
