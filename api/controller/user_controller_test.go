package controller_test

// import (
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"main/api/controller"
// 	"main/domain"
// 	"main/domain/mocks"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func setUserID(userID string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Set("x-user-id", userID)
// 		c.Next()
// 	}
// }

// func TestFetch(t *testing.T) {

// 	t.Run("success", func(t *testing.T) {
// 		mockProfile := &domain.Profile{
// 			Name:  "Test Name",
// 			Email: "test@gmail.com",
// 		}

// 		userObjectID := 1
// 		userID := string(rune(userObjectID))

// 		mockUserUsecase := new(mocks.UserUsecase)

// 		mockUserUsecase.On("GetProfileByID", mock.Anything, userID).Return(mockProfile, nil)

// 		gin := gin.Default()

// 		rec := httptest.NewRecorder()

// 		pc := &controller.UserController{
// 			UserUsecase: mockUserUsecase,
// 		}

// 		gin.Use(setUserID(userID))
// 		gin.GET("/profile", pc.Fetch)

// 		body, err := json.Marshal(mockProfile)
// 		assert.NoError(t, err)

// 		bodyString := string(body)

// 		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
// 		gin.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusOK, rec.Code)

// 		assert.Equal(t, bodyString, rec.Body.String())

// 		mockUserUsecase.AssertExpectations(t)
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		userObjectID := 2
// 		userID := string(rune(userObjectID))

// 		mockUserUsecase := new(mocks.UserUsecase)

// 		customErr := errors.New("Unexpected")

// 		mockUserUsecase.On("GetProfileByID", mock.Anything, userID).Return(nil, customErr)

// 		gin := gin.Default()

// 		rec := httptest.NewRecorder()

// 		pc := &controller.UserController{
// 			UserUsecase: mockUserUsecase,
// 		}

// 		gin.Use(setUserID(userID))
// 		gin.GET("/profile", pc.Fetch)

// 		body, err := json.Marshal(domain.ErrorResponse{Message: customErr.Error()})
// 		assert.NoError(t, err)

// 		bodyString := string(body)

// 		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
// 		gin.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusInternalServerError, rec.Code)

// 		assert.Equal(t, bodyString, rec.Body.String())

// 		mockUserUsecase.AssertExpectations(t)
// 	})

// }
