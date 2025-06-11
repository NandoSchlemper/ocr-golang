package ocr

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func convertToGrayscale(img image.Image, outputPath string) error {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// Conversão adequada para escala de cinza usando a fórmula NTSC
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			original := img.At(x, y)
			r, g, b, _ := original.RGBA()

			// Fórmula de luminosidade corrigida
			gray := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
			grayImg.Set(x, y, color.Gray{Y: uint8(gray)})
		}
	}

	// Cria diretório se necessário
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Aumenta a qualidade do JPEG
	opt := jpeg.Options{Quality: 90}
	return jpeg.Encode(outFile, grayImg, &opt)
}

func ImageGrayscaleLoader(path string) {
	imgFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	img, err := jpeg.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	err = convertToGrayscale(img, "grayscale"+path)
	if err != nil {
		panic(err)
	}
}

func PreprocessImage(imagePath string) error {
	img, err := imaging.Open(imagePath)
	if err != nil {
		return err
	}

	// 1. Preprocessamento da Imagem
	img = imaging.Resize(img, 1500, 2000, imaging.Lanczos)
	img = imaging.Rotate90(img)
	img = imaging.Sharpen(img, 2.0)
	img = imaging.AdjustContrast(img, 80)
	img = imaging.Grayscale(img)

	return imaging.Save(img, "grayscale2.0"+imagePath)
}
