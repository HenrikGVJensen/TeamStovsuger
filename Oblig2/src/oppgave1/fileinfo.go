package main

import (
	"fmt"
	"os"
	"log"
	"flag"

)

func main() {
	// først  go build fileinfo.go

	Peker := flag.String("f", "", "filnavn")

	flag.Parse()
	// fileinfo -f [filnavn] ofr å finne fil

	info := *Peker

	file, err := os.Lstat(*Peker) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	x := float64(file.Size())
	kb := x/1024
	mb := kb/1024
	gb := mb/1024
	fmt.Printf("Information about file %s: \n", info)
	fmt.Printf("Size : %.0fB, %fKB, %fMB, %.9fGB \n", x, kb, mb, gb)


	mode :=file.Mode()
	dir := mode.IsDir()
	reg := mode.IsRegular()
	append := mode&os.ModeAppend !=0
	dev := mode%os.ModeDevice !=0
	sym := mode%os.ModeSymlink !=0
	if dir == true{
		fmt.Println("is directory")
	} else {
		fmt.Println("is not directory")
	}
	if reg == true{
		fmt.Println("is regular file")
	} else {
		fmt.Println("is not regular file")
	}
	fmt.Println("Has Unix permission bits: ",mode)
	if append == true{
		fmt.Println("is append only")
	} else {
		fmt.Println("is not append only")
	}
	if dev == true {
		fmt.Println("is device file")
	} else {
		fmt.Println("is not device file")
	}
	if sym == true{
		fmt.Println("is symbolic link")

	} else {
		fmt.Println("is not symbolic link")
	}


}
