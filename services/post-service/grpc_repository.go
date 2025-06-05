package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/post"
	userpb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostgresPostServiceRepository struct {
	DB *pgxpool.Pool
}

func NewPostgresPostServiceRepository(pool *pgxpool.Pool) *PostgresPostServiceRepository {
	return &PostgresPostServiceRepository{DB: pool}
}

func (r *PostgresPostServiceRepository) GetPost(ctx context.Context, postId int64) (*pb.Post, error) {

	query := `
			WITH selected_posts AS (
		SELECT id AS post_id, user_id, body, created_at, updated_at FROM posts WHERE id = $1
	)
			SELECT
				u.id AS user_id,
				u.username,
				u.email,
				u.updated_at AS user_updated_at,
				
				sp.post_id,
				sp.body,
				sp.created_at,
				sp.updated_at AS post_updated_at

			FROM selected_posts sp
	JOIN users u ON u.id = sp.user_id;

	`
	var post pb.Post
	var user userpb.User
	var userUpdatedAt time.Time
	var postCreatedAt time.Time
	var postUpdatedAt time.Time
	row := r.DB.QueryRow(ctx, query, postId)
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &userUpdatedAt, &post.Id, &post.Body, &postCreatedAt, &postUpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to scan: %w", err)

	}

	user.UpdatedAt = timestamppb.New(userUpdatedAt)
	post.User = &user
	post.CreatedAt = timestamppb.New(postCreatedAt)
	post.UpdatedAt = timestamppb.New(postUpdatedAt)
	return &post, nil

}
