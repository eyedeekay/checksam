package checksam

import (
	"log"
	"time"

	"github.com/eyedeekay/sam3"
)

// CheckSAMAvailable tries a SAM connection and returns true if it succeeds.
func CheckSAMAvailable(yoursam string) bool {
	if yoursam == "" {
		yoursam = "127.0.0.1:7656"
	}
	sam, err := sam3.NewSAM(yoursam)
	if err != nil {
		return false
	}
	defer sam.Close()
	if _, err := sam.NewKeys(); err != nil {
		return false
	}
	return true
}

func WaitForSAM(yoursam string, timeout int) bool {
	loops := 0
	for !CheckSAMAvailable(yoursam) {
		// if loops is divisible by 5 output log
		if loops%5 == 0 {
		    log.Println("Waiting for SAM")
		}
		if loops == timeout {
			return false
		}
		time.Sleep(1 * time.Second)
		loops++
	}
	log.Println("Found SAM")
	return true
}