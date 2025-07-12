-- Knowledge Service Database Schema
-- This schema supports the test and quiz functionality for the knowledge microservice

-- Using BIGINT IDs for consistency with other services
-- No UUID extension needed

-- === CORE TABLES ===

-- Test categories table - for organizing tests (must be created first)
CREATE TABLE IF NOT EXISTS test_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Tests table - stores test metadata
CREATE TABLE IF NOT EXISTS tests (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category_id INTEGER REFERENCES test_categories(id) ON DELETE SET NULL,
    difficulty_level VARCHAR(20) NOT NULL CHECK (difficulty_level IN ('easy', 'medium', 'hard')),
    duration_minutes INTEGER NOT NULL CHECK (duration_minutes > 0),
    passing_score DECIMAL(5,2) NOT NULL CHECK (passing_score >= 0 AND passing_score <= 100),
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE,
    instructions TEXT,
    created_by BIGINT NOT NULL, -- Reference to users.id from auth service
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    metadata JSONB DEFAULT '{}'::jsonb NOT NULL
);

-- Questions table - stores individual questions for tests
CREATE TABLE IF NOT EXISTS questions (
    id BIGSERIAL PRIMARY KEY,
    test_id BIGINT NOT NULL REFERENCES tests(id) ON DELETE CASCADE,
    question_text TEXT NOT NULL,
    options JSONB NOT NULL, -- Array of options as JSON
    correct_answer INTEGER NOT NULL CHECK (correct_answer >= 0),
    explanation TEXT,
    points INTEGER DEFAULT 1 CHECK (points > 0),
    question_order INTEGER NOT NULL CHECK (question_order > 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    metadata JSONB DEFAULT '{}'::jsonb NOT NULL,

    -- Ensure unique ordering within each test
    UNIQUE(test_id, question_order)
);

-- Test submissions table - tracks when users take tests
CREATE TABLE IF NOT EXISTS test_submissions (
    id BIGSERIAL PRIMARY KEY,
    test_id BIGINT NOT NULL REFERENCES tests(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL, -- Reference to users.id from auth service
    started_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    submitted_at TIMESTAMP WITH TIME ZONE,
    score DECIMAL(5,2),
    passed BOOLEAN,
    time_taken_minutes INTEGER,
    is_completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    metadata JSONB DEFAULT '{}'::jsonb NOT NULL,

    -- Ensure one submission per user per test
    UNIQUE(test_id, user_id)
);

-- Answer submissions table - individual answers for each question
CREATE TABLE IF NOT EXISTS answer_submissions (
    id BIGSERIAL PRIMARY KEY,
    submission_id BIGINT NOT NULL REFERENCES test_submissions(id) ON DELETE CASCADE,
    question_id BIGINT NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    selected_answer INTEGER NOT NULL CHECK (selected_answer >= 0),
    is_correct BOOLEAN NOT NULL,
    points_earned INTEGER DEFAULT 0 CHECK (points_earned >= 0),
    answered_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,

    -- Ensure one answer per question per submission
    UNIQUE(submission_id, question_id)
);



-- === INDEXES FOR PERFORMANCE ===

-- Test categories indexes
CREATE INDEX IF NOT EXISTS idx_categories_name ON test_categories(name);

-- Tests indexes
CREATE INDEX IF NOT EXISTS idx_tests_created_by ON tests(created_by);
CREATE INDEX IF NOT EXISTS idx_tests_category_id ON tests(category_id);
CREATE INDEX IF NOT EXISTS idx_tests_difficulty ON tests(difficulty_level);
CREATE INDEX IF NOT EXISTS idx_tests_is_active ON tests(is_active);
CREATE INDEX IF NOT EXISTS idx_tests_expires_at ON tests(expires_at);
CREATE INDEX IF NOT EXISTS idx_tests_created_at ON tests(created_at);

-- Questions indexes
CREATE INDEX IF NOT EXISTS idx_questions_test_id ON questions(test_id);
CREATE INDEX IF NOT EXISTS idx_questions_order ON questions(test_id, question_order);

-- Test submissions indexes
CREATE INDEX IF NOT EXISTS idx_submissions_test_id ON test_submissions(test_id);
CREATE INDEX IF NOT EXISTS idx_submissions_user_id ON test_submissions(user_id);
CREATE INDEX IF NOT EXISTS idx_submissions_user_test ON test_submissions(user_id, test_id);
CREATE INDEX IF NOT EXISTS idx_submissions_completed ON test_submissions(is_completed);
CREATE INDEX IF NOT EXISTS idx_submissions_submitted_at ON test_submissions(submitted_at);

-- Answer submissions indexes
CREATE INDEX IF NOT EXISTS idx_answers_submission_id ON answer_submissions(submission_id);
CREATE INDEX IF NOT EXISTS idx_answers_question_id ON answer_submissions(question_id);
CREATE INDEX IF NOT EXISTS idx_answers_correct ON answer_submissions(is_correct);

-- Insert default test categories
INSERT INTO test_categories (name, description) VALUES
    ('Mathematics', 'Mathematical concepts and problem solving'),
    ('Science', 'Scientific principles and experiments'),
    ('Language Arts', 'Reading comprehension and writing skills'),
    ('History', 'Historical events and analysis'),
    ('Environmental Science', 'Environmental sustainability and conservation'),
    ('Computer Science', 'Programming and computer concepts'),
    ('General Knowledge', 'Mixed topics and general awareness')
ON CONFLICT (name) DO NOTHING;

-- === COMMENTS FOR DOCUMENTATION ===

COMMENT ON TABLE tests IS 'Core table storing test metadata and configuration';
COMMENT ON TABLE questions IS 'Individual questions belonging to tests with multiple choice options';
COMMENT ON TABLE test_submissions IS 'Tracking table for user test attempts and results';
COMMENT ON TABLE answer_submissions IS 'Individual answers submitted by users for each question';
COMMENT ON TABLE test_categories IS 'Hierarchical categorization of tests';



COMMENT ON COLUMN questions.options IS 'JSON array of possible answers for multiple choice questions';
COMMENT ON COLUMN questions.correct_answer IS 'Zero-based index of the correct answer in the options array';

COMMENT ON COLUMN answer_submissions.selected_answer IS 'Zero-based index of the answer selected by the user';

-- === FOREIGN KEY RELATIONSHIPS ===
-- Note: In microservices architecture, we don't create actual foreign keys across services
-- These are logical relationships that should be enforced at the application level:

-- tests.created_by -> auth_service.users.id (teacher who created the test)
-- test_submissions.user_id -> auth_service.users.id (student taking the test)

-- For role validation, the knowledge service should call the profiles service via gRPC:
-- - Before creating a test: verify user has teacher role
-- - Before taking a test: verify user has student role
-- - A user can have both roles (teaching assistant scenario)
