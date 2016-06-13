package main   

import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/gob"
	"encoding/base64"
	"bytes"
)

// Define our struct to store server details
type ServerDetails struct {
	Host 	string // Uppercase property name denotes public
	User 		string
	password	string // lowercase property name denotes private
}

func main() {

	// Define config path and config dir
	configPath := GetConfigPath()
	configDir := GetFFDirectory()

	// If film-fetch dir doesn't exist then create it
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.Mkdir(configDir, 0777)
	}

	// If there's no defined server details, ask
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		
		// Initialise a blank struct
		details := ServerDetails{}
		
		// Prompt user for details
		collectDetails(&details)

		// Encode struct as string and save to config
		base64struct := ToGOB64(details)
		d1 := []byte(base64struct)
		ioutil.WriteFile(configPath, d1, 0644)
	} else {
		// Load details
	    dat, _ := ioutil.ReadFile(configPath)
	    details := FromGOB64(string(dat))
	    printDetails(details)
	}	
	
}

// Prompt the user for required details
func collectDetails(details *ServerDetails) {
	fmt.Println("film-fetch has not been configured, you will be asked for some details now.\n")
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Host: ")
	details.Host, _ = reader.ReadString('\n')
	fmt.Print("Username: ")
	details.User, _ = reader.ReadString('\n')

	//@TODO: Ask for password, we won't save it though - golangs encoding doesn't encode private struct properties
}

// Print details to screen, just for debug
func printDetails(details ServerDetails) {
	fmt.Printf("Host: %s", details.Host)
	fmt.Printf("User: %s", details.User)
}

// Calculate where the config should be
// @TODO: Make this multiplatform and actually work
func GetConfigPath() string {
	return "/Users/cabbage/.film-fetch/config"
}

// Calculate where the FF directory should be
// @TODO: Make this multiplatform and actually work
func GetFFDirectory() string {
	return "/Users/cabbage/.film-fetch"
}

// Taken from http://stackoverflow.com/questions/28020070/golang-serialize-and-deserialize-back
// @TODO: Wrap this in something and move it elsewhere
func ToGOB64(m ServerDetails) string {
    b := bytes.Buffer{}
    e := gob.NewEncoder(&b)
    err := e.Encode(m)
    if err != nil { fmt.Println(`failed gob Encode`, err) }
    return base64.StdEncoding.EncodeToString(b.Bytes())
}

func FromGOB64(str string) ServerDetails {
	m := ServerDetails{}
    by, err := base64.StdEncoding.DecodeString(str)
    if err != nil { fmt.Println(`failed base64 Decode`, err); }
    b := bytes.Buffer{}
    b.Write(by)
    d := gob.NewDecoder(&b)
    err = d.Decode(&m)
    if err != nil { fmt.Println(`failed gob Decode`, err); }
    return m
}

