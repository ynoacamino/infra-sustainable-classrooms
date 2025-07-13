import { CreateTeacherProfilePayloadSchema } from "@/types/profiles/schemas/payload";
import type { Field } from "@/types/shared/field";
import type z from "zod";

export const registerTeacherFormSchema = CreateTeacherProfilePayloadSchema;
export const registerTeacherFormFields: Field<keyof z.infer<typeof registerTeacherFormSchema>>[] = [
  {
    name: "first_name",
    label: "First Name",
    type: "text",
    placeholder: "e.g. John",
    description: "Enter your first name.",
  },
  {
    name: "last_name",
    label: "Last Name",
    type: "text",
    placeholder: "e.g. Doe",
    description: "Enter your last name.",
  },
  {
    name: "email",
    label: "Email",
    type: "email",
    placeholder: "e.g. example12@gmail.com",
    description: "Enter your email address.",
  },
  {
    name: "phone",
    label: "Phone Number",
    type: "text",
    placeholder: "e.g. +1234567890",
    description: "Enter your phone number (optional).",
  },
  {
    name: "avatar_url",
    label: "Profile Picture URL",
    type: "text",
    placeholder: "e.g. https://example.com/avatar.jpg",
    description: "Enter the URL of your profile picture (optional).",
  },
  {
    name: "bio",
    label: "Biography",
    type: "textarea",
    placeholder: "Tell us about yourself...",
    description: "Write a short biography (optional).",
  },
  {
    name: "position",
    label: "Position",
    type: "text",
    placeholder: "e.g. Science Teacher",
    description: "Enter your current position or role.",
  },
]