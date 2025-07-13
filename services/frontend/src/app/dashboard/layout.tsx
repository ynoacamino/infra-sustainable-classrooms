import Header from '@/layout/shared/header';
import { authService } from '@/services/auth/service';
import { profilesService } from '@/services/profiles/service';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export default async function LayoutStudent({
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
  return (
    <>
      <Header profile={res.data} />
      <main className="flex flex-col pb-20 items-center w-full">
        <div className='w-full max-w-6xl px-4 sm:px-6 lg:px-8'>
          {children}
        </div>
      </main>
    </>
  );
}
