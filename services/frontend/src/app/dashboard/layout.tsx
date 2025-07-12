import Header from '@/layout/shared/header';
import { authService } from '@/services/auth/service';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export default async function LayoutStudent({
  children,
}: {
  children: React.ReactNode;
}) {
  const auth = await authService(cookies());
  const res = await auth.getUserProfile();
  if (!res.success) {
    redirect('/auth/verify');
  }
  return (
    <>
      <Header user={res.data} />
      <main className="flex flex-col pb-20 items-center">{children}</main>
    </>
  );
}
