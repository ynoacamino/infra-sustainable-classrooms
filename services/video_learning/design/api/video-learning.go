package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("video_learning", func() {
	Description("Video Learning microservice")

	HTTP(func() {
		Path("/video_learning")
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
	Error("permission_denied", String, "Permission denied for this operation")
	Error("profile_not_found", String, "User profile not found")

	//  Search videos by query
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

	// Get recommendations
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

	// Get video details
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

	// Get comments for video
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

	// Create comment
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

	// Get user's own videos
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

	// Initial video upload
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

	// Complete video upload
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

	// Upload thumbnail
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

	// Full getters for all tables - Video Categories
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

	// Additional useful endpoints

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

	// Get videos by category
	Method("GetVideosByCategory", func() {
		Description("Get videos by category with pagination")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "category_id", Int64, "Category ID", func() {
				Minimum(1)
			})
			Field(3, "page", Int, "Page number for pagination", func() {
				Default(1)
				Minimum(1)
			})
			Field(4, "page_size", Int, "Number of videos per page", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token", "category_id")
		})

		Result(VideoList)

		HTTP(func() {
			GET("/category/{category_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// Get similar videos
	Method("GetSimilarVideos", func() {
		Description("Get videos similar to a specific video")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID to find similar videos for", func() {
				Minimum(1)
			})
			Field(3, "page", Int, "Page number for pagination", func() {
				Default(1)
				Minimum(1)
			})
			Field(4, "page_size", Int, "Number of videos per page", func() {
				Default(10)
				Minimum(1)
				Maximum(50)
			})

			Required("session_token", "video_id")
		})

		Result(VideoList)

		HTTP(func() {
			GET("/video/{video_id}/similar")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// Delete comment
	Method("DeleteComment", func() {
		Description("Delete user's own comment")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "comment_id", Int64, "Comment ID", func() {
				Minimum(1)
			})

			Required("session_token", "comment_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			DELETE("/comment/{comment_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("permission_denied", StatusForbidden)
		})
	})

	// Create category
	Method("GetOrCreateCategory", func() {
		Description("Create a new video category or get existing one")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "name", String, "Category name", func() {
				MinLength(1)
				MaxLength(100)
			})

			Required("session_token", "name")
		})

		Result(VideoCategory)

		HTTP(func() {
			POST("/category")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// === TAG MANAGEMENT ===

	// Create or get tag
	Method("GetOrCreateTag", func() {
		Description("Create a new tag or get existing one")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "name", String, "Tag name", func() {
				MinLength(1)
				MaxLength(50)
			})

			Required("session_token", "name")
		})

		Result(VideoTag)

		HTTP(func() {
			POST("/tag")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
		})
	})

})
