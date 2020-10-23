/*
196. 删除重复的电子邮箱
编写一个 SQL 查询，来删除 Person 表中所有重复的电子邮箱，重复的邮箱里只保留 Id 最小 的那个。

+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
Id 是这个表的主键。
例如，在运行你的查询语句之后，上面的 Person 表应返回以下几行:

+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+


提示：

执行 SQL 之后，输出是整个 Person 表。
使用 delete 语句。
 */

/*

方法：使用 DELETE 和 WHERE 子句
算法

我们可以使用以下代码，将此表与它自身在电子邮箱列中连接起来。

MySQL

SELECT p1.*
FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email
;
然后我们需要找到其他记录中具有相同电子邮件地址的更大 ID。所以我们可以像这样给 WHERE 子句添加一个新的条件。

MySQL

SELECT p1.*
FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email AND p1.Id > p2.Id
;
因为我们已经得到了要删除的记录，所以我们最终可以将该语句更改为 DELETE。

MySQL

DELETE p1 FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email AND p1.Id > p2.Id

 */

/*

 方法：使用 DELETE 和 WHERE 子句
算法

我们可以使用以下代码，将此表与它自身在电子邮箱列中连接起来。

MySQL

SELECT p1.*
FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email
;
然后我们需要找到其他记录中具有相同电子邮件地址的更大 ID。所以我们可以像这样给 WHERE 子句添加一个新的条件。

MySQL

SELECT p1.*
FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email AND p1.Id > p2.Id
;
因为我们已经得到了要删除的记录，所以我们最终可以将该语句更改为 DELETE。

MySQL

DELETE p1 FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email AND p1.Id > p2.Id

官方sql是下面这样的👇

MySQL

DELETE p1 FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email AND p1.Id > p2.Id
当然这个sql是很ok的，简洁清晰，且用了自连接的方式。
有慢查询优化经验的同学会清楚，在实际生产中，面对千万上亿级别的数据，连接的效率往往最高，因为用到索引的概率较高。

因此，建议学习使用官方的题解，但是有两点，可能需要再解释下：

1、DELETE p1

在DELETE官方文档中，给出了这一用法，比如下面这个DELETE语句👇

DELETE t1 FROM t1 LEFT JOIN t2 ON t1.id=t2.id WHERE t2.id IS NULL;

这种DELETE方式很陌生，竟然和SELETE的写法类似。它涉及到t1和t2两张表，DELETE t1表示要删除t1的一些记录，具体删哪些，就看WHERE条件，满足就删；

这里删的是t1表中，跟t2匹配不上的那些记录。

所以，官方sql中，DELETE p1就表示从p1表中删除满足WHERE条件的记录。

2、p1.Id > p2.Id

继续之前，先简单看一下表的连接过程，这个搞懂了，理解WHERE条件就简单了👇

a. 从驱动表（左表）取出N条记录；
b. 拿着这N条记录，依次到被驱动表（右表）查找满足WHERE条件的记录；

所以，官方sql的过程就是👇

先把Person表搬过来( •̀∀•́ )


a. 从表p1取出3条记录；
b. 拿着第1条记录去表p2查找满足WHERE的记录，代入该条件p1.Email = p2.Email AND p1.Id > p2.Id后，发现没有满足的，所以不用删掉记录1；
c. 记录2同理；
d. 拿着第3条记录去表p2查找满足WHERE的记录，发现有一条记录满足，所以要从p1删掉记录3；
e. 3条记录遍历完，删掉了1条记录，这个DELETE也就结束了。

 */

delete p1 from Person as p1,Person p2 where p1.Email = p2.Email and p1.Id > p2.Id