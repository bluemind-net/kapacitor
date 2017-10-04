package pipeline

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestWindowNode_MarshalJSON(t *testing.T) {
	type fields struct {
		Period         time.Duration
		Every          time.Duration
		AlignFlag      bool
		FillPeriodFlag bool
		PeriodCount    int64
		EveryCount     int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "all fields set",
			fields: fields{
				Period:         time.Hour,
				Every:          time.Minute,
				AlignFlag:      true,
				FillPeriodFlag: true,
				PeriodCount:    1,
				EveryCount:     2,
			},
			want: `{"typeOf":"window","ID":"0","period":3600000000000,"every":60000000000,"align":true,"fillPeriod":true,"periodCount":1,"everyCount":2}`,
		},
		{
			name: "only period and every",
			fields: fields{
				Period: time.Hour,
				Every:  time.Minute,
			},
			want: `{"typeOf":"window","ID":"0","period":3600000000000,"every":60000000000,"align":false,"fillPeriod":false,"periodCount":0,"everyCount":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := newWindowNode()
			w.Period = tt.fields.Period
			w.Every = tt.fields.Every
			w.AlignFlag = tt.fields.AlignFlag
			w.FillPeriodFlag = tt.fields.FillPeriodFlag
			w.PeriodCount = tt.fields.PeriodCount
			w.EveryCount = tt.fields.EveryCount
			g, err := json.Marshal(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("WindowNode.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := string(g)
			if got != tt.want {
				t.Errorf("WindowNode.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestWindowNode_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *WindowNode
		wantErr bool
	}{
		{
			name:  "all fields set",
			input: `{"typeOf":"window","ID":"0","period":3600000000000,"every":60000000000,"align":true,"fillPeriod":true,"periodCount":1,"everyCount":2}`,
			want: &WindowNode{
				Period:         time.Hour,
				Every:          time.Minute,
				AlignFlag:      true,
				FillPeriodFlag: true,
				PeriodCount:    1,
				EveryCount:     2,
			},
		},
		{
			name:  "only period and every",
			input: `{"typeOf":"window","ID":"0","period":3600000000000,"every":60000000000,"align":false,"fillPeriod":false,"periodCount":0,"everyCount":0}`,
			want: &WindowNode{
				Period: time.Hour,
				Every:  time.Minute,
			},
		},
		{
			name:  "set id correctly",
			input: `{"typeOf":"window","ID":"5","period":3600000000000,"every":60000000000,"align":false,"fillPeriod":false,"periodCount":0,"everyCount":0}`,
			want: &WindowNode{
				chainnode: chainnode{
					node: node{
						id: 5,
					},
				},
				Period: time.Hour,
				Every:  time.Minute,
			},
		},
		{
			name:    "invalid data",
			input:   `{"typeOf":"window","ID":"0", "period": "invalid"}`,
			wantErr: true,
			want:    &WindowNode{},
		},
		{
			name:    "invalid node type",
			input:   `{"typeOf":"invalid","ID":"0"}`,
			wantErr: true,
			want:    &WindowNode{},
		},
		{
			name:    "invalid id type",
			input:   `{"typeOf":"window","ID":"invalid"}`,
			wantErr: true,
			want:    &WindowNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WindowNode{}
			err := json.Unmarshal([]byte(tt.input), w)
			if (err != nil) != tt.wantErr {
				t.Errorf("WindowNode.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(w, tt.want) {
				t.Errorf("WindowNode.UnmarshalJSON() =\n%#+v\nwant\n%#+v", w, tt.want)
			}
		})
	}

}
