import { User } from '@/modules/auth/types/user';

interface LoginCredentials {
  email: string;
  password: string;
}

interface AuthResponse {
  user?: User;
  error?: string;
}

class AuthService {
  private baseUrl = '/api/auth';

  async login(credentials: LoginCredentials): Promise<AuthResponse> {
    try {
      const response = await fetch(`${this.baseUrl}/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials),
      });

      const data = await response.json();

      if (!response.ok) {
        return { error: data.error || 'Login failed' };
      }

      return { user: data.user };
    } catch (error) {
      console.error('Login error:', error);
      return { error: 'Network error' };
    }
  }

  async logout(): Promise<void> {
    try {
      await fetch(`${this.baseUrl}/logout`, {
        method: 'POST',
      });
    } catch (error) {
      console.error('Logout error:', error);
    }
  }

  async getUser(): Promise<User | null> {
    try {
      const response = await fetch(`${this.baseUrl}/me`);

      if (!response.ok) {
        return null;
      }

      const data = await response.json();
      return data.user;
    } catch (error) {
      console.error('Get user error:', error);
      return null;
    }
  }
}

export const authService = new AuthService();
