import { UpdateCoursePayloadSchema } from "@/types/text/schemas/payload";
import type { Field } from "@/types/shared/field";
import type z from "zod";

// For the form, we omit course_id since it's passed separately
export const updateCourseFormSchema = UpdateCoursePayloadSchema;

export const updateCourseFormFields: Field<keyof z.infer<typeof updateCourseFormSchema>>[] = [
  {
    name: "title",
    label: "Course Title",
    type: "text",
    placeholder: "e.g. Advanced JavaScript Programming",
    description: "Enter the course title (3-150 characters).",
  },
  {
    name: "description",
    label: "Course Description",
    type: "textarea",
    placeholder: "e.g. Deep dive into JavaScript programming language features...",
    description: "Enter a detailed description of the course (10-300 characters).",
  },
  {
    name: "imageUrl",
    label: "Course Image URL",
    type: "text",
    placeholder: "e.g. https://example.com/updated-course-image.jpg",
    description: "Enter the URL of the course image (optional, 5-500 characters).",
  },
];
