package p2p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Get IP table Data

type IpAddresses struct {
	IpAddress []IpAddress `json:"ip_address"`
}

type IpAddress struct {
	Ipv4   string `json:"ipv4"`
	Latency  time.Duration `json:"latency"`
	Download float64    `json:"download"`
	Upload float64 `json:"upload"`
}

// Read data from Ip tables from json file
func ReadIpTable()(*IpAddresses ,error){

	jsonFile, err := os.Open("/etc/p2p-rendering/ip_table.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil,err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var ipAddresses IpAddresses

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &ipAddresses)

    return &ipAddresses, nil
}

// Write to IP table json file
func (i *IpAddresses) WriteIpTable() error {
	file, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("/etc/p2p-rendering/ip_table.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Print Ip table data for Cli
func PrintIpTable() error {
	table, err := ReadIpTable()

	if err != nil {
		return err
	}

	for i := 0; i < len(table.IpAddress); i++ {
		fmt.Printf("----------------------\nIP Address: %s\nLatency: %s\nDownload: %f\nUplaod: %f\n-----------" +
			"-----------\n",table.IpAddress[i].Ipv4,
			table.IpAddress[i].Latency,table.IpAddress[i].Download,table.IpAddress[i].Upload)
	}
	return nil
}