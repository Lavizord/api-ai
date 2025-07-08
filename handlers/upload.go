package handlers

import (
	"fmt"
	"net/http"

	"api-ai/internal/logger"
	"api-ai/internal/services"

	"go.uber.org/zap"
)

// UploadPDF godoc
// @Summary      Upload a PDF file with a prompt
// @Description  Accepts a PDF file and an optional prompt for processing
// @Tags         api
// @Accept       multipart/form-data
// @Produce      plain
// @Param        pdf     formData  file   true  "PDF file to upload"
// @Success      200     {string}  string "PDF uploaded"
// @Failure      400     {object}  models.ErrorResponse
// @Router       /api/upload [post]
// @Security BearerAuth
func UploadPDF(w http.ResponseWriter, r *http.Request, fh *services.FileHandler) {
	// Parse form (max 10MB)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		logger.Default.Warn("Invalid form data size", zap.Error(err))
		http.Error(w, "Invalid form data size", http.StatusBadRequest)
		return
	}
	file, fileHeader, err := r.FormFile("pdf")
	if err != nil {
		logger.Default.Warn("PDF missing in form", zap.Error(err))
		http.Error(w, "PDF not found in request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = fh.HandleFileUpload(fileHeader, file, "upload")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "PDF uploaded")
}
