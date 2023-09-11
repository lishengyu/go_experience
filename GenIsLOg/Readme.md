# 命令行示例
./GenIsLog -comid 2272120 -stime 20230628173910 -etime 20230628173912 -spath ../ -dpath /home/syli -lognum 2

# 补日志方式
1.在spath目录下进行查找，文件名称中时间字符在在给定stime和etime中间文件
2.打开1中匹配的文件名，进行字符串切片，查找是否匹配sep
    ctcc - 0x10 'C' 'U' 'D'
    cucc - 0x03 'X' '1' 'D'
3.在2中匹配的文件中，进行command id查找
4.在3中匹配command id的记录中，查找sport和time的偏移
5.生成在port[10000,55555]中的随机端口，生成在[stime,etime]中的随机时间
6.在4中匹配的记录中，依据sport和time的偏移，使用5中生成的值进行替换
7.生成指定条目数的记录，写入dpath目录下，文件名称使用[stime,etime]的随机值
8.日志文件生成后，创建空的.ok文件，uiap监控该事件取文件
9.文件内容格式为：(使用hexdump查看)
  header + sep + content * N
