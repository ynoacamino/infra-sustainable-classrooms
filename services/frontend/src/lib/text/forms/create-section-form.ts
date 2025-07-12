import { CreateSectionPayloadSchema } from "@/types/text/schemas/payload";
import type { Field } from "@/types/shared/field";
import type z from "zod";

export const createSectionFormSchema = CreateSectionPayloadSchema;

export const createSectionFormFields: Field<keyof z.infer<typeof createSectionFormSchema>>[] = [
  {
    name: "title",
    label: "Section Title",
    type: "text",
    placeholder: "e.g. Getting Started",
    description: "Enter the section title (3-100 characters).",
  },
  {
    name: "description",
    label: "Section Description",
    type: "textarea",
    placeholder: "e.g. Introduction to the course structure...",
    description: "Enter a description of the section (5-200 characters).",
  },
  {
    name: "order",
    label: "Section Order",
    type: "number",
    placeholder: "e.g. 1",
    description: "Enter the order of the section in the course (optional, will be auto-numbered if not set).",
  },
];
