import { CreateCoursePayloadSchema } from "@/types/text/schemas/payload";
import type { Field } from "@/types/shared/field";
import type z from "zod";

export const createCourseFormSchema = CreateCoursePayloadSchema;

export const createCourseFormFields: Field<keyof z.infer<typeof createCourseFormSchema>>[] = [
  {
    name: "title",
    label: "Course Title",
    type: "text",
    placeholder: "e.g. Introduction to JavaScript",
    description: "Enter the course title (3-150 characters).",
  },
  {
    name: "description",
    label: "Course Description",
    type: "textarea",
    placeholder: "e.g. Learn the basics of JavaScript programming language...",
    description: "Enter a detailed description of the course (10-300 characters).",
  },
  {
    name: "imageUrl",
    label: "Course Image URL",
    type: "text",
    placeholder: "e.g. https://example.com/course-image.jpg",
    description: "Enter the URL of the course image (optional, 5-500 characters).",
  },
];
