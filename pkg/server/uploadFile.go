package server

import (
	"api-channel/pkg/conf"
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path"

	"github.com/h2non/bimg"
)

func (s *Server) UploadFile(ctx context.Context, req *proto.FileRequest) (*proto.FileResponse, error) {
	// token := req.Token
	reqId := req.RequestId

	// Check token is exist
	// _, err := checkToken(token)
	// if err != nil {
	// 	return nil, err
	// }

	db := database.GetInstance()
	filepath := ""
	fileType := "test"
	fileSize := 0
	if info := req.GetInfo(); info != nil {
		filepath = path.Join(conf.UPLOAD_DIR, path.Base(info.Name))
		fileSize = int(info.Size)
		file, err := createFile(filepath)
		if err != nil {
			return nil, fmt.Errorf("error creating file: %w", err)
		}
		files[reqId] = file
	}
	if chunk := req.GetChunk(); chunk != nil {
		if f, ok := files[reqId]; ok {
			f.Write(chunk)
		}
	}
	if req.GetDone() {
		// Write file path in Attachments table
		var attachment models.Attachment
		attachment = models.Attachment{
			MessageID: uint(req.MessageId),
			FilePath:  filepath,
			FileType:  fileType,
			FileSize:  fileSize,
		}
		res := db.Create(&attachment)

		// If file is image, Generate thumbnail of it.
		if fileType == "image" {
			buffer, err := bimg.Read(filepath)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			image := bimg.NewImage(buffer)

			// generate small size 64x64
			imageBytes, err := image.ForceResize(64, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			db.Create(models.Thumbnail{
				AttachmentID: attachment.ID,
				Base64:       base64.StdEncoding.EncodeToString(imageBytes),
				Type:         "small",
			})

			// generate mini size 32x32
			imageBytes, err = image.ForceResize(32, 32)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			db.Create(models.Thumbnail{
				AttachmentID: attachment.ID,
				Base64:       base64.StdEncoding.EncodeToString(imageBytes),
				Type:         "mini",
			})

			// generate blured and cropped
			imageBytes, err := image.
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			db.Create(models.Thumbnail{
				AttachmentID: attachment.ID,
				Base64:       base64.StdEncoding.EncodeToString(imageBytes),
			})
		}

		if f, ok := files[reqId]; ok {
			f.Close()
		}
		delete(files, reqId)
	}

	return &proto.FileResponse{}, nil
}

var files = map[string]*os.File{}

func createFile(filepath string) (*os.File, error) {
	var err error
	if err := os.MkdirAll(path.Dir(filepath), 0777); err != nil {
		return nil, err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
