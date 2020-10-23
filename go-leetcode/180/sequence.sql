/*
180. 连续出现的数字
SQL架构
编写一个 SQL 查询，查找所有至少连续出现三次的数字。

+----+-----+
| Id | Num |
+----+-----+
| 1  |  1  |
| 2  |  1  |
| 3  |  1  |
| 4  |  2  |
| 5  |  1  |
| 6  |  2  |
| 7  |  2  |
+----+-----+
例如，给定上面的 Logs 表， 1 是唯一连续出现至少三次的数字。

+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+
 */

/*
方法：用 DISTINCT 和 WHERE 语句
算法

连续出现的意味着相同数字的 Id 是连着的，由于这题问的是至少连续出现 3 次，我们使用 Logs 并检查是否有 3 个连续的相同数字。


SELECT *
FROM
    Logs l1,
    Logs l2,
    Logs l3
WHERE
    l1.Id = l2.Id - 1
    AND l2.Id = l3.Id - 1
    AND l1.Num = l2.Num
    AND l2.Num = l3.Num
;



Id	Num	Id	Num	Id	Num
1	1	2	1	3	1
注意：前两列来自 l1 ，接下来两列来自 l2 ，最后两列来自 l3 。

然后我们从上表中选择任意的 Num 获得想要的答案。同时我们需要添加关键字 DISTINCT ，因为如果一个数字连续出现超过 3 次，会返回重复元素。

MySQL


SELECT DISTINCT
    l1.Num AS ConsecutiveNums
FROM
    Logs l1,
    Logs l2,
    Logs l3
WHERE
    l1.Id = l2.Id - 1
    AND l2.Id = l3.Id - 1
    AND l1.Num = l2.Num
    AND l2.Num = l3.Num
;

 */

select distinct l1.Num as ConsecutiveNums
from
Logs l1,
Logs l2,
Logs l3

where
l1.Id = l2.Id - 1
and l2.Id = l3.Id - 1
and l1.Num = l2.Num
and l2.Num = l3.Num;