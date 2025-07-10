'use client';

import { ReactNode } from 'react';
import Header from '@/modules/core/layouts/Header';
import type { User } from '@/modules/auth/types/user';

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
