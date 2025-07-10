import { ReactNode, HTMLAttributes } from 'react';

interface H1Props extends HTMLAttributes<HTMLHeadingElement> {
  children: ReactNode;
}

export default function H1({ children, ...rest }: H1Props) {
  return (
    <div className="w-full flex justify-center py-4 px-4 sm:px-6 md:px-8 max-w-5xl">
      <div className="flex w-full flex-col gap-4">
        <h1 className="text-[32px] font-bold" {...rest}>
          {children}
        </h1>
      </div>
    </div>
  );
}
