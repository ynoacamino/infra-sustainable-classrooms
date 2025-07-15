import { ProfileView } from '@/components/profiles/profile-view';
import { profilesService } from '@/services/profiles/service';
import { cookies } from 'next/headers';

export default async function ViewProfilePage() {
  const profiles = await profilesService(cookies());
  const profile = await profiles.getCompleteProfile();
  if (!profile.success) {
    return <div>Error loading profile: {profile.error.message}</div>;
  }

  return (
    <div className="flex flex-col items-center justify-center w-full h-full">
      <h1 className="text-2xl font-bold mb-4">My Profile</h1>
      <ProfileView profile={profile.data} />
    </div>
  );
}
