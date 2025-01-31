# 用户相关 API

## 1. 用户注册
### 接口描述
注册新用户账号

### 请求方法
POST

### 请求路径
/user/register

### 请求参数
| 参数名   | 类型   | 必填 | 说明         | 示例     |
|----------|--------|------|--------------|----------|
| username | string | 是   | 用户名(3-20位) | "zhangsan"|
| password | string | 是   | 密码(6-20位)  | "123456" |

### 请求示例
{
    "username": "zhangsan",
    "password": "123456"
}

### 响应参数
| 参数名 | 类型   | 说明     | 示例     |
|--------|--------|----------|----------|
| status | int    | 状态码   | 10000    |
| info   | string | 状态信息 | "success"|
| data   | object | 返回数据 | {"message": "注册成功"} |

### 响应示例
成功：
{
    "status": 10000,
    "info": "success",
    "data": {
        "message": "注册成功"
    }
}

失败：
{
    "status": 10003,
    "info": "用户名已存在",
    "data": null
}

## 2. 用户登录
### 接口描述
用户登录获取认证token

### 请求方法
POST

### 请求路径
/user/login

### 请求参数
| 参数名   | 类型   | 必填 | 说明     | 示例     |
|----------|--------|------|----------|----------|
| username | string | 是   | 用户名    | "zhangsan"|
| password | string | 是   | 密码      | "123456" |

### 请求示例
{
    "username": "zhangsan",
    "password": "123456"
}

### 响应参数
| 参数名 | 类型   | 说明     | 示例     |
|--------|--------|----------|----------|
| status | int    | 状态码   | 10000    |
| info   | string | 状态信息 | "success"|
| data   | object | 返回数据 | {"token": "xxx", "refresh_token": "xxx"} |

### 响应示例
成功：
{
    "status": 10000,
    "info": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
}

## 3. 刷新Token
### 接口描述
使用refresh_token刷新access_token

### 请求方法
POST

### 请求路径
/user/token/refresh

### 请求参数
| 参数名        | 类型   | 必填 | 说明          | 示例     |
|--------------|--------|------|---------------|----------|
| refresh_token| string | 是   | 刷新令牌      | "Bearer eyJhbG..." |

### 请求示例
{
    "refresh_token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

### 响应示例
成功：
{
    "status": 10000,
    "info": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
}

## 4. 修改密码
### 接口描述
修改用户密码

### 请求方法
POST

### 请求路径
/user/password/update

### 请求头
| 参数名        | 说明         | 示例                                    |
|--------------|--------------|----------------------------------------|
| Authorization| Bearer Token | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI... |

### 请求参数
| 参数名      | 类型   | 必填 | 说明          | 示例      |
|------------|--------|------|---------------|-----------|
| newPassword| string | 是   | 新密码(6-20位) | "654321"  |

### 请求示例
{
    "newPassword": "654321"
}

### 响应示例
成功：
{
    "status": 10000,
    "info": "success",
    "data": {
        "message": "更新成功"
    }
}