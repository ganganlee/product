## 商品管理API
> 1、添加商品
- /product
- post
- json

| 名称 | 必填 | 类型 | 说明 |
| ------ | ------ | ------ |------ |
| name | true | string | 商品名称 |
| num | true | uint | 商品数量 |
|image|true|string|商品图片|
|url|true|string|商品链接|
```cgo
//success
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "ok"
}
```

> 2、修改商品
- /product
- put
- json

| 名称 | 必填 | 类型 | 说明 |
| ------ | ------ | ------ |------ |
| id | true | int | 商品id |
| name | true | string | 商品名称 |
| num | true | uint | 商品数量 |
|image|true|string|商品图片|
|url|true|string|商品链接|
```cgo
//success
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "ok"
}
```

> 3、删除商品
- /product
- delete
- json

| 名称 | 必填 | 类型 | 说明 |
| ------ | ------ | ------ |------ |
| id | true | int | 商品id |
```cgo
//success
{
  "code": 200,
  "data": {
    "id": 1
  },
  "msg": "ok"
}
```

> 4、获取商品列表
- /product
- get
- json
```cgo
//success
{
  "code": 200,
  "data": [
    {
      "ID": 2,
      "CreatedAt": "2020-04-27T08:23:45+08:00",
      "UpdatedAt": "2020-04-27T08:23:45+08:00",
      "DeletedAt": null,
      "name": "添加项目添加项目001",
      "num": 100,
      "image": "http://www.baidu.com",
      "url": "http://www.baidu.com"
    },
    ...
  ],
  "msg": "ok"
}
```