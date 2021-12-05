package notes

import (
	"context"

	"github.com/KikwiSan/course_project/notes"
	pb "github.com/KikwiSan/course_project/notes/proto_files"
)

type NoteServer struct {
	repository notes.Repository
}

func (s *NoteServer) Get(ctx context.Context, noteReq *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	temp, _ := s.repository.Get(ctx, noteReq.Id)
	note := notes.UnconvertingNote(temp)
	return &pb.GetNoteResponse{
		Id:           note.GetId(),
		Name:         note.GetName(),
		Note_content: note.Note_content,
	}, nil
}

func (s *NoteServer) Create(ctx context.Context, newReq *pb.Note) (*pb.CreateNoteResponse, error) {
	_ = s.repository.Create(ctx, notes.ConvertingNote(newReq))
	return &pb.CreateNoteResponse{
		Note: newReq,
	}, nil
}

func (s *NoteServer) GetAll(ctx context.Context, allReq *pb.GetAllNotesRequest) (*pb.GetAllNotesResponse, error) {
	notes_, _ := s.repository.GetAll(ctx)
	notes := notes.UnconvertingAllNotes(notes_)
	return &pb.GetAllNotesResponse{
		AllNotes: notes,
	}, nil
}

func (s *NoteServer) GetNotesByName(ctx context.Context, nameReq *pb.GetNotesByNamerRequest) (*pb.GetNotesByNameResponse, error) {
	notes_, _ := s.repository.GetNotesByName(ctx, nameReq.Name)
	notes := notes.UnconvertingAllNotes(notes_)
	return &pb.GetNotesByNameResponse{
		Notes: notes,
	}, nil
}
