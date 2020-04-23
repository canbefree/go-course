package impl

import (
	"testing"
)

func TestRedisCache_Set(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name    string
		cache   *RedisCache
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				key: "xieyt01",
				val: "123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := &RedisCache{}
			if err := cache.Set(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("RedisCache.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisCache_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		cache   *RedisCache
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				key: "xieyt01",
			},
			want: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := &RedisCache{}
			got, err := cache.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RedisCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
