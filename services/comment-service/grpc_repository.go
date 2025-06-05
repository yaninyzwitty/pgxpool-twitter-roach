package main

import (
	"context"
	"fmt"
	"time"

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
			WITH selected_comments AS (
			SELECT 
				c.id AS comment_id,
				c.user_id,
				c.body,
				c.created_at,
				c.updated_at
			FROM comments c
			WHERE c.id = $1
		)
		SELECT 
			sc.comment_id,
			sc.body,
			sc.created_at,
			sc.updated_at,

			u.id AS user_id,
			u.username,
			u.email,
			u.updated_at AS user_updated_at
		FROM selected_comments sc
		JOIN users u ON sc.user_id = u.id;
	`

	var comment pb.Comment
	var user userpb.User
	var commentCreatedAt, commentUpdatedAt time.Time
	var userUpdatedAt time.Time

	row := r.DB.QueryRow(ctx, query, commentID)

	if err := row.Scan(&comment.Id, &comment.Body, &commentCreatedAt, &commentUpdatedAt, &user.Id, &user.Username, &user.Email, &userUpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to scan: %w", err)
	}
	comment.User.UpdatedAt = timestamppb.New(userUpdatedAt)
	comment.User = &user
	comment.CreatedAt = timestamppb.New(commentCreatedAt)
	comment.UpdatedAt = timestamppb.New(commentUpdatedAt)
	return &comment, nil

}
