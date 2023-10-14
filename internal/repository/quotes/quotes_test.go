package quotes

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"

	"random-quotes/internal/entity"
)

func TestRepository_GetQuotes(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(fs afero.Fs)
		want    []entity.Quotes
		wantErr bool
	}{
		{
			name:   "Return quotes",
			fields: fields{fs: afero.NewMemMapFs()},
			want: []entity.Quotes{
				{
					Quotes: "You can observe a lot just by watching.",
					Author: "Yogi Berra",
				},
			},
			prepare: func(fs afero.Fs) {
				err := fs.Mkdir("data", 0755)
				if err != nil {
					panic(err)
				}
				err = afero.WriteFile(fs, "data/quotes.json",
					[]byte("[{\"quotes\": \"You can observe a lot just by watching.\", \"author\": \"Yogi Berra\"}]"), 0644)
				if err != nil {
					panic(err)
				}
			},
		},
		{
			name:    "Return err",
			fields:  fields{fs: afero.NewMemMapFs()},
			prepare: func(fs afero.Fs) {},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &Repository{
				fs: tt.fields.fs,
			}

			tt.prepare(repository.fs)
			got, err := repository.GetQuotes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQuotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuotes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
