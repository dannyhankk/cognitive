package cognitive

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
)

func (s *MyServer) Chat2VoiceRequest(
	ctx context.Context, in *pb.ChatRequest) (
	out *pb.ChatResponse, err error) {

	return nil, nil
}