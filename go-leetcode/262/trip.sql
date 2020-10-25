/*
262. 行程和用户
SQL架构
Trips 表中存所有出租车的行程信息。
每段行程有唯一键 Id，Client_Id 和 Driver_Id 是 Users 表中 Users_Id 的外键。
Status 是枚举类型，枚举成员为 (‘completed’, ‘cancelled_by_driver’, ‘cancelled_by_client’)。

+----+-----------+-----------+---------+--------------------+----------+
| Id | Client_Id | Driver_Id | City_Id |        Status      |Request_at|
+----+-----------+-----------+---------+--------------------+----------+
| 1  |     1     |    10     |    1    |     completed      |2013-10-01|
| 2  |     2     |    11     |    1    | cancelled_by_driver|2013-10-01|
| 3  |     3     |    12     |    6    |     completed      |2013-10-01|
| 4  |     4     |    13     |    6    | cancelled_by_client|2013-10-01|
| 5  |     1     |    10     |    1    |     completed      |2013-10-02|
| 6  |     2     |    11     |    6    |     completed      |2013-10-02|
| 7  |     3     |    12     |    6    |     completed      |2013-10-02|
| 8  |     2     |    12     |    12   |     completed      |2013-10-03|
| 9  |     3     |    10     |    12   |     completed      |2013-10-03|
| 10 |     4     |    13     |    12   | cancelled_by_driver|2013-10-03|
+----+-----------+-----------+---------+--------------------+----------+
Users 表存所有用户。每个用户有唯一键 Users_Id。Banned 表示这个用户是否被禁止，
Role 则是一个表示（‘client’, ‘driver’, ‘partner’）的枚举类型。

+----------+--------+--------+
| Users_Id | Banned |  Role  |
+----------+--------+--------+
|    1     |   No   | client |
|    2     |   Yes  | client |
|    3     |   No   | client |
|    4     |   No   | client |
|    10    |   No   | driver |
|    11    |   No   | driver |
|    12    |   No   | driver |
|    13    |   No   | driver |
+----------+--------+--------+
写一段 SQL 语句查出 2013年10月1日 至 2013年10月3日 期间非禁止用户的取消率。
基于上表，你的 SQL 语句应返回如下结果，取消率（Cancellation Rate）保留两位小数。

取消率的计算方式如下：(被司机或乘客取消的非禁止用户生成的订单数量) / (非禁止用户生成的订单总数)

+------------+-------------------+
|     Day    | Cancellation Rate |
+------------+-------------------+
| 2013-10-01 |       0.33        |
| 2013-10-02 |       0.00        |
| 2013-10-03 |       0.50        |
+------------+-------------------+
致谢:
非常感谢 @cak1erlizhou 详细的提供了这道题和相应的测试用例。
 */

 /*

 统计每天非禁止用户的取消率，需要知道非禁止用户有哪些，总行程数，取消的行程数。

解法一
首先确定被禁止用户的行程记录，再剔除这些行程记录。

行程表中，字段 client_id 和 driver_id，都与用户表中的 users_id 关联。
因此只要 client_id 和 driver_id 中有一个被禁止了，此条行程记录要被剔除。

先说一种错误的找出没被禁止用户行程记录的方法。此方法很有迷惑性。

思路：


if (client_id = users_id 或 driver_id = users_id) 且 users_id没有被禁止
{
    此条记录没被禁止。
}
SQL 代码


SELECT *
FROM Trips AS T JOIN Users AS U
ON (T.client_id = U.users_id  OR T.driver_id = U.users_id )  AND U.banned ='No'
乍一看，思路是对。其实是错误的。因为，我们不知觉得肯定了一个假设—— client_id 与 driver_id 是相同的。
只有当两者相同时，才能用此条件排除被禁止用户的行程记录。

错误的结果：


+------+-----------+-----------+---------+---------------------+------------+----------+--------+--------+
| Id   | Client_Id | Driver_Id | City_Id | STATUS              | Request_at | Users_Id | Banned | Role   |
+------+-----------+-----------+---------+---------------------+------------+----------+--------+--------+
|    1 |         1 |        10 |       1 | completed           | 2013-10-01 |        1 | No     | client |
|    1 |         1 |        10 |       1 | completed           | 2013-10-01 |       10 | No     | driver |
|    2 |         2 |        11 |       1 | cancelled_by_driver | 2013-10-01 |       11 | No     | driver |
|    3 |         3 |        12 |       6 | completed           | 2013-10-01 |        3 | No     | client |
|    3 |         3 |        12 |       6 | completed           | 2013-10-01 |       12 | No     | driver |
|    4 |         4 |        13 |       6 | cancelled_by_client | 2013-10-01 |        4 | No     | client |
|    4 |         4 |        13 |       6 | cancelled_by_client | 2013-10-01 |       13 | No     | driver |
|    5 |         1 |        10 |       1 | completed           | 2013-10-02 |        1 | No     | client |
|    5 |         1 |        10 |       1 | completed           | 2013-10-02 |       10 | No     | driver |
|    6 |         2 |        11 |       6 | completed           | 2013-10-02 |       11 | No     | driver |
|    7 |         3 |        12 |       6 | completed           | 2013-10-02 |        3 | No     | client |
|    7 |         3 |        12 |       6 | completed           | 2013-10-02 |       12 | No     | driver |
|    8 |         2 |        12 |      12 | completed           | 2013-10-03 |       12 | No     | driver |
|    9 |         3 |        10 |      12 | completed           | 2013-10-03 |        3 | No     | client |
|    9 |         3 |        10 |      12 | completed           | 2013-10-03 |       10 | No     | driver |
|   10 |         4 |        13 |      12 | cancelled_by_driver | 2013-10-03 |        4 | No     | client |
|   10 |         4 |        13 |      12 | cancelled_by_driver | 2013-10-03 |       13 | No     | driver |
+------+-----------+-----------+---------+---------------------+------------+----------+--------+--------+
结果中，被禁止的 users_id = 2，其行程记录没被剔除掉。

明显， client_id 与 driver_id 不一定相同 。

正确的做法是对 client_id 和 driver_id 各自关联的 users_id，同时检测是否被禁止。


if (client_id = users_id_1 且 users_id_1没被禁止 并且 client_id = users_id_2 且 users_id_2没被禁止){
    此条记录没被禁止。
}
SQL 代码：


SELECT *
FROM Trips AS T
JOIN Users AS U1 ON (T.client_id = U1.users_id AND U1.banned ='No')
JOIN Users AS U2 ON (T.driver_id = U2.users_id AND U2.banned ='No')
在此基础上，按日期分组，统计每组的 总行程数，取消的行程数 。

每组的总行程数：COUNT(T.STATUS)。

每组的取消的行程数：


SUM(
	IF(T.STATUS = 'completed',0,1)
)
取消率 = 每组的取消的行程数 / 每组的总行程数

完整逻辑为:


SELECT T.request_at AS `Day`,
	ROUND(
			SUM(
				IF(T.STATUS = 'completed',0,1)
			)
			/
			COUNT(T.STATUS),
			2
	) AS `Cancellation Rate`
FROM Trips AS T
JOIN Users AS U1 ON (T.client_id = U1.users_id AND U1.banned ='No')
JOIN Users AS U2 ON (T.driver_id = U2.users_id AND U2.banned ='No')
WHERE T.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY T.request_at
其中 SUM 求和函数，COUNT 计数函数，ROUND 四舍五入函数。

解法二
思路与解法一相同。而采用不同的方法排除掉被禁止用户的行程记录。想到排除，就联想到集合差。

client_id 和 driver_id 的全部为集合 U。被禁止的 users_id 集合为 A。

U 减去 A 的结果为没被禁止的用户。


(
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
) AS A
好了，先演示一个错误的解法：

行程表连接表 A，排除掉被被禁止的行程。


SELECT *
FROM trips AS T,
(
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
) AS A
WHERE (T.Client_Id != A.users_id AND T.Driver_Id != A.users_id)
剩下的逻辑与解法一后部分相同，完善后的逻辑为：


SELECT T.request_at AS `Day`,
	ROUND(
			SUM(
				IF(T.STATUS = 'completed',0,1)
			)
			/
			COUNT(T.STATUS),
			2
	) AS `Cancellation Rate`
FROM trips AS T,
(
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
) AS A
WHERE (T.Client_Id != A.users_id AND T.Driver_Id != A.users_id) AND T.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY T.request_at
很可惜，当表 A 为空时，此方法的结果是空表。但是表 A 为空，可能是有用户但是没有被禁止的用户。因此方法是错误的。

正确的解法是：行程表 left join 表 A 两次，A.users_id 都为 NULL 的行都是没被排除的行。


SELECT *
FROM trips AS T LEFT JOIN
(
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
) AS A ON (T.Client_Id = A.users_id)
LEFT JOIN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
) AS A1
ON (T.Driver_Id = A1.users_id)
WHERE A.users_id IS NULL AND A1.users_id IS NULL
补上其它部分的逻辑为：


SELECT T.request_at AS `Day`,
	ROUND(
			SUM(
				IF(T.STATUS = 'completed',0,1)
			)
			/
			COUNT(T.STATUS),
			2
	) AS `Cancellation Rate`
FROM trips AS T LEFT JOIN
(
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
) AS A ON (T.Client_Id = A.users_id)
LEFT JOIN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
) AS A1
ON (T.Driver_Id = A1.users_id)
WHERE A.users_id IS NULL AND A1.users_id IS NULL AND T.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY T.request_at
解法三
与解法二思路相同。找出被禁止的用户后，不再连接行程表和用户表，直接从行程表中排除掉被被禁止用户的行程记录。

被禁止的用户用子查询：


(
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
)
行程表中 client_id 和 driver_id 都在此子查询结果中的行要剔除掉。


SELECT *
FROM trips AS T
WHERE
T.Client_Id NOT IN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
)
AND
T.Driver_Id NOT IN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
)
补上其它部分：


SELECT T.request_at AS `Day`,
	ROUND(
			SUM(
				IF(T.STATUS = 'completed',0,1)
			)
			/
			COUNT(T.STATUS),
			2
	) AS `Cancellation Rate`
FROM trips AS T
WHERE
T.Client_Id NOT IN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
)
AND
T.Driver_Id NOT IN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
)
AND T.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY T.request_at

  */

select T.request_at as `Day`,
round(sum(if (T.status = 'completed',0,1)) / count(T.status),2) as `Cancellation Rate`
from Trips as T
join Users as U1 on (T.client_id = U1.users_id and U1.banned = 'NO')
join Users as U2 on (T.driver_id = U2.users_id and U2.banned = 'NO')
where T.request_at between '2013-10-01' AND '2013-10-03'
group by T.request_at



select T.request_at as `Day`,
round(sum(if (T.status = 'completed',0,1)) / count(T.status),2) as `Cancellation Rate`
from Trips as T
left join
( select users_id from users where banned = 'Yes') as A on (T.client_id = A.users_id)
left join
(select users_id from users where banned = 'Yes') as A1 on (T.Driver_id = A1.users_id)
where A.users_id is null and A1.users_id is null and T.request_at between '2013-10-01' and '2013-10-03' group by T.request_at



select T.request_at as `Day`,
round(sum(if (T.status = 'completed',0,1)) / count(T.status),2) as `Cancellation Rate`
from Trips as T
where
T.client_id not in (select users_id from users where banned = 'Yes')
and
T.driver_id not in (select users_id from users where banned = 'Yes')
and
T.request_at between '2013-10-01' AND '2013-10-03'
group by T.request_at