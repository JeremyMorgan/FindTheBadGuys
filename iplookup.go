package main

import (
	"bufio"
	"bytes"
	"github.com/oschwald/geoip2-golang"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {

	// your bad guys file, list of all IP addresses of attackers
	// You can get this from centos with this script : https://gist.github.com/JeremyMorgan/94af88899785ea725a55a382f3fd209b
	var badguysfile = "[Your path to IP addresses]"

	// location of your geoliteDB file
	var geolitedb = "[Your path to GeoLite2-City.mmdb]"

	// files you want to store your data in:
	var contintentsFile = "continents.txt"
	var countriesFile = "countries.txt"
	var citiesFile = "cities.txt"
	var subdivisionsFile = "subdivisions.txt"

	/////////////    No need to edit past this line  ///////////////

	// create some buffers for our data
	var continents bytes.Buffer
        var countries bytes.Buffer
	var cities bytes.Buffer
	var subdivisions bytes.Buffer

	// open the GeoLite2 database
	db, err := geoip2.Open(geolitedb)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// open up bad guys file (list of IP addresses from attackers
	file, err := os.Open(badguysfile)

	if err != nil {
		log.Fatalf("We couldn't open your bad guys file: %s", err)
	}

	// parse it line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var ipaddresses []string

	for scanner.Scan() {
		ipaddresses = append(ipaddresses, scanner.Text())
	}

	// close the file, we don't need it anymore
	file.Close()

	// loop through our IP addresses and make a lookup to our database for each
	for _, ipaddress := range ipaddresses {

		ip := net.ParseIP(ipaddress)
		record, err := db.City(ip)
		if err != nil {
			log.Fatal(err)
		}

		// write continent name into buffer
		if len(record.Continent.Names["en"]) > 0 {
			continents.WriteString(record.Continent.Names["en"] + "\n")
		}

		// write country name into buffer
		if len(record.Country.Names["en"]) > 0 {
			countries.WriteString(record.Country.Names["en"] + "\n")
		}

		// write city name into buffer
		if len(record.City.Names["en"]) > 0 {
			cities.WriteString(record.City.Names["en"] + "\n")
		}

		// write subdivision into buffer
		if len(record.Subdivisions) > 0 {
			subdivisions.WriteString(record.Subdivisions[0].Names["en"] + "\n")
		}
	}

	// write the continents file
	err = ioutil.WriteFile(contintentsFile, continents.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	// write the countries file
	err = ioutil.WriteFile(countriesFile, countries.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	// write the cities file
	err = ioutil.WriteFile(citiesFile, cities.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	// write the subdivisions file
	err = ioutil.WriteFile(subdivisionsFile, subdivisions.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
