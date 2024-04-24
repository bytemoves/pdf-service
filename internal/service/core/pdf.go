package core

import (
	"context"
	"log/slog"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PdfService interface {
	Process (ctx context.Context,html string) (pdf []byte, err error)

}

type pdfService struct {
	log *slog.Logger
}

func NewPdfService ( log  *slog.Logger) PdfService {
	return pdfService{log: log}
}


func ( s pdfService) Process(ctx context.Context,html string) (pdf []byte, err error) {
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
	pdfGenerator , err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		s.log.ErrorContext(ctx, "error while creating a new pdf generator")
		return nil , err
	}

	pdfGenerator.AddPage(page)

	err = pdfGenerator.Create()
	if err != nil {
		s.log.ErrorContext(ctx, "error while creating ne pdf page")
		return nil,err
	}
	return pdfGenerator.Bytes(),nil
}