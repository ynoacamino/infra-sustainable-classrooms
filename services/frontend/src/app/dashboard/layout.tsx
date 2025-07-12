import Header from '@/layout/shared/header';
import { profilesService } from '@/services/profiles/service';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export default async function LayoutStudent({
  children,
}: {
  children: React.ReactNode;
}) {
  const profiles = await profilesService(cookies());
  const res = await profiles.getCompleteProfile();
  if (!res.success) {
    redirect('/auth/verify');
  }
  return (
    <>
      <Header profile={res.data} />
      <main className="flex flex-col pb-20 items-center">{children}</main>
    </>
  );
}
