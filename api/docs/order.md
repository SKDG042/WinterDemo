## 订单相关接口

### 1. 创建订单
#### 请求方法
POST

#### 请求路径
/order/create

#### 请求头
| 参数名        | 必填 | 说明                                  | 示例                                      |
|--------------|------|---------------------------------------|-------------------------------------------|
| Authorization| 是   | Bearer Token                          | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6...   |

#### 响应参数
| 参数名      | 类型    | 说明                                  |
|------------|---------|---------------------------------------|
| order_id   | uint    | 订单ID                                |
| status     | string  | 订单状态（待支付/已支付/已取消等）      |
| total_price| float64 | 订单总价                              |
| created_at | string  | 创建时间                              |
| products   | array   | 订单商品列表                          |

#### 响应示例
成功：
{
    "status": 10000,
    "info": "success",
    "data": {
        "order_id": 1,
        "status": "待支付",
        "total_price": 199.98,
        "created_at": "2024-01-01 12:00:00",
        "products": [
            {
                "product_id": 1,
                "name": "商品名称",
                "price": 99.99,
                "quantity": 2
            }
        ]
    }
}

失败：
{
    "status": 10101,
    "info": "请先添加商品到购物车",
    "data": null
}