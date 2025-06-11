package ocr

// Config representa a configuração do OCR
type Config struct {
	Language     string // Idioma para reconhecimento (ex: "por", "eng")
	PageSegMode  int    // Modo de segmentação de página (1-13)
	TessdataPath string // Caminho para os dados de treinamento (opcional)
}

type OCR interface {
	ExtractText(imagePath string) (string, error)
	Close() error
}

func New(config Config) (OCR, error) {
	executor, err := NewTesseractExecutor()
	if err != nil {
		return nil, err
	}

	return &ocrService{
		executor: executor,
		config:   config,
	}, nil
}

type ocrService struct {
	executor *TesseractExecutor
	config   Config
}

func (o *ocrService) ExtractText(imagePath string) (string, error) {
	return o.executor.ExtractText(imagePath, o.config.Language, o.config.PageSegMode)
}

func (o *ocrService) Close() error {
	return nil
}
