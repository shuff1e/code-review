/*

626. 换座位
SQL架构
小美是一所中学的信息科技老师，她有一张 seat 座位表，平时用来储存学生名字和与他们相对应的座位 id。

其中纵列的 id 是连续递增的

小美想改变相邻俩学生的座位。

你能不能帮她写一个 SQL query 来输出小美想要的结果呢？



示例：

+---------+---------+
|    id   | student |
+---------+---------+
|    1    | Abbot   |
|    2    | Doris   |
|    3    | Emerson |
|    4    | Green   |
|    5    | Jeames  |
+---------+---------+
假如数据输入的是上表，则输出结果如下：

+---------+---------+
|    id   | student |
+---------+---------+
|    1    | Doris   |
|    2    | Abbot   |
|    3    | Green   |
|    4    | Emerson |
|    5    | Jeames  |
+---------+---------+
注意：

如果学生人数是奇数，则不需要改变最后一个同学的座位。

 */

 /*

 方法一：使用 CASE【通过】
算法

对于所有座位 id 是奇数的学生，修改其 id 为 id+1，如果最后一个座位 id 也是奇数，则最后一个座位 id 不修改。对于所有座位 id 是偶数的学生，修改其 id 为 id-1。

首先查询座位的数量。

MySQL

SELECT
    COUNT(*) AS counts
FROM
    seat
然后使用 CASE 条件和 MOD 函数修改每个学生的座位 id。

MySQL

MySQL

SELECT
    (CASE
        WHEN MOD(id, 2) != 0 AND counts != id THEN id + 1
        WHEN MOD(id, 2) != 0 AND counts = id THEN id
        ELSE id - 1
    END) AS id,
    student
FROM
    seat,
    (SELECT
        COUNT(*) AS counts
    FROM
        seat) AS seat_counts
ORDER BY id ASC;



方法二：使用位操作和 COALESCE()【通过】
算法

使用 (id+1)^1-1 计算交换后每个学生的座位 id。

MySQL

SELECT id, (id+1)^1-1, student FROM seat;

| id | (id+1)^1-1 | student |
|----|------------|---------|
| 1  | 2          | Abbot   |
| 2  | 1          | Doris   |
| 3  | 4          | Emerson |
| 4  | 3          | Green   |
| 5  | 6          | Jeames  |
然后连接原来的座位表和更新 id 后的座位表。

MySQL

SELECT
    *
FROM
    seat s1
        LEFT JOIN
    seat s2 ON (s1.id+1)^1-1 = s2.id
ORDER BY s1.id;

| id | student | id | student |
|----|---------|----|---------|
| 1  | Abbot   | 2  | Doris   |
| 2  | Doris   | 1  | Abbot   |
| 3  | Emerson | 4  | Green   |
| 4  | Green   | 3  | Emerson |
| 5  | Jeames  |    |         |
注：前两列来自表 s1，后两列来自表 s2。

最后输出 s1.id 和 s2.student。
但是 id=5 的学生，s1.student 正确，s2.student 为 NULL。因此使用 COALESCE() 函数为最后一行记录生成正确的输出。

MySQL

MySQL

SELECT
    s1.id, COALESCE(s2.student, s1.student) AS student
FROM
    seat s1
        LEFT JOIN
    seat s2 ON ((s1.id + 1) ^ 1) - 1 = s2.id
ORDER BY s1.id;

  */

 select
 ( case
 when mod(id,2) != 0 and id != seat_counts.counts then id + 1
 when mod(id,2) != 0 and id = seat_counts.counts then id
 else id - 1
 end)
  as id,student

 from seat,(select count(*) as counts from seat) as seat_counts


 order by id asc

 /*
 COALESCE(value1,value2,...);

COALESCE函数需要许多参数，并返回第一个非NULL参数。如果所有参数都为NULL，则COALESCE函数返回NULL。
  */



  select s1.id,coalesce(s2.student,s1.student) as student

from seat s1 left join seat s2 on ((s1.id+1)^1) - 1 = s2.id

order by s1.id