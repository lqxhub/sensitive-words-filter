# 敏感词检查

一种用Go实现的敏感词过滤算法。使用DFA算法，进行匹配。

这个敏感词检测使用http 服务来获取输入,然后做校验

使用步骤:

1. `git clone https://github.com/lqxhub/sensitive-words-filter.git`
2. `cd sensitive-words-filter/example`
3. `go build`
4. 运行 example
5. 使用浏览器打开 127.0.01:8081 这个地址

在输入框里输入要检查的词就可以进行检查了

敏感词的规则在 rule.txt 这个文件里,每一行是一个规则

能做到自动去除非中文，英文字符后再进行比对

例如：

有个敏感词是 **我爱你**

那么 **我爱你** **我)爱__你**， **__我爱你__**， **我000爱888你** 等等这些词都会被当成敏感词被检测出来

