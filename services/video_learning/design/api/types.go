package design

import (
	. "goa.design/goa/v3/dsl"
)

// Video type for basic video information
var Video = Type("Video", func() {
	Description("Video information")

	Field(1, "id", Int64, "Video unique identifier")
	Field(2, "title", String, "Video title")
	Field(3, "author", String, "Video author/creator")
	Field(4, "views", Int, "Number of views")
	Field(5, "likes", Int, "Number of likes")
	Field(6, "thumbnail_url", String, "Minio presigned URL for thumbnail")

	Required("id", "title", "author", "views", "likes", "thumbnail_url")
})

// VideoDetails type for complete video information
var VideoDetails = Type("VideoDetails", func() {
	Description("Complete video information")

	Field(1, "id", Int64, "Video unique identifier")
	Field(2, "title", String, "Video title")
	Field(3, "description", String, "Video description")
	Field(4, "author", String, "Video author/creator")
	Field(5, "views", Int, "Number of views")
	Field(6, "likes", Int, "Number of likes")
	Field(7, "video_url", String, "Minio presigned URL for video")
	Field(8, "thumbnail_url", String, "Minio presigned URL for thumbnail")
	Field(9, "upload_date", Int64, "Upload timestamp in milliseconds")
	Field(10, "category", String, "Video category")
	Field(11, "tags", ArrayOf(String), "Video tags")
	Field(12, "similar_videos", ArrayOf(Video), "List of similar videos")

	Required("id", "title", "description", "author", "views", "likes", "video_url", "thumbnail_url", "upload_date", "category", "tags", "similar_videos")
})

// OwnVideo type for user's own videos
var OwnVideo = Type("OwnVideo", func() {
	Description("User's own video information")

	Field(1, "id", Int64, "Video unique identifier")
	Field(2, "title", String, "Video title")
	Field(3, "views", Int, "Number of views")
	Field(4, "likes", Int, "Number of likes")
	Field(5, "thumbnail_url", String, "Minio presigned URL for thumbnail")
	Field(6, "upload_date", Int64, "Upload timestamp in milliseconds")

	Required("id", "title", "views", "likes", "thumbnail_url", "upload_date")
})

// Comment type for video comments
var Comment = Type("Comment", func() {
	Description("Video comment information")

	Field(1, "id", Int64, "Comment unique identifier")
	Field(2, "author", String, "Comment author")
	Field(3, "date", Int64, "Comment publish date in milliseconds")
	Field(4, "title", String, "Comment title")
	Field(5, "body", String, "Comment content")

	Required("id", "author", "date", "title", "body")
})

// CommentReply type for comment replies
var CommentReply = Type("CommentReply", func() {
	Description("Comment reply information")

	Field(1, "id", Int64, "Reply unique identifier")
	Field(2, "author", String, "Reply author")
	Field(3, "date", Int64, "Reply publish date in milliseconds")
	Field(4, "body", String, "Reply content")

	Required("id", "author", "date", "body")
})

// VideoCategory type for video categories
var VideoCategory = Type("VideoCategory", func() {
	Description("Video category information")

	Field(1, "id", Int64, "Category unique identifier")
	Field(2, "name", String, "Category name")

	Required("id", "name")
})

// VideoTag type for video tags
var VideoTag = Type("VideoTag", func() {
	Description("Video tag information")

	Field(1, "id", Int64, "Tag unique identifier")
	Field(2, "name", String, "Tag name")

	Required("id", "name")
})

// UserCategoryLike type for user category preferences
var UserCategoryLike = Type("UserCategoryLike", func() {
	Description("User category preference information")

	Field(1, "user_id", Int64, "User unique identifier")
	Field(2, "category_id", Int64, "Category unique identifier")
	Field(3, "likes", Int, "Number of likes for this category")

	Required("user_id", "category_id", "likes")
})

// UploadResponse type for file upload responses
var UploadResponse = Type("UploadResponse", func() {
	Description("File upload response")

	Field(1, "object_name", String, "Minio object name")
	Field(2, "presigned_url", String, "Presigned URL for accessing the file")

	Required("object_name")
})

// ThumbnailResponse type for thumbnail generation responses
var ThumbnailResponse = Type("ThumbnailResponse", func() {
	Description("Thumbnail generation response")

	Field(1, "thumbnail_object_name", String, "Minio object name for thumbnail")
	Field(2, "presigned_url", String, "Presigned URL for accessing the thumbnail")

	Required("thumbnail_object_name", "presigned_url")
})

// SimpleResponse type for simple success/message responses
var SimpleResponse = Type("SimpleResponse", func() {
	Description("Simple response with success status and message")

	Field(1, "success", Boolean, "Operation success status")
	Field(2, "message", String, "Response message")
	Field(3, "id", Int64, "Created resource ID (when applicable)")

	Required("success", "message")
})

// VideoList type for paginated video lists
var VideoList = Type("VideoList", func() {
	Description("Paginated list of videos")

	Field(1, "videos", ArrayOf(Video), "List of videos")
	Field(2, "total", Int, "Total number of videos")

	Required("videos", "total")
})

// CommentList type for paginated comment lists
var CommentList = Type("CommentList", func() {
	Description("Paginated list of comments")

	Field(1, "comments", ArrayOf(Comment), "List of comments")
	Field(2, "total", Int, "Total number of comments")

	Required("comments", "total")
})

// UserSession type for user session information (from profiles microservice)
var UserSession = Type("UserSession", func() {
	Description("User session information from profiles microservice")

	Field(1, "user_id", Int64, "User unique identifier")
	Field(2, "username", String, "Username")
	Field(3, "email", String, "User email")
	Field(4, "session_valid", Boolean, "Session validity status")

	Required("user_id", "username", "session_valid")
})

// === INTER-SERVICE COMMUNICATION TYPES ===
var RoleValidationResponse = Type("RoleValidationResponse", func() {
	Description("Response for role validation")

	Field(1, "user_id", Int64, "User unique identifier")
	Field(2, "role", String, "User role (student, teacher)")

	Required("user_id", "role")
})
