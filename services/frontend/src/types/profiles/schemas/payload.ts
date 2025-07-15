import {
  ProfileSchema,
  StudentProfileSchema,
  TeacherProfileSchema,
} from '@/types/profiles/schemas/models';
import z from 'zod';

export const CreateStudentProfilePayloadSchema = StudentProfileSchema.omit({
  // This is a cookie
  // session: z.string(),
  role: true,
  user_id: true,
  created_at: true,
  updated_at: true,
  is_active: true,
});

export const CreateTeacherProfilePayloadSchema = TeacherProfileSchema.omit({
  // This is a cookie
  // session: z.string(),
  role: true,
  user_id: true,
  created_at: true,
  updated_at: true,
  is_active: true,
});

export const GetPublicProfileByIdPayloadSchema = ProfileSchema.pick({
  user_id: true,
});

const UpdateStudentSchema = StudentProfileSchema.omit({
  role: true,
  user_id: true,
  created_at: true,
  updated_at: true,
  is_active: true,
  major: true, // Is it a bug? On the server, this is not required
  // Is it a bug? On the server, this is not required
  grade_level: true,
}).partial();
const UpdateTeacherSchema = TeacherProfileSchema.omit({
  role: true,
  user_id: true,
  created_at: true,
  updated_at: true,
  is_active: true,
}).partial();

export const UpdateProfilePayloadSchema = z.union([
  UpdateStudentSchema,
  UpdateTeacherSchema,
]);
