package notes_collection

import (
	"context"

	_ "github.com/KikwiSan/course_project/notes/proto_files"
	pb "github.com/KikwiSan/course_project/notes/proto_files"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	Id   int64  `sql:"id"`
	Name string `sql:"name"`
}

type MongoRepository struct {
	db *mongo.Collection
}

type Repository interface {
	CreateNote(ctx context.Context, note *Note) error
	GetNote(ctx context.Context, id int64) (*Note, error)
	GetNotesByName(ctx context.Context, name string) ([]*Note, error)
	GetAllNotes(ctx context.Context) ([]*Note, error)
}

func (repos *MongoRepository) CreateNote(ctx context.Context, note *Note) {
	repos.db.InsertOne(ctx, note)
}

func (repos *MongoRepository) GetNote(ctx context.Context, id int64) *Note {
	var note *Note
	cursor, _ := repos.db.Find(ctx, bson.M{"id": id})
	cursor.All(ctx, &note)
	return note
}

func (repos *MongoRepository) GetAllNotes(ctx context.Context) []*Note {
	notes := make([]*Note, 0)
	cursor, _ := repos.db.Find(ctx, bson.M{})
	cursor.All(ctx, &notes)
	return notes
}

func (repos *MongoRepository) GetNotesByName(ctx context.Context, name string) []*Note {
	notes := make([]*Note, 0)
	cursor, _ := repos.db.Find(ctx, bson.M{"name": name})
	cursor.All(ctx, &notes)
	return notes
}

func ParsingNote(note *pb.Note) *Note {
	return &Note{
		Id:   note.Id,
		Name: note.Name,
	}
}

func UnparsingNote(note *Note) *pb.Note {
	return &pb.Note{
		Id:   note.Id,
		Name: note.Name,
	}
}

func ParsingAllNotes(notes []*pb.Note) []*Note {
	result := make([]*Note, len(notes))
	for _, val := range notes {
		result = append(result, ParsingNote(val))
	}
	return result
}

func UnparsingAllNotes(notes []*Note) []*pb.Note {
	result := make([]*pb.Note, len(notes))
	for _, val := range notes {
		result = append(result, UnparsingNote(val))
	}
	return result
}
