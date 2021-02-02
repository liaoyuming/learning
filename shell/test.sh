#!/bin/sh

:<<EOF
    #、## 表示从左边开始删除。一个 # 表示从左边删除到第一个指定的字符；两个 # 表示从左边删除到最后一个指定的字符。
    %、%% 表示从右边开始删除。一个 % 表示从右边删除到第一个指定的字符；两个 % 表示从左边删除到最后一个指定的字符。
    删除包括了指定的字符本身。
EOF
var=http://www.aaa.com/123.htm
echo ${var#*a}
echo ${var##*a}
echo ${var%a*}
echo ${var%%a*}
echo ${var:0:5}
echo ${var:7}
echo ${var:0-7:3}
echo ${var:0-7}

:<<EOF
#获取键盘输入信息
read -p "input a val:" a    #获取键盘输入的 a 变量数字
read -p "input b val:" b    #获取键盘输入的 b 变量数字
r=$[a+b]                    #计算a+b的结果 赋值给r  不能有空格
echo "result = ${r}"        #输出显示结果 r


echo "执行的文件名：$0";
echo "第一个参数为：$1";
echo "第二个参数为：$2";
echo "第三个参数为：$3";
echo "参数个数: $#"
echo '$* 结果:'
for i in "$*"; do
    echo $i
done
echo '$@ 结果:'
for i in "$@"; do
    echo $i
done
EOF

my_array=("A" "B" "C" "D")
echo ${my_array}
