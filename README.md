# img_bed-Go

这个程序使用golang实现，使用到了gin还有github.com/chai2010/webp

本程序不使用数据库，大小只有10MB，只接受jpg和png图片，上传后会把图片保存为webp格式(更优的图像数据压缩算法)，方便浏览器传输

## 安装

```bash
mkdir img_bed
cd img_bed
wget https://github.com/meowrain/img-bed-Go/releases/download/V1.0.0/img_bed-linux-amd64
chmod +x img_bed-linux-amd64

```

然后在这个目录下创建`config`目录

```bash
mkdir config
vim config/config.yaml
```

配置项共三项，domain项为你的域名(记住不要丢了https://)
port是开放的端口
auth是用来上传验证的，需要自行设置
![](https://static.meowrain.cn/i/2024/06/23/lpR8Wz1719126971898183265.webp)

## picgo配置

> https://github.com/Molunerfinn/PicGo/releases

先根据自己的系统版本下一个picgo，然后安装上（记住需要node环境，不然没办法安装插件)

然后我们到插件设置里安装插件

![](https://static.meowrain.cn/i/2024/06/23/ntfa3s1719127175848529110.webp)

![](https://static.meowrain.cn/i/2024/06/23/couQCh1719127201955181776.webp)

![](https://static.meowrain.cn/i/2024/06/23/b2lneF1719127334738136622.webp)

API地址填 https://yourdomain/upload
Post参数名填 file
自定义Body填 {"token":"your token"}

> 这里的token就是咱们上面配置的yaml文件里面的token

## 配置typora粘贴上传

![](https://static.meowrain.cn/i/2024/06/23/MPtE741719127389577367413.webp)
![](https://static.meowrain.cn/i/2024/06/23/y6cAwZ1719127411798653640.webp)
![](https://static.meowrain.cn/i/2024/06/23/hES8xj1719127433241819923.webp)

这样就能粘贴图片的同时上传图片到服务器了

## 配置开机自启动

```bash
sudo vim /etc/systemd/system/img-bed.service
```

```
[Unit]
Description=Image Bed Service
After=network.target

[Service]
ExecStart=可执行文件路径
Restart=always
User=your_user  # 替换为运行该服务的用户
Group=your_group  # 替换为运行该服务的用户组
WorkingDirectory=可执行文件路径

[Install]
WantedBy=multi-user.target

```

重新加载 systemd 配置：
保存并关闭文件后，重新加载 systemd 配置：

```bash
sudo systemctl daemon-reload
```

启用自启动

```bash
sudo systemctl enable img-bed.service
sudo systemctl start img-bed.service
```

检查服务状态：
确认服务是否正常启动：

```bash
sudo systemctl status img-bed.service
```