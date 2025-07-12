import H1 from '@/ui/h1';
import Section from '@/ui/section';
import { Courses } from '@/components/courses/student/courses-list';
import { Grades } from '@/components/courses/student/grades';

export default function DashboardPage() {
  return (
    <>
      <H1>My Dashboard</H1>
      <Section title="Course Progress">
        <Courses />
      </Section>
      <Section title="Grades">
        <Grades />
      </Section>
    </>
  );
}
