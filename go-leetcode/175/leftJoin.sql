
/*

175. 组合两个表
SQL架构
表1: Person

+-------------+---------+
| 列名         | 类型     |
+-------------+---------+
| PersonId    | int     |
| FirstName   | varchar |
| LastName    | varchar |
+-------------+---------+
PersonId 是上表主键
表2: Address

+-------------+---------+
| 列名         | 类型    |
+-------------+---------+
| AddressId   | int     |
| PersonId    | int     |
| City        | varchar |
| State       | varchar |
+-------------+---------+
AddressId 是上表主键

编写一个 SQL 查询，满足条件：无论 person 是否有地址信息，都需要基于上述两表提供 person 的以下信息：

FirstName, LastName, City, State
*/

/*
方法：使用 outer join
算法

因为表 Address 中的 personId 是表 Person 的外关键字，所以我们可以连接这两个表来获取一个人的地址信息。

考虑到可能不是每个人都有地址信息，我们应该使用 outer join 而不是默认的 inner join。
*/

select FirstName, LastName, City, State
from Person left join Address
on Person.PersonId = Address.PersonId
;

/*
假设你要join两个没有重复列的表，这是最常见的情况：

inner join  A 和 B 获得的是A和B的交集(intersect),即韦恩图(venn diagram) 相交的部分.

outer join A和B获得的是A和B的并集(union), 即韦恩图(venn diagram)的所有部分.

示例

假定有两张表，每张表只有一列，列数据如下：

A    B
-    -
1    3
2    4
3    5
4    6
注意(1,2)是A表独有的，(3,4) 两张共有， (5,6)是B独有的。

Inner join

使用等号进行inner join以获得两表的交集，即共有的行。

select * from a INNER JOIN b on a.a = b.b;
select a.*,b.*  from a,b where a.a = b.b;

a | b
--+--
3 | 3
4 | 4
Left outer join

 left outer join 除了获得B表中符合条件的列外，还将获得A表所有的列。

select * from a LEFT OUTER JOIN b on a.a = b.b;
select a.*,b.*  from a,b where a.a = b.b(+);

a |  b
--+-----
1 | null
2 | null
3 |    3
4 |    4
Full outer join

full outer join 得到A和B的交集,即A和B中所有的行.。如果A中的行在B中没有对应的部分,B的部分将是 null, 反之亦然。

select * from a FULL OUTER JOIN b on a.a = b.b;

 a   |  b
-----+-----
   1 | null
   2 | null
   3 |    3
   4 |    4
null |    6
null |    5
*/