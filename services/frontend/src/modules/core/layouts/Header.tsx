'use client';

import { IconWifi, IconBell } from '@tabler/icons-react';
import Link from 'next/link';
import { HeaderAuth } from '@/modules/auth/components/headerAuth';
import { routesConfig } from '@/modules/core/config/routes';
import type { User } from '@/modules/auth/types/user';

interface HeaderProps {
  user: User;
}

export default function Header({ user }: HeaderProps) {
  return (
    <header className="flex justify-between sticky top-0 z-50 bg-background py-2 md:py-4 px-4 sm:px-6 md:px-8">
      <Link href="/" className="gap-2 font-bold text-lg flex">
        <IconWifi size={24} />
        <span>StudyCentral</span>
      </Link>
      <nav className="text-sm flex gap-9 font-medium items-center">
        {routesConfig.studentRoutes.map((route) => (
          <Link
            key={route.path}
            href={route.path}
            className="hover:text-primary transition-colors duration-75"
            aria-label={route.name}
          >
            {route.name}
          </Link>
        ))}
        <button className="rounded-full bg-zinc-300 w-9 aspect-square flex items-center justify-center hover:bg-zinc-400/50 transition-colors">
          <IconBell size={18} stroke={2} />
        </button>
        <div className="min-w-30">
          <HeaderAuth user={user} />
        </div>
      </nav>
    </header>
  );
}
