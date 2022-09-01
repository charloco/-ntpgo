# 介绍

实现一个最简易的ntp客户端,减少对外部的依赖,简化部署, ntpgo不参与时间的交换, 只是从server同步时间

# 部署

参照 make test_install, 只需要ntpgo install --server <ntp server>, 填入想要的ntp server地址

# 启动

install之后, ntpgo.service已经部署和enable了,只需要ntpgo restart即可启动, 或者systemctl restart ntpgo

# 默认硬编码的配置

* 程序部署到/usr/bin/ntpgo
* 配置文件为/etc/ntpgo/ntp.yaml
* 5分钟同步一次时间
* 如果本地时间和ntpserver的时间差距超过2s, 就重新设置本地时间
