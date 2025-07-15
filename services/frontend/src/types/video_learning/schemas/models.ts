import z from 'zod';

export const VideoCategorySchema = z.object({
  id: z.number().int('Id must be an integer'),
  name: z.string().min(1, 'Name is required'),
});

export const VideoTagSchema = z.object({
  id: z.number().int('Id must be an integer'),
  name: z.string().min(1, 'Name is required'),
});

export const VideoSchema = z.object({
  id: z.number().int('Id must be an integer'),
  title: z.string().min(1, 'Title is required'),
  author: z.string().min(1, 'Author is required'),
  views: z
    .number()
    .int('Views must be an integer')
    .nonnegative('Views cannot be negative'),
  likes: z
    .number()
    .int('Likes must be an integer')
    .nonnegative('Likes cannot be negative'),
  thumbnail_url: z.string().url('Thumbnail URL must be a valid URL'),
});

export const VideoDetailsSchema = VideoSchema.extend({
  description: z.string(),
  video_url: z.string().url('Video URL must be a valid URL'),
  upload_date: z.number().int('Upload date must be an integer'),
  category_id: VideoCategorySchema.shape.id,
  tags_ids: z.array(VideoTagSchema.shape.id),
});

export const OwnVideoSchema = VideoSchema.omit({
  author: true,
}).extend({
  upload_date: z.number().int('Upload date must be an integer'),
});

export const CommentSchema = z.object({
  id: z.number().int('Id must be an integer'),
  author: z.string().min(1, 'Author is required'),
  date: z.number().int('Date must be an integer'),
  title: z.string().min(1, 'Title is required'),
  body: z.string().min(1, 'Body is required'),
  video_id: VideoSchema.shape.id,
});
