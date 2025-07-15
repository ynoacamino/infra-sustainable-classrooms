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


INSERT INTO video_categories (name) VALUES
  ('Education'),
  ('Technology'),
  ('Science'),
  ('Art'),
  ('Entertainment');

-- Tags
INSERT INTO video_tags (name) VALUES
  ('python'),
  ('docker'),
  ('minio'),
  ('sql'),
  ('tutorial');

-- Videos
INSERT INTO video (title, user_id, description, video_obj_name, thumb_obj_name, category_id)
VALUES
  ('Python Basics', 1, 'A beginner tutorial on Python.', 'v1.mp4', 'i1.webp', (SELECT id FROM video_categories WHERE name = 'Education')),
  ('Docker Guide', 2, 'Learn how to containerize apps with Docker.', 'v2.mp4', 'i2.webp', (SELECT id FROM video_categories WHERE name = 'Technology')),
  ('Space Explained', 3, 'An educational video about the universe.', 'v3.mp4', 'i3.webp', (SELECT id FROM video_categories WHERE name = 'Science')),
  ('Watercolor Painting', 1, 'Learn how to paint with watercolors.', 'v4.mp4', 'i4.webp', (SELECT id FROM video_categories WHERE name = 'Art')),
  ('Top 5 Movies of 2025', 4, 'Reviewing the best films of the year.', 'v1.mp4', 'i2.webp', (SELECT id FROM video_categories WHERE name = 'Entertainment'));

-- Tag Relations
INSERT INTO video_video_tags (video_id, tag_id)
VALUES
  (1, (SELECT id FROM video_tags WHERE name = 'python')),
  (1, (SELECT id FROM video_tags WHERE name = 'tutorial')),
  (2, (SELECT id FROM video_tags WHERE name = 'docker')),
  (2, (SELECT id FROM video_tags WHERE name = 'tutorial')),
  (3, (SELECT id FROM video_tags WHERE name = 'minio')), -- optional, just example
  (4, (SELECT id FROM video_tags WHERE name = 'sql')),
  (5, (SELECT id FROM video_tags WHERE name = 'tutorial'));

-- Video Comments
INSERT INTO video_comments (video_id, user_id, title, content)
VALUES
  (1, 2, 'Great tutorial', 'Really helpful introduction to Python.'),
  (2, 3, 'Loved it!', 'Docker finally makes sense.'),
  (3, 1, 'Wow', 'Beautiful visuals of the cosmos.'),
  (4, 2, 'Creative', 'Nice techniques with watercolors.'),
  (5, 1, 'Fun review', 'Totally agree with your top pick!');

-- Category Likes
INSERT INTO user_category_likes (user_id, category_id, likes)
VALUES
  (1, (SELECT id FROM video_categories WHERE name = 'Education'), 3),
  (1, (SELECT id FROM video_categories WHERE name = 'Technology'), 2),
  (1, (SELECT id FROM video_categories WHERE name = 'Science'), 1);
