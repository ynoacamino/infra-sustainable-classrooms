-- Tabla de categorías de videos
CREATE TABLE IF NOT EXISTS video_categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
);

-- Tabla de videos
CREATE TABLE IF NOT EXISTS video (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT,
    upload_date TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    views INTEGER DEFAULT 0 NOT NULL,
    likes INTEGER DEFAULT 0 NOT NULL,
    video_obj_name VARCHAR(50),
    thumb_obj_name VARCHAR(50),
    category_id BIGINT NOT NULL REFERENCES video_categories (id) ON DELETE RESTRICT,
);

CREATE INDEX IF NOT EXISTS idx_video_title ON video (title);

-- Tabla de etiquetas de videos
CREATE TABLE IF NOT EXISTS video_tags (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
);

-- Tabla de asociación entre videos y etiquetas
CREATE TABLE IF NOT EXISTS video_video_tags (
    video_id BIGINT NOT NULL REFERENCES video (id) ON DELETE CASCADE,
    tag_id BIGINT NOT NULL REFERENCES video_tags (id) ON DELETE CASCADE,
    PRIMARY KEY (video_id, tag_id)
);

-- Tabla de comentarios de videos
CREATE TABLE IF NOT EXISTS video_comments (
    id BIGSERIAL PRIMARY KEY,
    video_id BIGINT NOT NULL REFERENCES video (id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL,
    title VARCHAR(150) NOT NULL,
    content TEXT NOT NULL,
    publish_date TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
);

-- Tabla de respuestas a comentarios
CREATE TABLE IF NOT EXISTS video_comment_replies (
    id BIGSERIAL PRIMARY KEY,
    comment_id BIGINT NOT NULL REFERENCES video_comments (id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    publish_date TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
);

-- Tabla de asociación entre usuarios y categorias preferidas
CREATE TABLE IF NOT EXISTS user_category_likes (
    user_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL REFERENCES video_categories (id) ON DELETE CASCADE,
    likes INTEGER DEFAULT 0 NOT NULL,
    PRIMARY KEY (user_id, category_id)
);