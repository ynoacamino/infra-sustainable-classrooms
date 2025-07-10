'use client';

import Header from '@/layout/shared/header';
import type { User } from '@/types/auth/user';
import { type ReactNode } from 'react';

interface LayoutStudentProps {
  user: User;
  children: ReactNode;
}

export default function LayoutStudent({ user, children }: LayoutStudentProps) {
  return (
    <>
      <Header user={user} />
      <main className="flex flex-col pb-20 items-center">{children}</main>
    </>
  );
}
