'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import LayoutLogin from '@/modules/core/layouts/LayoutLogin';
import { LoginForm } from '@/modules/auth/components/LoginForm';
import { authService } from '@/lib/auth';
import Link from 'next/link';

export default function LoginPage() {
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
      <h1 className="font-bold text-3xl">Welcome back</h1>
      <div className="flex flex-col gap-y-2 items-center w-full">
        <LoginForm />
        <p className="mt-4 text-sm text-gray-600">
          Don&apos;t have an account?{' '}
          <Link href="/auth/register" className="text-blue-500 hover:underline">
            Sign up
          </Link>
        </p>
      </div>
    </LayoutLogin>
  );
}
