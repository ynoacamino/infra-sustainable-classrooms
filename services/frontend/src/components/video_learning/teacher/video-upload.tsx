'use client';

import { useState, useEffect } from 'react';
import {
  Upload,
  Video,
  Image as ImageIcon,
  X,
  CheckCircle,
} from 'lucide-react';
import Image from 'next/image';
import { Button } from '@/ui/button';
import { Input } from '@/ui/input';
import { Label } from '@/ui/label';
import { Textarea } from '@/ui/textarea';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/ui/select';
import { videoLearningService } from '@/services/video_learning/service';
import { cookies } from 'next/headers';
import type { VideoCategory, VideoTag } from '@/types/video_learning/models';
import { toast } from 'sonner';

interface UploadState {
  video: File | null;
  thumbnail: File | null;
  title: string;
  description: string;
  categoryId: string;
  selectedTags: number[];
}

export function VideoUpload() {
  const [uploadState, setUploadState] = useState<UploadState>({
    video: null,
    thumbnail: null,
    title: '',
    description: '',
    categoryId: '',
    selectedTags: [],
  });
  const [categories, setCategories] = useState<VideoCategory[]>([]);
  const [tags, setTags] = useState<VideoTag[]>([]);
  const [loading, setLoading] = useState(false);
  const [uploadProgress, setUploadProgress] = useState(0);
  const [currentStep, setCurrentStep] = useState<
    'form' | 'uploading' | 'complete'
  >('form');
  const [videoPreview, setVideoPreview] = useState<string | null>(null);
  const [thumbnailPreview, setThumbnailPreview] = useState<string | null>(null);

  useEffect(() => {
    loadCategoriesAndTags();
  }, []);

  const loadCategoriesAndTags = async () => {
    try {
      const service = await videoLearningService(cookies());
      const [categoriesResult, tagsResult] = await Promise.all([
        service.getAllCategories(),
        service.getAllTags(),
      ]);

      if (categoriesResult.success) {
        setCategories(categoriesResult.data);
      }
      if (tagsResult.success) {
        setTags(tagsResult.data);
      }
    } catch (error) {
      console.error('Failed to load categories and tags:', error);
      toast.error('Failed to load categories and tags');
    }
  };

  const handleFileChange = (type: 'video' | 'thumbnail', file: File) => {
    setUploadState((prev) => ({ ...prev, [type]: file }));

    // Create preview for video
    if (type === 'video') {
      const url = URL.createObjectURL(file);
      setVideoPreview(url);
    }

    // Create preview for thumbnail
    if (type === 'thumbnail') {
      const url = URL.createObjectURL(file);
      setThumbnailPreview(url);
    }
  };

  const handleTagToggle = (tagId: number) => {
    setUploadState((prev) => ({
      ...prev,
      selectedTags: prev.selectedTags.includes(tagId)
        ? prev.selectedTags.filter((id) => id !== tagId)
        : [...prev.selectedTags, tagId],
    }));
  };

  const validateForm = () => {
    if (!uploadState.video) {
      toast.error('Please select a video file');
      return false;
    }
    if (!uploadState.title.trim()) {
      toast.error('Please enter a video title');
      return false;
    }
    if (!uploadState.description.trim()) {
      toast.error('Please enter a video description');
      return false;
    }
    if (!uploadState.categoryId) {
      toast.error('Please select a category');
      return false;
    }
    return true;
  };

  const handleUpload = async () => {
    if (!validateForm()) return;

    setLoading(true);
    setCurrentStep('uploading');
    setUploadProgress(0);

    try {
      const service = await videoLearningService(cookies());

      // Step 1: Initial video upload
      setUploadProgress(20);
      const videoUploadResult = await service.initialUpload({
        file: uploadState.video!,
        title: uploadState.title,
        description: uploadState.description,
        category_id: parseInt(uploadState.categoryId),
        tags_ids: uploadState.selectedTags,
      });

      if (!videoUploadResult.success) {
        throw new Error('Failed to upload video');
      }

      // Step 2: Upload thumbnail if provided
      if (uploadState.thumbnail) {
        setUploadProgress(60);
        const thumbnailUploadResult = await service.uploadThumbnail({
          file: uploadState.thumbnail,
          video_id: videoUploadResult.data.video_id,
        });

        if (!thumbnailUploadResult.success) {
          console.warn('Thumbnail upload failed, continuing with video upload');
        }
      }

      // Step 3: Complete upload
      setUploadProgress(90);
      const completeResult = await service.completeUpload({
        video_id: videoUploadResult.data.video_id,
      });

      if (!completeResult.success) {
        throw new Error('Failed to complete video upload');
      }

      setUploadProgress(100);
      setCurrentStep('complete');
      toast.success('Video uploaded successfully!');

      // Reset form
      setTimeout(() => {
        setUploadState({
          video: null,
          thumbnail: null,
          title: '',
          description: '',
          categoryId: '',
          selectedTags: [],
        });
        setVideoPreview(null);
        setThumbnailPreview(null);
        setCurrentStep('form');
        setUploadProgress(0);
      }, 2000);
    } catch (error) {
      console.error('Upload failed:', error);
      toast.error('Upload failed. Please try again.');
      setCurrentStep('form');
      setUploadProgress(0);
    } finally {
      setLoading(false);
    }
  };

  const removeFile = (type: 'video' | 'thumbnail') => {
    setUploadState((prev) => ({ ...prev, [type]: null }));
    if (type === 'video') {
      setVideoPreview(null);
    } else {
      setThumbnailPreview(null);
    }
  };

  if (currentStep === 'uploading') {
    return (
      <div className="max-w-md mx-auto text-center space-y-4">
        <div className="w-16 h-16 mx-auto bg-primary/10 rounded-full flex items-center justify-center">
          <Upload className="h-8 w-8 text-primary" />
        </div>
        <h3 className="text-lg font-semibold">Uploading Video...</h3>
        <div className="w-full bg-muted rounded-full h-2">
          <div
            className="bg-primary h-2 rounded-full transition-all duration-300"
            style={{ width: `${uploadProgress}%` }}
          />
        </div>
        <p className="text-sm text-muted-foreground">
          {uploadProgress}% complete
        </p>
      </div>
    );
  }

  if (currentStep === 'complete') {
    return (
      <div className="max-w-md mx-auto text-center space-y-4">
        <div className="w-16 h-16 mx-auto bg-green-100 rounded-full flex items-center justify-center">
          <CheckCircle className="h-8 w-8 text-green-600" />
        </div>
        <h3 className="text-lg font-semibold">Upload Complete!</h3>
        <p className="text-sm text-muted-foreground">
          Your video has been uploaded successfully and is now available.
        </p>
      </div>
    );
  }

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {/* Video Upload */}
        <div className="space-y-4">
          <Label htmlFor="video-upload">Video File *</Label>
          <div className="border-2 border-dashed border-muted-foreground/25 rounded-lg p-6">
            {!uploadState.video ? (
              <div className="text-center">
                <Video className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
                <p className="text-sm text-muted-foreground mb-2">
                  Select a video file to upload
                </p>
                <Input
                  id="video-upload"
                  type="file"
                  accept="video/*"
                  className="hidden"
                  onChange={(e) => {
                    const file = e.target.files?.[0];
                    if (file) handleFileChange('video', file);
                  }}
                />
                <Button
                  variant="outline"
                  onClick={() =>
                    document.getElementById('video-upload')?.click()
                  }
                >
                  Choose Video
                </Button>
              </div>
            ) : (
              <div className="space-y-2">
                <div className="flex items-center justify-between">
                  <p className="text-sm font-medium">
                    {uploadState.video.name}
                  </p>
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => removeFile('video')}
                  >
                    <X className="h-4 w-4" />
                  </Button>
                </div>
                {videoPreview && (
                  <video
                    src={videoPreview}
                    controls
                    className="w-full h-32 object-cover rounded"
                  />
                )}
                <p className="text-xs text-muted-foreground">
                  {(uploadState.video.size / (1024 * 1024)).toFixed(2)} MB
                </p>
              </div>
            )}
          </div>
        </div>

        {/* Thumbnail Upload */}
        <div className="space-y-4">
          <Label htmlFor="thumbnail-upload">Thumbnail (Optional)</Label>
          <div className="border-2 border-dashed border-muted-foreground/25 rounded-lg p-6">
            {!uploadState.thumbnail ? (
              <div className="text-center">
                <ImageIcon className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
                <p className="text-sm text-muted-foreground mb-2">
                  Select a thumbnail image
                </p>
                <Input
                  id="thumbnail-upload"
                  type="file"
                  accept="image/*"
                  className="hidden"
                  onChange={(e) => {
                    const file = e.target.files?.[0];
                    if (file) handleFileChange('thumbnail', file);
                  }}
                />
                <Button
                  variant="outline"
                  onClick={() =>
                    document.getElementById('thumbnail-upload')?.click()
                  }
                >
                  Choose Image
                </Button>
              </div>
            ) : (
              <div className="space-y-2">
                <div className="flex items-center justify-between">
                  <p className="text-sm font-medium">
                    {uploadState.thumbnail.name}
                  </p>
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => removeFile('thumbnail')}
                  >
                    <X className="h-4 w-4" />
                  </Button>
                </div>
                {thumbnailPreview && (
                  <Image
                    src={thumbnailPreview}
                    alt="Thumbnail preview"
                    width={320}
                    height={180}
                    className="w-full h-32 object-cover rounded"
                  />
                )}
              </div>
            )}
          </div>
        </div>
      </div>

      {/* Video Details */}
      <div className="space-y-4">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="space-y-2">
            <Label htmlFor="title">Title *</Label>
            <Input
              id="title"
              value={uploadState.title}
              onChange={(e) =>
                setUploadState((prev) => ({ ...prev, title: e.target.value }))
              }
              placeholder="Enter video title"
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="category">Category *</Label>
            <Select
              value={uploadState.categoryId}
              onValueChange={(value) =>
                setUploadState((prev) => ({ ...prev, categoryId: value }))
              }
            >
              <SelectTrigger>
                <SelectValue placeholder="Select category" />
              </SelectTrigger>
              <SelectContent>
                {categories.map((category) => (
                  <SelectItem key={category.id} value={category.id.toString()}>
                    {category.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </div>

        <div className="space-y-2">
          <Label htmlFor="description">Description *</Label>
          <Textarea
            id="description"
            value={uploadState.description}
            onChange={(e) =>
              setUploadState((prev) => ({
                ...prev,
                description: e.target.value,
              }))
            }
            placeholder="Enter video description"
            rows={4}
          />
        </div>

        {/* Tags */}
        <div className="space-y-2">
          <Label>Tags</Label>
          <div className="flex flex-wrap gap-2">
            {tags.map((tag) => (
              <Button
                key={tag.id}
                variant={
                  uploadState.selectedTags.includes(tag.id)
                    ? 'default'
                    : 'outline'
                }
                size="sm"
                onClick={() => handleTagToggle(tag.id)}
              >
                {tag.name}
              </Button>
            ))}
          </div>
        </div>
      </div>

      {/* Upload Button */}
      <div className="flex justify-center">
        <Button
          onClick={handleUpload}
          disabled={
            loading ||
            !uploadState.video ||
            !uploadState.title ||
            !uploadState.description ||
            !uploadState.categoryId
          }
          className="px-8"
        >
          {loading ? 'Uploading...' : 'Upload Video'}
        </Button>
      </div>
    </div>
  );
}
