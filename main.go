package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/go-pdf/fpdf"
	_ "golang.org/x/image/webp"
)

func main() {
	if len(os.Args) < 2 {
		obsluzBlad(nil, "Nie podano argumentu! (Nie załączono pliku)")
	}
	sciezkaZdjecia := os.Args[1]

	plik, err := os.Open(sciezkaZdjecia)
	if err != nil {
		obsluzBlad(err, "Nie udało się otworzyć pliku!")
	}
	defer plik.Close()

	obrazConfig, format, err := image.DecodeConfig(plik)
	if err != nil {
		obsluzBlad(err, "Nie udało się odczytać pliku! (Błędny format - program akceptuje jedynie .png i .jpg)")
	}

	var orientacja string
	if obrazConfig.Height > obrazConfig.Width {
		orientacja = "P"
	} else {
		orientacja = "L"
	}

	dokument := fpdf.New(orientacja, "mm", "A4", "")

	dokument.AddPage()

	szerokosc, wysokosc := dokument.GetPageSize()
	opcje := fpdf.ImageOptions{
		ImageType: format,
	}
	dokument.ImageOptions(sciezkaZdjecia, 0, 0, szerokosc, wysokosc, false, opcje, 0, "")

	nazwaPdf := fmt.Sprintf("%v.pdf", strings.TrimSuffix(sciezkaZdjecia, filepath.Ext(sciezkaZdjecia)))
	err = dokument.OutputFileAndClose(nazwaPdf)
	if err != nil {
		obsluzBlad(err, "Nie udało się zapisać pliku pdf")
	}
}

func obsluzBlad(err error, tresc string) {
	fmt.Println("===== BŁĄD =====")
	fmt.Println(tresc)
	if err != nil {
		fmt.Printf("Informacja dla technika: %s\n", err)
	}
	fmt.Println("================\n\n")
	fmt.Println("\nNaciśnij enter aby zakończyć...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	os.Exit(1)
}
