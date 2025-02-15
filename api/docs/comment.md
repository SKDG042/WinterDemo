## 评论相关接口

### 1. 添加评论
#### 请求方法
POST

#### 请求路径
/comment/add/{product_id}

#### 请求参数
| 参数名      | 类型    | 必填 | 说明       | 示例     |
|-------------|---------|------|------------|----------|
| content     | string  | 是   | 评论内容   | "很好用" |
| parent_id   | uint    | 否   | 父评论ID   | 1        |
| is_anonymous| boolean | 否   | 是否匿名   | false    |

#### 请求示例
{
    "content": "很好用",
    "parent_id": 1,
    "is_anonymous": false
}

#### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": {
        "comment_id": 1,
        "message": "评论成功"
    }
}

### 2. 删除评论
#### 请求方法
DELETE

#### 请求路径
/comment/delete/{comment_id}

#### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": "删除评论成功"
}

### 3. 获取商品评论
#### 请求方法
GET

#### 请求路径
/comment/list/{product_id}

#### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": [
        {
            "comment_id": 1,
            "content": "评论内容",
            "user_id": 1,
            "nickname": "用户昵称",
            "avatar": "用户头像URL",
            "product_id": 1,
            "parent_id": 0,
            "created_at": "2024-01-01 12:00:00",
            "children": [
                {
                    "comment_id": 2,
                    "content": "回复内容",
                    "user_id": 2,
                    "nickname": "回复用户昵称",
                    "avatar": "用户头像URL",
                    "product_id": 1,
                    "parent_id": 1,
                    "created_at": "2024-01-01 12:01:00",
                    "children": []
                }
            ]
        }
    ]
}