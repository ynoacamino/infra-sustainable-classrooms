import z from 'zod';

/**
 * Schema for validating user identifiers (username or email).
 *
 * This schema enforces:
 * - Minimum length of 3 characters to prevent too short identifiers
 * - Maximum length of 100 characters to prevent database overflow
 * - String type validation
 *
 * @example
 * ```typescript
 * const result = UserIdentifier.safeParse("user@example.com");
 * if (result.success) {
 *   console.log("Valid identifier:", result.data);
 * }
 * ```
 *
 * @see {@link https://zod.dev/ | Zod Documentation}
 */
export const UserIdentifier = z
  .string()
  .min(3, 'Identifier must be at least 3 characters long')
  .max(100, 'Identifier must be at most 100 characters long')
  .describe('User identifier (username/email)');

/**
 * Schema for TOTP (Time-based One-Time Password) secret configuration.
 *
 * This schema validates the complete TOTP setup data returned from the server,
 * including the authenticator URL, backup codes, and issuer information.
 *
 * @example
 * ```typescript
 * const totpData = {
 *   totp_url: "otpauth://totp/MyApp:user@sample.com?secret=ABC123&issuer=MyApp",
 *   backup_codes: ["12345678", "87654321", "11111111"],
 *   issuer: "MyApp"
 * };
 *
 * const result = TOTPSecretSchema.safeParse(totpData);
 * ```
 *
 * @property {string} totp_url - The otpauth:// URL for authenticator apps (RFC 6238 compliant)
 * @property {string[]} backup_codes - Array of single-use backup codes for account recovery
 * @property {string} issuer - The application/service name displayed in authenticator apps
 */
export const TOTPSecretSchema = z.object({
  totp_url: z.string().url(),
  backup_codes: z.array(z.string()),
  issuer: z.string(),
});

/**
 * Schema for user data model with authentication-related information.
 *
 * Represents a complete user entity with authentication status, timestamps,
 * and extensible metadata. All numeric timestamps are Unix timestamps in seconds.
 *
 * @example
 * ```typescript
 * const user = {
 *   id: 12345,
 *   identifier: "john.doe@example.com",
 *   created_at: 1640995200, // Unix timestamp
 *   last_login: 1641081600, // Unix timestamp
 *   is_verified: true,
 *   metadata: { role: "admin", preferences: "dark_mode" }
 * };
 *
 * const result = UserSchema.safeParse(user);
 * ```
 *
 * @property {number} id - Unique integer identifier for the user (primary key)
 * @property {string} identifier - User's login identifier (validated by UserIdentifier schema)
 * @property {number} created_at - Unix timestamp when the user account was created
 * @property {number} last_login - Unix timestamp of the user's last successful login
 * @property {boolean} is_verified - Whether the user has completed email/phone verification
 * @property {Record<string, string>} metadata - Extensible key-value pairs for additional user data
 */
export const UserSchema = z.object({
  id: z.number().int(),
  identifier: UserIdentifier,
  created_at: z.number().int(),
  last_login: z.number().int(),
  is_verified: z.boolean(),
  metadata: z.record(z.string()),
});

/**
 * Schema for device and session information tracking.
 *
 * Captures device-specific information for security monitoring, session management,
 * and analytics. All fields are optional to handle cases where information
 * may not be available or privacy settings restrict data collection.
 *
 * @example
 * ```typescript
 * const deviceInfo = {
 *   user_agent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
 *   ip_address: "192.168.1.100",
 *   device_id: "device_abc123xyz",
 *   platform: "web"
 * };
 *
 * const result = DeviceInfoSchema.safeParse(deviceInfo);
 * ```
 *
 * @property {string} [user_agent] - Browser/client user agent string for compatibility tracking
 * @property {string} [ip_address] - Client IP address for geolocation and security monitoring
 * @property {string} [device_id] - Unique device identifier for device-specific features
 * @property {string} [platform] - Platform type (web, mobile, desktop, etc.)
 */
export const DeviceInfoSchema = z.object({
  user_agent: z.string().optional(),
  ip_address: z.string().optional(),
  device_id: z.string().optional(),
  platform: z.string().optional(),
});
