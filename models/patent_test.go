package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPatent_ExtractClaims(t *testing.T) {
	tests := []struct {
		name    string
		claims  string
		want    []string
		wantErr bool
	}{
		{
			name: "valid claims",
			claims: `[
				{"num": "1", "text": "A device comprising..."},
				{"num": "2", "text": "The device of claim 1..."}
			]`,
			want:    []string{"A device comprising...", "The device of claim 1..."},
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			claims:  `invalid json`,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty claims",
			claims:  `[]`,
			want:    []string{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Patent{
				Claims: tt.claims,
			}
			got, err := p.ExtractClaims()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
