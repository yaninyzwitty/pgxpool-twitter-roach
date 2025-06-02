package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostgresSocialServiceRepository struct {
	DB *pgxpool.Pool
}

func NewPostgresSocialServiceRepository(pool *pgxpool.Pool) *PostgresSocialServiceRepository {
	return &PostgresSocialServiceRepository{DB: pool}
}

func (r *PostgresSocialServiceRepository) GetUserById(ctx context.Context, userId uint64) (*pb.User, error) {
	query := `SELECT id, email, username, created_at FROM users WHERE id = $1`
	row := r.DB.QueryRow(ctx, query, userId)

	var u pb.User
	var createdAt time.Time
	if err := row.Scan(&u.Id, &u.Email, &u.Username, &createdAt); err != nil {
		return nil, err
	}

	u.CreatedAt = timestamppb.New(createdAt)
	return &u, nil
}
func (r *PostgresSocialServiceRepository) GetUserByEmail(ctx context.Context, email string) (*pb.User, error) {
	query := `SELECT id, email, username, created_at FROM users WHERE email = $1`
	row := r.DB.QueryRow(ctx, query, email)

	var u pb.User
	var createdAt time.Time
	if err := row.Scan(&u.Id, &u.Email, &u.Username, &createdAt); err != nil {
		return nil, err
	}

	u.CreatedAt = timestamppb.New(createdAt)
	return &u, nil
}
func (r *PostgresSocialServiceRepository) GetUsers(ctx context.Context, limit, offset int) ([]*pb.User, error) {
	query := `SELECT id, email, username, created_at FROM users ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := r.DB.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		var user pb.User
		var createdAt time.Time
		if err := rows.Scan(&user.Id, &user.Email, &user.Username, &createdAt); err != nil {
			return nil, err
		}
		user.CreatedAt = timestamppb.New(createdAt)
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
