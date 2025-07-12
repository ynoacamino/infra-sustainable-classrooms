import { textService } from "@/services/text/service";
import { cookies } from "next/headers";

export default async function CouresPage() {
  const text = await textService(cookies());
  const courses = await text.listCourses({});

  if (!courses.success) {
    return <div>Error loading courses: {courses.error.message}</div>;
  }

  return (
    <div className="flex flex-col items-center justify-center w-full h-full">
      <h1 className="text-2xl font-bold mb-4">Courses</h1>
      <ul className="list-disc">
        {courses.data.map(course => (
          <li key={course.id} className="mb-2">
            {course.title}
          </li>
        ))}
      </ul>
    </div>
  );
}