


## 分块上传

### 流程
1. api发送分块上传请求，传入文件的sha1值和文件大小
2. rpc生成分块上传的初始化信息，api得到uploadID、分块大小、分块数量信息，rpc服务端会将这些信息写入redis
3. api返回初始化信息给前端，等待前端上传分块
4. 前端上传分块文件：将文件分割成小文件，每个文件对应有index，通过api上传，api检查该index是否已经上传or没有上传
5. 全部上传结束，调用文件分块合并：检查分块完整性、合并
6. 判断文件保存方法：本地、cos、minio，调用对应rpc接口


```mermaid
graph TD
    Start[开始] --> Init[初始化请求<br/>(API层)]
    Init --> Generate[生成上传信息<br/>(RPC层)]
    Generate --> RedisWrite[写入Redis<br/>(uploadID、分块信息)]
    RedisWrite --> RespInit[返回初始化信息<br/>(前端)]
    RespInit --> Upload[前端分片上传<br/>(分块文件+index)]
    Upload --> CheckRedis{检查分块状态<br/>(Redis: 是否已上传?)}
    CheckRedis -->|是| Skip[跳过<br/>(返回已上传)]
    CheckRedis -->|否| SaveChunk[保存分块到本地<br/>(uploadID/index)]
    SaveChunk --> UpdateRedis[更新Redis状态<br/>(标记已上传)]
    UpdateRedis --> AllUploaded{所有分块上传完成?}
    AllUploaded -->|否| Upload
    AllUploaded -->|是| Merge[合并文件<br/>(检查完整性)]
    Merge --> Storage{选择存储方式<br/>(本地/Cos/Minio)}
    Storage --> Local[本地存储<br/>(直接保存)]
    Storage --> Cos[调用Cos RPC<br/>(云存储)]
    Storage --> Minio[调用Minio RPC<br/>(对象存储)]
    Local --> Clean[清理临时文件<br/>(删除分块+Redis数据)]
    Cos --> Clean
    Minio --> Clean
    Clean --> End[结束]
```


## 文件妙传

即查询数据库中文件sha1值是否存在，存在则直接建立文件元数据关系

