package services

import (
	"api-ai/ent"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
)

// This will be our main File Handler struct.
//
// It will contain any references to objects that will be used and handle most
// of the laber of handlign the file in our system (upload, possible interactions with AI interface)
type FileHandler struct {
	db *ent.Client
	r2 *R2Uploader
}

func NewFileHandler(db *ent.Client) (*FileHandler, error) {
	r2env, err := LoadR2Venv()
	if err != nil {
		return nil, fmt.Errorf("failed to load R2 env: %w", err)
	}
	r2c, err := NewR2Uploader(r2env)
	if err != nil {
		return nil, fmt.Errorf("failed to create R2 client: %w", err)
	}
	return &FileHandler{db: db, r2: r2c}, nil
}

func (f *FileHandler) CreateFileEntry(filename, source string) (int, error) {
	fileRecord, err := f.db.Files.Create().
		SetFileSource(source).
		SetFileName(filename).
		SetType("pdf").
		Save(context.Background())
	if err != nil {
		return 0, err
	}
	return fileRecord.ID, nil
}

func (f *FileHandler) CreateFileEntryWithData(filename, source, prompt string, fileBytes []byte) (int, error) {

	fileRecord, err := f.db.Files.Create().
		SetFileSource(source).
		SetFileName(filename).
		SetPromptUsed(prompt).
		SetType("pdf").
		SetFileData(fileBytes).
		Save(context.Background())
	if err != nil {
		return 0, err
	}
	return fileRecord.ID, nil
}

func (f *FileHandler) HandleFileUpload(fileHeader *multipart.FileHeader, file multipart.File, source string) error {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	fileid, err := f.CreateFileEntryWithData(fileHeader.Filename, source, f.GetFilePrompt(), fileBytes)
	if err != nil {
		return err
	}

	url, err := f.UploadFileToCloudflare(fileid, fileBytes)
	if err != nil {
		return fmt.Errorf("error uploading file to Cloudflare: %v", err)
	}

	_, err = f.UpdateFileUrl(fileid, url)
	if err != nil {
		return fmt.Errorf("error updating url field: %v", err)
	}
	// TODO: we then send the file to chatpdf.
	// TODO: we then return the response.

	return nil
}

func (f *FileHandler) UploadFileToCloudflare(fileid int, fileBytes []byte) (string, error) {
	key := "uploads/" + strconv.Itoa(fileid)
	url, err := f.r2.UploadFile(context.Background(), key, fileBytes)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (f *FileHandler) GetFilePrompt() string {
	return ""
}

func (f *FileHandler) UpdateFileUrl(id int, url string) (int, error) {
	fileRecord, err := f.db.Files.UpdateOneID(id).
		SetFileURL(url).
		Save(context.Background())
	if err != nil {
		return 0, err
	}
	return fileRecord.ID, nil
}
