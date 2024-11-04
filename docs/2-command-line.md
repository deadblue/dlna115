# 命令行

## 命令行格式

```bash
dlna115 <command> <command-arguments>
```

`command` 支持如下命令：

* `login`
* `daemon`

## 命令说明

### login | 登录

登录命令用来模拟桌面客户端的扫码登录，并将用户的凭证信息导出到文件中。

```bash
dlna115 login [-p <platform>] [-s <secret>] [credential-file]
```

参数说明：

**`-p`/`-platform`**

模拟登录的平台，支持如下平台：

- `web`（默认值）
- `android`
- `ios`
- `tv`
- `wechat`
- `alipay`
- `qandroid`

!!! Note

    115 限制帐号在每个平台最多只能登录一次，请选择一个自己不常用的平台。

---

**`-s`/`-secret`**

加密凭证文件的密钥。默认为空，表示不加密。

---

**`credential-file`**

保存凭证的文件路径。未传入此参数时，凭证内容会打印在终端中。

---

### daemon | 启动服务

启动 DLNA MediaServer 与 SSDP 服务。

```bash
dlna115 daemon -c <config-file.yaml>
```

参数说明：

**`-c`/`-config`**

配置文件路径。配置文件格式请参见：[配置文件](3-configuration.md)。

---

## 注册为系统服务

=== "Linux"

    ```ini
    [Unit]
    Description=115 DLNA Service
    After=network-online.target

    [Service]
    DynamicUser=yes
    ExecStart=/usr/local/bin/dlna115 daemon -c /usr/local/etc/dlna115/config.yaml
    Restart=on-failure

    [Install]
    WantedBy=multi-user.target
    ```

=== "mac"

    ```xml
    <?xml version="1.0" encoding="UTF-8"?>
    <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
    <plist version="1.0">
      <dict>
        <key>UserName</key>
        <string>nobody</string>
        <key>Umask</key>
        <string>0022</string>
        <key>KeepAlive</key>
        <true/>
        <key>RunAtLoad</key>
        <true/>
        <key>Label</key>
        <string>dlna115</string>
        <key>ProgramArguments</key>
        <array>
          <string>/usr/local/bin/dlna115</string>
          <string>daemon</string>
          <string>-c</string>
          <string>/usr/local/etc/dlna115/config.yaml</string>
        </array>
      </dict>
    </plist>
    ```