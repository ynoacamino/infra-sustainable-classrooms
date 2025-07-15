import { redirect } from 'next/navigation';

export default async function CodeLabDashboard() {
  redirect('/dashboard/codelab/exercises');
}
