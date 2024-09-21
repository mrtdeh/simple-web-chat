package server

import (
	"api-channel/pkg/conf"
	database "api-channel/pkg/db"
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

	msg := req.GetChunk()
	if msg == nil {
		panic("message is empty")
	}

	db := database.GetInstance()

	if info := req.GetInfo(); info != nil {
		file, err := createFile(info)
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
		if f, ok := files[reqId]; ok {
			f.Close()

		}
		delete(files, reqId)

		// Store file path and message into database

	}

	return &proto.FileResponse{}, nil
}

var files = map[string]*os.File{}

func createFile(info *proto.FileInfo) (*os.File, error) {
	var err error
	if err := os.MkdirAll(conf.UPLOAD_DIR, 0777); err != nil {
		return nil, err
	}
	filepath := path.Join(conf.UPLOAD_DIR, path.Base(info.Name))
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
