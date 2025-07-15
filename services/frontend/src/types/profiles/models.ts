import type {
  CompleteProfileSchema,
  ProfileSchema,
  PublicProfileSchema,
  StudentProfileSchema,
  TeacherProfileSchema,
} from '@/types/profiles/schemas/models';
import type z from 'zod';

export type Profile = z.infer<typeof ProfileSchema>;

export type PublicProfile = z.infer<typeof PublicProfileSchema>;

export type StudentProfile = z.infer<typeof StudentProfileSchema>;

export type TeacherProfile = z.infer<typeof TeacherProfileSchema>;

export type CompleteProfile = z.infer<typeof CompleteProfileSchema>;
