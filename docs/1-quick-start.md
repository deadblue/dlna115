# 快速启动

## 安装

下载对应系统的安装包，解压到任意目录：

[https://github.com/deadblue/dlna115/releases/latest](https://github.com/deadblue/dlna115/releases/latest)

## 配置与运行

### 导出凭证

在终端中执行：

```bash
dlna115 login <credential.txt>
```

运行后将会看到一个二维码，请通过 115 手机客户端扫码登录，登录完成后凭证信息会导出到指定的文件中。

!!! Note

    Windows 用户需要在 [Windows Terminal](https://apps.microsoft.com/store/detail/windows-terminal/9N0DX20HK701) 中执行此命令，否则在终端中打印的二维码将无法被扫描。


### 配置服务

编辑 `config-quickstart.yaml` ，将第一步导出的凭证文件的路径填写到`source`字段中。

> 路径可使用完整路径或相对主程序的路径。

```yaml
storage:
  credential-source:
    type: file
    source: "凭证文件路径"
```

### 启动服务

执行如下命令，启动服务。

```bash
dlna115 daemon -c config-quickstart.yaml
```

## 访问服务

在处于同一局域网的终端上，如手机、平板或TV上，启动支持 UPnP 的播放器，即可在本地网络中扫描到名叫 115 的服务。

进入该服务即可浏览和播放 115 网盘上的视频。

推荐的播放器：

| Player \ Platform | macOS | Linux | Windows | Android | iOS |
| ----------------- | ----- | ----- | ------- | ------- | --- |
| VLC media player  | ✓     | ✓     | ✓       | ✓       | ✓   |
| OPlayer           | ✗     | ✗     | ✗       | ✓       | ✓   |
