import { NextResponse } from 'next/server';
import { getUser } from '@/modules/auth/lib/authServer';

export async function GET() {
  try {
    const user = await getUser();

    if (!user) {
      return NextResponse.json({ error: 'User not found' }, { status: 401 });
    }

    return NextResponse.json({ user });
  } catch (error) {
    console.error('Get user error:', error);
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 },
    );
  }
}
