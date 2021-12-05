<<<<<<< HEAD
package main

import (
	"context"

	"github.com/KikwiSan/CourseProject/notes/notes_collection"
	pb "github.com/KikwiSan/CourseProject/notes/proto_files"
)

type NoteServer struct {
	repository notes_collection.Repository
}

func (s *NoteServer) GetNote(ctx context.Context, bReq *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	result, _ := s.repository.GetNote(ctx, bReq.Id)
	note := notes_collection.UnparsingNote(result)
	return &pb.GetNoteResponse{
		Title:   note.GetTitle(),
		Author:  note.GetAuthor(),
		IsOrder: note.IsOrder,
	}, nil
}

func (s *NoteServer) CreateNote(ctx context.Context, newReq *pb.Note) (*pb.CreateNoteResponse, error) {
	_ = s.repository.CreateNote(ctx, notes_collection.ParsingNote(newReq))
	return &pb.CreateNoteResponse{
		Note: newReq,
	}, nil
}

func (s *NoteServer) GetAllNotes(ctx context.Context, allReq *pb.GetAllNotesRequest) (*pb.GetAllNotesResponse, error) {
	notes_, _ := s.repository.GetAllNotes(ctx)
	notes := notes_collection.UnparsingAllNotes(notes_)
	return &pb.GetAllNotesResponse{
		AllNotes: notes,
	}, nil
}

func (s *NoteServer) GetNotesByAuthor(ctx context.Context, auReq *pb.GetNotesByAuthorRequest) (*pb.GetNotesByAuthorResponse, error) {
	notes_, _ := s.repository.GetNotesByAuthor(ctx, auReq.Author)
	notes := notes_collection.UnparsingAllNotes(notes_)
	return &pb.GetNotesByAuthorResponse{
		Notes: notes,
	}, nil
}
=======
package main

import (
	"context"

	"github.com/KikwiSan/CourseProject/notes/notes_collection"
	pb "github.com/KikwiSan/CourseProject/notes/proto_files"
)

type NoteServer struct {
	repository notes_collection.Repository
}

func (s *NoteServer) GetNote(ctx context.Context, bReq *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	result, _ := s.repository.GetNote(ctx, bReq.Id)
	note := notes_collection.UnparsingNote(result)
	return &pb.GetNoteResponse{
		Title:   note.GetTitle(),
		Author:  note.GetAuthor(),
		IsOrder: note.IsOrder,
	}, nil
}

func (s *NoteServer) CreateNote(ctx context.Context, newReq *pb.Note) (*pb.CreateNoteResponse, error) {
	_ = s.repository.CreateNote(ctx, notes_collection.ParsingNote(newReq))
	return &pb.CreateNoteResponse{
		Note: newReq,
	}, nil
}

func (s *NoteServer) GetAllNotes(ctx context.Context, allReq *pb.GetAllNotesRequest) (*pb.GetAllNotesResponse, error) {
	notes_, _ := s.repository.GetAllNotes(ctx)
	notes := notes_collection.UnparsingAllNotes(notes_)
	return &pb.GetAllNotesResponse{
		AllNotes: notes,
	}, nil
}

func (s *NoteServer) GetNotesByAuthor(ctx context.Context, auReq *pb.GetNotesByAuthorRequest) (*pb.GetNotesByAuthorResponse, error) {
	notes_, _ := s.repository.GetNotesByAuthor(ctx, auReq.Author)
	notes := notes_collection.UnparsingAllNotes(notes_)
	return &pb.GetNotesByAuthorResponse{
		Notes: notes,
	}, nil
}
>>>>>>> 9006d60a1538c3867b12bcf075dde1053ad2a85e
