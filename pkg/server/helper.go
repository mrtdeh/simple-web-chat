package server

// func init() {
// 	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
// }

// func errorResponse(err error) *proto.ServerResponse {
// 	return &proto.ServerResponse{
// 		Data: &proto.ServerResponse_ErrorResponse{
// 			ErrorResponse: &proto.ErrorResponse{
// 				Error: err.Error(),
// 			},
// 		},
// 	}
// }

// func (s *Server) sendToAll(msg *proto.ServerResponse) {
// 	wait := sync.WaitGroup{}
// 	done := make(chan int)
// 	grpcLog.Info("Sending message to: all")
// 	for _, conn := range s.Connections {
// 		wait.Add(1)

// 		go func(req *proto.ServerResponse, conn *Connection) {
// 			defer wait.Done()

// 			if conn.active {
// 				err := conn.stream.Send(req)

// 				if err != nil {
// 					grpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
// 					conn.active = false
// 					conn.error <- err
// 				}
// 			}

// 		}(msg, conn)

// 	}
// 	go func() {
// 		wait.Wait()
// 		close(done)
// 	}()

// 	<-done
// }

// func msgResponse(msg string, from *proto.User) *proto.ServerResponse {
// 	return &proto.ServerResponse{
// 		Data: &proto.ServerResponse_MsgResponse{
// 			MsgResponse: &proto.MessageResponse{
// 				From:      from.Username,
// 				Timestamp: time.Now().String(),
// 				Msg: &proto.Message{
// 					Data: &proto.Message_TextMsg{
// 						TextMsg: &proto.TextMessage{
// 							Content: msg,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func genToken() string {
// 	md5Sum := md5.Sum([]byte(time.Now().String()))
// 	return hex.EncodeToString(md5Sum[:])
// }
