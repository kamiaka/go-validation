package is

import "testing"

func TestIsIPv4CIDR(t *testing.T) {
	cases := []struct {
		cidr string
		want bool
	}{
		{
			cidr: "192.168.1.1/32",
			want: true,
		},
		{
			cidr: "192.168.1.1",
			want: false,
		},
		{
			cidr: "2001:db8::/32",
			want: false,
		},
		{
			cidr: "192.168.1.1/abcde",
			want: false,
		},
	}
	for i, tc := range cases {
		got := isIPv4CIDR(tc.cidr)
		if got != tc.want {
			t.Errorf("#%d: isIPv4CIDR(%#v) == %#v, want %#v", i, tc.cidr, got, tc.want)
		}
	}
}
