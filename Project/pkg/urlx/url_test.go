package urlx

import "testing"

func TestGetBasePath(t *testing.T) {
	type args struct {
		targetUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "基本用例",
			args: args{
				targetUrl: "https://mbd.baidu.com/newspage/data/landingsuper?context=%7B%22nid%22%3A%22news_9278476480154871491%22%7D&n_type=-1&p_from=-1",
			},
			want:    "landingsuper",
			wantErr: false,
		},
		{
			name: "无效的url示例",
			args: args{
				targetUrl: "/xxx/123",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBasePath(tt.args.targetUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
