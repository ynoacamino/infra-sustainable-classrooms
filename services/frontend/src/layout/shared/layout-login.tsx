import HeaderLogin from '@/layout/auth/header-login';
import type { ReactNode } from 'react';

interface LayoutLoginProps {
  children: ReactNode;
}

export default function LayoutLogin({ children }: LayoutLoginProps) {
  return (
    <>
      <HeaderLogin />
      <main className="flex flex-col items-center justify-center gap-y-12 flex-1 px-4">
        {children}
      </main>
    </>
  );
}
