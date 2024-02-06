package maps_test

import (
	"errors"
	"goModDemo/maps"
	mock_maps "goModDemo/maps/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
)

func Test_wrapperFunc(t *testing.T) {
	type args struct {
		db *mock_maps.MockiPostgresDB
	}
	add1 := maps.Home{State: "NC", Zip: 123}
	tests := []struct {
		name string
		args func(*args)
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: func(a *args) {
				a.db.EXPECT().PostRecord(add1).Return(
					nil, errors.New("insert error"),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			a := args{
				db: mock_maps.NewMockiPostgresDB(ctrl),
			}
			if tt.args != nil {
				tt.args(&a)
			}

			maps.WrapperFunc(a.db, add1)
		})
	}
}

type sqlmockResult struct{}

func (sqlmockResult) LastInsertId() (int64, error) {
	return 1, nil
}

func (sqlmockResult) RowsAffected() (int64, error) {
	return 1, nil
}

// RowsAffected returns the number of rows affected by an
// update, insert, or delete. Not every database or database
// driver may support this.

func Test_wrapperFunc2(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db2 := mock_maps.NewMockiPostgresDB(ctrl)

	add1 := maps.Home{State: "NC", Zip: 123}

	db2.EXPECT().PostRecord(add1).Return(
		sqlmockResult{}, nil,
	)

	r, err := maps.WrapperFunc(db2, add1)
	if err != nil {

		t.Errorf("expected error but didnt get any error: ")
	}
	p, _ := r.RowsAffected()
	if p != 1 {
		t.Errorf("expectd 1  ")

	}

}
