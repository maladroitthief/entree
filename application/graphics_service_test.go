package application_test

import (
	"image"
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/sprite"
)

func TestNewGraphicsService(t *testing.T) {
	type args struct {
		logger logs.Logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				logger: &logger{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := application.NewGraphicsService(tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGraphicsService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_graphicsService_LoadSpriteSheet(t *testing.T) {
	type fields struct {
		log logs.Logger
	}
	type args struct {
		ss sprite.SpriteSheet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewGraphicsService(tt.fields.log)
			if err != nil {
				t.Errorf("GraphicsService.LoadSpriteSheet() failed to create a NewGraphicsService()")
				return
			}

			svc.LoadSpriteSheet(tt.args.ss)
		})
	}
}

func Test_graphicsService_GetSprite(t *testing.T) {
	type fields struct {
		log logs.Logger
	}
	type args struct {
		sheetName  string
		spriteName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    image.Rectangle
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewGraphicsService(tt.fields.log)
			if err != nil {
				t.Errorf("GraphicsService.GetSprite() failed to create a NewGraphicsService()")
				return
			}

			got, err := svc.GetSprite(tt.args.sheetName, tt.args.spriteName)
			if (err != nil) != tt.wantErr {
				t.Errorf("graphicsService.GetSprite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("graphicsService.GetSprite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_graphicsService_GetSpriteSheet(t *testing.T) {
	type fields struct {
		log logs.Logger
	}
	type args struct {
		sheetName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sprite.SpriteSheet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewGraphicsService(tt.fields.log)
			if err != nil {
				t.Errorf("GraphicsService.GetSpriteSheet() failed to create a NewGraphicsService()")
				return
			}

			got, err := svc.GetSpriteSheet(tt.args.sheetName)
			if (err != nil) != tt.wantErr {
				t.Errorf("graphicsService.GetSpriteSheet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("graphicsService.GetSpriteSheet() = %v, want %v", got, tt.want)
			}
		})
	}
}
