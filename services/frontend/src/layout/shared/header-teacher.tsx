import { IconWifi } from '@tabler/icons-react';
import { routesConfig } from '@/config/shared/routes';
import { HeaderAuth } from '@/layout/auth/header-auth';
import { Link } from '@/ui/link';
import { ThemeToggle } from '@/ui/theme-toggle';
import type { Profile } from '@/types/profiles/models';

interface HeaderProps {
  profile: Profile;
}

export default function HeaderTeacher({ profile }: HeaderProps) {
  return (
    <header className="flex flex-col h-screen justify-between sticky top-0 z-50 bg-background py-2 md:py-4 px-4 sm:px-6 w-full max-w-54">
      <div className="flex flex-col items-center gap-2">
        <Link href="/" className="gap-2 font-bold text-lg flex w-full">
          <IconWifi size={24} />
          <span>StudyCentral</span>
        </Link>
        <Link href="/" className="gap-2 font-bold text-lg flex w-full">
          <span>Teacher View</span>
        </Link>
      </div>
      <nav className="text-sm flex flex-col gap-4 font-medium items-center">
        <div className="min-w-30 flex justify-center items-center mb-4">
          <HeaderAuth profile={profile} />
        </div>
        <div className="flex justify-center mb-2">
          <ThemeToggle />
        </div>
        {routesConfig.teacherRoutes.map((route) => (
          <Link
            key={route.path}
            href={route.path}
            className="transition-colors duration-75 w-40"
            aria-label={route.name}
          >
            <route.icon className="size-6" />
            {route.name}
          </Link>
        ))}
      </nav>
    </header>
  );
}
