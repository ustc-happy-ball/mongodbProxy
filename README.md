该工程是一个database proxy，其他服务通过gRPC调用的方式来对数据库进行操作。

#### 数据库collection结构

##### Account collection

| 字段名         | 类型     | 注解                      |
| -------------- | -------- | ------------------------- |
| _id            | objectId |                           |
| player_id      | int32    | 该account对应的player的id |
| login_password | string   | 账户密码                  |
| delete         | bool     | 当前账号是否注销          |
| phone          | string   | 电话                      |
| recent_login   | int64    | 最近一次登录的时间        |
| create_at      | int64    | 记录创建时间              |
| update_at      | int64    | 记录上一次更新时间        |

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gpskfyx13ej30qo0ccq5j.jpg" alt="image-20210422154405709" style="zoom:50%;" />

##### Player collection

| 字段名        | 类型   | 注解                    |
| ------------- | ------ | ----------------------- |
| player_id     | int32  | player的id（全局唯一）  |
| account_id    | string | 与该player对应的account |
| highest_score | int32  | 该玩家的历史最高分      |
| highest_rank  | int32  | 该玩家的历史最高排名    |
| create_at     | int64  | 该记录的创建时间        |
| update_at     | int64  | 该记录的更新时间        |

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gpskfohftyj30k409ugn7.jpg" alt="image-20210422154345042" style="zoom:50%;" />

#### gRPC接口

##### Account collection 操作接口

1. 通过手机号来查找account信息

   ```protobuf
   // 请求格式
   message AccountFindByPhoneRequest { 
     string phone = 1;
   }
   // 返回格式
   message AccountFindByPhoneResponse {
     Account account = 1;
   }
   ```

2. 添加新的account记录

   ```protobuf
   message AccountAddRequest {
     Account account = 1;
   }
   
   message AccountAddResponse {
   }
   ```

Account的protobuf定义：

```protobuf
message Account {
  string objectId = 1;
  int32 playerId = 2;
  string loginPassword = 3;
  bool delete = 4;
  string phone = 5;
  int64 recentLogin = 6;
  int64 createAt = 7;
  int64 updateAt = 8;
}
```

##### Player collection 操作接口

1. 通过playerId查找player信息

   ```protobuf
   message PlayerFindByPlayerIdRequest {
     int32 playerId = 1;
   }
   
   message PlayerFindByPlayerIdResponse {
     Player player = 1;
   }
   ```

2. 添加新的player记录

   ```protobuf
   message PlayerAddRequest {
     Player player = 1;
   }
   
   message PlayerAddResponse {
   }
   ```

Player的protobuf定义：

```protobuf
message Player {
  int32 playerId = 2;
  string accountId = 3;
  int32 highestScore = 4;
  int32 highestRank = 5;
  int64 createAt = 6;
  int64 updateAt = 7;
}
```

