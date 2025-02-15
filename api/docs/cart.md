## 购物车相关接口

### 1. 添加商品到购物车
#### 请求方法
POST

#### 请求路径
/cart/add

#### 请求参数
| 参数名     | 类型 | 必填 | 说明     | 示例 |
|------------|------|------|----------|------|
| product_id | uint | 是   | 商品ID   | 1    |
| quantity   | int  | 是   | 购买数量 | 1    |

#### 请求示例
{
    "product_id": 1,
    "quantity": 1
}

#### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": "添加购物车成功"
}

### 2. 获取购物车列表
#### 请求方法
GET

#### 请求路径
/cart/list

#### 响应参数
| 参数名      | 类型    | 说明         |
|-------------|---------|--------------|
| account     | float64 | 总价         |
| cart        | array   | 购物车商品列表|

#### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": {
        "account": 199.98,
        "cart": [
            {
                "product_id": 1,
                "name": "商品名称",
                "type": "商品类型",
                "price": 99.99,
                "cover": "商品封面URL",
                "link": "商品链接",
                "num": 2
            }
        ]
    }
}

### 3. 清空购物车
#### 请求方法
DELETE

#### 请求路径
/cart/clear

#### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": "清空购物车成功"
}
