package main

import (
	"fmt"
	"log"
	"loki/ocr"
)

func main() {
	// Configuração básica
	config := ocr.Config{
		Language:    "por", // pt brzada
		PageSegMode: 1,     // Segmentação automática
	}

	ocr.PreprocessImage("nheco.jpg")

	ocrEngine, err := ocr.New(config)
	if err != nil {
		log.Fatalf("Falha ao criar OCR: %v", err)
	}
	defer ocrEngine.Close()
	text, err := ocrEngine.ExtractText("grayscale2.0nheco.jpg")
	if err != nil {
		log.Fatalf("Falha ao extrair texto: %v", err)
	}

	fmt.Println("Texto reconhecido (normal):")
	fmt.Println(text)

}
