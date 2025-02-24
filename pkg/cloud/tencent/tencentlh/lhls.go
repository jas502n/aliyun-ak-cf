package tencentlh

import (
	"encoding/json"
	"github.com/teamssix/cf/pkg/cloud/tencent/tencentcvm"
	"github.com/teamssix/cf/pkg/util/errutil"
	lh "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/teamssix/cf/pkg/cloud"
	"github.com/teamssix/cf/pkg/util/cmdutil"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

var (
	header = []string{"序号 (SN)", "实例ID (Instance ID)", "实例名称 (Instance Name)", "系统名称 (OS Name)", "系统类型 (OS Type)", "状态 (Status)", "私有 IP (Private IP)", "公网 IP (Public IP)", "区域 ID (Region ID)"}
)

type Instances struct {
	InstanceId       string
	InstanceName     string
	OSName           string
	OSType           string
	Status           string
	PrivateIpAddress string
	PublicIpAddress  string
	RegionId         string
}

func DescribeInstances(region string, running bool, SpecifiedInstanceID string) []Instances {
	var out []Instances
	request := lh.NewDescribeInstancesRequest()
	if running {
		request.Filters = []*lh.Filter{
			{
				Name:   common.StringPtr("instance-state"),
				Values: common.StringPtrs([]string{"RUNNING"}),
			},
		}
	}
	if SpecifiedInstanceID != "all" {
		request.InstanceIds = common.StringPtrs([]string{SpecifiedInstanceID})
	}
	response, err := LHClient(region).DescribeInstances(request)
	errutil.HandleErr(err)
	InstancesList := response.Response.InstanceSet
	log.Debugf("正在 %s 区域中查找实例 (Looking for instances in the %s region)", region, region)
	if len(InstancesList) != 0 {
		log.Debugf("在 %s 区域下找到 %d 个实例 (Found %d instances in %s region)", region, len(InstancesList), len(InstancesList), region)
		var (
			PrivateIpAddressList []string
			PublicIpAddressList  []string
			PrivateIpAddress     string
			PublicIpAddress      string
			OSType               string
		)
		for _, v := range InstancesList {
			for _, m := range v.PrivateAddresses {
				PrivateIpAddressList = append(PrivateIpAddressList, *m)
			}
			for _, m := range v.PublicAddresses {
				PublicIpAddressList = append(PublicIpAddressList, *m)
			}
			a, _ := json.Marshal(PrivateIpAddressList)
			if len(PrivateIpAddressList) == 1 {
				PrivateIpAddress = PrivateIpAddressList[0]
			} else {
				PrivateIpAddress = string(a)
			}
			b, _ := json.Marshal(PublicIpAddressList)
			if len(PublicIpAddressList) == 1 {
				PublicIpAddress = PublicIpAddressList[0]
			} else {
				PublicIpAddress = string(b)
			}
			newOSname := strings.Split(*v.OsName, " ")[0]
			if find(tencentcvm.LinuxSet, newOSname) {
				OSType = "linux"
			} else {
				OSType = "windows"
			}
			errutil.HandleErr(err)
			obj := Instances{
				InstanceId:       *v.InstanceId,
				InstanceName:     *v.InstanceName,
				OSName:           *v.OsName,
				OSType:           OSType,
				Status:           *v.InstanceState,
				PrivateIpAddress: PrivateIpAddress,
				PublicIpAddress:  PublicIpAddress,
				RegionId:         *v.Zone,
			}
			out = append(out, obj)
		}
	}
	return out
}

func ReturnInstancesList(region string, running bool, specifiedInstanceID string) []Instances {
	var InstancesList []Instances
	var Instance []Instances
	if region == "all" {
		for _, j := range GetLHRegions() {
			region := *j.Region
			Instance = DescribeInstances(region, running, specifiedInstanceID)
			for _, i := range Instance {
				InstancesList = append(InstancesList, i)
			}
		}
	} else {
		InstancesList = DescribeInstances(region, running, specifiedInstanceID)
	}
	return InstancesList
}

func PrintInstancesListRealTime(region string, running bool, specifiedInstanceID string) {
	InstancesList := ReturnInstancesList(region, running, specifiedInstanceID)
	var data = make([][]string, len(InstancesList))
	for i, o := range InstancesList {
		SN := strconv.Itoa(i + 1)
		data[i] = []string{SN, o.InstanceId, o.InstanceName, o.OSName, o.OSType, o.Status, o.PrivateIpAddress, o.PublicIpAddress, o.RegionId}
	}
	var td = cloud.TableData{Header: header, Body: data}
	if len(data) == 0 {
		log.Info("未发现 LH 实例 (No LH instances found)")
	} else {
		Caption := "LH 资源 (LH resources)"
		cloud.PrintTable(td, Caption)
	}
	cmdutil.WriteCacheFile(td, "tencent", "lh", region, specifiedInstanceID)
}

func PrintInstancesListHistory(region string, running bool, specifiedInstanceID string) {
	cmdutil.PrintECSCacheFile(header, region, specifiedInstanceID, "tencent", "LH", running)
}

func PrintInstancesList(region string, running bool, specifiedInstanceID string, lhFlushCache bool) {
	if lhFlushCache {
		PrintInstancesListRealTime(region, running, specifiedInstanceID)
	} else {
		PrintInstancesListHistory(region, running, specifiedInstanceID)
	}
}
