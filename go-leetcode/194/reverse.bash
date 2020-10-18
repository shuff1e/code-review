#给定一个文件 file.txt，转置它的内容。
#
#你可以假设每行列数相同，并且每个字段由 ' ' 分隔.
#
#示例:
#
#假设 file.txt 文件内容如下：
#
#name age
#alice 21
#ryan 30
#应当输出：
#
#name alice ryan
#age 21 30

#!/bin/bash

column=$(awk '{print NF}' file.txt |uniq)

for ((i=1;i<=column;i++))
do
    cut -d ' ' -f $i file.txt | xargs
done

awk '{for (i=1;i<=NF;i++) {if (NR==1) {row[i]=$i} else {row[i]=row[i]" "$i}}} END {for (a in row){print row[a]}}' file.txt