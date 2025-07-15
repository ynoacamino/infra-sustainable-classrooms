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

-- Seed data for exercises
DO $$
DECLARE
  i INT;
  exercise_id BIGINT;
  answer_id BIGINT;
  titles TEXT[] := ARRAY[
    'Sumar números desde un string',
    'Restar dos números',
    'Multiplicar dos números',
    'Dividir dos números',
    'Potencia de un número'
  ];
  descriptions TEXT[] := ARRAY[
    'Dado un string como "2,3", devuelve la suma como string.',
    'Dado un string como "10,4", devuelve la resta como string.',
    'Dado un string como "3,5", devuelve la multiplicación como string.',
    'Dado un string como "20,4", devuelve la división como string.',
    'Dado un string como "2,4", devuelve la potencia como string.'
  ];
  difficulties TEXT[] := ARRAY[
    'easy', 'easy', 'medium', 'medium', 'hard'
  ];
  solutions TEXT[] := ARRAY[
    'function solution(input) { const [a, b] = input.split(",").map(Number); return String(a + b); }',
    'function solution(input) { const [a, b] = input.split(",").map(Number); return String(a - b); }',
    'function solution(input) { const [a, b] = input.split(",").map(Number); return String(a * b); }',
    'function solution(input) { const [a, b] = input.split(",").map(Number); return String(a / b); }',
    'function solution(input) { const [a, b] = input.split(",").map(Number); return String(Math.pow(a, b)); }'
  ];
  incorrect_solutions TEXT[] := ARRAY[
    'function solution(input) { return "incorrecto"; }',
    'function solution(input) { return input; }',
    'function solution(input) { return "42"; }',
    'function solution(input) { return null; }',
    'function solution(input) { return ""; }'
  ];
BEGIN
  FOR i IN 1..5 LOOP
    -- Insertar ejercicio
    INSERT INTO exercises (
      title, description, initial_code, solution, difficulty, created_by
    ) VALUES (
      titles[i],
      descriptions[i],
      'function solution(input) {\n  // tu código aquí\n}',
      solutions[i],
      difficulties[i],
      1
    ) RETURNING id INTO exercise_id;

    -- Insertar tests
    INSERT INTO tests (input, output, public, exercise_id) VALUES
      ('2,3', '5', TRUE, exercise_id),
      ('10,20', '30', TRUE, exercise_id),
      ('-5,5', '0', FALSE, exercise_id);

    -- Insertar respuesta
    INSERT INTO answers (exercise_id, user_id, completed)
    VALUES (exercise_id, 1, TRUE)
    RETURNING id INTO answer_id;

    -- Insertar intento exitoso
    INSERT INTO attempts (answer_id, code, success)
    VALUES (
      answer_id,
      solutions[i],
      TRUE
    );

    -- Insertar intento fallido
    INSERT INTO attempts (answer_id, code, success)
    VALUES (
      answer_id,
      incorrect_solutions[i],
      FALSE
    );
  END LOOP;
END;
$$;
