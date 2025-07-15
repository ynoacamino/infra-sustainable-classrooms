-- Knowledge Service Database Schema - Simplified
-- Simple form system: teachers create forms, students submit once and get a score

-- Tests table - stores form metadata (simplified)
DROP TABLE IF EXISTS tests CASCADE;
CREATE TABLE tests (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created_by BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Questions table - stores questions for tests (simplified, no JSON)
CREATE TABLE IF NOT EXISTS questions (
    id BIGSERIAL PRIMARY KEY,
    test_id BIGINT NOT NULL REFERENCES tests(id) ON DELETE CASCADE,
    question_text TEXT NOT NULL,
    option_a TEXT NOT NULL,
    option_b TEXT NOT NULL,
    option_c TEXT NOT NULL,
    option_d TEXT NOT NULL,
    correct_answer INTEGER NOT NULL CHECK (correct_answer BETWEEN 0 AND 3), -- 0=A, 1=B, 2=C, 3=D
    question_order INTEGER NOT NULL CHECK (question_order > 0),

    -- Ensure unique ordering within each test
    UNIQUE(test_id, question_order)
);

-- Test submissions table - tracks when users take tests (simplified)
CREATE TABLE IF NOT EXISTS test_submissions (
    id BIGSERIAL PRIMARY KEY,
    test_id BIGINT NOT NULL REFERENCES tests(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL, -- Reference to users.id from auth service
    score DECIMAL(5,2) NOT NULL, -- percentage score
    submitted_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,

    -- Ensure one submission per user per test
    UNIQUE(test_id, user_id)
);

-- Answer submissions table - individual answers for each question (simplified)
CREATE TABLE IF NOT EXISTS answer_submissions (
    id BIGSERIAL PRIMARY KEY,
    submission_id BIGINT NOT NULL REFERENCES test_submissions(id) ON DELETE CASCADE,
    question_id BIGINT NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    selected_answer INTEGER NOT NULL CHECK (selected_answer BETWEEN 0 AND 3),
    is_correct BOOLEAN NOT NULL,

    -- Ensure one answer per question per submission
    UNIQUE(submission_id, question_id)
);

-- === INDEXES FOR PERFORMANCE ===
CREATE INDEX IF NOT EXISTS idx_questions_test_id ON questions(test_id);
CREATE INDEX IF NOT EXISTS idx_submissions_test_id ON test_submissions(test_id);
CREATE INDEX IF NOT EXISTS idx_submissions_user_id ON test_submissions(user_id);
CREATE INDEX IF NOT EXISTS idx_answers_submission_id ON answer_submissions(submission_id);


-- seed data for tests
DO $$
DECLARE
  test_id BIGINT;
BEGIN
  -- 1. Matemáticas básicas
  INSERT INTO tests (title, created_by) VALUES
    ('Test de Matemáticas Básicas', 1)
    RETURNING id INTO test_id;

  INSERT INTO questions (test_id, question_text, option_a, option_b, option_c, option_d, correct_answer, question_order) VALUES
    (test_id, '¿Cuánto es 7 + 5?', '10', '12', '13', '11', 1, 1),
    (test_id, '¿Cuál es el doble de 8?', '14', '12', '16', '18', 2, 2),
    (test_id, '¿Cuánto es 9 x 3?', '27', '18', '21', '24', 0, 3),
    (test_id, '¿Cuál es la raíz cuadrada de 81?', '7', '8', '9', '10', 2, 4);

  -- 2. Historia universal
  INSERT INTO tests (title, created_by) VALUES
    ('Test de Historia Universal', 1)
    RETURNING id INTO test_id;

  INSERT INTO questions (test_id, question_text, option_a, option_b, option_c, option_d, correct_answer, question_order) VALUES
    (test_id, '¿En qué año ocurrió la Revolución Francesa?', '1789', '1492', '1810', '1776', 0, 1),
    (test_id, '¿Quién fue el primer emperador romano?', 'Julio César', 'Nerón', 'Augusto', 'Trajano', 2, 2),
    (test_id, '¿Qué civilización construyó las pirámides de Egipto?', 'Romana', 'Egipcia', 'Maya', 'Griega', 1, 3),
    (test_id, '¿Cuál fue el detonante de la Primera Guerra Mundial?', 'La caída de Berlín', 'El asesinato del archiduque Francisco Fernando', 'La Revolución Rusa', 'La invasión a Polonia', 1, 4);

  -- 3. Ciencias naturales
  INSERT INTO tests (title, created_by) VALUES
    ('Test de Ciencias Naturales', 1)
    RETURNING id INTO test_id;

  INSERT INTO questions (test_id, question_text, option_a, option_b, option_c, option_d, correct_answer, question_order) VALUES
    (test_id, '¿Cuál es el planeta más cercano al Sol?', 'Venus', 'Mercurio', 'Tierra', 'Marte', 1, 1),
    (test_id, '¿Qué órgano bombea sangre por todo el cuerpo?', 'Pulmón', 'Riñón', 'Hígado', 'Corazón', 3, 2),
    (test_id, '¿Qué gas respiramos del aire para vivir?', 'Oxígeno', 'Hidrógeno', 'Nitrógeno', 'Dióxido de carbono', 0, 3),
    (test_id, '¿Cómo se llama el proceso por el cual las plantas hacen su alimento?', 'Digestión', 'Evaporación', 'Fotosíntesis', 'Respiración', 2, 4);

  -- 4. Lengua y gramática
  INSERT INTO tests (title, created_by) VALUES
    ('Test de Lengua y Gramática', 1)
    RETURNING id INTO test_id;

  INSERT INTO questions (test_id, question_text, option_a, option_b, option_c, option_d, correct_answer, question_order) VALUES
    (test_id, '¿Cuál es un sinónimo de “feliz”?', 'Triste', 'Contento', 'Aburrido', 'Cansado', 1, 1),
    (test_id, '¿Cuál es el sujeto en la frase “El perro corre rápido”?', 'Corre', 'El', 'Perro', 'Rápido', 2, 2),
    (test_id, '¿Qué tipo de palabra es “rápidamente”?', 'Adjetivo', 'Sustantivo', 'Verbo', 'Adverbio', 3, 3),
    (test_id, '¿Qué signo se usa al final de una pregunta?', 'Punto', 'Coma', 'Signo de interrogación', 'Dos puntos', 2, 4);

  -- 5. Tecnología y computación
  INSERT INTO tests (title, created_by) VALUES
    ('Test de Tecnología y Computación', 1)
    RETURNING id INTO test_id;

  INSERT INTO questions (test_id, question_text, option_a, option_b, option_c, option_d, correct_answer, question_order) VALUES
    (test_id, '¿Qué significa “CPU”?', 'Unidad Central de Procesamiento', 'Computadora Personal Única', 'Procesador Universal', 'Centro de Usuario Principal', 0, 1),
    (test_id, '¿Qué lenguaje se usa comúnmente para páginas web?', 'Python', 'HTML', 'C++', 'SQL', 1, 2),
    (test_id, '¿Cuál es un sistema operativo?', 'Chrome', 'Facebook', 'Linux', 'Google', 2, 3),
    (test_id, '¿Qué es un byte?', 'Unidad de imagen', 'Tipo de virus', 'Unidad de almacenamiento', 'Archivo musical', 2, 4);
END;
$$;
