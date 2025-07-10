'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import LayoutStudent from '@/layout/shared/layout-student';
import H1 from '@/ui/h1';
import Section from '@/ui/section';
import { Courses } from '@/components/courses/student/courses-list';
import { Grades } from '@/components/courses/student/grades';
import { authService } from '@/services/auth/auth';
import type { User } from '@/types/auth/user';

export default function DashboardPage() {
  const router = useRouter();
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const checkAuth = async () => {
      const currentUser = await authService.getUser();
      if (!currentUser) {
        router.replace('/auth/login');
        return;
      }
      setUser(currentUser);
      setLoading(false);
    };
    checkAuth();
  }, [router]);

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="text-lg">Loading...</div>
      </div>
    );
  }

  if (!user) {
    return null;
  }

  return (
    <LayoutStudent user={user}>
      <H1>My Dashboard</H1>
      <Section title="Course Progress">
        <Courses />
      </Section>
      <Section title="Grades">
        <Grades />
      </Section>
    </LayoutStudent>
  );
}
