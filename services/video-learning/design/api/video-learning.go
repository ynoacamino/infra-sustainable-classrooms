package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("video-learning", func() {
	Description("Video Learning microservice")

	HTTP(func() {
		Path("/video-learning")
	})

	GRPC(func() {
		// gRPC service configuration for microservice communication
	})

	// Global error definitions for the service
	Error("invalid_input", String, "Invalid input parameters")
	Error("video_not_found", String, "Video not found")
	Error("unauthorized", String, "Unauthorized access")
	Error("upload_failed", String, "File upload failed")
	Error("service_unavailable", String, "Service temporarily unavailable")
	Error("invalid_session", String, "Invalid or expired session")

	// 1. Search videos by query
	Method("SearchVideos", func() {
		Description("Search videos by query string and category")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "query", String, "Search query string", func() {
				Example("machine learning")
				MinLength(1)
				MaxLength(200)
			})
			Field(3, "category_id", Int64, "Category ID to filter by", func() {
				Minimum(1)
			})
			Field(4, "page", Int, "Page number for pagination", func() {
				Default(1)
				Minimum(1)
			})
			Field(5, "page_size", Int, "Number of videos per page", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token", "query")
		})

		Result(VideoList)

		HTTP(func() {
			POST("/search")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	// 2. Get recommendations
	Method("GetRecommendations", func() {
		Description("Get recommended videos for user")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "page", Int, "Page number for pagination", func() {
				Default(1)
				Minimum(1)
			})
			Field(3, "page_size", Int, "Number of videos per page", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token")
		})

		Result(VideoList)

		HTTP(func() {
			POST("/recommendations")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	// 3. Get video details
	Method("GetVideo", func() {
		Description("Get complete video information including presigned URL")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID", func() {
				Minimum(1)
			})

			Required("session_token", "video_id")
		})

		Result(VideoDetails)

		HTTP(func() {
			GET("/video/{video_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// 4. Get comments for video
	Method("GetComments", func() {
		Description("Get paginated comments for a video")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID", func() {
				Minimum(1)
			})
			Field(3, "page", Int, "Page number for pagination", func() {
				Default(1)
				Minimum(1)
			})
			Field(4, "page_size", Int, "Number of comments per page", func() {
				Default(10)
				Minimum(1)
				Maximum(50)
			})

			Required("session_token", "video_id")
		})

		Result(CommentList)

		HTTP(func() {
			POST("/comments")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// 5. Create comment
	Method("CreateComment", func() {
		Description("Create a new comment on a video")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID", func() {
				Minimum(1)
			})
			Field(3, "title", String, "Comment title", func() {
				MinLength(1)
				MaxLength(150)
			})
			Field(4, "body", String, "Comment content", func() {
				MinLength(1)
				MaxLength(2000)
			})

			Required("session_token", "video_id", "title", "body")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/comment/create")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// 6. Create comment reply
	Method("CreateCommentReply", func() {
		Description("Create a reply to an existing comment")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "comment_id", Int64, "Comment ID to reply to", func() {
				Minimum(1)
			})
			Field(3, "body", String, "Reply content", func() {
				MinLength(1)
				MaxLength(1000)
			})

			Required("session_token", "comment_id", "body")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/comment/reply")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// 7. Get user's own videos
	Method("GetOwnVideos", func() {
		Description("Get authenticated user's uploaded videos")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "page", Int, "Page number for pagination", func() {
				Default(1)
				Minimum(1)
			})
			Field(3, "page_size", Int, "Number of videos per page", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token")
		})

		Result(ArrayOf(OwnVideo))

		HTTP(func() {
			GET("/my-videos")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
		})
	})

	// 8. Initial video upload
	Method("InitialUpload", func() {
		Description("Upload video file and get object name")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "file", Bytes, "Video file data")
			Field(3, "filename", String, "Original filename")

			Required("session_token", "file", "filename")
		})

		Result(UploadResponse)

		HTTP(func() {
			POST("/upload/video")
			Cookie("session_token:session")
			MultipartRequest()
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("upload_failed", StatusInternalServerError)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// 9. Complete video upload
	Method("CompleteUpload", func() {
		Description("Complete video upload with metadata")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "title", String, "Video title", func() {
				MinLength(1)
				MaxLength(150)
			})
			Field(3, "description", String, "Video description", func() {
				MaxLength(2000)
			})
			Field(4, "category_id", Int64, "Video category ID", func() {
				Minimum(1)
			})
			Field(5, "tags", ArrayOf(String), "Video tags")
			Field(6, "thumbnail_object_name", String, "Thumbnail object name in Minio")
			Field(7, "video_object_name", String, "Video object name in Minio")

			Required("session_token", "title", "category_id", "video_object_name")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/upload/complete")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// 10. Upload thumbnail
	Method("UploadThumbnail", func() {
		Description("Upload custom thumbnail for video")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "file", Bytes, "Thumbnail image file")
			Field(3, "filename", String, "Original filename")

			Required("session_token", "file", "filename")
		})

		Result(UploadResponse)

		HTTP(func() {
			POST("/upload/thumbnail")
			Cookie("session_token:session")
			MultipartRequest()
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("upload_failed", StatusInternalServerError)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// 11. Randomly pick thumbnail from video
	Method("GenerateThumbnail", func() {
		Description("Generate thumbnail from video file")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_object_name", String, "Video object name in Minio")

			Required("session_token", "video_object_name")
		})

		Result(ThumbnailResponse)

		HTTP(func() {
			POST("/thumbnail/generate")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("upload_failed", StatusInternalServerError)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// 12. Full getters for all tables - Video Categories
	Method("GetAllCategories", func() {
		Description("Get all video categories")

		Payload(func() {
			Field(1, "session_token", String, "User session token")

			Required("session_token")
		})

		Result(ArrayOf(VideoCategory))

		HTTP(func() {
			GET("/categories")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
		})
	})

	// Get all tags
	Method("GetAllTags", func() {
		Description("Get all video tags")

		Payload(func() {
			Field(1, "session_token", String, "User session token")

			Required("session_token")
		})

		Result(ArrayOf(VideoTag))

		HTTP(func() {
			GET("/tags")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
		})
	})

	// Get all videos (admin/debug purpose)
	Method("GetAllVideos", func() {
		Description("Get all videos with pagination (admin purpose)")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "page", Int, "Page number for pagination", func() {
				Default(1)
				Minimum(1)
			})
			Field(3, "page_size", Int, "Number of videos per page", func() {
				Default(50)
				Minimum(1)
				Maximum(200)
			})

			Required("session_token")
		})

		Result(VideoList)

		HTTP(func() {
			GET("/admin/videos")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
		})
	})

	// Get all comments (admin/debug purpose)
	Method("GetAllComments", func() {
		Description("Get all comments with pagination (admin purpose)")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "page", Int, "Page number for pagination", func() {
				Default(1)
				Minimum(1)
			})
			Field(3, "page_size", Int, "Number of comments per page", func() {
				Default(50)
				Minimum(1)
				Maximum(200)
			})

			Required("session_token")
		})

		Result(CommentList)

		HTTP(func() {
			GET("/admin/comments")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
		})
	})

	// Get user category preferences
	Method("GetUserCategoryPreferences", func() {
		Description("Get user's category preferences")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "user_id", Int64, "User ID", func() {
				Minimum(1)
			})

			Required("session_token", "user_id")
		})

		Result(ArrayOf(UserCategoryLike))

		HTTP(func() {
			GET("/user/{user_id}/preferences")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// 13. Additional useful endpoints

	// Like/Unlike video
	Method("ToggleVideoLike", func() {
		Description("Toggle like status for a video")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID", func() {
				Minimum(1)
			})

			Required("session_token", "video_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/video/{video_id}/like")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// TODO this could be done in get video details and avoid the extra request
	// Increment video views
	Method("IncrementViews", func() {
		Description("Increment view count for a video")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID", func() {
				Minimum(1)
			})

			Required("session_token", "video_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/video/{video_id}/view")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// Delete video (for content creators)
	Method("DeleteVideo", func() {
		Description("Delete user's own video")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID", func() {
				Minimum(1)
			})

			Required("session_token", "video_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			DELETE("/video/{video_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// Update video metadata
	Method("UpdateVideo", func() {
		Description("Update video metadata (title, description, etc.)")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID", func() {
				Minimum(1)
			})
			Field(3, "title", String, "New video title", func() {
				MinLength(1)
				MaxLength(150)
			})
			Field(4, "description", String, "New video description", func() {
				MaxLength(2000)
			})
			Field(5, "category_id", Int64, "New category ID", func() {
				Minimum(1)
			})
			Field(6, "tags", ArrayOf(String), "New video tags")

			Required("session_token", "video_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			PUT("/video/{video_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// gRPC method for validating user sessions (inter-service communication)
	Method("ValidateUserRole", func() {
		Description("Validate user role for inter-service communication")

		Payload(func() {
			Field(1, "user_id", Int64, "User ID to validate")

			Required("user_id")
		})

		Result(RoleValidationResponse)

		// This method is only for gRPC inter-service communication
		GRPC(func() {
			Response(CodeOK)
			Response("profile_not_found", CodeNotFound)
			Response("permission_denied", CodePermissionDenied)
		})
	})

})
