#!/bin/bash

column=$(awk '{print NF}' file.txt |uniq)

for ((i=1;i<=column;i++))
do
    cut -d ' ' -f $i file.txt | xargs
done

awk '{for (i=1;i<=NF;i++) {if (NR==1) {row[i]=$i} else {row[i]=row[i]" "$i}}} END {for (a in row){print row[a]}}' file.txt