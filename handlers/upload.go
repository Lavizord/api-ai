package handlers

import (
	"fmt"
	"net/http"

	"api-ai/internal/logger"
)

func UploadPDF(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, "PDF not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	logger.Default.Info("PDF received")
	fmt.Fprintln(w, "PDF uploaded")
}
