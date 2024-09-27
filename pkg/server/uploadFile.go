package server

import (
	"api-channel/pkg/conf"
	database "api-channel/pkg/db"
	"api-channel/pkg/helper"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
	"fmt"
	"os"
	"path"
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
	fileType := "image"
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
		// Check and validate file type
		// ....

		// Write file path in Attachments table
		attachment := models.Attachment{
			MessageID: uint(req.MessageId),
			FilePath:  filepath,
			FileType:  fileType,
			FileSize:  fileSize,
		}
		db.Create(&attachment)

		// If file is image, Generate thumbnail of it.
		if fileType == "image" {
			img, err := helper.OpenImage(filepath)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			// generate small size 64x64
			thum64 := helper.Thumbnail(img, 64)
			db.Create(models.Thumbnail{
				AttachmentID: attachment.ID,
				Base64:       helper.Base64Image(thum64, 1),
				Type:         "small",
			})

			// generate mini size 32x32
			thum32 := helper.Thumbnail(img, 32)
			db.Create(models.Thumbnail{
				AttachmentID: attachment.ID,
				Base64:       helper.Base64Image(thum32, 1),
				Type:         "mini",
			})

			// generate image placeholder
			placeholder := helper.Placeholder(img)
			db.Create(models.Thumbnail{
				AttachmentID: attachment.ID,
				Base64:       helper.Base64Image(placeholder, 1),
				Type:         "placeholder",
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
