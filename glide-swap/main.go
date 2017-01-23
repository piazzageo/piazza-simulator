package main

import "log"

func main() {
	// read in glide.yaml
	// read in glide.lock

	_, err := readGlideLock()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
