<p align="center">
<img width="500" src="https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202207022209168.png"><br><br>
<a href="https://github.com/teamssix/cf/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/teamssix/cf"/></a>
<a href="https://github.com/teamssix/cf/releases"><img alt="GitHub releases" src="https://img.shields.io/github/release/teamssix/cf"/></a>
<a href="https://github.com/teamssix/cf/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"/></a>
<a href="https://github.com/teamssix/cf/releases"><img alt="Downloads" src="https://img.shields.io/github/downloads/teamssix/cf/total?color=brightgreen"/></a>
<a href="https://goreportcard.com/report/github.com/teamssix/cf"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/teamssix/cf"/></a>
<a href="https://twitter.com/intent/tweet/?text=CF%2C%20an%20amazing%20cloud%20exploitation%20framework%0Ahttps%3A%2F%2Fgithub.com%2Fteamssix%2Fcf%0A%23cloud%20%23security%20%23cloudsecurity%20%23cybersecurtiy"><img alt="tweet" src="https://img.shields.io/twitter/url?url=https://github.com/teamssix/cf" /></a>
<a href="https://twitter.com/teamssix"><img alt="Twitter" src="https://img.shields.io/twitter/follow/teamssix?label=Followers&style=social" /></a>
<a href="https://github.com/teamssix"><img alt="Github" src="https://img.shields.io/github/followers/TeamsSix?style=social" /></a><br></br>
<a href="README.md">中文</a> | English
</p>




---

CF is a cloud exploitation framework, It can facilitate the work of the red team after obtaining access key.

CF releases: [github.com/teamssix/cf/releases](https://github.com/teamssix/cf/releases)

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209232137924.png)

Current Supported Clouds:

- [x] Alibaba Cloud
- [x] Tencent Cloud
- [ ] AWS (Expected to be supported by October 14, 2022)
- [ ] Huawei Cloud (Expected to be supported by December 14, 2022)

For the scheduling of functions you can refer to: [github.com/teamssix/cf/discussions/130](https://github.com/teamssix/cf/discussions/130)

## Manual

For detailed manuals, please visit: [wiki.teamssix.com/cf](https://wiki.teamssix.com/cf)

> The manual currently supports Chinese only

[![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209081100625.png)](https://wiki.teamssix.com/cf)

## Cases

|               Title                | Version |                             URL                              |   Time    |
| :--------------------------------: | :-----: | :----------------------------------------------------------: | :-------: |
| 《记录一次平平无奇的云上攻防过程》 | v0.4.0  | [https://zone.huoxian.cn/d/2557](https://zone.huoxian.cn/d/2557) | 2022.9.14 |
|   《我用 CF 打穿了他的云上内网》   | v0.2.4  | [https://zone.huoxian.cn/d/1341-cf](https://zone.huoxian.cn/d/1341-cf) | 2022.7.13 |

## Easy to start

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737405.png)

> Here is the example of Alibaba Cloud, other more operations can be viewed in the user manual.

Configuration

```bash
cf config
```

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737407.png)

One-click access to current access credentials

```bash
cf alibaba perm
```

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737408.png)

One-click to take over the console

```bash
cf alibaba console
```

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737409.png)

One-click listing of cloud service resources with current access credentials

```bash
cf alibaba ls
```

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737410.png)

View the help information for the operation of the command executed by CF for the instance

```bash
cf alibaba ecs exec -h
```

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737411.png)

One-click command to execute proof of privilege for all instances

```bash
cf alibaba ecs exec -b
```

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737412.png)

One-click access to temporary access credential data in instances

```bash
cf alibaba ecs exec -m
```

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737413.png)

One-click download of OSS object storage data

```bash
cf alibaba oss obj get
```

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737414.png)

One-Click Upgrade CF Version

```bash
cf upgrade
```

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209071737416.png)

If it feels good, maybe you can give me a Star ~

## Contributor

Thank you for your contributions to CF, A note on contributions: [CONTRIBUTING](https://github.com/teamssix/cf/blob/main/CONTRIBUTING.md)

<div align=center>
<table>
    <tr>
        <td align="center"><a href="https://github.com/teamssix"><img alt="TeamsSix"
                                src="https://avatars.githubusercontent.com/u/49087564?v=4" style="width: 100px;"/><br/>TeamsSix</a></td>
        <td align="center"><a href="https://github.com/Amzza0x00"><img alt="Amzza0x00"
                                src="https://avatars.githubusercontent.com/u/32904523?v=4"  style="width: 100px;"/><br/>Amzza0x00</a></td>
        <td align="center"><a href="https://github.com/Esonhugh"><img alt="Esonhugh"
                                src="https://avatars.githubusercontent.com/u/32677240?v=4"  style="width: 100px;"/><br/>Esonhugh</a></td>
        <td align="center"><a href="https://github.com/Dawnnnnnn"><img alt="Dawnnnnnn"
                                src="https://avatars.githubusercontent.com/u/24506421?v=4"  style="width: 100px;"/><br/>Dawnnnnnn</a></td>
        <td align="center"><a href="https://github.com/Belos-pretender"><img alt="Belos-pretender"
                                src="https://avatars.githubusercontent.com/u/52148409?v=4"  style="width: 100px;"/><br/>Belos-pretender</a></td>
</table>
</div>

<div align=center><a href="https://github.com/teamssix"><img src="https://repobeats.axiom.co/api/embed/30b8de6c059cbe83fe0ba44fff91136270a39ab9.svg"></a></div>



## 404Starlink

<img src="https://github.com/knownsec/404StarLink/raw/master/Images/logo.png" width="40%">

CF has joined [404Starlink](https://github.com/knownsec/404StarLink)

## More

If you are interested in cloud security, you can see my other project [Awesome Cloud Security](https://github.com/teamssix/awesome-cloud-security) , many cloud security resources are included here.

If these cloud security resources are still not enough for you, check out my [cloud security knowledge base](https://wiki.teamssix.com/)), where I have a lot of notes and articles in the direction of cloud security.

Finally, the following is my personal wechat official accounts, welcome to follow ~

<div align=center><a href="https://github.com/teamssix"><img width="700" src="https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202204152148071.png"></a></div>

If you would like to work with me on this, you can join the team by sending your resume to admin@wgpsec.org.

![](https://cdn.jsdelivr.net/gh/teamssix/BlogImages/imgs/202209151716790.png)

<div align=center><a href="https://github.com/teamssix"><img src="https://api.star-history.com/svg?repos=teamssix/cf&type=Timeline"></a></div>



## Warning

* This tool can only be used in legal scenarios and is strictly forbidden to be used in illegal scenarios.
* The risks involved in this tool are the responsibility of the tenant and not the cloud providers.

