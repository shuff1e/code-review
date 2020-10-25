/*
595. 大的国家
SQL架构
这里有张 World 表

+-----------------+------------+------------+--------------+---------------+
| name            | continent  | area       | population   | gdp           |
+-----------------+------------+------------+--------------+---------------+
| Afghanistan     | Asia       | 652230     | 25500100     | 20343000      |
| Albania         | Europe     | 28748      | 2831741      | 12960000      |
| Algeria         | Africa     | 2381741    | 37100000     | 188681000     |
| Andorra         | Europe     | 468        | 78115        | 3712000       |
| Angola          | Africa     | 1246700    | 20609294     | 100990000     |
+-----------------+------------+------------+--------------+---------------+
如果一个国家的面积超过 300 万平方公里，或者人口超过 2500 万，那么这个国家就是大国家。

编写一个 SQL 查询，输出表中所有大国家的名称、人口和面积。

例如，根据上表，我们应该输出:

+--------------+-------------+--------------+
| name         | population  | area         |
+--------------+-------------+--------------+
| Afghanistan  | 25500100    | 652230       |
| Algeria      | 37100000    | 2381741      |
+--------------+-------------+--------------+
 */

 /*

 方法一：使用 WHERE 子句和 OR【通过】
思路

使用 WHERE 子句过滤所有记录，获得满足条件的国家。

算法

根据定义，大国家至少满足以下两个条件中的一个：

面积超过 300 万平方公里。
人口超过 2500 万。
使用下面语句获得满足条件 1 的大国家。

MySQL

SELECT name, population, area FROM world WHERE area > 3000000
使用下面语句获得满足条件 2 的大国家。

MySQL

SELECT name, population, area FROM world WHERE population > 25000000
使用 OR 将两个子查询合并在一起。

MySQL

MySQL

SELECT
    name, population, area
FROM
    world
WHERE
    area > 3000000 OR population > 25000000
;
方法二：使用 WHERE 子句和 UNION【通过】
算法

该方法思路与 方法一 一样，但是使用 UNION 连接子查询。

MySQL

MySQL

SELECT
    name, population, area
FROM
    world
WHERE
    area > 3000000

UNION

SELECT
    name, population, area
FROM
    world
WHERE
    population > 25000000
;
注：方法二 比 方法一 运行速度更快，但是它们没有太大差别。

  */

select name,population,area
from world
where
area > 3000000

union

select name,population,area
from world
where
population > 25000000






select
name ,population,area
from world
where
area > 3000000 or population > 25000000