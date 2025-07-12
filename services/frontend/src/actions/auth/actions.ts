'use server';

import { authService } from '@/services/auth/service';
import type {
  GenerateSecretPayload,
  VerifyTOTPPayload,
} from '@/types/auth/payload';
import { cookies } from 'next/headers';

export async function generateSecretAction(payload: GenerateSecretPayload) {
  const auth = await authService(cookies());
  const res = await auth.generateSecret(payload);
  if (!res.success) {
    return res;
  }
  const { totp_url, backup_codes, issuer } = res.data;
  (await cookies()).set('totp_url', totp_url, { httpOnly: true });
  (await cookies()).set('backup_codes', JSON.stringify(backup_codes), {
    httpOnly: true,
  });
  (await cookies()).set('issuer', issuer, { httpOnly: true });
  return res;
}

export async function verifyTotpAction(payload: VerifyTOTPPayload) {
  const auth = await authService(cookies());
  return await auth.verifyTOTP(payload);
}

export async function logoutAction() {
  const auth = await authService(cookies());
  const res = await auth.logout();
  if (res.success) {
    (await cookies()).delete('session');
  }
  return res;
}
