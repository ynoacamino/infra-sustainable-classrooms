import type { Metadata } from 'next';
import '@fontsource-variable/lexend';
import '@/app/globals.css';
import { Toaster } from 'sonner';

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
    <html lang="en">
      <head>
        <meta charSet="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      </head>
      <body>
        <main>{children}</main>
        <Toaster />
      </body>
    </html>
  );
}
