-- Tabla de categorías de videos
CREATE TABLE IF NOT EXISTS video_categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_video_categories_name ON video_categories (name);

-- Tabla de videos
CREATE TABLE IF NOT EXISTS video (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    user_id BIGINT NOT NULL,
    description TEXT,
    views INTEGER DEFAULT 0 NOT NULL,
    likes INTEGER DEFAULT 0 NOT NULL,
    video_obj_name VARCHAR(50),
    thumb_obj_name VARCHAR(50),
    category_id BIGINT NOT NULL REFERENCES video_categories (id) ON DELETE RESTRICT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_video_title ON video (title);

-- Tabla de etiquetas de videos
CREATE TABLE IF NOT EXISTS video_tags (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_video_tags_name ON video_tags (name);

-- Tabla de asociación entre videos y etiquetas
CREATE TABLE IF NOT EXISTS video_video_tags (
    video_id BIGINT NOT NULL REFERENCES video (id) ON DELETE CASCADE,
    tag_id BIGINT NOT NULL REFERENCES video_tags (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    PRIMARY KEY (video_id, tag_id)
);

-- Tabla de comentarios de videos
CREATE TABLE IF NOT EXISTS video_comments (
    id BIGSERIAL PRIMARY KEY,
    video_id BIGINT NOT NULL REFERENCES video (id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL,
    title VARCHAR(150) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Tabla de asociación entre usuarios y categorias preferidas
CREATE TABLE IF NOT EXISTS user_category_likes (
    user_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL REFERENCES video_categories (id) ON DELETE CASCADE,
    likes INTEGER DEFAULT 0 NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    PRIMARY KEY (user_id, category_id)
);

-- Tabla de asociación entre usuarios y videos que les gustan
CREATE TABLE IF NOT EXISTS user_video_likes (
    user_id BIGINT NOT NULL,
    video_id BIGINT NOT NULL REFERENCES video (id) ON DELETE CASCADE,
    liked BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    PRIMARY KEY (user_id, video_id)
);