package awselbv2

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
