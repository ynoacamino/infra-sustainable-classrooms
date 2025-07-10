import Editor from '@/components/post/admin/editor';
import Preview from '@/components/post/admin/preview';
import { cn } from '@/lib/shared/utils';
import { Button } from '@/ui/button';
import { Link } from '@/ui/link';
import { Skeleton } from '@/ui/skeleton';
import { IconCancel, IconDeviceFloppy } from '@tabler/icons-react';

export function ModeSwitchSkeleton() {
  return (
    <div className="w-full flex-col gap-8 flex pb-10 max-w-5xl">
      <div className="w-full flex border-border border-b pb-2">
        <Skeleton className="h-8 w-20" />
        <Skeleton className="h-8 w-20 ml-4" />
      </div>
      <Skeleton className="h-96 w-full" />
    </div>
  );
}

export default function ModeSwitch({
  viewOrEdit,
  setViewOrEdit,
  content,
  setContent,
  className,
}: {
  viewOrEdit: 'preview' | 'edit';
  setViewOrEdit: (mode: 'preview' | 'edit') => void;
  content: string;
  setContent: (value: string) => void;
  className?: string;
}) {
  return (
    <div className="w-full flex gap-10 relative">
      <div
        className={cn('w-full flex-col gap-8 flex pb-10 max-w-5xl', className)}
      >
        <div className="w-full border-border border-b">
          <button
            className={cn(
              'font-bold text-sm text-foreground/60 px-6 py-3 cursor-pointer',
              {
                'text-foreground': viewOrEdit === 'edit',
              },
            )}
            onClick={() => setViewOrEdit('edit')}
          >
            Edit
          </button>
          <button
            className={cn(
              'font-bold text-sm text-foreground/60 px-6 py-3 cursor-pointer',
              {
                'text-foreground': viewOrEdit === 'preview',
              },
            )}
            onClick={() => setViewOrEdit('preview')}
          >
            Preview
          </button>
        </div>
        <div className="flex w-full">
          <div className="w-full">
            {viewOrEdit === 'preview' ? (
              <Preview content={content} />
            ) : (
              <Editor
                content={content}
                onChange={(value) => setContent(value)}
              />
            )}
          </div>
        </div>
      </div>
      <div className="flex flex-col gap-3 flex-1 sticky top-8 h-full">
        <Button>
          <IconDeviceFloppy />
          Save
        </Button>
        <Link href="/teacher/post" variant="outline" className="">
          <IconCancel />
          Cancel
        </Link>
      </div>
    </div>
  );
}
