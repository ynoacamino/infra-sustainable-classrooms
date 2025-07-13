'use server';

import { profilesService } from "@/services/profiles/service";
import type { CreateStudentProfilePayloadSchema, CreateTeacherProfilePayloadSchema, UpdateProfilePayloadSchema } from "@/types/profiles/schemas/payload";
import { cookies } from 'next/headers';
import type z from "zod";

export async function createTeacherProfileAction(payload: z.infer<typeof CreateTeacherProfilePayloadSchema>) {
  const profiles = await profilesService(cookies());
  return  profiles.createTeacherProfile(payload);
}

export async function createStudentProfileAction(payload: z.infer<typeof CreateStudentProfilePayloadSchema>) {
  const profiles = await profilesService(cookies());
  return profiles.createStudentProfile(payload);
}

export async function updateProfileAction(payload: z.infer<typeof UpdateProfilePayloadSchema>) {
  const profiles = await profilesService(cookies());
  return profiles.updateProfile(payload);
}