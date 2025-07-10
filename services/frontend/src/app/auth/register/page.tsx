'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import LayoutLogin from '@/layout/shared/layout-login';
import { authService } from '@/services/auth/auth';
import { RegisterForm } from '@/components/auth/forms/register-form';

export default function RegisterPage() {
  const router = useRouter();

  useEffect(() => {
    // Check if user is already logged in
    const checkAuth = async () => {
      const user = await authService.getUser();
      if (user) {
        router.replace('/dashboard');
      }
    };
    checkAuth();
  }, [router]);

  return (
    <LayoutLogin>
      <h1 className="font-bold text-3xl">Create your account</h1>
      <div className="flex flex-col gap-y-2 items-center w-full">
        <RegisterForm />
        <p className="mt-4 text-sm text-gray-600">
          Already have an account?{' '}
          <Link href="/auth/login" className="text-blue-500 hover:underline">
            Sign in
          </Link>
        </p>
      </div>
    </LayoutLogin>
  );
}
