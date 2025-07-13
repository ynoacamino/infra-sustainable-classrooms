'use client';

import { ChevronDown, LogOut, UserPen } from 'lucide-react';
import { Popover, PopoverContent, PopoverTrigger } from '@/ui/popover';
import { redirect, RedirectType } from 'next/navigation';
import { Button } from '@/ui/button';
import { logoutAction } from '@/actions/auth/actions';
import { toast } from 'sonner';
import type { Profile } from '@/types/profiles/models';
import Link from 'next/link';

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
          <img
            src={profile.avatar_url}
            alt="Profile Avatar"
            className="w-24 h-24 rounded-full object-cover mb-2"
          />
          <span className="text-sm text-foreground/70">{profile.email}</span>
          <span className="text-lg font-medium text-center">
            {profile.first_name}
          </span>
          <div className="flex flex-col items-center justify-start w-full gap-1 text-secondary">
            <Button variant="ghost" asChild>
              <Link
                href="/dashboard/profiles/update"
                className="text-foreground w-full flex"
              >
                <UserPen />
                <span className="flex-1 w-full text-center">
                  Update profile
                </span>
              </Link>
            </Button>
            <Button
              variant="ghost"
              className="text-foreground w-full flex"
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
