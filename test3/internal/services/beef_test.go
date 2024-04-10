package services

import (
	"7solutionstest3/internal/models"
	mockerys "7solutionstest3/internal/repositories/adapters/rest/mocks"
	"errors"
	"reflect"
	"testing"

	"golang.org/x/net/context"
)

func TestBeefSummaryService_BeefSummary(t *testing.T) {
	ctx := context.Background()
	beefAPIMock := mockerys.NewBeefAPICaller(t)

	type fields struct {
		beefAPICaller BeefAPICaller
	}
	tests := []struct {
		name    string
		fields  fields
		ctx     context.Context
		want    *models.BeefResponse
		wantErr bool
		mock    func()
	}{
		{
			name:   "Test case Happy flow",
			fields: fields{beefAPICaller: beefAPIMock},
			ctx:    ctx,
			want: &models.BeefResponse{
				Beef: map[string]int{
					"t-bone":   4,
					"fatback":  1,
					"pastrami": 1,
					"pork":     1,
					"meatloaf": 1,
					"jowl":     1,
					"enim":     1,
					"bresaola": 1,
				},
			},
			wantErr: false,
			mock: func() {
				mockData := `Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.`
				beefAPIMock.EXPECT().FetchBeefData(ctx).Return(&mockData, nil).Times(1)
			},
		},
		{
			name:    "Test case Failed",
			fields:  fields{beefAPICaller: beefAPIMock},
			ctx:     ctx,
			want:    nil,
			wantErr: true,
			mock: func() {
				beefAPIMock.EXPECT().FetchBeefData(ctx).Return(nil, errors.New("error fetch beef data, status code 500")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BeefSummaryService{
				beefAPICaller: tt.fields.beefAPICaller,
			}

			tt.mock()

			got, err := s.BeefSummary(tt.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeefSummaryService.BeefSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeefSummaryService.BeefSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}
