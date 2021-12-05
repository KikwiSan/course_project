<<<<<<< HEAD
package api

import (
	"context"
	"log"

	_ "github.com/KikwiSan/CourseProject/users"
	pb "github.com/KikwiSan/CourseProject/users"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int64  `sql:"id"`
	Name     string `sql:"name"`
	Email    string `sql:"email"`
	Password string `sql:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
}

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func (repos *PostgresRepository) GetUser(ctx context.Context, id int64) (*User, error) {
	var user *User
	if err := repos.db.GetContext(ctx, &user, "select * from users where id = $1", id); err != nil {
		return nil, err
	}
	return user, nil
}

func (repos *PostgresRepository) CreateUser(ctx context.Context, user *User) error {
	log.Println(user)
	query := "insert into users (id, name, email, password, ordered_books) values ($1, $2, $3, $4, $5)"
	_, err := repos.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Password)
	return err
}

func (repos *PostgresRepository) GetAllUsers(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := repos.db.GetContext(ctx, users, "select * from users"); err != nil {
		return users, err
	}
	return users, nil
}

func ParsingUser(user *pb.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UnparsingUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ParsingAllUsers(users []*pb.User) []*User {
	result := make([]*User, len(users))
	for _, val := range users {
		result = append(result, ParsingUser(val))
	}
	return result
}

func UnparsingAllUsers(users []*User) []*pb.User {
	result := make([]*pb.User, len(users))
	for _, val := range users {
		result = append(result, UnparsingUser(val))
	}
	return result
}
=======
package api

import (
	"context"
	"log"

	_ "github.com/KikwiSan/CourseProject/users"
	pb "github.com/KikwiSan/CourseProject/users"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int64  `sql:"id"`
	Name     string `sql:"name"`
	Email    string `sql:"email"`
	Password string `sql:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
}

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func (repos *PostgresRepository) GetUser(ctx context.Context, id int64) (*User, error) {
	var user *User
	if err := repos.db.GetContext(ctx, &user, "select * from users where id = $1", id); err != nil {
		return nil, err
	}
	return user, nil
}

func (repos *PostgresRepository) CreateUser(ctx context.Context, user *User) error {
	log.Println(user)
	query := "insert into users (id, name, email, password, ordered_books) values ($1, $2, $3, $4, $5)"
	_, err := repos.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Password)
	return err
}

func (repos *PostgresRepository) GetAllUsers(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := repos.db.GetContext(ctx, users, "select * from users"); err != nil {
		return users, err
	}
	return users, nil
}

func ParsingUser(user *pb.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UnparsingUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ParsingAllUsers(users []*pb.User) []*User {
	result := make([]*User, len(users))
	for _, val := range users {
		result = append(result, ParsingUser(val))
	}
	return result
}

func UnparsingAllUsers(users []*User) []*pb.User {
	result := make([]*pb.User, len(users))
	for _, val := range users {
		result = append(result, UnparsingUser(val))
	}
	return result
}
>>>>>>> 9006d60a1538c3867b12bcf075dde1053ad2a85e
