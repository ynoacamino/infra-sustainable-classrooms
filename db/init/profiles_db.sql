-- Tabla principal de perfiles
CREATE TABLE IF NOT EXISTS profiles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE, -- Referencia al usuario en auth_db
    role VARCHAR(20) NOT NULL CHECK (role IN ('student', 'teacher')),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    avatar_url TEXT,
    bio TEXT,
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Tabla de perfiles de estudiantes
CREATE TABLE IF NOT EXISTS student_profiles (
    id BIGSERIAL PRIMARY KEY,
    profile_id BIGINT NOT NULL REFERENCES profiles(id) ON DELETE CASCADE,
    grade_level VARCHAR(50) NOT NULL,
    major VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Tabla de perfiles de profesores
CREATE TABLE IF NOT EXISTS teacher_profiles (
    id BIGSERIAL PRIMARY KEY,
    profile_id BIGINT NOT NULL REFERENCES profiles(id) ON DELETE CASCADE,
    position VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Índices para optimizar consultas
CREATE INDEX IF NOT EXISTS idx_profiles_user_id ON profiles(user_id);
CREATE INDEX IF NOT EXISTS idx_profiles_role ON profiles(role);
CREATE INDEX IF NOT EXISTS idx_profiles_email ON profiles(email);
CREATE INDEX IF NOT EXISTS idx_profiles_is_active ON profiles(is_active);
CREATE INDEX IF NOT EXISTS idx_profiles_created_at ON profiles(created_at);

CREATE INDEX IF NOT EXISTS idx_student_profiles_profile_id ON student_profiles(profile_id);
CREATE INDEX IF NOT EXISTS idx_student_profiles_grade_level ON student_profiles(grade_level);
CREATE INDEX IF NOT EXISTS idx_student_profiles_major ON student_profiles(major);

CREATE INDEX IF NOT EXISTS idx_teacher_profiles_profile_id ON teacher_profiles(profile_id);
CREATE INDEX IF NOT EXISTS idx_teacher_profiles_position ON teacher_profiles(position);

-- Función para validar que el perfil coincida con el rol
CREATE OR REPLACE FUNCTION validate_profile_role()
RETURNS TRIGGER AS $$
BEGIN
    -- Validar que solo estudiantes tengan student_profiles
    IF TG_TABLE_NAME = 'student_profiles' THEN
        IF NOT EXISTS (
            SELECT 1 FROM profiles 
            WHERE id = NEW.profile_id AND role = 'student'
        ) THEN
            RAISE EXCEPTION 'Cannot create student profile for non-student user';
        END IF;
    END IF;
    
    -- Validar que solo profesores tengan teacher_profiles
    IF TG_TABLE_NAME = 'teacher_profiles' THEN
        IF NOT EXISTS (
            SELECT 1 FROM profiles 
            WHERE id = NEW.profile_id AND role = 'teacher'
        ) THEN
            RAISE EXCEPTION 'Cannot create teacher profile for non-teacher user';
        END IF;
    END IF;
    
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Aplicar triggers de validación
CREATE TRIGGER validate_student_profile_role BEFORE INSERT OR UPDATE ON student_profiles
    FOR EACH ROW EXECUTE FUNCTION validate_profile_role();

CREATE TRIGGER validate_teacher_profile_role BEFORE INSERT OR UPDATE ON teacher_profiles
    FOR EACH ROW EXECUTE FUNCTION validate_profile_role();
