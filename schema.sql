CREATE TABLE users (
    id BIGINT PRIMARY KEY DEFAULT unique_rowid(),
    username STRING NOT NULL,
    email STRING UNIQUE NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS posts (
    id BIGINT PRIMARY KEY DEFAULT unique_rowid(),
    user_id BIGINT NOT NULL REFERENCES users(id),
    body STRING NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


TABLE (users u JOIN posts p ON u.id = p.user_id); -- gives all data from joined tables

CREATE TABLE IF NOT EXISTS comments (
    id BIGINT PRIMARY KEY DEFAULT unique_rowid(),
    user_id BIGINT NOT NULL REFERENCES users(id),
    post_id BIGINT NOT NULL REFERENCES posts(id),
    body STRING NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


-- get users, count of post, comments
with 
post_counts as (
    select user_id, count(*) as post_count from posts
),
comment_counts as (
    select user_id, count(*) as comment_count from comments
)
select u.username, u.email, coalesce(pc.post_count, 0) as total_posts, coalesce(cc.comment_count, 0) as total_comments  from users u left join post_counts pc on u.id = pc.user_id left join comment_counts cc on u.id = cc.user_id ORDER BY total_comments desc, total_posts desc;
