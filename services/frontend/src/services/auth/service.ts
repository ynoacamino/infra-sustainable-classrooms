import { AuthInterceptor } from '@/services/auth/interceptor';
import type { SimpleResponse } from '@/services/shared/response';
import { Service } from '@/services/shared/service';
import type { TOTPSecret, User } from '@/types/auth/models';
import type {
  GenerateSecretPayload,
  VerifyBackupCodePayload,
  VerifyTOTPPayload,
} from '@/types/auth/payload';
import type { AuthResponse, BackupCodeResponse } from '@/types/auth/responses';
import {
  GenerateSecretPayloadSchema,
  VerifyBackupCodePayloadSchema,
  VerifyTOTPPayloadSchema,
} from '@/types/auth/schemas/payload';
import type { AsyncResult } from '@/types/shared/services/result';
import type { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';

class AuthService extends Service {
  constructor() {
    super('auth');
  }

  async generateSecret(
    payload: GenerateSecretPayload,
  ): AsyncResult<TOTPSecret> {
    return await this.post<TOTPSecret>({
      endpoint: ['totp', 'generate'],
      payload: {
        schema: GenerateSecretPayloadSchema,
        data: payload,
      },
    });
  }

  async verifyTOTP(payload: VerifyTOTPPayload): AsyncResult<AuthResponse> {
    return await this.post<AuthResponse>({
      endpoint: ['totp', 'verify'],
      payload: {
        schema: VerifyTOTPPayloadSchema,
        data: payload,
      },
    });
  }

  async verifyBackupCode(
    payload: VerifyBackupCodePayload,
  ): AsyncResult<BackupCodeResponse> {
    return await this.post<BackupCodeResponse>({
      endpoint: ['backup', 'verify'],
      payload: {
        schema: VerifyBackupCodePayloadSchema,
        data: payload,
      },
    });
  }

  async refreshSession(): AsyncResult<SimpleResponse> {
    return await this.post<SimpleResponse>({
      endpoint: ['session', 'refresh'],
    });
  }

  async logout(): AsyncResult<SimpleResponse> {
    return await this.post<SimpleResponse>({
      endpoint: 'logout',
    });
  }

  async getUserProfile(): AsyncResult<User> {
    return await this.get<User>({
      endpoint: 'profile',
    });
  }
}

// Factory function to create an instance of AuthService with the interceptor
export const authService = async (cookies: Promise<ReadonlyRequestCookies>) => {
  const service = new AuthService();
  service.addInterceptor(new AuthInterceptor(await cookies));
  return service;
};
