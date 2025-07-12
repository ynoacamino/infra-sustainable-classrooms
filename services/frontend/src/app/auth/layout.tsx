import HeaderLogin from '@/layout/auth/header-login';

export default function LayoutLogin({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <>
      <HeaderLogin />
      <main className="flex flex-col items-center justify-center gap-y-12 flex-1 px-4">
        {children}
      </main>
    </>
  );
}
