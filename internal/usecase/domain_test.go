package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/rsdmike/rps/internal/usecase"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func translation(t *testing.T) (*usecase.DomainsUseCase, *MockDomain) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockDomain(mockCtl)

	translation := usecase.NewDomainsUseCase(repo)

	return translation, repo
}

func TestHistory(t *testing.T) {
	t.Parallel()

	translation, repo := translation(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetCount(context.Background(), "").Return(0, nil)
			},
			res: 0,
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetCount(context.Background(), "").Return(0, errInternalServErr)
			},
			res: 0,
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := translation.GetCount(context.Background(), "")

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
