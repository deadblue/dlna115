# 命令行说明

## 命令行格式

```bash
dlna115 <command> <command-arguments>
```

`command` 支持如下命令：

* `login`
* `daemon`

## login | 登录

登录命令用来模拟桌面客户端的扫码登录，并将用户的凭证信息导出到文件中。

```bash
dlna115 login [-p <platform>] [-s <secret>] [credential-file]
```

参数说明：

**`-p`/`-platform`**

模拟登录的平台，支持 `mac`，`windows` 和 `linux`（默认值）。

!!! Note

    115 限制帐号在每个平台最多只能登录一次。
    
    如果用户平常使用 115 的 Linux 桌面客户端，请通过 -p 参数指定一个自己不常用的平台，否则桌面客户端的登录将被顶掉。

---

**`-s`/`-secret`**

加密凭证文件的密钥。默认为空，表示不加密。

---

**`credential-file`**

保存凭证的文件路径。

---

## daemon | 启动服务

启动 DLNA MediaServer 与 SSDP 服务。

```bash
dlna115 daemon -c <config-file.yaml>
```

参数说明：

**`-c`/`-config`**

配置文件路径。配置文件格式请参见：[配置文件](3-configuration.md)。

---


### 注册为系统服务

=== "Linux"

    ```ini
    [Unit]
    Description=115 DLNA Service
    After=network.target nss-lookup.target

    [Service]
    User=nobody
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