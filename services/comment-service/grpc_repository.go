package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/comment"
	userpb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostgresCommentServiceRepository struct {
	DB *pgxpool.Pool
}

func NewPostgresCommentServiceRepository(pool *pgxpool.Pool) *PostgresCommentServiceRepository {
	return &PostgresCommentServiceRepository{DB: pool}
}

func (r *PostgresCommentServiceRepository) GetComment(ctx context.Context, commentID int64) (*pb.Comment, error) {
	query := `
		SELECT 
			c.id AS comment_id,
			c.post_id,
			c.user_id,
			c.body,
			c.created_at,
			c.updated_at,
			u.id AS user_id,
			u.username AS user_name,
			u.email AS user_email,
			u.created_at AS user_created_at
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.id = $1;
	`

	row := r.DB.QueryRow(ctx, query, commentID)

	var comment pb.Comment
	var user userpb.User
	var commentCreatedAt, commentUpdatedAt time.Time
	var userCreatedAt time.Time

	err := row.Scan(
		&comment.Id,
		&comment.PostId,
		new(any), // skipping c.user_id, already in user
		&comment.Body,
		&commentCreatedAt,
		&commentUpdatedAt,
		&user.Id,
		&user.Username,
		&user.Email,
		&userCreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no rows found: %w", err)
		}
		return nil, err
	}

	comment.CreatedAt = timestamppb.New(commentCreatedAt)
	comment.UpdatedAt = timestamppb.New(commentUpdatedAt)
	user.CreatedAt = timestamppb.New(userCreatedAt)
	comment.User = &user

	return &comment, nil
}
