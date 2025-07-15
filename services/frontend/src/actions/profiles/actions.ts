'use server';

import { profilesService } from '@/services/profiles/service';
import type {
  CreateStudentProfilePayload,
  CreateTeacherProfilePayload,
  UpdateProfilePayload,
} from '@/types/profiles/payload';
import { cookies } from 'next/headers';

export async function createTeacherProfileAction(
  payload: CreateTeacherProfilePayload,
) {
  const profiles = await profilesService(cookies());
  return profiles.createTeacherProfile(payload);
}

export async function createStudentProfileAction(
  payload: CreateStudentProfilePayload,
) {
  const profiles = await profilesService(cookies());
  return profiles.createStudentProfile(payload);
}

export async function updateProfileAction(payload: UpdateProfilePayload) {
  const profiles = await profilesService(cookies());
  return profiles.updateProfile(payload);
}
