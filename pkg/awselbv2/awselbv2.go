package awselbv2

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

func ListeneRule(arnstr string) string {
	if !arn.IsARN(arnstr) {
		return "not arn"
	}

	parn, _ := arn.Parse(arnstr)
	res := strings.Split(parn.Resource, "/")

	url := "https://"
	url += parn.Region
	url += ".console.aws.amazon.com/ec2"
	url += "/v2/home?region="
	url += parn.Region
	url += "#ELBRules:type=app;loadBalancerName="
	url += res[2]
	url += ";loadBalancerId="
	url += res[3]
	url += ";listenerId="
	url += res[4]
	url += ";accountId="
	url += parn.AccountID

	//fmt.Println(url)
	return url

}
