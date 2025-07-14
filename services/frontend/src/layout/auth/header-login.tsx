import { ThemeToggle } from '@/ui/theme-toggle';
import { IconWifi } from '@tabler/icons-react';
import Link from 'next/link';

export default function HeaderLogin() {
  return (
    <header className="flex justify-between sticky top-0 z-50 bg-background py-2 md:py-4 px-4 sm:px-6 md:px-8">
      <Link href="/auth/login" className="gap-2 font-bold text-lg flex">
        <IconWifi size={24} />
        <span>StudyCentral</span>
      </Link>
      <ThemeToggle />
    </header>
  );
}
