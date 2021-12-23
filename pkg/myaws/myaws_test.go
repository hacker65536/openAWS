package myaws

import (
	"testing"
)

func TestListeneRule(t *testing.T) {
	// https://docs.aws.amazon.com/ja_jp/general/latest/gr/aws-arns-and-namespaces.html
	str := "arn:aws:elasticloadbalancing:us-west-2:123456789012:listener-rule/app/my-load-balancer/50dc6c495c0c9188/f2f7dc8efc522ab2/9683b2d02a6cabee"

	want := "https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#ELBRules:type=app;loadBalancerName=my-load-balancer;loadBalancerId=50dc6c495c0c9188;listenerId=f2f7dc8efc522ab2;accountId=123456789012"
	got := ListeneRule(str)

	if got != want {
		t.Errorf("ListenerRule()  \ngot=\n%v \nwant=\n%v", got, want)
	}
}

func TestOpenS3(t *testing.T) {
	a := []struct {
		input string
		want  string
	}{
		{
			input: "cf-templates-xxxxxxxxxxx-ap-northeast-1",
			want:  "https://s3.console.aws.amazon.com/s3/buckets/cf-templates-xxxxxxxxxxx-ap-northeast-1?tab=objects",
		},
		{
			input: "cf-templates-xxxxxxxxx-ap-northeast-1/20130701Ac-cf2.txt",
			want:  "https://s3.console.aws.amazon.com/s3/object/cf-templates-xxxxxxxxx-ap-northeast-1?prefix=20130701Ac-cf2.txt",
		},
	}

	for _, v := range a {
		got := OpenS3(v.input)

		if got != v.want {
			t.Errorf("OpenS3()  \ninput=%v\n got=%v\n want=%v\n", v.input, got, v.want)
		}
	}
}
