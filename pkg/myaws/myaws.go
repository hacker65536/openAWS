package myaws

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	log "github.com/sirupsen/logrus"
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

	// fmt.Println(url)
	return url
}

func OpenS3(s3args []string) string {
	u, _ := url.Parse("https://s3.console.aws.amazon.com/s3/")
	q := u.Query()
	u.Path = "s3/"

	s3path := s3args[0]
	r := regexp.MustCompile(`/`)
	str := r.Split(s3path, 2)
	log.WithFields(
		log.Fields{
			"s3args": s3args,
			"str":    str,
		}).Debug()
	if len(str) > 1 {
		if str[1][len(str[1])-1:len(str[1])] == "/" {
			// buckets
			u.Path += "buckets/" + str[0]
		} else {
			// objects
			u.Path += "object/" + str[0]
		}
		q.Set("prefix", str[1])
	} else {
		// buckets
		u.Path += "buckets/" + str[0]
		q.Set("tab", "objects")
	}

	if len(s3args) > 1 {
		q.Set("tab", s3args[1])
	}
	u.RawQuery = q.Encode()
	return u.String()
}
