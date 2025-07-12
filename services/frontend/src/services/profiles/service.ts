import { SessionInterceptor } from '@/services/auth/interceptor';
import { Service } from '@/services/shared/service';
import type {
  CompleteProfile,
  Profile,
  PublicProfile,
  StudentProfile,
  TeacherProfile,
} from '@/types/profiles/models';
import type {
  CreateStudentProfilePayload,
  CreateTeacherProfilePayload,
  GetPublicProfileByIdPayload,
  UpdateProfilePayload,
} from '@/types/profiles/payload';
import {
  CreateStudentProfilePayloadSchema,
  GetPublicProfileByIdPayloadSchema,
  UpdateProfilePayloadSchema,
} from '@/types/profiles/schemas/payload';
import type { AsyncResult } from '@/types/shared/services/result';
import type { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';

class ProfilesService extends Service {
  constructor() {
    super('profiles');
  }

  async createStudentProfile(
    payload: CreateStudentProfilePayload,
  ): AsyncResult<StudentProfile> {
    return this.post<StudentProfile>({
      endpoint: 'student',
      payload: {
        schema: CreateStudentProfilePayloadSchema,
        data: payload,
      },
    });
  }

  async createTeacherProfile(
    payload: CreateTeacherProfilePayload,
  ): AsyncResult<TeacherProfile> {
    return this.post<TeacherProfile>({
      endpoint: 'teacher',
      payload: {
        schema: CreateStudentProfilePayloadSchema,
        data: payload,
      },
    });
  }

  async getCompleteProfile(): AsyncResult<CompleteProfile> {
    return this.get<CompleteProfile>({
      endpoint: 'me',
    });
  }

  async getPublicProfileById(
    payload: GetPublicProfileByIdPayload,
  ): AsyncResult<PublicProfile> {
    return this.get<PublicProfile>({
      endpoint: ['public', payload.user_id],
      payload: {
        schema: GetPublicProfileByIdPayloadSchema,
        data: payload,
      },
    });
  }

  async updateProfile(payload: UpdateProfilePayload): AsyncResult<Profile> {
    return this.put<Profile>({
      endpoint: 'me',
      payload: {
        schema: UpdateProfilePayloadSchema,
        data: payload,
      },
    });
  }

  isStudentProfile(profile: CompleteProfile): profile is StudentProfile {
    return profile.role === 'student';
  }
  isTeacherProfile(profile: CompleteProfile): profile is TeacherProfile {
    return profile.role === 'teacher';
  }
}

export const profilesService = async (
  cookies: Promise<ReadonlyRequestCookies>,
) => {
  const service = new ProfilesService();
  service.addInterceptor(new SessionInterceptor(await cookies));
  return service;
};
