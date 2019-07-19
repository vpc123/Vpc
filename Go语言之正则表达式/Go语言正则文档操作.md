### Go语言正则表达式

##### 1    匹配中文姓名的go语言正则

> req.RealName是中文的字符串

> NameReg,_ := regexp.Match("^[\u4E00-\u9FA5]{2,5}$", []byte(req.RealName))
> 
>    if NameReg{  
>    fmt.Println("正则匹配成功,姓名")  
> }else {  
>    return ErrUserRealNameCode  
> }

##### 2    匹配中国地区的身份证件的正则表达式

> IdCaerdReg1 := regexp.MustCompile(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
> 
> IdCaerdReg2 := regexp.MustCompile(`^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}[0-9Xx]$`)
> 
> if IdCaerdReg1.MatchString(req.IdCardNum)||IdCaerdReg2.MatchString(req.IdCardNum) {  
>    fmt.Println("正则匹配成功,身份证")  
> }else{  
>    return ErrUserRealIdCardCode  
> }

##### 3    匹配中国地区简单的银行卡格式正则

> BankCardNumReg := regexp.MustCompile(`^(\d{16}|\d{17}|\d{18}|\d{19})$`)
> 
> if BankCardNumReg.MatchString(req.BankCardNum){  
>       fmt.Println("银行卡号匹配成功!")  
> }else {  
>    return UserBankCardRepOpCode  
> }

说明:此处举例说明了3种比较常用的正则验证格式表达，身份证号，银行卡，中文姓名的验证，通过正则化的条件判断我们大概了解了go语言解析函数模块使用的一般形式和方法。

#### 正则基础

    用法：
        单一：
         .                   匹配任意一个字符，如果设置 s = true，则可以匹配换行符
    
        [字符类]            匹配“字符类”中的一个字符，“字符类”见后面的说明
        [^字符类]           匹配“字符类”外的一个字符，“字符类”见后面的说明
    
        \小写Perl标记       匹配“Perl类”中的一个字符，“Perl类”见后面的说明
        \大写Perl标记       匹配“Perl类”外的一个字符，“Perl类”见后面的说明
    
        [:ASCII类名:]       匹配“ASCII类”中的一个字符，“ASCII类”见后面的说明
        [:^ASCII类名:]      匹配“ASCII类”外的一个字符，“ASCII类”见后面的说明
    
        \pUnicode普通类名   匹配“Unicode类”中的一个字符(仅普通类)，“Unicode类”见后面的说明
        \PUnicode普通类名   匹配“Unicode类”外的一个字符(仅普通类)，“Unicode类”见后面的说明
    
        \p{Unicode类名}     匹配“Unicode类”中的一个字符，“Unicode类”见后面的说明
        \P{Unicode类名}     匹配“Unicode类”外的一个字符，“Unicode类”见后面的说明

---

复合：

        xy             匹配 xy（x 后面跟随 y）
        x|y            匹配 x 或 y (优先匹配 x)

---

重复：

        x*             匹配零个或多个 x，优先匹配更多(贪婪)
        x+             匹配一个或多个 x，优先匹配更多(贪婪)
        x?             匹配零个或一个 x，优先匹配一个(贪婪)
        x{n,m}         匹配 n 到 m 个 x，优先匹配更多(贪婪)
        x{n,}          匹配 n 个或多个 x，优先匹配更多(贪婪)
        x{n}           只匹配 n 个 x
        x*?            匹配零个或多个 x，优先匹配更少(非贪婪)
        x+?            匹配一个或多个 x，优先匹配更少(非贪婪)
        x??            匹配零个或一个 x，优先匹配零个(非贪婪)
        x{n,m}?        匹配 n 到 m 个 x，优先匹配更少(非贪婪)
        x{n,}?         匹配 n 个或多个 x，优先匹配更少(非贪婪)
        x{n}?          只匹配 n 个 x

---

分组：

        (子表达式)            被捕获的组，该组被编号 (子匹配)
        (?P<命名>子表达式)    被捕获的组，该组被编号且被命名 (子匹配)
        (?:子表达式)          非捕获的组 (子匹配)
        (?标记)               在组内设置标记，非捕获，标记影响当前组后的正则表达式
        (?标记:子表达式)      在组内设置标记，非捕获，标记影响当前组内的子表达式
    
        标记的语法是：
        xyz  (设置 xyz 标记)
        -xyz (清除 xyz 标记)
        xy-z (设置 xy 标记, 清除 z 标记)
    
        可以设置的标记有：
        i              不区分大小写 (默认为 false)
        m              多行模式：让 ^ 和 $ 匹配整个文本的开头和结尾，而非行首和行尾(默认为 false)
        s              让 . 匹配 \n (默认为 false)
        U              非贪婪模式：交换 x* 和 x*? 等的含义 (默认为 false)

---

位置标记：

        ^              如果标记 m=true 则匹配行首，否则匹配整个文本的开头（m 默认为 false）
        $              如果标记 m=true 则匹配行尾，否则匹配整个文本的结尾（m 默认为 false）
        \A             匹配整个文本的开头，忽略 m 标记
        \b             匹配单词边界
        \B             匹配非单词边界
        \z             匹配整个文本的结尾，忽略 m 标记

---

转义序列：

        \a             匹配响铃符    （相当于 \x07）
                       注意：正则表达式中不能使用 \b 匹配退格符，因为 \b 被用来匹配单词边界，
                       可以使用 \x08 表示退格符。
        \f             匹配换页符    （相当于 \x0C）
        \t             匹配横向制表符（相当于 \x09）
        \n             匹配换行符    （相当于 \x0A）
        \r             匹配回车符    （相当于 \x0D）
        \v             匹配纵向制表符（相当于 \x0B）
        \123           匹配 8  進制编码所代表的字符（必须是 3 位数字）
        \x7F           匹配 16 進制编码所代表的字符（必须是 3 位数字）
        \x{10FFFF}     匹配 16 進制编码所代表的字符（最大值 10FFFF  ）
        \Q...\E        匹配 \Q 和 \E 之间的文本，忽略文本中的正则语法
    
        \\             匹配字符 \
        \^             匹配字符 ^
        \$             匹配字符 $
        \.             匹配字符 .
        \*             匹配字符 *
        \+             匹配字符 +
        \?             匹配字符 ?
        \{             匹配字符 {
        \}             匹配字符 }
        \(             匹配字符 (
        \)             匹配字符 )
        \[             匹配字符 [
        \]             匹配字符 ]
        \|             匹配字符 |

---

可以将“命名字符类”作为“字符类”的元素：

        [\d]           匹配数字 (相当于 \d)
        [^\d]          匹配非数字 (相当于 \D)
        [\D]           匹配非数字 (相当于 \D)
        [^\D]          匹配数字 (相当于 \d)
        [[:name:]]     命名的“ASCII 类”包含在“字符类”中 (相当于 [:name:])
        [^[:name:]]    命名的“ASCII 类”不包含在“字符类”中 (相当于 [:^name:])
        [\p{Name}]     命名的“Unicode 类”包含在“字符类”中 (相当于 \p{Name})
        [^\p{Name}]    命名的“Unicode 类”不包含在“字符类”中 (相当于 \P{Name})

---

说明：

---

“字符类”取值如下（“字符类”包含“Perl类”、“ASCII类”、“Unicode类”）：
    x                    单个字符
    A-Z                  字符范围(包含首尾字符)
    \小写字母            Perl类
    [:ASCII类名:]        ASCII类
    \p{Unicode脚本类名}  Unicode类 (脚本类)
    \pUnicode普通类名    Unicode类 (普通类)

---

“Perl 类”取值如下：

    \d             数字 (相当于 [0-9])
    \D             非数字 (相当于 [^0-9])
    \s             空白 (相当于 [\t\n\f\r ])
    \S             非空白 (相当于[^\t\n\f\r ])
    \w             单词字符 (相当于 [0-9A-Za-z_])
    \W             非单词字符 (相当于 [^0-9A-Za-z_])

---

“ASCII 类”取值如下

    [:alnum:]      字母数字 (相当于 [0-9A-Za-z])
    [:alpha:]      字母 (相当于 [A-Za-z])
    [:ascii:]      ASCII 字符集 (相当于 [\x00-\x7F])
    [:blank:]      空白占位符 (相当于 [\t ])
    [:cntrl:]      控制字符 (相当于 [\x00-\x1F\x7F])
    [:digit:]      数字 (相当于 [0-9])
    [:graph:]      图形字符 (相当于 [!-~])
    [:lower:]      小写字母 (相当于 [a-z])
    [:print:]      可打印字符 (相当于 [ -~] 相当于 [ [:graph:]])
    [:punct:]      标点符号 (相当于 [!-/:-@[-反引号{-~])
    [:space:]      空白字符(相当于 [\t\n\v\f\r ])
    [:upper:]      大写字母(相当于 [A-Z])
    [:word:]       单词字符(相当于 [0-9A-Za-z_])
    [:xdigit:]     16 進制字符集(相当于 [0-9A-Fa-f])

---

“Unicode 类”取值如下---普通类：

    C                 -其他-          (other)
    Cc                控制字符        (control)
    Cf                格式            (format)
    Co                私人使用区      (private use)
    Cs                代理区          (surrogate)
    L                 -字母-          (letter)
    Ll                小写字母        (lowercase letter)
    Lm                修饰字母        (modifier letter)
    Lo                其它字母        (other letter)
    Lt                首字母大写字母  (titlecase letter)
    Lu                大写字母        (uppercase letter)
    M                 -标记-          (mark)
    Mc                间距标记        (spacing mark)
    Me                关闭标记        (enclosing mark)
    Mn                非间距标记      (non-spacing mark)
    N                 -数字-          (number)
    Nd                十進制数字      (decimal number)
    Nl                字母数字        (letter number)
    No                其它数字        (other number)
    P                 -标点-          (punctuation)
    Pc                连接符标点      (connector punctuation)
    Pd                破折号标点符号  (dash punctuation)
    Pe                关闭的标点符号  (close punctuation)
    Pf                最后的标点符号  (final punctuation)
    Pi                最初的标点符号  (initial punctuation)
    Po                其他标点符号    (other punctuation)
    Ps                开放的标点符号  (open punctuation)
    S                 -符号-          (symbol)
    Sc                货币符号        (currency symbol)
    Sk                修饰符号        (modifier symbol)
    Sm                数学符号        (math symbol)
    So                其他符号        (other symbol)
    Z                 -分隔符-        (separator)
    Zl                行分隔符        (line separator)
    Zp                段落分隔符      (paragraph separator)
    Zs                空白分隔符      (space separator)

---





#### 代码:

    // 示例
    func main() {
     text := `Hello 世界！123 Go.`
    // 查找连续的小写字母
    reg := regexp.MustCompile(`[a-z]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["ello" "o"]
    
    // 查找连续的非小写字母
    reg = regexp.MustCompile(`[^a-z]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["H" " 世界！123 G" "."]
    
    // 查找连续的单词字母
    reg = regexp.MustCompile(`[\w]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello" "123" "Go"]
    
    // 查找连续的非单词字母、非空白字符
    reg = regexp.MustCompile(`[^\w\s]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["世界！" "."]
    
    // 查找连续的大写字母
    reg = regexp.MustCompile(`[[:upper:]]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["H" "G"]
    
    // 查找连续的非 ASCII 字符
    reg = regexp.MustCompile(`[[:^ascii:]]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["世界！"]
    
    // 查找连续的标点符号
    reg = regexp.MustCompile(`[\pP]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["！" "."]
    
    // 查找连续的非标点符号字符
    reg = regexp.MustCompile(`[\PP]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello 世界" "123 Go"]
    
    // 查找连续的汉字
    reg = regexp.MustCompile(`[\p{Han}]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["世界"]
    
    // 查找连续的非汉字字符
    reg = regexp.MustCompile(`[\P{Han}]+`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello " "！123 Go."]
    
    // 查找 Hello 或 Go
    reg = regexp.MustCompile(`Hello|Go`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello" "Go"]
    
    // 查找行首以 H 开头，以空格结尾的字符串
    reg = regexp.MustCompile(`^H.*\s`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello 世界！123 "]
    
    // 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
    reg = regexp.MustCompile(`(?U)^H.*\s`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello "]
    
    // 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
    reg = regexp.MustCompile(`(?i:^hello).*Go`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello 世界！123 Go"]
    
    // 查找 Go.
    reg = regexp.MustCompile(`\QGo.\E`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Go."]
    
    // 查找从行首开始，以空格结尾的字符串（非贪婪模式）
    reg = regexp.MustCompile(`(?U)^.* `)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello "]
    
    // 查找以空格开头，到行尾结束，中间不包含空格字符串
    reg = regexp.MustCompile(` [^ ]*$`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // [" Go."]
    
    // 查找“单词边界”之间的字符串
    reg = regexp.MustCompile(`(?U)\b.+\b`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello" " 世界！" "123" " " "Go"]
    
    // 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
    reg = regexp.MustCompile(`[^ ]{1,4}o`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello" "Go"]
    
    // 查找 Hello 或 Go
    reg = regexp.MustCompile(`(?:Hell|G)o`)
    fmt.Printf("%q\n", reg.FindAllString(text, -1))
    // ["Hello" "Go"]
    
    // 查找 Hello 或 Go，替换为 Hellooo、Gooo
    reg = regexp.MustCompile(`(?PHell|G)o`)
    fmt.Printf("%q\n", reg.ReplaceAllString(text, "${n}ooo"))
    // "Hellooo 世界！123 Gooo."
    
    // 交换 Hello 和 Go
    reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
    fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
    // "Go 世界！123 Hello."
    
    // 特殊字符的查找
    reg = regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]\|]`)
    fmt.Printf("%q\n", reg.ReplaceAllString("\f\t\n\r\v\123\x7F\U0010FFFF\\^$.*+?{}()[]|", "-"))
    // "----------------------"
    }

### 总结:

Go语言的正则表达式非常容易在工作中使用到，通过正则表达式的书写方式深入了解Go语言的具体使用技巧增加对传入参数的判断要求，精确参数的合法正规性，在以后的工作生活中将会大有裨益。
