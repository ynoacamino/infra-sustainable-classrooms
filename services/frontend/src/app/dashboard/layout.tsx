import Header from '@/layout/shared/header';
import { profilesService } from '@/services/profiles/service';
import { cookies } from 'next/headers';

export default async function LayoutStudent({
  children,
}: {
  children: React.ReactNode;
}) {
  const profiles = await profilesService(cookies());
  const res = await profiles.getCompleteProfile();
  console.log(res);
  if (!res.success) {
    return <h1>Create profile here</h1>;
  }
  return (
    <>
      <Header profile={res.data} />
      <main className="flex flex-col pb-20 items-center">{children}</main>
    </>
  );
}
