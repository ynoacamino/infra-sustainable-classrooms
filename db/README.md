# Database Configuration

This directory contains the database initialization scripts and schemas for the Sustainable Classrooms platform. The project uses a multi-database architecture with PostgreSQL, where each microservice has its own dedicated database.

## 📊 Database Architecture

### Multi-Database Strategy

The platform uses database-per-service pattern for better:

- **Isolation**: Each service owns its data
- **Scalability**: Services can scale independently
- **Reliability**: Failures are contained to individual services
- **Technology Choice**: Services can use different storage technologies

### Database List

| Database            | Service                | Purpose                             | Schema File                                         |
| ------------------- | ---------------------- | ----------------------------------- | --------------------------------------------------- |
| `auth_db`           | Auth Service           | User authentication and sessions    | [auth_db.sql](init/auth_db.sql)                     |
| `profiles_db`       | Profiles Service       | User profiles and role management   | [profiles_db.sql](init/profiles_db.sql)             |
| `text_db`           | Text Service           | Text content and version control    | [text_db.sql](init/text_db.sql)                     |
| `video_learning_db` | Video Learning Service | Video content and progress tracking | [video_learning_db.sql](init/video_learning_db.sql) |
| `knowledge_db`      | Knowledge Service      | Knowledge base and learning paths   | [knowledge_db.sql](init/knowledge_db.sql)           |
| `codelab_db`        | Codelab Service        | Interactive coding environments     | [codelab_db.sql](init/codelab_db.sql)               |

## 🔧 Configuration

### Environment Variables

Configure MinIO through environment files:

#### Development (.env.dev)

```bash
POSTGRES_DB=infrastructure_db
POSTGRES_USER=infrastructure_user
POSTGRES_PASSWORD=infrastructure_pass
POSTGRES_SSL_MODE=disable
POSTGRES_MAX_CONNECTIONS=150
POSTGRES_SHARED_BUFFERS=512MB
POSTGRES_EFFECTIVE_CACHE_SIZE=2GB
```

#### Production (.env.prod)

```bash
POSTGRES_DB=infrastructure_db
POSTGRES_USER=9fec08e4574bcb96d535c9616c
POSTGRES_PASSWORD=fa9dde3f017498917f2febdb0a
POSTGRES_SSL_MODE=disable
POSTGRES_MAX_CONNECTIONS=175
POSTGRES_SHARED_BUFFERS=512MB
POSTGRES_EFFECTIVE_CACHE_SIZE=2GB
```

## 🗂️ Schema Documentation

### Auth Database (`auth_db`)

Contains authentication-related tables:

- `users`: User accounts and credentials
- `sessions`: Active user sessions
- `backup_codes`: Two-factor authentication backup codes
- `auth_attempts`: Login attempt tracking

### Profiles Database (`profiles_db`)

Contains user profile information:

- `profiles`: Basic user profile data
- `student_profiles`: Student-specific information
- `teacher_profiles`: Teacher-specific information
- `profile_activities`: Profile activity logs

### Text Database (`text_db`)

Contains text content management:

- `content`: Text articles, lessons, assignments
- `content_versions`: Version history
- `categories`: Content categorization
- `content_collaborators`: Collaboration permissions
- `content_comments`: Review and feedback system

### Video Learning Database (`video_learning_db`)

Contains video-related data:

- `videos`: Video metadata and information
- `video_files`: Multiple quality/format files
- `video_progress`: User viewing progress
- `video_annotations`: Interactive annotations
- `video_playlists`: Curated video collections

### Knowledge Database (`knowledge_db`)

Contains knowledge management:

- `knowledge_articles`: Knowledge base articles
- `knowledge_categories`: Article categorization
- `knowledge_nodes`: Knowledge graph nodes
- `knowledge_relationships`: Node relationships
- `learning_paths`: Structured learning sequences

### Codelab Database (`codelab_db`)

Contains coding environment data:

- `coding_labs`: Programming exercises
- `code_submissions`: Student submissions
- `code_executions`: Execution results
- `collaboration_sessions`: Real-time coding sessions
- `code_templates`: Reusable code templates
