# Frontend Service File Organization

This document explains the file organization structure for the frontend service of the Sustainable Classrooms platform.

## Directory Structure

```
src/
├── app/                   # Next.js App Router pages and API routes
├── components/            # Feature-based React components
├── services/              # API service layer
├── lib/                   # Utility functions and shared logic
├── types/                 # TypeScript type definitions
├── hooks/                 # Custom React hooks
├── layout/                # Layout components
├── ui/                    # Reusable UI components (Design System)
├── config/                # Configuration files
└── actions/               # Server actions (if any)
```

## Detailed Directory Breakdown

### `/app` - Next.js App Router

Contains pages, layouts, and API routes following Next.js 13+ App Router conventions.

### `/components` - Feature Components

Organized by feature/domain for better maintainability and scalability.

```
components/
├── auth/                  # Authentication components
│   └── forms/
├── courses/               # Course-related components
│   ├── admin/            # Admin-specific course components
│   └── student/          # Student-specific course components
└── post/                 # Post/content components
    ├── admin/
    └── student/
```

**Conventions:**
- Components are organized by business domain
- Each feature has its own directory
- Role-based subdirectories (student, teacher)
- Component file names use kebab case with `.tsx` extension, but the component name is in PascalCase

### `/services` - API Service Layer

Contains service classes for API communication.

```
services/
└── auth/
    └── auth.ts           # Authentication service
```

**Conventions:**
- Service files export an instance
- Use class-based services for stateful operations
- Async methods return consistent response formats

### `/lib` - Utility Functions and Shared Logic

Shared utilities, form schemas, constants, and helper functions.

```
lib/
├── auth/
│   ├── enums/           # Authentication enums
│   └── forms/           # Form schemas and validation
├── courses/
│   ├── admin/
│   ├── enums/
│   └── mock/            # Mock data for development
├── post/
├── shared/
│   ├── enums/
│   ├── time.ts          # Time utilities
│   └── utils.ts         # General utilities
```

**Conventions:**
- Domain-specific utilities in their own directories
- Shared utilities in the `shared/` directory
- Enums in dedicated `enums/` subdirectories
- Mock data in `mock/` subdirectories

### `/types` - TypeScript Types

Centralized type definitions organized by domain.

```
types/
├── auth/
│   ├── role.ts
│   └── user.ts
├── courses/
│   ├── courses.ts
│   └── grades.ts
├── post/
│   └── post.ts
└── shared/
    ├── field.ts
    └── select-option.ts
```

**Conventions:**
- Types are organized by business domain
- Shared types in the `shared/` directory
- Interface names use PascalCase
- Type names use PascalCase with `Type` suffix when needed

### `/hooks` - Custom React Hooks

Custom hooks for data fetching and state management.

```
hooks/
├── courses/
│   ├── use-courses.ts
│   └── use-grades.ts
└── post/
    ├── use-posts.ts
    └── use-view-or-edit.ts
```

**Conventions:**
- Hooks are organized by feature domain
- Hook names start with `use`
- Data fetching hooks typically use SWR
- Return consistent object structures

### `/layout` - Layout Components

Reusable layout components for different parts of the application.

```
layout/
├── auth/
│   ├── header-auth.tsx
│   └── header-login.tsx
└── shared/
    ├── header.tsx
    ├── layout-login.tsx
    └── layout-student.tsx
```

**Conventions:**
- Layout components are organized by context
- Shared layouts in the `shared/` directory
- Layout names describe their purpose clearly

### `/ui` - Design System Components

Reusable UI components that form the design system.

```
ui/
├── button.tsx
├── form.tsx
├── input.tsx
├── select.tsx
├── table.tsx
├── skeleton.tsx
├── popover.tsx
└── ...
```

**Conventions:**
- Low-level, reusable components
- No business logic

### `/config` - Configuration Files

Application configuration and constants.

```
config/
└── shared/
    └── routes.ts         # Route configuration
```

**Conventions:**
- Configuration organized by scope
- Export const objects for constants
- Environment-specific configurations when needed

## File Naming Conventions

### Components
- **PascalCase** for component files: `LoginForm.tsx`
- **kebab-case** for directories: `auth-forms/`
- **camelCase** for functions and variables

### Non-Components
- **kebab-case** for utility files: `use-courses.ts`
- **camelCase** for functions and variables
- **PascalCase** for types and interfaces
- **SCREAMING_SNAKE_CASE** for constants

## Import/Export Patterns

### Named Exports (Preferred)
```typescript
// Good
export function LoginForm() { ... }
export { LoginForm };

// Usage
import { LoginForm } from '@/components/auth/forms/login-form';
```

### Default Exports (Limited Use)
```typescript
// Only for page components and some layout components
export default function LoginPage() { ... }
```

## Path Aliases

The project uses TypeScript path aliases for cleaner imports:

```typescript
// tsconfig.json paths
{
  "@/*": ["./src/*"]
}

// Usage
import { Button } from '@/ui/button';
import { authService } from '@/services/auth/auth';
import type { User } from '@/types/auth/user';
```

## Best Practices

1. **Keep components focused**: Each component should have a single responsibility
2. **Use TypeScript**: All files should be properly typed
3. **Consistent naming**: Follow the established naming conventions
4. **Organize imports**: Group imports logically (external, internal, types)
5. **Documentation**: Add JSDoc comments for complex functions
6. **Testing**: Co-locate test files with their components when applicable

## Adding New Features

When adding a new feature:

1. Create a new directory in `/components` for the feature
2. Add corresponding types in `/types`
3. Create hooks in `/hooks` if needed
4. Add service methods in `/services`
5. Update route configuration in `/config` if needed
6. Follow the established patterns and conventions

This structure ensures maintainability, scalability, and developer productivity while keeping the codebase organized and easy to navigate.
