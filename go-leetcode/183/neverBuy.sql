/*
183. 从不订购的客户
SQL架构
某网站包含两个表，Customers 表和 Orders 表。编写一个 SQL 查询，找出所有从不订购任何东西的客户。

Customers 表：

+----+-------+
| Id | Name  |
+----+-------+
| 1  | Joe   |
| 2  | Henry |
| 3  | Sam   |
| 4  | Max   |
+----+-------+
Orders 表：

+----+------------+
| Id | CustomerId |
+----+------------+
| 1  | 3          |
| 2  | 1          |
+----+------------+
例如给定上述表格，你的查询应返回：

+-----------+
| Customers |
+-----------+
| Henry     |
| Max       |
+-----------+
 */

 select customers.name as 'Customers' from customers
 where customers.id not in (
 select customerid from orders
 )

 /*

 方法：使用子查询和 NOT IN 子句
算法

如果我们有一份曾经订购过的客户名单，就很容易知道谁从未订购过。

我们可以使用下面的代码来获得这样的列表。


select customerid from orders;
然后，我们可以使用 NOT IN 查询不在此列表中的客户。

MySQL

select customers.name as 'Customers'
from customers
where customers.id not in
(
    select customerid from orders
);

  */