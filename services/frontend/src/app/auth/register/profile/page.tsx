import { redirect } from 'next/navigation';
import Link from 'next/link';
import LayoutLogin from '@/layout/shared/layout-login';
import { authService } from '@/services/auth/service';
import { cookies } from 'next/headers';
import { profilesService } from '@/services/profiles/service';
import { GraduationCap, User } from 'lucide-react';

export default async function RegisterProfilePage() {
  const auth = await authService(cookies());
  const authRes = await auth.getUserProfile();
  const profiles = await profilesService(cookies());
  const profileRes = await profiles.getCompleteProfile();

  if (!authRes.success) {
    redirect('/auth/verify');
  }

  if (profileRes.success) {
    redirect('/dashboard');
  }

  return (
    <LayoutLogin>
      <h1 className="font-bold text-3xl">Create your profile</h1>
      <p>Select teacher or student to create your profile</p>
      <div className="flex flex-col gap-y-2 items-center w-full">
        <div className="flex items-center justify-between w-full">
          <Link
            href="/auth/register/profile/teacher"
            className="flex flex-col gap-2 items-center mt-2"
          >
            <GraduationCap className="size-16 stroke-1" />
            Register as Teacher
          </Link>
          <Link
            href="/auth/register/profile/student"
            className="flex flex-col gap-2 items-center mt-2"
          >
            <User className="size-16 stroke-1" />
            Register as Student
          </Link>
        </div>
      </div>
    </LayoutLogin>
  );
}
