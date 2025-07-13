import HeaderTeacher from '@/layout/shared/header-teacher';
import { authService } from '@/services/auth/service';
import { profilesService } from '@/services/profiles/service';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export default async function LayoutTeacher({
  children,
}: {
  children: React.ReactNode;
}) {
  const profiles = await profilesService(cookies());
  const auth = await authService(cookies());
  const res = await profiles.getCompleteProfile();

  if (!res.success) {
    const res = await auth.getUserProfile();
    if (!res.success) {
      redirect('/auth/verify');
    } else {
      redirect('/auth/register/profile');
    }
  }

  const isTeacher = profiles.isTeacherProfile(res.data);
  if (!isTeacher) {
    redirect('/dashboard');
  }

  return (
    <div className="w-full flex justify-start">
      <HeaderTeacher profile={res.data} />
      <main className="flex flex-1 flex-col pb-20 justify-start min-h-screen">
        {children}
      </main>
    </div>
  );
}
