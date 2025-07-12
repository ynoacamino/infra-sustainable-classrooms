-- Tabla de cursos
CREATE TABLE courses (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description VARCHAR(300) NOT NULL,
    image_url VARCHAR(500),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Tabla de secciones
CREATE TABLE sections (
    id BIGSERIAL PRIMARY KEY,
    course_id BIGINT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(200) NOT NULL,
    "order" INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Tabla de artículos
CREATE TABLE articles (
    id BIGSERIAL PRIMARY KEY,
    section_id BIGINT NOT NULL REFERENCES sections(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Índices para optimización de consultas
CREATE INDEX idx_sections_course_id ON sections(course_id);
CREATE INDEX idx_articles_section_id ON articles(section_id);
CREATE INDEX idx_courses_title ON courses(title);
CREATE INDEX idx_sections_title ON sections(title);
CREATE INDEX idx_articles_title ON articles(title);
