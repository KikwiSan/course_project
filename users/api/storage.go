package api

import (
	"context"

	_ "github.com/KikwiSan/course_project/users"
	pb "github.com/KikwiSan/course_project/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	hostname      = "localhost"
	host_port     = 5432
	username      = "123"
	password      = "123"
	database_name = "users"
)

func Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, err
}

type User struct {
	ID       int32
	Name     string
	Password string
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int32) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
}

type MongoRepository struct {
	db *mongo.Collection
}

func (repos *MongoRepository) CreateUser(ctx context.Context, user *User) {
	repos.db.InsertOne(ctx, user)
}

func (repos *MongoRepository) GetUser(ctx context.Context, id int32) *User {
	var user *User
	cursor, _ := repos.db.Find(ctx, bson.M{"id": id})
	cursor.All(ctx, &user)
	return user
}

func (repos *MongoRepository) GetAllUsers(ctx context.Context) []*User {
	users := make([]*User, 0)
	cursor, _ := repos.db.Find(ctx, bson.M{})
	cursor.All(ctx, &users)
	return users
}

func ConvertingUser(user *pb.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Password: user.Password,
	}
}

func UnconvertingUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}
}

func ConvertingAllUsers(users []*pb.User) []*User {
	temp := make([]*User, len(users))
	for _, val := range users {
		temp = append(temp, ConvertingUser(val))
	}
	return temp
}

func UnconvertingAllUsers(users []*User) []*pb.User {
	temp := make([]*pb.User, len(users))
	for _, val := range users {
		temp = append(temp, UnconvertingUser(val))
	}
	return temp
}