package usecase

import (
	"atwell/domain"
	mocks "atwell/mocks/domain"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTweetUsecase_Get(t *testing.T) {
	// mock
	userRepository := new(mocks.UserRepository)
	email := "test_get@email.com"
	mockUser := domain.User{Email: email}
	userRepository.On("Get", email).Return(
		mockUser,
		nil,
	)

	tweetRepository := new(mocks.TweetRepository)
	mockedTwList := make([]domain.Tweet, 0)
	mockedTwList = append(
		mockedTwList,
		domain.Tweet{Comment: "test_get"},
	)
	from := time.Now()
	to := time.Now()
	tweetRepository.On("Get", mockUser, from, to).Return(mockedTwList, nil).Once()

	// call function
	u := NewTweetUsecase(
		tweetRepository,
		userRepository,
	)
	twList, err := u.Get(
		email,
		from,
		to,
	)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test_get", twList[0].Comment)
}

func TestTweetUsecase_Get_WhenUserRepositoryError(t *testing.T) {
	// mock
	userRepository := new(mocks.UserRepository)
	email := "test_get@email.com"
	mockUser := domain.User{Email: email}
	userRepository.On("Get", email).Return(
		mockUser,
		errors.New("some error"), // error occurred
	).Once()

	tweetRepository := new(mocks.TweetRepository)
	mockedTwList := make([]domain.Tweet, 0)
	mockedTwList = append(
		mockedTwList,
		domain.Tweet{Comment: "test_get"},
	)
	from := time.Now()
	to := time.Now()
	tweetRepository.On("Get", mockUser, from, to).Return(mockedTwList, nil).Once()

	// call function
	u := NewTweetUsecase(
		tweetRepository,
		userRepository,
	)
	_, err := u.Get(
		email,
		from,
		to,
	)
	if err == nil {
		t.Fatal("error should be returned when some error occurred")
	}
}

func TestTweetUsecase_Create(t *testing.T) {
	// mock
	userRepository := new(mocks.UserRepository)
	email := "test_get@email.com"
	mockUser := domain.User{Email: email}
	userRepository.On("Get", email).Return(
		mockUser,
		nil,
	)

	tweetRepository := new(mocks.TweetRepository)
	comment := "tweet_usecase_create_test"
	tweet := domain.Tweet{Comment: comment}
	tweetRepository.On("Create", mockUser, comment).Return(tweet, nil).Once()

	// call function
	u := NewTweetUsecase(
		tweetRepository,
		userRepository,
	)
	resTweet, err := u.Create(email, comment)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, comment, resTweet.Comment)
}

func TestTweetUsecase_Create_WhenUserRepositoryError(t *testing.T) {
	// mock
	userRepository := new(mocks.UserRepository)
	email := "test_get@email.com"
	mockUser := domain.User{Email: email}
	userRepository.On("Get", email).Return(
		mockUser,
		errors.New("some error"), // when error occurred
	)

	tweetRepository := new(mocks.TweetRepository)
	comment := "tweet_usecase_create_test"
	tweet := domain.Tweet{Comment: comment}
	tweetRepository.On("Create", mockUser, comment).Return(tweet, nil).Once()

	// call function
	u := NewTweetUsecase(
		tweetRepository,
		userRepository,
	)
	_, err := u.Create(email, comment)
	assert.NotNil(t, err)
}

func TestTweetUsecase_Delete(t *testing.T) {
	// mock
	userRepository := new(mocks.UserRepository)
	email := "test_test@email.com"
	mockUser := domain.User{Email: email}
	userRepository.On("Get", email).Return(
		mockUser,
		nil,
	)

	tweetRepository := new(mocks.TweetRepository)
	deleteID := uint(111)
	tweetRepository.On("Delete", mockUser, deleteID).Return(nil).Once()

	// call function
	u := NewTweetUsecase(
		tweetRepository,
		userRepository,
	)
	err := u.Delete(email, deleteID)
	assert.Nil(t, err)
}

func TestTweetUsecase_Delete_WhenUserRepositoryError(t *testing.T) {
	// mock
	userRepository := new(mocks.UserRepository)
	email := "test_test@email.com"
	mockUser := domain.User{Email: email}
	userRepository.On("Get", email).Return(
		mockUser,
		errors.New("some error"), // error occurred
	)

	tweetRepository := new(mocks.TweetRepository)
	deleteID := uint(111)
	tweetRepository.On("Delete", mockUser, deleteID).Return(nil).Once()

	// call function
	u := NewTweetUsecase(
		tweetRepository,
		userRepository,
	)
	err := u.Delete(email, deleteID)
	assert.NotNil(t, err)
}
