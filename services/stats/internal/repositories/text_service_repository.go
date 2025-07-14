package repositories

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/internal/ports"
	"google.golang.org/grpc"

	// Importaciones para el cliente gRPC del servicio text
	genTextClient "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/grpc/text/client"
	textGen "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"
)

type TextServiceRepository struct {
	client *textGen.Client
}

func NewTextServiceRepository(grpcConn *grpc.ClientConn) ports.TextServiceRepository {
	grpcClient := genTextClient.NewClient(grpcConn)

	client := &textGen.Client{
		// Aquí mapearemos los endpoints cuando estén disponibles en gRPC
		CheckArticleCompletedEndpoint: grpcClient.CheckArticleCompleted(),
	}

	return &TextServiceRepository{
		client: client,
	}
}

func (r *TextServiceRepository) GetUserProgressForCourse(ctx context.Context, userID int64, courseID int64) (*ports.UserProgressForCourseResponse, error) {
	// Por ahora implementaremos la lógica básica
	// En el futuro esto se conectará directamente con los endpoints gRPC del servicio text
	return &ports.UserProgressForCourseResponse{
		Articles: []ports.ArticleProgressData{},
	}, nil
}

func (r *TextServiceRepository) GetUserCompletedArticles(ctx context.Context, userID int64) (*ports.UserCompletedArticlesResponse, error) {
	// Por ahora implementaremos la lógica básica
	return &ports.UserCompletedArticlesResponse{
		Articles: []ports.CompletedArticleData{},
	}, nil
}

func (r *TextServiceRepository) GetCourseCompletionStats(ctx context.Context, userID int64, courseID int64) (*ports.CourseCompletionStatsResponse, error) {
	// Por ahora implementaremos la lógica básica
	return &ports.CourseCompletionStatsResponse{
		TotalArticles:        0,
		CompletedArticles:    0,
		CompletionPercentage: 0,
	}, nil
}
