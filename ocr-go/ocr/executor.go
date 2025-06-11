package ocr

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type TesseractExecutor struct {
	tesseractPath string
}

func NewTesseractExecutor() (*TesseractExecutor, error) {
	path, err := findTesseract()
	if err != nil {
		return nil, err
	}
	return &TesseractExecutor{tesseractPath: path}, nil
}

func findTesseract() (string, error) {
	if path, err := exec.LookPath("TESSERACT"); err == nil {
		return path, nil
	}

	if runtime.GOOS == "windows" {
		paths := []string{
			"C:\\Program Files\\Tesseract\\tesseract.exe", // change it (depende do path do tesseract do brother)
			"C:/Program Files/Tesseract/tesseract.exe",    // change it (depende do path do tesseract do brother)
		}
		for _, path := range paths {
			if _, err := os.Stat(path); err == nil {
				return path, nil
			}
		}
	}

	return "", errors.New("tesseract não encontrado no PATH ou nos locais padrão")
}

func (t *TesseractExecutor) ExtractText(imagePath string, lang string, psm int) (string, error) {
	args := []string{imagePath, "stdout"}

	if lang != "" {
		args = append(args, "-l", lang)
	}

	if psm > 0 {
		args = append(args, "--psm", fmt.Sprint(psm))
	}

	cmd := exec.Command(t.tesseractPath, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", errors.New(stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}
