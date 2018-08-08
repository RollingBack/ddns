package main

import (
	"github.com/denverdino/aliyungo/dns"
	"github.com/denverdino/aliyungo/common"
	"local/ddns/ifconfig"
	"flag"
	"os"
	"fmt"
)

type Client struct {
	client *dns.Client
}

func (client *Client)GetDomainRecords(domain string) *dns.DescribeDomainRecordsNewResponse {
	paginate := common.Pagination{ PageNumber: 1, PageSize: 100}
	response, err := client.client.DescribeDomainRecordsNew(
		&dns.DescribeDomainRecordsNewArgs{
			DomainName: domain, Pagination: paginate})
	if err != nil {
		panic(err)
	}
	return response
}

func (client *Client)UpdateDomainRecord(recordID string, args dns.UpdateDomainRecordArgs) bool {
	_, err := client.client.UpdateDomainRecord(&args)
	if err != nil {
		return false
	}
	return true
}

func Required(key string, s string)  {
	if s == "" {
		os.Stderr.Write([]byte(fmt.Sprintf("%s is required param", key)))
		os.Exit(0)
	}
}

func main(){
	AccessKeyId := flag.String("i", "", "aliyun access key id")
	AccessKeySecret := flag.String("s", "", "aliyun key secret")
	domainName := flag.String("d", "", "")

	flag.Parse()

	Required("access_key_id", *AccessKeyId)
	Required("access_key_secret", *AccessKeySecret)
	Required("domain_name", *domainName)

	client := dns.NewClientNew(*AccessKeyId, *AccessKeySecret)
	client.SetRegionID(common.Hangzhou)
	myClient := Client{client:client}
	response := myClient.GetDomainRecords(*domainName)
	publicIP := ifconfig.GetPublicIP()
	for _, eachRecord := range response.DomainRecords.Record {
		myClient.UpdateDomainRecord(eachRecord.RecordId, dns.UpdateDomainRecordArgs{
			RecordId: eachRecord.RecordId,
			RR: eachRecord.RR,
			Type: eachRecord.Type,
			Value: publicIP})
	}

}
