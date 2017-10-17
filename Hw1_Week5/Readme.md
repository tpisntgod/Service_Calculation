### 设计思路
用 flag 包对命令行参数进行解析，用os，bufio.NewReader对文件、os.Stdin进行读取，输出使用os.Stdout.Write()，标准错误输出使用 fmt.Fprintf(os.Stderr, "error_message")。  
1. 命令行参数检查，比较容易，检查参数数量，参数值是否符合要求。
2. 读取部分，通过参数检查判断从键盘读入还是从文件读入，确定rd = bufio.NewReader(os.Stdin) 还是   
rd = bufio.NewReader(file)。通过line, err := rd.ReadString('\n')按行读取文件。  
如果以'\f'作为分页符，则line, err := rd.ReadString('\f')，且加上一句去掉分页符的语句，line = strings.Trim(line, "\f")，将分页符改为换行符，
输出更好看，二者差别不大。其余就是统计页数啊什么的。
3. 管道，使用了 os/exec 包来建立用于进程间通信的管道。
