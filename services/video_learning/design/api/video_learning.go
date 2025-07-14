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
	Error("category_not_found", String, "Video category not found")
	Error("tag_not_found", String, "Video tag not found")

	//  Search videos by query
	// DONE in frontend
	// NOT TESTED
	Method("SearchVideos", func() {
		Description("Search videos by query string and category, paginated")

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
			GET("/search")
			Params(func() {
				Param("query")
				Param("category_id")
				Param("page")
				Param("page_size")
			})
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	// Get recommendations
	// DONE in frontend
	// NOT TESTED
	Method("GetRecommendations", func() {
		Description("Get random set of recommended videos for user")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "amount", Int, "How many videos to get", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token")
		})

		Result(VideoList)

		HTTP(func() {
			GET("/recommendations")
			Params(func() {
				Param("amount")
			})
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	// Get video details
	// DONE in frontend
	// NOT TESTED
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
	// DONE in frontend
	// NOT TESTED
	Method("GetComments", func() {
		Description("Get paginated comments for a video")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID to get comments for", func() {
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

			Required("video_id", "session_token")
		})

		Result(CommentList)

		HTTP(func() {
			GET("/comments/{video_id}")
			Params(func() {
				Param("page")
				Param("page_size")
			})
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// Create comment
	// DONE in frontend
	// NOT TESTED
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
			POST("/comments/{video_id}/create")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// Get user's own videos
	// DONE in frontend
	// NOT TESTED
	Method("GetOwnVideos", func() {
		Description("Get authenticated user's uploaded videos, paginated")

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
			Params(func() {
				Param("page")
				Param("page_size")
			})
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
		})
	})

	// Initial video upload
	// DONE in frontend
	// NOT TESTED
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
	// DONE in frontend
	// NOT TESTED
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
			// BUG: Fails here, send tags ids
			Field(5, "tags", ArrayOf(String), "Video tags")
			Field(6, "thumbnail_object_name", String, "Thumbnail object name in Minio")
			Field(7, "video_object_name", String, "Video object name in Minio")

			Required("session_token", "title", "category_id", "video_object_name", "thumbnail_object_name")
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
	// DONE in frontend
	// NOT TESTED
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
	// DONE in frontend
	// NOT TESTED
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
	// DONE in frontend
	// NOT TESTED
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
	// DONE in frontend
	// NOT TESTED
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
			PUT("/video/{video_id}/like")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// Delete video (for content creators)
	// DONE in frontend
	// NOT TESTED
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

	// Get videos by category
	// DONE in frontend
	// NOT TESTED
	Method("GetVideosByCategory", func() {
		Description("Get random set of videos of a category")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "category_id", Int64, "Category ID", func() {
				Minimum(1)
			})
			Field(3, "amount", Int, "How many videos to get", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token", "category_id")
		})

		Result(VideoList)

		HTTP(func() {
			GET("/category/{category_id}")
			Params(func() {
				Param("amount")
			})
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// Get similar videos
	// DONE in frontend
	// NOT TESTED
	Method("GetSimilarVideos", func() {
		Description("Get random set of videos similar to a specific video")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "video_id", Int64, "Video ID to find similar videos for", func() {
				Minimum(1)
			})
			Field(3, "amount", Int, "How many similar videos to get", func() {
				Default(10)
				Minimum(1)
				Maximum(50)
			})

			Required("session_token", "video_id")
		})

		Result(VideoList)

		HTTP(func() {
			GET("/video/{video_id}/similar")
			Params(func() {
				Param("amount")
			})
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("video_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// Delete comment
	// DONE in frontend
	// NOT TESTED
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

	// Get or create category
	// DONE in frontend
	// NOT TESTED
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

	// Get category by ID
	Method("GetCategoryById", func() {
		Description("Get video category by ID")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "category_id", Int64, "Category ID", func() {
				Minimum(1)
			})

			Required("session_token", "category_id")
		})

		Result(VideoCategory)

		HTTP(func() {
			GET("/category/{category_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("category_not_found", StatusNotFound)
		})
	})

	// === TAG MANAGEMENT ===

	// Create or get tag
	// DONE in frontend
	// NOT TESTED
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

	// Get tag by ID
	Method("GetTagById", func() {
		Description("Get video tag by ID")

		Payload(func() {
			Field(1, "session_token", String, "User session token")
			Field(2, "tag_id", Int64, "Tag ID", func() {
				Minimum(1)
			})

			Required("session_token", "tag_id")
		})

		Result(VideoTag)

		HTTP(func() {
			GET("/tag/{tag_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_session", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("tag_not_found", StatusNotFound)
		})
	})
})
