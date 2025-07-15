import Link from 'next/link';
import { redirect } from 'next/navigation';
import { authService } from '@/services/auth/service';
import { cookies } from 'next/headers';
import { BackupForm } from '@/components/auth/forms/backup-form';

export default async function BackupPage() {
  const auth = await authService(cookies());
  const res = await auth.getUserProfile();
  if (res.success) {
    redirect('/dashboard');
  }

  return (
    <>
      <h1 className="font-bold text-3xl">Use Backup Code</h1>
      <p className="text-gray-600 text-center max-w-md">
        Enter one of your backup codes to sign in. Each backup code can only be
        used once.
      </p>
      <div className="flex flex-col gap-y-2 items-center w-full">
        <BackupForm />
        <p className="mt-4 text-sm text-gray-600">
          Have access to your authenticator app?{' '}
          <Link href="/auth/verify" className="text-blue-500 hover:underline">
            Use TOTP code instead
          </Link>
        </p>
        <p className="text-sm text-gray-600">
          Don&apos;t have an account?{' '}
          <Link href="/auth/register" className="text-blue-500 hover:underline">
            Sign up
          </Link>
        </p>
      </div>
    </>
  );
}
