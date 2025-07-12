import { TotpSetupComponent } from '@/components/auth/TotpSetupComponent';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export default async function SaveTotpPage() {
  // Get the TOTP secret and backup codes from cookies
  const totpUrl = (await cookies()).get('totp_url');
  const backupCodes = (await cookies()).get('backup_codes');
  const issuer = (await cookies()).get('issuer');

  if (!totpUrl || !backupCodes || !issuer) {
    redirect('/auth/register');
  }

  const parsedBackupCodes = backupCodes.value
    ? JSON.parse(backupCodes.value)
    : [];

  if (!Array.isArray(parsedBackupCodes) || parsedBackupCodes.length === 0) {
    redirect('/auth/register');
  }

  // delete cookies after use
  (await cookies()).delete('totp_url');
  (await cookies()).delete('backup_codes');
  (await cookies()).delete('issuer');

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Set up Two-Factor Authentication
        </h2>
        <p className="mt-2 text-center text-sm text-gray-600">
          Complete your account setup by configuring 2FA
        </p>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-2xl">
        <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
          <TotpSetupComponent
            totpUrl={totpUrl.value}
            backupCodes={parsedBackupCodes}
            issuer={issuer.value}
          />
        </div>
      </div>
    </div>
  );
}
