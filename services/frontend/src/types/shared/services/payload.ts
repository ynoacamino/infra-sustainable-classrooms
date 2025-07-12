import type { z, ZodSchema } from 'zod';

export interface ServicePayload<T extends ZodSchema = ZodSchema> {
  schema: T;
  data: z.infer<T>;
}
