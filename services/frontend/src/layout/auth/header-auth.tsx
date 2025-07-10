import type { User } from '@/types/auth/user';
import { ChevronDown, LogOut } from 'lucide-react';
import { Popover, PopoverContent, PopoverTrigger } from '@/ui/popover';
import { useRouter } from 'next/navigation';
import { authService } from '@/services/auth/auth';
import Image from 'next/image';
import { Button } from '@/ui/button';

interface HeaderAuthProps {
  user: User;
}

export function HeaderAuth({ user }: HeaderAuthProps) {
  const router = useRouter();

  const handleLogout = async () => {
    try {
      await authService.logout();
      router.push('/auth/login');
    } catch (error) {
      console.error('Error logging out:', error);
    }
  };
  return (
    <Popover>
      <PopoverTrigger asChild>
        <div className="flex gap-3 items-center hover:cursor-pointer">
          <Image
            alt="name"
            src={user.photo}
            className="rounded-full w-9 aspect-square"
          />
          <ChevronDown className="w-5 h-5" />
        </div>
      </PopoverTrigger>
      <PopoverContent className="w-72">
        <div className="flex flex-col gap-3 items-center pt-5">
          <Image
            alt="name"
            src={user.photo}
            className="rounded-full aspect-square bg-secondary w-20"
          />
          <span className="text-lg font-medium text-center">
            {user.name} ({user.email})
          </span>
          <div className="flex flex-col items-center justify-start w-full gap-1 text-secondary">
            <Button
              variant="ghost"
              className="text-base w-full flex"
              onClick={handleLogout}
            >
              <LogOut />
              <span className="flex-1 w-full">Cerrar Sesion</span>
            </Button>
          </div>
        </div>
      </PopoverContent>
    </Popover>
  );
}
