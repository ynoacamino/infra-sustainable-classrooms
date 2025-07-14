import { Skeleton } from '@/ui/skeleton';

export default function TestsLoading() {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center justify-between mb-8">
        <div>
          <Skeleton className="h-9 w-24 mb-2" />
          <Skeleton className="h-5 w-64" />
        </div>
        <Skeleton className="h-9 w-32" />
      </div>

      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        {Array.from({ length: 6 }).map((_, i) => (
          <div key={i} className="border border-gray-200 rounded-lg p-6">
            <div className="flex items-start justify-between mb-4">
              <Skeleton className="h-6 w-48" />
              <Skeleton className="h-5 w-5 rounded-full" />
            </div>

            <div className="space-y-2 mb-4">
              <div className="flex items-center">
                <Skeleton className="h-4 w-4 mr-2" />
                <Skeleton className="h-4 w-24" />
              </div>
              <div className="flex items-center">
                <Skeleton className="h-4 w-4 mr-2" />
                <Skeleton className="h-4 w-32" />
              </div>
              <div className="flex items-center">
                <Skeleton className="h-4 w-4 mr-2" />
                <Skeleton className="h-4 w-20" />
              </div>
            </div>

            <Skeleton className="h-9 w-full" />
            <Skeleton className="h-4 w-32 mt-3" />
          </div>
        ))}
      </div>
    </div>
  );
}
