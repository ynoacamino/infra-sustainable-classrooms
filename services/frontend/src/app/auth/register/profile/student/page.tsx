import { redirect } from 'next/navigation';
import LayoutLogin from '@/layout/shared/layout-login';
import { authService } from '@/services/auth/service';
import { cookies } from 'next/headers';
import { profilesService } from '@/services/profiles/service';
import { RegisterStudentForm } from '@/components/profiles/forms/register-student-form';

export default async function RegisterStudentPage() {
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
      <h1 className="font-bold text-3xl">Create Student Profile</h1>
      <div className="flex flex-col gap-y-2 items-center w-full">
        <RegisterStudentForm />
      </div>
    </LayoutLogin>
  );
}
