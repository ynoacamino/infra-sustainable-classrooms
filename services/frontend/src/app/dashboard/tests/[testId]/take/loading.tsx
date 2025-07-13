import { Skeleton } from '@/ui/skeleton';

export default function TakeTestLoading() {
  return (
    <div className="container mx-auto px-4 py-8 max-w-4xl">
      {/* Header */}
      <div className="flex items-center gap-4 mb-8">
        <Skeleton className="h-8 w-24" />
        <div>
          <Skeleton className="h-9 w-64 mb-2" />
          <Skeleton className="h-5 w-80" />
        </div>
      </div>

      {/* Test Info */}
      <div className="bg-blue-50 border border-blue-200 rounded-lg p-6 mb-8">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div className="text-center">
            <Skeleton className="h-5 w-16 mx-auto mb-1" />
            <Skeleton className="h-4 w-12 mx-auto" />
          </div>
          <div className="text-center">
            <Skeleton className="h-5 w-20 mx-auto mb-1" />
            <Skeleton className="h-4 w-16 mx-auto" />
          </div>
          <div className="text-center">
            <Skeleton className="h-5 w-24 mx-auto mb-1" />
            <Skeleton className="h-4 w-20 mx-auto" />
          </div>
        </div>
      </div>

      {/* Progress */}
      <div className="mb-6">
        <div className="flex justify-between items-center mb-2">
          <Skeleton className="h-4 w-24" />
          <Skeleton className="h-4 w-16" />
        </div>
        <Skeleton className="h-2 w-full rounded" />
      </div>

      {/* Question */}
      <div className="border border-gray-200 rounded-lg p-8 mb-8">
        <div className="mb-6">
          <Skeleton className="h-5 w-24 mb-3" />
          <Skeleton className="h-6 w-full mb-2" />
          <Skeleton className="h-6 w-3/4" />
        </div>

        <div className="space-y-3">
          {Array.from({ length: 4 }).map((_, i) => (
            <div key={i} className="flex items-center p-3 border border-gray-200 rounded-lg">
              <Skeleton className="h-4 w-4 mr-3 rounded-full" />
              <Skeleton className="h-4 w-48" />
            </div>
          ))}
        </div>
      </div>

      {/* Navigation */}
      <div className="flex justify-between">
        <Skeleton className="h-10 w-24" />
        <div className="flex gap-2">
          <Skeleton className="h-10 w-20" />
          <Skeleton className="h-10 w-24" />
        </div>
      </div>
    </div>
  );
}
