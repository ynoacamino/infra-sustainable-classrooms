import type { ServicePayload } from '@/types/shared/services/payload';
import type { ZodSchema } from 'zod';

export interface ServiceRequest<B extends ZodSchema = ZodSchema> {
  endpoint: (string | number) | (string | number)[];
  payload?: ServicePayload<B>;
  options?: RequestInit;
}
