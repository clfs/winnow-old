package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagInput   = flag.String("input", "", "read a message from FILE")
	flagOutput  = flag.String("output", "", "write a message to FILE")
	flagKey     = flag.String("key", "", "read a key from FILE")
	flagKeygen  = flag.String("keygen", "", "write a key to FILE")
	flagDeflate = flag.Bool("deflate", false, "deflate input")
)

const keyLength = 64

func main() {
	flag.Parse()

	if *flagKeygen != "" {
		err := generateKey(*flagKeygen)
		if err != nil {
			log.Fatal(err)
		}
	} else if *flagInput != "" && *flagOutput != "" && *flagKey != "" {
		if *flagDeflate {
			err := deflate(*flagInput, *flagOutput, *flagKey)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := inflate(*flagInput, *flagOutput, *flagKey)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		flag.Usage()
	}
}

func generateKey(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return fmt.Errorf("could not open keygen file %q: %v", path, err)
	}
	defer f.Close()
	key := make([]byte, keyLength)
	if _, err := rand.Read(key); err != nil {
		return fmt.Errorf("could not generate key: %v", err)
	}
	if _, err := f.Write(key); err != nil {
		return fmt.Errorf("could not write key to file %q: %v", path, err)
	}
	return nil
}

func deflate(inPath, outPath, keyPath string) error {
	return nil
}

func inflate(inPath, outPath, keyPath string) error {
	return nil
}

/*
	// Try to open all files.
	if path := *flagKeygen; path != "" {
		keygenFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			log.Fatalf("error: failed to open keygen file %q: %v", path, err)
		}
		defer .Close()
	}




	switch {
	case *flagKeygen != "":
		f, err := os.OpenFile(*flagKeygen, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			log.Fatalf("error: failed to open keygen file %q: %v", *flagKeygen, err)
		}
		defer f.Close()
		generateKey(f)
	case *flagInput != "" && *flagOutput != "" && *flagKey != "":

	}

	if  {

	}

		if  {
		switch {
		case *flagDeflate:
			deflate(*flag
		}
	}

	// Handle inflation and deflation.
	if inPath, outPath := *flagInput, *flag

	// Handle key generation.
	if path := *flagKeygen; path != "" {
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			log.Fatalf("error: failed to open keygen file %q: %v", path, err)
		}
		defer f.Close()
		if err := keygen(f); err != nil {
			log.Fatalf("error: %v", err)
		}
		return
	}

}

func handleKeyGeneration(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		log.Fatalf("error: failed to open keygen file %q: %v", path, err)
	}
	defer f.Close()
	if err := keygen(f); err != nil {
		log.Fatalf("error: %v", err)
	}
	return nil
}

func keygen(w io.Writer) error {
	key := make([]byte, keyLength)
	if _, err := rand.Read(key); err != nil {
		return fmt.Errorf("could not generate key: %v", err)
	}
	if _, err := w.Write(key); err != nil {
		return fmt.Errorf("could not write key to file: %v", err)
	}
	return nil
}
*/
