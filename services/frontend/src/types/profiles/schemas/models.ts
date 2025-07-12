import z from 'zod';

export const ProfileSchema = z.object({
  user_id: z
    .number()
    .int('User ID must be an integer')
    .describe('User identifier'),
  role: z
    .enum(['student', 'teacher'], {
      message: 'Role must be either student or teacher',
    })
    .describe('User role (student, teacher)'),
  first_name: z
    .string()
    .min(1, 'First name is required')
    .describe('First name'),
  last_name: z.string().min(1, 'Last name is required').describe('Last name'),
  email: z.string().email('Invalid email address').describe('Email address'),
  phone: z.string().optional().describe('Phone number'),
  avatar_url: z
    .string()
    .url('Invalid URL')
    .optional()
    .describe('Profile picture URL'),
  bio: z
    .string()
    .max(500, 'Bio cannot exceed 500 characters')
    .optional()
    .describe('Biography/description'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe('Profile creation timestamp'),
  updated_at: z
    .number()
    .int('Last update timestamp must be an integer')
    .describe('Last update timestamp'),
  is_active: z.boolean().describe('Whether profile is active'),
});

export const PublicProfileSchema = ProfileSchema.pick({
  user_id: true,
  role: true,
  first_name: true,
  last_name: true,
  avatar_url: true,
  bio: true,
  is_active: true,
});

export const StudentProfileSchema = ProfileSchema.extend({
  role: z.literal('student'),
  grade_level: z.string().describe('Student grade level'),
  major: z.string().optional().describe('Student major or field of study'),
});

export const TeacherProfileSchema = ProfileSchema.extend({
  role: z.literal('teacher'),
  position: z.string().describe('Teacher position or title'),
});

export const CompleteProfileSchema = z.discriminatedUnion('role', [
  StudentProfileSchema.extend({ role: z.literal('student') }),
  TeacherProfileSchema.extend({ role: z.literal('teacher') }),
]);
