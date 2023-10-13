package invoicer

import (
	context "context"
	"log"
)

type Server struct {
	UnimplementedInvoicerServer
}

func (s *Server) Create(crx context.Context, req *CreateRequest) (*CreateResponse, error) {
	log.Printf("Received message from client")

	return &CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}
