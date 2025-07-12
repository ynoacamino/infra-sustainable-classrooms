import Link from 'next/link';
import { redirect } from 'next/navigation';
import { authService } from '@/services/auth/service';
import { cookies } from 'next/headers';
import { VerifyForm } from '@/components/auth/forms/verify-form';

export default async function LoginPage() {
  const auth = await authService(cookies());
  const res = await auth.getUserProfile();
  if (res.success) {
    redirect('/dashboard');
  }

  return (
    <>
      <h1 className="font-bold text-3xl">Welcome back</h1>
      <div className="flex flex-col gap-y-2 items-center w-full">
        <VerifyForm />
        <p className="mt-4 text-sm text-gray-600">
          Don&apos;t have an account?{' '}
          <Link href="/auth/register" className="text-blue-500 hover:underline">
            Sign up
          </Link>
        </p>
        <p className="text-sm text-gray-600">
          Don&apos;t have access to your code?{' '}
          <Link href="/auth/backup" className="text-blue-500 hover:underline">
            Use backup code
          </Link>
        </p>
      </div>
    </>
  );
}
