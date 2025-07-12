'use client';

import { ChevronDown, LogOut } from 'lucide-react';
import { Popover, PopoverContent, PopoverTrigger } from '@/ui/popover';
import { redirect, RedirectType } from 'next/navigation';
import { Button } from '@/ui/button';
import { logoutAction } from '@/actions/auth/actions';
import { toast } from 'sonner';
import type { Profile } from '@/types/profiles/models';

interface HeaderAuthProps {
  profile: Profile;
}

export function HeaderAuth({ profile }: HeaderAuthProps) {
  // TODO: Logout may be a server action
  const handleLogout = async () => {
    const res = await logoutAction();
    if (!res.success) {
      toast.error(res.error.message);
      return;
    }
    toast.success('Logout successful');
    redirect('/dashboard', RedirectType.push);
  };
  return (
    <Popover>
      <PopoverTrigger asChild>
        <div className="flex gap-3 items-center hover:cursor-pointer">
          {profile.first_name}
          <ChevronDown className="w-5 h-5" />
        </div>
      </PopoverTrigger>
      <PopoverContent className="w-72">
        <div className="flex flex-col gap-3 items-center pt-5">
          <span className="text-lg font-medium text-center">{profile.em}</span>
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
