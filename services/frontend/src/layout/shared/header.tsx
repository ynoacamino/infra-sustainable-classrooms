import { IconWifi, IconBell } from '@tabler/icons-react';
import { routesConfig } from '@/config/shared/routes';
import { HeaderAuth } from '@/layout/auth/header-auth';
import { Link } from '@/ui/link';
import { Button } from '@/ui/button';
import { ThemeToggle } from '@/ui/theme-toggle';
import type { Profile } from '@/types/profiles/models';

interface HeaderProps {
  profile: Profile;
}

export default function Header({ profile }: HeaderProps) {
  return (
    <header className="flex justify-between sticky top-0 z-50 bg-background py-2 md:py-4 px-4 sm:px-6 md:px-8">
      <Link href="/" className="gap-2 font-bold text-lg flex" variant="ghost">
        <IconWifi size={24} />
        <span>StudyCentral</span>
      </Link>
      <nav className="text-sm flex gap-5 font-medium items-center">
        {profile.role === 'teacher' && (
          <Link
            href="/teacher"
            className="transition-colors duration-75"
            aria-label="Teacher Dashboard"
            variant="outline"
          >
            Teacher Dashboard
          </Link>
        )}
        {routesConfig.studentRoutes.map((route) => (
          <Link
            key={route.path}
            href={route.path}
            className="transition-colors duration-75"
            aria-label={route.name}
            variant="ghost"
          >
            {route.name}
          </Link>
        ))}
        <Button className="rounded-full" size="icon" variant="secondary">
          <IconBell />
        </Button>
        <ThemeToggle />
        <div className="min-w-30">
          <HeaderAuth profile={profile} />
        </div>
      </nav>
    </header>
  );
}
