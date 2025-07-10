import { authService } from '@/services/auth/auth';
import { NextResponse } from 'next/server';

export async function POST() {
  try {
    await authService.logout();
    return NextResponse.json({ success: true });
  } catch (error) {
    console.error('Logout error:', error);
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 },
    );
  }
}
