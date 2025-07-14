-- Knowledge Service Database Schema - Simplified
-- Simple form system: teachers create forms, students submit once and get a score

-- Tests table - stores form metadata (simplified)
CREATE TABLE IF NOT EXISTS tests (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created_by BIGINT NOT NULL, -- Reference to users.id from auth service
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
