package server

import (
	"api-channel/pkg/conf"
	database "api-channel/pkg/db"
	"api-channel/pkg/helper"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path"
)

type File struct {
	d   *os.File
	bin bool
}

var files = map[string]*File{}

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
	fileType := ""
	fileSize := 0

	// Read file details and create file descriptor
	if info := req.GetInfo(); info != nil {
		filepath = path.Join(conf.UPLOAD_DIR, path.Base(info.Name))
		fileSize = int(info.Size)
		fileType = info.Type
		file, err := createFile(filepath)
		if err != nil {
			return nil, fmt.Errorf("error creating file: %w", err)
		}
		files[reqId] = &File{file, isBinary(fileType)}
	}
	// Read file chunck from client and write to file
	if chunk := req.GetChunk(); chunk != "" {
		if f, ok := files[reqId]; ok {

			var chunkByte = []byte(chunk)
			// If file is binary, decode base64 to bytes
			if f.bin {
				var err error
				chunkByte, err = base64.StdEncoding.DecodeString(chunk)
				if err != nil {
					return nil, fmt.Errorf("error decoding chunk: %w", err)
				}
			}
			f.d.Write(chunkByte)
		}
	}
	// Process whole file and save into database
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
			f.d.Close()
		}
		delete(files, reqId)
	}

	return &proto.FileResponse{}, nil
}

func isBinary(fileType string) bool {
	switch fileType {
	case "image", "video", "audio":
		return true
	case "txt", "text":
		return false
	}
	panic("unknown file type: " + fileType)
}
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
