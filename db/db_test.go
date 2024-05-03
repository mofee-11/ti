package db

import "testing"

func TestCheck(t *testing.T) {
	cases := []struct {
		in struct {
			sub  string
			data map[string]interface{}
		}
		want bool
	}{
		{
			in: struct {
				sub  string
				data map[string]interface{}
			}{
				sub: "value",
				data: map[string]interface{}{
					"name": "value",
				},
			},
			want: true,
		},
		{
			in: struct {
				sub  string
				data map[string]interface{}
			}{
				sub: "name",
				data: map[string]interface{}{
					"name": "value",
				},
			},
			want: false,
		},
		{
			in: struct {
				sub  string
				data map[string]interface{}
			}{
				sub: "a",
				data: map[string]interface{}{
					"tag": []string{
						"a",
						"b",
						"c",
					},
				},
			},
			want: true,
		},
		{
			in: struct {
				sub  string
				data map[string]interface{}
			}{
				sub: "deep",
				data: map[string]interface{}{
					"tree": []interface{}{
						"a",
						"b",
						map[string]interface{}{"lili": "deep"},
					},
				},
			},
			want: true,
		},
		{
			in: struct {
				sub  string
				data map[string]interface{}
			}{
				sub: "吃",
				data: map[string]interface{}{
					"tag": []interface{}{
						"日志/吃",
					},
				},
			},
			want: true,
		},
	}

	for _, c := range cases {
		got := check(c.in.data, c.in.sub)
		if got != c.want {
			t.Errorf("check(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
