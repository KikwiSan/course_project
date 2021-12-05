package notes

import (
	"context"

	_ "github.com/KikwiSan/course_project/notes/proto_files"
	pb "github.com/KikwiSan/course_project/notes/proto_files"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, err
}

type Note struct {
	Id           int32  `sql:"id"`
	Name         string `sql:"name"`
	Note_content string `sql:"note_content"`
}

type MongoRepository struct {
	db *mongo.Collection
}

type Repository interface {
	Create(ctx context.Context, note *Note) error
	Get(ctx context.Context, id int64) (*Note, error)
	GetNotesByName(ctx context.Context, name string) ([]*Note, error)
	GetAllNotes(ctx context.Context) ([]*Note, error)
}

func (repos *MongoRepository) Create(ctx context.Context, note *Note) {
	repos.db.InsertOne(ctx, note)
}

func (repos *MongoRepository) Get(ctx context.Context, id int32) *Note {
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

func ConvertingNote(note *pb.Note) *Note {
	return &Note{
		Id:           note.Id,
		Name:         note.Name,
		Note_content: note.Note_content,
	}
}

func UnconvertingNote(note *Note) *pb.Note {
	return &pb.Note{
		Id:           note.Id,
		Name:         note.Name,
		Note_content: note.Note_content,
	}
}

func ConveringAllNotes(notes []*pb.Note) []*Note {
	temp := make([]*Note, len(notes))
	for _, val := range notes {
		temp = append(temp, ConvertingNote(val))
	}
	return temp
}

func UnconvertingAllNotes(notes []*Note) []*pb.Note {
	temp := make([]*pb.Note, len(notes))
	for _, val := range notes {
		temp = append(temp, UnconvertingNote(val))
	}
	return temp
}
