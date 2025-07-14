-- Create exercises table
CREATE TABLE IF NOT EXISTS exercises (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    initial_code TEXT NOT NULL,
    solution TEXT NOT NULL,
    difficulty VARCHAR(20) NOT NULL CHECK (difficulty IN ('easy', 'medium', 'hard')),
    created_by BIGINT NOT NULL, -- Reference to users.id from auth service
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tests (
    id BIGSERIAL PRIMARY KEY,
    input TEXT NOT NULL,
    output TEXT NOT NULL,
    public BOOLEAN NOT NULL DEFAULT FALSE,
    exercise_id BIGINT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create answers table
CREATE TABLE IF NOT EXISTS answers (
    id BIGSERIAL PRIMARY KEY,
    exercise_id BIGINT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL, -- Reference to users.id from auth service
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (exercise_id, user_id)
);

-- Create attempts table
CREATE TABLE IF NOT EXISTS attempts (
    id BIGSERIAL PRIMARY KEY,
    answer_id BIGINT NOT NULL REFERENCES answers(id) ON DELETE CASCADE,
    code TEXT NOT NULL,
    success BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_exercises_difficulty ON exercises(difficulty);
CREATE INDEX IF NOT EXISTS idx_exercises_created_by ON exercises(created_by);
CREATE INDEX IF NOT EXISTS idx_exercises_created_at ON exercises(created_at);

CREATE INDEX IF NOT EXISTS idx_tests_exercise_id ON tests(exercise_id);
CREATE INDEX IF NOT EXISTS idx_tests_public ON tests(public);

CREATE INDEX IF NOT EXISTS idx_answers_exercise_id ON answers(exercise_id);
CREATE INDEX IF NOT EXISTS idx_answers_user_id ON answers(user_id);
CREATE INDEX IF NOT EXISTS idx_answers_completed ON answers(completed);

CREATE INDEX IF NOT EXISTS idx_attempts_answer_id ON attempts(answer_id);
CREATE INDEX IF NOT EXISTS idx_attempts_success ON attempts(success);
CREATE INDEX IF NOT EXISTS idx_attempts_created_at ON attempts(created_at);

