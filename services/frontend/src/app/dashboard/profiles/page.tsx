import { redirect } from 'next/navigation';

export default async function ProfilesPage() {
  // Por defecto redirigir a ver el perfil
  redirect('/dashboard/profiles/view');
}
