# 配置文件

## 完整格式

```yaml
# 115 storage settings
storage:

  # Source to get 115 credential.
  credential-source:
    # Credential source type: "file" or "url"
    type: "file"
    # For "file" type, source should be full path of the file.
    # For "url" type, source should be the URL.
    source: "/path/to/credential.txt"
    # Secret key to decrypt credential.
    secret: ""

  # Top folders which will be shown under root directory.
  top-folders:

    # A virtual folder which contains stared files.
    - type: "star"
      # You can set name to override default name.
      name: "Favorites"

    # A virtual folder which contains files with given label.
    - type: "label"
      # Target must be set, and should be a label name.
      target: "label-name"
      # Default value is the label name.
      name: ""

    # A directory on storage.
    - type: "dir"
      # Target must be set, and should be a the directory path.
      target: "parent/child1/child2"
      # Default value is the base name of directory.
      name: ""

  # Do not use HLS for video.
  # Set this to true if you UPnP player does not support HLS.
  disable-hls: false

# Media server settings
media:
  # Listening port, default value is 8115.
  port: 8115
  # Unique ID, server will automatically generate one if you do not set.
  uuid: ""
  # Server friendly name, default value is "115"
  name: ""

```

## 配置项说明

### storage | 115网盘配置

**disable-hls: `Boolean`**

> 是否禁用 HLS，默认值 `false`。
> 
> 当使用非 web 平台的凭证时，此项强制为 `true`。

> 禁用 HLS 后， MediaServer 将从网盘下载视频内容，并推送给客户端，对 CPU 和内存的占用会增大。

#### credential-source | 凭证来源

**type: `String`**

> 凭证来源类型：
>
> - file: 本地路径。
> - url: 网络地址。

**source: `String`**

> 凭证来源地址
>
> - 当 type="file" 时，应该为一个本地的路径。
> - 当 type="url" 时，应该为一个 HTTP 或 HTTPS 的 URL。

**secret: `String`**

> 凭证密钥（可选）
>
> 解密凭证信息的密钥，仅当凭证文件被加密时使用。

#### top-folders | 顶级目录列表

本配置项用于定义 DLNA 服务返回的顶级目录列表，若不配置，则只返回收藏文件。

本配置项为一个数组，每个数组元素是一个顶级目录的配置，它包含如下属性。

**type: `String`**

> 目录类型
>
> - dir: 一个网盘目录
> - label: 标签
> - star: 收藏

**target: `String`**

> 目录来源
>
> - 当 type="dir" 时，为网盘上目录的完整路径，如：“我的接收/视频”
> - 当 type="label" 时，为标签名字
> - 当 type="star" 时，此项将被忽略。

**name: `String`**

> 目录显示名称（可选）
>
> 用于自定义顶级目录的显示名称，未设置时，将显示目录原始的名称。

### media | 媒体服务配置

**port: `Integer`**

> 可选。媒体服务监听端口，默认 8115。

**uuid: `String`**

> 可选。媒体服务 UUID。未设置时，会在每次启动时自动生成，建议设置。

**name: `String`**

> 可选。媒体服务名称，将显示在客户端的服务列表中，默认为 "115"。
