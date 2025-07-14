import type { Metadata } from 'next';
import '@fontsource-variable/lexend';
import '@/app/globals.css';
import { Toaster } from 'sonner';
import { ThemeProvider } from '@/components/shared/theme-provider';

export const metadata: Metadata = {
  title: 'Study Central',
  description: 'Sustainable Classrooms Learning Platform',
  icons: {
    icon: '/favicon.svg',
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <head>
        <meta charSet="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      </head>
      <body>
        <ThemeProvider>
          <main>{children}</main>
          <Toaster />
        </ThemeProvider>
      </body>
    </html>
  );
}
