package tests

import (
	"api-channel/proto"
	"context"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// تعریف اینترفیس دیتابیس
type Database interface {
	SaveMessage(msg string) error
}

// پیاده‌سازی تست
func TestCreateMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock دیتابیس
	mockDB := NewMockDatabase(ctrl)
	mockDB.EXPECT().SaveMessage("Hello World").Return(nil)

	// سرویس با mock دیتابیس
	service := &MessageService{db: mockDB}

	// درخواست gRPC
	req := &proto.MessageRequest{Content: "Hello World"}
	res, err := service.CreateMessage(context.Background(), req)

	// بررسی نتیجه
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.Success)
}
