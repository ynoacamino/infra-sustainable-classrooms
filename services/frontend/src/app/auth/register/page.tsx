import { redirect } from 'next/navigation';
import Link from 'next/link';
import LayoutLogin from '@/layout/shared/layout-login';
import { RegisterForm } from '@/components/auth/forms/register-form';
import { authService } from '@/services/auth/service';
import { cookies } from 'next/headers';

export default async function RegisterPage() {
  const auth = await authService(cookies());
  const res = await auth.getUserProfile();
  if (res.success) {
    redirect('/dashboard');
  }

  return (
    <LayoutLogin>
      <h1 className="font-bold text-3xl">Create your account</h1>
      <div className="flex flex-col gap-y-2 items-center w-full">
        <RegisterForm />
        <p className="mt-4 text-sm text-gray-600">
          Already have an account?{' '}
          <Link href="/auth/verify" className="text-blue-500 hover:underline">
            Sign in
          </Link>
        </p>
      </div>
    </LayoutLogin>
  );
}
