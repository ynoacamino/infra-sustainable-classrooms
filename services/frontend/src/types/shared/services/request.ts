import type { ServicePayload } from '@/types/shared/services/payload';
import type { ZodSchema } from 'zod';

// TODO: Remove ZodSchema as defult to get complete type inference
export interface ServiceRequest<B extends ZodSchema = ZodSchema> {
  endpoint: (string | number) | (string | number)[];
  payload?: ServicePayload<B>;
  query?: (keyof ServicePayload<B>['data'])[];
  options?: RequestInit;
}
