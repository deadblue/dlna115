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

## storage | 115网盘配置

**disable-hls**

是否禁用 HLS，默认值 `false`。

当使用不支持 HLS 的播放器时，可将此选项设置为 `true`。

禁用 HLS 后， MediaServer 将从网盘下载视频内容，并推送给客户端，对 CPU 和内存的占用会增大。

### credential-source | 凭证来源



### top-folders | 顶级目录



