import type {
  CreateStudentProfilePayloadSchema,
  CreateTeacherProfilePayloadSchema,
  GetPublicProfileByIdPayloadSchema,
  UpdateProfileSchema,
} from '@/types/profiles/schemas/payload';
import type z from 'zod';

export type CreateStudentProfilePayload = z.infer<
  typeof CreateStudentProfilePayloadSchema
>;

export type CreateTeacherProfilePayload = z.infer<
  typeof CreateTeacherProfilePayloadSchema
>;

export type GetPublicProfileByIdPayload = z.infer<
  typeof GetPublicProfileByIdPayloadSchema
>;

export type UpdateProfilePayload = z.infer<typeof UpdateProfileSchema>;
