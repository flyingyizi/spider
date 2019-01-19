## 0."一个三十" 这种如何分开
Jan 9, 2019
```text
 `text = "给我办一个三十的新流量业务"
result = []
for (word, start, end) in jieba.tokenize(text):
pseg_data = [(w, f) for (w, f) in pseg.cut(word)]
print((pseg_data, start, end))
result.append((pseg_data, start, end))
`
([('给', 'p')], 0, 1)
([('我', 'r')], 1, 2)
([('办', 'v')], 2, 3)
([('一个三十', 'm')], 3, 7)
([('的', 'uj')], 7, 8)
([('新', 'a')], 8, 9)
([('流量', 'n')], 9, 11)
([('业务', 'n')], 11, 13)
如果是"两个三十"就可以分开
```
## 1.频率最大算法的正向与反向
Jan 6, 2019
```text
 您好；
我看到您代码的时候，您的动态规划求频率最大是反向的（也就是从最后一个字符开始算），我自己写了一下，正向的频率最大（也就是从第一个字符开始算）我跑了跑，发现结果是一样的，请问一下这个可以证明吗？还是我操作失误了，还有就是不是说汉语的重点在后面吗？是因为这个用的反向频率最大吗？谢谢
```
## 2.print(jieba.lcut("a b c"))
Dec 17, 2018
```text
 ['a', ' ', 'b', ' ', 'c']
出来空格。。
```
## 3.如何评估 extract_tags 结果的表现
Dec 17, 2018
```text
 关键词提取无论如何总能得到一些结果，在语料非常多的时候，人工标记不太具备可行性，那么有哪些推荐的算法用于评估提取内容的准确性和相关性呢？
```
## 4.suggest_freq无法处理特殊符号
Dec 11, 2018
```text
 import jieba

jieba.suggest_freq("{UM}", tune=True)
jieba.lcut("{UM}卡上写的地址就是那个", HMM=False)
# ['{', 'UM', '}', '卡', '上', '写', '的', '地址', '就是', '那个']
jieba.suggest_freq("卡上", tune=True)
jieba.lcut("{UM}卡上写的地址就是那个", HMM=False)
# ['{', 'UM', '}', '卡上', '写', '的', '地址', '就是', '那个']
suggest_freq只能处理全部为中文的情况, 如果希望把特殊符号识别成一个词就会出错.
```
## 5.怎么保存自己添加的词到词典？
Dec 6, 2018
```text
 我想把新的词汇保存到txt文档里，请问大神怎么做？？？
```
## 6.jieba.initialize()函数无法处理带中文路径
Nov 25, 2018
```text
 通过 jieba.set_dictionary("."+os.sep+"dict.txt") 指定dict.txt 文件位于一个带中文的父目录时，使用jieba.initialize() 进行初始化报错，报错为：
Traceback (most recent call last):
File "", line 26, in 
File "F:\Qiyi\build\getcalldatafrom_txt\out00-PYZ.pyz\jieba", line 111, in initialize
UnicodeDecodeError: 'ascii' codec can't decode byte 0xcc in position 3: ordinal not in range(128)
定位于语句： jieba_init.py_ 中的 111行： default_logger.debug("Building prefix dict from %s ..." % (abs_path or 'the default dictionary'))
将程序至于全英文目录时，运行正常。
```
## 7.TF_IDF分析三国演义的地名，为何结果完全不对？
Nov 21, 2018
```text
 部分代码：
sentence = getText('三国演义.txt')
keywords = jieba.analyse.extract_tags(sentence, topK=20, withWeight=True, allowPOS=('ns'))
按地名统计前20，出来的结果是：
将军 0.11860080694708297
丞相 0.0921638281251496
主公 0.0715118726461463
军士 0.061001657055284146
商议 0.05922492016471381
云长 0.054936383725351035
军马 0.05441077997191448
大喜 0.05001738396254889
后主 0.04707253654544958
先主 0.04402888433316952
都督 0.04296257804954991
众将 0.04167988860499132
天下 0.03893417374252234
陛下 0.03882608055724891
太守 0.035023324694504726
人马 0.03359843706149093
城上 0.03265662810340311
天子 0.03212681223959663
后人 0.03164288695617565
众官 0.03082541531255808
这是怎么统计出来的地名？完全不对啊！
```
## 8.jieba似乎无法把长痘痘这个词的长和痘痘分开？
Nov 20, 2018
```text
 把痘痘加入自定词典后，也无法把长痘痘中的长和痘痘分开。各位大佬有什么办法嘛？
```
## 9.jieba 分词格式有那些，并且怎么动态加载分词呢？
Nov 19, 2018
```text
 No description provided.
```
## 10.某些领域中，自己设定的特定分词效果如何加入
Nov 17, 2018
```text
 比如某个句子的分词效果，结巴的结果和预期有出入
我想把人工标注的结果加入到结巴，如何？
```
## 11.请教: 自定义字典里的分词不能被识别的问题
Nov 10, 2018
```text
 原字典自带的dict.txt里面 '阳台' 的词频是 242 , 我在我的自定义字典里, 定义:  阳台栏杆 500  阳台天花 500
但是在分词的时候,依然会把阳台有关的词分词为:  阳台 / 栏杆  阳台/天花
```
## 12.这个ns(地名)词性的识别有好多问题啊
Nov 10, 2018
```text
 比如"中断" "满格" 等会莫名奇妙的识别为ns,ns词性是地名词性，这也太假了
```
## 13.请问支持先分词，再基于分词结果进行词性标注吗？
Nov 7, 2018
```text
 我想先分完词，再基于分词结果做词性标注，请问结巴分词支持吗？我看哈工大的pyltp是支持的（[https://blog.csdn.net/baidu_15113429/article/details/78909666]），但是我不想使用pyltp
```
```text
 這裡好像快倒了 沒啥人氣 慘....
```
```text
 https://blog.csdn.net/enter89/article/details/80619805
```
```text
 參考我修正的項目 #670
```
## 14.如何验证分词的准确性问题
Nov 4, 2018
```text
 如何验证分词的准确性
请问用jieba分完词后，如何去验证分词的准确性？
```
```text
 Same question. What is the training data? 人民日报语料吗？
```
```text
 Same question. What is the training data? 人民日报语料吗？
不是人民日报的，是网上爬取的一些评论数据，最近看了数学之美的分词，有一些启发
```
## 15.jieba分词是否支持依存句法分析？
Oct 30, 2018
```text
 如题，多谢！
```
## 16.特殊分词控制
Oct 27, 2018
```text
 sent = "基础上的"
我想把它分词成 基础上/的 。
但是不切分“上的”。
应该怎么实现这样的功能呢？
类似的词还有 “既要有” “越要有”
想分成 既要/有 越要/有
```
```text
 drop word list





来自钉钉专属商务邮箱------------------------------------------------------------------
发件人：OrangeIUH<notifications@github.com>
日　期：2018年10月27日 15:45:09
收件人：fxsjy/jieba<jieba@noreply.github.com>
抄　送：Subscribed<subscribed@noreply.github.com>
主　题：[fxsjy/jieba] 特殊分词控制 (#680)

sent = "基础上的"
 我想把它分词成 基础上/的 。
 但是不切分“上的”。
 应该怎么实现这样的功能呢？
 类似的词还有 “既要有” “越要有”
 想分成 既要/有 越要/有
—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
## 17.能否替换jieba使用的词库，而不是补充jieba的词库？
Oct 24, 2018
```text
 需求：
要求匹配出内容的指定的一些关键词（我自己的词库里的词），忽略jieba里千千万万不相关的词，对于我来说，这些词没有用。
实现：
我想的是，把jieba默认的词库替换掉。可是，我看文档里没有提到如何替换jieba的词库
还有什么好的办法实现吗？
谢谢各位
```
```text
 你可以用jieba的自定义字典实现准确识别

中铁建资产管理有限公司 高正  18601064889
…
 在 2018年10月24日，11:11，Bakkan Hwang ***@***.***> 写道：

 需求：
 要求匹配出内容的指定的一些关键词（我自己的词库里的词），忽略jieba里千千万万不相关的词，对于我来说，这些词没有用。

 实现：
 我想的是，把jieba默认的词库替换掉。可是，我看文档里没有提到如何替换jieba的词库
 还有什么好的办法实现吗？

 谢谢各位

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 jieba有实现延时加载机制
jieba 采用延迟加载，import jieba 和 jieba.Tokenizer() 不会立即触发词典的加载，一旦有必要才开始加载词典构建前缀字典。如果你想手工初始 jieba，也可以手动初始化。
import jieba
jieba.initialize()  # 手动初始化（可选）
在 0.28 之前的版本是不能指定主词典的路径的，有了延迟加载机制后，你可以改变主词典的路径:
jieba.set_dictionary('data/dict.txt.big')
例子： https://github.com/fxsjy/jieba/blob/master/test/test_change_dictpath.py
```
## 18.TFIDF算法要取多少topK
Oct 20, 2018
```text
 如題，TFIDF算法一般來說topK這個參數要取多少比較合適?我目前使用的題材是金庸的小說，目前正在切神鵰俠侶這本。
```
## 19.请问，官方什么时候能出一个PHP版本呢？因为分词可能是对网页传来的信息直接就分词了
Oct 18, 2018
```text
 请问，官方什么时候能出一个PHP版本呢？因为分词可能是对网页传来的信息直接就分词了，而不是一直读数据库的
```
## 20.地名部分词性不准确
Oct 14, 2018
```text
 部分地名词性为nr，比如“孙吴县”。建议按照民政部最新的行政区划表更新词库。
2018行政区划代码
```
## 21.如何将识别出的新词加入用户字典？
Oct 8, 2018
```text
 `



seg_list = jieba.cut("他来到了网易杭研大厦")  # 默认是精确模式
print(", ".join(seg_list))
他, 来到, 了, 网易, 杭研, 大厦`



但是在分词过程中判断词性
`



words = pseg.cut("来到了网易杭研大厦")
for word, flag in words:
...  print('%s %s' % (word, flag))`



依旧会把“杭研”分为两个词  杭 j 研 vn  ；
如何判断“杭研”是否为新词？
```
```text
 可以參考我修正的錯誤 #670
```
## 22.请问如何定义如“建设银行”=“建行”这类，以便在统计或者模糊匹配的时候被定义为同一个单词？
Oct 1, 2018
```text
 rt
```
## 23.添加自定义词 为何切不出来呢
Sep 25, 2018
```text
 jieba.add_word("β细胞", freq=True, tag='n')
a = '揭晓！胰岛β细胞功能可以恢复吗?'
WORD = jieba.lcut(a,cut_all=False)
print ([word for word in WORD])
['揭晓', '！', '胰岛', 'β', '细胞', '功能', '可以', '恢复', '吗', '?']
```
```text
 freq的值设置高一点？
```
```text
 freq的值设置高一点？

谢谢，查到了 好像是和特殊字符有关 修改了源代码
```
## 24.使用 add_word 和 suggest_freq 后，词语还是被拆分
Sep 21, 2018
```text
 In [94]: s = '乌鲁木齐爱家超市南门店'

In [95]: jieba.cut(s)
Out[95]: <generator object Tokenizer.cut at 0x10e24fde0>

In [96]: jieba.lcut(s)
Out[96]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [97]: jieba.add_word('南门')

In [98]: jieba.lcut(s)
Out[98]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [99]: jieba.suggest_freq('南门', True)
Out[99]: 833

In [100]: jieba.lcut(s)
Out[100]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [101]: jieba.lcut(s, HMM=False)
Out[101]: ['乌鲁木齐', '爱家', '超市', '南', '门店']
期望结果
['乌鲁木齐', '爱家', '超市', '南门', '店']
```
```text
 In [94]: s = '乌鲁木齐爱家超市南门店'

In [95]: jieba.cut(s)
Out[95]: <generator object Tokenizer.cut at 0x10e24fde0>

In [96]: jieba.lcut(s)
Out[96]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [97]: jieba.add_word('南门')

In [98]: jieba.lcut(s)
Out[98]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [99]: jieba.suggest_freq('南门', True)
Out[99]: 833

In [100]: jieba.lcut(s)
Out[100]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [101]: jieba.lcut(s, HMM=False)
Out[101]: ['乌鲁木齐', '爱家', '超市', '南', '门店']
期望结果
['乌鲁木齐', '爱家', '超市', '南门', '店']

@4ft35t 设置的频率还是太低了，使用jieba.add_word("南门", freq=1000)设置高一点就可以了。
```
## 25.研究发现，5-HT功能活动降低与抑郁症患者的抑郁心境、食欲减退、失眠、昼夜节律紊乱、内分泌功能紊乱、性功能障碍、焦虑不安、不能应付应激、活动减少等密切相关；而5-HT功能增高与躁狂症的发病有关。
Sep 6, 2018
```text
 Does anyone have a list of the POS tags, and what they stand for?  They're supposed to be compatible with ictclas.  Thanks.
```
```text
 使用：
jieba.add_word("5-HT功能", freq=True, tag='n')
jieba.add_word("食欲减退", freq=True, tag='n')
jieba.add_word("昼夜节律紊乱", freq=True, tag='n')
然后分词处理：
print("/".join(jieba.cut(sta, HMM=True)))
发现结果为：
研究/发现/，/5/-/HT/功能/活动/降低/与/抑郁症/患者/的/抑郁/心境/、/食欲减退/、/失眠/、/昼夜节律紊乱/、/内分泌/功能/紊乱/、/性功能
障碍/、/焦虑不安/、/不能/应付/应激/、/活动/减少/等/密切相关/；/而/5/-/HT/功能/增高/与/躁狂症/的/发病/有关/。
预计结果是：
研究/发现/，/5-HT功能/活动/降低/与/抑郁症/患者/的/抑郁/心境/、/食欲减退/、/失眠/、/昼夜节律紊乱/、/内分泌/功能/紊乱/、/性功能
障碍/、/焦虑不安/、/不能/应付/应激/、/活动/减少/等/密切相关/；/而/5-HT功能/增高/与/躁狂症/的/发病/有关/。
“5-HT功能”，为什么没有用？
```
```text
 因为有“-”，所以对两边的数据进行了切分
```
```text
 那怎么解决这个问题。因为有些专业名里面包含这种“-”



liwy@futongdf.com.cn

发件人： LuoZe
发送时间： 2018-09-06 21:22
收件人： fxsjy/jieba
抄送： lwy1111111; Author
主题： Re: [fxsjy/jieba] 研究发现，5-HT功能活动降低与抑郁症患者的抑郁心境、食欲减退、失眠、昼夜节律紊乱、内分泌功能紊乱、性功能障碍、焦虑不安、不能应付应激、活动减少等密切相关；而5-HT功能增高与躁狂症的发病有关。 (#666)
因为有“-”，所以对两边的数据进行了切分
—
You are receiving this because you authored the thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 from pyhanlp import *

CustomDictionary.insert('5-HT功能', '医疗术语 10')
print(HanLP.segment('研究发现，5-HT功能活动降低与抑郁症患者的抑郁心境'))
[研究/vn, 发现/v, ，/w, 5-HT功能/医疗术语, 活动/vn, 降低/v, 与/cc, 抑郁症/nhd, 患者/n, 的/ude1, 抑郁/a, 心境/n]
```
```text
 谢谢，但是我问的是结巴的解决方法！
```
```text
 谢谢，但是我问的是结巴的解决方法！

之前有帖子写了，特殊字符，可以考虑修改
jieba.re_han_default = re.compile(r'([\u4e00-\u9fa5a-zA-Z0-9+#&._%/-]+)', re.UNICODE)
```
## 26.连在一起的数词和量词可以分到一起去吗？
Sep 1, 2018
```text
 您好；
我看到您代码的时候，您的动态规划求频率最大是反向的（也就是从最后一个字符开始算），我自己写了一下，正向的频率最大（也就是从第一个字符开始算）我跑了跑，发现结果是一样的，请问一下这个可以证明吗？还是我操作失误了，还有就是不是说汉语的重点在后面吗？是因为这个用的反向频率最大吗？谢谢
```
```text
 比如：
s = u'''出租 珠江新城 13楼独立90方 2房2 全配套 月8400元''' cut = jieba.cut(s) print ','.join(cut)
结果是
出租, ,珠江新城, ,13,楼,独立,90,方, ,2,房,2, ,全,配套, ,月,8400,元
有可能把13楼、90方、2房、2房2、月8400元给单独划分出来吗？
```
```text
 可以,但需要你自己单独提供词库

lfol <notifications@github.com> 于2018年9月1日周六 下午4:36写道：
…
 比如：
 s = u'''出租 珠江新城 13楼独立90方 2房2 全配套 月8400元''' cut = jieba.cut(s) print
 ','.join(cut)

 结果是
 出租, ,珠江新城, ,13,楼,独立,90,方, ,2,房,2, ,全,配套, ,月,8400,元

 有可能把13楼、90方、2房、2房2、月8400元给单独划分出来吗？

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#665>, or mute the thread
 <https://github.com/notifications/unsubscribe-auth/AaXy0hBiwe3XSNWR8GFIFhy4w56w2teuks5uWkb_gaJpZM4WWIsV>
 .
```
```text
 from pyhanlp import *

segment = HanLP.newSegment().enableNumberQuantifierRecognize(True)
sentences = [
    "十九元套餐包括什么",
    "九千九百九十九朵玫瑰",
    "壹佰块都不给我",
    "９０１２３４５６７８只蚂蚁",
    "牛奶三〇〇克*2",
    "ChinaJoy“扫黄”细则露胸超2厘米罚款",
]
for sentence in sentences:
    print(segment.seg(sentence))
[十九元/mq, 套餐/n, 包括/v, 什么/ry]
[九千九百九十九朵/mq, 玫瑰/n]
[壹佰块/mq, 都/d, 不/d, 给/p, 我/rr]
[９０１２３４５６７８只/mq, 蚂蚁/n]
[牛奶/nf, 三〇〇克/mq, */w, 2/m]
[ChinaJoy/nx, “/w, 扫黄/vi, ”/w, 细则/n, 露/v, 胸/ng, 超/v, 2厘米/mq, 罚款/vi]
```
## 27.为什么不给posseg添加cut_for_search?
Aug 15, 2018
```text
 我想把新的词汇保存到txt文档里，请问大神怎么做？？？
```
```text
 No description provided.
```
## 28.自定义词包含字母与数字导致新词不生效
Aug 13, 2018
```text
 比如将XXXX年、XX月视作一个词
或者我需要将书名号内所有内容视作整个词
能不能进行自定义配置？
```
```text
 用字典不就行了。匹配正则，想下这效率得有多慢。
```
```text
 Environment:

jieba v0.39

Code:
import jieba
jieba.add_word("现代汉语文本切分与词性标注规范Ｖ1.0");
seg_list = jieba.cut("北大计算语言学研究所从 1992 年开始进行汉语语料库的多级加工研究。第一步是对原\n" +
                "始语料进行切分和词性标注。1994 年制订了《现代汉语文本切分与词性标注规范Ｖ1.0》。")
print(','.join(seg_list))

output:
北大,计算,语言学,研究所,从, ,1992, ,年,开始,进行,汉语,语料库,的,多级,加工,研究,。,第一步,是,对,原,
,始,语料,进行,切分,和,词性,标注,。,1994, ,年,制订,了,《,现代汉语,文本,切分,与,词性,标注,规范,Ｖ,1.0,》,。
```
```text
 根据个人需要修改jieba包init.py中几个正则表达式，使其支持数字特殊字符。re_han_default = re.compile(“(.+)”, re.U)
```
```text
 from pyhanlp import *

segment = HanLP.newSegment('感知机')
CustomDictionary.insert('现代汉语文本切分与词性标注规范Ｖ1.0')
print(segment.analyze("北大计算语言学研究所从1992年开始进行汉语语料库的多级加工研究。第一步是对原" +
                      "始语料进行切分和词性标注。1994年制订了《现代汉语文本切分与词性标注规范Ｖ1.0》。"))

[北大/j 计算/v 语言学/n 研究所/n]/nt 从/p 1992年/t 开始/v 进行/v 汉语/nz 语料库/n 的/u 多/m 级/q 加工/v 研究/v 。/w 第一/m 步/q 是/v 对/p 原始/a 语/Ng 料/v 进行/v 切分/vn 和/c 词性/n 标注/v 。/w 1994年/t 制订/v 了/u 《/w 现代汉语文本切分与词性标注规范Ｖ1.0/nz 》/w 。/w
```
## 29.形容词某些情况下变成了名词的问题反馈
Aug 13, 2018
```text
 作为一个大四学生，刚刚学过python，想看看一些成熟的项目具体代码，透彻地进行分析，不知道jieba适不适合呢？
```
```text
 個人感覺挺好的, http://www.cnblogs.com/zhbzz2007/tag/Natural%20language%20processing/
```
```text
 形容词某些情况下变成了名词了 如下所示
# 新 a
>>> pseg.lcut("新材料装备制造业的循环产业体系")
[pair('新', 'a'), pair('材料', 'n'), pair('装备', 'n'), pair('制造业', 'n'), pair('的', 'uj'), pair('循环', 'vn'), pair('产业', 'n'), pair('体系', 'n')]
# 新 n
>>> pseg.lcut("青海格尔木工业园镁锂新材料产业基地")
[pair('青海', 'ns'), pair('格尔木', 'nr'), pair('工业园', 'n'), pair('镁锂', 'nz'), pair('新', 'n'), pair('材料', 'n'), pair('产业基地', 'n')]


# 大 a
>>> pseg.lcut("大工业自由")
[pair('大', 'a'), pair('工业', 'n'), pair('自由', 'a')]
# 大 n
>>> pseg.lcut("7个国家的大公司中")
[pair('7', 'm'), pair('个', 'm'), pair('国家', 'n'), pair('的', 'uj'), pair('大', 'n'), pair('公司', 'n'), pair('中', 'f')]

有什么办法可以使其保持一致吗？
```
```text
 参考hanlp
```
```text
 可以參考我修正的錯誤 #670
測試完 您上述的語句都正常
(可能跟我修正的錯誤有關係)
#'新', 'a'
>>> pseg.lcut("新材料装备制造业的循环产业体系")
[pair('新', 'a'), pair('材料', 'n'), pair('装备', 'n'), pair('制造业', 'n'), pair('的', 'uj'), pair('循环', 'vn'), pair('产业', 'n'), pair('体系', 'n')]
#新,'n' -> #'新', 'a'
>>> pseg.lcut("青海格尔木工业园镁锂新材料产业基地")
[pair('青海', 'ns'), pair('格尔木', 'nr'), pair('工业园', 'n'), pair('镁', 'n'), pair('锂', 'n'), pair('新', 'a'), pair('材料', 'n'), pair('产业基地', 'n')]

#'大', 'a'
pseg.lcut("大工业自由")
[pair('大', 'a'), pair('工业', 'n'), pair('自由', 'a')]

#'大', 'n' -> #'大', 'a'
pseg.lcut("7个国家的大公司中")
[pair('7', 'm'), pair('个', 'm'), pair('国家', 'n'), pair('的', 'uj'), pair('大', 'a'), pair('公司', 'n'), pair('中', 'f')]
```
## 30.请问下同义词该怎么处理呢？
Aug 10, 2018
```text
 import jieba

jieba.suggest_freq("{UM}", tune=True)
jieba.lcut("{UM}卡上写的地址就是那个", HMM=False)
# ['{', 'UM', '}', '卡', '上', '写', '的', '地址', '就是', '那个']
jieba.suggest_freq("卡上", tune=True)
jieba.lcut("{UM}卡上写的地址就是那个", HMM=False)
# ['{', 'UM', '}', '卡上', '写', '的', '地址', '就是', '那个']
suggest_freq只能处理全部为中文的情况, 如果希望把特殊符号识别成一个词就会出错.
```
```text
 请问下同义词该怎么处理呢？
如：
北京大学,北大,pku
清华大学,清华,Tsinghua University
```
## 31.同时加载多个user_dict，不同词性的相同词语存在覆盖的情况，怎么解决？
Aug 6, 2018
```text
 如圖所示

我是用簡體中文維基百科的語料，經過處理之後把符號換成 space，再將整個文本當成一個當成一個 string
餵進去。
下面是我的代碼
import time

start_time = time.time()

jieba.enable_parallel(16)
jieba_words = {word for word in jieba.cut_for_search(text_processed, HMM=True) if len(word) > 1 and len(word) <= 7}

print("Segmentizing took {} seconds.".format(time.time() - start_time))

謝謝幫忙！
```
```text
 一直在等待一个Ruby版本的jieba，求大神来开发一个！
```
```text
 No description provided.
```
```text
 比如：
高级Java工程师 100 position
高级Java工程师 110 certificate
词频是乱填的，两者来自不同的user_dict,会产生覆盖的情况
```
## 32.结巴多线程分词怎么比单线程更慢呢？
Aug 1, 2018
```text
 参考Issue: #387
如果直接在spark-stream的flatMap中执行jieba.cut()分词，则因“发现该dict加载后是以python字典的数据类型存在于spark的driver里面，但是worker进程无法共享这段内存。By @liaicheng ”，导致每次执行时多次“Building prefix dict from the default dictionary ……”，影响了分词性能。
如果“建议用jieba.Tokenizer 得到一个分词器实例，然后再调用相关方法。By @fxsjy ”，则对tokr实例进行广播变量（sc.broadcast(tokr)）时，报“can't pickle thread.lock objects”的错误。
考虑到spark是采用pickle进行持久化，参考pyahocorasick（https://pypi.python.org/pypi/pyahocorasick/ ）的实现方式，
You can also create an eventually large automaton ahead of time and pickle it to re-load later. Here we just pickle to a string. You would typically pickle to a file instead:



import cPickle
pickled = cPickle.dumps(A)
B = cPickle.dumps(pickled)
B.get('he')
(0, 'he')



jieba也能否优化以更好地支持序列化从而进一步提高在spark中的执行效率？
```
```text
 结巴不支持多线程分词吗？为什么多线程分词还要更慢呢？
缘起
对14206篇文章进行相似性去重 首先要分词 总用时12分钟 发现光分词就用了8分钟
所以想通过多线程分词来加快处理速度 如5个线程 每个线程处理2841篇文章
如下所示
def segment(texts_, start, end):
    print(start, end)
    texts_cut = []
    for t in texts_[start:end]:
        if not isinstance(t, str):
            t = ""
        t = regex.sub("", t)
        cut = list(jieba.cut(t))
        texts_cut.append(cut)

    return texts_cut

max_workers = 5
batch_count = len(texts)//max_workers
logging.info("start")
with concurrent.futures.ThreadPoolExecutor(max_workers=max_workers) as executor:
    future_2_start = {executor.submit(segment, texts, s*batch_count, (s+1)*batch_count if s < max_workers-1 else len(texts)) : s for s in range(max_workers)}
    start_2_data = {}
    for future in concurrent.futures.as_completed(future_2_start):
        start = future_2_start[future]
        data = future.result()
        start_2_data[start] = data

    texts_cut = []
    for start in sorted(start_2_data):
        texts_cut.extend(start_2_data[start])

    print(len(texts_cut))
    
logging.info("done")

于是用了5193篇文章做了实验
单线程（max_workers=1）用时约43秒
2018-08-01 19:10:14,529 : INFO : start
2018-08-01 19:10:57,450 : INFO : done

5个线程（max_workers=5）用时约50秒
2018-08-01 19:11:38,841 : INFO : start
2018-08-01 19:12:29,616 : INFO : done

为什么多线程分词还要更慢呢？
```
```text
 原因

The GIL is an interpreter-level lock. This lock prevents execution of multiple threads at once in the Python interpreter. Each thread that wants to run must wait for the GIL to be released by the other thread, which means your multi-threaded Python application is actually single threaded. The GIL prevents simultaneous access to Python objects by multiple threads.
摘自：https://medium.com/practo-engineering/threading-vs-multiprocessing-in-python-7b57f224eadb

解决方法：

But of course, we would want to use the ProcessPoolExecutor for CPU intensive tasks. The ThreadPoolExecutor is better suited for network operations or I/O
摘自： http://masnun.com/2016/03/29/python-a-quick-introduction-to-the-concurrent-futures-module.html

上述代码 ThreadPoolExecutor ==》 ProcessPoolExecutor 即可
注意：
先显式初始化 否则每个进程都要初始化
jieba.initialize()

用时
# 多进程 ProcessPoolExecutor 8个进程
2018-08-11 10:38:34,569 : INFO : start
2018-08-11 10:38:47,999 : INFO : done took time: 12.43

2018-08-11 10:40:23,051 : INFO : start
2018-08-11 10:40:35,909 : INFO : done took time: 11.95
```
```text
 其实jieba自带了并行分词，只需要jieba.enable_parallel()
```
## 33.更新词典，缓存jieba.cache是否会更新?
Jul 30, 2018
```text
 【词频省略时使用自动计算的能保证分出该词的词频】这句话不太明白。
词频省略的时候，要如何自动计算词频呢（详细），谢谢~
```
```text
 Hi  抱歉咨询一个简单的问题， 自定义词典更新后，缓存文件jieba.cache是否会更新?
```
```text
 我发现没有更新，好郁闷！不知道如何更新。
```
## 34.请问大家有好的停用词词典么
Jul 27, 2018
```text
 我在py2.7下
re_han = re.compile("([\u4E00-\u9FD5]+)")对中英文混合分词时出现问题
比如输入 中国tfboy说唱,篱笆女人等。
re_han.split结果为[u‘中国’，u'tfbo' ，u'y说唱', ....]
将re_han 改为 re.compile(u"([\u4E00-\u9FD5]+)")能正常分组
```
```text
 个人在做简历分类，请问大家，有适合的中文文本分类的停用词词典么
```
```text
 A possible one if you really need it in NLP task: https://github.com/kn45/seg-jb/blob/master/segjb/stopwords.dat
```
```text
 https://github.com/hankcs/HanLP/blob/master/data/dictionary/stopwords.txt
```
## 35.识别含未登录词的句子沿用有向无环图和最大概率路径的方法
Jul 24, 2018
```text
 用英文试了一下结巴的分词，但是出来的单词很奇怪，很多好像都被切掉了。
如图

想问下可以支持用自己的分词结果去做比如extract嘛？就比如传list而不是text
```
```text
 希望在生产环境用结巴，但生产环境对于获取绝对路径、产生cache文件等操作有沙盒控制。想请问cache文件在jieba中是否有不可替代的作用，我可否修改代码做一个不用cache的版本（略降低效率也可）？
```
```text
 刚刚接触本项目，注意到对于含未登录词的句子引用cut(..., HMM=True)方法时，会将整句作为observable放入HMM模型，用viterbi算法推测分词情况。请问这一方法对于未分词的句子的已识别的部分，是否在模型内体现了用cut(..., HMM=False)的方法？
如果能确定句子中仅包含已登录词的部分，及其分词情况，进而直接定义好HMM模型中对应的隐藏状态，需要计算概率的HMM的隐藏状态段减少，是否能提高分段预测情况？
```
## 36.词频高却没有被切出来
Jul 24, 2018
```text
 Hi I'm using Jieba with Java and I have a question.
My purpose to use Jieba is getting words from sentences and add mean per word. But sometimes the result is not what I expected.
sentence : 我要九个
result : 我要 / 九个
exptected : 我 / 要 / 九 / 个
sentence : 吃三丸
result : "吃三丸"
expected : 吃 / 三 / 丸
Because I have a Chinese English dictionary like this.
我 : I
要 : want
九 : nine
个 :  individual; measure word for people or objects
But for current result I have to have all possible cases in the dictionary.
ex）一个， 两个，三个，四个，六个，七个，八个，九个，是个。。。。
Thanks in advance.
```
```text
 代码如下：
#encoding=utf-8
import jieba
import jieba.posseg as posseg

jieba.load_userdict("/data/paper/keywords/user_dict.txt")

words = posseg.cut("在现代社会中，免洗消毒液的使用越来越普遍，消毒液中所含的肉豆蔻酸异丙酯、丙二醇和乙醇等成分可增强皮肤的渗透性。")

for w in words:
    print('%s %s' % (w.word, w.flag))

报错
Building prefix dict from /usr/lib/python2.7/site-packages/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.296606063843 seconds.
Prefix dict has been built succesfully.
Traceback (most recent call last):
File "./jieba-test.py", line 5, in 
jieba.load_userdict("/data/paper/keywords/user_dict.txt")
File "/usr/lib/python2.7/site-packages/jieba/init.py", line 119, in wrapped
return fn(_args, *_kwargs)
File "/usr/lib/python2.7/site-packages/jieba/init.py", line 312, in load_userdict
add_word(_tup)
File "/usr/lib/python2.7/site-packages/jieba/init.py", line 116, in wrapped
return fn(_args, **kwargs)
TypeError: add_word() takes at most 3 arguments (4 given)
之前还工作的好好的。user_dict.txt 类似于这样的：
酞 10 n
 酞酐 20 n
 酞磺胺噻唑 50 n
 酞磺醋胺 40 n
 酞基 20 n
 酞己炔酯 40 n
 酞腈 20 n
 酞菁 20 n
 酞菁二(吡啶)铁(II)络合物 150 n
 酞菁镉 30 n
 酞菁蓝 30 n
 酞菁锂 30 n
 酞菁镁 30 n
 酞菁染料 40 n
 酞菁素艳蓝 50 n
 酞菁锡(II) 70 n
 酞菁亚铁 40 n
 酞菁铟 30 n
 酞菁银 30 n
 酞类染料 40 n
 酞嗪 20 n
 酞醛 20 n
 酞酸 20 n
 酞酸丁苄酯 50 n
 酞酸二苯酯 50 n
 酞酸二丙酯 50 n
 酞酸二环己酯 60 n
 酞酸二甲酯 50 n
 酞酸二壬酯 50 n
 酞酸二辛酯 50 n
 酞酸二乙酯 50 n
 酞酸二异癸酯 60 n
 酞酮酸 30 n
 酞醯胺醯基 50 n
 酞醯基 30 n
 酞醯亚胺基 50 n
 酞酰胺 30 n
 酞酰氯 30 n
 酞亚胺 30 n

其中优先级是按词语长度用脚本自动生成的，给予长词更高的优先级。
```
```text
 我发现了。似乎是自定义词库本身的问题
```
```text
 以下示例中 words 是默认词典中所有的词，feqs 是对应的词频：
freqs[words.index('北京市')]  # 3392
freqs[words.index('人民政府')]  # 15227
freqs[words.index('北京市人民政府')]  # 135

list(jieba.cut('北京市人民政府发布了一条政策'))
#  ['北京市人民政府', '发布', '了', '一条', '政策']
如上，北京市 和 人民政府 的词频要比 北京市人民政府 高出很多，但是 北京市 和 人民政府 却没有被切开，这是为什么呢？
```
```text
 注意源码中calc()函数：

def calc(self, sentence, DAG, route): 
N = len(sentence)

route[N] = (0, 0) 

logtotal = log(self.total) 

for idx in xrange(N - 1, -1, -1):

    route[idx] = max((log(self.FREQ.get(sentence[idx:x + 1]) or 1) - logtotal + route[x + 1][0], x) for x in DAG[idx])


最后一行的运算，实际上是log(单词词频) - log(全词词频) = log(单词词频 / 全词词频)，也即这个词与所有词词频之比。因此词数多时原概率会多除以一个分母即全词词频，概率上反而变小。此处并非单词词频简单相加。
个人理解，欢迎交流。
```
```text
 你能以我上面给的例子解释下吗？我看文档上说的是词频高就能被切开。
多多指教。
```
```text
 假设所有词词频之和为total（比如1000000）。当进行动态规划分词时，当进行“北”字子运算时，比较的是log(“北京市”->词频 / total) + log("人民政府"->词频 / total) = log("北京市"->词频 * “人民政府"->词频 / total^2)，和log("北京市人民政府"->词频 / total)。虽然“北京市”和“人民政府”词频相对高，但由于多除了一个全词词频之和的分母，概率运算的结果实际上是变小了的。据我的理解，实际上分词方法倾向于分割为字典中匹配的尽可能长的词。
另外，此处用log函数而非乘除法直接运算似乎是为了解决浮点数对于此处概率极小时不够精确的情况。
```
```text
 多谢指教，我再看看
```
## 37.TextRank 算法疑问？
Jul 22, 2018
```text
 我在自定义词典里加入了带有罗马数字的词，如“纽约Ⅰ线”，并且为其赋予很高的词频，但是总是被分隔开。请问我要如何才能强制把“纽约Ⅰ线”这样的自定义词完整切分出来？？？
词典相关内容如下：
纽约Ⅰ线 100000 ns
python版本3.6.2
代码大致如下：
line = "并明确控制纽约Ⅰ线功率"
jieba.load_userdict("dict.txt")
print(jieba.get_FREQ("纽约Ⅰ线"))
print(jieba.get_FREQ("纽约"))
print(jieba.get_FREQ("Ⅰ"))
print(jieba.get_FREQ("线"))
print(jieba.get_FREQ("Ⅰ线"))
words = posseg.cut(line)
for word in words:
print(word)
输出结果如下：
100000
1758
None
7688
None
并/c
明确/ad
控制/updown
纽约/ns
Ⅰ/x
线/n
功率/n
```
```text
 In [39]: from jieba.analyse import ChineseAnalyzer
ImportError                               Traceback (most recent call last)
 in ()
----> 1 from jieba.analyse import ChineseAnalyzer
ImportError: cannot import name ChineseAnalyzer
请问是什么原因？
```
```text
 看看ChineseTokenizer这个类依赖的库是不是没install，比如whoosh这个库，pip install whoosh一下
```
```text
 在统计词与邻近词的连通数的时候，为什么只统计后面的词，按道理说前后的词都要统计的吧？
```
## 38.如何动态加载自定义词典
Jul 5, 2018
```text
 No description provided.
```
```text
 https://www.cnblogs.com/adienhsuan/p/5674033.html，可以参考一下
```
```text
 我也遇到这个问题了，求告知一个比较全的词性介绍。
```
```text
 找到一个词性介绍，见jieba
```
```text
 官方应该添加一份词性列表到文档中
```
```text
 同问
```
```text
 应该和这个一样
http://ltp.readthedocs.io/zh_CN/latest/appendix.html#id3
```
```text
 生成关键词，里面有一些符号如何处理，如()空格等一些符号
```
```text
 例如：全角符号，在全模式下不会出现。但是在精确模式下就会出现。
```
```text
 你是想要清掉標點符號嗎？ 像這篇？ #169
```
```text
 在一个已经在运行的进程中，我想能够自动加载自定义词典以实现实时对关键词的提取。之前我已经尝试过直接实时判断本地词典文件是否有更新，更新则使用load_userdict来加载词典，但是这个方法好像不管用，新添加进去的关键词还是不能被识别，请问大佬们有什么办法破吗？
```
```text
 CustomDictionary.reload() 可以重新加载词典，但是在加载期间， 分词功能不知道还能不能正常使用。
```
```text
 @csyjgu 敢问大侠CustomDictionary是自定义词典？还是什么？有reload方法吗？
```
```text
 我也有类似的问题，能否在一个工程中能否载入多个字典，然后根据方法需要动态选择使用哪个？
```
```text
 我也有相同的疑問，能否動態選擇要加載哪一種詞典呢?
目前想到兩種做法，

每次更換自定義字典時 使用 jieba.initialize() 初始化，再切換辭典，不過這個方法挺耗時的。
創造多個 jieba實例對象分別載入自定義辭典，再依據需求選用不同的對象。

雖然可行，但都不是最佳解！
```
```text
 CustomDictionary.reload() 可以重新加载词典，但是在加载期间， 分词功能不知道还能不能正常使用。

这是HanLP中的方法吧。
```
```text
 可以用jieba.add_word 和 jieba.del_word 进行动态修改吧, 我现在的问题是web应用程序是多进程运行的，add和del之后只对当前处理请求的进程生效 。
```
## 39.cut_all为True时反而切不出正常模式下的词
Jul 4, 2018
```text
 ：将目标文本按行分隔后，把各行文本分配到多个python进程并行分词，然后归并结果，从而获得分词速度的可观提升
基于python自带的multiprocessing模块，目前暂不支持windows
```
```text
 You can try this. It works in Windows.
from path import Path
from multiprocessing import Pool
import argparse
import time

LINE_PER_CORE = 5000
NUM_CORE = 30
FLOOR_COUNT = 10
CEIL_COUNT = 200

import jieba


def process_one(_in):
    r_list = []
    for l in _in:
        new_l = ' '.join(jieba.cut(l))
        r_list.append(new_l.strip())
    return r_list


def do(l_list, writer):
    pool = Pool(NUM_CORE)
    r_list=pool.map(process_one,[l_list[it:it+LINE_PER_CORE] for it in range(0,len(l_list),LINE_PER_CORE)])
    pool.close()
    pool.join()
    for lr in r_list:
        for line in lr:
            writer.write(line + '\n')

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("-i","--input", help="input folder", default=".")
    parser.add_argument("-o", "--output", help="output folder", default="w_process")
    parser.add_argument("--LINE_PER_CORE", help="# lines per core", type=int, default=20000)
    parser.add_argument("--NUM_CORE", help="# of cores", type=int, default=30)
    parser.add_argument("--coding", type=str, default="utf-8")

    args = parser.parse_args()
    print("Args :", args)
    input_folder = args.input
    output_folder = args.output
    LINE_PER_CORE = args.LINE_PER_CORE
    NUM_CORE = args.NUM_CORE
    coding = args.coding

    if not Path(output_folder).exists():
        Path(output_folder).mkdir()
    for f in Path(input_folder).files('*.txt'):
        print(f.basename(), time.strftime('%Y-%m-%d %X', time.localtime()))
        with open(output_folder + '/%s.output.txt' % (f.namebase,),'w', encoding='utf-8') as f_out:
            with open(f.abspath(),'r', encoding='utf-8') as f_in:
                l_list=[]
                all_dict = {}
                for l in f_in:
                    if len(l_list)<NUM_CORE*LINE_PER_CORE:
                        l_list.append(l)
                    else:
                        do(l_list, f_out)
                        print(f.basename(), time.strftime('%Y-%m-%d %X', time.localtime()))
                        l_list=[]
                if len(l_list)>0:
                    do(l_list, f_out)
    print(time.strftime('%Y-%m-%d %X', time.localtime()))
```
```text
 jieba.load_userdict('entity_dic.txt')
Traceback (most recent call last):
File "", line 1, in 
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 119, in wrapped
return fn(_args, *_kwargs)
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 334, in load_userdict
logger.debug('%s at line %s %s' % (f_name, lineno, line))
NameError: global name 'f_name' is not defined
我是用pip安装的。。。
```
```text
 应该是 #257 的问题，还是升级版本吧。
```
```text
 现在版本是0.36
```
```text
 这个问题出现在 0.36 整个版本，你可以自己改正：
jieba/__init__.py第 334 行：
logger.debug('%s at line %s %s' % (f.name, line_no, line))
```
```text
 
```
## 40.请问HMM如果分词不准确了，该怎么修改呢？
Jul 4, 2018
```text
 No description provided.
```
```text
 我是用正则表达式处理的，new_sentence = re.sub(r'[^\u4e00-\u9fa5]', ' ', old_sentence) 然后再进行分词的, \u4e00-\u9fa5这个是utf-8中，中文编码的范围
```
```text
 @cbzhuang  非常谢谢你的回复！ 我用了这个，不知道可对。#169
```
```text
 Actually, CJK characters are encoded together so there's no critical range for Chinese characters. A punctuation dict could be used to do the filtering.
```
```text
 db中有数据条“酒”，“白酒”，"红酒"
搜索“酒" ,只能出现一条记录”酒"
希望能同时出现“白酒”,"红酒"
```
```text
 这种问题不归这里管吧
```
```text
 rt，其实有时候拼音纠错还是很大作用的，不支持拼音就意味着不支持同音字，非常希望能加入这个功能
```
```text
 假如我们使用自己的语料来训练HMM，如果HMM的分词不准确，该怎么修改呢？
方法1. 我们在HMM的语料中人为的增加这个词的出现几率吗？
但是这样我们就会陷入一个循环：1.手工修改语料--->  2.训练HMM----->3.测试分词结果   ----如果不行继续回到第一步？
```
```text
 有趣的问题。
你可以想想看其他的学习方法

Brent

On Jul 4, 2018, at 17:42, 王晓珂 <notifications@github.com<mailto:notifications@github.com>> wrote:


假如我们使用自己的语料来训练HMM，如果HMM的分词不准确，该怎么修改呢？
方法1. 我们在HMM的语料中人为的增加这个词的出现几率吗？
但是这样我们就会陷入一个循环：1.手工修改语料---> 2.训练HMM----->3.测试分词结果 ----如果不行继续回到第一步？

—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub<#644>, or mute the thread<https://github.com/notifications/unsubscribe-auth/AM-fgWuc7VQrNuSN_aU5hoxSOpLVjjktks5uDI38gaJpZM4VCRlN>.
```
```text
 用感知机，CRF之类的高级算法
```
## 41.复合词，这种怎么分词处理
Jun 26, 2018
```text
 参考Issue: #387
如果直接在spark-stream的flatMap中执行jieba.cut()分词，则因“发现该dict加载后是以python字典的数据类型存在于spark的driver里面，但是worker进程无法共享这段内存。By @liaicheng ”，导致每次执行时多次“Building prefix dict from the default dictionary ……”，影响了分词性能。
如果“建议用jieba.Tokenizer 得到一个分词器实例，然后再调用相关方法。By @fxsjy ”，则对tokr实例进行广播变量（sc.broadcast(tokr)）时，报“can't pickle thread.lock objects”的错误。
考虑到spark是采用pickle进行持久化，参考pyahocorasick（https://pypi.python.org/pypi/pyahocorasick/ ）的实现方式，
You can also create an eventually large automaton ahead of time and pickle it to re-load later. Here we just pickle to a string. You would typically pickle to a file instead:



import cPickle
pickled = cPickle.dumps(A)
B = cPickle.dumps(pickled)
B.get('he')
(0, 'he')



jieba也能否优化以更好地支持序列化从而进一步提高在spark中的执行效率？
```
```text
 结巴不支持多线程分词吗？为什么多线程分词还要更慢呢？
缘起
对14206篇文章进行相似性去重 首先要分词 总用时12分钟 发现光分词就用了8分钟
所以想通过多线程分词来加快处理速度 如5个线程 每个线程处理2841篇文章
如下所示
def segment(texts_, start, end):
    print(start, end)
    texts_cut = []
    for t in texts_[start:end]:
        if not isinstance(t, str):
            t = ""
        t = regex.sub("", t)
        cut = list(jieba.cut(t))
        texts_cut.append(cut)

    return texts_cut

max_workers = 5
batch_count = len(texts)//max_workers
logging.info("start")
with concurrent.futures.ThreadPoolExecutor(max_workers=max_workers) as executor:
    future_2_start = {executor.submit(segment, texts, s*batch_count, (s+1)*batch_count if s < max_workers-1 else len(texts)) : s for s in range(max_workers)}
    start_2_data = {}
    for future in concurrent.futures.as_completed(future_2_start):
        start = future_2_start[future]
        data = future.result()
        start_2_data[start] = data

    texts_cut = []
    for start in sorted(start_2_data):
        texts_cut.extend(start_2_data[start])

    print(len(texts_cut))
    
logging.info("done")

于是用了5193篇文章做了实验
单线程（max_workers=1）用时约43秒
2018-08-01 19:10:14,529 : INFO : start
2018-08-01 19:10:57,450 : INFO : done

5个线程（max_workers=5）用时约50秒
2018-08-01 19:11:38,841 : INFO : start
2018-08-01 19:12:29,616 : INFO : done

为什么多线程分词还要更慢呢？
```
```text
 原因

The GIL is an interpreter-level lock. This lock prevents execution of multiple threads at once in the Python interpreter. Each thread that wants to run must wait for the GIL to be released by the other thread, which means your multi-threaded Python application is actually single threaded. The GIL prevents simultaneous access to Python objects by multiple threads.
摘自：https://medium.com/practo-engineering/threading-vs-multiprocessing-in-python-7b57f224eadb

解决方法：

But of course, we would want to use the ProcessPoolExecutor for CPU intensive tasks. The ThreadPoolExecutor is better suited for network operations or I/O
摘自： http://masnun.com/2016/03/29/python-a-quick-introduction-to-the-concurrent-futures-module.html

上述代码 ThreadPoolExecutor ==》 ProcessPoolExecutor 即可
注意：
先显式初始化 否则每个进程都要初始化
jieba.initialize()

用时
# 多进程 ProcessPoolExecutor 8个进程
2018-08-11 10:38:34,569 : INFO : start
2018-08-11 10:38:47,999 : INFO : done took time: 12.43

2018-08-11 10:40:23,051 : INFO : start
2018-08-11 10:40:35,909 : INFO : done took time: 11.95
```
```text
 其实jieba自带了并行分词，只需要jieba.enable_parallel()
```
```text
 待分词的内容：

敏感词|XX功
大笨蛋
XX功
敏感词
法轮功

问题：无法分出“敏感词”、“大笨蛋”、“XX功”
```
```text
 自定义词典不能解决吗
```
## 42.为什么对gen_pfdict并行化没有加速效果呢
Jun 18, 2018
```text
 无论是否翻墙，均无法访问，也ping不到
```
```text
 @Pzoom522 这个链接已经崩了很久了。
```
```text
 这个分词里面多少只有一种词性吗。
```
```text
 逻辑上 第一个多少应该分为 多 和 少
```
```text
 @sidealice , 这个难度挺高的。先mark一下。
```
```text
 这个好难哈哈
```
```text
 这种问题估计本身就有问题：比如我在你那句接一句。你说的多少是多少。只有说得人才知道哪个是多少，哪个是多+少了，甚至乎自己都不知道。这是语言表达本身就不够准确规范，让人都猜不透的，估计就不是难的问题了吧？
```
```text
 这个属于语意消除歧义，以现阶段的计算能力，是很难的。
```
```text
 <type 'exceptions.IOError'> [Errno 5] Input/output error   for term in terms:\n', '  File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 255, in cut\n    for word in cut_block(blk):\n', '  File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 195, in cut_DAG\n    DAG = get_DAG(sentence)\n', '  File "/usr/local/lib/python2.7/dist-packages/jieba/__init.py", line 115, in wrapped\n    initialize(DICTIONARY)\n', '  File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 65, in initialize\n    print >> sys.stderr, "Building Trie..., from " + abs_path\n'
难道是由于后台运行后，console退出了，导致print出差？
```
```text
 在遍历文件每一行那里，使用threading多线程处理，无论是静态分配还是动态分配（不同块大小），都没有直接串行的效率高，词典规模增大三倍后仍是这样。
这是为什么呢？
```
```text
 python 有一个全局多线程锁，对于计算密集型任务应该使用多进程。需要加速可以试试这个repo https://github.com/deepcs233/jieba_fast/tree/master/jieba_fast。
```
## 43.可以处理文言文吗？
Jun 17, 2018
```text
 在使用pseg.cut来获取分词词性的时候，发现“云计算”、“石墨烯”这些词语的词性居然是x，x我看介绍通常是一些标点符号，为什么会出现这种情况，还有其他词语也会有这种情况吗？
```
```text
 因为云计算，石墨烯是新词，词典里找不到，估计就默认设置成词性x了，你需要自己自定义词典并且加入这些新词的词性。
```
```text
 在使用用户词典时，运行样例，若是有词包含空格，例如样例中的
Edu Trust认证 2000
则运行时会出现如下的错误：
Prefix dict has been built succesfully.
Traceback (most recent call last):
File "test.py", line 6, in 
jieba.load_userdict("userdict.txt")
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 381, in load_userdict
f.name, lineno, line))
ValueError
```
```text
 pypi 上的版本没更新
```
```text
 试了一下，赶脚效果不太好
```
```text
 字典要换，用这个 https://github.com/The-Orizon/nlputils/tree/master/jiebazhc
```
## 44.在Windows和linux下，jieba.posseg用于词性提取英语，词性不一样
Jun 14, 2018
```text
 請問我們想分詞出來的結果是兩個英文組合字, 在自建詞庫裡有"Apple book", 但分詞出來後還是"Apple"和 "book"分開的兩個字, 請問要如何處理才能達到我們要的結果? 謝謝
```
```text
 我們希望看到的是"Apple book", 兩個字組合成的詞, 請各位大俠開導一下,謝謝 @fxsjy
```
```text
 求大神進來看一下唄... @felixonmars  , 好像都沒人進來看???奇怪???
```
```text
 請問這個討論群是不是關了?都沒有人回覆耶????
```
```text
 这个问题得从底层改（最近没人想来填坑的样子
```
```text
 能透過改這個檔案來達成目的嗎?  init.py
```
```text
 可以
```
```text
 @gumblex, @fxsjy  , 請教大俠如何達成? 如何改? 拜託指導一下, 大大的感謝
```
```text
 您好，如何添加停用词，常见的停词表在哪里可以下载呢
```
```text
 停用词表推荐这个：https://github.com/dongxiexidian/Chinese/blob/master/stopwords.dat
```
```text
 多谢！
```
```text
 @tukeyone  , here is the example for adding the stop words
# build stop word list
def stopwordslist(filepath):
    stopwords = [line.strip() for line in open(filepath, 'r').readlines()]
    return stopwords

def seg_sentence(sentence):
    sentence_seged = jieba.cut(sentence.strip())
    # loading stop word list
    stopwords = stopwordslist("your_stop_word_txt_file")  
    outstr = ''
    for word in sentence_seged:
        if word not in stopwords:
            if word != '\t':
                outstr += word
                outstr += " "
    return outstr
```
```text
 Thank you ！ this problem has been solved .
```
```text
 无自定义词典
str="QFII和RQFII迎政策利好 外资更敢来玩了？"
在Windows下，seglist=jieba.posseg.cut(str)得到 QFII eng，RQFII eng
在linux下，seglist=jieba.posseg.cut(str)得到 FII x，FII x
```
```text
 还是用户词典导致了 分词出错了
```
## 45.通过文件名的形式加载自定义用户词典时，若省略掉词频和词性，每一个单词后面必须要加一个空格，不然会报错
Jun 7, 2018
```text
 代码如下：
string="这是一个test的行。和一些无意义的dm、jksajdfl"
jieba.set_dictionary('dict.txt.big')#自定义词库
res=jieba.lcut(string,HMM=False)
print(res)
['test', 'dm', 'jksajdfl']
自定义的dict里面只有一个测试的汉字词。
```
```text
 自己解决了。
init.py 破坏掉re_eng = re.compile('[a-zA-Z0-9]', re.U) 就可以了。。。。
比如re_eng = re.compile('aaaaaaaa', re.U)
```
```text
 清华最近放出来的THULAC看起来很赞的样子，文档中提到了训练模型所用的语料库，官方团队的资源获取能力实在是太厉害了。不知道可否和清华方面沟通一下，用他们的语料库来训练jieba的模型。
THULAC

1.THULAC工具包提供的模型是如何得到的？
THULAC工具包随包附带的分词模型Model_1以及分词和词性标注模型Model_2是由人民日报语料库训练得到的。这份语料库中包含已标注的字数约为一千二百万字。
同时，我们还提供更复杂、完善和精确的分词和词性标注联合模型Model_3和分词词表。该模型是由多语料联合训练训练得到（语料包括来自多文体的标注文本和人民日报标注文本等）。这份语料包含已标注的字数约为五千八百万字。由于模型较大，如有机构或个人需要，请填写“doc/资源申请表.doc”，并发送至 thunlp@gmail.com ，通过审核后我们会将相关资源发送给联系人。
```
```text
 😄 说实话并没有觉得很赞……
```
```text
 不知道是不是bug

得出结果，通过pip 安装的包为旧版0.39  请作者更新一下
```
```text
 
```
## 46.词性标注不正确
Jun 6, 2018
```text
 Does anyone have a list of the POS tags, and what they stand for?  They're supposed to be compatible with ictclas.  Thanks.
```
```text
 使用：
jieba.add_word("5-HT功能", freq=True, tag='n')
jieba.add_word("食欲减退", freq=True, tag='n')
jieba.add_word("昼夜节律紊乱", freq=True, tag='n')
然后分词处理：
print("/".join(jieba.cut(sta, HMM=True)))
发现结果为：
研究/发现/，/5/-/HT/功能/活动/降低/与/抑郁症/患者/的/抑郁/心境/、/食欲减退/、/失眠/、/昼夜节律紊乱/、/内分泌/功能/紊乱/、/性功能
障碍/、/焦虑不安/、/不能/应付/应激/、/活动/减少/等/密切相关/；/而/5/-/HT/功能/增高/与/躁狂症/的/发病/有关/。
预计结果是：
研究/发现/，/5-HT功能/活动/降低/与/抑郁症/患者/的/抑郁/心境/、/食欲减退/、/失眠/、/昼夜节律紊乱/、/内分泌/功能/紊乱/、/性功能
障碍/、/焦虑不安/、/不能/应付/应激/、/活动/减少/等/密切相关/；/而/5-HT功能/增高/与/躁狂症/的/发病/有关/。
“5-HT功能”，为什么没有用？
```
```text
 因为有“-”，所以对两边的数据进行了切分
```
```text
 那怎么解决这个问题。因为有些专业名里面包含这种“-”



liwy@futongdf.com.cn

发件人： LuoZe
发送时间： 2018-09-06 21:22
收件人： fxsjy/jieba
抄送： lwy1111111; Author
主题： Re: [fxsjy/jieba] 研究发现，5-HT功能活动降低与抑郁症患者的抑郁心境、食欲减退、失眠、昼夜节律紊乱、内分泌功能紊乱、性功能障碍、焦虑不安、不能应付应激、活动减少等密切相关；而5-HT功能增高与躁狂症的发病有关。 (#666)
因为有“-”，所以对两边的数据进行了切分
—
You are receiving this because you authored the thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 from pyhanlp import *

CustomDictionary.insert('5-HT功能', '医疗术语 10')
print(HanLP.segment('研究发现，5-HT功能活动降低与抑郁症患者的抑郁心境'))
[研究/vn, 发现/v, ，/w, 5-HT功能/医疗术语, 活动/vn, 降低/v, 与/cc, 抑郁症/nhd, 患者/n, 的/ude1, 抑郁/a, 心境/n]
```
```text
 谢谢，但是我问的是结巴的解决方法！
```
```text
 谢谢，但是我问的是结巴的解决方法！

之前有帖子写了，特殊字符，可以考虑修改
jieba.re_han_default = re.compile(r'([\u4e00-\u9fa5a-zA-Z0-9+#&._%/-]+)', re.UNICODE)
```
```text
 我想要用jieba对content一列进行分词：

结果总是出现“AttributeError: 'float' object has no attribute 'decode'，为了解决这个问题，把这个series变成str之后又无法loop，请问我应该怎么处理呢？谢谢！
```
```text
 我如果猜测没错，你想写的代码应该是
t_list = []
t_list += list(jieba.cut(sample['content']))

如果不是你应该把你具体哪句出错列出来，不过很可能是python 语法不熟造成的
```
```text
 您好，谢谢您的热心回答！第一次在github发问题不是很清楚问法，给您造成麻烦啦～

问题我已经解决啦，是我数据本身的问题，让我加入一个conditional loop就OK啦～谢谢您！

祝好～

2018-04-19 15:47 GMT+08:00 MoreFreeze <notifications@github.com>:
…
 我如果猜测没错，你想写的代码应该是

 t_list = []
 t_list += list(jieba.cut(sample['content']))

 如果不是你应该把你具体哪句出错列出来，不过很可能是python 语法不熟造成的

 —
 You are receiving this because you authored the thread.
 Reply to this email directly, view it on GitHub
 <#617 (comment)>, or mute
 the thread
 <https://github.com/notifications/unsubscribe-auth/AhgNuVNwsxwnvtduCq8KfQGDIgnVnRlhks5tqEEhgaJpZM4TPapa>
 .
```
```text
 为什么"超市"被标记成动词了(v)
```
```text
 HanLP可以标对
```
## 47.’查询卡’如何才能切分成‘查询’和‘卡’呢？
May 29, 2018
```text
 No description provided.
```
```text
 在基于TF-IDF进行特征提取时，因为文本背景是某一具体行业，不适合使用通用的IDF语料库，我觉得应该使用自定义的基于该行业背景的IDF语料库。请问如何生成自定义IDF语料库呢？
我现在有的数据是几十万个该行业的文档，初步想法是：对每个文档分词去重，把所有文档分词结果汇集去重后形成一个分词集，然后对于分词集里的每一个词语w，按**f=log(该行业文档总数/(含有w的文档数量+1))**公式求出词语w的IDF值f，最后txt文件里每一行放一个(w, f)
是这样吗？谢啦~
```
```text
 还有2个问题：假设通用IDF语料库里有A B C三个词语及其idf值，我自定义IDF语料库里有A B D及其idf值，那么请问，在添加自定义的IDF语料库后：

自定义IDF语料库里的A和B及其相应idf值就直接覆盖通用IDF语料库里的A和B吧？
通用IDF语料库里原先的C及其idf值，现在还有吗？

（其实问题只有就1个：添加自定义IDF语料库后，是整个文件替换，还是说只有那些重复的词语才被替换？）
```
```text
 求助求助求助，没有朋友知道吗？？？
```
```text
 我也想做一个词库,满足自己的需求,自带的字库里有很多类似一一二/一一分/一三六八之类意义不大的词
```
```text
 line 是单个文档
all_dict = {}
for line in lines:
    temp_dict = {}
    total += 1
    cut_line = jieba.cut(line, cut_all=False)
    for word in cut_line:
        temp_dict[word] = 1
    for key in temp_dict:
        num = all_dict.get(key, 0)
        all_dict[key] = num + 1
for key in all_dict:
    w = key.encode('utf-8')
    p = '%.10f' % (math.log10(total/(all_dict[key] + 1)))
```
```text
 @M2shad0w 非常感谢！还有一个问题：
假设通用IDF语料库里有A B C三个词语及其idf值，我自定义的IDF语料库里有A B D及其idf值，那么请问，在添加自定义的IDF语料库后：

自定义IDF语料库里的A和B及其相应idf值就直接覆盖通用IDF语料库里的A和B吧？
通用IDF语料库里原先的C及其idf值，现在还有吗？
```
```text
 @siberiawolf61
我看了一下 结巴库中 load idf path 的代码
https://github.com/fxsjy/jieba/blob/master/jieba/analyse/tfidf.py#L65
class TFIDF(KeywordExtractor):

    def __init__(self, idf_path=None):
        self.tokenizer = jieba.dt
        self.postokenizer = jieba.posseg.dt
        self.stop_words = self.STOP_WORDS.copy()
        self.idf_loader = IDFLoader(idf_path or DEFAULT_IDF)
        self.idf_freq, self.median_idf = self.idf_loader.get_idf()

...

self.idf_loader = IDFLoader(idf_path or DEFAULT_IDF)

应该是覆盖了，c值的 idf 也没有了
```
```text
 @M2shad0w 好的，谢谢啊！
```
```text
 感谢！
```
```text
 运行test目录下的demo.py时出现了
1. 分词
Building prefix dict from /usr/local/lib/python2.7/dist-packages/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.000440835952759 seconds.
Prefix dict has been built succesfully.
Full Mode: 我/ 来/ 到/ 北/ 京/ 清/ 华/ 大/ 学
Traceback (most recent call last):
File "demo.py", line 22, in 
print("Default Mode: " + "/ ".join(seg_list))  # 默认模式
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 274, in cut
for word in cut_block(blk):
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 199, in cut_DAG
calc(sentence, DAG, route=route)
File "/usr/local/lib/python2.7/dist-packages/jieba/__init.py", line 144, in calc
logtotal = log(total)
ValueError: math domain error
这是因为什么？
这是用python2还是python3写的？
```
```text
 词典好像没成功加载，删掉 /tmp/jieba.cache，下载最新版试试。
最新版 Python 2/3 均兼容
```
```text
 已经手动把结巴自带的词典dict.txt中‘查询卡’这个词删掉了，并且结巴自带词典含有‘查询’和‘卡’两个词，为什么使用lcut命令还是切分成’查询卡’呢？只有del_word('查询卡’)后才能被正确切分成‘查询’和‘卡’？请问导致出现这样问题的原理是什么，不用del_word的话怎么解决这个问题
```
## 48.請問是否能更詳細解釋DAG的建立機制
May 28, 2018
```text
 Command "/usr/bin/python2 -u -c "import setuptools, tokenize;file='/tmp/pip-build-6jExFt/jieba/setup.py';f=getattr(tokenize, 'open', open)(file);code=f.read().replace('\r\n', '\n');f.close();exec(compile(code, file, 'exec'))" install --record /tmp/pip-xAfrwY-record/install-record.txt --single-version-externally-managed --compile" failed with error code -9 in /tmp/pip-build-6jExFt/jieba/
```
```text
 Can you provide full error message?
```
```text
 请问dict.txt 词库是否都来自搜狗的2006版免费词库? #24
```
```text
 如果词典中有个关键词是ck，英文单词lock会被切分为lo、ck。
```
```text
 数词+量词+单字名词 很多都有这个问题……
```
```text
 jieba是默认把字母与数字同汉字分开的
```
```text
 「基于前缀词典实现高效的词图扫描，生成句子中汉字所有可能成词情况所构成的有向无环图 (DAG)」
從辭庫到Trie乃至於DAG的形成過程有所不解
若我的目標為「男朋友」，那麼會是 {0: [0, 1, 2]}嗎?
那這樣應該會有「男」、「男朋」、「男朋友」3種可能不是嗎?
但為何只會出現「男」、「男朋友」2種?
還是其實是{0: [0, 2]} ?
謝謝
```
## 49.AttributeError: module 'jieba' has no attribute 'lcut'
May 18, 2018
```text
 jieba.enable_parallel()
jieba.load_userdict("./aux/dict")
如上所示，如果先开启并行，再载入自定义词典，会导致自定义的词典没有产生效果
如果先载入自定义词典再进行并行计算开启，自定义词典就是有效的，不知道这算不算bug
```
```text
 确实如此。
先开启并行，再载入自定义词典。此时，从jieba.dt.get_FREQ和jieba.dt.user_word_tag_tab中都能查到用户词典中的词，说明已导入用户词典中的词，但分词结果中并未产生效果。
```
```text
 这是bug吧？同遇到了，定位了半天问题，求修复啊
```
```text
 No description provided.
```
```text
 keywords = jieba.analyse.textrank(content, topK=20, withWeight=True)
for keyword in keywords:
print keyword[0], keyword[1]
和TF-idf一样的方法，这样就可以了。
```
```text
 python3上面使用这种方法将会得到分开的字，不是score.
```
```text
 比如动词 名词之类的
```
```text
 https://github.com/fxsjy/jieba#4-词性标注
```
```text
 词性标注的缩写有没有一个对照表，虽然大部分能猜出来
```
```text
 问题1：
在jieba首页说明中显示自定义词典词频和词性不必须，但是我使用textrank，更改词典会报错！
请问，二者是必须的吗？
问题2：
词频和词性对于textrank算法有何影响？
```
```text
 执行python脚本时报错，AttributeError: module 'jieba' has no attribute 'lcut'，代码如下：
#对句子经行分词，并去掉换行符
def tokenizer(text):
''' Simple Parser converting each document to lower-case, then
removing the breaks for new lines and finally splitting on the
whitespace
'''
text = [jieba.lcut(document.replace('\n', '')) for document in text]
return text
```
## 50.项目不维护了吗
May 17, 2018
```text
 比如"中断" "满格" 等会莫名奇妙的识别为ns,ns词性是地名词性，这也太假了
```
```text
 No description provided.
```
```text
 我在维护一个加速版的jieba，有issue和pr欢迎
https://github.com/deepcs233/jieba_fast
```
## 51.字典中的“出生”无法识别的bug
May 10, 2018
```text
 通过 jieba.set_dictionary("."+os.sep+"dict.txt") 指定dict.txt 文件位于一个带中文的父目录时，使用jieba.initialize() 进行初始化报错，报错为：
Traceback (most recent call last):
File "", line 26, in 
File "F:\Qiyi\build\getcalldatafrom_txt\out00-PYZ.pyz\jieba", line 111, in initialize
UnicodeDecodeError: 'ascii' codec can't decode byte 0xcc in position 3: ordinal not in range(128)
定位于语句： jieba_init.py_ 中的 111行： default_logger.debug("Building prefix dict from %s ..." % (abs_path or 'the default dictionary'))
将程序至于全英文目录时，运行正常。
```
```text
 我在使用jieba分词时，遇到了如下情况：
待分词句子：奥布瑞·普拉扎（Aubrey Plaza），1984年6月26日出生于美国特拉华州威尔明顿，演员。
分词结果：
jieba.cut:
奥/布/瑞/·/普拉/扎/（/Aubrey/ /Plaza/）/，/1984/年/6/月/26/日出/生于/美国/特拉华州/威尔/明/顿/，/演员
posseg.cut:
奥 nr 布 nr 瑞 ns · x 普拉 nrt 扎 v （ x Aubrey eng   x Plaza eng ） x ， x 1984 eng 年 m 6 eng 月 m 26 eng 日出 v 生于 v 美国 ns 特拉华州 ns 威尔 nrt 明 a 顿 q ， x 演员 n
出生这个词一直无法正确分出来，我发现词典中已经包含了该词：出生 1979 v，而日出频率低于出生：日出 394 v；然后我尝试自己再一次将“出生”添加到词典，以及关闭HMM，均没有作用，请问这是什么问题呢，谢谢！
```
```text
 找到了原因，日出的频率是394，出生频率是1979，生于的频率是4690，导致分为日出/生于
解决办法：
jieba.del_word('日出')
jieba.add_word('出生于')
jieba.add_word('日出')
```
## 52.带 “-”的英文自定义单词问题
May 2, 2018
```text
 现在我cut
我想看电视
结果是
我 想 看电视
看了下默认的词库里有 看电视 这个词。。
我想搞成
我 想 看 电视
```
```text
 jieba.add_word("看电视",0)#或者jieba.del_word("看电视")#发现主干这个del_word的bug已经修了
jieba.add_word("看",100,"v")
jieba.add_word("电视",100,"n")
```
```text
 我自定义了一个词典，有一些包含“-”的词，比如 “SK-II”。加载这个词库时候发现只有这类词语没法生效，会被分为 SK，-，II,而不是SK-II。请问这个问题应该如何解决
```
```text
 同问，自定义词典中有5%,但分词会分开，词性标注的会分开 但正常分词不会被分开
```
```text
 找到解决方法了
修改jieba/init.py代码中的re_han_default 正则表达式为如下值：
re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._%\-]+)", re.U)
```
```text
 Can solve this: #692
```
## 53.停词如何添加
Apr 26, 2018
```text
 請問我們想分詞出來的結果是兩個英文組合字, 在自建詞庫裡有"Apple book", 但分詞出來後還是"Apple"和 "book"分開的兩個字, 請問要如何處理才能達到我們要的結果? 謝謝
```
```text
 我們希望看到的是"Apple book", 兩個字組合成的詞, 請各位大俠開導一下,謝謝 @fxsjy
```
```text
 求大神進來看一下唄... @felixonmars  , 好像都沒人進來看???奇怪???
```
```text
 請問這個討論群是不是關了?都沒有人回覆耶????
```
```text
 这个问题得从底层改（最近没人想来填坑的样子
```
```text
 能透過改這個檔案來達成目的嗎?  init.py
```
```text
 可以
```
```text
 @gumblex, @fxsjy  , 請教大俠如何達成? 如何改? 拜託指導一下, 大大的感謝
```
```text
 您好，如何添加停用词，常见的停词表在哪里可以下载呢
```
```text
 停用词表推荐这个：https://github.com/dongxiexidian/Chinese/blob/master/stopwords.dat
```
```text
 多谢！
```
```text
 @tukeyone  , here is the example for adding the stop words
# build stop word list
def stopwordslist(filepath):
    stopwords = [line.strip() for line in open(filepath, 'r').readlines()]
    return stopwords

def seg_sentence(sentence):
    sentence_seged = jieba.cut(sentence.strip())
    # loading stop word list
    stopwords = stopwordslist("your_stop_word_txt_file")  
    outstr = ''
    for word in sentence_seged:
        if word not in stopwords:
            if word != '\t':
                outstr += word
                outstr += " "
    return outstr
```
```text
 Thank you ！ this problem has been solved .
```
## 54.第三方库与jieba的兼容问题
Apr 23, 2018
```text
 如題，TFIDF算法一般來說topK這個參數要取多少比較合適?我目前使用的題材是金庸的小說，目前正在切神鵰俠侶這本。
```
```text
 遇到了一个奇怪的情况，我在Ubuntu16.04下通过pip安装了jieba0.39，可以正常使用，但后来可能又装了一些第三方的库，然后jieba.lcut就报错了，显示没有这个方法，由于后来装的第三方库比较多，正在生产环境上使用，不能直接卸载，感觉这个问题是因为某个第三方库中引入了jieba旧版本的代码，以至于我现在只要使用jieba，用的旧的代码，但通过pip list查询，jieba的版本还是0.39，请问应该怎样解决？谢谢~
```
```text
 我也遇到这个问题了，我是因为同时装了jieba3k，换成jieba就可以了，jieba.lcut只有在jieba中可以使用
```
## 55.请问如何添加表情做分割
Apr 21, 2018
```text
 No description provided.
```
```text
 你没看代码吗？
```
```text
 哎，作者就是不爱写注释，哎。
```
```text
 @ilqxejraha 谢谢您，可能是我表达有误，我是想问词汇库里面的词汇和词频是人工统计的吗？还是通过其它的方法。
```
```text
 统计过来的。最后得到了就是这么一个模型。
```
```text
 你在源码中看到词频的使用了吗？
```
```text
 具体的统计词频，词频的作用可能体现在，一个词存在多个意思。
比如英语中，经常有一个词会有很多个意思。
比如出现一个词，并且，这个词有很多种解释，这时候词频可能会对词意的选择有一定帮助。
具体的其他算法hmm的我还没看。
```
```text
 @KevinDotW 据说是基于人民日报的语料库，我也想知道怎么才能创建自己的词典
```
```text
 词性标注是否可以指定分词模式；
想指定搜索模式
谢谢。
```
```text
 比如微博的表情文字是   [泪]
我想拿到这个表情统计个数，请问可以怎么做？
```
## 56.extract_tags中分词不合理
Apr 17, 2018
```text
 如何验证分词的准确性
请问用jieba分完词后，如何去验证分词的准确性？
```
```text
 Same question. What is the training data? 人民日报语料吗？
```
```text
 Same question. What is the training data? 人民日报语料吗？
不是人民日报的，是网上爬取的一些评论数据，最近看了数学之美的分词，有一些启发
```
```text
 一些不字开头的词语在短句中，为啥被extract_tags后"不"字就不见了，将"不"与后面词语分开，这样提取的词的意思都变了，不合理吧
比如

不合适
不适合
不太合适
不太适合
不满意
不想买
不想要
```
## 57.source of the word dictionary
Apr 16, 2018
```text
 需求：
要求匹配出内容的指定的一些关键词（我自己的词库里的词），忽略jieba里千千万万不相关的词，对于我来说，这些词没有用。
实现：
我想的是，把jieba默认的词库替换掉。可是，我看文档里没有提到如何替换jieba的词库
还有什么好的办法实现吗？
谢谢各位
```
```text
 你可以用jieba的自定义字典实现准确识别

中铁建资产管理有限公司 高正  18601064889
…
 在 2018年10月24日，11:11，Bakkan Hwang ***@***.***> 写道：

 需求：
 要求匹配出内容的指定的一些关键词（我自己的词库里的词），忽略jieba里千千万万不相关的词，对于我来说，这些词没有用。

 实现：
 我想的是，把jieba默认的词库替换掉。可是，我看文档里没有提到如何替换jieba的词库
 还有什么好的办法实现吗？

 谢谢各位

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 jieba有实现延时加载机制
jieba 采用延迟加载，import jieba 和 jieba.Tokenizer() 不会立即触发词典的加载，一旦有必要才开始加载词典构建前缀字典。如果你想手工初始 jieba，也可以手动初始化。
import jieba
jieba.initialize()  # 手动初始化（可选）
在 0.28 之前的版本是不能指定主词典的路径的，有了延迟加载机制后，你可以改变主词典的路径:
jieba.set_dictionary('data/dict.txt.big')
例子： https://github.com/fxsjy/jieba/blob/master/test/test_change_dictpath.py
```
```text
 In [4]: print '/'.join(jieba.cut("我来到清华大学",cut_all=True))
我/来到/清华/清华大学/华大/大学
In [5]: jieba.add_word('到清')
In [6]: print '/'.join(jieba.cut("我来到清华大学",cut_all=True))
我/来到/到清/清华/清华大学/华大/大学
In [7]: jieba.add_word('学')
In [8]: print '/'.join(jieba.cut("我来到清华大学",cut_all=True))
我/来到/到清/清华/清华大学/华大/大学
希望高手能给予帮助!
```
```text
 同问
```
```text
 #408
```
```text
 Hello, does anybody know from where the dictionary of traditional Chinese characters originated? Who created that dictionary or from where was it ported?
```
## 58.想问下jieba.load_userdict(userdic)是否会覆盖
Apr 14, 2018
```text
 孙中山 孙中山 nr
毛泽东 毛泽东 nr
周恩来 周恩来 t
邓小平 邓小平 nr
"周恩来"被标注成时间词
```
```text
 关于jieba.load_userdict(dic)这个加载自定义词典，有如下疑问，烦请解答：

加载了自定义词典会不会覆盖默认的词典
2.如果我有多个txt的自定义词典，可否加载多个，机制是加载的多个并存，还是说后加载的覆盖先加载的？
3.如问题2，如果是后加载的覆盖先加载的，那么我就需要把多个txt先合并再加载，如果是并存的那么我可以运行多次jieba.load_userdict(dic)来加载多个词典，反正会并存。
4.之所以有如上的问题，是因为我从搜狗词库里面下载了很多日常用语secl词典

非常感谢，开发人员能解答下。
```
```text
 從這段code來看，他是去呼叫add_word() 方法把字典裡的字一個一個載入，所以應該不存在覆蓋問題，也應該可以多次載入
https://github.com/fxsjy/jieba/blob/master/jieba/__init__.py#L392
```
```text
 load_userdict最后也是调用的是add_word()方法，不会覆盖
```
## 59.结巴分词中 自定义词典里面的新词的tfidf是怎么算的呢？
Apr 13, 2018
```text
 jieba.posseg.cut是可以分割词并得到tag类型，但是却把百分比分割成两个部分，tag都是'x'，jieba.cut就能准确分成"2%"，如果需要tag还需要百分比，应该怎么写
```
```text
 同样遇到了这个问题
```
```text
 同问
```
```text
 已解决，见https://blog.csdn.net/liuhongyue/article/details/80498195
```
```text
 似乎没有添加自己的停用词可能，是不是可以考虑添加
```
```text
 结巴分词中 自定义词典里面的新词的tfidf是怎么算的呢？
```
```text
 自问自答吧 ， 没有出现的词，idf是 所有词idf的平均是11左右。 tf就是这个句子中这个新词的词频。
详细分析：http://www.cnblogs.com/zle1992/p/8822832.html
```
```text
 自问自答吧 ， 没有出现的词，idf是 所有词idf的平均是11左右。 tf就是这个句子中这个新词的词频。
详细分析：http://www.cnblogs.com/zle1992/p/8822832.html

如果是平均值，为什么不是标准平均值，而是有差异呢？
```
## 60.为jieba.posseg增加cut_all选项允许全模式下词性标注
Apr 13, 2018
```text
 No description provided.
```
```text
 需求非常大
```
```text
 自己搞个thrift之类的service... 转一下... = =
```
```text
 ls +1
```
```text
 +1
```
```text
 虽然有做广告的嫌疑，但是还是发一下把：
jieba的c++版本有提供http接口供php使用啊：https://github.com/aszxqw/cppjieba
```
```text
 @raykwok @tonicbupt @tonyseek @binaryoung @yanyiwu @fxsjy 我這個週末心血來潮翻譯了 jieba 的 PHP 版本，fukuball/jieba-php，目前是從 python 的 0.16 版翻譯過來的，未來再慢慢往上升級，效能也需要再改善，請有興趣的開發者一起加入開發！
```
```text
 @fukuball 非常感谢
```
```text
 建议为jieba.posseg增加cut_all选项，允许采用全模式分词状态下再进行词性标注以得到更全的结果
```
```text
 我也需要啊，现在还没有这个api？
```
## 61.jieba无法正确识别人名啊！
Apr 11, 2018
```text
 比如将XXXX年、XX月视作一个词
或者我需要将书名号内所有内容视作整个词
能不能进行自定义配置？
```
```text
 用字典不就行了。匹配正则，想下这效率得有多慢。
```
```text
 Environment:

jieba v0.39

Code:
import jieba
jieba.add_word("现代汉语文本切分与词性标注规范Ｖ1.0");
seg_list = jieba.cut("北大计算语言学研究所从 1992 年开始进行汉语语料库的多级加工研究。第一步是对原\n" +
                "始语料进行切分和词性标注。1994 年制订了《现代汉语文本切分与词性标注规范Ｖ1.0》。")
print(','.join(seg_list))

output:
北大,计算,语言学,研究所,从, ,1992, ,年,开始,进行,汉语,语料库,的,多级,加工,研究,。,第一步,是,对,原,
,始,语料,进行,切分,和,词性,标注,。,1994, ,年,制订,了,《,现代汉语,文本,切分,与,词性,标注,规范,Ｖ,1.0,》,。
```
```text
 根据个人需要修改jieba包init.py中几个正则表达式，使其支持数字特殊字符。re_han_default = re.compile(“(.+)”, re.U)
```
```text
 from pyhanlp import *

segment = HanLP.newSegment('感知机')
CustomDictionary.insert('现代汉语文本切分与词性标注规范Ｖ1.0')
print(segment.analyze("北大计算语言学研究所从1992年开始进行汉语语料库的多级加工研究。第一步是对原" +
                      "始语料进行切分和词性标注。1994年制订了《现代汉语文本切分与词性标注规范Ｖ1.0》。"))

[北大/j 计算/v 语言学/n 研究所/n]/nt 从/p 1992年/t 开始/v 进行/v 汉语/nz 语料库/n 的/u 多/m 级/q 加工/v 研究/v 。/w 第一/m 步/q 是/v 对/p 原始/a 语/Ng 料/v 进行/v 切分/vn 和/c 词性/n 标注/v 。/w 1994年/t 制订/v 了/u 《/w 现代汉语文本切分与词性标注规范Ｖ1.0/nz 》/w 。/w
```
```text
 nr===========
['耿直', '易怒', '易怒', '鲁莽', '安静', '耿直', '公正', '公允', '明朗', '妙语连珠', '言相告', '赤诚', '赤诚', '金石为开', '子虚', '阳奉阴违', '恩尽义', '向壁虚造', '向壁虚构', '空中楼阁', '阳奉阴违', '殷勤', '和蔼', '和善', '和悦', '过谦', '周全', '和蔼可亲', '关怀备至', '殷勤', '大智若愚', '一谦四益', '纳谏如流', '博采', '刚毅', '麻利', '英勇', '英勇', '耿直', '雷厉风飞', '和蔼', '和悦', '仁慈', '慈善', '慈祥', '和善', '温顺', '温驯', '温雅', '慈爱', '蔼然可亲', '和蔼可亲', '温文', '慈悲为怀', '仁慈', '仁慈', '温柔敦厚', '束', '安静', '安静', '安闲', '安逸', '安祥', '安然', '安神', '静谧', '静默', '若素', '和蔼可亲']
这些词的flag全部都是nr，什么情况？难道只要是中文姓氏开头的词都是人名吗？那准确率太低了吧。。。
```
```text
 查了几个，在dict.txt都是人名。可能是用其他词性工具生成不够准确吧。
```
## 62.AttributeError when use jieba: 'float' object has no attribute 'decode'
Apr 11, 2018
```text
 Does anyone have a list of the POS tags, and what they stand for?  They're supposed to be compatible with ictclas.  Thanks.
```
```text
 使用：
jieba.add_word("5-HT功能", freq=True, tag='n')
jieba.add_word("食欲减退", freq=True, tag='n')
jieba.add_word("昼夜节律紊乱", freq=True, tag='n')
然后分词处理：
print("/".join(jieba.cut(sta, HMM=True)))
发现结果为：
研究/发现/，/5/-/HT/功能/活动/降低/与/抑郁症/患者/的/抑郁/心境/、/食欲减退/、/失眠/、/昼夜节律紊乱/、/内分泌/功能/紊乱/、/性功能
障碍/、/焦虑不安/、/不能/应付/应激/、/活动/减少/等/密切相关/；/而/5/-/HT/功能/增高/与/躁狂症/的/发病/有关/。
预计结果是：
研究/发现/，/5-HT功能/活动/降低/与/抑郁症/患者/的/抑郁/心境/、/食欲减退/、/失眠/、/昼夜节律紊乱/、/内分泌/功能/紊乱/、/性功能
障碍/、/焦虑不安/、/不能/应付/应激/、/活动/减少/等/密切相关/；/而/5-HT功能/增高/与/躁狂症/的/发病/有关/。
“5-HT功能”，为什么没有用？
```
```text
 因为有“-”，所以对两边的数据进行了切分
```
```text
 那怎么解决这个问题。因为有些专业名里面包含这种“-”



liwy@futongdf.com.cn

发件人： LuoZe
发送时间： 2018-09-06 21:22
收件人： fxsjy/jieba
抄送： lwy1111111; Author
主题： Re: [fxsjy/jieba] 研究发现，5-HT功能活动降低与抑郁症患者的抑郁心境、食欲减退、失眠、昼夜节律紊乱、内分泌功能紊乱、性功能障碍、焦虑不安、不能应付应激、活动减少等密切相关；而5-HT功能增高与躁狂症的发病有关。 (#666)
因为有“-”，所以对两边的数据进行了切分
—
You are receiving this because you authored the thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 from pyhanlp import *

CustomDictionary.insert('5-HT功能', '医疗术语 10')
print(HanLP.segment('研究发现，5-HT功能活动降低与抑郁症患者的抑郁心境'))
[研究/vn, 发现/v, ，/w, 5-HT功能/医疗术语, 活动/vn, 降低/v, 与/cc, 抑郁症/nhd, 患者/n, 的/ude1, 抑郁/a, 心境/n]
```
```text
 谢谢，但是我问的是结巴的解决方法！
```
```text
 谢谢，但是我问的是结巴的解决方法！

之前有帖子写了，特殊字符，可以考虑修改
jieba.re_han_default = re.compile(r'([\u4e00-\u9fa5a-zA-Z0-9+#&._%/-]+)', re.UNICODE)
```
```text
 我想要用jieba对content一列进行分词：

结果总是出现“AttributeError: 'float' object has no attribute 'decode'，为了解决这个问题，把这个series变成str之后又无法loop，请问我应该怎么处理呢？谢谢！
```
```text
 我如果猜测没错，你想写的代码应该是
t_list = []
t_list += list(jieba.cut(sample['content']))

如果不是你应该把你具体哪句出错列出来，不过很可能是python 语法不熟造成的
```
```text
 您好，谢谢您的热心回答！第一次在github发问题不是很清楚问法，给您造成麻烦啦～

问题我已经解决啦，是我数据本身的问题，让我加入一个conditional loop就OK啦～谢谢您！

祝好～

2018-04-19 15:47 GMT+08:00 MoreFreeze <notifications@github.com>:
…
 我如果猜测没错，你想写的代码应该是

 t_list = []
 t_list += list(jieba.cut(sample['content']))

 如果不是你应该把你具体哪句出错列出来，不过很可能是python 语法不熟造成的

 —
 You are receiving this because you authored the thread.
 Reply to this email directly, view it on GitHub
 <#617 (comment)>, or mute
 the thread
 <https://github.com/notifications/unsubscribe-auth/AhgNuVNwsxwnvtduCq8KfQGDIgnVnRlhks5tqEEhgaJpZM4TPapa>
 .
```
## 63.关于 re_han_cut_all 和 re_han_default 所设定的正则不同的问题
Apr 7, 2018
```text
 根据不同的参数名称调用不同的自定义词典和stopwords来分词,自定义词典和stopwords全部初始化加载到内存中,实际分词的时候根据参数来替换dict和stopwords的变量值,这样能实现吗?
```
```text
 I use spark cluster to seg word and submit the program with jieba to the cluster as following:
/home/hadoop/spark-1.2.0-bin-hadoop2.4/bin/spark-submit --master spark://hz0124:7077 --driver-memory 1g --executor-memory 5g --total-executor-cores 60 --py-files ../sparkcode/dependency/dependency.zip seg
word_test.py
the jieba is zipped as the file ../sparkcode/dependency/dependency.zip
then an error occured:

Anyone meets this problem? how to deal with it?
```
```text
 Using zip will cause problems to read the dictionary because it uses regular file IO.
You can unzip the dependencies, or put the dict.txt elsewhere (not in a zip) and then manually initialize jieba with jieba.initialize('/path/to/dictionary').
```
```text
 发现在cut_all模式下，字符串中的特殊字符不能得到很好的返回，查看代码发现这两种模式下的正则表达式设定是不同的，想请教一下，为何两种模式下，设定的正则是不同的
测试结果如图所示：
精确模式

全分割模式
```
```text
 同遇到该问题，怎么回事
```
## 64.3只鸡 被分成了 [3,只鸡]
Apr 3, 2018
```text
 Command "/usr/bin/python2 -u -c "import setuptools, tokenize;file='/tmp/pip-build-6jExFt/jieba/setup.py';f=getattr(tokenize, 'open', open)(file);code=f.read().replace('\r\n', '\n');f.close();exec(compile(code, file, 'exec'))" install --record /tmp/pip-xAfrwY-record/install-record.txt --single-version-externally-managed --compile" failed with error code -9 in /tmp/pip-build-6jExFt/jieba/
```
```text
 Can you provide full error message?
```
```text
 请问dict.txt 词库是否都来自搜狗的2006版免费词库? #24
```
```text
 如果词典中有个关键词是ck，英文单词lock会被切分为lo、ck。
```
```text
 数词+量词+单字名词 很多都有这个问题……
```
```text
 jieba是默认把字母与数字同汉字分开的
```
## 65.如何实现单个字的分词及统计，希望能画出单个字的云图
Apr 2, 2018
```text
 如果我输入 “15路公交” 或“150平米以上的房子”
现在的切割是 15/公交 150/平米/以上/的/房子
我自定义词典里词频调成 0 或100都不行，谢谢
```
```text
 同问
能不能加入正则规则？
```
```text
 能否考虑，先切分再用某种方法合并？
```
```text
 嗯嗯...
```
```text
 您解决了么？
```
```text
 大工程啊 现在没有好用的中文引擎
```
```text
 我也想问这个问题，Coreseek是基于sphinx的。。
```
```text
 那几个分词太让人纠结了
```
```text
 @fxsjy 老大考虑下
```
```text
 很多问题都是针对2个字以上词语进行分词的，如何利用JIEBA进行单个字的分词，研究需要研究单个字的出现词频，不需要词组词频，请指点
```
```text
 单字切分不需要结巴了吧？先去停用词，去标点符号
corpus = '我永远喜欢结城明日奈'
list = []
for each in corpus:
list.append(each)
```
## 66.3.6的协程语法完全无用,多协程无提升
Mar 25, 2018
```text
 wiki里“对Python中文分词模块结巴分词算法过程的理解和分析”跳转到黄网。
```
```text
 哈哈哈哈哈哈哈

2017-11-02 11:37 GMT+08:00 kercker <notifications@github.com>:
…
 wiki里“对Python中文分词模块结巴分词算法过程的理解和分析”跳转到黄网。

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#541>, or mute the thread
 <https://github.com/notifications/unsubscribe-auth/ABD5_fZrOpJSHG1DaFGVx-iT7X4t9A95ks5syTkJgaJpZM4QPIGV>
 .
```
```text
 6666666666老铁网站发出来呀
```
```text
 又骗我点击增加csdn浏览量。。。。。。。。。。。。。。。。。
```
```text
 我看到CRF分词算法的介绍，http://blog.csdn.net/ifengle/article/details/3849852
感觉还行。不知道jieba分词是怎么看CRF的？或者已经用了CRF?
```
```text
 从效果上看，CRF可以有更好的切分效果。百度已经在工程上应用了。
```
```text
 @chenweican 真的吗？感觉crf资料不是很多啊。
```
```text
 import jieba
Traceback (most recent call last):
File "", line 1, in
File "/Users/mengbin/jython2.7.0/Lib/site-packages/jieba/init.py", line 15, in
from ._compat import *
ImportError: No module named _compat
```
```text
 No description provided.
```
```text
 如果是始终占用着cpu当然无用，如果是http请求或者其他i/o操作，你就能看到协程的真正强大之处了
```
## 67.为什么会建立一个缓存文件,单独一个0.889秒,会不会每次都建立缓存.
Mar 25, 2018
```text
 在使用 load_userdict() 时，并不会主动关闭文件，导致程序警告
ResourceWarning: unclosed file <_io.BufferedReader name='dict.txt'>

不知道没有关闭文件的意义是什么？还有后续操作嘛？
```
```text
 jieba 的词典是用空格作为词，词频，词性之间的分隔符的，但是当要在句子中添加“CH125 Spacy”这种词的时候就会识别不了。
如果词典使用“\t”是不是会相对较好？
```
```text
 @rockyzhengwu, @fxsjy 您好, 請問您有解決嗎?我也想要自訂字典中加有空白的英文組合字，但都識別不出來，　都會拆成好幾個字...
```
```text
 我知道dict.txt是jieba的詞庫，請問當程式用jieba lib分析一篇文章時，
如何列出該篇文章不在dict.txt的所有關鍵詞?
底下是我想到的簡單作法:
是先用cut把文章做分詞之後，然後再一一與dict.txt的詞比較，不知道有沒有其他方式?
```
```text
 我知道dict.txt是jieba的词库，请问当程式用jieba的lib分析一篇文章时，
如何列出该篇文章不在dict.txt的所有关键词？
底下是我想到的简单作法：
是先用cut把文章做分詞之後，然後再一一與dict.txt的詞比較，不知道有沒有其他方式?
```
```text
 No description provided.
```
## 68.千与千寻和浪客剑心居然是成语
Mar 21, 2018
```text
 No description provided.
```
```text
 jieba.0.38

不加载自定义词时，分词正常

from jieba import posseg as pseg
pseg.cut(u"为何 MOTO 在华销售的绝大多数")

为何/r MOTO/eng 销售/vn


加载自定义词时，分词异常

from jieba import posseg as pseg
jieba.load_userdict("???")

pseg.cut(u"为何 MOTO 在华销售的绝大多数")

为何/r M/x OTO/x 销售/vn 

自定义词典中含有 OTO
```
```text
 使用jieba.pooseg.cut发现千与千寻和浪客剑心的词性标记都为i（成语）。我将信将疑的查了下发现这两个词都仅作为动漫/电影名词出现。
查看后发现，dict.txt中千与千寻和浪客剑心的标记都是3 i。我认为改成3 n会不会比较合适呢？毕竟dict.txt中大部分动漫名词都被标记为n。
```
## 69.bug? add_word失败了
Mar 19, 2018
```text
 例如：
[笑CRY]  500
将上述整体加入词库中，并进行切分。
目前看，做不到。
是需要做哪些调整吗？词典或代码？
```
```text
 补充：上例中的500是词频。
```
```text
 我改了一下，目前支持词库中的符号和空格匹配了 https://github.com/WalkerWang731/jieba
```
```text
 你好，我很想知道，jieba如何去除标点符号
```
```text
 纯Python实现，无需jieba
punct = set(u''':!),.:;?]}¢'"、。〉》」』】〕〗〞︰︱︳﹐､﹒
﹔﹕﹖﹗﹚﹜﹞！），．：；？｜｝︴︶︸︺︼︾﹀﹂﹄﹏､～￠
々‖•·ˇˉ―--′’”([{£¥'"‵〈《「『【〔〖（［｛￡￥〝︵︷︹︻
︽︿﹁﹃﹙﹛﹝（｛“‘-—_…''')
# 对str/unicode
filterpunt = lambda s: ''.join(filter(lambda x: x not in punct, s))
# 对list
filterpuntl = lambda l: list(filter(lambda x: x not in punct, l))
```
```text
 楼上的回答看上去好酷炫。
```
```text
 吊到无法直视。谢了！
------------------ 原始邮件 ------------------
发件人: "Yanyi Wu";notifications@github.com;
发送时间: 2014年7月19日(星期六) 下午5:46
收件人: "fxsjy/jieba"jieba@noreply.github.com;
抄送: "发如雪止"734133872@qq.com;
主题: Re: [jieba] 关于标点符号 (#169)
楼上的回答看上去好酷炫。
—
Reply to this email directly or view it on GitHub.
```
```text
 我覺得可以幫 jieba 增加一個 jieba.trimPunct(content) 的 method 讓有需要的人可以使用
```
```text
 不好意思，想請問一下類似的問題
由於我目前在處理網頁資料，爬取下來的內容會有一些雜訊，類似
&nbsp;和&gt;

我該如何將其濾掉呢？
```
```text
 不好意思，现在才回复，我觉得你可以先做一遍文本过滤再用jieba分词。
可以先把里面的标点符号过滤掉。
acemoon0301@163.com
发件人： 2153030
发送时间： 2014-12-25 02:06
收件人： fxsjy/jieba
抄送： AcemoonMa
主题： Re: [jieba] 关于标点符号 (#169)
不好意思，想請問一下類似的問題
由於我目前在處理網頁資料，爬取下來的內容會有" "、">" 等雜訊
我該如何將其濾掉呢？
—
Reply to this email directly or view it on GitHub.
```
```text
 @AcemoonMa 不好意思，因為抓取的是html裡的內容，裡面的" "、">"是以字串呈現
空白是 "&+n+b+s+p+;" 將加號都去除，然而用gumblex大的方法會將所有的"&","n","b","s","p",";"都濾掉
昨天我找到了一個方法，可以濾掉字串裡的字串，應該也可以運用在分詞前的濾掉停用字
http://stackoverflow.com/questions/6116978/python-replace-multiple-strings
import re

rep = {"condition1": "", "condition2": "text"} # define desired replacements here

# use these three lines to do the replacement
rep = dict((re.escape(k), v) for k, v in rep.iteritems())
pattern = re.compile("|".join(rep.keys()))
text = pattern.sub(lambda m: rep[re.escape(m.group(0))], text)
For example:
>>> pattern.sub(lambda m: rep[re.escape(m.group(0))], "(condition1) and --condition2--")
'() and --text--'
還沒有實際測試，測試過後跟大家回報
```
```text
 @2153030 HTML 的话最好用 BeautifulSoup 之类的解析库预处理提取字符串（如它的 stripped_strings），再进行分词等自然语言方面的操作。
```
```text
 如果是html解析有很多库都可以提取你需要的text,tag,attribute这些数据啊。
至于网页本身的架构也是可以获取完整的。
例如lxml,beautifulsoap以及python自带的库。
如果你获取后的数据中仍有这些字符，你可以考虑直接写一个字符集合，然后用最基础的循环过滤出来啊。
或者直接用unicode编码过滤，把除了中文，英文，数字以外的都过滤掉就可以了。
acemoon0301@163.com
发件人： 2153030
发送时间： 2014-12-26 15:36
收件人： fxsjy/jieba
抄送： AcemoonMa
主题： Re: [jieba] 关于标点符号 (#169)
@AcemoonMa 不好意思，因為抓取的是html裡的內容，裡面的" "、">"是以字串呈現
空白是 "&+n+b+s+p+;" 將加號都去除，然而用gumblex大的方法會將所有的"&","n","b","s","p",";"都濾掉
昨天我找到了一個方法，可以濾掉字串裡的字串，應該也可以運用在分詞前的濾掉停用字
http://stackoverflow.com/questions/6116978/python-replace-multiple-strings
import re
rep = {"condition1": "", "condition2": "text"} # define desired replacements here
use these three lines to do the replacement
rep = dict((re.escape(k), v) for k, v in rep.iteritems())
pattern = re.compile("|".join(rep.keys()))
text = pattern.sub(lambda m: rep[re.escape(m.group(0))], text)
For example:



pattern.sub(lambda m: rep[re.escape(m.group(0))], "(condition1) and --condition2--")
'() and --text--'
—
Reply to this email directly or view it on GitHub.
```
```text
 你如果有常用的社交或者通讯软件，你可以发软件名和ID，我加你好友，共同探讨，邮件不太方便。
acemoon0301@163.com
发件人： 2153030
发送时间： 2014-12-26 15:36
收件人： fxsjy/jieba
抄送： AcemoonMa
主题： Re: [jieba] 关于标点符号 (#169)
@AcemoonMa 不好意思，因為抓取的是html裡的內容，裡面的" "、">"是以字串呈現
空白是 "&+n+b+s+p+;" 將加號都去除，然而用gumblex大的方法會將所有的"&","n","b","s","p",";"都濾掉
昨天我找到了一個方法，可以濾掉字串裡的字串，應該也可以運用在分詞前的濾掉停用字
http://stackoverflow.com/questions/6116978/python-replace-multiple-strings
import re
rep = {"condition1": "", "condition2": "text"} # define desired replacements here
use these three lines to do the replacement
rep = dict((re.escape(k), v) for k, v in rep.iteritems())
pattern = re.compile("|".join(rep.keys()))
text = pattern.sub(lambda m: rep[re.escape(m.group(0))], text)
For example:



pattern.sub(lambda m: rep[re.escape(m.group(0))], "(condition1) and --condition2--")
'() and --text--'
—
Reply to this email directly or view it on GitHub.
```
```text
 @gumblex @AcemoonMa 謝謝兩位大大的提示
這樣看起來似乎讓實作方便些，我去研究一下BeutifulSoup
@AcemoonMa 郵件已發
```
```text
 楼主 python菜鸟问个简单但急需回答的问题  请问你这个怎么用呢？能写个例子吗？
```
```text
 我想我知道啦  多谢楼主。
```
```text
 import re

def delete_punctuation(text):
    """删除标点符号"""
    text = re.sub(r'[^0-9A-Za-z\u4E00-\u9FFF]+', ' ', text)
    return text

这个满足你的要求吗？
```
```text
 （1）我看有人说“将用户词典覆盖jieba/dict.txt 即可”，有人说“用load.userdict方法添加用户词典”。所以请问是不是上述两种方法都可行？
（2）如果使用后一种方法的话呢，用户词典和jieba默认自带的dict.txt 是同时起作用？还是只有用户词典起作用呢？  “起作用”的意思==“jieba用谁”
（3）如果一个单词，例如“天安门”，同时在dict.txt 和用户词典中出现并且带有不同的词频，那么用词频计算时，是选择哪一个词频呢？不用引起冲突么？
```
```text
 O_O
Out[4]: '<song>丑女七回头</song>'
jieba.lcut("<song>丑女七回头</song>")
Out[5]: ['<', 'song', '>', '丑女', '七', '回头', '<', '/', 'song', '>']
jieba.add_word(O_O)
jieba.lcut("<song>丑女七回头</song>")
Out[7]: ['<', 'song', '>', '丑女', '七', '回头', '<', '/', 'song', '>']
jieba.lcut(O_O)
Out[8]: ['<', 'song', '>', '丑女', '七', '回头', '<', '/', 'song', '>']
```
## 70.带中横杆"-"的词如何把它当做一个词
Mar 16, 2018
```text
 文档中写到，词频省略时使用自动计算能保证分出该词的词频。
例如，“吉林省延边朝鲜族自治州”，期望分词为“吉林省”和“延边朝鲜族自治州”。添加自定义词典并省略词频，并没有成功分出期望的结果。
In [126]: jieba.lcut('吉林省延边朝鲜族自治州',  HMM=False)
Out[126]: ['吉林省延边朝鲜族自治州']
请指点！
```
```text
 jieba.del_word('吉林省延边朝鲜族自治州')
```
```text
 @Silencezjl 这样一个一个del_word也不是办法...似乎得完全自定义词典
```
```text
 使用jieba提取readme内容中的关键词，返回的结果是：
jieba,the,cut,to,list
这个结果中是否可以将the，to剔除掉，有没有什么选项？
```
```text
 @yukaizhao, 主要是idf.txt中只有中文词汇。我会加强这一块，需要分析英文语料。
```
```text
 @fxsjy 我也碰到这个问题，但我试着添加了英文的idf数据，但是仍然有the,to这样的英文出现，能否添加stop words？
```
```text
 @fay, @yukaizhao , 最新版本已经对关键词提取功能做了一些改进，也加了一些英文的stop words。 抽取效果Demo: http://jiebademo.ap01.aws.af.cm/extract
```
```text
 试了一下，效果好多了，正好有件事儿请教，现在的master版本是稳定版本吗，我想换一个小点的词典，更新词典后发现和以前的jieba版本不兼容。
```
```text
 @yukaizhao , master版本不是稳定版，但是应该没有兼容性的问题吧。出错信息是什么？
```
```text
 确实是词典的格式不一样，新的词典添加词性，老词典没有词性
老词典：
一万三千五百斤 4
新词典：
她 134035 r
请问哪儿可以下载最新的稳定版呀？
```
```text
 https://pypi.python.org/pypi/jieba
```
```text
 我自己在网上爬了1500个网页评论，训练了一个idf.txt，然后读取前面的名词，也就是字典里有 "n"的词，效果还很不错。有需要我可以给你。
On Thu, Aug 22, 2013 at 1:26 PM, Zoey Young notifications@github.comwrote:

抽取关键词中含...
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/27#issuecomment-23069212
.
```
```text
 @jannson 求教怎么训练idf
```
```text
 如何将一些带中横杆"-"的词如何把它当做一个词，比如一些药品名：PD-1, PDL-1 这种词当做一个词，而不是分成："PD", "-", "1"
```
```text
 我也没找到什么好的方法，只能自己定义了一个字典，添加这些带-的词，并将频率设置高一些。
```
```text
 添加了并且词频设置的比较高，还是会拆分成开
```
```text
 那你可以试着改一下源码，参考如下链接：http://blog.csdn.net/wangpei1949/article/details/57077007
```
```text
 好的，谢谢！
```
## 71.使用词典一定能把词典中的词分开吗？
Mar 13, 2018
```text
 sent = "基础上的"
我想把它分词成 基础上/的 。
但是不切分“上的”。
应该怎么实现这样的功能呢？
类似的词还有 “既要有” “越要有”
想分成 既要/有 越要/有
```
```text
 drop word list





来自钉钉专属商务邮箱------------------------------------------------------------------
发件人：OrangeIUH<notifications@github.com>
日　期：2018年10月27日 15:45:09
收件人：fxsjy/jieba<jieba@noreply.github.com>
抄　送：Subscribed<subscribed@noreply.github.com>
主　题：[fxsjy/jieba] 特殊分词控制 (#680)

sent = "基础上的"
 我想把它分词成 基础上/的 。
 但是不切分“上的”。
 应该怎么实现这样的功能呢？
 类似的词还有 “既要有” “越要有”
 想分成 既要/有 越要/有
—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 @fxsjy @tuang
不知道是否有好的解决方法？
```
```text
 目前除了加自定义词典，还不能自动识别。我先做个记号，纳入开发计划。
```
```text
 AT&T是可以的啊
```
```text
 = =
```
```text
 T恤也可以啊。http://jiebademo.ap01.aws.af.cm/
```
```text
 嗯嗯，对的。打开方式问题。感谢耐心指正！
```
```text
 @fxsjy T-shirt 和 google+ 貌似还识别不了
```
```text
 使用suggest_freq或者add_word、load_userdict是否能保证词典中的词一定分出来呢？因为有很多医疗术语，如果不能保证的话，感觉会有很多词会分错。
已经有个词典，能否优先根据词典先按照正向最大匹配（先匹配词典中最长的词，再考虑长度短的词），然后再把词典中没有的词进行分词呢？
@fxsjy
```
## 72.如何删除已添加的用户词典
Mar 12, 2018
```text
 部分地名词性为nr，比如“孙吴县”。建议按照民政部最新的行政区划表更新词库。
2018行政区划代码
```
```text
 如何解决向字典添加“leap motion”这种中间带空格的关键词
```
```text
 @microcao , 目前不行，得改代码。 词典里面的分割符就是空格，此外还要改一处代码里面的正则表达式，把空格也作为成词的合法字符。
```
```text
 @fxsjy 能帮忙搞一下带空格的英文单词的实现吗？感谢！
```
```text
 同问。@fxsjy
```
```text
 @geekan 这会带来另一个问题。如果“leap motion”作为一个词的话，那么就要考虑：搜索引擎模式下leap和motion怎样切分出来？
```
```text
 @anderscui 不了解这里的逻辑，不过这个看起来像是一个通用的问题：

长短词同时存在时，搜索引擎模式能否全切分？我觉得是必须要全切分的。
```
```text
 @geekan 对，所以我在考虑，能否先按照现有方式切分，然后根据一个配置文件将包含空格的词merge，而不是把这样的词直接加入自定义词典。
```
```text
 @anderscui 不是很明白，为啥要单独定义配置文件呢？
是因为搜索引擎模式对空格有特殊处理吗？
```
```text
 进程是长时间运行进程，这个过程中自定义词典会更新，需要在词典更新后使用新的自定义词典。
文档中没有关于这个功能的说明。是否可以增加这个功能呢？
```
```text
 你这个解决了吗？
我现在做法是重新调用load_userdict方法来重新加载自定义词典
```
```text
 请问一下，添加后自定义词典后如何删掉已添加的词典？比如针对项目特性填加如下所示的词典，设置了中特异的词性jp，项目结束后不想使用了，如何清除已添加的词典。关闭软件再打开后发现还是能分词出jp词性的词并给予jp词性
用户词典：
子宫颈阴道上部 jp
子宫口 jp
子宫阔韧带 jp
子宫内膜 jp
```
## 73.加入了字典后在本地环境下测试没问题，但使用了flask的web框架后似乎就不行了
Mar 9, 2018
```text
 在我的程序中import jieba成功，并且能够调用分词cut，但是import jieba.analyse失败，无法提取关键词，提示错误为ImportError: No module named jieba.analyse，但是我打印jieba文件夹目录如下，目录中含有analyse。
['finalseg', 'dict.txt', 'init.pyc', 'analyse', '_compat.py', 'main.py', 'posseg', 'init.py', '_compat.pyc']
采用from jieba import analyse的方式能够import成功，但是无法调用extract_tags.@fxsjy
```
```text
 @ShaoyongHua 请问是如何解决的？
```
```text
 我可以啊
```
```text
 我也有这个问题，不过我是from jieba import posseg错误
```
```text
 @ocsponge 请问您的问题解决了吗？
```
```text
 在pyspark程序中，引入jieba进行分词，报no module named jieba, 但是在jupyter notebook中，import  jieba的时候是没有任何问题的，求解？
```
```text
 这是spark的问题，不是jieba的问题，第三方包要用特殊方式载入spark的namespace，具体请查spark的文档。
…
On Sun, Oct 7, 2018, 20:31 guangdi ***@***.***> wrote:
 在pyspark程序中，引入jieba进行分词，报no module named jieba, 但是在jupyter
 notebook中，import jieba的时候是没有任何问题的，求解？

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#406 (comment)>, or mute
 the thread
 <https://github.com/notifications/unsubscribe-auth/AA0SqphEvkPSy1dtZjNnLs6MGxgrQ9eqks5uifQFgaJpZM4Ki_yv>
 .
```
```text
 你好，请问下日期；数字这些问题应该怎么解决呢？
```
```text
 不知道你说的是什么数字问题,如果是需要将2017-09-01这样的日期切分成一个单词的话,可以考虑将其这样格式的日期添加到自定义词典中.我目前就是这么处理的效果不错
```
```text
 @aliray ，将这些格式的日期添加到自定义词典中，就可以处理吗？比如加入词库是2017-09-01，想要将2017-10-01分词到一起可能吗？貌似我这边不行啊
```
```text
 [2018-02-02 12:00:00,000] ERROR in app: Exception on /api/getWordFreq [GET]
Traceback (most recent call last):
File "e:\app\python35\lib\site-packages\flask\app.py", line 1982, in wsgi_app
response = self.full_dispatch_request()
File "e:\app\python35\lib\site-packages\flask\app.py", line 1615, in full_dis
atch_request
return self.finalize_request(rv)
File "e:\app\python35\lib\site-packages\flask\app.py", line 1630, in finalize
request
response = self.make_response(rv)
File "e:\app\python35\lib\site-packages\flask\app.py", line 1740, in make_res
onse
rv = self.response_class.force_type(rv, request.environ)
File "e:\app\python35\lib\site-packages\werkzeug\wrappers.py", line 921, in f
rce_type
response = BaseResponse(*_run_wsgi_app(response, environ))
File "e:\app\python35\lib\site-packages\werkzeug\wrappers.py", line 59, in _r
n_wsgi_app
return _run_wsgi_app(*args)
File "e:\app\python35\lib\site-packages\werkzeug\test.py", line 923, in run_w
gi_app
app_rv = app(environ, start_response)
TypeError: 'dict' object is not callable
127.0.0.1 - - [02/Feb/2018 12:00:00] "GET /api/getWordFreq HTTP/1.1" 500 -��
是因为没法取到字典的绝对路径吗？��字典放在站点根目录，加载时写的是： jieba.load_userdict("dict.txt")
�
```
## 74.如何只使用某个自定义的词典?
Mar 8, 2018
```text
 `text = "给我办一个三十的新流量业务"
result = []
for (word, start, end) in jieba.tokenize(text):
pseg_data = [(w, f) for (w, f) in pseg.cut(word)]
print((pseg_data, start, end))
result.append((pseg_data, start, end))
`
([('给', 'p')], 0, 1)
([('我', 'r')], 1, 2)
([('办', 'v')], 2, 3)
([('一个三十', 'm')], 3, 7)
([('的', 'uj')], 7, 8)
([('新', 'a')], 8, 9)
([('流量', 'n')], 9, 11)
([('业务', 'n')], 11, 13)
如果是"两个三十"就可以分开
```
```text
 你好，我在使用结巴分词时，出现同同一词汇词性不同的情况，我使用了动态加词汇，修改dict，导入自己的词汇表，都会出现这种情况，比如我定义：街 1000000000000 ns 我的目标是分割出的“街”都是“ns”，但是在实际中会出现：(m)欢乐(a)街(n)怎么(r)走(v)？  去(v)美食(n)街(ns) 这两种不同的情况，请问是怎么回事呢？
ps:我只是在使用jieba,并不懂里面的理论，如果有基本错误还请原谅。
```
```text
 对同一段文本使用上述两个函数进行分词，分词结果不一样？
如：弄正
在jieba.cut()中分成了”弄“和”正“
在jieba.posseg.cut()中则保持”弄正“
```
```text
 如题，对于不同的用户，我想使用不同的词典，请问要如何才可以做到?并且其它的词典对这一次分词没有作用.
```
## 75.為什麼多進程時只有一個進程在跑？
Mar 2, 2018
```text
 如圖所示

我是用簡體中文維基百科的語料，經過處理之後把符號換成 space，再將整個文本當成一個當成一個 string
餵進去。
下面是我的代碼
import time

start_time = time.time()

jieba.enable_parallel(16)
jieba_words = {word for word in jieba.cut_for_search(text_processed, HMM=True) if len(word) > 1 and len(word) <= 7}

print("Segmentizing took {} seconds.".format(time.time() - start_time))

謝謝幫忙！
```
## 76.请问能否自定义正则匹配
Feb 25, 2018
```text
 比如将XXXX年、XX月视作一个词
或者我需要将书名号内所有内容视作整个词
能不能进行自定义配置？
```
```text
 用字典不就行了。匹配正则，想下这效率得有多慢。
```
## 77.关于“云计算”、“石墨烯”等词语的词性标注问题
Feb 25, 2018
```text
 在使用pseg.cut来获取分词词性的时候，发现“云计算”、“石墨烯”这些词语的词性居然是x，x我看介绍通常是一些标点符号，为什么会出现这种情况，还有其他词语也会有这种情况吗？
```
```text
 因为云计算，石墨烯是新词，词典里找不到，估计就默认设置成词性x了，你需要自己自定义词典并且加入这些新词的词性。
```
## 78.Questiong) extract words from sentence with word base.
Feb 25, 2018
```text
 Hi I'm using Jieba with Java and I have a question.
My purpose to use Jieba is getting words from sentences and add mean per word. But sometimes the result is not what I expected.
sentence : 我要九个
result : 我要 / 九个
exptected : 我 / 要 / 九 / 个
sentence : 吃三丸
result : "吃三丸"
expected : 吃 / 三 / 丸
Because I have a Chinese English dictionary like this.
我 : I
要 : want
九 : nine
个 :  individual; measure word for people or objects
But for current result I have to have all possible cases in the dictionary.
ex）一个， 两个，三个，四个，六个，七个，八个，九个，是个。。。。
Thanks in advance.
```
## 79."书品相"切词问题
Feb 13, 2018
```text
 "书品相”，无论用add word，还是大幅提高"品相"，均无法切出”书“和”品相“这一结果。
```
## 80.word_add
Feb 8, 2018
```text
 请问add_word(word, freq=None, tag=None) 这个函数增加的词语只能在当前的python jieba 编译环境下生效，python退出重新登陆时，修改的无效。
现在我需要永久性新增或者删除词语，如何实现永久性热更新操作？
谢谢
```
```text
 add_word()只添加到内存中，jieba好像没有提供固化更新字典的实现。但字典格式好像是通用的，可以用一些其他字典生成工具配套使用。
```
```text
 希望支持数据库词库..
```
## 81.demo failed
Feb 2, 2018
```text
 http://jiebademo.ap01.aws.af.cm/
not work
```
## 82.英文分词单词被切分
Feb 2, 2018
```text
 用英文试了一下结巴的分词，但是出来的单词很奇怪，很多好像都被切掉了。
如图

想问下可以支持用自己的分词结果去做比如extract嘛？就比如传list而不是text
```
## 83.`add_word` doesn't work with punctuation marks
Feb 1, 2018
```text
 I would like to treat !! as one token instead of 2.
jieba.add_word(u'!!')
for x in jieba.cut(u'巨大的恶魔!!'):
    print x
巨大
的
恶魔
!
!
```
## 84.是否提供清除suggest_freq结果的函数?
Jan 31, 2018
```text
 在自己的代码中用suggest_freq(word)操作确保特殊词不被拆分, 但同时代码其他部分引用的外部模块也同时引用了jieba, suggest_freq(word)的操作会影响外部模块分词操作(并不是希望看到的). 对此是否提供清除suggest_freq的操作?
```
```text
 import importlib
importlib.reload(jieba) is okey...
```
```text
 by using importlib.reload method, you may encounter this scenario in iPython/jupyter:
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
...
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.

Once more you reload jieba module, there is one more line output in each section when jieba.dt.initialize() is implemented.
I'm not sure if there's any influence to other codes.
```
## 85.和Stanford Word Segmenter比较，以及如何在论文中引用？
Jan 24, 2018
```text
 打扰，请问这个工具用的算法和Stanford Word Segmenter用的算法有什么主要区别吗？
以及，如果使用了这个工具，需要引用你们的论文吗？（或者说这个工具有基于某篇论文吗？）——我也可以直接引用github链接。。。
```
## 86.jieba.analyse.textrank()方法提取英文文本的关键词没有输出数据，但是tfidf()方法则可以处理英文文本。
Jan 24, 2018
```text
 jieba.analyse.textrank()方法提取英文文本的关键词没有输出数据，但是jieba.analyse.tfidf()方法则可以处理英文文本。中文文本两个都可以处理。so ,问题是，，，
```
```text
 你看看源代码就知道了，tokenizer和postokenizer的区别。
```
## 87.re_num判断数字的问题
Jan 17, 2018
```text
 比如要分词的字符串：http://192.168.0.121/login.view
因为：re_num = re.compile("[.0-9]+")这样写，
分词的结果为：
login eng
. m
view eng
"."被判断成数字，建议修改一下这个正则
```
```text
 这个“.”  很有必要被当做数字，信我。
```
## 88.词频省略时使用自动计算的能保证分出该词的词频
Jan 12, 2018
```text
 文档中写到，词频省略时使用自动计算能保证分出该词的词频。
例如，“吉林省延边朝鲜族自治州”，期望分词为“吉林省”和“延边朝鲜族自治州”。添加自定义词典并省略词频，并没有成功分出期望的结果。
In [126]: jieba.lcut('吉林省延边朝鲜族自治州',  HMM=False)
Out[126]: ['吉林省延边朝鲜族自治州']
请指点！
```
```text
 jieba.del_word('吉林省延边朝鲜族自治州')
```
```text
 @Silencezjl 这样一个一个del_word也不是办法...似乎得完全自定义词典
```
## 89.如何中途换词典user dict?
Jan 11, 2018
```text
 刚开始导入库时设定了自定义词典:
import jieba
jieba.load_userdict(
    "/home/weiwu/share/deep_learning/data/dict/finance_dict.txt")


如果程序中途想更换另一个词典，请问该怎么做？
```
```text
 reload(jieba)?
```
## 90.分词词典无效
Jan 10, 2018
```text
 我尝试对恒大相关的文章进行分词，结果错误率奇高，结果如下
'''
按照原计划，凯赛尔在西班牙学习三年后就将回国，但对于志向高远的“凯撒”来说，他更希望凭借未来三年的表现能留在西班牙继续深造，从而拉开自己的职业生涯，“**/随恒大/足校在西班牙学习三年后，我希望能留在这里，并开启自己的职业生涯
保监会再发狠/招恒大/人寿被禁止投资股票
这些/对恒大/概念股的影响有多大
然后在10月底三季报披露散户看/到恒大/**买了4.95%是不是要举牌了，许老板就把货都卖给你们了
'''
我将自定义词典中”恒大“的词频调整到了10000都没有任何变化，尝试jieba.suggest_freq('恒大', True)也没有用；尝试将HMM关掉，结果恒大这个词会被一直分成”恒/大“。
```
```text
 能不能再切之前，直接添加：
jieba.add_word("恒大")
```
## 91.请问，如何能将“男性”分词成“男”和“性”呢？
Jan 4, 2018
```text
 No description provided.
```
```text
 用 del_word('男性') 直接将“男性”这个词删掉？
```
## 92.怎么实现jieba在内存中替换自定义词典和stopwords的值?
Jan 3, 2018
```text
 根据不同的参数名称调用不同的自定义词典和stopwords来分词,自定义词典和stopwords全部初始化加载到内存中,实际分词的时候根据参数来替换dict和stopwords的变量值,这样能实现吗?
```
## 93.sensitive
Jan 1, 2018
```text
 No description provided.
```
## 94.请问怎么对日期时间类词进行准确分词
Dec 27, 2017
```text
 jieba.lcut('2017年10月5日或2017-10-03或12:21和12点30分还有十二点三十分')
分出来
['2017年', '10月', '5日', '或', '2017', '-', '10', '-', '03', '或', '12', ':', '21', '和', '12点', '30分', '还有', '十二点', '三十分']
如何分成
['2017年10月5日', '或', '2017-10-3', '或', '12:21', '和', '12点30分', '还有', '十二点三十分']
```
```text
 把这些词汇加入到词典中

发自我的vivo智能手机

sugarZ <notifications@github.com>编写：
…
jieba.lcut('2017年10月5日或2017-10-03或12:21和12点30分还有十二点三十分')
分出来
['2017年', '10月', '5日', '或', '2017', '-', '10', '-', '03', '或', '12', ':', '21', '和', '12点', '30分', '还有', '十二点', '三十分']
如何分成
['2017年10月5日', '或', '2017-10-3', '或', '12:21', '和', '12点30分', '还有', '十二点三十分']
—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 之前也想到了,如果固定的几个还可以,但是非常大量的日期时间格式,没法全部加到字典,这方法太笨了.
```
```text
 @sugarZ，请问lcut可以将“2017年”分词到一起吗，貌似我这边是“2017”，“年”
```
```text
 @JiaWenqi 好像不行,我是加的自定义词典,把最近的年份都添加了进去,还有1-12月0-24点(钟),但是如果把时间也加上就太麻烦了,如果jieba支持特定格式分词配置就好了.
```
```text
 目前想到一个方案，将待分词文本用时间正则进行分割后分段进行分词
```
## 95.怎么将颜文字添加到user dict呢
Dec 26, 2017
```text
 例如（￣▽￣），(｀・ω・´)等颜文字表情。
在处理的时候，即使添加 add_word 或者 load_userdict ，结果都会被分开成：
（   ￣   ▽   ￣   ）
```
## 96.使用jieba3k时，导入自定义user_dict文件时出错
Dec 24, 2017
```text
 Traceback (most recent call last):
File "", line 1, in 
File "/opt/anaconda3/lib/python3.6/site-packages/jieba/init.py", line 119, in wrapped
return fn(*args, **kwargs)
File "/opt/anaconda3/lib/python3.6/site-packages/jieba/init.py", line 307, in load_userdict
word, freq = tup[0], tup[1]
IndexError: list index out of range
```
## 97.如何重新训练模型？
Dec 22, 2017
```text
 结巴的字典是几年前的，不知道能否加入一些新的词重新训练一份？
使用了自定义词典感觉效果不太好。
```
```text
 By theory, jieba is not a trainable model i'm afraid
```
```text
 可以训练啊，你拿别的人工或机器分词/标注的语料，统计一下词频就行。
```
```text
 @gumblex 你好 能详细说说这个过程吗, 我想做一个自己的关键字提取的idf.txt, 但是不知道如何下手, idf.txt中的第二列, 那一串数字, 不知道怎么得出  比如
劳动防护 13.900677652
勞動防護 13.900677652
生化学 13.900677652
```
## 98.请问新发现的词可以导出吗？
Dec 20, 2017
```text
 请问打开新词发现后，发现的新词可以导出吗？
```
```text
 同问，pynlpir有一个get_newword用法，可以print出新词，结巴有没有类似功能？
```
## 99.jieba.posseg.cut 分词将百分比分割成数字和符号，并且标注tag为'x'
Dec 20, 2017
```text
 jieba.posseg.cut是可以分割词并得到tag类型，但是却把百分比分割成两个部分，tag都是'x'，jieba.cut就能准确分成"2%"，如果需要tag还需要百分比，应该怎么写
```
```text
 同样遇到了这个问题
```
```text
 同问
```
```text
 已解决，见https://blog.csdn.net/liuhongyue/article/details/80498195
```
## 100.jieba0.39 add_word添加词需要在并行之前添加
Dec 19, 2017
```text
 jieba.add_word("石墨烯",100,"nr")
jieba.add_word("凱特琳",100,"nr")
jieba.add_word("莫那娄氏",10,"n")
jieba.del_word("自定义词")
不添加词典，利用add_word和del_word仍然不起作用
先开启并行，不起作用。如果先添加词，再进行并行就可以。。。。还真是bug
jieba.enable_parallel()
```
```text
 #547
```
## 101.使用c/c++重写了计算DAG和HMM中的vitrebi函数，速度大幅提升
Dec 18, 2017
```text
 除了函数的替换外，基本上没有改动别处的代码。经过大量测试分词的结果与您的结果相同，开启HMM的普通模式时间缩短了60%左右，关闭HMM的普通模式时间缩短了50%左右。
  请问我可以把这些独立成一个仓库吗，经完整的测试后再request你的readme
```
```text
 我想，你可以把这些c++的程序封装成一个静态 so文件，然后可以增加到这个项目里面，然后py脚本直接调用你的c++so 代码，这样可以提高效率。也可以让这个项目更加的完美。
```
```text
 增加一个 c++的模块在里面，然后，py 脚本函数调用弄一个开关，选择是用py代码还是用c++代码。
```
```text
 *unix下编译.so没啥问题，windows下目前兼容还不好还在调试
```
```text
 windows就不要支持了，你可以弄一个说明，就说，c++代码库只支持linux系统，不对windows提供。
windows就让他们用py代码得了。反正又不跑服务端.
```
```text
 用 except ImportError 这种吧，你先 fork 到自己仓库，作者估计也不怎么管了。
```
```text
 已经封装好，详情见https://github.com/deepcs233/jieba_fast
可使用 pip install jieba_fast安装
```
```text
 安装出错啊，报错少文件。
gcc: 错误：source/jieba_fast_functions_wrap_py3_wrap.c：没有那个文件或目录
gcc: 致命错误：没有输入文件
编译中断。
```
```text
 @george-sq 已经fix，有问题建议到issues里问
```
## 102.centos 7 安装时候报错
Dec 16, 2017
```text
 Command "/usr/bin/python2 -u -c "import setuptools, tokenize;file='/tmp/pip-build-6jExFt/jieba/setup.py';f=getattr(tokenize, 'open', open)(file);code=f.read().replace('\r\n', '\n');f.close();exec(compile(code, file, 'exec'))" install --record /tmp/pip-xAfrwY-record/install-record.txt --single-version-externally-managed --compile" failed with error code -9 in /tmp/pip-build-6jExFt/jieba/
```
```text
 Can you provide full error message?
```
## 103.请问dict.txt是通过什么规则得到的呢？
Dec 8, 2017
```text
 No description provided.
```
```text
 你没看代码吗？
```
```text
 哎，作者就是不爱写注释，哎。
```
```text
 @ilqxejraha 谢谢您，可能是我表达有误，我是想问词汇库里面的词汇和词频是人工统计的吗？还是通过其它的方法。
```
```text
 统计过来的。最后得到了就是这么一个模型。
```
```text
 你在源码中看到词频的使用了吗？
```
```text
 具体的统计词频，词频的作用可能体现在，一个词存在多个意思。
比如英语中，经常有一个词会有很多个意思。
比如出现一个词，并且，这个词有很多种解释，这时候词频可能会对词意的选择有一定帮助。
具体的其他算法hmm的我还没看。
```
```text
 @KevinDotW 据说是基于人民日报的语料库，我也想知道怎么才能创建自己的词典
```
## 104.希望能出一个适用Elasticsearch的插件
Dec 5, 2017
```text
 No description provided.
```
## 105.词性使用的是什么标准？zg是什么词性？
Dec 4, 2017
```text
 No description provided.
```
```text
 https://www.cnblogs.com/adienhsuan/p/5674033.html，可以参考一下
```
```text
 我也遇到这个问题了，求告知一个比较全的词性介绍。
```
```text
 找到一个词性介绍，见jieba
```
```text
 官方应该添加一份词性列表到文档中
```
```text
 同问
```
```text
 应该和这个一样
http://ltp.readthedocs.io/zh_CN/latest/appendix.html#id3
```
## 106.用户自定义字典词频受默认词典的影响权重
Nov 27, 2017
```text
 请问：
结巴分词中用户自定义的字典受默认字典词频的影响是和单个词频的值有关么，例如我自定义词典中增加一个词“汉武帝”，词频为30 ，而默认词典中“汉”和“武帝”的词频都是100，那是不是说我这个词一定会被分成默认字典“汉/武帝”，而不是我自定义词表中的“汉武帝”，如果是这样的话，我后期自定义的字典如果批量加入，我怎么定词频的值合适，还是我需要对一类词进入默认词表中先检查一下，在定义自己的字典，不知道有没有什么好的办法，望回复，感谢！！
```
```text
 我也有这个问题，请问您解决了吗？
```
```text
 我解决了，你在用户自定义词典中不要规定词频，如下
体外药物释放
增效剂
航道整治
肌腱膜纤维肉瘤癌基因同系物A
这样就会自动生成一个合适的词频，从而可以把这个词分出来
```
```text
 @我也碰到了类似的问题，使用楼上的@grandmoi 的方法不写词频和词性会报下面这个错误
IndexError: list index out of range
不写词性，词频写多高都分不出词
难不成真的要自己一个个加到字典里……
```
```text
 @grandmoi 我也有这个问题。我的自定义词典中定义了‘国债’，没有设置词频。然后输入‘中国债’，想分出来‘中’、‘国债’。但是实际输出还是‘中国’、‘债’
```
## 107.自定义词库可否包含中文和英文的组合
Nov 21, 2017
```text
 用户自定义词典定义了中文和英文结合的词 ，分词没用效果
比如词典里定义了一个a字群
分词好像没有效果
```
## 108.jieba分词如何对一句话进行一元分词？比如：“我来到北京清华大学”  分成： 我 / 来 / 到 / 北 / 京 / 清 / 华 / 大 / 学
Nov 21, 2017
```text
 比如：“我来到北京清华大学”  分成： 我 / 来 / 到 / 北 / 京 / 清 / 华 / 大 / 学
```
```text
 这种事儿也用分词么？

2017-11-21 20:00 GMT+08:00 jasmineol <notifications@github.com>:
…
 比如：“我来到北京清华大学” 分成： 我 / 来 / 到 / 北 / 京 / 清 / 华 / 大 / 学

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#552>, or mute the thread
 <https://github.com/notifications/unsubscribe-auth/ABD5_ejuZPwUklkvnwSV_5J3028wkVO3ks5s4rtUgaJpZM4QluiD>
 .
```
```text
 这个网上随便一搜都是分字的
比如把中文字分隔转为list
```
```text
 补充一下，c++如何使用jieba分字，涉及到gbk，utf8这种编码的转换，所以想看下jieba有没有现成的
```
```text
 这种直接获取编码就是一个词吧，之前看过汉子转拼音。
```
```text
 设置一个空词典，关闭HMM新词联想，分出来全是一个个字的
```
## 109.自定义词库中，可否包含“标点符号+中文”的组合？
Nov 15, 2017
```text
 例如：
[笑CRY]  500
将上述整体加入词库中，并进行切分。
目前看，做不到。
是需要做哪些调整吗？词典或代码？
```
```text
 补充：上例中的500是词频。
```
```text
 我改了一下，目前支持词库中的符号和空格匹配了 https://github.com/WalkerWang731/jieba
```
## 110.并行和加入自定义词典的调用顺序对自定义词典的载入有很大的影响
Nov 15, 2017
```text
 jieba.enable_parallel()
jieba.load_userdict("./aux/dict")
如上所示，如果先开启并行，再载入自定义词典，会导致自定义的词典没有产生效果
如果先载入自定义词典再进行并行计算开启，自定义词典就是有效的，不知道这算不算bug
```
```text
 确实如此。
先开启并行，再载入自定义词典。此时，从jieba.dt.get_FREQ和jieba.dt.user_word_tag_tab中都能查到用户词典中的词，说明已导入用户词典中的词，但分词结果中并未产生效果。
```
```text
 这是bug吧？同遇到了，定位了半天问题，求修复啊
```
## 111.是否应该关闭文件
Nov 15, 2017
```text
 在使用 load_userdict() 时，并不会主动关闭文件，导致程序警告
ResourceWarning: unclosed file <_io.BufferedReader name='dict.txt'>

不知道没有关闭文件的意义是什么？还有后续操作嘛？
```
## 112.Add LICENSE to MANIFEST.in
Nov 9, 2017
```text
 Please add the LICENSE file to the MANIFEST.in in both jieba and jieba3k so that it's included in the source distribution.
```
## 113.在线demo似乎崩了
Nov 5, 2017
```text
 无论是否翻墙，均无法访问，也ping不到
```
```text
 @Pzoom522 这个链接已经崩了很久了。
```
## 114.有没有接口返回分词结果的可信度？
Nov 4, 2017
```text
 某些新词或不常用词拆开和合并的分词质量比较？
未来会不会有这种功能的API，方便定量研究？
```
## 115.wiki里“3）对Python中文分词模块结巴分词算法过程的理解和分析”跳转到黄网
Nov 2, 2017
```text
 wiki里“对Python中文分词模块结巴分词算法过程的理解和分析”跳转到黄网。
```
```text
 哈哈哈哈哈哈哈

2017-11-02 11:37 GMT+08:00 kercker <notifications@github.com>:
…
 wiki里“对Python中文分词模块结巴分词算法过程的理解和分析”跳转到黄网。

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#541>, or mute the thread
 <https://github.com/notifications/unsubscribe-auth/ABD5_fZrOpJSHG1DaFGVx-iT7X4t9A95ks5syTkJgaJpZM4QPIGV>
 .
```
```text
 6666666666老铁网站发出来呀
```
```text
 又骗我点击增加csdn浏览量。。。。。。。。。。。。。。。。。
```
## 116.Why do you cut some unicode symbols
Oct 26, 2017
```text
 First of all thank you for your work, you did so far!
I have a question for this specific line https://github.com/fxsjy/jieba/blob/cb0de2973b2fafaa67a0245a14206d8be70db515/jieba/posseg/init.py#L17. Why do you use this specific range of unicode literals? For my specific case(russian text) your app not splitting words which is not good at all.
For example:
>>> re_han_internal = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._]+)")
>>> re_han_internal.split("""Простой и безопасный способ делиться терминалом: обзор инструмента warp""")
['Простой и безопасный способ делиться терминалом: обзор инструмента ', 'warp', '']

Increasing unicode range helps:
>>> re_han_internal = re.compile("([\u0041-\u9FD5a-zA-Z0-9+#&\._]+)")
>>> re_han_internal.split("""Простой и безопасный способ делиться терминалом: обзор инструмента warp""")
['', 'Простой', ' ', 'и', ' ', 'безопасный', ' ', 'способ', ' ', 'делиться', ' ', 'терминалом', ': ', 'обзор', ' ', 'инструмента', ' ', 'warp', '']
```
```text
 it's based on the dict which is ONLY for chinese by default, u wanna use the program, u should add a special russian dict.
```
## 117.多线程导致载入自定义词典不一致问题
Oct 25, 2017
```text
 我在加载了自定义词典之前enable_parallel，发现分词结果并没有使用词典。如果在加载自定义词典之后再enable_parallel则能按照预期结果分词。
#encoding=utf-8
import sys
import jieba
jieba.enable_parallel(8)
jieba.load_userdict('user.dict')

words = jieba.cut("热水器加热时间太长")
print('/'.join(words))

结果是 “热水器/加热/时间/太/长”
#encoding=utf-8
import sys
import jieba
jieba.load_userdict('user.dict') 
jieba.enable_parallel(8)      # 放在load_userdict后面了

words = jieba.cut("热水器加热时间太长")
print('/'.join(words))

结果是*“热水器/加热时间/太长”*
词典里是有“加热时间”，所以可以猜测第一个代码分词前没有载入自定义词典，可能是因为enable_parallel放load_userdict前面的。这不一致的情况是不是多线程读入词典的bug？
```
## 118.jieba.add_word怎么不生效
Oct 23, 2017
```text
 比如我有一句话，"2015年度本公司营业利润较2014年度减少-3480.83万元，降幅为-26.95%，主要原因系国内宏观经济增速放缓，下游市场需求不足，市场景气程度有所降低，公司销售收入整体下降。"
我先对结巴做了add_word设置
jieba.add_word("-26.95%")
jieba.add_word(u"-3480.83万元")
然后分词结果是：
2015年度|本|公司|营业利润|较|2014年度|减少|-|3480.83万元|，|降幅|为|-|26.95%|，|主要|原因|系|国内|宏观经济|增速|放缓|，|下游|市场需求|不足|，|市场|景气|程度|有所|降低|，|公司|销售收入|整体|下降|。|
带符号的就不行
```
```text
 遇到同样问题。
加入用户字典    test-test 词 依然会被切分成 test/-/test。
提高词频依然没用，不知道是不是 -  会被固定切出来，不收其他影响。
```
```text
 同样的问题，比如：
jieba.add_word("弗里德里克·泰勒",freq=888888); jieba.add_word("冯·诺伊曼",freq=888888); jieba.suggest_freq(('弗里德里克·泰勒', '冯·诺伊曼'), True); txt = "这个外国人叫弗里德里克·泰勒，是美国前国防部部长。冯·诺伊曼是公认的计算机发明人"; print(" | ".join(jieba.cut(txt, HMM=True, cut_all=False)))
结果还是：
'这个 | 外国人 | 叫 | 弗里德里克 | · | 泰勒 | ， | 是 | 美国 | 前 | 国防部 | 部长 | 。 | 冯 | · | 诺伊曼 | 是 | 公认 | 的 | 计算机 | 发明人'
大家有什么好办法？
```
```text
 有同样的需求，需要保证指定的词绝对不分。不知该怎么做。
```
```text
 代码里面有如下的正则表达式，用来提取中文

re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&._%]+)", re.U)
也就是说对于”弗里德里克·泰勒”中的“·”不认为是中文，而是作为类似逗号和句号来处理。

一个不太好但是可以用的办法就是，修改上面的正则表达式，将“·”加入。其中\xb7就是“·”。

re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&.\xb7%]+)", re.U)

测试结果：
》 这个 | 外国人 | 叫 | 弗里德里克·泰勒 | ， | 是 | 美国 | 前 | 国防部 | 部长 | 。 | 冯·诺伊曼 | 是 | 公认 | 的 | 计算机 | 发明人
```
## 119.用户实例化`Tokenizer()`对象,`suggest_freq`不起作用
Oct 18, 2017
```text
 class Tokenizer(object):
    ..
    def suggest_freq(self, segment, tune=False):
        ..
        if tune:
            add_word(word, freq)
此处,add_word(word, freq)应为self.add_word(word, freq),否则更新的是默认分词器jieba.dt,而不是用户的new_dt = Tokenizer().
```
## 120.可以支持特殊规则不分词吗
Oct 17, 2017
```text
 如 我的需求是""号里面的话不分词
```
## 121.为什么不能用自己的语料训练新的HMM模型
Oct 6, 2017
```text
 java版本的分词Ansj和HanLP都支持用自己的语料再训练，为什么结巴分词这么久也不支持？
```
## 122.POS list
Oct 4, 2017
```text
 Does anyone have a list of the POS tags, and what they stand for?  They're supposed to be compatible with ictclas.  Thanks.
```
## 123.怎么去除“Building prefix dict from the default dictionary ...”的提示
Sep 28, 2017
```text
 您好，每一次我用jieba进行分词的时候，都会有
Building prefix dict from the default dictionary ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.128 seconds.
Prefix dict has been built succesfully.

这样的提示。怎么去除这些提示呢？
```
```text
 找到jieba库的__init__.py， 大约28~30行左右，
log_console = logging.StreamHandler(sys.stderr)
default_logger = logging.getLogger(__name__)
default_logger.setLevel(logging.DEBUG)
default_logger.addHandler(log_console)

将default_logger.setLevel(logging.DEBUG) 改为 default_logger.setLevel(logging.INFO) 试试。
```
```text
 @oisc it shows error:
Traceback (most recent call last):
  File "xxx.py", line 29, in <module>
    jieba_logger.handlers.clear()
AttributeError: 'list' object has no attribute 'clear'
```
```text
 @IvyGongoogle
try del jieba_logger.handlers[:] if you are using Python2
```
```text
 @aqua7regia it works. thank you~~
```
```text
 jieba.setLogLevel(20)
或
import logging
jieba.setLogLevel(logging.INFO)
```
```text
 @oisc  sorry, What do you mean for del jieba_logger.handlers[:]?
```
```text
 @gaoyangthu it works, but It seems to be jieba.setLogLevel(logging.INFO)
```
```text
 @IvyGongoogle You're right!
setLogLevel(0) equals setLogLevel(logging.NOTSET)
setLogLevel(10) equals setLogLevel(logging.DEBUG)
setLogLevel(20) equals setLogLevel(logging.INFO)
setLogLevel(30) equals setLogLevel(logging.WARNING)
setLogLevel(40) equals setLogLevel(logging.ERROR)
setLogLevel(50) equals setLogLevel(logging.CRITICAL)
```
```text
 @gaoyangthu  yes, thank you very much~~
```
## 124.我想用jieba分词后，只想提出里面的中文分词，不要标点符号，怎么用python处理啊 谢谢
Sep 27, 2017
```text
 No description provided.
```
```text
 我是用正则表达式处理的，new_sentence = re.sub(r'[^\u4e00-\u9fa5]', ' ', old_sentence) 然后再进行分词的, \u4e00-\u9fa5这个是utf-8中，中文编码的范围
```
```text
 @cbzhuang  非常谢谢你的回复！ 我用了这个，不知道可对。#169
```
```text
 Actually, CJK characters are encoded together so there's no critical range for Chinese characters. A punctuation dict could be used to do the filtering.
```
## 125.有安卓版本吗,java版本里面有些类,sdk舍弃了
Sep 22, 2017
```text
 No description provided.
```
## 126.带有罗马数字的自定义词无法正确识别
Sep 21, 2017
```text
 我在自定义词典里加入了带有罗马数字的词，如“纽约Ⅰ线”，并且为其赋予很高的词频，但是总是被分隔开。请问我要如何才能强制把“纽约Ⅰ线”这样的自定义词完整切分出来？？？
词典相关内容如下：
纽约Ⅰ线 100000 ns
python版本3.6.2
代码大致如下：
line = "并明确控制纽约Ⅰ线功率"
jieba.load_userdict("dict.txt")
print(jieba.get_FREQ("纽约Ⅰ线"))
print(jieba.get_FREQ("纽约"))
print(jieba.get_FREQ("Ⅰ"))
print(jieba.get_FREQ("线"))
print(jieba.get_FREQ("Ⅰ线"))
words = posseg.cut(line)
for word in words:
print(word)
输出结果如下：
100000
1758
None
7688
None
并/c
明确/ad
控制/updown
纽约/ns
Ⅰ/x
线/n
功率/n
```
## 127.如何優化關鍵詞提取時間
Sep 21, 2017
```text
 每一次提取一句話的關鍵詞的速度平均都在6s左右如何優化這個過程
```
## 128.关于textrank.py中rank()方法的探讨
Sep 21, 2017
```text
 textrank中的rank()方法参考了pagerank的思路，并且在程序中iterate 10次。但是对于jieba案例，可否考虑把阻尼因子设定为1，基于这一点，可以推断出每一个节点的weight就是这个点的连线数，没有必要iterate 10反复迭代了。修改方法如下：
    for n, out in self.graph.items():
        ws[n] = wsdef
        outSum[n] = sum((e[2] for e in out), 0.0)

    # this line for build stable iteration
    # sorted_keys = sorted(self.graph.keys())
    # for x in xrange(10):  # 10 iters
    #     for n in sorted_keys:
    #         s = 0
    #         for e in self.graph[n]:
    #             s += e[2] / outSum[e[1]] * ws[e[1]]
    #         ws[n] = (1 - self.d) + self.d * s
    ws=outSum
    (min_rank, max_rank) = (sys.float_info[0], sys.float_info[3])
```
## 129.我想比较两个文件内容的相似度，算tfidf做weight，但是怎么先算出df呢
Sep 9, 2017
```text
 No description provided.
```
## 130.extract_tags 忽略单字不合理
Sep 6, 2017
```text
 “车、门、窗” 这类单字词无法抽取，这个可以加一个可选参数来来允许单字关键字呢？
```
## 131.如何保证公交和公交线路名不被切割
Sep 1, 2017
```text
 如果我输入 “15路公交” 或“150平米以上的房子”
现在的切割是 15/公交 150/平米/以上/的/房子
我自定义词典里词频调成 0 或100都不行，谢谢
```
```text
 同问
能不能加入正则规则？
```
```text
 能否考虑，先切分再用某种方法合并？
```
## 132.为什么结巴对繁体字的pos-tagging支持的不好呢？
Aug 31, 2017
```text
 举例：
sen = u"我钟爱法国"
我/r 钟爱/nr 法国/ns
sen = u"我鍾愛法國 "
我/r 鍾/nr 愛/v 法國/ns
有什么方法可以改善么？比如可以把简体字的词库改成繁体字词库么？
```
```text
 改用dict.big.txt的繁體字典後, 並參考我修正的錯誤 #670
```
## 133.如何将iphone 5，xbox one 等分成一个词
Aug 31, 2017
```text
 No description provided.
```
```text
 英文的你要中文分词怎么做，加入词典看看
```
```text
 同样的问题 可以用同义词 但是试了不好用
```
```text
 这里可以将Python ， C++分出来。
```
## 134.增加了自定义词典后，一个完整的单词被强行分开了
Aug 30, 2017
```text
 在我的自定义词典中有这样一个专有名词"ens"，在分词时却将“license”强行分成了“lic/ens/e”。这个要怎么处理这种情况？
```
```text
 在词典里加上“license”，再给个比“ens”较高的数
```
```text
 我觉得jieba应该尊重英文词的边界,当自定义词典里起始或结束是英文字母,例如:



jieba.add_word('w底')
jieba.lcut("太好了w底出现了")  # 这里"w底"的w不是英文词的连续,可以切分
['太好了', 'w底', '出现', '了']
jieba.lcut("wow底出现了") # 这里"w底"的w是英文词的连续,不应该切分
['wo', 'w底', '出现', '了']
```
```text
 @chunsheng-chen 那是不是得有个字典存英文单词，感觉会很大哦
```
```text
 不知道具体细节，但我猜测jieba对英文词的分解是基于类似"[a-zA-Z0-9]*"的模式，所以不需要英文字典，例如:



jieba.lcut("this is a 1test1-abc2! call 911")
['this', ' ', 'is', ' ', 'a', ' ', '1test1', '-', 'abc2', '!', ' ', 'call', ' ', '911']



如果能尊重英文词的自然分割方式，就不会出现上面的情况了: license是一个完整的词，wow是一个完整的词。
```
```text
 自定义词典中部分词含有日文的假名，但是分词出来好像全部无视了。
```
## 135.jieba可以做关键词分配吗？
Aug 30, 2017
```text
 就是我有一个给定的关键词库，然后新来一篇文档，从词库里面找出几个词语作为这篇文档的关键词。
```
```text
 目前我就在做类似的工作，效果挺不错。
你的这种情况，可以这么做。

先将文档按句子切分成多个句子，然后计算关键词库中的每个关键词在这篇文档中的句子集合；
接着计算关键词库中两两关键词的相似性（可以用Jaccard相似性度量），这样构成了一个相似性矩阵；
接着对相似性矩阵进行特征分解，然后对特征值进行归一化；
对归一化的特征值从大到小排序，并累计求和（cumsum），选取前<=0.8的特征值对应的关键词作为这篇文档的关键词

以上是一个基本的版本，直接用，效果一般。因此，需要考虑关键词的tf-idf。我是这么做的，在上面的步骤3时，对特征值进行tfidf加权。最后实验结果很好。
以上，楼主可以试试。
```
```text
 @MacQing 非常感谢，我试一下。
```
## 136.请将许可文件并入pypi包中。 Please add the license file to the pypi package.
Aug 25, 2017
```text
 No description provided.
```
## 137.自定义词典词频计算
Aug 18, 2017
```text
 【词频省略时使用自动计算的能保证分出该词的词频】这句话不太明白。
词频省略的时候，要如何自动计算词频呢（详细），谢谢~
```
## 138.发邮件 为什么不能分成 ‘发 v’  和 ‘邮件 n’
Aug 9, 2017
```text
 No description provided.
```
## 139.请教：jieba适合初学者学习其源码吗？
Aug 8, 2017
```text
 作为一个大四学生，刚刚学过python，想看看一些成熟的项目具体代码，透彻地进行分析，不知道jieba适不适合呢？
```
```text
 個人感覺挺好的, http://www.cnblogs.com/zhbzz2007/tag/Natural%20language%20processing/
```
## 140.想分的词却合成一个词了
Aug 7, 2017
```text
 你好，在做基于情感词典的情绪分析，结果发现，结巴分词，会把很多词合在一起。比如：你真笨；
我想要的结果是：你，真，笨，
但是返回的却是：你，真笨
还有：你太笨，返回的也是：你，太笨
感觉很不科学呐。真和太明明是一种程度词，怎么和笨合成一个词了。
同时调研了其它多个分词开源工具，发现只有你们会合在一起。
```
```text
 是不是只有自定义的词比你们已有的词长才会有效，而比你们的词短的就不生效？
```
```text
 你可以试试lcut(seq,HMM=false)

Brent

On Aug 7, 2017, at 18:55, jiangchao123 <notifications@github.com<mailto:notifications@github.com>> wrote:


是不是只有自定义的词比你们已有的词长才会有效，而比你们的词短的就不生效？

—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub<#503 (comment)>, or mute the thread<https://github.com/notifications/unsubscribe-auth/AM-fgZNs2YqKUicyDBIwL5pGAZrIcKnbks5sVu07gaJpZM4OvHe9>.
```
```text
 试了一下，这样在有些词是可以的，有些依然不行，并且如果HMM设为FALSE的话，好多人名都会被切开了，比如：小明被切成了小，明；李小福被切成了李，小，福；
找代驾我想分成找，代驾；但是HMM设为True为找代驾；HMM设为False为找，代，驾
看来是不能两全？？？
```
## 141.某些关键词抽取不到
Aug 2, 2017
```text
 代码入下：
content="环境设施都很好，服务周到！"
aspect_content_array = jieba.analyse.extract_tags(content, topK=50)
print ",".join(aspect_content_array)
抽取结果：
服务周到,设施,环境
我想得到“好”这个关键词，但是抽取不出来
另外能否在抽关键词的时候，把“服务周到” 抽取成 “服务”和“周到”两个词
```
## 142.AttributeError: 'int' object has no attribute 'decode'
Aug 2, 2017
```text
 上图是我的数据前一部分，我的目的是对 titles 一列进行分词，分词的代码如下。现在遇到的问题是AttributeError: 'int' object has no attribute 'decode'，所以我认为是 titles 中有 int 所致，所以添加了一个判断条件，但是代码执行的结果依旧是报之前的错。请问这是什么原因？
def jiebait(text):
    seglist = jieba.cut(text, cut_all = True)
    fenci = []
    for word in seglist:
        if (not isinstance(word, int)) and (len(word) >= 2):
            fenci.append(word)
    # 如用搜索引擎模式：
    #seglist = jieba.cut_for_search(text)
    return ' '.join(fenci)
```
```text
 def jiebait(text):
    seglist = [str(w) for w in jieba.cut(text, cut_all = True)]
    fenci = []
    for word in seglist:
        if (not isinstance(word, int)) and (len(word) >= 2):
            fenci.append(word)
    # 如用搜索引擎模式：
    #seglist = jieba.cut_for_search(text)
    return ' '.join(fenci)
```
```text
 AttributeError                            Traceback (most recent call last)
 in ()
2     result_line = ""
3     # segment
----> 4     seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]
5     # remove special character
6     temp = re.sub("[\s+.!_,$%^(+"')]+|[+——()?【】“”！，。？、~@#￥%……&（）]+", "",("/".join(seg_list)))
 in (.0)
2     result_line = ""
3     # segment
----> 4     seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]
5     # remove special character
6     temp = re.sub("[\s+.!_,$%^(+"')]+|[+——()?【】“”！，。？、~@#￥%……&（）]+", "",("/".join(seg_list)))
D:\Anaconda\lib\site-packages\jieba_init_.py in cut(self, sentence, cut_all, HMM)
280             - HMM: Whether to use the Hidden Markov Model.
281         '''
--> 282         sentence = strdecode(sentence)
283
284         if cut_all:
D:\Anaconda\lib\site-packages\jieba_compat.py in strdecode(sentence)
35     if not isinstance(sentence, text_type):
36         try:
---> 37             sentence = sentence.decode('utf-8')
38         except UnicodeDecodeError:
39             sentence = sentence.decode('gbk', 'ignore')
AttributeError: 'int' object has no attribute 'decode'
```
```text
 你的 text 本身是不是 int
```
```text
 是的，本身是int，但是seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]这一步，应该都转成string了
```
```text
 解决了 我这边数据转码的问题 不好意思。。。感谢
```
## 143.请问提取关键词时debug信息如何关闭
Aug 1, 2017
```text
 Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/t8/p9fyd38n58v35j4fxrndsyyw0000gn/T/jieba.cache
Loading model cost 0.357 seconds.
Prefix dict has been built succesfully.
```
## 144.Maybe the demo site is down?
Aug 1, 2017
```text
 No description provided.
```
## 145.能自定义词性吗？
Jul 26, 2017
```text
 能自定义词性吗？
将一部分特殊词放在自定义词典里，标注自定义词性(例如:unc)，jieba.cut后得到特殊词的词性？
例如：先生，女士，经理，标注为称谓词
```
```text
 明显可以
你可以把下面的加到自定义字典里
先生 w 1000
女士 w 1000
经理 w 1000

在这里
单行字典的结构是
word pos freq
词 词性 词频
```
```text
 这个格式正确吗？
word freq pos?
还有就是会出现自定义词典后，词的词性变为x
```
```text
 word freq pos 是这个格式啊
```
## 146.请问jieba有没有中文纠错的功能呢？
Jul 26, 2017
```text
 比如遇到本身就有语义或者拼写的错误的中文：咳嗽->咳嗽等等
想从语义和词性上通过字典来实现可以吗？
```
```text
 没有的吧，据我使用的经验来说是没有的。
这里评价体系比较难弄
据我的使用经验而言我们是自己建立了一套转换体系
```
## 147.for help,  jython+jieba failed, ValueError: insecure string pickle
Jul 19, 2017
```text
 met a problem when using jython+jieba
jython version :jython-standalone 2.7.1
jieba :0.38
Exception in thread "main" java.lang.ExceptionInInitializerError Caused by: Traceback (most recent call last): File "jiebademo.py", line 7, in <module> import jieba File "/usr/local/lib/python2.7/site-packages/jieba/__init__.py", line 16, in <module> from . import finalseg File "/usr/local/lib/python2.7/site-packages/jieba/finalseg/__init__.py", line 30, in <module> start_P, trans_P, emit_P = load_model() File "/usr/local/lib/python2.7/site-packages/jieba/finalseg/__init__.py", line 24, in load_model start_p = pickle.load(get_module_res("finalseg", PROB_START_P)) File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 1378, in load File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 858, in load File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 858, in load File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 966, in load_string ValueError: insecure string pickle
any help??
```
```text
 这个需要你把
/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py
line 966
上下文贴出来看看
if you can show the context of
/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py
line 966
maybe i can help you
```
## 148.jieba/__init__.py 48行
Jul 13, 2017
```text
 def setLogLevel(log_level):
global logger
default_logger.setLevel(log_level)
logger 是否应该是 default_logger?
```
## 149.为什么"今天上午"的标注是nr ？
Jul 5, 2017
```text
 hi,
我在jieba.dict.utf8这个字典里面等26556行发现下面的内容
今天上午 3 nr
今天下午 3 nr
nr不是人名的标注吗？ 为什么这俩的标注是nr？
```
## 150.如何给筛选出来的关键词标注词性？
Jul 4, 2017
```text
 某些新词或不常用词拆开和合并的分词质量比较？
未来会不会有这种功能的API，方便定量研究？
```
```text
 
```
```text
 返回值是列表的形式，如何用jieba的词性标注给列表中的每个关键词标注词性。
```
```text
 参数控制 withFlag=True
```
## 151.jieba詞庫来源于何处？
Jun 29, 2017
```text
 Command "/usr/bin/python2 -u -c "import setuptools, tokenize;file='/tmp/pip-build-6jExFt/jieba/setup.py';f=getattr(tokenize, 'open', open)(file);code=f.read().replace('\r\n', '\n');f.close();exec(compile(code, file, 'exec'))" install --record /tmp/pip-xAfrwY-record/install-record.txt --single-version-externally-managed --compile" failed with error code -9 in /tmp/pip-build-6jExFt/jieba/
```
```text
 Can you provide full error message?
```
```text
 请问dict.txt 词库是否都来自搜狗的2006版免费词库? #24
```
## 152.可以使用数据库保存用户自定义词典吗？
Jun 27, 2017
```text
 请问，官方什么时候能出一个PHP版本呢？因为分词可能是对网页传来的信息直接就分词了，而不是一直读数据库的
```
```text
 现在结巴的用户自定义词典是保存成文本文件的，我想问问是否可以增加功能变为使用数据库保存自定义词典呢？这样也可以直接在数据库里面添加新词，维护词典也更方便一些。
```
```text
 much slower when init
```
## 153.mac 下 python3 的命令行分词中，选项中加了分割符选项却不指定的话，程序不会自动结束。
Jun 26, 2017
```text
 Traceback (most recent call last):
File "", line 1, in 
File "/opt/anaconda3/lib/python3.6/site-packages/jieba/init.py", line 119, in wrapped
return fn(*args, **kwargs)
File "/opt/anaconda3/lib/python3.6/site-packages/jieba/init.py", line 307, in load_userdict
word, freq = tup[0], tup[1]
IndexError: list index out of range
```
```text
 如图，结巴加载后不会往下执行，而是停在最后一行。
(py3env1) ➜  nlpTest python -m jieba -d cutTest.txt > cuted2.txt
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/fv/9t0ldhvx03j4ch64rv3j9qh40000gn/T/jieba.cache
Loading model cost 0.855 seconds.
Prefix dict has been built succesfully.
```
## 154.ChineseAnalyzer导入失败
Jun 23, 2017
```text
 我在自定义词典里加入了带有罗马数字的词，如“纽约Ⅰ线”，并且为其赋予很高的词频，但是总是被分隔开。请问我要如何才能强制把“纽约Ⅰ线”这样的自定义词完整切分出来？？？
词典相关内容如下：
纽约Ⅰ线 100000 ns
python版本3.6.2
代码大致如下：
line = "并明确控制纽约Ⅰ线功率"
jieba.load_userdict("dict.txt")
print(jieba.get_FREQ("纽约Ⅰ线"))
print(jieba.get_FREQ("纽约"))
print(jieba.get_FREQ("Ⅰ"))
print(jieba.get_FREQ("线"))
print(jieba.get_FREQ("Ⅰ线"))
words = posseg.cut(line)
for word in words:
print(word)
输出结果如下：
100000
1758
None
7688
None
并/c
明确/ad
控制/updown
纽约/ns
Ⅰ/x
线/n
功率/n
```
```text
 In [39]: from jieba.analyse import ChineseAnalyzer
ImportError                               Traceback (most recent call last)
 in ()
----> 1 from jieba.analyse import ChineseAnalyzer
ImportError: cannot import name ChineseAnalyzer
请问是什么原因？
```
```text
 看看ChineseTokenizer这个类依赖的库是不是没install，比如whoosh这个库，pip install whoosh一下
```
## 155.offset值应该和原输入保持一致
Jun 22, 2017
```text
 某些新词或不常用词拆开和合并的分词质量比较？
未来会不会有这种功能的API，方便定量研究？
```
```text
 
```
```text
 返回值是列表的形式，如何用jieba的词性标注给列表中的每个关键词标注词性。
```
```text
 参数控制 withFlag=True
```
```text
 输入： hello world
输出：
……        start           end
hello      0                5
world     5                10
这个offset位置标识其实忽略了原文中 hello 和 world 中的空格位置。这样在做高亮显示场景时是错误的。
```
## 156.jython can't use jieba
Jun 18, 2017
```text
 wiki里“对Python中文分词模块结巴分词算法过程的理解和分析”跳转到黄网。
```
```text
 哈哈哈哈哈哈哈

2017-11-02 11:37 GMT+08:00 kercker <notifications@github.com>:
…
 wiki里“对Python中文分词模块结巴分词算法过程的理解和分析”跳转到黄网。

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#541>, or mute the thread
 <https://github.com/notifications/unsubscribe-auth/ABD5_fZrOpJSHG1DaFGVx-iT7X4t9A95ks5syTkJgaJpZM4QPIGV>
 .
```
```text
 6666666666老铁网站发出来呀
```
```text
 又骗我点击增加csdn浏览量。。。。。。。。。。。。。。。。。
```
```text
 我看到CRF分词算法的介绍，http://blog.csdn.net/ifengle/article/details/3849852
感觉还行。不知道jieba分词是怎么看CRF的？或者已经用了CRF?
```
```text
 从效果上看，CRF可以有更好的切分效果。百度已经在工程上应用了。
```
```text
 @chenweican 真的吗？感觉crf资料不是很多啊。
```
```text
 import jieba
Traceback (most recent call last):
File "", line 1, in
File "/Users/mengbin/jython2.7.0/Lib/site-packages/jieba/init.py", line 15, in
from ._compat import *
ImportError: No module named _compat
```
## 157.可否在论文中引用您的分词器
Jun 12, 2017
```text
 比如"中断" "满格" 等会莫名奇妙的识别为ns,ns词性是地名词性，这也太假了
```
```text
 No description provided.
```
```text
 我在维护一个加速版的jieba，有issue和pr欢迎
https://github.com/deepcs233/jieba_fast
```
```text
 您好, 我想在我的论文中引用您的分词器, 请问我可以直接引用github的链接吗? 或者您有发表的作品我需要一并引用? 十分感谢!
```
## 158.dict是如何生成的？生成dict时词频统计是如何统计的？
May 23, 2017
```text
 "书品相”，无论用add word，还是大幅提高"品相"，均无法切出”书“和”品相“这一结果。
```
```text
 我想知道jieba里初始的dict是如何生成的？怎么统计的词频？得到dict的过程是从语料中从零开始训练生成，还是说也是借助了新华字典这种初始的字典作为辅助原料？还有就是词频是如何进行统计的？
```
```text
 #393
```
```text
 同问
```
## 159.请问能否在已分出词的基础上进一步发现新词
May 17, 2017
```text
 No description provided.
```
```text
 在Hadoop上怎么加载自定义字典啊
```
```text
 例如：
中共中央/总书记/，/国家/主席/，/中央军委/主席
合并成三个词：中共中央总书记/国家主席/中央军委主席
或者
一带/一路
合并成：一带一路
```
```text
 同问,在关键词抽取这个场景下,太细粒度的分词后抽取到的关键词,不一定是最准确的,这种组合起来的分词更符合关键词提取需求,NLPIR分词后会发现行业新词,因此抽取效果上感觉更好
```
## 160.避免含標點的數字被錯誤分詞
May 15, 2017
```text
 比如遇到本身就有语义或者拼写的错误的中文：咳嗽->咳嗽等等
想从语义和词性上通过字典来实现可以吗？
```
```text
 没有的吧，据我使用的经验来说是没有的。
这里评价体系比较难弄
据我的使用经验而言我们是自己建立了一套转换体系
```
```text
 Hello,
請問應該如何避免句子「我賺了123,244.2元！」被錯誤的分隔成
['我', '賺', '了', '123', ',', '244.2', '元', '！']

謝謝！
```
## 161.“挺不错” 为什么词性是 i
May 4, 2017
```text
 除了函数的替换外，基本上没有改动别处的代码。经过大量测试分词的结果与您的结果相同，开启HMM的普通模式时间缩短了60%左右，关闭HMM的普通模式时间缩短了50%左右。
  请问我可以把这些独立成一个仓库吗，经完整的测试后再request你的readme
```
```text
 我想，你可以把这些c++的程序封装成一个静态 so文件，然后可以增加到这个项目里面，然后py脚本直接调用你的c++so 代码，这样可以提高效率。也可以让这个项目更加的完美。
```
```text
 增加一个 c++的模块在里面，然后，py 脚本函数调用弄一个开关，选择是用py代码还是用c++代码。
```
```text
 *unix下编译.so没啥问题，windows下目前兼容还不好还在调试
```
```text
 windows就不要支持了，你可以弄一个说明，就说，c++代码库只支持linux系统，不对windows提供。
windows就让他们用py代码得了。反正又不跑服务端.
```
```text
 用 except ImportError 这种吧，你先 fork 到自己仓库，作者估计也不怎么管了。
```
```text
 已经封装好，详情见https://github.com/deepcs233/jieba_fast
可使用 pip install jieba_fast安装
```
```text
 安装出错啊，报错少文件。
gcc: 错误：source/jieba_fast_functions_wrap_py3_wrap.c：没有那个文件或目录
gcc: 致命错误：没有输入文件
编译中断。
```
```text
 @george-sq 已经fix，有问题建议到issues里问
```
```text
 您好，当我用 jieba.analyse.extract_tags 方法提取关键词时，当数据量比较大的时候，这个方法运行时间太长了，请问有没有并行加速的措施呢？  @fxsjy
```
```text
 同问。
```
```text
 个人感觉拆分开来更合理。
另外jieba.posseg.cut 可以选择为full mode 或者索引引擎模式吗
```
```text
 同求full mode
```
## 162.可否加入情緒分析功能?
May 4, 2017
```text
 我想先分完词，再基于分词结果做词性标注，请问结巴分词支持吗？我看哈工大的pyltp是支持的（[https://blog.csdn.net/baidu_15113429/article/details/78909666]），但是我不想使用pyltp
```
```text
 這裡好像快倒了 沒啥人氣 慘....
```
```text
 https://blog.csdn.net/enter89/article/details/80619805
```
```text
 參考我修正的項目 #670
```
```text
 每次都这样...
Building prefix dict from /usr/local/lib/python2.7/site-packages/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.704 seconds.
Prefix dict has been built succesfully.
```
```text
 同样遇到加载慢的问题。通常是首次加载，或者长时间空闲后首次加载。
有时会慢至 6、7 秒甚至十几秒。
为什么？
那个 cache 文件很小啊。
```
```text
 可否加入情緒分析功能?
目前有的情緒辭典有:
NTU SD
NTU Sentiment Dictionary
https://docs.google.com/forms/d/e/1FAIpQLSe20EyOE3bp9cKT0gF6R4DodTHOmriIGegkGYa03oHYejhi9g/viewform?c=0&w=1
大连理工大学信息检索研究室
http://ir.dlut.edu.cn/EmotionOntologyDownload
```
```text
 同问.
```
## 163.词性标注可以只用词性的大类么？
May 2, 2017
```text
 需求有些特别：要解决的问题的起点并不是文章。而是经过正确分词后的list。有没有什么办法在不经过jieba分词功能的情况下，直接使用关键词提取的功能（调用关键词提取方法之前，传入自定义的分词结果－list格式）
```
```text
 你是不是问题没描述明白？
你自定义的分词结果是什么东西分出来的？如果是jieba分出来的，那就是可以用自定义词典。
如果是其他方式分出来的，可以A+B。合并列表。
```
```text
 字典是中文分词的核心数据。分词的好坏，往往由字典的正确率和是否与时俱进决定。设想中，应该有这么一个机制，允许Jieba的使用者一起来贡献字典的更新。
1、比如，让Jieba的垂直行业用户，可以把针对他们的行业已经tain过的字典上传上来，分享给大家
2、再比如，Jieba自己的字典，是否可以通过对于当前一段时间互联网的最新热词/短语的流行潮流进行自我更新？
```
```text
 单开一个 repo，专门存放分享上来的字典？
通过 pull request 来分享、更新
不过这样要有人来维护这个 repo，要确保字典质量，要花精力的啊
```
```text
 @anjianshi 众筹词典？
```
```text
 众筹是大家出钱然后由某个人 / 组织负责提供、维护字典吗？不知道好不好操作 :)
```
```text
 导入自定义词库可以满足个人需求。
```
```text
 抽取句子中的序列规则，因为词性标注得太细的原因，很多规则冗余，请问可以修改源码的哪一部分？或者直接有这个功能么？
```
## 164.‘今天下午’被分成了人名
Apr 30, 2017
```text
 如题，多谢！
```
```text
 No description provided.
```
```text
 用正则直接匹配可不可以？
```
```text
 正则是可以，如果词库中包含电影电视名 估计能直接拆分出来
```
```text
 用 jieba.add_word
```
```text
 This simple sentence:
今天下午三点提醒我开会,
when I am using posseg but pos this sentence, got this result:
[pair('今天下午', 'nr'), pair('三点', 'm'), pair('提醒', 'v'), pair('我', 'r'), pair('开会', 'v')]

Obviously, this not we want, 今天下午 should be 't' rather than 'nr'.
Even I am using user_dict to separate 今天 and 下午， no result.
```
```text
 这个jieba词性烂的很，不建议用，应该是训练数据问题
```
```text
 当然跟分词效果也有很大关系，你这个明显是分词不对，应该是今天/下午
```
```text
 真的有这么烂？？
```
## 165.HMM模型参数具体是怎样训练出来的？
Apr 30, 2017
```text
 import jieba

jieba.suggest_freq("{UM}", tune=True)
jieba.lcut("{UM}卡上写的地址就是那个", HMM=False)
# ['{', 'UM', '}', '卡', '上', '写', '的', '地址', '就是', '那个']
jieba.suggest_freq("卡上", tune=True)
jieba.lcut("{UM}卡上写的地址就是那个", HMM=False)
# ['{', 'UM', '}', '卡上', '写', '的', '地址', '就是', '那个']
suggest_freq只能处理全部为中文的情况, 如果希望把特殊符号识别成一个词就会出错.
```
```text
 请问下同义词该怎么处理呢？
如：
北京大学,北大,pku
清华大学,清华,Tsinghua University
```
```text
 具体指的是prob_*.py文件里面的数据
是不是对训练的样本的每个字都分别贴上‘BMES’的标签？然后再统计出三个概率表？
```
```text
 模型的数据是如何生成的？ #7
```
## 166.为什么我用jieba切出来的是一个个的字，而不是词。
Apr 27, 2017
```text
 Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/t8/p9fyd38n58v35j4fxrndsyyw0000gn/T/jieba.cache
Loading model cost 0.357 seconds.
Prefix dict has been built succesfully.
```
```text
 jieba 能用于elasticsearch吗？
```
```text
 我用淘宝一个搜索词的所有商品名字组成一个字符串，交给jieba 切出来的都是一个个的字。用正常的句子试了试也是一样。用的是pip install jieba下载的
```
## 167.自定义词典中单个词语中包含空格和"-"
Apr 26, 2017
```text
 “我很帅”这句话词性标注时，“帅”这个词标注为了形容词a，但是在“我非常帅”中“帅”这个词标注为了nr。词典中的“帅”词性为nr。这是什么原因啊，我刚开始使用词性标注工具，原理还不太懂。麻烦懂的解答一下。
```
```text
 Jieba对中文名的支持不是很好啊,例如：
中国商务部国际贸易谈判副代表俞建华表示，2001年在上海，中方成功举办了APEC第九次领导人非正式会议。时隔13年，中方非常荣幸再次成为APEC东道主。中国政府高度重视这一盛事，目前各项准备工作已全面展开。
分词为
中国/ns 商务部/nt 国际/n 贸易谈判/n 副/b 代表/n 俞/zg 建华/nz 表示/v ，/x 2001/m 年/m 在/p 上海/ns ，/x 中方/f 成功/a 举办/v 了/ul APEC/eng 第九次/m 领导人/n 非正式/b 会议/n 。/x 时隔/n 13/m 年/m ，/x 中方/f 非常/d 荣幸/nr 再次/d 成为/v APEC/eng 东道主/nr 。/x 中国政府/nt 高度重视/l 这/r 一/m 盛事/n ，/x 目前/t 各项/r 准备/v 工作/vn 已/d 全面/n 展开/v 。/x
其中人名“俞建华”被分成两个词了"俞/zg 建华/nz ”,有计划改进吗？
```
```text
 如果只用posseg中的viterbi方法进行分词，“中国商务部国际贸易谈判副代表俞建华表示” 被分得非常完美
中国商务部/nt, 国际/n, 贸易/vn, 谈判/vn, 副/b, 代表/n, 俞建华/nr, 表示/v
```
```text
 为了更好匹配使用了自定义词典，因为有些涉及到了产品型号，如5w-30，需要切成一个词，但是现在切出来还分成3个单词。修改了init中的正则
re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&._-\ ]+)", re.U)，还是不能成功，请问有方法吗？
```
```text
 @jianjianxin
正则表达式写错了吧
```
## 168.自定义词库解析...
Apr 20, 2017
```text
 除了函数的替换外，基本上没有改动别处的代码。经过大量测试分词的结果与您的结果相同，开启HMM的普通模式时间缩短了60%左右，关闭HMM的普通模式时间缩短了50%左右。
  请问我可以把这些独立成一个仓库吗，经完整的测试后再request你的readme
```
```text
 我想，你可以把这些c++的程序封装成一个静态 so文件，然后可以增加到这个项目里面，然后py脚本直接调用你的c++so 代码，这样可以提高效率。也可以让这个项目更加的完美。
```
```text
 增加一个 c++的模块在里面，然后，py 脚本函数调用弄一个开关，选择是用py代码还是用c++代码。
```
```text
 *unix下编译.so没啥问题，windows下目前兼容还不好还在调试
```
```text
 windows就不要支持了，你可以弄一个说明，就说，c++代码库只支持linux系统，不对windows提供。
windows就让他们用py代码得了。反正又不跑服务端.
```
```text
 用 except ImportError 这种吧，你先 fork 到自己仓库，作者估计也不怎么管了。
```
```text
 已经封装好，详情见https://github.com/deepcs233/jieba_fast
可使用 pip install jieba_fast安装
```
```text
 安装出错啊，报错少文件。
gcc: 错误：source/jieba_fast_functions_wrap_py3_wrap.c：没有那个文件或目录
gcc: 致命错误：没有输入文件
编译中断。
```
```text
 @george-sq 已经fix，有问题建议到issues里问
```
```text
 您好，当我用 jieba.analyse.extract_tags 方法提取关键词时，当数据量比较大的时候，这个方法运行时间太长了，请问有没有并行加速的措施呢？  @fxsjy
```
```text
 同问。
```
```text
 个人感觉拆分开来更合理。
另外jieba.posseg.cut 可以选择为full mode 或者索引引擎模式吗
```
```text
 同求full mode
```
```text
 您好,
最近想透过R做断字断词及语意情感分析, 找了不少分词辞典及情感辞典
看到网上另一个词库分享, 但是里面的字段格式不太能理解, 不知道各位是否可以指点一下呢?
Download URL:
http://down.51cto.com/data/269758
档案字段格式长这样:
1		扭在	nz	6ff026e67cc327c2			2	930	1	0	3
2		拟在	nz	3ad73d9dc29b7c54			2	10092	0	0	3
3		捻针	nz	52w76148h1f9cei9			2	308	1	0	3
4		怒发冲冠	nfcg	9jue6c3a96b5eoif			4	9313	1	0	3
5		农副产品	nfcp	adc3aa31df8f47dd			4	7450	1	0	3
6		女房东	nfd	78foi563e45ga896			3	7108	1	0	3
7		暖风机	nfj	bbe96g73c89c3298			3	5116	1	0	3
8		年富力强	nflq	6df5a2e8ba64c9a3			4	13740	1	0	3
9		逆耳忠言	nezy	8h65g473e5e5g52e			4	2285	1	0	3
10		难分难解	nfnj	47a6ce306f3i3d2w			4	7382	1	0	3
11		难分难舍	nfns	7i3eb71865g69aa5			4	6718	1	0	3
12		闹翻天	nft	cbe4d1c47ie345a2			3	2694	1	0	3
13		女服务员	nfwy	a9cc81f8f08fac43			4	12386	1	0	3
14		逆反心理	nfxl	a3i3ba1d2a8ed348			4	6096	1	0	3
15		农副业	nfy	c1969cd63ic682bb			3	5468	1	0	3
16		年复一年	nfyn	fd18eb2b7afbc1ed			4	27804	1	0	3
```
## 169.請問如何判斷斷詞後的詞，是否有出現在字典內?
Apr 20, 2017
```text
 "书品相”，无论用add word，还是大幅提高"品相"，均无法切出”书“和”品相“这一结果。
```
```text
 我想知道jieba里初始的dict是如何生成的？怎么统计的词频？得到dict的过程是从语料中从零开始训练生成，还是说也是借助了新华字典这种初始的字典作为辅助原料？还有就是词频是如何进行统计的？
```
```text
 #393
```
```text
 同问
```
```text
 我想要判斷一篇文章，斷詞後，哪些詞是目前字典中並沒有收錄的   (生詞)
官方說明有提到，Jieba核心程式，預設遇到生詞，會自動用HMM來辨識
而我想擷取出這些生詞
請問有何辦法呢?
感激!
```
## 170.jieba.suggest_freq()中没有这个方法
Apr 15, 2017
```text
 met a problem when using jython+jieba
jython version :jython-standalone 2.7.1
jieba :0.38
Exception in thread "main" java.lang.ExceptionInInitializerError Caused by: Traceback (most recent call last): File "jiebademo.py", line 7, in <module> import jieba File "/usr/local/lib/python2.7/site-packages/jieba/__init__.py", line 16, in <module> from . import finalseg File "/usr/local/lib/python2.7/site-packages/jieba/finalseg/__init__.py", line 30, in <module> start_P, trans_P, emit_P = load_model() File "/usr/local/lib/python2.7/site-packages/jieba/finalseg/__init__.py", line 24, in load_model start_p = pickle.load(get_module_res("finalseg", PROB_START_P)) File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 1378, in load File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 858, in load File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 858, in load File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 966, in load_string ValueError: insecure string pickle
any help??
```
```text
 这个需要你把
/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py
line 966
上下文贴出来看看
if you can show the context of
/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py
line 966
maybe i can help you
```
```text
 请问如何使用新词发现功能？以您举得例子来说，"他来到了网易杭研大厦"，分词后确实出现了“杭研”，但我怎么判断“杭研”是个新词？
```
```text
 同问，怎么样爬取，学习新词？

Regards,
Leon Lee
(86)1382-510-1940
Skype:lt.leon0519
2015-06-18 15:39 GMT+08:00 maybeluo notifications@github.com:

请问如何使用新词发现功能？以您举得例子来说，"他来到了网易杭研大厦"，分词后确实出现了“杭研”，但我怎么判断“杭研”是个新词？
—
Reply to this email directly or view it on GitHub
#273.
```
```text
 【新词】这个概念吧，基本上就是说词典里面没有的就是新词，如果词典里面有，就不是新词。。。。
```
```text
 jieba是否有接口获取发现的新词列表？或者有接口获取已有的词典用来判断哪些是新词？
```
```text
 确实是词典里面没有就是新词，但是在词典很大的情况下一个个搜索分出来的词是不是“新词”也是不现实的。
```
```text
 jieba.suggest_freq()中没有这个方法和add_word()，composer 拉取的结巴分词，为什么？
```
## 171.咋样才能不对网址进行切分
Apr 15, 2017
```text
 我在加载了自定义词典之前enable_parallel，发现分词结果并没有使用词典。如果在加载自定义词典之后再enable_parallel则能按照预期结果分词。
#encoding=utf-8
import sys
import jieba
jieba.enable_parallel(8)
jieba.load_userdict('user.dict')

words = jieba.cut("热水器加热时间太长")
print('/'.join(words))

结果是 “热水器/加热/时间/太/长”
#encoding=utf-8
import sys
import jieba
jieba.load_userdict('user.dict') 
jieba.enable_parallel(8)      # 放在load_userdict后面了

words = jieba.cut("热水器加热时间太长")
print('/'.join(words))

结果是*“热水器/加热时间/太长”*
词典里是有“加热时间”，所以可以猜测第一个代码分词前没有载入自定义词典，可能是因为enable_parallel放load_userdict前面的。这不一致的情况是不是多线程读入词典的bug？
```
```text
 如题
```
```text
 @gausszh ，jieba.cut没有对停用词的处理。 jieba.analyse.ChineseAnalyzer有停用词的处理：https://github.com/fxsjy/jieba/blob/master/jieba/analyse/analyzer.py
```
```text
 只能通过这样的方法加停用词么？如果停用词表过大怎么办呢？可不可以通过自定义词典给超高词频，来削弱停用词以达到过滤的效果？
```
```text
 @Jaybeka 我做了一些尝试，请你参考！
你需要到对应的文件夹里找到analyzer.py并做相应的修改即可，其中stopword.txt是用户定义的停用词表
import codecs

if os.path.exists(r"/Library/Python/2.7/site-packages/jieba/analyse/stopword.txt"):
    print "Using your own stopwords.\n"
    STOP_WORDS = frozenset(( line.rstrip() for line in codecs.open('/Library/Python/2.7/site-packages/jieba/analyse/stopword.txt','r','utf-8') ))
else:
    print "Using jieba's stopwords.\n"
    STOP_WORDS = frozenset(('a', 'an', 'and', 'are', 'as', 'at', 'be', 'by', 'can',
                        'for', 'from', 'have', 'if', 'in', 'is', 'it', 'may',
                        'not', 'of', 'on', 'or', 'tbd', 'that', 'the', 'this',
                        'to', 'us', 'we', 'when', 'will', 'with', 'yet',
                        'you', 'your',u'的',u'了',u'和'))
```
```text
 @fxsjy 我在实验效果的时候看到这样的情况：

我想应该是analyzer此时使用了搜索引擎模式，我最近在做文本分析，想要把结果中“人民大会堂”去掉（即不含“人民”）；请问需要如何做呢？
```
```text
 @zihaolucky
谢谢你的建议，我现在在研究关键词提取的问题，而原作者@fxsjy 好像没有把analyzer和analyse的停词表统一，所以我对analyse/init.py进行了修改，加了支持导入自定义停词表和自定义语料库（用于得到tf-idf），修改后的文件如下
https://github.com/Jaybeka/jieba/blob/master/jieba/analyse/__init__.py
因为python新手，可能写得比较复杂，还请多多指教。‘
但我现在遇到的问题是用微博爬取的数据，关注了什么值得买后，关键词前N个一定是与之相关的。
我不知道如果调低其tf-idf，是算法问题？还是语料库不够大？
测试程序如下，备选文本还有log.txt，进行相应替换就好了
https://github.com/Jaybeka/jieba/blob/master/test/test_stopword_tfidf.py
```
```text
 @Jaybeka 关键词提取的问题要请 @fxsjy 同学来解答了。如果你可以把例子截图放上来应该就更清楚了^^;
```
```text
 @fxsjy ，如何升级到最新版的jieba
```
```text
 @just4thu , 最简单的方法是 pip install -U jieba
```
```text
 @fxsjy ，thx very much
```
```text
 @fxsjy 請問如何在斷詞之後使用停用詞處理??
```
```text
 修改https://github.com/fxsjy/jieba/blob/master/jieba/analyse/analyzer.py这个链接给出的对应文件中ChineseTokenizer类就可以了，把mode改成default
```
```text
 @gausszh ，jieba.cut没有对停用词的处理。 jieba.analyse.ChineseAnalyzer有停用词的处理：https://github.com/fxsjy/jieba/blob/master/jieba/analyse/analyzer.py

请问为什么cut/cut_for_search 没有停用词的支持呢？我发现用jieba.analyse.extract_tags 抽取关键词虽然支持停用词字典，但原本就会把一些较不重要的词过滤掉，而被过滤掉的词可能对搜索会有用，比如谁的动物的脚谁是最长的返回['最长', '动物']，而'脚' 对于搜索引擎来说其实也是一个关键词。
或者我直接将停用词加载到一个集合中，自己在cut/cut_for_search后过滤掉？
```
```text
 单个字的词结巴分不出的，所以脚字不会出
```
```text
 @ShenDezhou 其实我想问为什么cut/cut_for_search 中不能加入停用词？
```
```text
 @morefreeze jieba做的是中文分词，先用(\r\n|\s)正则分出来句子后，再进行前缀匹配或者用HMM从概率分布图上计算最短路径，在最新版本引入的ChineseAnalyzer模块作者用了whoosh模块的StopAnalyzer来剔除分词后出现的停用词。如果你想在分词前剔除停用词，可以在修改re_skip_default和re_skip_cut_all正则，但可能会影响jieba分词的F1Measure结果。
```
```text
 @ShenDezhou 感谢回复！
```
```text
 错误输出例：
>>> list(pseg.cut('你们啊，naïve！'))
[pair('你们', 'r'), pair('啊', 'zg'), pair('，', 'x'), pair('na', 'eng'), pair('ï', 'x'), pair('ve', 'eng'), pair('！', 'x')]
>>> list(pseg.cut('傻“bī”的“bī”，究竟是什么字？'))
[pair('傻', 'a'), pair('“', 'x'), pair('b', 'x'), pair('ī', 'x'), pair('”', 'x'), pair('的', 'uj'), pair('“', 'x'), pair('b', 'x'), pair('ī', 'x'), pair('”', 'x'), pair('，', 'x'), pair('究竟', 'd'), pair('是', 'v'), pair('什么', 'r'), pair('字', 'n'), pair('？', 'x')]
要对范围有个大概感觉的话，可以看看 https://en.wikipedia.org/wiki/Latin_script_in_Unicode 里面长得不像符号的东西。
```
```text
 目前结巴把 www.baidu.com    等url切分成了['www', '.' ,'baidu', 'com']但是我不想让它切开  该咋么处理。
还有my_name等样带下划线，也不想让它切开。该咋么做。
```
## 172.import jieba.analyse 失败
Apr 12, 2017
```text
 我想先分完词，再基于分词结果做词性标注，请问结巴分词支持吗？我看哈工大的pyltp是支持的（[https://blog.csdn.net/baidu_15113429/article/details/78909666]），但是我不想使用pyltp
```
```text
 這裡好像快倒了 沒啥人氣 慘....
```
```text
 https://blog.csdn.net/enter89/article/details/80619805
```
```text
 參考我修正的項目 #670
```
```text
 每次都这样...
Building prefix dict from /usr/local/lib/python2.7/site-packages/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.704 seconds.
Prefix dict has been built succesfully.
```
```text
 同样遇到加载慢的问题。通常是首次加载，或者长时间空闲后首次加载。
有时会慢至 6、7 秒甚至十几秒。
为什么？
那个 cache 文件很小啊。
```
```text
 可否加入情緒分析功能?
目前有的情緒辭典有:
NTU SD
NTU Sentiment Dictionary
https://docs.google.com/forms/d/e/1FAIpQLSe20EyOE3bp9cKT0gF6R4DodTHOmriIGegkGYa03oHYejhi9g/viewform?c=0&w=1
大连理工大学信息检索研究室
http://ir.dlut.edu.cn/EmotionOntologyDownload
```
```text
 同问.
```
```text
 import jieba.analyse 失败 提示：ImportError: No module named 'jieba.analyse'; 'jieba' is not a package
而在命令行中却可以导入，不知道什么原因
环境：python3.5
```
```text
 先确定你自己的文件名不叫 jieba.py
```
```text
 好的 谢谢@gumblex
```
## 173.extract_tags() got an unexpected keyword argument 'allowPOS'
Apr 8, 2017
```text
 “车、门、窗” 这类单字词无法抽取，这个可以加一个可选参数来来允许单字关键字呢？
```
```text
 结巴分词好像有停用字典，我在调用的时候会自动去除停用词么？如果不能，应该怎么调用？
```
```text
 我觉得更应该关心怎么使用结巴的停用词表
```
```text
 @DonMillion @chenweican 似乎已有类似的Issue了，请移步#77
```
```text
 现在只能提取关键词 仍然不能提取频数.....
```
```text
 @fxsjy 楼主可以开放统计频数不？这个接口很重要啊!
```
```text
 使用extract_tags函数时，提示函数没有 allowPOS 参数？？？
> from jieba import analyse
> s = '精确模式，试图将句子最精确地切开，适合文本分析；#此模式为默认模式'
> analyse.extract_tags(s,allowPOS=('n','nv'))
Traceback (most recent call last):
File "", line 1, in 
analyse.extract_tags(s,allowPOS=('n','nv'))
TypeError: extract_tags() got an unexpected keyword argument 'allowPOS'
```
```text
 +1
```
```text
 我也遇见了这个问题,昨天还好好的,可以运行,没有改动,今天就不行了,去掉这个参数就可以运行.
```
```text
 希望再提供一些信息:

在 jupyter notebook遇见此问题,配置信息如下:
-- MaC 10.13.4, Python 3.6.5,
但是在终端中是没有问题,可以运行.
```
```text
 +1
```
## 174.识别URL
Apr 7, 2017
```text
 比如"中断" "满格" 等会莫名奇妙的识别为ns,ns词性是地名词性，这也太假了
```
```text
 No description provided.
```
```text
 我在维护一个加速版的jieba，有issue和pr欢迎
https://github.com/deepcs233/jieba_fast
```
```text
 您好, 我想在我的论文中引用您的分词器, 请问我可以直接引用github的链接吗? 或者您有发表的作品我需要一并引用? 十分感谢!
```
```text
 提交一个小bug
demo的分词结果是：
头发/n 长长/n 了/ul ，/x 变成/v 了长/v 头发/n 。/x
结果好像不是太好。
请问是否能用自定义词库或者什么方便纠正？
```
```text
 您好，发现结巴对URL的识别好像不太好，，，，要是从微博中提取出URL这个有办法么？谢谢
```
```text
 识别URL不是分词应该处理的问题呀，使用正则表达式匹配？
```
## 175.词性标注是否能指定分词模式
Apr 5, 2017
```text
 No description provided.
```
```text
 你没看代码吗？
```
```text
 哎，作者就是不爱写注释，哎。
```
```text
 @ilqxejraha 谢谢您，可能是我表达有误，我是想问词汇库里面的词汇和词频是人工统计的吗？还是通过其它的方法。
```
```text
 统计过来的。最后得到了就是这么一个模型。
```
```text
 你在源码中看到词频的使用了吗？
```
```text
 具体的统计词频，词频的作用可能体现在，一个词存在多个意思。
比如英语中，经常有一个词会有很多个意思。
比如出现一个词，并且，这个词有很多种解释，这时候词频可能会对词意的选择有一定帮助。
具体的其他算法hmm的我还没看。
```
```text
 @KevinDotW 据说是基于人民日报的语料库，我也想知道怎么才能创建自己的词典
```
```text
 词性标注是否可以指定分词模式；
想指定搜索模式
谢谢。
```
## 176.请问可以输出 textrank 抽取得到的每个关键词的score吗
Apr 2, 2017
```text
 jieba.enable_parallel()
jieba.load_userdict("./aux/dict")
如上所示，如果先开启并行，再载入自定义词典，会导致自定义的词典没有产生效果
如果先载入自定义词典再进行并行计算开启，自定义词典就是有效的，不知道这算不算bug
```
```text
 确实如此。
先开启并行，再载入自定义词典。此时，从jieba.dt.get_FREQ和jieba.dt.user_word_tag_tab中都能查到用户词典中的词，说明已导入用户词典中的词，但分词结果中并未产生效果。
```
```text
 这是bug吧？同遇到了，定位了半天问题，求修复啊
```
```text
 No description provided.
```
```text
 keywords = jieba.analyse.textrank(content, topK=20, withWeight=True)
for keyword in keywords:
print keyword[0], keyword[1]
和TF-idf一样的方法，这样就可以了。
```
```text
 python3上面使用这种方法将会得到分开的字，不是score.
```
## 177.像这种有明显语法错误的POS，要怎样修正？
Mar 31, 2017
```text
 比如我有一句话，"2015年度本公司营业利润较2014年度减少-3480.83万元，降幅为-26.95%，主要原因系国内宏观经济增速放缓，下游市场需求不足，市场景气程度有所降低，公司销售收入整体下降。"
我先对结巴做了add_word设置
jieba.add_word("-26.95%")
jieba.add_word(u"-3480.83万元")
然后分词结果是：
2015年度|本|公司|营业利润|较|2014年度|减少|-|3480.83万元|，|降幅|为|-|26.95%|，|主要|原因|系|国内|宏观经济|增速|放缓|，|下游|市场需求|不足|，|市场|景气|程度|有所|降低|，|公司|销售收入|整体|下降|。|
带符号的就不行
```
```text
 遇到同样问题。
加入用户字典    test-test 词 依然会被切分成 test/-/test。
提高词频依然没用，不知道是不是 -  会被固定切出来，不收其他影响。
```
```text
 同样的问题，比如：
jieba.add_word("弗里德里克·泰勒",freq=888888); jieba.add_word("冯·诺伊曼",freq=888888); jieba.suggest_freq(('弗里德里克·泰勒', '冯·诺伊曼'), True); txt = "这个外国人叫弗里德里克·泰勒，是美国前国防部部长。冯·诺伊曼是公认的计算机发明人"; print(" | ".join(jieba.cut(txt, HMM=True, cut_all=False)))
结果还是：
'这个 | 外国人 | 叫 | 弗里德里克 | · | 泰勒 | ， | 是 | 美国 | 前 | 国防部 | 部长 | 。 | 冯 | · | 诺伊曼 | 是 | 公认 | 的 | 计算机 | 发明人'
大家有什么好办法？
```
```text
 有同样的需求，需要保证指定的词绝对不分。不知该怎么做。
```
```text
 代码里面有如下的正则表达式，用来提取中文

re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&._%]+)", re.U)
也就是说对于”弗里德里克·泰勒”中的“·”不认为是中文，而是作为类似逗号和句号来处理。

一个不太好但是可以用的办法就是，修改上面的正则表达式，将“·”加入。其中\xb7就是“·”。

re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&.\xb7%]+)", re.U)

测试结果：
》 这个 | 外国人 | 叫 | 弗里德里克·泰勒 | ， | 是 | 美国 | 前 | 国防部 | 部长 | 。 | 冯·诺伊曼 | 是 | 公认 | 的 | 计算机 | 发明人
```
```text
 小明nr 比p 小红nr 可爱v ，x 但c 都d 没有v 珍珍n 可爱v
```
## 178.no module named jieba
Mar 28, 2017
```text
 http://jiebademo.ap01.aws.af.cm/
not work
```
```text
 求助；
postgresql添加jieba扩展，提示找不到jieba的 module ，我的Postgresql的版本是8.4，看到github上有pg_jieba的项目，要求postgresql的版本都是9.1以上，请问Jieba对Postgresql有版本限制吗？紧急求助，谢谢！
```
## 179.在线演示网站域名错误
Mar 15, 2017
```text
 部分代码：
sentence = getText('三国演义.txt')
keywords = jieba.analyse.extract_tags(sentence, topK=20, withWeight=True, allowPOS=('ns'))
按地名统计前20，出来的结果是：
将军 0.11860080694708297
丞相 0.0921638281251496
主公 0.0715118726461463
军士 0.061001657055284146
商议 0.05922492016471381
云长 0.054936383725351035
军马 0.05441077997191448
大喜 0.05001738396254889
后主 0.04707253654544958
先主 0.04402888433316952
都督 0.04296257804954991
众将 0.04167988860499132
天下 0.03893417374252234
陛下 0.03882608055724891
太守 0.035023324694504726
人马 0.03359843706149093
城上 0.03265662810340311
天子 0.03212681223959663
后人 0.03164288695617565
众官 0.03082541531255808
这是怎么统计出来的地名？完全不对啊！
```
```text
 No description provided.
```
```text
 在线演示还有人在维护吗？
```
## 180.可以处理纯英文的文章么？
Mar 14, 2017
```text
 java版本的分词Ansj和HanLP都支持用自己的语料再训练，为什么结巴分词这么久也不支持？
```
```text
 No description provided.
```
```text
 我只能告诉你这玩意看得懂空格和标点……（
```
```text
 用 https://github.com/moses-smt/mosesdecoder/blob/master/scripts/tokenizer/tokenizer.perl
```
## 181.使用相同的自定义词库，在不同机器上分词结果不同
Mar 13, 2017
```text
 jieba.add_word("石墨烯",100,"nr")
jieba.add_word("凱特琳",100,"nr")
jieba.add_word("莫那娄氏",10,"n")
jieba.del_word("自定义词")
不添加词典，利用add_word和del_word仍然不起作用
先开启并行，不起作用。如果先添加词，再进行并行就可以。。。。还真是bug
jieba.enable_parallel()
```
```text
 #547
```
```text
 Is it possible to use the jieba's tokenizer in TfidfVectorizer? [http://scikit-learn.org/stable/modules/generated/sklearn.feature_extraction.text.TfidfVectorizer.html].
sklearn.feature_extraction.text.TfidfVectorizer asks for a callable tokenizer, and I am wondering which function in jieba can be passed here.
I would like to cluster or classify some Chinese documents.
```
```text
 jieba.cut()
or
jieba.cut_for_search()
jieba_1_分词
```
```text
 tfidf_vectorizer = TfidfVectorizer(tokenizer=jieba.cut, lowercase=False, stop_words=stopwords)
```
```text
 比如“黄焖鸡米饭”，在词典中我添加了“黄焖鸡”，本机分词正确。但在其他机器上，加载我的词典后，分词仍然为“黄焖”，请问是调用方式错了吗？
代码如下：
import jieba
jieba.load_userdict('new.dic')
split_names=jieba.cut(value)
```
```text
 检查一下那台机器的字典是否正确加载了吧。。暂时没遇到过这种情况
```
## 182.哈喽LZ&all，想用结巴做语义分析，有什么好的建议么？
Mar 10, 2017
```text
 No description provided.
```
```text
 首先感谢 @fxsjy 的开源共享精神，让我们能处理中文分词，真的很棒！
那么问题来了，很多人都想用机器分析用户输入的中文语义，然后做后续的应用开发，比如机器人、人工智能、客服bot……
有什么好办法去做语义分析吗
比如，我想知道用户的一句话是一般陈述句还是一般疑问句，还是特殊疑问句，主谓宾是谁
这样才能让机器去检索数据库，才能回答给具体主语还是宾语这个主体
目前google开源了他的SyntaxNet，但是只对英文有效，有谁研究过这个么
```
## 183.jieba 能用于elasticsearch吗？
Mar 9, 2017
```text
 Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/t8/p9fyd38n58v35j4fxrndsyyw0000gn/T/jieba.cache
Loading model cost 0.357 seconds.
Prefix dict has been built succesfully.
```
```text
 jieba 能用于elasticsearch吗？
```
## 184.自定义词库问题
Mar 8, 2017
```text
 如果词语中有空格，应该怎么写词典条目？
```
```text
 修改正则re_han_default，里面加个空格就能把有空格的词条分出来了
```
```text
 看了一下源码，对于用户词典用的解析方法是re_han_default，对于默认词典的解析方法是直接split空格，这样是不是有问题。
```
```text
 我也有这样的需求，不知道该怎么解决？
我的自定义词典：
Criminal Minds
Murdoch Mysteries
Iron Fist
中间带空格的词条都被分成单个的单词了。
```
```text
 我使用set_dictionary自定义词库 但是发现每个词语后面都需要添加词语，词频，词性
在README中介绍到   词频，词性 和词性是可以省略的 检查发现在__init__ 文件77行对每一行进行分词后报错 请问怎么解决有人知道吗
```
```text
 jieba.load_userdict 加载词典不起作用
jieba.add_word 一个一个的加载起作用
```
```text
 @qxde01 我也遇到了这个问题。
```
```text
 @1000sprites
   dict_words=pd.read_csv(dictfile)
   n=dict_words.shape[0]
   for i in range(0,n):
            jieba.add_word(dict_words.word[i],int(dict_words.freq[i]),'nz')
```
```text
 @qxde01 thanks so much.
```
## 185.几处 re_han 的汉字可以向 BMP 外拓展
Mar 5, 2017
```text
 请问，官方什么时候能出一个PHP版本呢？因为分词可能是对网页传来的信息直接就分词了，而不是一直读数据库的
```
```text
 现在结巴的用户自定义词典是保存成文本文件的，我想问问是否可以增加功能变为使用数据库保存自定义词典呢？这样也可以直接在数据库里面添加新词，维护词典也更方便一些。
```
```text
 much slower when init
```
```text
 基本上是各种\U000xxxxx的节奏，范围大概是 https://en.wikipedia.org/wiki/Template:CJK_ideographs_in_Unicode 里面的 Unified 和 Compatibility 那样。
```
## 186."eng" 词类应拓展到包括 ï、ī 在内的诸多拉丁字母
Mar 5, 2017
```text
 我在加载了自定义词典之前enable_parallel，发现分词结果并没有使用词典。如果在加载自定义词典之后再enable_parallel则能按照预期结果分词。
#encoding=utf-8
import sys
import jieba
jieba.enable_parallel(8)
jieba.load_userdict('user.dict')

words = jieba.cut("热水器加热时间太长")
print('/'.join(words))

结果是 “热水器/加热/时间/太/长”
#encoding=utf-8
import sys
import jieba
jieba.load_userdict('user.dict') 
jieba.enable_parallel(8)      # 放在load_userdict后面了

words = jieba.cut("热水器加热时间太长")
print('/'.join(words))

结果是*“热水器/加热时间/太长”*
词典里是有“加热时间”，所以可以猜测第一个代码分词前没有载入自定义词典，可能是因为enable_parallel放load_userdict前面的。这不一致的情况是不是多线程读入词典的bug？
```
```text
 如题
```
```text
 @gausszh ，jieba.cut没有对停用词的处理。 jieba.analyse.ChineseAnalyzer有停用词的处理：https://github.com/fxsjy/jieba/blob/master/jieba/analyse/analyzer.py
```
```text
 只能通过这样的方法加停用词么？如果停用词表过大怎么办呢？可不可以通过自定义词典给超高词频，来削弱停用词以达到过滤的效果？
```
```text
 @Jaybeka 我做了一些尝试，请你参考！
你需要到对应的文件夹里找到analyzer.py并做相应的修改即可，其中stopword.txt是用户定义的停用词表
import codecs

if os.path.exists(r"/Library/Python/2.7/site-packages/jieba/analyse/stopword.txt"):
    print "Using your own stopwords.\n"
    STOP_WORDS = frozenset(( line.rstrip() for line in codecs.open('/Library/Python/2.7/site-packages/jieba/analyse/stopword.txt','r','utf-8') ))
else:
    print "Using jieba's stopwords.\n"
    STOP_WORDS = frozenset(('a', 'an', 'and', 'are', 'as', 'at', 'be', 'by', 'can',
                        'for', 'from', 'have', 'if', 'in', 'is', 'it', 'may',
                        'not', 'of', 'on', 'or', 'tbd', 'that', 'the', 'this',
                        'to', 'us', 'we', 'when', 'will', 'with', 'yet',
                        'you', 'your',u'的',u'了',u'和'))
```
```text
 @fxsjy 我在实验效果的时候看到这样的情况：

我想应该是analyzer此时使用了搜索引擎模式，我最近在做文本分析，想要把结果中“人民大会堂”去掉（即不含“人民”）；请问需要如何做呢？
```
```text
 @zihaolucky
谢谢你的建议，我现在在研究关键词提取的问题，而原作者@fxsjy 好像没有把analyzer和analyse的停词表统一，所以我对analyse/init.py进行了修改，加了支持导入自定义停词表和自定义语料库（用于得到tf-idf），修改后的文件如下
https://github.com/Jaybeka/jieba/blob/master/jieba/analyse/__init__.py
因为python新手，可能写得比较复杂，还请多多指教。‘
但我现在遇到的问题是用微博爬取的数据，关注了什么值得买后，关键词前N个一定是与之相关的。
我不知道如果调低其tf-idf，是算法问题？还是语料库不够大？
测试程序如下，备选文本还有log.txt，进行相应替换就好了
https://github.com/Jaybeka/jieba/blob/master/test/test_stopword_tfidf.py
```
```text
 @Jaybeka 关键词提取的问题要请 @fxsjy 同学来解答了。如果你可以把例子截图放上来应该就更清楚了^^;
```
```text
 @fxsjy ，如何升级到最新版的jieba
```
```text
 @just4thu , 最简单的方法是 pip install -U jieba
```
```text
 @fxsjy ，thx very much
```
```text
 @fxsjy 請問如何在斷詞之後使用停用詞處理??
```
```text
 修改https://github.com/fxsjy/jieba/blob/master/jieba/analyse/analyzer.py这个链接给出的对应文件中ChineseTokenizer类就可以了，把mode改成default
```
```text
 @gausszh ，jieba.cut没有对停用词的处理。 jieba.analyse.ChineseAnalyzer有停用词的处理：https://github.com/fxsjy/jieba/blob/master/jieba/analyse/analyzer.py

请问为什么cut/cut_for_search 没有停用词的支持呢？我发现用jieba.analyse.extract_tags 抽取关键词虽然支持停用词字典，但原本就会把一些较不重要的词过滤掉，而被过滤掉的词可能对搜索会有用，比如谁的动物的脚谁是最长的返回['最长', '动物']，而'脚' 对于搜索引擎来说其实也是一个关键词。
或者我直接将停用词加载到一个集合中，自己在cut/cut_for_search后过滤掉？
```
```text
 单个字的词结巴分不出的，所以脚字不会出
```
```text
 @ShenDezhou 其实我想问为什么cut/cut_for_search 中不能加入停用词？
```
```text
 @morefreeze jieba做的是中文分词，先用(\r\n|\s)正则分出来句子后，再进行前缀匹配或者用HMM从概率分布图上计算最短路径，在最新版本引入的ChineseAnalyzer模块作者用了whoosh模块的StopAnalyzer来剔除分词后出现的停用词。如果你想在分词前剔除停用词，可以在修改re_skip_default和re_skip_cut_all正则，但可能会影响jieba分词的F1Measure结果。
```
```text
 @ShenDezhou 感谢回复！
```
```text
 错误输出例：
>>> list(pseg.cut('你们啊，naïve！'))
[pair('你们', 'r'), pair('啊', 'zg'), pair('，', 'x'), pair('na', 'eng'), pair('ï', 'x'), pair('ve', 'eng'), pair('！', 'x')]
>>> list(pseg.cut('傻“bī”的“bī”，究竟是什么字？'))
[pair('傻', 'a'), pair('“', 'x'), pair('b', 'x'), pair('ī', 'x'), pair('”', 'x'), pair('的', 'uj'), pair('“', 'x'), pair('b', 'x'), pair('ī', 'x'), pair('”', 'x'), pair('，', 'x'), pair('究竟', 'd'), pair('是', 'v'), pair('什么', 'r'), pair('字', 'n'), pair('？', 'x')]
要对范围有个大概感觉的话，可以看看 https://en.wikipedia.org/wiki/Latin_script_in_Unicode 里面长得不像符号的东西。
```
## 187.试试GANS 自行生成分界 如何
Mar 3, 2017
```text
 现在我cut
我想看电视
结果是
我 想 看电视
看了下默认的词库里有 看电视 这个词。。
我想搞成
我 想 看 电视
```
```text
 jieba.add_word("看电视",0)#或者jieba.del_word("看电视")#发现主干这个del_word的bug已经修了
jieba.add_word("看",100,"v")
jieba.add_word("电视",100,"n")
```
```text
 我自定义了一个词典，有一些包含“-”的词，比如 “SK-II”。加载这个词库时候发现只有这类词语没法生效，会被分为 SK，-，II,而不是SK-II。请问这个问题应该如何解决
```
```text
 同问，自定义词典中有5%,但分词会分开，词性标注的会分开 但正常分词不会被分开
```
```text
 找到解决方法了
修改jieba/init.py代码中的re_han_default 正则表达式为如下值：
re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._%\-]+)", re.U)
```
```text
 Can solve this: #692
```
```text
 ganerating advanced nets
```
```text
 generative adversarial networks?
```
## 188.jieba 分词 extract_tags 遗漏分词结果
Mar 1, 2017
```text
 No description provided.
```
```text
 用 del_word('男性') 直接将“男性”这个词删掉？
```
```text
 for i in jieba.analyse.extract_tags('我爱上海的水'):
print i
这个例子里面只能输出上海
我，爱，水，都没了？ 请问为什么呀
```
```text
 seg_list = jieba.cut('我爱上海的水',cut_all=True)
print("Full Mode: " + "/ ".join(seg_list))  # 全模式
```
## 189.结巴分词的准确度
Feb 23, 2017
```text
 I would like to treat !! as one token instead of 2.
jieba.add_word(u'!!')
for x in jieba.cut(u'巨大的恶魔!!'):
    print x
巨大
的
恶魔
!
!
```
```text
 默认的dict.txt中有一个词'一一分', 这个词从来没见过, 是个词吗?
```
```text
 你好，
请问一下，关于结巴的分词功能，有没有进行过什么准确度的测试吗？在哪里可以找到结果？
谢谢
```
## 190.demo地址失效
Feb 15, 2017
```text
 我尝试对恒大相关的文章进行分词，结果错误率奇高，结果如下
'''
按照原计划，凯赛尔在西班牙学习三年后就将回国，但对于志向高远的“凯撒”来说，他更希望凭借未来三年的表现能留在西班牙继续深造，从而拉开自己的职业生涯，“**/随恒大/足校在西班牙学习三年后，我希望能留在这里，并开启自己的职业生涯
保监会再发狠/招恒大/人寿被禁止投资股票
这些/对恒大/概念股的影响有多大
然后在10月底三季报披露散户看/到恒大/**买了4.95%是不是要举牌了，许老板就把货都卖给你们了
'''
我将自定义词典中”恒大“的词频调整到了10000都没有任何变化，尝试jieba.suggest_freq('恒大', True)也没有用；尝试将HMM关掉，结果恒大这个词会被一直分成”恒/大“。
```
```text
 能不能再切之前，直接添加：
jieba.add_word("恒大")
```
```text
 我想把只有在指定的字典有的关键词提出来。
import jieba
jieba.initialize()
jieba.set_dictionary('mydic.txt')
但是出来的词还是包含一堆不在字典中的。请问我该怎么做？谢谢
```
```text
 如果 jieba.analyse.extract_tags 能只返回指定字典里的就更好了
```
```text
 http://jiebademo.ap01.aws.af.cm/ 打不开，console里提示net::ERR_NAME_NOT_RESOLVED
```
## 191.默认的字典词性错误
Feb 8, 2017
```text
 jieba.add_word("石墨烯",100,"nr")
jieba.add_word("凱特琳",100,"nr")
jieba.add_word("莫那娄氏",10,"n")
jieba.del_word("自定义词")
不添加词典，利用add_word和del_word仍然不起作用
先开启并行，不起作用。如果先添加词，再进行并行就可以。。。。还真是bug
jieba.enable_parallel()
```
```text
 #547
```
```text
 Is it possible to use the jieba's tokenizer in TfidfVectorizer? [http://scikit-learn.org/stable/modules/generated/sklearn.feature_extraction.text.TfidfVectorizer.html].
sklearn.feature_extraction.text.TfidfVectorizer asks for a callable tokenizer, and I am wondering which function in jieba can be passed here.
I would like to cluster or classify some Chinese documents.
```
```text
 jieba.cut()
or
jieba.cut_for_search()
jieba_1_分词
```
```text
 tfidf_vectorizer = TfidfVectorizer(tokenizer=jieba.cut, lowercase=False, stop_words=stopwords)
```
```text
 比如“黄焖鸡米饭”，在词典中我添加了“黄焖鸡”，本机分词正确。但在其他机器上，加载我的词典后，分词仍然为“黄焖”，请问是调用方式错了吗？
代码如下：
import jieba
jieba.load_userdict('new.dic')
split_names=jieba.cut(value)
```
```text
 检查一下那台机器的字典是否正确加载了吧。。暂时没遇到过这种情况
```
```text
 你好，我用jieba做一些实体识别，发现很多默认的词性都很奇怪
宣传教育 nr
不知道这个是否有办法能都改正下，或者有啥建议
```
## 192.jieba.lcut()的并行分词问题
Feb 4, 2017
```text
 如果词语中有空格，应该怎么写词典条目？
```
```text
 修改正则re_han_default，里面加个空格就能把有空格的词条分出来了
```
```text
 看了一下源码，对于用户词典用的解析方法是re_han_default，对于默认词典的解析方法是直接split空格，这样是不是有问题。
```
```text
 我也有这样的需求，不知道该怎么解决？
我的自定义词典：
Criminal Minds
Murdoch Mysteries
Iron Fist
中间带空格的词条都被分成单个的单词了。
```
```text
 我使用set_dictionary自定义词库 但是发现每个词语后面都需要添加词语，词频，词性
在README中介绍到   词频，词性 和词性是可以省略的 检查发现在__init__ 文件77行对每一行进行分词后报错 请问怎么解决有人知道吗
```
```text
 jieba.load_userdict 加载词典不起作用
jieba.add_word 一个一个的加载起作用
```
```text
 @qxde01 我也遇到了这个问题。
```
```text
 @1000sprites
   dict_words=pd.read_csv(dictfile)
   n=dict_words.shape[0]
   for i in range(0,n):
            jieba.add_word(dict_words.word[i],int(dict_words.freq[i]),'nz')
```
```text
 @qxde01 thanks so much.
```
```text
 简单查看了一下jieba.enable_parallel()部分的源代码
def enable_parallel(processnum=None):
    """
    Change the module's `cut` and `cut_for_search` functions to the
    parallel version.
    Note that this only works using dt, custom Tokenizer
    instances are not supported.
    """
    global pool, dt, cut, cut_for_search
    from multiprocessing import cpu_count
    if os.name == 'nt':
        raise NotImplementedError(
            "jieba: parallel mode only supports posix system")
    else:
        from multiprocessing import Pool
    dt.check_initialized()
    if processnum is None:
        processnum = cpu_count()
    pool = Pool(processnum)
    cut = _pcut
    cut_for_search = _pcut_for_search
其中只是改写了cut和cut_for_search，请问可不可以理解为即使调用了enable_parallel()，使用lcut类的函数还是单进程运行？
以及，Windows系统下Python同样支持from multiprocessing import Pool，请问为什么一定要限制在os.name != 'nt'下才能使用？
```
```text
 自己回答吧
虽然不知道为什么原作者不在Windows下开启多进程模式，但是自己改写了一个粗略版。
思想就是暴力多进程，而且有可能严重消耗内存，适合数据量大同时主机内存空间较大使用
#-*-coding:utf-8-*- 
import jieba
from multiprocessing import Pool,cpu_count

def cut(sentence):
	if sentence!=None:
		sentence = jieba.lcut(sentence,cut_all=False)
		return [i for i in sentence]
	else :
		return None

if __name__ == '__main__':
        path = raw_input("Enter the path: ")
        f = open(path,'r')
        # 读取数据
	data = f.readlines()
        f.close()
        # 创建进程池
	pool = Pool(cpu_count())
	data = pool.map(cut, data)
	pool.close()
	pool.join()
这里要求data为可迭代对象，返回值为多维列表
```
```text
 自己回答吧
虽然不知道为什么原作者不在Windows下开启多进程模式，但是自己改写了一个粗略版。
思想就是暴力多进程，而且有可能严重消耗内存，适合数据量大同时主机内存空间较大使用
#-*-coding:utf-8-*- 
import jieba
from multiprocessing import Pool,cpu_count

def cut(sentence):
	if sentence!=None:
		sentence = jieba.lcut(sentence,cut_all=False)
		return [i for i in sentence]
	else :
		return None

if __name__ == '__main__':
        path = raw_input("Enter the path: ")
        f = open(path,'r')
        # 读取数据
	data = f.readlines()
        f.close()
        # 创建进程池
	pool = Pool(cpu_count())
	data = pool.map(cut, data)
	pool.close()
	pool.join()
这里要求data为可迭代对象，返回值为多维列表

我测试了下，修改了jieba并行分词的源码，如果在win下运行并行分词，会出现错误，而错误貌似是python-x86/x64下才会出现的问题
```
## 193.运行报错：AttributeError: module 'jieba.finalseg' has no attribute 'cut'
Jan 25, 2017
```text
 孙中山 孙中山 nr
毛泽东 毛泽东 nr
周恩来 周恩来 t
邓小平 邓小平 nr
"周恩来"被标注成时间词
```
```text
 关于jieba.load_userdict(dic)这个加载自定义词典，有如下疑问，烦请解答：

加载了自定义词典会不会覆盖默认的词典
2.如果我有多个txt的自定义词典，可否加载多个，机制是加载的多个并存，还是说后加载的覆盖先加载的？
3.如问题2，如果是后加载的覆盖先加载的，那么我就需要把多个txt先合并再加载，如果是并存的那么我可以运行多次jieba.load_userdict(dic)来加载多个词典，反正会并存。
4.之所以有如上的问题，是因为我从搜狗词库里面下载了很多日常用语secl词典

非常感谢，开发人员能解答下。
```
```text
 從這段code來看，他是去呼叫add_word() 方法把字典裡的字一個一個載入，所以應該不存在覆蓋問題，也應該可以多次載入
https://github.com/fxsjy/jieba/blob/master/jieba/__init__.py#L392
```
```text
 load_userdict最后也是调用的是add_word()方法，不会覆盖
```
```text
 运行报错，在_init_.py 和 finalseg文件里都找不到
File "C:\XXXXXXX\Python\Python35\lib\site-packages\jieba_init_.py", line 265, in __cut_DAG
recognized = finalseg.cut(buf)
AttributeError: module 'jieba.finalseg' has no attribute 'cut'
看了下源码，如下，请大神帮忙了!
else:
if not self.FREQ.get(buf):
recognized = finalseg.cut(buf)
for t in recognized:
yield t
else:
for elem in buf:
yield elem
buf = ''
```
```text
 遇到了同样的问题
```
## 194.繁體與簡體的轉換結果不同
Jan 25, 2017
```text
 rt
```
```text
 为什么简体字文章用jieba.analyse.extract_tags提取关键字是繁体字
```
```text
 我几个月前下了jieba使用，发现抽取关键词时，停用词未做处理，所以在网上找了几个停用词表merge以后，手工做了处理。
现在看已经增加了自定义停用词表功能，不过我个人建议是：
1.需要有一个默认的停用词表，在不附带任何选项的情况下，analyse.extract_tags()应该返回经过默认停用词表处理后的结果。
2.假如有人不愿意要停用词，可以在该函数增加一个选项禁用停用词表。
3.假如有人想用自己的停用词表，和现在一样以增设自定义词典的处理即可。
附件为我目前使用的停用词表
stop_words.txt
```
```text
 这个改动建议和各选项的使用频率有关。90%以上的使用者在关键词抽取时是希望删去停用词的，因此应该作为默认选项使用。而不是让大家自己各显神通去找停用词表。
```
```text
 print(jieba.lcut('伊莎貝拉海灘'))
['伊莎貝拉海灘']
print(jieba.lcut('伊莎贝拉海滩'))
['伊莎贝拉', '海滩']
請問同樣的字串在簡體跟繁體下的拆解結果不同的原因，如果想要拆解繁體得到簡體的結果，該如何調整？
```
```text
 自問自答：掛上 dict.txt.big 之後繁體與簡體的結果就一樣了。
```
## 195.你好 结巴分词模型是如何训练的 提供接口了吗
Jan 24, 2017
```text
 您好；
我看到您代码的时候，您的动态规划求频率最大是反向的（也就是从最后一个字符开始算），我自己写了一下，正向的频率最大（也就是从第一个字符开始算）我跑了跑，发现结果是一样的，请问一下这个可以证明吗？还是我操作失误了，还有就是不是说汉语的重点在后面吗？是因为这个用的反向频率最大吗？谢谢
```
```text
 比如：
s = u'''出租 珠江新城 13楼独立90方 2房2 全配套 月8400元''' cut = jieba.cut(s) print ','.join(cut)
结果是
出租, ,珠江新城, ,13,楼,独立,90,方, ,2,房,2, ,全,配套, ,月,8400,元
有可能把13楼、90方、2房、2房2、月8400元给单独划分出来吗？
```
```text
 可以,但需要你自己单独提供词库

lfol <notifications@github.com> 于2018年9月1日周六 下午4:36写道：
…
 比如：
 s = u'''出租 珠江新城 13楼独立90方 2房2 全配套 月8400元''' cut = jieba.cut(s) print
 ','.join(cut)

 结果是
 出租, ,珠江新城, ,13,楼,独立,90,方, ,2,房,2, ,全,配套, ,月,8400,元

 有可能把13楼、90方、2房、2房2、月8400元给单独划分出来吗？

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#665>, or mute the thread
 <https://github.com/notifications/unsubscribe-auth/AaXy0hBiwe3XSNWR8GFIFhy4w56w2teuks5uWkb_gaJpZM4WWIsV>
 .
```
```text
 from pyhanlp import *

segment = HanLP.newSegment().enableNumberQuantifierRecognize(True)
sentences = [
    "十九元套餐包括什么",
    "九千九百九十九朵玫瑰",
    "壹佰块都不给我",
    "９０１２３４５６７８只蚂蚁",
    "牛奶三〇〇克*2",
    "ChinaJoy“扫黄”细则露胸超2厘米罚款",
]
for sentence in sentences:
    print(segment.seg(sentence))
[十九元/mq, 套餐/n, 包括/v, 什么/ry]
[九千九百九十九朵/mq, 玫瑰/n]
[壹佰块/mq, 都/d, 不/d, 给/p, 我/rr]
[９０１２３４５６７８只/mq, 蚂蚁/n]
[牛奶/nf, 三〇〇克/mq, */w, 2/m]
[ChinaJoy/nx, “/w, 扫黄/vi, ”/w, 细则/n, 露/v, 胸/ng, 超/v, 2厘米/mq, 罚款/vi]
```
```text
 我用THULAC（THU Lexical Analyzer for Chinese）由清华大学自然语言处理与社会人文计算实验室研制推出的一套中文词法分析工具包 训练模型能用在结巴上吗
多谢
```
## 196.年、月、日以及分、秒这种不应该是量词q吗，为什么jieba的dict里面都是标称数词m呢
Jan 18, 2017
```text
 某些新词或不常用词拆开和合并的分词质量比较？
未来会不会有这种功能的API，方便定量研究？
```
```text
 
```
```text
 返回值是列表的形式，如何用jieba的词性标注给列表中的每个关键词标注词性。
```
```text
 参数控制 withFlag=True
```
```text
 输入： hello world
输出：
……        start           end
hello      0                5
world     5                10
这个offset位置标识其实忽略了原文中 hello 和 world 中的空格位置。这样在做高亮显示场景时是错误的。
```
```text
 怪怪
byte-compiling /opt/venv27/lib/python2.7/site-packages/jieba/posseg/prob_emit.py to prob_emit.pyc
Killed
```
```text
 这是一个很大的数据文件，估计内存很小不够编译成 pyc 文件被杀了。你有多少可用内存？
```
```text
 多谢多谢，怪不得，我的机器上只有300多m的可用内存，估计是不够用了
```
```text
 不过你可以直接复制文件进去，现在 300m 应该够载入标准词典运行 jieba 的。这点内存应该也够编译啊，除非是 pypy。
```
```text
 在aws上，小于1G内存都编译不了。说的就是那个老版最小的实例
```
```text
 好像还是这样，我 pip install jieba --log log.txt ，log 最后一句还是：Killed ，我 ram 应该是够得都没怎么用。
```
```text
 No description provided.
```
```text
 支持日期不好
```
## 197.词语是Unicode对词语做文本分析存在的问题
Jan 8, 2017
```text
 原字典自带的dict.txt里面 '阳台' 的词频是 242 , 我在我的自定义字典里, 定义:  阳台栏杆 500  阳台天花 500
但是在分词的时候,依然会把阳台有关的词分词为:  阳台 / 栏杆  阳台/天花
```
```text
 运行demo出现这个。。。
Building prefix dict from /Library/Python/2.7/site-packages/jieba/dict.txt ...
Loading model from cache /var/folders/hv/rnxvykcd4dq98fm44qqh7_sm0000gn/T/jieba.cache
Loading model cost 0.584 seconds.
Prefix dict has been built succesfully.
Traceback (most recent call last):
File "test_userdict.py", line 6, in 
jieba.load_userdict('userdict.txt') #载入字典及字典文件名称
File "/Library/Python/2.7/site-packages/jieba/init.py", line 381, in load_userdict
f.name, lineno, line))
ValueError
```
```text
 用结巴对文本分完词之后，由于词语是以Unicode形式存在，每当我要将其导出为txt或CSV文件，就会存在编码的问题。我想了一个本办法，把打印出来的结果直接复制粘贴，可是对于有几万行的输出，有什么别的办法吗？
```
```text
 编一下码不就好了么？比如：data.encode("utf-8")
```
## 198.请问一个词有多个词性，自定义字典该如何定义
Jan 3, 2017
```text
 比如将XXXX年、XX月视作一个词
或者我需要将书名号内所有内容视作整个词
能不能进行自定义配置？
```
```text
 用字典不就行了。匹配正则，想下这效率得有多慢。
```
```text
 Environment:

jieba v0.39

Code:
import jieba
jieba.add_word("现代汉语文本切分与词性标注规范Ｖ1.0");
seg_list = jieba.cut("北大计算语言学研究所从 1992 年开始进行汉语语料库的多级加工研究。第一步是对原\n" +
                "始语料进行切分和词性标注。1994 年制订了《现代汉语文本切分与词性标注规范Ｖ1.0》。")
print(','.join(seg_list))

output:
北大,计算,语言学,研究所,从, ,1992, ,年,开始,进行,汉语,语料库,的,多级,加工,研究,。,第一步,是,对,原,
,始,语料,进行,切分,和,词性,标注,。,1994, ,年,制订,了,《,现代汉语,文本,切分,与,词性,标注,规范,Ｖ,1.0,》,。
```
```text
 根据个人需要修改jieba包init.py中几个正则表达式，使其支持数字特殊字符。re_han_default = re.compile(“(.+)”, re.U)
```
```text
 from pyhanlp import *

segment = HanLP.newSegment('感知机')
CustomDictionary.insert('现代汉语文本切分与词性标注规范Ｖ1.0')
print(segment.analyze("北大计算语言学研究所从1992年开始进行汉语语料库的多级加工研究。第一步是对原" +
                      "始语料进行切分和词性标注。1994年制订了《现代汉语文本切分与词性标注规范Ｖ1.0》。"))

[北大/j 计算/v 语言学/n 研究所/n]/nt 从/p 1992年/t 开始/v 进行/v 汉语/nz 语料库/n 的/u 多/m 级/q 加工/v 研究/v 。/w 第一/m 步/q 是/v 对/p 原始/a 语/Ng 料/v 进行/v 切分/vn 和/c 词性/n 标注/v 。/w 1994年/t 制订/v 了/u 《/w 现代汉语文本切分与词性标注规范Ｖ1.0/nz 》/w 。/w
```
```text
 nr===========
['耿直', '易怒', '易怒', '鲁莽', '安静', '耿直', '公正', '公允', '明朗', '妙语连珠', '言相告', '赤诚', '赤诚', '金石为开', '子虚', '阳奉阴违', '恩尽义', '向壁虚造', '向壁虚构', '空中楼阁', '阳奉阴违', '殷勤', '和蔼', '和善', '和悦', '过谦', '周全', '和蔼可亲', '关怀备至', '殷勤', '大智若愚', '一谦四益', '纳谏如流', '博采', '刚毅', '麻利', '英勇', '英勇', '耿直', '雷厉风飞', '和蔼', '和悦', '仁慈', '慈善', '慈祥', '和善', '温顺', '温驯', '温雅', '慈爱', '蔼然可亲', '和蔼可亲', '温文', '慈悲为怀', '仁慈', '仁慈', '温柔敦厚', '束', '安静', '安静', '安闲', '安逸', '安祥', '安然', '安神', '静谧', '静默', '若素', '和蔼可亲']
这些词的flag全部都是nr，什么情况？难道只要是中文姓氏开头的词都是人名吗？那准确率太低了吧。。。
```
```text
 查了几个，在dict.txt都是人名。可能是用其他词性工具生成不够准确吧。
```
```text
 中文中有很多词既是名词又是形容词，这种做词性标注时该怎么处理
```
```text
 最近也遇到这样的问题，能一起讨论下解决方案吗？
```
```text
 同问
```
```text
 南京市长江大桥
```
## 199.可以直接对切分语料用baum算法求模型参数吗？
Dec 22, 2016
```text
 作为一个大四学生，刚刚学过python，想看看一些成熟的项目具体代码，透彻地进行分析，不知道jieba适不适合呢？
```
```text
 個人感覺挺好的, http://www.cnblogs.com/zhbzz2007/tag/Natural%20language%20processing/
```
```text
 形容词某些情况下变成了名词了 如下所示
# 新 a
>>> pseg.lcut("新材料装备制造业的循环产业体系")
[pair('新', 'a'), pair('材料', 'n'), pair('装备', 'n'), pair('制造业', 'n'), pair('的', 'uj'), pair('循环', 'vn'), pair('产业', 'n'), pair('体系', 'n')]
# 新 n
>>> pseg.lcut("青海格尔木工业园镁锂新材料产业基地")
[pair('青海', 'ns'), pair('格尔木', 'nr'), pair('工业园', 'n'), pair('镁锂', 'nz'), pair('新', 'n'), pair('材料', 'n'), pair('产业基地', 'n')]


# 大 a
>>> pseg.lcut("大工业自由")
[pair('大', 'a'), pair('工业', 'n'), pair('自由', 'a')]
# 大 n
>>> pseg.lcut("7个国家的大公司中")
[pair('7', 'm'), pair('个', 'm'), pair('国家', 'n'), pair('的', 'uj'), pair('大', 'n'), pair('公司', 'n'), pair('中', 'f')]

有什么办法可以使其保持一致吗？
```
```text
 参考hanlp
```
```text
 可以參考我修正的錯誤 #670
測試完 您上述的語句都正常
(可能跟我修正的錯誤有關係)
#'新', 'a'
>>> pseg.lcut("新材料装备制造业的循环产业体系")
[pair('新', 'a'), pair('材料', 'n'), pair('装备', 'n'), pair('制造业', 'n'), pair('的', 'uj'), pair('循环', 'vn'), pair('产业', 'n'), pair('体系', 'n')]
#新,'n' -> #'新', 'a'
>>> pseg.lcut("青海格尔木工业园镁锂新材料产业基地")
[pair('青海', 'ns'), pair('格尔木', 'nr'), pair('工业园', 'n'), pair('镁', 'n'), pair('锂', 'n'), pair('新', 'a'), pair('材料', 'n'), pair('产业基地', 'n')]

#'大', 'a'
pseg.lcut("大工业自由")
[pair('大', 'a'), pair('工业', 'n'), pair('自由', 'a')]

#'大', 'n' -> #'大', 'a'
pseg.lcut("7个国家的大公司中")
[pair('7', 'm'), pair('个', 'm'), pair('国家', 'n'), pair('的', 'uj'), pair('大', 'a'), pair('公司', 'n'), pair('中', 'f')]
```
```text
 最近研究 中文分词，准备自己做一个，采用双向匹配分词和HMM处理未登录词、削歧义。
不过看了jieba，感觉细节已经做了很多。
准备实现 whoosh 的分词接口，就用在下一个项目中。
不知道能不能提供些 jieba设计方面的资料
```
```text
 可以提供源码 ;)
```
```text
 @terminus318 , 最近工作较忙，还未研究whoosh。 在网上搜索时发现有人对whoosh和jieba做了集成。先mark一下： http://blog.csdn.net/wenxuansoft/article/details/8170714
```
```text
 http://www.verydemo.com/demo_c230_i1903.html
```
```text
 @terminus318 , @oldcai , 结巴0.30版已经添加了用于Whoosh的分词接口：ChineseAnalyzer。
用法：https://github.com/fxsjy/jieba/blob/master/test/test_whoosh.py
```
```text
 哈哈，感谢。
另：工信处女干事好忙，每次测试都要请她过来。
```
```text
 No description provided.
```
## 200.自定义主词库，关闭了HMM，还是会分出英文。无论长短
Dec 20, 2016
```text
 代码如下：
string="这是一个test的行。和一些无意义的dm、jksajdfl"
jieba.set_dictionary('dict.txt.big')#自定义词库
res=jieba.lcut(string,HMM=False)
print(res)
['test', 'dm', 'jksajdfl']
自定义的dict里面只有一个测试的汉字词。
```
```text
 自己解决了。
init.py 破坏掉re_eng = re.compile('[a-zA-Z0-9]', re.U) 就可以了。。。。
比如re_eng = re.compile('aaaaaaaa', re.U)
```
## 201.【分享】好多人需要的：关键词带空格和特殊字符方法~~
Dec 20, 2016
```text
 免费分享，造损免责。
打开默认词典（根目录）或自定义词典，把所有用来间隔词频和词性的空格间隔符改成@@
（选用@@是因为一般关键词里遇到这个分隔符的几率比较小吧）
继续，打开jieba根目录下init.py

搜索
re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._]+)", re.U)
改成
re_han_default = re.compile("(.+)", re.U)



搜索
re_userdict = re.compile('^(.+?)( [0-9]+)?( [a-z]+)?$', re.U)
改成
re_userdict = re.compile('^(.+?)(\u0040\u0040[0-9]+)?(\u0040\u0040[a-z]+)?$', re.U)



搜索
word, freq = line.split(' ')[:2]
改成
word, freq = line.split('\u0040\u0040')[:2]



补充：若用的全模式继续改。
搜索
re_han_cut_all = re.compile("([\u4E00-\u9FD5]+)", re.U)
改成
re_han_cut_all = re.compile("(.+)", re.U)
```
```text
 您好，我正好碰见这个问题，我想把 迈克尔·乔丹 不要拆开，试了您的方法，
将默认的dict改成了如下：
AT&T@@3@@nz
B超@@3@@n
c#@@3@@nz
C#@@3@@nz
c++@@3@@nz
将user_dict变成了如下：
小皇帝
科比布莱恩特
迈克尔·乔丹
init.py也按照您的方法改了，没用全模式所以没改最后一个
加载报错：
ValueError: invalid POS dictionary entry in C:\Tools\anaconda\lib\site-packages\jieba\dict_gai.txt at Line 1: AT&T@@3@@nz
请您指教！感激不尽！
```
```text
 @xusong123 我试了一下，没有问题啊，是不是用用了记事本打开txt导致加上了bom头。建议用notepad++创建编辑文本。
```
```text
 您好，我试了试，还是不行，方便的话我把这个贴上来给您看看
dict_gai3.txt
我加载的是dict_gai3.txt
```
```text
 @xusong123  你试一下，不用我的方法，然后把你字典里的点·删掉，然后看看报错不。如果也报错，就证明你后来新加的词，编码上有问题。
```
```text
 谢谢，已经解决，最好用python3
```
```text
 @xusong123 请问怎么解决的？我也出现同样的错误了，谢谢
```
```text
 我也遇到了这个问题。通常是/tmp/jieba.cache未即时更新引起的。
解决方法：删除jieba.cache，把默认字典（dict.txt）中的空格替换为@@即可。
@wonderfreer
```
## 202.可不可以只输出用户词典里有的词？
Dec 9, 2016
```text
 应该可以通过 添加 用户词典，然后判断词是否在 词典里吧？
但是添加用的是jieba.load_userdict
所以如果我要用if word in 这个dict叫什么呢？。。。。
```
```text
 添加了一个dict但是一直乱码啊，好奇怪啊，utf-8格式，别的用同样方法添加的停用词都不会乱码。一旦我想要匹配  输出词和dict就会输出乱码
```
```text
 用这个，定义主词库，jieba.set_dictionary('dict.txt.big')
自定义的字典，用Notepad++创建，别用win的记事本
```
```text
 这个是添加用户词典里的词，从而可以被识别和区分，但是并不是只输出词典里有的词吧？


Sent from Mail Master
在2016年12月20日 19:15，mycrystalgirl 写道:

用这个，定义主词库，jieba.set_dictionary('dict.txt.big')
自定义的字典，用Notepad++创建，别用win的记事本

—
You are receiving this because you authored the thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 你都已经定义了主词典，当然所有分词结果都是都主词典抽取的，所以只输出主词典有的词。
记住几点
1.记得别用关键词提取，而是用分词功能jieba.cut
2.关闭HMM
即便这样，还是匹配出字符串里任何的英文和数字字符串，所以
3.修改init.py把里面的eng正则破坏掉，我另一个帖子讲的
我也是刚刚用结巴分词，需求跟你的一样，才摸索出的。
```
```text
 你说需要破坏init.py
“”修改init.py把里面的eng正则破坏掉，我另一个帖子讲的“”
能不能把你的帖子地址分享一下。我也遇到了非常类似的问题。谢谢
```
```text
 HMM=false
```
## 203.进行关键词提取之前，能否传入自己的分词结果
Dec 8, 2016
```text
 需求有些特别：要解决的问题的起点并不是文章。而是经过正确分词后的list。有没有什么办法在不经过jieba分词功能的情况下，直接使用关键词提取的功能（调用关键词提取方法之前，传入自定义的分词结果－list格式）
```
```text
 你是不是问题没描述明白？
你自定义的分词结果是什么东西分出来的？如果是jieba分出来的，那就是可以用自定义词典。
如果是其他方式分出来的，可以A+B。合并列表。
```
## 204.关于词语中有空格的情况
Dec 7, 2016
```text
 如果词语中有空格，应该怎么写词典条目？
```
```text
 修改正则re_han_default，里面加个空格就能把有空格的词条分出来了
```
```text
 看了一下源码，对于用户词典用的解析方法是re_han_default，对于默认词典的解析方法是直接split空格，这样是不是有问题。
```
```text
 我也有这样的需求，不知道该怎么解决？
我的自定义词典：
Criminal Minds
Murdoch Mysteries
Iron Fist
中间带空格的词条都被分成单个的单词了。
```
## 205.您好，请问jieba现在有没有支持windows系统的并行模块呢？
Dec 7, 2016
```text
 ：将目标文本按行分隔后，把各行文本分配到多个python进程并行分词，然后归并结果，从而获得分词速度的可观提升
基于python自带的multiprocessing模块，目前暂不支持windows
```
```text
 You can try this. It works in Windows.
from path import Path
from multiprocessing import Pool
import argparse
import time

LINE_PER_CORE = 5000
NUM_CORE = 30
FLOOR_COUNT = 10
CEIL_COUNT = 200

import jieba


def process_one(_in):
    r_list = []
    for l in _in:
        new_l = ' '.join(jieba.cut(l))
        r_list.append(new_l.strip())
    return r_list


def do(l_list, writer):
    pool = Pool(NUM_CORE)
    r_list=pool.map(process_one,[l_list[it:it+LINE_PER_CORE] for it in range(0,len(l_list),LINE_PER_CORE)])
    pool.close()
    pool.join()
    for lr in r_list:
        for line in lr:
            writer.write(line + '\n')

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("-i","--input", help="input folder", default=".")
    parser.add_argument("-o", "--output", help="output folder", default="w_process")
    parser.add_argument("--LINE_PER_CORE", help="# lines per core", type=int, default=20000)
    parser.add_argument("--NUM_CORE", help="# of cores", type=int, default=30)
    parser.add_argument("--coding", type=str, default="utf-8")

    args = parser.parse_args()
    print("Args :", args)
    input_folder = args.input
    output_folder = args.output
    LINE_PER_CORE = args.LINE_PER_CORE
    NUM_CORE = args.NUM_CORE
    coding = args.coding

    if not Path(output_folder).exists():
        Path(output_folder).mkdir()
    for f in Path(input_folder).files('*.txt'):
        print(f.basename(), time.strftime('%Y-%m-%d %X', time.localtime()))
        with open(output_folder + '/%s.output.txt' % (f.namebase,),'w', encoding='utf-8') as f_out:
            with open(f.abspath(),'r', encoding='utf-8') as f_in:
                l_list=[]
                all_dict = {}
                for l in f_in:
                    if len(l_list)<NUM_CORE*LINE_PER_CORE:
                        l_list.append(l)
                    else:
                        do(l_list, f_out)
                        print(f.basename(), time.strftime('%Y-%m-%d %X', time.localtime()))
                        l_list=[]
                if len(l_list)>0:
                    do(l_list, f_out)
    print(time.strftime('%Y-%m-%d %X', time.localtime()))
```
## 206.［请教］如何将类似“abc123456”的文本切分为“abc”和“123456”？
Dec 6, 2016
```text
 如何将类似“abc123456”的文本切分为“abc”和“123456”？默认设置是切分不开的。
```
```text
 给代码里的几处re_eng戳个洞，分出来一类数字看看？
```
```text
 谢谢啦！
```
## 207.自定義詞典生成
Dec 3, 2016
```text
 如果我想自己產生詞典, 我的資料只有
结巴中文分词
這句話
那我的詞典應該是這樣嗎
结 1
巴 1
中 1
文 1
分 1
词 1
结巴 1
巴中 1
中文 1
文分 1
分词 1
结巴中 1
.
.
.
```
## 208.关于finalseg模块中re_han的问题
Nov 29, 2016
```text
 我在py2.7下
re_han = re.compile("([\u4E00-\u9FD5]+)")对中英文混合分词时出现问题
比如输入 中国tfboy说唱,篱笆女人等。
re_han.split结果为[u‘中国’，u'tfbo' ，u'y说唱', ....]
将re_han 改为 re.compile(u"([\u4E00-\u9FD5]+)")能正常分组
```
## 209.词性标注出的词性与词典中不同
Nov 15, 2016
```text
 “我很帅”这句话词性标注时，“帅”这个词标注为了形容词a，但是在“我非常帅”中“帅”这个词标注为了nr。词典中的“帅”词性为nr。这是什么原因啊，我刚开始使用词性标注工具，原理还不太懂。麻烦懂的解答一下。
```
## 210.词性标注
Nov 14, 2016
```text
 孙中山 孙中山 nr
毛泽东 毛泽东 nr
周恩来 周恩来 t
邓小平 邓小平 nr
"周恩来"被标注成时间词
```
## 211.词性 eng 是啥？   为什么官方没有词性对照表？
Nov 10, 2016
```text
 词性 eng 是啥？   为什么官方没有词性对照表？  好纠结， 上网查了资料也找不到eng是啥
```
```text
 http://blog.csdn.net/syani/article/details/52276282
```
```text
 eng是英语的意思。
我也没找到官方的词性对照表。楼上那个词性对照表是很久之前的东西了，现在使用的词性集又加入了一些新的词性。
```
```text
 词性参考ICTCLAS呀：
POS = {
    "n": {  # 1. 名词  (1个一类，7个二类，5个三类)
        "n": "名词",
        "nr": "人名",
        "nr1": "汉语姓氏",
        "nr2": "汉语名字",
        "nrj": "日语人名",
        "nrf": "音译人名",
        "ns": "地名",
        "nsf": "音译地名",
        "nt": "机构团体名",
        "nz": "其它专名",
        "nl": "名词性惯用语",
        "ng": "名词性语素"
    },
    "t": {  # 2. 时间词(1个一类，1个二类)
        "t": "时间词",
        "tg": "时间词性语素"
    },
    "s": {  # 3. 处所词(1个一类)
        "s": "处所词"
    },
    "f": {  # 4. 方位词(1个一类)
        "f": "方位词"
    },
    "v": {  # 5. 动词(1个一类，9个二类)
        "v": "动词",
        "vd": "副动词",
        "vn": "名动词",
        "vshi": "动词“是”",
        "vyou": "动词“有”",
        "vf": "趋向动词",
        "vx": "形式动词",
        "vi": "不及物动词（内动词）",
        "vl": "动词性惯用语",
        "vg": "动词性语素"
    },
    "a": {  # 6. 形容词(1个一类，4个二类)
        "a": "形容词",
        "ad": "副形词",
        "an": "名形词",
        "ag": "形容词性语素",
        "al": "形容词性惯用语"
    },
    "b": {  # 7. 区别词(1个一类，2个二类)
        "b": "区别词",
        "bl": "区别词性惯用语"
    },
    "z": {  # 8. 状态词(1个一类)
        "z": "状态词"
    },
    "r": {  # 9. 代词(1个一类，4个二类，6个三类)
        "r": "代词",
        "rr": "人称代词",
        "rz": "指示代词",
        "rzt": "时间指示代词",
        "rzs": "处所指示代词",
        "rzv": "谓词性指示代词",
        "ry": "疑问代词",
        "ryt": "时间疑问代词",
        "rys": "处所疑问代词",
        "ryv": "谓词性疑问代词",
        "rg": "代词性语素"
    },
    "m": {  # 10. 数词(1个一类，1个二类)
        "m": "数词",
        "mq": "数量词"
    },
    "q": {  # 11. 量词(1个一类，2个二类)
        "q": "量词",
        "qv": "动量词",
        "qt": "时量词"
    },
    "d": {  # 12. 副词(1个一类)
        "d": "副词"
    },
    "p": {  # 13. 介词(1个一类，2个二类)
        "p": "介词",
        "pba": "介词“把”",
        "pbei": "介词“被”"
    },
    "c": {  # 14. 连词(1个一类，1个二类)
        "c": "连词",
        "cc": "并列连词"
    },
    "u": {  # 15. 助词(1个一类，15个二类)
        "u": "助词",
        "uzhe": "着",
        "ule": "了 喽",
        "uguo": "过",
        "ude1": "的 底",
        "ude2": "地",
        "ude3": "得",
        "usuo": "所",
        "udeng": "等 等等 云云",
        "uyy": "一样 一般 似的 般",
        "udh": "的话",
        "uls": "来讲 来说 而言 说来",
        "uzhi": "之",
        "ulian": "连 "  # （“连小学生都会”）
    },
    "e": {  # 16. 叹词(1个一类)
        "e": "叹词"
    },
    "y": {  # 17. 语气词(1个一类)
        "y": "语气词(delete yg)"
    },
    "o": {  # 18. 拟声词(1个一类)
        "o": "拟声词"
    },
    "h": {  # 19. 前缀(1个一类)
        "h": "前缀"
    },
    "k": {  # 20. 后缀(1个一类)
        "k": "后缀"
    },
    "x": {  # 21. 字符串(1个一类，2个二类)
        "x": "字符串",
        "xx": "非语素字",
        "xu": "网址URL"
    },
    "w": {   # 22. 标点符号(1个一类，16个二类)
        "w": "标点符号",
        "wkz": "左括号",  # （ 〔  ［  ｛  《 【  〖 〈   半角：( [ { <
        "wky": "右括号",  # ） 〕  ］ ｝ 》  】 〗 〉 半角： ) ] { >
        "wyz": "全角左引号",  # “ ‘ 『
        "wyy": "全角右引号",  # ” ’ 』
        "wj": "全角句号",  # 。
        "ww": "问号",  # 全角：？ 半角：?
        "wt": "叹号",  # 全角：！ 半角：!
        "wd": "逗号",  # 全角：， 半角：,
        "wf": "分号",  # 全角：； 半角： ;
        "wn": "顿号",  # 全角：、
        "wm": "冒号",  # 全角：： 半角： :
        "ws": "省略号",  # 全角：……  …
        "wp": "破折号",  # 全角：——   －－   ——－   半角：---  ----
        "wb": "百分号千分号",  # 全角：％ ‰   半角：%
        "wh": "单位符号"  # 全角：￥ ＄ ￡  °  ℃  半角：$
    }
}
```
```text
 see: https://gist.github.com/luw2007/6016931
same as: #453
```
## 212.词性冲突
Nov 6, 2016
```text
 当我自定义词典里面的词语的词性与原词典里面相同词语词性冲突时，如何设置让他返回我定义的词性？ 谢谢
```
```text
 请问这个问题有解决吗
```
## 213.結巴自建詞庫中如何找出有空白的英文組合字?
Nov 3, 2016
```text
 請問我們想分詞出來的結果是兩個英文組合字, 在自建詞庫裡有"Apple book", 但分詞出來後還是"Apple"和 "book"分開的兩個字, 請問要如何處理才能達到我們要的結果? 謝謝
```
```text
 我們希望看到的是"Apple book", 兩個字組合成的詞, 請各位大俠開導一下,謝謝 @fxsjy
```
```text
 求大神進來看一下唄... @felixonmars  , 好像都沒人進來看???奇怪???
```
```text
 請問這個討論群是不是關了?都沒有人回覆耶????
```
```text
 这个问题得从底层改（最近没人想来填坑的样子
```
```text
 能透過改這個檔案來達成目的嗎?  init.py
```
```text
 可以
```
```text
 @gumblex, @fxsjy  , 請教大俠如何達成? 如何改? 拜託指導一下, 大大的感謝
```
## 214.如何用代码修改默认词库？
Nov 2, 2016
```text
 现在我cut
我想看电视
结果是
我 想 看电视
看了下默认的词库里有 看电视 这个词。。
我想搞成
我 想 看 电视
```
```text
 jieba.add_word("看电视",0)#或者jieba.del_word("看电视")#发现主干这个del_word的bug已经修了
jieba.add_word("看",100,"v")
jieba.add_word("电视",100,"n")
```
## 215.jieba优化以支持spark-stream高效率分词
Nov 1, 2016
```text
 参考Issue: #387
如果直接在spark-stream的flatMap中执行jieba.cut()分词，则因“发现该dict加载后是以python字典的数据类型存在于spark的driver里面，但是worker进程无法共享这段内存。By @liaicheng ”，导致每次执行时多次“Building prefix dict from the default dictionary ……”，影响了分词性能。
如果“建议用jieba.Tokenizer 得到一个分词器实例，然后再调用相关方法。By @fxsjy ”，则对tokr实例进行广播变量（sc.broadcast(tokr)）时，报“can't pickle thread.lock objects”的错误。
考虑到spark是采用pickle进行持久化，参考pyahocorasick（https://pypi.python.org/pypi/pyahocorasick/ ）的实现方式，
You can also create an eventually large automaton ahead of time and pickle it to re-load later. Here we just pickle to a string. You would typically pickle to a file instead:



import cPickle
pickled = cPickle.dumps(A)
B = cPickle.dumps(pickled)
B.get('he')
(0, 'he')



jieba也能否优化以更好地支持序列化从而进一步提高在spark中的执行效率？
```
## 216.spark中import jieba.analyse失败
Oct 28, 2016
```text
 在我的程序中import jieba成功，并且能够调用分词cut，但是import jieba.analyse失败，无法提取关键词，提示错误为ImportError: No module named jieba.analyse，但是我打印jieba文件夹目录如下，目录中含有analyse。
['finalseg', 'dict.txt', 'init.pyc', 'analyse', '_compat.py', 'main.py', 'posseg', 'init.py', '_compat.pyc']
采用from jieba import analyse的方式能够import成功，但是无法调用extract_tags.@fxsjy
```
```text
 @ShaoyongHua 请问是如何解决的？
```
```text
 我可以啊
```
```text
 我也有这个问题，不过我是from jieba import posseg错误
```
```text
 @ocsponge 请问您的问题解决了吗？
```
```text
 在pyspark程序中，引入jieba进行分词，报no module named jieba, 但是在jupyter notebook中，import  jieba的时候是没有任何问题的，求解？
```
```text
 这是spark的问题，不是jieba的问题，第三方包要用特殊方式载入spark的namespace，具体请查spark的文档。
…
On Sun, Oct 7, 2018, 20:31 guangdi ***@***.***> wrote:
 在pyspark程序中，引入jieba进行分词，报no module named jieba, 但是在jupyter
 notebook中，import jieba的时候是没有任何问题的，求解？

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#406 (comment)>, or mute
 the thread
 <https://github.com/notifications/unsubscribe-auth/AA0SqphEvkPSy1dtZjNnLs6MGxgrQ9eqks5uifQFgaJpZM4Ki_yv>
 .
```
## 217.词典分隔符用空格就不能添加含有空格的词了
Oct 22, 2016
```text
 在使用 load_userdict() 时，并不会主动关闭文件，导致程序警告
ResourceWarning: unclosed file <_io.BufferedReader name='dict.txt'>

不知道没有关闭文件的意义是什么？还有后续操作嘛？
```
```text
 jieba 的词典是用空格作为词，词频，词性之间的分隔符的，但是当要在句子中添加“CH125 Spacy”这种词的时候就会识别不了。
如果词典使用“\t”是不是会相对较好？
```
```text
 @rockyzhengwu, @fxsjy 您好, 請問您有解決嗎?我也想要自訂字典中加有空白的英文組合字，但都識別不出來，　都會拆成好幾個字...
```
## 218.python 无法添加单个词汇
Oct 20, 2016
```text
 需求：
要求匹配出内容的指定的一些关键词（我自己的词库里的词），忽略jieba里千千万万不相关的词，对于我来说，这些词没有用。
实现：
我想的是，把jieba默认的词库替换掉。可是，我看文档里没有提到如何替换jieba的词库
还有什么好的办法实现吗？
谢谢各位
```
```text
 你可以用jieba的自定义字典实现准确识别

中铁建资产管理有限公司 高正  18601064889
…
 在 2018年10月24日，11:11，Bakkan Hwang ***@***.***> 写道：

 需求：
 要求匹配出内容的指定的一些关键词（我自己的词库里的词），忽略jieba里千千万万不相关的词，对于我来说，这些词没有用。

 实现：
 我想的是，把jieba默认的词库替换掉。可是，我看文档里没有提到如何替换jieba的词库
 还有什么好的办法实现吗？

 谢谢各位

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 jieba有实现延时加载机制
jieba 采用延迟加载，import jieba 和 jieba.Tokenizer() 不会立即触发词典的加载，一旦有必要才开始加载词典构建前缀字典。如果你想手工初始 jieba，也可以手动初始化。
import jieba
jieba.initialize()  # 手动初始化（可选）
在 0.28 之前的版本是不能指定主词典的路径的，有了延迟加载机制后，你可以改变主词典的路径:
jieba.set_dictionary('data/dict.txt.big')
例子： https://github.com/fxsjy/jieba/blob/master/test/test_change_dictpath.py
```
```text
 In [4]: print '/'.join(jieba.cut("我来到清华大学",cut_all=True))
我/来到/清华/清华大学/华大/大学
In [5]: jieba.add_word('到清')
In [6]: print '/'.join(jieba.cut("我来到清华大学",cut_all=True))
我/来到/到清/清华/清华大学/华大/大学
In [7]: jieba.add_word('学')
In [8]: print '/'.join(jieba.cut("我来到清华大学",cut_all=True))
我/来到/到清/清华/清华大学/华大/大学
希望高手能给予帮助!
```
```text
 同问
```
```text
 #408
```
## 219.当数据量大时，extract_tags 方法非常慢
Oct 19, 2016
```text
 除了函数的替换外，基本上没有改动别处的代码。经过大量测试分词的结果与您的结果相同，开启HMM的普通模式时间缩短了60%左右，关闭HMM的普通模式时间缩短了50%左右。
  请问我可以把这些独立成一个仓库吗，经完整的测试后再request你的readme
```
```text
 我想，你可以把这些c++的程序封装成一个静态 so文件，然后可以增加到这个项目里面，然后py脚本直接调用你的c++so 代码，这样可以提高效率。也可以让这个项目更加的完美。
```
```text
 增加一个 c++的模块在里面，然后，py 脚本函数调用弄一个开关，选择是用py代码还是用c++代码。
```
```text
 *unix下编译.so没啥问题，windows下目前兼容还不好还在调试
```
```text
 windows就不要支持了，你可以弄一个说明，就说，c++代码库只支持linux系统，不对windows提供。
windows就让他们用py代码得了。反正又不跑服务端.
```
```text
 用 except ImportError 这种吧，你先 fork 到自己仓库，作者估计也不怎么管了。
```
```text
 已经封装好，详情见https://github.com/deepcs233/jieba_fast
可使用 pip install jieba_fast安装
```
```text
 安装出错啊，报错少文件。
gcc: 错误：source/jieba_fast_functions_wrap_py3_wrap.c：没有那个文件或目录
gcc: 致命错误：没有输入文件
编译中断。
```
```text
 @george-sq 已经fix，有问题建议到issues里问
```
```text
 您好，当我用 jieba.analyse.extract_tags 方法提取关键词时，当数据量比较大的时候，这个方法运行时间太长了，请问有没有并行加速的措施呢？  @fxsjy
```
```text
 同问。
```
## 220.在关键字提取中withWeight返回的是权重值，如何返回关键字的词频
Oct 12, 2016
```text
 当我自定义词典里面的词语的词性与原词典里面相同词语词性冲突时，如何设置让他返回我定义的词性？ 谢谢
```
```text
 请问这个问题有解决吗
```
```text
 withWeight返回的是权重值，如何返回关键字的词频，计数？
虽然可以通过先分词，再通过关键词计数，但希望直接在
jieba.analyse.extract_tags(sentence, topK=20, withWeight=False, allowPOS=()) 函数中
返回关键词出现的次数
```
```text
 TF*IDF
```
## 221.cache文件在分词里起了什么作用？可否有无cache的模式
Oct 10, 2016
```text
 用英文试了一下结巴的分词，但是出来的单词很奇怪，很多好像都被切掉了。
如图

想问下可以支持用自己的分词结果去做比如extract嘛？就比如传list而不是text
```
```text
 希望在生产环境用结巴，但生产环境对于获取绝对路径、产生cache文件等操作有沙盒控制。想请问cache文件在jieba中是否有不可替代的作用，我可否修改代码做一个不用cache的版本（略降低效率也可）？
```
## 222.jieba.load_userdict()的问题
Sep 29, 2016
```text
 请问：
结巴分词中用户自定义的字典受默认字典词频的影响是和单个词频的值有关么，例如我自定义词典中增加一个词“汉武帝”，词频为30 ，而默认词典中“汉”和“武帝”的词频都是100，那是不是说我这个词一定会被分成默认字典“汉/武帝”，而不是我自定义词表中的“汉武帝”，如果是这样的话，我后期自定义的字典如果批量加入，我怎么定词频的值合适，还是我需要对一类词进入默认词表中先检查一下，在定义自己的字典，不知道有没有什么好的办法，望回复，感谢！！
```
```text
 我也有这个问题，请问您解决了吗？
```
```text
 我解决了，你在用户自定义词典中不要规定词频，如下
体外药物释放
增效剂
航道整治
肌腱膜纤维肉瘤癌基因同系物A
这样就会自动生成一个合适的词频，从而可以把这个词分出来
```
```text
 @我也碰到了类似的问题，使用楼上的@grandmoi 的方法不写词频和词性会报下面这个错误
IndexError: list index out of range
不写词性，词频写多高都分不出词
难不成真的要自己一个个加到字典里……
```
```text
 @grandmoi 我也有这个问题。我的自定义词典中定义了‘国债’，没有设置词频。然后输入‘中国债’，想分出来‘中’、‘国债’。但是实际输出还是‘中国’、‘债’
```
```text
 我按照规范建立了自己的字典，因为业务需要，我创建了两个字典，分别名字为a.txt ，b.txt 这两个文件分表包括了两个数据表格里面的数据，自己测试了一下不能如果采用以下的方法加载两个文件，好像不能同时生效


import jieba
jieba.load_userdict('a.txt')
jieba.load_userdict('a.txt')


请问结巴能否同时加载两个用户字典呢？？
```
```text
 今天又遇到了问题，我在我的flask web app 中使用了jieba的加载自定义字典功能，然后用下面的命令启动


gunicorn -w 4 -p gevent -b 0.0.0.0:9999 --reload run:app
发现jieba连续不断的吐出下面的提示，我觉得应该是gunicorn开启了多个线程导致了这个问题，我想请教下，该如何解决？


loading model from cache /tmp/jieba.cache
loading model cost 2.44023799896 seconds.
Trie has been built succesfully.
[2016-09-29 17:05:37 +0000] [32528] [INFO] Booting worker with pid: 32528
Building Trie..., from /root/py27/lib/python2.7/site-packages/jieba/dict.txt
loading model from cache /tmp/jieba.cache
loading model cost 2.28571200371 seconds.
Trie has been built succesfully.
[2016-09-29 17:06:06 +0000] [32556] [INFO] Booting worker with pid: 32556
Building Trie..., from /root/py27/lib/python2.7/site-packages/jieba/dict.txt
loading model from cache /tmp/jieba.cache
loading model cost 2.27150511742 seconds.
Trie has been built succesfully.
[2016-09-29 17:06:10 +0000] [32560] [INFO] Booting worker with pid: 32560
Building Trie..., from /root/py27/lib/python2.7/site-packages/jieba/dict.txt
loading model from cache /tmp/jieba.cache
```
```text
 gunicorn会fork多个进程，但是jieba是lazy加载词典的。你可以在import jieba后，调用一下jieba.initialize()。 这样就不会多次加载了。
```
```text
 同样也是jieba load_dict的问题，我发现我自己在词典中添加了一个词并设定了参数比如：萌萌哒 50 a，但是使用posseg分词的结果却是 萌萌哒 x，这是版本问题还是其他设定的问题？
```
```text
 @fxsjy
具体的代码用到了这几个部分
import jieba
jieba.initialize()
import os
if os.path.exists('cbi360.txt'):
jieba.load_userdict('cbi360.txt')
import jieba.posseg as peg
其中 cbi360.txt是我的自己的字典，而且我还用刀了jieba.posseg 的方法，请问这个具体的顺序是怎么样的啊？
```
## 223.，开兄控欧拉5555566555555555，mbvcxzzxz/
Sep 20, 2016
```text
 No description provided.
```
```text
 我想用wiki做自定义词库，已经用gensim训练出一个模型。
不知道用wiki是否比人民日报的效果好？
不知道自定义词库格式要求是怎么样的？
有没有现成的工具可以制作词库？
```
```text
 文档中可以jieba支持自定义加载用户词库的功能，而且自定义的词库的格式也比较简单，我自己是写了一个python脚本来建立自己的词库的。
```
```text
 谢谢 @super1-chen
请问用了user_dict是否会覆盖原来的默认词库？
```
```text
 No description provided.
```
## 224.默认词典有词存疑
Sep 20, 2016
```text
 I would like to treat !! as one token instead of 2.
jieba.add_word(u'!!')
for x in jieba.cut(u'巨大的恶魔!!'):
    print x
巨大
的
恶魔
!
!
```
```text
 默认的dict.txt中有一个词'一一分', 这个词从来没见过, 是个词吗?
```
## 225.wiki自定义词库替代默认词库
Aug 24, 2016
```text
 No description provided.
```
```text
 我想用wiki做自定义词库，已经用gensim训练出一个模型。
不知道用wiki是否比人民日报的效果好？
不知道自定义词库格式要求是怎么样的？
有没有现成的工具可以制作词库？
```
```text
 文档中可以jieba支持自定义加载用户词库的功能，而且自定义的词库的格式也比较简单，我自己是写了一个python脚本来建立自己的词库的。
```
```text
 谢谢 @super1-chen
请问用了user_dict是否会覆盖原来的默认词库？
```
## 226.是否有Ruby版本的计划？
Aug 24, 2016
```text
 如圖所示

我是用簡體中文維基百科的語料，經過處理之後把符號換成 space，再將整個文本當成一個當成一個 string
餵進去。
下面是我的代碼
import time

start_time = time.time()

jieba.enable_parallel(16)
jieba_words = {word for word in jieba.cut_for_search(text_processed, HMM=True) if len(word) > 1 and len(word) <= 7}

print("Segmentizing took {} seconds.".format(time.time() - start_time))

謝謝幫忙！
```
```text
 一直在等待一个Ruby版本的jieba，求大神来开发一个！
```
## 227.如何生成自定义的逆向文件频率（IDF）文本语料库？
Aug 23, 2016
```text
 No description provided.
```
```text
 在基于TF-IDF进行特征提取时，因为文本背景是某一具体行业，不适合使用通用的IDF语料库，我觉得应该使用自定义的基于该行业背景的IDF语料库。请问如何生成自定义IDF语料库呢？
我现在有的数据是几十万个该行业的文档，初步想法是：对每个文档分词去重，把所有文档分词结果汇集去重后形成一个分词集，然后对于分词集里的每一个词语w，按**f=log(该行业文档总数/(含有w的文档数量+1))**公式求出词语w的IDF值f，最后txt文件里每一行放一个(w, f)
是这样吗？谢啦~
```
```text
 还有2个问题：假设通用IDF语料库里有A B C三个词语及其idf值，我自定义IDF语料库里有A B D及其idf值，那么请问，在添加自定义的IDF语料库后：

自定义IDF语料库里的A和B及其相应idf值就直接覆盖通用IDF语料库里的A和B吧？
通用IDF语料库里原先的C及其idf值，现在还有吗？

（其实问题只有就1个：添加自定义IDF语料库后，是整个文件替换，还是说只有那些重复的词语才被替换？）
```
```text
 求助求助求助，没有朋友知道吗？？？
```
```text
 我也想做一个词库,满足自己的需求,自带的字库里有很多类似一一二/一一分/一三六八之类意义不大的词
```
```text
 line 是单个文档
all_dict = {}
for line in lines:
    temp_dict = {}
    total += 1
    cut_line = jieba.cut(line, cut_all=False)
    for word in cut_line:
        temp_dict[word] = 1
    for key in temp_dict:
        num = all_dict.get(key, 0)
        all_dict[key] = num + 1
for key in all_dict:
    w = key.encode('utf-8')
    p = '%.10f' % (math.log10(total/(all_dict[key] + 1)))
```
```text
 @M2shad0w 非常感谢！还有一个问题：
假设通用IDF语料库里有A B C三个词语及其idf值，我自定义的IDF语料库里有A B D及其idf值，那么请问，在添加自定义的IDF语料库后：

自定义IDF语料库里的A和B及其相应idf值就直接覆盖通用IDF语料库里的A和B吧？
通用IDF语料库里原先的C及其idf值，现在还有吗？
```
```text
 @siberiawolf61
我看了一下 结巴库中 load idf path 的代码
https://github.com/fxsjy/jieba/blob/master/jieba/analyse/tfidf.py#L65
class TFIDF(KeywordExtractor):

    def __init__(self, idf_path=None):
        self.tokenizer = jieba.dt
        self.postokenizer = jieba.posseg.dt
        self.stop_words = self.STOP_WORDS.copy()
        self.idf_loader = IDFLoader(idf_path or DEFAULT_IDF)
        self.idf_freq, self.median_idf = self.idf_loader.get_idf()

...

self.idf_loader = IDFLoader(idf_path or DEFAULT_IDF)

应该是覆盖了，c值的 idf 也没有了
```
```text
 @M2shad0w 好的，谢谢啊！
```
```text
 感谢！
```
## 228.简体字文章提取关键字是繁体字
Aug 23, 2016
```text
 rt
```
```text
 为什么简体字文章用jieba.analyse.extract_tags提取关键字是繁体字
```
## 229.自定义词典有怎么支持标点符号或者空格
Aug 19, 2016
```text
 上图是我的数据前一部分，我的目的是对 titles 一列进行分词，分词的代码如下。现在遇到的问题是AttributeError: 'int' object has no attribute 'decode'，所以我认为是 titles 中有 int 所致，所以添加了一个判断条件，但是代码执行的结果依旧是报之前的错。请问这是什么原因？
def jiebait(text):
    seglist = jieba.cut(text, cut_all = True)
    fenci = []
    for word in seglist:
        if (not isinstance(word, int)) and (len(word) >= 2):
            fenci.append(word)
    # 如用搜索引擎模式：
    #seglist = jieba.cut_for_search(text)
    return ' '.join(fenci)
```
```text
 def jiebait(text):
    seglist = [str(w) for w in jieba.cut(text, cut_all = True)]
    fenci = []
    for word in seglist:
        if (not isinstance(word, int)) and (len(word) >= 2):
            fenci.append(word)
    # 如用搜索引擎模式：
    #seglist = jieba.cut_for_search(text)
    return ' '.join(fenci)
```
```text
 AttributeError                            Traceback (most recent call last)
 in ()
2     result_line = ""
3     # segment
----> 4     seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]
5     # remove special character
6     temp = re.sub("[\s+.!_,$%^(+"')]+|[+——()?【】“”！，。？、~@#￥%……&（）]+", "",("/".join(seg_list)))
 in (.0)
2     result_line = ""
3     # segment
----> 4     seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]
5     # remove special character
6     temp = re.sub("[\s+.!_,$%^(+"')]+|[+——()?【】“”！，。？、~@#￥%……&（）]+", "",("/".join(seg_list)))
D:\Anaconda\lib\site-packages\jieba_init_.py in cut(self, sentence, cut_all, HMM)
280             - HMM: Whether to use the Hidden Markov Model.
281         '''
--> 282         sentence = strdecode(sentence)
283
284         if cut_all:
D:\Anaconda\lib\site-packages\jieba_compat.py in strdecode(sentence)
35     if not isinstance(sentence, text_type):
36         try:
---> 37             sentence = sentence.decode('utf-8')
38         except UnicodeDecodeError:
39             sentence = sentence.decode('gbk', 'ignore')
AttributeError: 'int' object has no attribute 'decode'
```
```text
 你的 text 本身是不是 int
```
```text
 是的，本身是int，但是seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]这一步，应该都转成string了
```
```text
 解决了 我这边数据转码的问题 不好意思。。。感谢
```
```text
 如题
```
```text
 同問，請教 @fxsjy
```
```text
 同問，請教 @fxsjy
```
```text
 请问有人解决了吗？
```
```text
 我改了一下，目前支持词库中的符号和空格匹配了 https://github.com/WalkerWang731/jieba
```
## 230.同一文本中相同词汇词性不同？
Aug 18, 2016
```text
 `text = "给我办一个三十的新流量业务"
result = []
for (word, start, end) in jieba.tokenize(text):
pseg_data = [(w, f) for (w, f) in pseg.cut(word)]
print((pseg_data, start, end))
result.append((pseg_data, start, end))
`
([('给', 'p')], 0, 1)
([('我', 'r')], 1, 2)
([('办', 'v')], 2, 3)
([('一个三十', 'm')], 3, 7)
([('的', 'uj')], 7, 8)
([('新', 'a')], 8, 9)
([('流量', 'n')], 9, 11)
([('业务', 'n')], 11, 13)
如果是"两个三十"就可以分开
```
```text
 你好，我在使用结巴分词时，出现同同一词汇词性不同的情况，我使用了动态加词汇，修改dict，导入自己的词汇表，都会出现这种情况，比如我定义：街 1000000000000 ns 我的目标是分割出的“街”都是“ns”，但是在实际中会出现：(m)欢乐(a)街(n)怎么(r)走(v)？  去(v)美食(n)街(ns) 这两种不同的情况，请问是怎么回事呢？
ps:我只是在使用jieba,并不懂里面的理论，如果有基本错误还请原谅。
```
## 231.textrank如何调整滑动窗口?
Aug 17, 2016
```text
 ['a', ' ', 'b', ' ', 'c']
出来空格。。
```
```text
 基于 TextRank 算法的关键词抽取
jieba.analyse.textrank(sentence, topK=20, withWeight=False, allowPOS=('ns', 'n', 'vn', 'v')) 直接使用，接口相同，注意默认过滤词性。
jieba.analyse.TextRank() 新建自定义 TextRank 实例
算法论文： TextRank: Bringing Order into Texts
基本思想:
将待抽取关键词的文本进行分词
以固定窗口大小(默认为5，通过span属性调整)，词之间的共现关系，构建图
计算图中节点的PageRank，注意是无向带权图
我在使用textrank的时候想调整滑动窗口为2,上面说窗口大小默认为5,但找不到span在哪里,求帮助!
```
```text
 搞掂了,原来要新建一个对象,然后改属性,例如:
text = jieba.analyse.TextRank()
text.span = 2
```
## 232.Spark引用jieba问题
Aug 15, 2016
```text
 就是我有一个给定的关键词库，然后新来一篇文档，从词库里面找出几个词语作为这篇文档的关键词。
```
```text
 目前我就在做类似的工作，效果挺不错。
你的这种情况，可以这么做。

先将文档按句子切分成多个句子，然后计算关键词库中的每个关键词在这篇文档中的句子集合；
接着计算关键词库中两两关键词的相似性（可以用Jaccard相似性度量），这样构成了一个相似性矩阵；
接着对相似性矩阵进行特征分解，然后对特征值进行归一化；
对归一化的特征值从大到小排序，并累计求和（cumsum），选取前<=0.8的特征值对应的关键词作为这篇文档的关键词

以上是一个基本的版本，直接用，效果一般。因此，需要考虑关键词的tf-idf。我是这么做的，在上面的步骤3时，对特征值进行tfidf加权。最后实验结果很好。
以上，楼主可以试试。
```
```text
 @MacQing 非常感谢，我试一下。
```
```text
 我在Pyspark中引用jieba，用jieba.load_userdict('xxx.txx')加载了自己的词典，但是发现RDD.map操作jieba的时候，该词典没有加载成功。于是，我定义了一个jieba_instance = jieba.Tokernice()，通过传参的方式将定义个一个jieba实例传递到RDD的Map操作函数里面。操作如下:
ss = rdd_data.map(lambda x:judge_disease(x,column,disease_connect,table_name,jieba_instance ))
judge_disease方法如下：
def judge_disease(x,column,disease_connect,table_name,jieba_instance ):
    arr = []
    if not column:
        return (x,None,table_name)
    seg = jieba_instance .cut(x[column.encode('utf-8')])
    for item in seg:
        item = item.encode('utf-8')
        if disease_connect.has_key(item):
            arr.append(unicode(item,'utf-8'))
    arr = list(set(arr))
    res = (x,"$$".join(arr).split("$$"),table_name)

    return res

报如下错误：
‘UnpicklingError: NEWOBJ class argument has NULL tp_new’
求助：首先，我该如何通过传参的方式，将一个jieba实例传入函数？
另外，有碰到相关问题的其它解决方法吗？
```
```text
 机器运行环境：
jieba + python2.7+spark1.6
在用jieba load user dict的时候，发现该dict加载后是以python 字典的数据类型存在于spark的driver里面，但是worker进程无法共享这段内存。同时，我也发现jieba分词调用的源数据是jieba.cache和FREQ（dictionary）。cache文件当set_dictionary的时候，是生成jieba.xxxxxxxxxxxx.cache，该文件在worker进程也无法访问。所以，spark里面，运用常规方法无法访问jieba自己的词典。
所以，特提供自己的意见，看能否在版本中进行修改。
1.可以直接加载python字典类型的函数，这样，我可以将该字典在spark里面共享，然后在spark的worker进程独立加载。
2.增加一种模式，使得在类似spark这种方式里面，可以直接选择如何生成jieba.cache，比如，可以直接将add_word或者load_userdict新增的词，放到该jieba.cache里面。
以上是我个人的浅见，希望继续交流。
```
```text
 不用传jieba实例，在judge_disease中初始化jieba实例呢？
```
```text
 @snowlord 那相当于每次都得新建jieba实例，对spark的map操作来说，太耗费资源和时间。
```
```text
 请问，楼主解决了吗？怎么解决的呢？ @liaicheng
```
```text
 把默认的dict.txt替换成自己的dict.txt。
```
```text
 我也尝试了这么做，但是dict.txt文件需要 词频 吧？ 我测试的，如果不添加词频，报错。
那么 对于自己词典中的词，没有词频，楼主添加的默认的吗？ 比如 3？ 还是怎么解决的呢？ @liaicheng
```
```text
 没有用过spark，估计是不是由于全局变量引起的一些问题？
建议用jieba.Tokenizer 得到一个分词器实例，然后再调用相关方法。
jieba.Tokenizer(dictionary=DEFAULT_DICT) 新建自定义分词器，可用于同时使用不同词典。jieba.dt 为默认分词器，所有全局分词相关函数都是该分词器的映射。
例子：https://github.com/fxsjy/jieba/blob/0243d568e9421ab7d3c75f49e9adfc230810e0a3/test/test_lock.py
```
```text
 @fxsjy Spark中使用tokr.cut()会报can't pickle thread.lock objects错误
@liaicheng 把默认的dict.txt替换成自己的dict.txt，你的意思是jieba.set_dictionary("dict\dict.txt")
jieba.initialize() 这样？但是我试了还是不行，每次在map中使用jieba时，还是会执行Building prefix dict from the default dictionary ……
```
```text
 请问我是将jieba分词包装到worker上了，这么做的话会有什么问题吗？效率会低吗？？然后用udf函数去解决这个问题
split_precision = udf(lambda value:[token for token in (jieba.cut(value.encode('utf-8'), cut_all=False))],ArrayType(StringType()))
split_result = df.withColumn("tokens"+'_'+column, split_precision(col(column)))
最后我想问一个问题，如果我想增加某个词的权重以让该词能够分出来的话，我在driver的程序里写了jieba.add_word(df_dict.iloc[i][0],freq=10000)
但最后没有将词分出来，猜想是没有将词加载到worker的词库里，所以没有将词分出来。。那么该怎么去做呢？
```
```text
 我也遇到這個問題，我的解法是：

把 /usr/local/lib/python2.7/dist-packages/jieba/ 裡面的 dict.txt 替換成自己的字典
把 /tmp/ 裡面的  jieba.cache 清掉

然後就可以了，參考看看
```
```text
 有解决方案了吗@liaicheng
```
```text
 我把版本为调低到0.36就不报‘UnpicklingError: NEWOBJ class argument has NULL tp_new’ 了
```
```text
 同样遇到pyspark加载用户自定义词典，未生效，有什么解决方案吗？
```
```text
 16年的bug，到现在都没解决么
```
```text
 我也遇到這個問題，我的解法是：

把 /usr/local/lib/python2.7/dist-packages/jieba/ 裡面的 dict.txt 替換成自己的字典
把 /tmp/ 裡面的  jieba.cache 清掉

然後就可以了，參考看看

这个cache需要去每一台机器上清理么？
```
```text
 问题：jieba分词器spark分布式环境中加载词典失效(不会报错，但是不会加载词典分出想要的词，这个也是我用了一年的pyspark+jieba之后才突然发现自己用错了一年)
尝试过的方法：

修改默认词典dict.txt。
失败，测试发现jieba并不加载自己的默认词典！！！
将jieba作为参数传入柯里化的udf。
失败，jieba中有线程锁，无法序列化。
新建一个jieba的Tokenizer实例，作为参数传入udf或者作为全局变量。
失败原因同2
使用.rdd.mapPartiton算子, 在每个partiton中加载词典。
失败原因同2
在udf中通过jieba.dt.initialized判断是否需要加载词典
成功。
原因：通过本地测试（这样executor日志也会痛driver日志一起打印出来），
发现，结巴会在每个partiton中分别初始化一个默认分词器。
jieba采用延迟加载机制，在没有使用jieba去分词或加载词典时，
jieba中的默认分词器，jieba.dt不会初始化。
jieba.dt.initialized属性在初始化后会从False变为True，
所以可以依据这个来判断jieba是否初始化。从而决定是否加载词典。

import jieba
from pyspark.sql import SparkSession, functions as F
from pyspark.sql.types import ArrayType, StringType

spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

df = spark.createDataFrame([
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
], ["id", "content"])
df.show()

jieba.load_userdict(
    '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')


def _wrong_tokenize(content):
    """错误的分词函数。此udf不会加载词典，"不用加水光",
    其中的"水光"会被切开，结果为[不用, 加水, 光, ？]

    :param content: {str} 要分词的句子
    :return: list[word, word, ...]
    """
    return [i for i in jieba.cut(content)]


wrong_tokenize = F.udf(_wrong_tokenize, ArrayType(StringType()))
df.select(wrong_tokenize('content').alias('words')).show()
print('df 的partition 数为：', df.rdd.getNumPartitions())
输出如下，“水光”一词被切开了，说明加载词典失效了。但是惊喜地发现jieba在每个partiton中只初始化了一次。
+---+--------+
| id| content|
+---+--------+
|  0|  不用加水光？|
|  0|  不用加水光？|
|  0|  不用加水光？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
+---+--------+

[Stage 2:>                                                          (0 + 1) / 1]Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Loading model cost 1.588 seconds.
Prefix dict has been built succesfully.
[Stage 3:>                                                          (0 + 3) / 3]Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Loading model cost 5.942 seconds.
Prefix dict has been built succesfully.
[Stage 3:===================>                                       (1 + 2) / 3]Loading model cost 6.045 seconds.
Prefix dict has been built succesfully.
[Stage 3:=======================================>                   (2 + 1) / 3]Loading model cost 6.320 seconds.
Prefix dict has been built succesfully.
+--------------------+
|               words|
+--------------------+
|      [不用, 加水, 光, ？]|
|      [不用, 加水, 光, ？]|
|      [不用, 加水, 光, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|      [单纯, 水, 光, 多少]|
|      [单纯, 水, 光, 多少]|
|      [单纯, 水, 光, 多少]|
+--------------------+
df 的partition 数为： 4
import jieba
from pyspark.sql import SparkSession, functions as F
from pyspark.sql.types import ArrayType, StringType

spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

df = spark.createDataFrame([
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
], ["id", "content"])
df.show()


def _tokenize(content):
    """此udf在每个partiton中加载一次词典，
    "不用加水光", 其中的"水光"会被保留，结果为[不用, 加, 水光, ？]

    :param content: {str} 要分词的内容
    :return: list[word, word, ...]
    """
    if not jieba.dt.initialized:
        # 词典中有"水光"这个词
        jieba.load_userdict(
            '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')
    return [i for i in jieba.cut(content)]


tokenize = F.udf(_tokenize, ArrayType(StringType()))
df.select(tokenize('content').alias('words')).show()
print('df 的partition 数为：', df.rdd.getNumPartitions())

+---+--------+
| id| content|
+---+--------+
|  0|  不用加水光？|
|  0|  不用加水光？|
|  0|  不用加水光？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
+---+--------+

Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
[Stage 2:>                                                          (0 + 1) / 1]Loading model cost 0.804 seconds.
Prefix dict has been built succesfully.
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
[Stage 3:>                                                          (0 + 3) / 3]Loading model cost 1.204 seconds.
Prefix dict has been built succesfully.
Loading model cost 1.200 seconds.
Prefix dict has been built succesfully.
Loading model cost 1.224 seconds.
Prefix dict has been built succesfully.
+------------------+
|             words|
+------------------+
|    [不用, 加, 水光, ？]|
|    [不用, 加, 水光, ？]|
|    [不用, 加, 水光, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|      [单纯, 水光, 多少]|
|      [单纯, 水光, 多少]|
|      [单纯, 水光, 多少]|
+------------------+

df 的partition 数为： 4
@fxsjy @liaicheng 至此，这个issue应该可以关闭了，jieba没有问题，saprk也没有问题，但是spark+jieba就产生了问题。建议作者大大把下面这个udf的写法，添加到官方的readme中，毕竟这个bug不会报错，难以感知，有很多想我一样的小白，错误地使用了一年的pyspark+jieba却没有发现问题。
def _tokenize(content):
    """此udf在每个partiton中加载一次词典"""
    if not jieba.dt.initialized:
        jieba.load_userdict('user_dict.txt')
    return [i for i in jieba.cut(content)]


tokenize = F.udf(_tokenize, ArrayType(StringType()))
```
```text
 This is an awesome answer. You could write a sample program and request a
merge.

Actually this is referring to understanding of spark behavior and your
sample is valuablly helpful.
…
On Wed, Dec 5, 2018, 10:27 Chant00 ***@***.*** wrote:

 问题：jieba分词器spark分布式环境中加载词典失效(不会报错，但是不会加载词典分出想要的词，这个也是我用了一年的pyspark+jieba之后才突然发现自己用错了一年)
 尝试过的方法：

    1. 修改默认词典dict.txt。
    失败，测试发现jieba并不加载自己的默认词典！！！
    2. 将jieba作为参数传入柯里化的udf。
    失败，jieba中有线程锁，无法序列化。
    3. 新建一个jieba的Tokenizer实例，作为参数传入udf或者作为全局变量。
    失败原因同2
    4. 使用.rdd.mapPartiton算子, 在每个partiton中加载词典。
    失败原因同2
    5. 在udf中通过jieba.dt.initialized判断是否需要加载词典
    成功。
    原因：通过本地测试（这样executor日志也会痛driver日志一起打印出来），
    发现，结巴会在每个partiton中分别初始化一个默认分词器。
    jieba采用延迟加载机制，在没有使用jieba去分词或加载词典时，
    jieba中的默认分词器，jieba.dt不会初始化。
    jieba.dt.initialized属性在初始化后会从False变为True，
    所以可以依据这个来判断jieba是否初始化。从而决定是否加载词典。

 import jiebafrom pyspark.sql import SparkSession, functions as Ffrom pyspark.sql.types import ArrayType, StringType

 spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

 df = spark.createDataFrame([
     (0, '不用加水光？'),
     (0, '不用加水光？'),
     (0, '不用加水光？'),
     (1, '加水光多少钱啊？'),
     (1, '加水光多少钱啊？'),
     (1, '加水光多少钱啊？'),
     (2, '单纯水光多少'),
     (2, '单纯水光多少'),
     (2, '单纯水光多少'),
 ], ["id", "content"])
 df.show()

 jieba.load_userdict(
     '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')

 def _wrong_tokenize(content):
     """错误的分词函数。此udf不会加载词典，"不用加水光",    其中的"水光"会被切开，结果为[不用, 加水, 光, ？]    :param content: {str} 要分词的句子    :return: list[word, word, ...]    """
     return [i for i in jieba.cut(content)]


 wrong_tokenize = F.udf(_wrong_tokenize, ArrayType(StringType()))
 df.select(wrong_tokenize('content').alias('words')).show()print('df 的partition 数为：', df.rdd.getNumPartitions())

 输出如下，“水光”一词被切开了，说明加载词典失效了。但是惊喜地发现jieba在每个partiton中只初始化了一次。
 +---+--------+| id| content|
 +---+--------+|  0|  不用加水光？||  0|  不用加水光？||  0|  不用加水光？||  1|加水光多少钱啊？||  1|加水光多少钱啊？||  1|加水光多少钱啊？||  2|  单纯水光多少||  2|  单纯水光多少||  2|  单纯水光多少|
 +---+--------+

 [Stage 2:>                                                          (0 + 1) / 1]Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Loading model cost 1.588 seconds.
 Prefix dict has been built succesfully.
 [Stage 3:>                                                          (0 + 3) / 3]Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Loading model cost 5.942 seconds.
 Prefix dict has been built succesfully.
 [Stage 3:===================>                                       (1 + 2) / 3]Loading model cost 6.045 seconds.
 Prefix dict has been built succesfully.
 [Stage 3:=======================================>                   (2 + 1) / 3]Loading model cost 6.320 seconds.
 Prefix dict has been built succesfully.
 +--------------------+|               words|
 +--------------------+|      [不用, 加水, 光, ？]||      [不用, 加水, 光, ？]||      [不用, 加水, 光, ？]||[加水, 光, 多少, 钱, 啊, ？]||[加水, 光, 多少, 钱, 啊, ？]||[加水, 光, 多少, 钱, 啊, ？]||      [单纯, 水, 光, 多少]||      [单纯, 水, 光, 多少]||      [单纯, 水, 光, 多少]|
 +--------------------+
 df 的partition 数为： 4

 import jiebafrom pyspark.sql import SparkSession, functions as Ffrom pyspark.sql.types import ArrayType, StringType

 spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

 df = spark.createDataFrame([
     (0, '不用加水光？'),
     (0, '不用加水光？'),
     (0, '不用加水光？'),
     (1, '加水光多少钱啊？'),
     (1, '加水光多少钱啊？'),
     (1, '加水光多少钱啊？'),
     (2, '单纯水光多少'),
     (2, '单纯水光多少'),
     (2, '单纯水光多少'),
 ], ["id", "content"])
 df.show()

 def _tokenize(content):
     """此udf在每个partiton中加载一次词典，    "不用加水光", 其中的"水光"会被保留，结果为[不用, 加, 水光, ？]    :param content: {str} 要分词的内容    :return: list[word, word, ...]    """
     if not jieba.dt.initialized:
         # 词典中有"水光"这个词
         jieba.load_userdict(
             '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')
     return [i for i in jieba.cut(content)]


 tokenize = F.udf(_tokenize, ArrayType(StringType()))
 df.select(tokenize('content').alias('words')).show()print('df 的partition 数为：', df.rdd.getNumPartitions())

 +---+--------+| id| content|
 +---+--------+|  0|  不用加水光？||  0|  不用加水光？||  0|  不用加水光？||  1|加水光多少钱啊？||  1|加水光多少钱啊？||  1|加水光多少钱啊？||  2|  单纯水光多少||  2|  单纯水光多少||  2|  单纯水光多少|
 +---+--------+

 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 [Stage 2:>                                                          (0 + 1) / 1]Loading model cost 0.804 seconds.
 Prefix dict has been built succesfully.
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 [Stage 3:>                                                          (0 + 3) / 3]Loading model cost 1.204 seconds.
 Prefix dict has been built succesfully.
 Loading model cost 1.200 seconds.
 Prefix dict has been built succesfully.
 Loading model cost 1.224 seconds.
 Prefix dict has been built succesfully.
 +------------------+|             words|
 +------------------+|    [不用, 加, 水光, ？]||    [不用, 加, 水光, ？]||    [不用, 加, 水光, ？]||[加, 水光, 多少钱, 啊, ？]||[加, 水光, 多少钱, 啊, ？]||[加, 水光, 多少钱, 啊, ？]||      [单纯, 水光, 多少]||      [单纯, 水光, 多少]||      [单纯, 水光, 多少]|
 +------------------+

 df 的partition 数为： 4

 @fxsjy <https://github.com/fxsjy> @liaicheng
 <https://github.com/liaicheng>
 至此，这个issue应该可以关闭了，jieba没有问题，saprk也没有问题，但是spark+jieba就产生了问题。建议作者大大把下面这个udf的写法，添加到官方的readme中，毕竟这个bug不会报错，难以感知，有很多想我一样的小白，错误地使用了一年的pyspark+jieba却没有发现问题。

 def _tokenize(content):
     """此udf在每个partiton中加载一次词典"""
     if not jieba.dt.initialized:
         jieba.load_userdict('user_dict.txt')
     return [i for i in jieba.cut(content)]


 tokenize = F.udf(_tokenize, ArrayType(StringType()))

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#387 (comment)>, or mute
 the thread
 <https://github.com/notifications/unsubscribe-auth/AA0SqpifJt8gBxevt5TSMZ-cvH2UuJz7ks5u1y8PgaJpZM4JkN9I>
 .
```
```text
 问题：jieba分词器spark分布式环境中加载词典失效(不会报错，但是不会加载词典分出想要的词，这个也是我用了一年的pyspark+jieba之后才突然发现自己用错了一年)
尝试过的方法：

修改默认词典dict.txt。
失败，测试发现jieba并不加载自己的默认词典！！！
将jieba作为参数传入柯里化的udf。
失败，jieba中有线程锁，无法序列化。
新建一个jieba的Tokenizer实例，作为参数传入udf或者作为全局变量。
失败原因同2
使用.rdd.mapPartiton算子, 在每个partiton中加载词典。
失败原因同2
在udf中通过jieba.dt.initialized判断是否需要加载词典
成功。
原因：通过本地测试（这样executor日志也会痛driver日志一起打印出来），
发现，结巴会在每个partiton中分别初始化一个默认分词器。
jieba采用延迟加载机制，在没有使用jieba去分词或加载词典时，
jieba中的默认分词器，jieba.dt不会初始化。
jieba.dt.initialized属性在初始化后会从False变为True，
所以可以依据这个来判断jieba是否初始化。从而决定是否加载词典。

import jieba
from pyspark.sql import SparkSession, functions as F
from pyspark.sql.types import ArrayType, StringType

spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

df = spark.createDataFrame([
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
], ["id", "content"])
df.show()

jieba.load_userdict(
    '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')


def _wrong_tokenize(content):
    """错误的分词函数。此udf不会加载词典，"不用加水光",
    其中的"水光"会被切开，结果为[不用, 加水, 光, ？]

    :param content: {str} 要分词的句子
    :return: list[word, word, ...]
    """
    return [i for i in jieba.cut(content)]


wrong_tokenize = F.udf(_wrong_tokenize, ArrayType(StringType()))
df.select(wrong_tokenize('content').alias('words')).show()
print('df 的partition 数为：', df.rdd.getNumPartitions())
输出如下，“水光”一词被切开了，说明加载词典失效了。但是惊喜地发现jieba在每个partiton中只初始化了一次。
+---+--------+
| id| content|
+---+--------+
|  0|  不用加水光？|
|  0|  不用加水光？|
|  0|  不用加水光？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
+---+--------+

[Stage 2:>                                                          (0 + 1) / 1]Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Loading model cost 1.588 seconds.
Prefix dict has been built succesfully.
[Stage 3:>                                                          (0 + 3) / 3]Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Loading model cost 5.942 seconds.
Prefix dict has been built succesfully.
[Stage 3:===================>                                       (1 + 2) / 3]Loading model cost 6.045 seconds.
Prefix dict has been built succesfully.
[Stage 3:=======================================>                   (2 + 1) / 3]Loading model cost 6.320 seconds.
Prefix dict has been built succesfully.
+--------------------+
|               words|
+--------------------+
|      [不用, 加水, 光, ？]|
|      [不用, 加水, 光, ？]|
|      [不用, 加水, 光, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|      [单纯, 水, 光, 多少]|
|      [单纯, 水, 光, 多少]|
|      [单纯, 水, 光, 多少]|
+--------------------+
df 的partition 数为： 4
import jieba
from pyspark.sql import SparkSession, functions as F
from pyspark.sql.types import ArrayType, StringType

spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

df = spark.createDataFrame([
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
], ["id", "content"])
df.show()


def _tokenize(content):
    """此udf在每个partiton中加载一次词典，
    "不用加水光", 其中的"水光"会被保留，结果为[不用, 加, 水光, ？]

    :param content: {str} 要分词的内容
    :return: list[word, word, ...]
    """
    if not jieba.dt.initialized:
        # 词典中有"水光"这个词
        jieba.load_userdict(
            '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')
    return [i for i in jieba.cut(content)]


tokenize = F.udf(_tokenize, ArrayType(StringType()))
df.select(tokenize('content').alias('words')).show()
print('df 的partition 数为：', df.rdd.getNumPartitions())
+---+--------+
| id| content|
+---+--------+
|  0|  不用加水光？|
|  0|  不用加水光？|
|  0|  不用加水光？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
+---+--------+

Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
[Stage 2:>                                                          (0 + 1) / 1]Loading model cost 0.804 seconds.
Prefix dict has been built succesfully.
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
[Stage 3:>                                                          (0 + 3) / 3]Loading model cost 1.204 seconds.
Prefix dict has been built succesfully.
Loading model cost 1.200 seconds.
Prefix dict has been built succesfully.
Loading model cost 1.224 seconds.
Prefix dict has been built succesfully.
+------------------+
|             words|
+------------------+
|    [不用, 加, 水光, ？]|
|    [不用, 加, 水光, ？]|
|    [不用, 加, 水光, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|      [单纯, 水光, 多少]|
|      [单纯, 水光, 多少]|
|      [单纯, 水光, 多少]|
+------------------+

df 的partition 数为： 4
@fxsjy @liaicheng 至此，这个issue应该可以关闭了，jieba没有问题，saprk也没有问题，但是spark+jieba就产生了问题。建议作者大大把下面这个udf的写法，添加到官方的readme中，毕竟这个bug不会报错，难以感知，有很多想我一样的小白，错误地使用了一年的pyspark+jieba却没有发现问题。
def _tokenize(content):
    """此udf在每个partiton中加载一次词典"""
    if not jieba.dt.initialized:
        jieba.load_userdict('user_dict.txt')
    return [i for i in jieba.cut(content)]


tokenize = F.udf(_tokenize, ArrayType(StringType()))

想请问你这个是分布式集群环境做的吗？
```
```text
 @Chant00
用此方法成功了~~感谢！
sc.addPyFile('file_path/userdict.txt') 
def seg_sentence(sentence, stopwords): 
sentence = re.sub("[A-Za-z0-9\!\%\[\]\,\。\@]", "", sentence) 
if not jieba.dt.initialized:
 jieba.load_userdict("userdict.txt") 
sentence_seged = jieba.cut(sentence.strip()) 
outstr = [] 
for word in sentence_seged:
 if word not in stopwords: 
if (word != '\t') or (word != "，"): 
outstr.append(word) 
return outstr
```
## 233.怎样在jieba中用自己领域内的词典？
Aug 11, 2016
```text
 例如：
[笑CRY]  500
将上述整体加入词库中，并进行切分。
目前看，做不到。
是需要做哪些调整吗？词典或代码？
```
```text
 补充：上例中的500是词频。
```
```text
 我改了一下，目前支持词库中的符号和空格匹配了 https://github.com/WalkerWang731/jieba
```
```text
 你好，我很想知道，jieba如何去除标点符号
```
```text
 纯Python实现，无需jieba
punct = set(u''':!),.:;?]}¢'"、。〉》」』】〕〗〞︰︱︳﹐､﹒
﹔﹕﹖﹗﹚﹜﹞！），．：；？｜｝︴︶︸︺︼︾﹀﹂﹄﹏､～￠
々‖•·ˇˉ―--′’”([{£¥'"‵〈《「『【〔〖（［｛￡￥〝︵︷︹︻
︽︿﹁﹃﹙﹛﹝（｛“‘-—_…''')
# 对str/unicode
filterpunt = lambda s: ''.join(filter(lambda x: x not in punct, s))
# 对list
filterpuntl = lambda l: list(filter(lambda x: x not in punct, l))
```
```text
 楼上的回答看上去好酷炫。
```
```text
 吊到无法直视。谢了！
------------------ 原始邮件 ------------------
发件人: "Yanyi Wu";notifications@github.com;
发送时间: 2014年7月19日(星期六) 下午5:46
收件人: "fxsjy/jieba"jieba@noreply.github.com;
抄送: "发如雪止"734133872@qq.com;
主题: Re: [jieba] 关于标点符号 (#169)
楼上的回答看上去好酷炫。
—
Reply to this email directly or view it on GitHub.
```
```text
 我覺得可以幫 jieba 增加一個 jieba.trimPunct(content) 的 method 讓有需要的人可以使用
```
```text
 不好意思，想請問一下類似的問題
由於我目前在處理網頁資料，爬取下來的內容會有一些雜訊，類似
&nbsp;和&gt;

我該如何將其濾掉呢？
```
```text
 不好意思，现在才回复，我觉得你可以先做一遍文本过滤再用jieba分词。
可以先把里面的标点符号过滤掉。
acemoon0301@163.com
发件人： 2153030
发送时间： 2014-12-25 02:06
收件人： fxsjy/jieba
抄送： AcemoonMa
主题： Re: [jieba] 关于标点符号 (#169)
不好意思，想請問一下類似的問題
由於我目前在處理網頁資料，爬取下來的內容會有" "、">" 等雜訊
我該如何將其濾掉呢？
—
Reply to this email directly or view it on GitHub.
```
```text
 @AcemoonMa 不好意思，因為抓取的是html裡的內容，裡面的" "、">"是以字串呈現
空白是 "&+n+b+s+p+;" 將加號都去除，然而用gumblex大的方法會將所有的"&","n","b","s","p",";"都濾掉
昨天我找到了一個方法，可以濾掉字串裡的字串，應該也可以運用在分詞前的濾掉停用字
http://stackoverflow.com/questions/6116978/python-replace-multiple-strings
import re

rep = {"condition1": "", "condition2": "text"} # define desired replacements here

# use these three lines to do the replacement
rep = dict((re.escape(k), v) for k, v in rep.iteritems())
pattern = re.compile("|".join(rep.keys()))
text = pattern.sub(lambda m: rep[re.escape(m.group(0))], text)
For example:
>>> pattern.sub(lambda m: rep[re.escape(m.group(0))], "(condition1) and --condition2--")
'() and --text--'
還沒有實際測試，測試過後跟大家回報
```
```text
 @2153030 HTML 的话最好用 BeautifulSoup 之类的解析库预处理提取字符串（如它的 stripped_strings），再进行分词等自然语言方面的操作。
```
```text
 如果是html解析有很多库都可以提取你需要的text,tag,attribute这些数据啊。
至于网页本身的架构也是可以获取完整的。
例如lxml,beautifulsoap以及python自带的库。
如果你获取后的数据中仍有这些字符，你可以考虑直接写一个字符集合，然后用最基础的循环过滤出来啊。
或者直接用unicode编码过滤，把除了中文，英文，数字以外的都过滤掉就可以了。
acemoon0301@163.com
发件人： 2153030
发送时间： 2014-12-26 15:36
收件人： fxsjy/jieba
抄送： AcemoonMa
主题： Re: [jieba] 关于标点符号 (#169)
@AcemoonMa 不好意思，因為抓取的是html裡的內容，裡面的" "、">"是以字串呈現
空白是 "&+n+b+s+p+;" 將加號都去除，然而用gumblex大的方法會將所有的"&","n","b","s","p",";"都濾掉
昨天我找到了一個方法，可以濾掉字串裡的字串，應該也可以運用在分詞前的濾掉停用字
http://stackoverflow.com/questions/6116978/python-replace-multiple-strings
import re
rep = {"condition1": "", "condition2": "text"} # define desired replacements here
use these three lines to do the replacement
rep = dict((re.escape(k), v) for k, v in rep.iteritems())
pattern = re.compile("|".join(rep.keys()))
text = pattern.sub(lambda m: rep[re.escape(m.group(0))], text)
For example:



pattern.sub(lambda m: rep[re.escape(m.group(0))], "(condition1) and --condition2--")
'() and --text--'
—
Reply to this email directly or view it on GitHub.
```
```text
 你如果有常用的社交或者通讯软件，你可以发软件名和ID，我加你好友，共同探讨，邮件不太方便。
acemoon0301@163.com
发件人： 2153030
发送时间： 2014-12-26 15:36
收件人： fxsjy/jieba
抄送： AcemoonMa
主题： Re: [jieba] 关于标点符号 (#169)
@AcemoonMa 不好意思，因為抓取的是html裡的內容，裡面的" "、">"是以字串呈現
空白是 "&+n+b+s+p+;" 將加號都去除，然而用gumblex大的方法會將所有的"&","n","b","s","p",";"都濾掉
昨天我找到了一個方法，可以濾掉字串裡的字串，應該也可以運用在分詞前的濾掉停用字
http://stackoverflow.com/questions/6116978/python-replace-multiple-strings
import re
rep = {"condition1": "", "condition2": "text"} # define desired replacements here
use these three lines to do the replacement
rep = dict((re.escape(k), v) for k, v in rep.iteritems())
pattern = re.compile("|".join(rep.keys()))
text = pattern.sub(lambda m: rep[re.escape(m.group(0))], text)
For example:



pattern.sub(lambda m: rep[re.escape(m.group(0))], "(condition1) and --condition2--")
'() and --text--'
—
Reply to this email directly or view it on GitHub.
```
```text
 @gumblex @AcemoonMa 謝謝兩位大大的提示
這樣看起來似乎讓實作方便些，我去研究一下BeutifulSoup
@AcemoonMa 郵件已發
```
```text
 楼主 python菜鸟问个简单但急需回答的问题  请问你这个怎么用呢？能写个例子吗？
```
```text
 我想我知道啦  多谢楼主。
```
```text
 import re

def delete_punctuation(text):
    """删除标点符号"""
    text = re.sub(r'[^0-9A-Za-z\u4E00-\u9FFF]+', ' ', text)
    return text

这个满足你的要求吗？
```
```text
 （1）我看有人说“将用户词典覆盖jieba/dict.txt 即可”，有人说“用load.userdict方法添加用户词典”。所以请问是不是上述两种方法都可行？
（2）如果使用后一种方法的话呢，用户词典和jieba默认自带的dict.txt 是同时起作用？还是只有用户词典起作用呢？  “起作用”的意思==“jieba用谁”
（3）如果一个单词，例如“天安门”，同时在dict.txt 和用户词典中出现并且带有不同的词频，那么用词频计算时，是选择哪一个词频呢？不用引起冲突么？
```
## 234.自定义词库不支持特殊符号
Aug 9, 2016
```text
 打扰，请问这个工具用的算法和Stanford Word Segmenter用的算法有什么主要区别吗？
以及，如果使用了这个工具，需要引用你们的论文吗？（或者说这个工具有基于某篇论文吗？）——我也可以直接引用github链接。。。
```
```text
 自定义词库：
*ST舜船
代码：
title_list = jieba.cut("*ST舜船")
print "/ ".join(title_list)
分词结果：
*/ ST/ 舜船
版本0.38    查看Changelog：4. 修复load_userdict加载用户词典不能识别含有空格等特殊字符的问题， by @gumblex;
```
```text
 我也遇到这个问题
```
```text
 同遇到这个问题。
# 结巴自定义词典不支持特殊符号
import jieba
jieba.add_word("<NUM>")
jieba.add_word('<ENG>')
jieba.add_word('<UNK>')
s = "我今天买了<NUM>块钱的<UNK>，<ENG>还不错"
print(jieba.lcut(s))
实际输出
['我', '今天', '买', '了', '<', 'NUM', '>', '块钱', '的', '<', 'UNK', '>', '，', '<', 'ENG', '>', '还', '不错']

期望输出
['我', '今天', '买', '了', '<NUM>', '块钱', '的', '<UNK>', '，', '<ENG'>', '还', '不错']
```
## 235.How about reorganizing the testcases and setup Travis CI?
Aug 8, 2016
```text
 如 我的需求是""号里面的话不分词
```
```text
 Currently the testcases is in a mess. I wondering if it's ok to make a PR to reorganize all the testcases based on pytest and setup TravisCI to enable continuous integration.
```
## 236.繁体中文分词时词性错误
Aug 6, 2016
```text
 把痘痘加入自定词典后，也无法把长痘痘中的长和痘痘分开。各位大佬有什么办法嘛？
```
```text
 jieba.possseg.cut的分词不能支持和jieba.cut一样的可选cut_all参数，这样两种模式下分词结果不一样。
```
```text
 繁体文本分词的时候用的是官方提供的dict.txt.big。可以看到里面“清洁机”和“清潔機”的词频、词性都是相同的。但是切分同一个、仅是简繁体不同的句子时：
“这台清洁机引出一条胶管。”
“這台清潔機引出一條膠管。”
“清洁机”和“清潔機”得出的词性却不同，而且“清潔機”的词性是“x”，即“非语素字”。不止一个词有这样的情况。
```
```text
 我也遇到這個問題，而且是默認字典裡面的詞才會有這樣的問題
如 “攝影機” 在dict.txt.big是n，但是切出來的詞性卻是x。
```
## 237.日期
Aug 2, 2016
```text
 在我的程序中import jieba成功，并且能够调用分词cut，但是import jieba.analyse失败，无法提取关键词，提示错误为ImportError: No module named jieba.analyse，但是我打印jieba文件夹目录如下，目录中含有analyse。
['finalseg', 'dict.txt', 'init.pyc', 'analyse', '_compat.py', 'main.py', 'posseg', 'init.py', '_compat.pyc']
采用from jieba import analyse的方式能够import成功，但是无法调用extract_tags.@fxsjy
```
```text
 @ShaoyongHua 请问是如何解决的？
```
```text
 我可以啊
```
```text
 我也有这个问题，不过我是from jieba import posseg错误
```
```text
 @ocsponge 请问您的问题解决了吗？
```
```text
 在pyspark程序中，引入jieba进行分词，报no module named jieba, 但是在jupyter notebook中，import  jieba的时候是没有任何问题的，求解？
```
```text
 这是spark的问题，不是jieba的问题，第三方包要用特殊方式载入spark的namespace，具体请查spark的文档。
…
On Sun, Oct 7, 2018, 20:31 guangdi ***@***.***> wrote:
 在pyspark程序中，引入jieba进行分词，报no module named jieba, 但是在jupyter
 notebook中，import jieba的时候是没有任何问题的，求解？

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#406 (comment)>, or mute
 the thread
 <https://github.com/notifications/unsubscribe-auth/AA0SqphEvkPSy1dtZjNnLs6MGxgrQ9eqks5uifQFgaJpZM4Ki_yv>
 .
```
```text
 你好，请问下日期；数字这些问题应该怎么解决呢？
```
```text
 不知道你说的是什么数字问题,如果是需要将2017-09-01这样的日期切分成一个单词的话,可以考虑将其这样格式的日期添加到自定义词典中.我目前就是这么处理的效果不错
```
```text
 @aliray ，将这些格式的日期添加到自定义词典中，就可以处理吗？比如加入词库是2017-09-01，想要将2017-10-01分词到一起可能吗？貌似我这边不行啊
```
## 238.同义词
Aug 2, 2016
```text
 当我自定义词典里面的词语的词性与原词典里面相同词语词性冲突时，如何设置让他返回我定义的词性？ 谢谢
```
```text
 请问这个问题有解决吗
```
```text
 withWeight返回的是权重值，如何返回关键字的词频，计数？
虽然可以通过先分词，再通过关键词计数，但希望直接在
jieba.analyse.extract_tags(sentence, topK=20, withWeight=False, allowPOS=()) 函数中
返回关键词出现的次数
```
```text
 TF*IDF
```
```text
 请问下同义词该怎么处理呢？
```
## 239.可以自定义排除某些词性的词吗
Jul 28, 2016
```text
 jieba.add_word("β细胞", freq=True, tag='n')
a = '揭晓！胰岛β细胞功能可以恢复吗?'
WORD = jieba.lcut(a,cut_all=False)
print ([word for word in WORD])
['揭晓', '！', '胰岛', 'β', '细胞', '功能', '可以', '恢复', '吗', '?']
```
```text
 freq的值设置高一点？
```
```text
 freq的值设置高一点？

谢谢，查到了 好像是和特殊字符有关 修改了源代码
```
```text
 jieba是否支持增量导入用户自定义词典？
即在调用load_userdict()之后，再增量地调用某个函数增加用户自定义词？
```
```text
 比如我使用关键词提取的时候不想要动词，有选项可以直接排除吗？
另外，我发现关键词提取会提取到数字，我不想要怎么办？
```
```text
 同问，楼楼有答案吗？
```
```text
 @dandan0503 搞不定呢。
```
## 240.用户自定义词典必须添加词频和词性吗
Jul 26, 2016
```text
 jieba.enable_parallel()
jieba.load_userdict("./aux/dict")
如上所示，如果先开启并行，再载入自定义词典，会导致自定义的词典没有产生效果
如果先载入自定义词典再进行并行计算开启，自定义词典就是有效的，不知道这算不算bug
```
```text
 确实如此。
先开启并行，再载入自定义词典。此时，从jieba.dt.get_FREQ和jieba.dt.user_word_tag_tab中都能查到用户词典中的词，说明已导入用户词典中的词，但分词结果中并未产生效果。
```
```text
 这是bug吧？同遇到了，定位了半天问题，求修复啊
```
```text
 No description provided.
```
```text
 keywords = jieba.analyse.textrank(content, topK=20, withWeight=True)
for keyword in keywords:
print keyword[0], keyword[1]
和TF-idf一样的方法，这样就可以了。
```
```text
 python3上面使用这种方法将会得到分开的字，不是score.
```
```text
 比如动词 名词之类的
```
```text
 https://github.com/fxsjy/jieba#4-词性标注
```
```text
 词性标注的缩写有没有一个对照表，虽然大部分能猜出来
```
```text
 问题1：
在jieba首页说明中显示自定义词典词频和词性不必须，但是我使用textrank，更改词典会报错！
请问，二者是必须的吗？
问题2：
词频和词性对于textrank算法有何影响？
```
## 241.请问如何自定义切词的粒度！！！！！！
Jul 20, 2016
```text
 class Tokenizer(object):
    ..
    def suggest_freq(self, segment, tune=False):
        ..
        if tune:
            add_word(word, freq)
此处,add_word(word, freq)应为self.add_word(word, freq),否则更新的是默认分词器jieba.dt,而不是用户的new_dt = Tokenizer().
```
```text
 Every time I reboot the Jieba the setting I added using jieba.add_word(word, freq=None, tag=None) will be washed away. Can I keep it somewhere making the setting permanent?
```
```text
 You can write a custom dictionary, and then load the dictionary at start.
```
```text
 当前需要hack Haystack，有点丑：
http://ashin.sinaapp.com/article/119/
http://blog.csdn.net/wenxuansoft/article/details/8170714
```
```text
 No description provided.
```
## 242.在线演示挂了？
Jul 19, 2016
```text
 如 我的需求是""号里面的话不分词
```
```text
 Currently the testcases is in a mess. I wondering if it's ok to make a PR to reorganize all the testcases based on pytest and setup TravisCI to enable continuous integration.
```
```text
 http://jiebademo.ap01.aws.af.cm/
在线演示挂啦，直接跳PasS平台咯。
```
## 243.怎样可以强制提高一些关键词中一些词的权重呀。
Jul 19, 2016
```text
 例如（￣▽￣），(｀・ω・´)等颜文字表情。
在处理的时候，即使添加 add_word 或者 load_userdict ，结果都会被分开成：
（   ￣   ▽   ￣   ）
```
```text
 你好我在我的一个小项目里使用了结巴分词，首先对您的努力表示感谢
我在使用py2exe打包我的程序的时候，其他的模块都没有出现问题，打包能够成功，但在运行的时候一直显示
IOError: [Errno 2] No such file or directory: 'E:\python\links\romeo\dist\l
ibrary.zip\jieba\finalseg\prob_start.py'
我单独写了一个.py文件，只有一句话：
import jieba
seg=jieba.cut('hello jieba')
打包，但运行的时候也是上面的那个结果。

我不知道您是否打包过使用了结巴的程序，如果有，您遇到过这样的问题吗？
或者您知道问题可能出现在哪吗？
谢谢啦~
```
```text
 @lisnb , 我没有用过py2exe。看错误信息，'E:\python\links\romeo\dist\l
ibrary.zip\jieba\finalseg\prob_start.py' 这个py文件被打包在zip里面了，所以读不出来。等会我装个py2exe试一试
```
```text
 @fxsjy  嗯，非常感谢您的回复
您试过之后应该就会看到，所有的依赖项都会被打包在一个zip文件里（如果没有其他设置的话）
也有可能是py2exe的问题，那么您用过其他的什么打包的工具打包过使用了结巴的程序吗？
```
```text
 @lisnb ，试了一下，确实不行。 路径问题不好处理啊，数据文件应该怎么放？
```
```text
 @fxsjy
网上是这么说的，不知道你问的是不是这个
from distutils.core import setup
import glob
import py2exe
setup(console=["myscript.py"],
data_files=[("bitmaps",
["bm/large.gif", "bm/small.gif"]),
("fonts",
glob.glob("fonts*.fnt"))],
)
说明：data_files选项将创建一个子目录dist\bitmaps，其中包含两个.gif文件；一个子目录dist\fonts，其中包含了所有的.fnt文件。
我在想的是，好像分词的字典没有被收进包里，为什么代码文件也没有呢，好奇怪。我以前也没有打包过，我也研究一下...
```
```text
 问题还在，我尝试手工把prob_start.py插入到zip包中，可这也不是解决方法，而且不管用。
```
```text
 @yangboz , 这个问题解决了，你git pull更新一下。 我用cxfree试验成功了。
测试程序： hello.py
#encoding=utf-8
import jieba
jieba.set_dictionary("./dict.txt")
jieba.initialize()
s="我研究生命起源。"
print " ".join(jieba.cut(s))
dict.txt 放在hello.py 同级目录。
```
```text
 pull下来安装最新的jieba 代码后，引用“jieba.set_dictionary("./dict.txt")jieba.initialize()” 路径错误信息是：jieba\posseg..dict.txt;
不用“jieba.set_dictionary("./dict.txt")jieba.initialize()” 后的错误信息是：jieba\posseg\dict.txt;
```
```text
 @yangboz , sorry，昨天的更新没有解决posseg读词典的路径问题，麻烦你再pull一下试一试。
下面是我的测试程序：
#encoding=utf-8
import jieba
jieba.set_dictionary("./dict.txt") #相对路径
#jieba.set_dictionary("c:/tmp/dict.txt")  #也支持绝对路径
jieba.initialize()

from jieba import posseg

s="我研究生命起源。"
print " ".join(jieba.cut(s))
for w in posseg.cut(s):
    print w.word,w.flag
```
```text
 pull到最新代码，最后还有提示提示 File "genericpath.pyc", line 54,in getmtime
WindowError: [Error 3] " "xxx\dist\library.zip\jieba\dict.txt"
我用的是 py2exe，和cxtree应该差不多。
```
```text
 @yangboz , 必须调用jieba.set_dictionary明确指出词典的路径，否则会在默认位置寻找dict.txt。但是被py2exe打包之后就找不到了。
```
```text
 pull到最新的，代码有看到有变化，pseg_cut 有过滤一些空白词汇了，但是最后打包成exe的时候，同时“明确指出词典的路径”后，原来utf-8编码到exe运行的时候成了“cp936”了，还不确定是py2exe的问题。
```
```text
 车库咖啡会场现场解决了:-)
总结如下：
str()操作中文字符+操作，python转换尝试寻找系统默认的编码;
dict定义中文需要"\u";
文件io输出也需要"encode('utf-8)", py2exe中文问题解决.
```
```text
 @yangboz ， 昨晚很有意思，大牛们在上面讲，咱俩在调程序。
```
```text
 你好，这个问题解决了吗？我也碰到相同的问题，很懊恼，jieba官网上说，支持py2exe的，但是制作后，还是不行
```
```text
 在上次来源大会现场作者帮助下，解决了，麻烦你贴下你的错误信息！
Send from Yangbo's iPhone.

On 2013年10月23日, at 15:17, AlgorithmFan notifications@github.com wrote:
你好，这个问题解决了吗？我也碰到相同的问题，很懊恼，jieba官网上说，支持py2exe的，但是制作后，还是不行
—
Reply to this email directly or view it on GitHub.
```
```text
 你好，按照你上面说的方法，加上
jieba.set_dictionary("dict.txt")
jieba.initialize()
仍然出现下面的错误。
Traceback (most recent call last):
File "mainWidget.py", line 6, in 
File "docKeyword.pyc", line 13, in 
File "jieba\posseg__init__.pyc", line 61, in 
File "jieba\posseg__init__.pyc", line 20, in load_model
IOError: [Errno 2] No such file or directory: 'E:\crawWenkuBC\dist\library.zip\jieba\dict.txt'
```
```text
 看错误信息还是这个dict.txt的路径问题，这个文档一直说的很晦涩，你先参考下我这打包程序中的路径 https://github.com/yangboz/hairy-avenger/tree/master/KingOfProgrammer/src/Q2
```
```text
 谢谢，已经解决，犯了低级错误，应该将
jieba.set_dictionary("dict.txt")
jieba.initialize()
直接放在import jieba后面，然后再进行import jieba.posseg as pseg,这样才可以在程序中修改dict.txt的路径，再次表示感谢。
```
```text
 Building prefix dict from C:\Users\game\Desktop\qt\test\dist\dict.txt ...
Loading model from cache c:\users\game\appdata\local\temp\jieba.ua7197607c9829a7854ca3e54b4005544.cache
Loading model cost 0.355 seconds.
Prefix dict has been built succesfully.
Traceback (most recent call last):
File "test.py", line 7, in 
from jieba import analyse
File "jieba\analyse_init_.pyo", line 9, in 
File "jieba\analyse\tfidf.pyo", line 65, in init
File "jieba\analyse\tfidf.pyo", line 42, in init
File "jieba\analyse\tfidf.pyo", line 47, in set_new_path
IOError: [Errno 2] No such file or directory: 'C:\Users\game\Desktop\qt\test\dist\library.zip\jieba\analyse\idf.txt
idf.txt还是不行
`# -- coding: utf-8 --
import sys
sys.path.append("../")
import jieba
jieba.set_dictionary("./dict.txt")
jieba.initialize()
from jieba import analyse
jieba.analyse.set_idf_path("./idf.txt")
if name == 'main':
seg_list = jieba.cut("我来到北京清华大学", cut_all=False)
print("Default Mode: " + "/ ".join(seg_list))  # 默认模式
s = "此外，公司拟对全资子公司吉林欧亚置业有限公司增资4.3亿元，增资后，吉林欧亚置业注册资本由7000万元增加到5亿元。吉林欧亚置业主要经营范围为房地产开发及百货零售等业务。目前在建吉林欧亚城市商业综合体项目。2013年，实现营业收入0万元，实现净利润-139.13万元。"
for x, w in jieba.analyse.extract_tags(s, withWeight=True):
    print('%s %s' % (x, w))`

很简单的代码
```
```text
 还是不行哈
```
```text
 No description provided.
```
## 244.分词编码问题
Jul 18, 2016
```text
 能自定义词性吗？
将一部分特殊词放在自定义词典里，标注自定义词性(例如:unc)，jieba.cut后得到特殊词的词性？
例如：先生，女士，经理，标注为称谓词
```
```text
 明显可以
你可以把下面的加到自定义字典里
先生 w 1000
女士 w 1000
经理 w 1000

在这里
单行字典的结构是
word pos freq
词 词性 词频
```
```text
 这个格式正确吗？
word freq pos?
还有就是会出现自定义词典后，词的词性变为x
```
```text
 word freq pos 是这个格式啊
```
```text
 jieba有官方的QQ群之类的吗？还有jieba有考虑文本摘要的问题么，最近在做这方面的。
```
```text
 File "D:/PythonCodes/fenciExp/jiebaExp/main.py", line 31, in start
for x in jieba.cut(i):
File "D:\��װ\Anaconda2\lib\site-packages\jieba__init__.py", line 301, in cut
for word in cut_block(blk):
File "D:\��װ\Anaconda2\lib\site-packages\jieba__init__.py", line 233, in cut_DAG
DAG = self.get_DAG(sentence)
File "D:\��װ\Anaconda2\lib\site-packages\jieba__init.py", line 179, in get_DAG
self.check_initialized()
File "D:\��װ\Anaconda2\lib\site-packages\jieba__init__.py", line 168, in check_initialized
self.initialize()
File "D:\��װ\Anaconda2\lib\site-packages\jieba__init__.py", line 143, in initialize
self.FREQ, self.total = self.gen_pfdict(self.get_dict_file())
File "D:\��װ\Anaconda2\lib\site-packages\jieba__init__.py", line 352, in get_dict_file
return get_module_res(DEFAULT_DICT_NAME)
File "D:\��װ\Anaconda2\lib\site-packages\jieba_compat.py", line 8, in 
os.path.join(*res))
File "D:\��װ\Anaconda2\lib\site-packages\setuptools-23.0.0-py2.7.egg\pkg_resources__init__.py", line 1178, in resource_stream
File "D:\��װ\Anaconda2\lib\site-packages\setuptools-23.0.0-py2.7.egg\pkg_resources__init__.py", line 1577, in get_resource_stream
File "D:\��װ\Anaconda2\lib\site-packages\setuptools-23.0.0-py2.7.egg\pkg_resources__init__.py", line 1530, in _fn
File "D:\��װ\Anaconda2\lib\ntpath.py", line 85, in join
result_path = result_path + p_path
UnicodeDecodeError: 'ascii' codec can't decode byte 0xb0 in position 1: ordinal not in range(
就是一个普通的测试：
···
import jieba
if name=="main":
str = '我爱北京天安门'
for i in jieba.cut(str):
print i
···
```
```text
 改为
str = u'我爱北京天安门'
Python 处理 UTF-8 的字符串应该在前面加u。
```
```text
 不是这里的问题
读取字典时传入的字符串是 unicode；库提供的绝对路径是 str，该路径包含中文，所以最终 str + unicode 就出现了解码问题。方便的解决方案就是把你的 Anaconda 放在非中文目录。（这种 Python 2 问题代码里不好解决）
```
## 245.单个英文单词的关键词提取
Jul 14, 2016
```text
 比如我有一句话，"2015年度本公司营业利润较2014年度减少-3480.83万元，降幅为-26.95%，主要原因系国内宏观经济增速放缓，下游市场需求不足，市场景气程度有所降低，公司销售收入整体下降。"
我先对结巴做了add_word设置
jieba.add_word("-26.95%")
jieba.add_word(u"-3480.83万元")
然后分词结果是：
2015年度|本|公司|营业利润|较|2014年度|减少|-|3480.83万元|，|降幅|为|-|26.95%|，|主要|原因|系|国内|宏观经济|增速|放缓|，|下游|市场需求|不足|，|市场|景气|程度|有所|降低|，|公司|销售收入|整体|下降|。|
带符号的就不行
```
```text
 遇到同样问题。
加入用户字典    test-test 词 依然会被切分成 test/-/test。
提高词频依然没用，不知道是不是 -  会被固定切出来，不收其他影响。
```
```text
 同样的问题，比如：
jieba.add_word("弗里德里克·泰勒",freq=888888); jieba.add_word("冯·诺伊曼",freq=888888); jieba.suggest_freq(('弗里德里克·泰勒', '冯·诺伊曼'), True); txt = "这个外国人叫弗里德里克·泰勒，是美国前国防部部长。冯·诺伊曼是公认的计算机发明人"; print(" | ".join(jieba.cut(txt, HMM=True, cut_all=False)))
结果还是：
'这个 | 外国人 | 叫 | 弗里德里克 | · | 泰勒 | ， | 是 | 美国 | 前 | 国防部 | 部长 | 。 | 冯 | · | 诺伊曼 | 是 | 公认 | 的 | 计算机 | 发明人'
大家有什么好办法？
```
```text
 有同样的需求，需要保证指定的词绝对不分。不知该怎么做。
```
```text
 代码里面有如下的正则表达式，用来提取中文

re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&._%]+)", re.U)
也就是说对于”弗里德里克·泰勒”中的“·”不认为是中文，而是作为类似逗号和句号来处理。

一个不太好但是可以用的办法就是，修改上面的正则表达式，将“·”加入。其中\xb7就是“·”。

re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&.\xb7%]+)", re.U)

测试结果：
》 这个 | 外国人 | 叫 | 弗里德里克·泰勒 | ， | 是 | 美国 | 前 | 国防部 | 部长 | 。 | 冯·诺伊曼 | 是 | 公认 | 的 | 计算机 | 发明人
```
```text
 小明nr 比p 小红nr 可爱v ，x 但c 都d 没有v 珍珍n 可爱v
```
```text
 如题
```
```text
 单个的英文单词不能通过取关键词取到吗？类似于c这个单词，即使把c加入到自定义的词典也不行啊
```
```text
 tfidf算法抽取，w就是待抽取的单词，如果单词长度小于2,就不处理；
if len(w.strip()) < 2 or w.lower() in self.stop_words:
    continue

textrank算法抽取，只有单词长度大于等于2,才会进行处理；
def pairfilter(self, wp):
    return (wp.flag in self.pos_filt and len(wp.word.strip()) >= 2
            and wp.word.lower() not in self.stop_words)

两个算法，都会较短的单词（字符数少于2）不进行抽取；
```
```text
 @zhbzz2007 很感谢
```
## 246.特殊符号分词错误
Jul 14, 2016
```text
 请问add_word(word, freq=None, tag=None) 这个函数增加的词语只能在当前的python jieba 编译环境下生效，python退出重新登陆时，修改的无效。
现在我需要永久性新增或者删除词语，如何实现永久性热更新操作？
谢谢
```
```text
 add_word()只添加到内存中，jieba好像没有提供固化更新字典的实现。但字典格式好像是通用的，可以用一些其他字典生成工具配套使用。
```
```text
 希望支持数据库词库..
```
```text
 支持这样 cut = jieba.lcut(text.txt) 输入文件的方式吗？试了下，说 NameError: name 'text' is not defined，可这文件同目录下面有的。
```
```text
 lcut 里的是字符串，把文件读进来再用
```
```text
 比如😂这个特殊符号，分词会将这一个符号切成了两个，出现乱码？
```
## 247.分词中，符号的权重是怎样的？
Jun 23, 2016
```text
 上图是我的数据前一部分，我的目的是对 titles 一列进行分词，分词的代码如下。现在遇到的问题是AttributeError: 'int' object has no attribute 'decode'，所以我认为是 titles 中有 int 所致，所以添加了一个判断条件，但是代码执行的结果依旧是报之前的错。请问这是什么原因？
def jiebait(text):
    seglist = jieba.cut(text, cut_all = True)
    fenci = []
    for word in seglist:
        if (not isinstance(word, int)) and (len(word) >= 2):
            fenci.append(word)
    # 如用搜索引擎模式：
    #seglist = jieba.cut_for_search(text)
    return ' '.join(fenci)
```
```text
 def jiebait(text):
    seglist = [str(w) for w in jieba.cut(text, cut_all = True)]
    fenci = []
    for word in seglist:
        if (not isinstance(word, int)) and (len(word) >= 2):
            fenci.append(word)
    # 如用搜索引擎模式：
    #seglist = jieba.cut_for_search(text)
    return ' '.join(fenci)
```
```text
 AttributeError                            Traceback (most recent call last)
 in ()
2     result_line = ""
3     # segment
----> 4     seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]
5     # remove special character
6     temp = re.sub("[\s+.!_,$%^(+"')]+|[+——()?【】“”！，。？、~@#￥%……&（）]+", "",("/".join(seg_list)))
 in (.0)
2     result_line = ""
3     # segment
----> 4     seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]
5     # remove special character
6     temp = re.sub("[\s+.!_,$%^(+"')]+|[+——()?【】“”！，。？、~@#￥%……&（）]+", "",("/".join(seg_list)))
D:\Anaconda\lib\site-packages\jieba_init_.py in cut(self, sentence, cut_all, HMM)
280             - HMM: Whether to use the Hidden Markov Model.
281         '''
--> 282         sentence = strdecode(sentence)
283
284         if cut_all:
D:\Anaconda\lib\site-packages\jieba_compat.py in strdecode(sentence)
35     if not isinstance(sentence, text_type):
36         try:
---> 37             sentence = sentence.decode('utf-8')
38         except UnicodeDecodeError:
39             sentence = sentence.decode('gbk', 'ignore')
AttributeError: 'int' object has no attribute 'decode'
```
```text
 你的 text 本身是不是 int
```
```text
 是的，本身是int，但是seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]这一步，应该都转成string了
```
```text
 解决了 我这边数据转码的问题 不好意思。。。感谢
```
```text
 如题
```
```text
 同問，請教 @fxsjy
```
```text
 同問，請教 @fxsjy
```
```text
 请问有人解决了吗？
```
```text
 我改了一下，目前支持词库中的符号和空格匹配了 https://github.com/WalkerWang731/jieba
```
```text
 在应用时，往往被切分为4, 英寸，7.5, ml
```
```text
 当前我希望获得这种连词的时候，是直接判断量词前是否有数词。有数词自动连接。
但我觉得这是个dirty trick，需要原生支持。
```
```text
 对于这种量词我所呆过公司的做法一般是将这类词当成特殊词语对待，就像品牌有专门的品牌词典和品牌同义词等。所以量词也被当成一个特殊的环节来用特殊的方法对待。一般是用词典和正则。
个人愚见，期待更好的回答。
发自我的 iPhone

在 31 Mar 2014，01:13，"geekan, FSE(Full StackOverflow Engineer)" notifications@github.com 写道：
当前我希望获得这种连词的时候，是直接判断量词前是否有数词。有数词自动连接。
但我觉得这是个dirty trick，需要原生支持。
—
Reply to this email directly or view it on GitHub.
```
```text
 这种量词的可能不是很常见，有一类是AT&T、T恤等常见的品牌或者是中英标点符号共同成词，尽管userdict里有这个词，也无法正确分割。
```
```text
 @aszxqw 正则是对的，但会加倍扫描次数
我猜合并到库里会有比较好的性能
```
```text
 @tuang 这是另外一个问题，我也遇到了。
你开一个issue？
```
```text
 @tuang 我开了一个issue，你看需要怎么补充一下
```
```text
 @geekan新开的issue在哪？
```
```text
 @tuang #145
```
```text
 问题来源：对 'mmol/L9.69' 字符串进行分词
str = 'mmol/L9.69'
jieba.add_word('mmol/L')
jieba.suggest_freq('mmol/L', tune=True)
res = jieba.cut(des, cut_all=True)
结果为：["mmol", "L9", "69"]
若是cut_all=False
res = jieba.cut(des, cut_all=False)
结果为：["mmol", "/", "L9", ".", "69"]
我的目的是：将mmol/l分成一个词。
其他测试：我将'/'改为其他符号，除了'.'号外，其他均不能得到我想要的结果。
有两个问题：1）jieba分词中是不是给符号分了比较高的权重？ 2）如何达成我的目的？
谢谢 :)
```
```text
 我也有这个问题，但是开发人员好像无意优化这个，估计要修改挺多内容的
```
## 248.怎么解决"歼-10"被拆成“歼”，“-”，“10”的问题呀？
Jun 13, 2016
```text
 No description provided.
```
```text
 https://www.cnblogs.com/adienhsuan/p/5674033.html，可以参考一下
```
```text
 我也遇到这个问题了，求告知一个比较全的词性介绍。
```
```text
 找到一个词性介绍，见jieba
```
```text
 官方应该添加一份词性列表到文档中
```
```text
 同问
```
```text
 应该和这个一样
http://ltp.readthedocs.io/zh_CN/latest/appendix.html#id3
```
```text
 生成关键词，里面有一些符号如何处理，如()空格等一些符号
```
```text
 例如：全角符号，在全模式下不会出现。但是在精确模式下就会出现。
```
```text
 你是想要清掉標點符號嗎？ 像這篇？ #169
```
```text
 在一个已经在运行的进程中，我想能够自动加载自定义词典以实现实时对关键词的提取。之前我已经尝试过直接实时判断本地词典文件是否有更新，更新则使用load_userdict来加载词典，但是这个方法好像不管用，新添加进去的关键词还是不能被识别，请问大佬们有什么办法破吗？
```
```text
 CustomDictionary.reload() 可以重新加载词典，但是在加载期间， 分词功能不知道还能不能正常使用。
```
```text
 @csyjgu 敢问大侠CustomDictionary是自定义词典？还是什么？有reload方法吗？
```
```text
 我也有类似的问题，能否在一个工程中能否载入多个字典，然后根据方法需要动态选择使用哪个？
```
```text
 我也有相同的疑問，能否動態選擇要加載哪一種詞典呢?
目前想到兩種做法，

每次更換自定義字典時 使用 jieba.initialize() 初始化，再切換辭典，不過這個方法挺耗時的。
創造多個 jieba實例對象分別載入自定義辭典，再依據需求選用不同的對象。

雖然可行，但都不是最佳解！
```
```text
 CustomDictionary.reload() 可以重新加载词典，但是在加载期间， 分词功能不知道还能不能正常使用。

这是HanLP中的方法吧。
```
```text
 可以用jieba.add_word 和 jieba.del_word 进行动态修改吧, 我现在的问题是web应用程序是多进程运行的，add和del之后只对当前处理请求的进程生效 。
```
```text
 自定义词典里写上“歼-10”了，也没有起作用。
```
```text
 添加到自定义词库
```
```text
 同问！大家有什么办法没呢？
```
## 249.结巴所用的POStag集的详细文档能够附上一份吗？
Jun 11, 2016
```text
 免费分享，造损免责。
打开默认词典（根目录）或自定义词典，把所有用来间隔词频和词性的空格间隔符改成@@
（选用@@是因为一般关键词里遇到这个分隔符的几率比较小吧）
继续，打开jieba根目录下init.py

搜索
re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._]+)", re.U)
改成
re_han_default = re.compile("(.+)", re.U)



搜索
re_userdict = re.compile('^(.+?)( [0-9]+)?( [a-z]+)?$', re.U)
改成
re_userdict = re.compile('^(.+?)(\u0040\u0040[0-9]+)?(\u0040\u0040[a-z]+)?$', re.U)



搜索
word, freq = line.split(' ')[:2]
改成
word, freq = line.split('\u0040\u0040')[:2]



补充：若用的全模式继续改。
搜索
re_han_cut_all = re.compile("([\u4E00-\u9FD5]+)", re.U)
改成
re_han_cut_all = re.compile("(.+)", re.U)
```
```text
 您好，我正好碰见这个问题，我想把 迈克尔·乔丹 不要拆开，试了您的方法，
将默认的dict改成了如下：
AT&T@@3@@nz
B超@@3@@n
c#@@3@@nz
C#@@3@@nz
c++@@3@@nz
将user_dict变成了如下：
小皇帝
科比布莱恩特
迈克尔·乔丹
init.py也按照您的方法改了，没用全模式所以没改最后一个
加载报错：
ValueError: invalid POS dictionary entry in C:\Tools\anaconda\lib\site-packages\jieba\dict_gai.txt at Line 1: AT&T@@3@@nz
请您指教！感激不尽！
```
```text
 @xusong123 我试了一下，没有问题啊，是不是用用了记事本打开txt导致加上了bom头。建议用notepad++创建编辑文本。
```
```text
 您好，我试了试，还是不行，方便的话我把这个贴上来给您看看
dict_gai3.txt
我加载的是dict_gai3.txt
```
```text
 @xusong123  你试一下，不用我的方法，然后把你字典里的点·删掉，然后看看报错不。如果也报错，就证明你后来新加的词，编码上有问题。
```
```text
 谢谢，已经解决，最好用python3
```
```text
 @xusong123 请问怎么解决的？我也出现同样的错误了，谢谢
```
```text
 我也遇到了这个问题。通常是/tmp/jieba.cache未即时更新引起的。
解决方法：删除jieba.cache，把默认字典（dict.txt）中的空格替换为@@即可。
@wonderfreer
```
```text
 我需要加载千万级的词典，词典才300M，但是加载后内存高达5G。
self.FREQ[word] = freq
self.total += freq
if tag:
self.user_word_tag_tab[word] = tag
for ch in xrange(len(word)):
wfrag = word[:ch + 1]
if wfrag not in self.FREQ:
self.FREQ[wfrag] =
FREQ
user_word_tag_tab
用了多个字典去存储数据，是不是浪费内存啊，这里可以合并成一个词典吧
```
```text
 @zhangzhenhu ， 可以试一试，发个PR吧 :-)
```
```text
 @fxsjy 128m内存直接跑不起来，256的才勉强能跑。
我用的bae。
```
```text
 是的，我也遇到一样的问题。加载词典太耗内存了！
```
```text
 请教一个词性标注的问题
现在的jieba.posseg.cut标注的词性w.flag，是不是都是从字典里面取出来的？
但是字典里面只定义了一种词性，而很多词是多词性的。
例如'连'这个词，在金山词霸网站可以查到，'连'有多重词性
http://www.iciba.com/%E8%BF%9E

本义：（动）相连；连接。（副）连续；接续；继续（表示前后没有断开）：～演三十场戏。
本义：（动）相连；连接。（介）包括在内：～我三个人｜～根拔。
本义：（动）相连；连接。（名）军队的编制单位；由若干排组成。
本义：（动）相连；连接。（Lián）姓。

ltp是哈工大在线工具分析结果：  http://www.ltp-cloud.com/demo/
例句一：连小学生都会。
ltp:   [连/u  小学生/n  都/d  会/v  。/wp]
jieba: [连/nr 小学生/nr 都/d 会/v 。/x]       -> 字典： 连 23315 nr
例句二：这是一个连的兵力。
ltp:   [这/r 是/v  一个/m  连/n  的/u  兵力/n  。/wp]
jieba: [这/r 是/v  一个/m  连的/d 兵力/n 。/x]
【问题】怎样定义多性词？比如'连'在第一个句子里应该是u，第二个句子里应该是n
谢谢。
```
```text
 目前jieba的POS Tagging，基于词库分词后，对词库中存在的词，直接取用词典中的词性（第三列为词性）；对于未登录词，再用HMM序列标注来同时完成新词发现和词性标注。
对于多性词，jieba的方案比较差，而中文中多性词约占了23.6%，比例还是很大的。
哈工大的LTP并不死绑词典中的词性，而考虑了上下文，并针对未登录词做了不少优化，相较之下好很多。
```
```text
 因为ICTCLAS的旧文档已经不知去向。
而且ICT POS 3.0已经不同了。
谢谢
我有找到过一个网站的说明，不知上面所说的关于jieba分词的词性标注集是否还适用

http://www.hankcs.com/nlp/corpus/several-revenue-segmentation-system-used-set-of-source-tagging.html
```
## 250.jieba加载速度能否有优化方式.
Jun 6, 2016
```text
 我想先分完词，再基于分词结果做词性标注，请问结巴分词支持吗？我看哈工大的pyltp是支持的（[https://blog.csdn.net/baidu_15113429/article/details/78909666]），但是我不想使用pyltp
```
```text
 這裡好像快倒了 沒啥人氣 慘....
```
```text
 https://blog.csdn.net/enter89/article/details/80619805
```
```text
 參考我修正的項目 #670
```
```text
 每次都这样...
Building prefix dict from /usr/local/lib/python2.7/site-packages/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.704 seconds.
Prefix dict has been built succesfully.
```
```text
 同样遇到加载慢的问题。通常是首次加载，或者长时间空闲后首次加载。
有时会慢至 6、7 秒甚至十几秒。
为什么？
那个 cache 文件很小啊。
```
## 251.import jieba.analyse => ImportError: No module named analyse
May 24, 2016
```text
 `



seg_list = jieba.cut("他来到了网易杭研大厦")  # 默认是精确模式
print(", ".join(seg_list))
他, 来到, 了, 网易, 杭研, 大厦`



但是在分词过程中判断词性
`



words = pseg.cut("来到了网易杭研大厦")
for word, flag in words:
...  print('%s %s' % (word, flag))`



依旧会把“杭研”分为两个词  杭 j 研 vn  ；
如何判断“杭研”是否为新词？
```
```text
 可以參考我修正的錯誤 #670
```
```text
 用demo里的程序就没有问题，自己新建一个python文件，复制里面的import，执行时总是提示ImportError: No module named analyse
什么原因
C:\Python27\Lib\site-packages\jieba-0.38-py2.7.egg-info
jieba-0.38-py2.7.egg-info\jieba
都存在
```
```text
 应该是包结构的问题，看看你init.py是否正确，试试不用没有包结构的情况
```
## 252.似乎没有添加自己的停用词可能，是不是可以考虑添加
May 19, 2016
```text
 jieba.posseg.cut是可以分割词并得到tag类型，但是却把百分比分割成两个部分，tag都是'x'，jieba.cut就能准确分成"2%"，如果需要tag还需要百分比，应该怎么写
```
```text
 同样遇到了这个问题
```
```text
 同问
```
```text
 已解决，见https://blog.csdn.net/liuhongyue/article/details/80498195
```
```text
 似乎没有添加自己的停用词可能，是不是可以考虑添加
```
## 253.带符号的专有名词如(C++, C#, node.js objective-c)
May 19, 2016
```text
 jieba.analyse.textrank()方法提取英文文本的关键词没有输出数据，但是jieba.analyse.tfidf()方法则可以处理英文文本。中文文本两个都可以处理。so ,问题是，，，
```
```text
 你看看源代码就知道了，tokenizer和postokenizer的区别。
```
```text
 re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&._]+)", re.U)
既然对C++, C#多了特判, 为什么不加入 objective-c特判
```
## 254.【关于字典的更新】Jieba缺省带了字典，希望可以提供一个机制，让Jieba使用者也可以提交字典更新
May 10, 2016
```text
 需求有些特别：要解决的问题的起点并不是文章。而是经过正确分词后的list。有没有什么办法在不经过jieba分词功能的情况下，直接使用关键词提取的功能（调用关键词提取方法之前，传入自定义的分词结果－list格式）
```
```text
 你是不是问题没描述明白？
你自定义的分词结果是什么东西分出来的？如果是jieba分出来的，那就是可以用自定义词典。
如果是其他方式分出来的，可以A+B。合并列表。
```
```text
 字典是中文分词的核心数据。分词的好坏，往往由字典的正确率和是否与时俱进决定。设想中，应该有这么一个机制，允许Jieba的使用者一起来贡献字典的更新。
1、比如，让Jieba的垂直行业用户，可以把针对他们的行业已经tain过的字典上传上来，分享给大家
2、再比如，Jieba自己的字典，是否可以通过对于当前一段时间互联网的最新热词/短语的流行潮流进行自我更新？
```
```text
 单开一个 repo，专门存放分享上来的字典？
通过 pull request 来分享、更新
不过这样要有人来维护这个 repo，要确保字典质量，要花精力的啊
```
```text
 @anjianshi 众筹词典？
```
```text
 众筹是大家出钱然后由某个人 / 组织负责提供、维护字典吗？不知道好不好操作 :)
```
```text
 导入自定义词库可以满足个人需求。
```
## 255.請問如何避免非中英文的詞(如Citroën)會被切開
May 5, 2016
```text
 比如某个句子的分词效果，结巴的结果和预期有出入
我想把人工标注的结果加入到结巴，如何？
```
```text
 非中英文的詞，如Citroën會被切成Citro+ë+n。請問可以如何避免？
看似系統將ë等特殊字看作符號。我有機會修改source code，將原本的英數字集加入此些特殊字嗎？目前看到相關的設定是jieba/posseg/init.py檔案的re_skip_detail = re.compile("([.0-9]+|[a-zA-Z0-9]+)")。
或是有其它建議？
謝謝。
```
## 256.有关于词性的吗？
Apr 26, 2016
```text
 jieba.enable_parallel()
jieba.load_userdict("./aux/dict")
如上所示，如果先开启并行，再载入自定义词典，会导致自定义的词典没有产生效果
如果先载入自定义词典再进行并行计算开启，自定义词典就是有效的，不知道这算不算bug
```
```text
 确实如此。
先开启并行，再载入自定义词典。此时，从jieba.dt.get_FREQ和jieba.dt.user_word_tag_tab中都能查到用户词典中的词，说明已导入用户词典中的词，但分词结果中并未产生效果。
```
```text
 这是bug吧？同遇到了，定位了半天问题，求修复啊
```
```text
 No description provided.
```
```text
 keywords = jieba.analyse.textrank(content, topK=20, withWeight=True)
for keyword in keywords:
print keyword[0], keyword[1]
和TF-idf一样的方法，这样就可以了。
```
```text
 python3上面使用这种方法将会得到分开的字，不是score.
```
```text
 比如动词 名词之类的
```
```text
 https://github.com/fxsjy/jieba#4-词性标注
```
```text
 词性标注的缩写有没有一个对照表，虽然大部分能猜出来
```
## 257.如何在论文中引用结巴分词？
Apr 25, 2016
```text
 关键词提取无论如何总能得到一些结果，在语料非常多的时候，人工标记不太具备可行性，那么有哪些推荐的算法用于评估提取内容的准确性和相关性呢？
```
```text
 长标点符号应该被分为一个词。
例如：
longpunct = set(u'——', #EM dash U+2014, 常用
u'--', #半角-
u'－－', #U+FF0D, 全角－
u'――', #Horizontal bar, U+2015
u'──', #Box Drawing Light Horizontal, U+2500
u'……',
u'......', # '.'*6
u'...', # '.'*3
) #etc.
```
```text
 么么哒~
```
## 258.分词结果异常
Apr 20, 2016
```text
 No description provided.
```
```text
 jieba.0.38

不加载自定义词时，分词正常

from jieba import posseg as pseg
pseg.cut(u"为何 MOTO 在华销售的绝大多数")

为何/r MOTO/eng 销售/vn


加载自定义词时，分词异常

from jieba import posseg as pseg
jieba.load_userdict("???")

pseg.cut(u"为何 MOTO 在华销售的绝大多数")

为何/r M/x OTO/x 销售/vn 

自定义词典中含有 OTO
```
## 259.关键词抽取处理停用词时的一个UE设计问题
Apr 9, 2016
```text
 rt
```
```text
 为什么简体字文章用jieba.analyse.extract_tags提取关键字是繁体字
```
```text
 我几个月前下了jieba使用，发现抽取关键词时，停用词未做处理，所以在网上找了几个停用词表merge以后，手工做了处理。
现在看已经增加了自定义停用词表功能，不过我个人建议是：
1.需要有一个默认的停用词表，在不附带任何选项的情况下，analyse.extract_tags()应该返回经过默认停用词表处理后的结果。
2.假如有人不愿意要停用词，可以在该函数增加一个选项禁用停用词表。
3.假如有人想用自己的停用词表，和现在一样以增设自定义词典的处理即可。
附件为我目前使用的停用词表
stop_words.txt
```
```text
 这个改动建议和各选项的使用频率有关。90%以上的使用者在关键词抽取时是希望删去停用词的，因此应该作为默认选项使用。而不是让大家自己各显神通去找停用词表。
```
## 260.添加自定义词典不起作用
Apr 1, 2016
```text
 jieba.lcut('2017年10月5日或2017-10-03或12:21和12点30分还有十二点三十分')
分出来
['2017年', '10月', '5日', '或', '2017', '-', '10', '-', '03', '或', '12', ':', '21', '和', '12点', '30分', '还有', '十二点', '三十分']
如何分成
['2017年10月5日', '或', '2017-10-3', '或', '12:21', '和', '12点30分', '还有', '十二点三十分']
```
```text
 把这些词汇加入到词典中

发自我的vivo智能手机

sugarZ <notifications@github.com>编写：
…
jieba.lcut('2017年10月5日或2017-10-03或12:21和12点30分还有十二点三十分')
分出来
['2017年', '10月', '5日', '或', '2017', '-', '10', '-', '03', '或', '12', ':', '21', '和', '12点', '30分', '还有', '十二点', '三十分']
如何分成
['2017年10月5日', '或', '2017-10-3', '或', '12:21', '和', '12点30分', '还有', '十二点三十分']
—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 之前也想到了,如果固定的几个还可以,但是非常大量的日期时间格式,没法全部加到字典,这方法太笨了.
```
```text
 @sugarZ，请问lcut可以将“2017年”分词到一起吗，貌似我这边是“2017”，“年”
```
```text
 @JiaWenqi 好像不行,我是加的自定义词典,把最近的年份都添加了进去,还有1-12月0-24点(钟),但是如果把时间也加上就太麻烦了,如果jieba支持特定格式分词配置就好了.
```
```text
 目前想到一个方案，将待分词文本用时间正则进行分割后分段进行分词
```
```text
 在词典中添加a/d,d/a分词的结果仍然是三部分，不能形成一个词，
权重调了也没有用，估计是特殊字符的问题。
```
```text
 与这个类似：#363
```
## 261.是否可以用清华THULAC使用的语料库来训练jieba的模型？
Mar 31, 2016
```text
 代码如下：
string="这是一个test的行。和一些无意义的dm、jksajdfl"
jieba.set_dictionary('dict.txt.big')#自定义词库
res=jieba.lcut(string,HMM=False)
print(res)
['test', 'dm', 'jksajdfl']
自定义的dict里面只有一个测试的汉字词。
```
```text
 自己解决了。
init.py 破坏掉re_eng = re.compile('[a-zA-Z0-9]', re.U) 就可以了。。。。
比如re_eng = re.compile('aaaaaaaa', re.U)
```
```text
 清华最近放出来的THULAC看起来很赞的样子，文档中提到了训练模型所用的语料库，官方团队的资源获取能力实在是太厉害了。不知道可否和清华方面沟通一下，用他们的语料库来训练jieba的模型。
THULAC

1.THULAC工具包提供的模型是如何得到的？
THULAC工具包随包附带的分词模型Model_1以及分词和词性标注模型Model_2是由人民日报语料库训练得到的。这份语料库中包含已标注的字数约为一千二百万字。
同时，我们还提供更复杂、完善和精确的分词和词性标注联合模型Model_3和分词词表。该模型是由多语料联合训练训练得到（语料包括来自多文体的标注文本和人民日报标注文本等）。这份语料包含已标注的字数约为五千八百万字。由于模型较大，如有机构或个人需要，请填写“doc/资源申请表.doc”，并发送至 thunlp@gmail.com ，通过审核后我们会将相关资源发送给联系人。
```
```text
 😄 说实话并没有觉得很赞……
```
## 262.如何多次载入自定义词典？
Mar 28, 2016
```text
 部分地名词性为nr，比如“孙吴县”。建议按照民政部最新的行政区划表更新词库。
2018行政区划代码
```
```text
 如何解决向字典添加“leap motion”这种中间带空格的关键词
```
```text
 @microcao , 目前不行，得改代码。 词典里面的分割符就是空格，此外还要改一处代码里面的正则表达式，把空格也作为成词的合法字符。
```
```text
 @fxsjy 能帮忙搞一下带空格的英文单词的实现吗？感谢！
```
```text
 同问。@fxsjy
```
```text
 @geekan 这会带来另一个问题。如果“leap motion”作为一个词的话，那么就要考虑：搜索引擎模式下leap和motion怎样切分出来？
```
```text
 @anderscui 不了解这里的逻辑，不过这个看起来像是一个通用的问题：

长短词同时存在时，搜索引擎模式能否全切分？我觉得是必须要全切分的。
```
```text
 @geekan 对，所以我在考虑，能否先按照现有方式切分，然后根据一个配置文件将包含空格的词merge，而不是把这样的词直接加入自定义词典。
```
```text
 @anderscui 不是很明白，为啥要单独定义配置文件呢？
是因为搜索引擎模式对空格有特殊处理吗？
```
```text
 进程是长时间运行进程，这个过程中自定义词典会更新，需要在词典更新后使用新的自定义词典。
文档中没有关于这个功能的说明。是否可以增加这个功能呢？
```
```text
 你这个解决了吗？
我现在做法是重新调用load_userdict方法来重新加载自定义词典
```
## 263.jieba.cut()和jieba.posseg.cut()的分词结果不一致？
Mar 28, 2016
```text
 `text = "给我办一个三十的新流量业务"
result = []
for (word, start, end) in jieba.tokenize(text):
pseg_data = [(w, f) for (w, f) in pseg.cut(word)]
print((pseg_data, start, end))
result.append((pseg_data, start, end))
`
([('给', 'p')], 0, 1)
([('我', 'r')], 1, 2)
([('办', 'v')], 2, 3)
([('一个三十', 'm')], 3, 7)
([('的', 'uj')], 7, 8)
([('新', 'a')], 8, 9)
([('流量', 'n')], 9, 11)
([('业务', 'n')], 11, 13)
如果是"两个三十"就可以分开
```
```text
 你好，我在使用结巴分词时，出现同同一词汇词性不同的情况，我使用了动态加词汇，修改dict，导入自己的词汇表，都会出现这种情况，比如我定义：街 1000000000000 ns 我的目标是分割出的“街”都是“ns”，但是在实际中会出现：(m)欢乐(a)街(n)怎么(r)走(v)？  去(v)美食(n)街(ns) 这两种不同的情况，请问是怎么回事呢？
ps:我只是在使用jieba,并不懂里面的理论，如果有基本错误还请原谅。
```
```text
 对同一段文本使用上述两个函数进行分词，分词结果不一样？
如：弄正
在jieba.cut()中分成了”弄“和”正“
在jieba.posseg.cut()中则保持”弄正“
```
## 264.在线演示出错了，需要升级
Mar 11, 2016
```text
 代码入下：
content="环境设施都很好，服务周到！"
aspect_content_array = jieba.analyse.extract_tags(content, topK=50)
print ",".join(aspect_content_array)
抽取结果：
服务周到,设施,环境
我想得到“好”这个关键词，但是抽取不出来
另外能否在抽关键词的时候，把“服务周到” 抽取成 “服务”和“周到”两个词
```
```text
 刚试了下自定义词典，里边加入这两行：
打底裤 10000
打底衫 10000
发现这两个词仍然分开了。debug之后发现这些词并没有加载。问题在这里：
jieba/init.py:
def load_userdict(f):
......
if not line.rstrip():
continue
tup = line.split(" ")
word, freq = tup[0], tup[1]
if freq.isdigit() is False:
continue
windows下userdict结尾是"\r\n", 因为我的词典，词频“1000”后面立刻是"\r\n", "\n"被split吃掉了，"\r"还在line结尾。所以tup里，freq="1000\r"，判断 isdigit()为False，这几个词都没被加载。
怀疑前面的
if not line.rstrip():
continue
是不是应该是
line = line.rstrip()
if not line:
continue
这么改之后，词是加载了（在load_userdict()里用print看了），而且“打底衫”分好了，但“打底裤”仍然分成了“打/底裤”。我在词频后面加两个零，结果还是一样。那是什么问题呢？
```
```text
 试试 Git 里的新版本
```
```text
 @gumblex 找到原因了。。原来是我画蛇添足给文件加了个BOM头。老版本并没有删掉这个头。现在好了 :)
```
```text
 @gumblex  java版的结巴， 支持添加自定义词典吗？
```
```text
 @myccc456 Jython 应该是支持的。如果你说的是 这个 Java 版本，根据 README 也应该是支持的。
```
```text
 @gumblex   谢谢， 试了下，通过如下代码即可自定义添加词典 java版
    WordDictionary dictAdd = WordDictionary.getInstance();

    File file = new File("D:/jieba-analysis/conf/user.dict");
    dictAdd.loadUserDict(file);
```
```text
 Thank you for visiting AppFog
The final sunset date for AppFog v1 is March 15, 2016, (located here at appfog.com). After March 15, 2016 AppFog v1 will no longer be available.
AppFog v2, located at ctl.io/appfog, has replaced AppFog v1.
For more information on this migration and needed action for AppFog v1 users, please see our original communication and migration guides.
```
## 265.在提取关键词的时候，如何去除掉停止词
Mar 9, 2016
```text
 需求：
要求匹配出内容的指定的一些关键词（我自己的词库里的词），忽略jieba里千千万万不相关的词，对于我来说，这些词没有用。
实现：
我想的是，把jieba默认的词库替换掉。可是，我看文档里没有提到如何替换jieba的词库
还有什么好的办法实现吗？
谢谢各位
```
```text
 你可以用jieba的自定义字典实现准确识别

中铁建资产管理有限公司 高正  18601064889
…
 在 2018年10月24日，11:11，Bakkan Hwang ***@***.***> 写道：

 需求：
 要求匹配出内容的指定的一些关键词（我自己的词库里的词），忽略jieba里千千万万不相关的词，对于我来说，这些词没有用。

 实现：
 我想的是，把jieba默认的词库替换掉。可是，我看文档里没有提到如何替换jieba的词库
 还有什么好的办法实现吗？

 谢谢各位

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 jieba有实现延时加载机制
jieba 采用延迟加载，import jieba 和 jieba.Tokenizer() 不会立即触发词典的加载，一旦有必要才开始加载词典构建前缀字典。如果你想手工初始 jieba，也可以手动初始化。
import jieba
jieba.initialize()  # 手动初始化（可选）
在 0.28 之前的版本是不能指定主词典的路径的，有了延迟加载机制后，你可以改变主词典的路径:
jieba.set_dictionary('data/dict.txt.big')
例子： https://github.com/fxsjy/jieba/blob/master/test/test_change_dictpath.py
```
```text
 In [4]: print '/'.join(jieba.cut("我来到清华大学",cut_all=True))
我/来到/清华/清华大学/华大/大学
In [5]: jieba.add_word('到清')
In [6]: print '/'.join(jieba.cut("我来到清华大学",cut_all=True))
我/来到/到清/清华/清华大学/华大/大学
In [7]: jieba.add_word('学')
In [8]: print '/'.join(jieba.cut("我来到清华大学",cut_all=True))
我/来到/到清/清华/清华大学/华大/大学
希望高手能给予帮助!
```
```text
 同问
```
```text
 #408
```
```text
 Hello, does anybody know from where the dictionary of traditional Chinese characters originated? Who created that dictionary or from where was it ported?
```
```text
 在提取关键词的时候：
1、如何去除掉自定义的停止词；
2、如何增加自定义的词库，自定义的词库的格式是什么？
```
```text
 词库格式：词语 词频 词性，空格分割，可以省略词频、词性，可以同时省略词频和词性
文件必须为UTF-8无BOM编码
格式可参考userdict.txt
具体请参考readme中的相关说明
```
```text
 停止词列表可以自定义
readme中有写

关键词提取所使用停止词（Stop Words）文本语料库可以切换成自定义语料库的路径
用法： jieba.analyse.set_stop_words(file_name) # file_name为自定义语料库的路径
自定义语料库示例：https://github.com/fxsjy/jieba/blob/master/extra_dict/stop_words.txt
用法示例：https://github.com/fxsjy/jieba/blob/master/test/extract_tags_stop_words.py
```
```text
 但是0.36版以后analyse这个模块已经无法调用，没有写在方法里
```
```text
 @ynyxxy 我这一切正常，版本：0.38，测试代码：
import jieba.analyse
jieba.analyse.set_stop_words("D:/a.txt")
```
```text
 @1354092549 我再看看，似乎是安装除了问题
```
## 266.jieba.dt找不到
Mar 8, 2016
```text
 比如遇到本身就有语义或者拼写的错误的中文：咳嗽->咳嗽等等
想从语义和词性上通过字典来实现可以吗？
```
```text
 没有的吧，据我使用的经验来说是没有的。
这里评价体系比较难弄
据我的使用经验而言我们是自己建立了一套转换体系
```
```text
 Hello,
請問應該如何避免句子「我賺了123,244.2元！」被錯誤的分隔成
['我', '賺', '了', '123', ',', '244.2', '元', '！']

謝謝！
```
```text
 import jieba

File "C:\Python27\lib\site-packages\jieba__init__.py", line 9, in 
import logging
File "C:\Python27\lib\logging__init__.py", line 26, in 
import sys, os, time, cStringIO, traceback, warnings, weakref, collections
File "C:\Python27\lib\collections.py", line 10, in 
from keyword import iskeyword as iskeyword
File "C:\Users\2200621\Dropbox\Learning\Python\py\spider\keyword.py", line 4,
in 
import jieba.analyse
File "C:\Python27\lib\site-packages\jieba\analyse__init_.py", line 2, in 
from .tfidf import TFIDF
File "C:\Python27\lib\site-packages\jieba\analyse\tfidf.py", line 5, in 
import jieba.posseg
File "C:\Python27\lib\site-packages\jieba\posseg__init__.py", line 257, in 
dt = POSTokenizer(jieba.dt)
AttributeError: 'module' object has no attribute 'dt'
```
```text
 如果我没记错，应该是 0.37 之后才加上 dt ，不妨检查下 jieba 的版本
```
```text
 同样遇到这个问题。
Traceback (most recent call last):
File "demo.py", line 6, in 
import jieba
File "C:\Python27\lib\site-packages\jieba__init__.py", line 9, in 
import logging
File "C:\Python27\lib\logging__init__.py", line 26, in 
import sys, os, time, cStringIO, traceback, warnings, weakref, collections
File "C:\Python27\lib\collections.py", line 10, in 
from keyword import iskeyword as iskeyword
File "D:\python\jieba\keyword.py", line 7, in 
import jieba.posseg
File "C:\Python27\lib\site-packages\jieba\posseg__init_.py", line 257, in 
dt = POSTokenizer(jieba.dt)
AttributeError: 'module' object has no attribute 'dt'
```
```text
 我也碰到楼上的问题，之前碰到的是找不到jieba.analyse模块。
重新安装之后就遇到楼上的问题，也许依赖顺序的问题？
我是用的是python2.7．
```
```text
 @seaguest 和 Python 版本没关系，和 jieba 版本有关系， import jieba 后执行 print jieba.__version__ 和 print jieba.__file__ 看看就知道了
```
```text
 @Linusp
经过试验证明，是我的一个文件命名的问题。
我有一个文件叫keyword.py，这个文件导致import jieba出错，当我吧名字改成keywords.py,问题就解决了。不知道这背后是什么原因。
上面的那个错误是发生在import jieba的时候，所以无法执行 print jieba.version 和 print jieba.file。
```
```text
 @seaguest , thanks. your solution works. It happens to me as well.
```
## 267.如何自定义IDF文件
Mar 5, 2016
```text
 用户自定义词典定义了中文和英文结合的词 ，分词没用效果
比如词典里定义了一个a字群
分词好像没有效果
```
```text
 import jieba
a=jieba.cut("阿里妈妈1");print(" ".join(a))
Building prefix dict from /usr/lib/python2.6/site-packages/jieba-0.37-py2.6.egg/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.256 seconds.
Prefix dict has been built succesfully.
阿里 妈妈 1
jieba.add_word('阿里妈妈1',100,'nz')
a=jieba.cut("阿里妈妈1");print(" ".join(a))
阿里妈妈1
jieba.del_word('阿里妈妈1')
a=jieba.cut("阿里妈妈1");print(" ".join(a))
阿里妈妈1
```
```text
 这个问题有解决吗?
```
```text
 为什么这两行的判断条件一致呢？
if len(obs_states)==0: obs_states = prev_states_expect_next
if len(obs_states)==0: obs_states = all_states
```
```text
 在抽取关键词 时，需用到IDF_big 文件，请根据需要用到的语料库，生成自己的IDF文件，用于计算TFIDF
```
## 268.中英文数字混合切分
Mar 3, 2016
```text
 No description provided.
```
```text
 首先感谢 @fxsjy 的开源共享精神，让我们能处理中文分词，真的很棒！
那么问题来了，很多人都想用机器分析用户输入的中文语义，然后做后续的应用开发，比如机器人、人工智能、客服bot……
有什么好办法去做语义分析吗
比如，我想知道用户的一句话是一般陈述句还是一般疑问句，还是特殊疑问句，主谓宾是谁
这样才能让机器去检索数据库，才能回答给具体主语还是宾语这个主体
目前google开源了他的SyntaxNet，但是只对英文有效，有谁研究过这个么
```
```text
 中英文数字混合切分时会出现英文字母和长串数字切在一起的情况。
eg:
h123123123
应该切分为h 123123123
```
## 269.结巴分词怎么运行整个文件啊
Feb 29, 2016
```text
 In [94]: s = '乌鲁木齐爱家超市南门店'

In [95]: jieba.cut(s)
Out[95]: <generator object Tokenizer.cut at 0x10e24fde0>

In [96]: jieba.lcut(s)
Out[96]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [97]: jieba.add_word('南门')

In [98]: jieba.lcut(s)
Out[98]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [99]: jieba.suggest_freq('南门', True)
Out[99]: 833

In [100]: jieba.lcut(s)
Out[100]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [101]: jieba.lcut(s, HMM=False)
Out[101]: ['乌鲁木齐', '爱家', '超市', '南', '门店']
期望结果
['乌鲁木齐', '爱家', '超市', '南门', '店']
```
```text
 In [94]: s = '乌鲁木齐爱家超市南门店'

In [95]: jieba.cut(s)
Out[95]: <generator object Tokenizer.cut at 0x10e24fde0>

In [96]: jieba.lcut(s)
Out[96]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [97]: jieba.add_word('南门')

In [98]: jieba.lcut(s)
Out[98]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [99]: jieba.suggest_freq('南门', True)
Out[99]: 833

In [100]: jieba.lcut(s)
Out[100]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [101]: jieba.lcut(s, HMM=False)
Out[101]: ['乌鲁木齐', '爱家', '超市', '南', '门店']
期望结果
['乌鲁木齐', '爱家', '超市', '南门', '店']

@4ft35t 设置的频率还是太低了，使用jieba.add_word("南门", freq=1000)设置高一点就可以了。
```
```text
 I wonder the efficiency of Jieba python edition.
I have a set of corpus around 1.5G size, with 3 of 4 cores parallel processing, using Jieba python and have waited for more than 2hr to finish the segment task.
Is this normal or anything wrong?
My platform is MBP '13 spring with 10G RAM, MacOS 10.9., python 3.4.2.
Thanks a lot.
```
```text
 No description provided.
```
## 270.中文标点符号如何处理？
Feb 29, 2016
```text
 现在我cut
我想看电视
结果是
我 想 看电视
看了下默认的词库里有 看电视 这个词。。
我想搞成
我 想 看 电视
```
```text
 jieba.add_word("看电视",0)#或者jieba.del_word("看电视")#发现主干这个del_word的bug已经修了
jieba.add_word("看",100,"v")
jieba.add_word("电视",100,"n")
```
```text
 我自定义了一个词典，有一些包含“-”的词，比如 “SK-II”。加载这个词库时候发现只有这类词语没法生效，会被分为 SK，-，II,而不是SK-II。请问这个问题应该如何解决
```
```text
 同问，自定义词典中有5%,但分词会分开，词性标注的会分开 但正常分词不会被分开
```
```text
 找到解决方法了
修改jieba/init.py代码中的re_han_default 正则表达式为如下值：
re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._%\-]+)", re.U)
```
```text
 Can solve this: #692
```
```text
 ganerating advanced nets
```
```text
 generative adversarial networks?
```
```text
 切词后会有中文标点，我想要去掉， 有没有相应的词表？
```
```text
 切词之前正则过滤吧，方便快捷。
表达式随手google一下就有
```
## 271.jieba.load_userdict 可以支持tuple吗？
Feb 16, 2016
```text
 比如要分词的字符串：http://192.168.0.121/login.view
因为：re_num = re.compile("[.0-9]+")这样写，
分词的结果为：
login eng
. m
view eng
"."被判断成数字，建议修改一下这个正则
```
```text
 这个“.”  很有必要被当做数字，信我。
```
```text
 查看源码,代码中会对文本进行decode,最终生成的分词结果并没有encode回来?
如果后续用其他工具对分词结果进行处理,会出UnicodeEncodeError,如果对分词结果手动encode就没问题.
猜想decode是为了把一个字当做一个字符处理,防止一个字的长度大于1.但是切词完之后不encode回来算是隐性的bug?
```
```text
 No description provided.
```
```text
 @54xiaosu 您好。
我之前仿照楼主大神的算法写了一个，https://github.com/aszxqw/cppjieba
经线上测试，挺稳定。效果蛮好。不知道是否对您有所帮助 :)
```
```text
 @aszxqw ，线上使用了啊，方便透漏是哪家网站吗？
```
```text
 @fxsjy 已私聊。
```
```text
 jieba.load_userdict 可以支持tuple吗？我想从数据库导出后一把调这个函数，省得到file再绕一圈，或者反复suggest
```
```text
 循环用 add_word(word, freq=None, tag=None)，freq 不填就是 suggest
```
```text
 谢谢
```
```text
 请教以下add_word是永久添加到词典里还是临时添加啊？
```
```text
 我试了一下，是临时 @natsuapo
```
## 272.AttributeError: module 'jieba' has no attribute 'cut'
Jan 23, 2016
```text
 Hi I'm using Jieba with Java and I have a question.
My purpose to use Jieba is getting words from sentences and add mean per word. But sometimes the result is not what I expected.
sentence : 我要九个
result : 我要 / 九个
exptected : 我 / 要 / 九 / 个
sentence : 吃三丸
result : "吃三丸"
expected : 吃 / 三 / 丸
Because I have a Chinese English dictionary like this.
我 : I
要 : want
九 : nine
个 :  individual; measure word for people or objects
But for current result I have to have all possible cases in the dictionary.
ex）一个， 两个，三个，四个，六个，七个，八个，九个，是个。。。。
Thanks in advance.
```
```text
 代码如下：
#encoding=utf-8
import jieba
import jieba.posseg as posseg

jieba.load_userdict("/data/paper/keywords/user_dict.txt")

words = posseg.cut("在现代社会中，免洗消毒液的使用越来越普遍，消毒液中所含的肉豆蔻酸异丙酯、丙二醇和乙醇等成分可增强皮肤的渗透性。")

for w in words:
    print('%s %s' % (w.word, w.flag))

报错
Building prefix dict from /usr/lib/python2.7/site-packages/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.296606063843 seconds.
Prefix dict has been built succesfully.
Traceback (most recent call last):
File "./jieba-test.py", line 5, in 
jieba.load_userdict("/data/paper/keywords/user_dict.txt")
File "/usr/lib/python2.7/site-packages/jieba/init.py", line 119, in wrapped
return fn(_args, *_kwargs)
File "/usr/lib/python2.7/site-packages/jieba/init.py", line 312, in load_userdict
add_word(_tup)
File "/usr/lib/python2.7/site-packages/jieba/init.py", line 116, in wrapped
return fn(_args, **kwargs)
TypeError: add_word() takes at most 3 arguments (4 given)
之前还工作的好好的。user_dict.txt 类似于这样的：
酞 10 n
 酞酐 20 n
 酞磺胺噻唑 50 n
 酞磺醋胺 40 n
 酞基 20 n
 酞己炔酯 40 n
 酞腈 20 n
 酞菁 20 n
 酞菁二(吡啶)铁(II)络合物 150 n
 酞菁镉 30 n
 酞菁蓝 30 n
 酞菁锂 30 n
 酞菁镁 30 n
 酞菁染料 40 n
 酞菁素艳蓝 50 n
 酞菁锡(II) 70 n
 酞菁亚铁 40 n
 酞菁铟 30 n
 酞菁银 30 n
 酞类染料 40 n
 酞嗪 20 n
 酞醛 20 n
 酞酸 20 n
 酞酸丁苄酯 50 n
 酞酸二苯酯 50 n
 酞酸二丙酯 50 n
 酞酸二环己酯 60 n
 酞酸二甲酯 50 n
 酞酸二壬酯 50 n
 酞酸二辛酯 50 n
 酞酸二乙酯 50 n
 酞酸二异癸酯 60 n
 酞酮酸 30 n
 酞醯胺醯基 50 n
 酞醯基 30 n
 酞醯亚胺基 50 n
 酞酰胺 30 n
 酞酰氯 30 n
 酞亚胺 30 n

其中优先级是按词语长度用脚本自动生成的，给予长词更高的优先级。
```
```text
 我发现了。似乎是自定义词库本身的问题
```
```text
 以下示例中 words 是默认词典中所有的词，feqs 是对应的词频：
freqs[words.index('北京市')]  # 3392
freqs[words.index('人民政府')]  # 15227
freqs[words.index('北京市人民政府')]  # 135

list(jieba.cut('北京市人民政府发布了一条政策'))
#  ['北京市人民政府', '发布', '了', '一条', '政策']
如上，北京市 和 人民政府 的词频要比 北京市人民政府 高出很多，但是 北京市 和 人民政府 却没有被切开，这是为什么呢？
```
```text
 注意源码中calc()函数：

def calc(self, sentence, DAG, route): 
N = len(sentence)

route[N] = (0, 0) 

logtotal = log(self.total) 

for idx in xrange(N - 1, -1, -1):

    route[idx] = max((log(self.FREQ.get(sentence[idx:x + 1]) or 1) - logtotal + route[x + 1][0], x) for x in DAG[idx])


最后一行的运算，实际上是log(单词词频) - log(全词词频) = log(单词词频 / 全词词频)，也即这个词与所有词词频之比。因此词数多时原概率会多除以一个分母即全词词频，概率上反而变小。此处并非单词词频简单相加。
个人理解，欢迎交流。
```
```text
 你能以我上面给的例子解释下吗？我看文档上说的是词频高就能被切开。
多多指教。
```
```text
 假设所有词词频之和为total（比如1000000）。当进行动态规划分词时，当进行“北”字子运算时，比较的是log(“北京市”->词频 / total) + log("人民政府"->词频 / total) = log("北京市"->词频 * “人民政府"->词频 / total^2)，和log("北京市人民政府"->词频 / total)。虽然“北京市”和“人民政府”词频相对高，但由于多除了一个全词词频之和的分母，概率运算的结果实际上是变小了的。据我的理解，实际上分词方法倾向于分割为字典中匹配的尽可能长的词。
另外，此处用log函数而非乘除法直接运算似乎是为了解决浮点数对于此处概率极小时不够精确的情况。
```
```text
 多谢指教，我再看看
```
```text
 python3.5，新建一个py，把给的第一个用法示例复制进去运行，就报这个错误。
```
```text
 我发现了，是起名字的问题，不要命名为jieba。。。
```
```text
 果然 我也遇到这个问题
```
```text
 我的文件名不是jieba也报这个错了，这是为什么
```
```text
 文件名不要用jieba
```
```text
 文件名不要用jieba,一改就好了，神奇
```
```text
 文件名也不要用tokenize.py, 否则也是同样的错误。
```
```text
 @mrkingsun 感谢回答，问题已经解决的话，请close该issue
```
## 273.在一个句子中出现两个相同的组合的词
Jan 12, 2016
```text
 如何将类似“abc123456”的文本切分为“abc”和“123456”？默认设置是切分不开的。
```
```text
 给代码里的几处re_eng戳个洞，分出来一类数字看看？
```
```text
 谢谢啦！
```
```text
 以下句子分词可改进：
句子：今天天气不错
jieba result: 今天天气    不错
期望结果：今天    天气     不错
句子：如果放到post中将出错。
jieba result: 如果 放到 post 中将   出错
期望结果：如果 放到 post 中    将    出错
```
```text
 fxsjy closed this in #351 on 16 Mar
```
```text
 In add_word, it should be:
freq = int(freq) if freq is not None else self.suggest_freq(word, False)
in order to be compatible to:
def del_word(self, word):
self.add_word(word, 0)
```
```text
 You can make a PR.
```
```text
 例如这样一句话：
我要听我要这首歌
我把 我要 这个词加到自定义词典中，但是开始的我和要也会 切分成 我要,
我想要的是 我/要/听/我要/这首歌这个要如何处理？
```
## 274.jieba.set_dictionary 會忽略詞性
Jan 12, 2016
```text
 No description provided.
```
```text
 用golang 实现了一下jieba, 在 https://github.com/pastebt/gieba , 欢迎指教
```
```text
 set_dictionary :
jieba.set_dictionary('data/dict.txt.big')
會有 #103 的問題發生
```
## 275.use jieba in sklearn.feature_extraction.text.TfidfVectorizer
Dec 30, 2015
```text
 jieba.add_word("石墨烯",100,"nr")
jieba.add_word("凱特琳",100,"nr")
jieba.add_word("莫那娄氏",10,"n")
jieba.del_word("自定义词")
不添加词典，利用add_word和del_word仍然不起作用
先开启并行，不起作用。如果先添加词，再进行并行就可以。。。。还真是bug
jieba.enable_parallel()
```
```text
 #547
```
```text
 Is it possible to use the jieba's tokenizer in TfidfVectorizer? [http://scikit-learn.org/stable/modules/generated/sklearn.feature_extraction.text.TfidfVectorizer.html].
sklearn.feature_extraction.text.TfidfVectorizer asks for a callable tokenizer, and I am wondering which function in jieba can be passed here.
I would like to cluster or classify some Chinese documents.
```
```text
 jieba.cut()
or
jieba.cut_for_search()
jieba_1_分词
```
```text
 tfidf_vectorizer = TfidfVectorizer(tokenizer=jieba.cut, lowercase=False, stop_words=stopwords)
```
## 276.(How) can I permanently change the POS tag for a word?
Dec 25, 2015
```text
 class Tokenizer(object):
    ..
    def suggest_freq(self, segment, tune=False):
        ..
        if tune:
            add_word(word, freq)
此处,add_word(word, freq)应为self.add_word(word, freq),否则更新的是默认分词器jieba.dt,而不是用户的new_dt = Tokenizer().
```
```text
 Every time I reboot the Jieba the setting I added using jieba.add_word(word, freq=None, tag=None) will be washed away. Can I keep it somewhere making the setting permanent?
```
```text
 You can write a custom dictionary, and then load the dictionary at start.
```
## 277.Efficiency of Jieba python?
Dec 22, 2015
```text
 In [94]: s = '乌鲁木齐爱家超市南门店'

In [95]: jieba.cut(s)
Out[95]: <generator object Tokenizer.cut at 0x10e24fde0>

In [96]: jieba.lcut(s)
Out[96]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [97]: jieba.add_word('南门')

In [98]: jieba.lcut(s)
Out[98]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [99]: jieba.suggest_freq('南门', True)
Out[99]: 833

In [100]: jieba.lcut(s)
Out[100]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [101]: jieba.lcut(s, HMM=False)
Out[101]: ['乌鲁木齐', '爱家', '超市', '南', '门店']
期望结果
['乌鲁木齐', '爱家', '超市', '南门', '店']
```
```text
 In [94]: s = '乌鲁木齐爱家超市南门店'

In [95]: jieba.cut(s)
Out[95]: <generator object Tokenizer.cut at 0x10e24fde0>

In [96]: jieba.lcut(s)
Out[96]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [97]: jieba.add_word('南门')

In [98]: jieba.lcut(s)
Out[98]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [99]: jieba.suggest_freq('南门', True)
Out[99]: 833

In [100]: jieba.lcut(s)
Out[100]: ['乌鲁木齐', '爱家', '超市', '南', '门店']

In [101]: jieba.lcut(s, HMM=False)
Out[101]: ['乌鲁木齐', '爱家', '超市', '南', '门店']
期望结果
['乌鲁木齐', '爱家', '超市', '南门', '店']

@4ft35t 设置的频率还是太低了，使用jieba.add_word("南门", freq=1000)设置高一点就可以了。
```
```text
 I wonder the efficiency of Jieba python edition.
I have a set of corpus around 1.5G size, with 3 of 4 cores parallel processing, using Jieba python and have waited for more than 2hr to finish the segment task.
Is this normal or anything wrong?
My platform is MBP '13 spring with 10G RAM, MacOS 10.9., python 3.4.2.
Thanks a lot.
```
## 278.f.name, lineno, line)) ValueError
Dec 17, 2015
```text
 原字典自带的dict.txt里面 '阳台' 的词频是 242 , 我在我的自定义字典里, 定义:  阳台栏杆 500  阳台天花 500
但是在分词的时候,依然会把阳台有关的词分词为:  阳台 / 栏杆  阳台/天花
```
```text
 运行demo出现这个。。。
Building prefix dict from /Library/Python/2.7/site-packages/jieba/dict.txt ...
Loading model from cache /var/folders/hv/rnxvykcd4dq98fm44qqh7_sm0000gn/T/jieba.cache
Loading model cost 0.584 seconds.
Prefix dict has been built succesfully.
Traceback (most recent call last):
File "test_userdict.py", line 6, in 
jieba.load_userdict('userdict.txt') #载入字典及字典文件名称
File "/Library/Python/2.7/site-packages/jieba/init.py", line 381, in load_userdict
f.name, lineno, line))
ValueError
```
## 279.運行中修改字典，重新load cache
Dec 17, 2015
```text
 在自己的代码中用suggest_freq(word)操作确保特殊词不被拆分, 但同时代码其他部分引用的外部模块也同时引用了jieba, suggest_freq(word)的操作会影响外部模块分词操作(并不是希望看到的). 对此是否提供清除suggest_freq的操作?
```
```text
 import importlib
importlib.reload(jieba) is okey...
```
```text
 by using importlib.reload method, you may encounter this scenario in iPython/jupyter:
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
...
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.

Once more you reload jieba module, there is one more line output in each section when jieba.dt.initialize() is implemented.
I'm not sure if there's any influence to other codes.
```
```text
 在cut句子以後，如果變更字典集，不會更新到cache，有什麼方法可以在initialize還可以重新建立一次cache嗎？
ex.
jieba.cut("我的靴子裡有蛇") #我/ 的/ 靴子/ 裡有/ 蛇
jieba.add_word("我的靴子")
jieba.cut("我的靴子裡有蛇") #還是... 我/ 的/ 靴子/ 裡有/ 蛇
```
## 280.建议优化词典内存结构
Dec 16, 2015
```text
 免费分享，造损免责。
打开默认词典（根目录）或自定义词典，把所有用来间隔词频和词性的空格间隔符改成@@
（选用@@是因为一般关键词里遇到这个分隔符的几率比较小吧）
继续，打开jieba根目录下init.py

搜索
re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._]+)", re.U)
改成
re_han_default = re.compile("(.+)", re.U)



搜索
re_userdict = re.compile('^(.+?)( [0-9]+)?( [a-z]+)?$', re.U)
改成
re_userdict = re.compile('^(.+?)(\u0040\u0040[0-9]+)?(\u0040\u0040[a-z]+)?$', re.U)



搜索
word, freq = line.split(' ')[:2]
改成
word, freq = line.split('\u0040\u0040')[:2]



补充：若用的全模式继续改。
搜索
re_han_cut_all = re.compile("([\u4E00-\u9FD5]+)", re.U)
改成
re_han_cut_all = re.compile("(.+)", re.U)
```
```text
 您好，我正好碰见这个问题，我想把 迈克尔·乔丹 不要拆开，试了您的方法，
将默认的dict改成了如下：
AT&T@@3@@nz
B超@@3@@n
c#@@3@@nz
C#@@3@@nz
c++@@3@@nz
将user_dict变成了如下：
小皇帝
科比布莱恩特
迈克尔·乔丹
init.py也按照您的方法改了，没用全模式所以没改最后一个
加载报错：
ValueError: invalid POS dictionary entry in C:\Tools\anaconda\lib\site-packages\jieba\dict_gai.txt at Line 1: AT&T@@3@@nz
请您指教！感激不尽！
```
```text
 @xusong123 我试了一下，没有问题啊，是不是用用了记事本打开txt导致加上了bom头。建议用notepad++创建编辑文本。
```
```text
 您好，我试了试，还是不行，方便的话我把这个贴上来给您看看
dict_gai3.txt
我加载的是dict_gai3.txt
```
```text
 @xusong123  你试一下，不用我的方法，然后把你字典里的点·删掉，然后看看报错不。如果也报错，就证明你后来新加的词，编码上有问题。
```
```text
 谢谢，已经解决，最好用python3
```
```text
 @xusong123 请问怎么解决的？我也出现同样的错误了，谢谢
```
```text
 我也遇到了这个问题。通常是/tmp/jieba.cache未即时更新引起的。
解决方法：删除jieba.cache，把默认字典（dict.txt）中的空格替换为@@即可。
@wonderfreer
```
```text
 我需要加载千万级的词典，词典才300M，但是加载后内存高达5G。
self.FREQ[word] = freq
self.total += freq
if tag:
self.user_word_tag_tab[word] = tag
for ch in xrange(len(word)):
wfrag = word[:ch + 1]
if wfrag not in self.FREQ:
self.FREQ[wfrag] =
FREQ
user_word_tag_tab
用了多个字典去存储数据，是不是浪费内存啊，这里可以合并成一个词典吧
```
```text
 @zhangzhenhu ， 可以试一试，发个PR吧 :-)
```
```text
 @fxsjy 128m内存直接跑不起来，256的才勉强能跑。
我用的bae。
```
```text
 是的，我也遇到一样的问题。加载词典太耗内存了！
```
## 281.建议词典分隔符空格改成\t
Dec 16, 2015
```text
 `



seg_list = jieba.cut("他来到了网易杭研大厦")  # 默认是精确模式
print(", ".join(seg_list))
他, 来到, 了, 网易, 杭研, 大厦`



但是在分词过程中判断词性
`



words = pseg.cut("来到了网易杭研大厦")
for word, flag in words:
...  print('%s %s' % (word, flag))`



依旧会把“杭研”分为两个词  杭 j 研 vn  ；
如何判断“杭研”是否为新词？
```
```text
 可以參考我修正的錯誤 #670
```
```text
 用demo里的程序就没有问题，自己新建一个python文件，复制里面的import，执行时总是提示ImportError: No module named analyse
什么原因
C:\Python27\Lib\site-packages\jieba-0.38-py2.7.egg-info
jieba-0.38-py2.7.egg-info\jieba
都存在
```
```text
 应该是包结构的问题，看看你init.py是否正确，试试不用没有包结构的情况
```
```text
 很多专有词语中间是有空格的
例如：OpenGL ES
词典分隔符用空格不太合理，建议改成\t
```
```text
 因为多义词的原因，建议支持词典中支持多词性，方便后续设计更复杂的算法
```
## 282.分词的问题
Dec 15, 2015
```text
 我尝试对恒大相关的文章进行分词，结果错误率奇高，结果如下
'''
按照原计划，凯赛尔在西班牙学习三年后就将回国，但对于志向高远的“凯撒”来说，他更希望凭借未来三年的表现能留在西班牙继续深造，从而拉开自己的职业生涯，“**/随恒大/足校在西班牙学习三年后，我希望能留在这里，并开启自己的职业生涯
保监会再发狠/招恒大/人寿被禁止投资股票
这些/对恒大/概念股的影响有多大
然后在10月底三季报披露散户看/到恒大/**买了4.95%是不是要举牌了，许老板就把货都卖给你们了
'''
我将自定义词典中”恒大“的词频调整到了10000都没有任何变化，尝试jieba.suggest_freq('恒大', True)也没有用；尝试将HMM关掉，结果恒大这个词会被一直分成”恒/大“。
```
```text
 能不能再切之前，直接添加：
jieba.add_word("恒大")
```
```text
 我想把只有在指定的字典有的关键词提出来。
import jieba
jieba.initialize()
jieba.set_dictionary('mydic.txt')
但是出来的词还是包含一堆不在字典中的。请问我该怎么做？谢谢
```
```text
 如果 jieba.analyse.extract_tags 能只返回指定字典里的就更好了
```
## 283.分词错误
Dec 14, 2015
```text
 我想把新的词汇保存到txt文档里，请问大神怎么做？？？
```
```text
 No description provided.
```
```text
 文本（'构建了以技术为主的体系'）的分词选择jieba.posseg.cut(data,HMM=True)和jieba.cut(data,HMM=True)结果不一样，是bug吗？求修复
另外posseg/init.py 237行elif re_eng.match(x):对吗？
```
## 284.用户词典出现错误
Dec 7, 2015
```text
 在使用pseg.cut来获取分词词性的时候，发现“云计算”、“石墨烯”这些词语的词性居然是x，x我看介绍通常是一些标点符号，为什么会出现这种情况，还有其他词语也会有这种情况吗？
```
```text
 因为云计算，石墨烯是新词，词典里找不到，估计就默认设置成词性x了，你需要自己自定义词典并且加入这些新词的词性。
```
```text
 在使用用户词典时，运行样例，若是有词包含空格，例如样例中的
Edu Trust认证 2000
则运行时会出现如下的错误：
Prefix dict has been built succesfully.
Traceback (most recent call last):
File "test.py", line 6, in 
jieba.load_userdict("userdict.txt")
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 381, in load_userdict
f.name, lineno, line))
ValueError
```
```text
 pypi 上的版本没更新
```
## 285.你们的demo挂了
Dec 6, 2015
```text
 比如我有一句话，"2015年度本公司营业利润较2014年度减少-3480.83万元，降幅为-26.95%，主要原因系国内宏观经济增速放缓，下游市场需求不足，市场景气程度有所降低，公司销售收入整体下降。"
我先对结巴做了add_word设置
jieba.add_word("-26.95%")
jieba.add_word(u"-3480.83万元")
然后分词结果是：
2015年度|本|公司|营业利润|较|2014年度|减少|-|3480.83万元|，|降幅|为|-|26.95%|，|主要|原因|系|国内|宏观经济|增速|放缓|，|下游|市场需求|不足|，|市场|景气|程度|有所|降低|，|公司|销售收入|整体|下降|。|
带符号的就不行
```
```text
 遇到同样问题。
加入用户字典    test-test 词 依然会被切分成 test/-/test。
提高词频依然没用，不知道是不是 -  会被固定切出来，不收其他影响。
```
```text
 同样的问题，比如：
jieba.add_word("弗里德里克·泰勒",freq=888888); jieba.add_word("冯·诺伊曼",freq=888888); jieba.suggest_freq(('弗里德里克·泰勒', '冯·诺伊曼'), True); txt = "这个外国人叫弗里德里克·泰勒，是美国前国防部部长。冯·诺伊曼是公认的计算机发明人"; print(" | ".join(jieba.cut(txt, HMM=True, cut_all=False)))
结果还是：
'这个 | 外国人 | 叫 | 弗里德里克 | · | 泰勒 | ， | 是 | 美国 | 前 | 国防部 | 部长 | 。 | 冯 | · | 诺伊曼 | 是 | 公认 | 的 | 计算机 | 发明人'
大家有什么好办法？
```
```text
 有同样的需求，需要保证指定的词绝对不分。不知该怎么做。
```
```text
 代码里面有如下的正则表达式，用来提取中文

re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&._%]+)", re.U)
也就是说对于”弗里德里克·泰勒”中的“·”不认为是中文，而是作为类似逗号和句号来处理。

一个不太好但是可以用的办法就是，修改上面的正则表达式，将“·”加入。其中\xb7就是“·”。

re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&.\xb7%]+)", re.U)

测试结果：
》 这个 | 外国人 | 叫 | 弗里德里克·泰勒 | ， | 是 | 美国 | 前 | 国防部 | 部长 | 。 | 冯·诺伊曼 | 是 | 公认 | 的 | 计算机 | 发明人
```
```text
 小明nr 比p 小红nr 可爱v ，x 但c 都d 没有v 珍珍n 可爱v
```
```text
 如题
```
## 286.jieba.lcut 文件怎么用法
Nov 28, 2015
```text
 请问add_word(word, freq=None, tag=None) 这个函数增加的词语只能在当前的python jieba 编译环境下生效，python退出重新登陆时，修改的无效。
现在我需要永久性新增或者删除词语，如何实现永久性热更新操作？
谢谢
```
```text
 add_word()只添加到内存中，jieba好像没有提供固化更新字典的实现。但字典格式好像是通用的，可以用一些其他字典生成工具配套使用。
```
```text
 希望支持数据库词库..
```
```text
 支持这样 cut = jieba.lcut(text.txt) 输入文件的方式吗？试了下，说 NameError: name 'text' is not defined，可这文件同目录下面有的。
```
```text
 lcut 里的是字符串，把文件读进来再用
```
## 287.Failed building wheel for jieba
Nov 24, 2015
```text
 jieba.analyse.textrank()方法提取英文文本的关键词没有输出数据，但是jieba.analyse.tfidf()方法则可以处理英文文本。中文文本两个都可以处理。so ,问题是，，，
```
```text
 你看看源代码就知道了，tokenizer和postokenizer的区别。
```
```text
 re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&._]+)", re.U)
既然对C++, C#多了特判, 为什么不加入 objective-c特判
```
```text
 成功安装，但是报错信息Failed building wheel for jieba，可能是这个错误的原因，jieba模块不能正常运行
```
```text
 给详细日志啊
```
```text
 D:\learn-python>pip install jieba
Collecting jieba
Using cached jieba-0.37.zip
Building wheels for collected packages: jieba
Running setup.py bdist_wheel for jieba
Complete output from command c:\users\administrator\appdata\local\programs\python\python35-32\python.exe -c "import setuptools;file='C:\Users\ADMINI1\AppData\Local\Temp\pip-build-ibpo4da7\j
ieba\setup.py';exec(compile(open(file).read().replace('\r\n', '\n'), file, 'exec'))" bdist_wheel -d C:\Users\ADMINI1\AppData\Local\Temp\tmpf6wrfjs2pip-wheel-:
Traceback (most recent call last):
File "", line 1, in 
UnicodeDecodeError: 'gbk' codec can't decode byte 0x80 in position 100: illegal multibyte sequence

Failed building wheel for jieba
Failed to build jieba
Installing collected packages: jieba
Running setup.py install for jieba
Successfully installed jieba-0.37
给这个你看行不行
```
```text
 感觉是 pip 的锅：终端是 GBK 编码，setup.py 是 UTF-8，然后读入直接运行
你的运行环境是什么？（我打算去报 bug）
```
```text
 windows7 64位系统
D:\learn-python>python --version
Python 3.5.0
```
```text
 这个现在有解决方案吗
```
## 288.能提供一下统计指定关键词词频的接口吗？
Nov 23, 2015
```text
 如題，TFIDF算法一般來說topK這個參數要取多少比較合適?我目前使用的題材是金庸的小說，目前正在切神鵰俠侶這本。
```
```text
 遇到了一个奇怪的情况，我在Ubuntu16.04下通过pip安装了jieba0.39，可以正常使用，但后来可能又装了一些第三方的库，然后jieba.lcut就报错了，显示没有这个方法，由于后来装的第三方库比较多，正在生产环境上使用，不能直接卸载，感觉这个问题是因为某个第三方库中引入了jieba旧版本的代码，以至于我现在只要使用jieba，用的旧的代码，但通过pip list查询，jieba的版本还是0.39，请问应该怎样解决？谢谢~
```
```text
 我也遇到这个问题了，我是因为同时装了jieba3k，换成jieba就可以了，jieba.lcut只有在jieba中可以使用
```
```text
 No description provided.
```
## 289.Installation failure
Nov 22, 2015
```text
 比如某个句子的分词效果，结巴的结果和预期有出入
我想把人工标注的结果加入到结巴，如何？
```
```text
 非中英文的詞，如Citroën會被切成Citro+ë+n。請問可以如何避免？
看似系統將ë等特殊字看作符號。我有機會修改source code，將原本的英數字集加入此些特殊字嗎？目前看到相關的設定是jieba/posseg/init.py檔案的re_skip_detail = re.compile("([.0-9]+|[a-zA-Z0-9]+)")。
或是有其它建議？
謝謝。
```
```text
 当我尝试使用“pip install jieba"命令来安装它时，我无法正确安装，日志中显示为:"error: could not create '/usr/local/lib/python2.7/dist-packages/jieba': Permission denied",请问，我应该怎样来解决这个问题？？？使用的系统是Ubuntu14.04版本！
```
```text
 Chmod file with privilege
Sent from my iPhone

On 22 Nov 2015, at 20:53, hippieZhou notifications@github.com wrote:
When I use "pip install jieba",I could't install it ,the log shows that:"error: could not create '/usr/local/lib/python2.7/dist-packages/jieba': Permission denied",what should i do ???
—
Reply to this email directly or view it on GitHub.
```
```text
 Thank you !
```
## 290.bug in del_word
Nov 19, 2015
```text
 如何将类似“abc123456”的文本切分为“abc”和“123456”？默认设置是切分不开的。
```
```text
 给代码里的几处re_eng戳个洞，分出来一类数字看看？
```
```text
 谢谢啦！
```
```text
 以下句子分词可改进：
句子：今天天气不错
jieba result: 今天天气    不错
期望结果：今天    天气     不错
句子：如果放到post中将出错。
jieba result: 如果 放到 post 中将   出错
期望结果：如果 放到 post 中    将    出错
```
```text
 fxsjy closed this in #351 on 16 Mar
```
```text
 In add_word, it should be:
freq = int(freq) if freq is not None else self.suggest_freq(word, False)
in order to be compatible to:
def del_word(self, word):
self.add_word(word, 0)
```
```text
 You can make a PR.
```
## 291.命令行输入输出的方式，可以标注词性吗
Nov 17, 2015
```text
 我想把新的词汇保存到txt文档里，请问大神怎么做？？？
```
```text
 No description provided.
```
```text
 文本（'构建了以技术为主的体系'）的分词选择jieba.posseg.cut(data,HMM=True)和jieba.cut(data,HMM=True)结果不一样，是bug吗？求修复
另外posseg/init.py 237行elif re_eng.match(x):对吗？
```
```text
 另外如果我想使用 php 调用结巴分词的话是不是只能生成文件然后相互传递了……
```
```text
 词性标注功能已加上。结巴有 PHP 版本 https://github.com/fukuball/jieba-php ；至少你可以用管道吧
```
```text
 感谢大神！！
```
## 292.另外问一下jieba可以分句吗？
Nov 16, 2015
```text
 No description provided.
```
```text
 我想用wiki做自定义词库，已经用gensim训练出一个模型。
不知道用wiki是否比人民日报的效果好？
不知道自定义词库格式要求是怎么样的？
有没有现成的工具可以制作词库？
```
```text
 文档中可以jieba支持自定义加载用户词库的功能，而且自定义的词库的格式也比较简单，我自己是写了一个python脚本来建立自己的词库的。
```
```text
 谢谢 @super1-chen
请问用了user_dict是否会覆盖原来的默认词库？
```
```text
 No description provided.
```
```text
 No description provided.
```
## 293.词性标注的分词和Tokenize分词不一致？
Nov 16, 2015
```text
 我在py2.7下
re_han = re.compile("([\u4E00-\u9FD5]+)")对中英文混合分词时出现问题
比如输入 中国tfboy说唱,篱笆女人等。
re_han.split结果为[u‘中国’，u'tfbo' ，u'y说唱', ....]
将re_han 改为 re.compile(u"([\u4E00-\u9FD5]+)")能正常分组
```
```text
 个人在做简历分类，请问大家，有适合的中文文本分类的停用词词典么
```
```text
 A possible one if you really need it in NLP task: https://github.com/kn45/seg-jb/blob/master/segjb/stopwords.dat
```
```text
 https://github.com/hankcs/HanLP/blob/master/data/dictionary/stopwords.txt
```
```text
 import jieba.posseg as pseg
words = pseg.cut(text)



result = jieba.tokenize(text)
发现对text的分词有些不同，比如“抄的”在前面就是一个词，第二种就是2个词.请问哪里出了问题？
```
```text
 我也发现这个问题了，没人回复啊
```
```text
 @jiffies , poseg和jieba.cut使用的模型不一样，所以分词的效果有差别。
```
```text
 哪个好呢？
```
## 294.关于TFIDF的几个问题
Nov 13, 2015
```text
 通过 jieba.set_dictionary("."+os.sep+"dict.txt") 指定dict.txt 文件位于一个带中文的父目录时，使用jieba.initialize() 进行初始化报错，报错为：
Traceback (most recent call last):
File "", line 26, in 
File "F:\Qiyi\build\getcalldatafrom_txt\out00-PYZ.pyz\jieba", line 111, in initialize
UnicodeDecodeError: 'ascii' codec can't decode byte 0xcc in position 3: ordinal not in range(128)
定位于语句： jieba_init.py_ 中的 111行： default_logger.debug("Building prefix dict from %s ..." % (abs_path or 'the default dictionary'))
将程序至于全英文目录时，运行正常。
```
```text
 我在使用jieba分词时，遇到了如下情况：
待分词句子：奥布瑞·普拉扎（Aubrey Plaza），1984年6月26日出生于美国特拉华州威尔明顿，演员。
分词结果：
jieba.cut:
奥/布/瑞/·/普拉/扎/（/Aubrey/ /Plaza/）/，/1984/年/6/月/26/日出/生于/美国/特拉华州/威尔/明/顿/，/演员
posseg.cut:
奥 nr 布 nr 瑞 ns · x 普拉 nrt 扎 v （ x Aubrey eng   x Plaza eng ） x ， x 1984 eng 年 m 6 eng 月 m 26 eng 日出 v 生于 v 美国 ns 特拉华州 ns 威尔 nrt 明 a 顿 q ， x 演员 n
出生这个词一直无法正确分出来，我发现词典中已经包含了该词：出生 1979 v，而日出频率低于出生：日出 394 v；然后我尝试自己再一次将“出生”添加到词典，以及关闭HMM，均没有作用，请问这是什么问题呢，谢谢！
```
```text
 找到了原因，日出的频率是394，出生频率是1979，生于的频率是4690，导致分为日出/生于
解决办法：
jieba.del_word('日出')
jieba.add_word('出生于')
jieba.add_word('日出')
```
```text
 你好。
在 /jieda/init.py 的 gen_trie(f_name) 函数中，目前只是简单地读出所有文件内容，然后以 \n 切分，然后再对每一行切片。
这样会有一个问题，当碰到一个空行时，程序就会出错（split(' ') 之后的三个变量赋值会不匹配），并且，这个错误处理起来还比较麻烦（事实上我改了一下默认的字典文件后就出错了，我也找不出字典哪里有问题）。
我建议是不是对空行处理一下，比如：
def gen_trie(f_name):
    lfreq = {}
    trie = {}
    ltotal = 0.0

    with open(f_name, 'rb') as f:
        for l in f:
            l = l.decode('utf-8').strip()
            if not l:
                continue
            word, freq, _ = l.split(' ')
            freq = float(freq)
            lfreq[word] = freq
            ltotal+=freq
            p = trie
            for c in word:
                if not c in p:
                    p[c] ={}
                p = p[c]
            p['']='' #ending flag

    return trie, lfreq,ltotal
```
```text
 嗯，可以考虑。不过现在加载词典已经很慢了，我需要测试加上这些条件判断后，加载速度会不会下降。
```
```text
 默认词典直接写成.py文件会不会快一些
```
```text
 @xluer ，现在的机制是第一次从文本加载，然后会用marshal.dump把词典序列化到临时磁盘文件。以后再启动时，会比较dump文件和词典文本文件的时间戳，如果dump文件更新，就不会读取文本文件了。
为什么用文本文件？-- 最初的想法是便于用于修改词典文件。
```
```text
 不知道有什么通用的方法来生成语料库？
我用几个语料连接到了一起，通过extract_tags生成了权重，作为IDF来使用，但似乎不好使。
IDF中没有的词是否可以识别？
我看如果是新词，指定了IDF后似乎都不能extract_tags出来。
```
```text
 好吧，看了下算法，我偷懒了。
不过还是比较希望有独立的IDF生成实现
```
```text
 可以参考一下 nltk 的实现，先计算所有词的 IDF
```
```text
 @gumblex 参考另外一个gist实现了一个TFIDF，等下看看效果
```
```text
 @geekan 求idf实现方法
```
## 295.怎么自动过滤停用词和标点符号?
Nov 6, 2015
```text
 您好，每一次我用jieba进行分词的时候，都会有
Building prefix dict from the default dictionary ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.128 seconds.
Prefix dict has been built succesfully.

这样的提示。怎么去除这些提示呢？
```
```text
 找到jieba库的__init__.py， 大约28~30行左右，
log_console = logging.StreamHandler(sys.stderr)
default_logger = logging.getLogger(__name__)
default_logger.setLevel(logging.DEBUG)
default_logger.addHandler(log_console)

将default_logger.setLevel(logging.DEBUG) 改为 default_logger.setLevel(logging.INFO) 试试。
```
```text
 @oisc it shows error:
Traceback (most recent call last):
  File "xxx.py", line 29, in <module>
    jieba_logger.handlers.clear()
AttributeError: 'list' object has no attribute 'clear'
```
```text
 @IvyGongoogle
try del jieba_logger.handlers[:] if you are using Python2
```
```text
 @aqua7regia it works. thank you~~
```
```text
 jieba.setLogLevel(20)
或
import logging
jieba.setLogLevel(logging.INFO)
```
```text
 @oisc  sorry, What do you mean for del jieba_logger.handlers[:]?
```
```text
 @gaoyangthu it works, but It seems to be jieba.setLogLevel(logging.INFO)
```
```text
 @IvyGongoogle You're right!
setLogLevel(0) equals setLogLevel(logging.NOTSET)
setLogLevel(10) equals setLogLevel(logging.DEBUG)
setLogLevel(20) equals setLogLevel(logging.INFO)
setLogLevel(30) equals setLogLevel(logging.WARNING)
setLogLevel(40) equals setLogLevel(logging.ERROR)
setLogLevel(50) equals setLogLevel(logging.CRITICAL)
```
```text
 @gaoyangthu  yes, thank you very much~~
```
```text
 提取的内容，方便排除部分标点，英文和数字么，需要调整一下吗？
```
```text
 你可以看看test.py的样本和结果
lizhengfu notifications@github.com编写：

提取的内容，方便排除部分标点，英文和数字么，需要调整一下吗？
—
Reply to this email directly or view it on GitHub.￼
```
```text
 @youyudehexie , 你是说stopwords？ 现在还没有提供接口，你可以先试一试效果：https://github.com/fxsjy/jieba/blob/master/test/extract_tags.py
```
```text
 THX??ֹͣ?? ?? ???ˣ???Ϊ?ҷ????????δ??뷢?ֻ???ȡ?????ֺͳ??ȴ???1?ı??㣬?Լ???????д??һ?£??ҵ??????ǣ???ȡ??ʱ?????????ֺͲ??ֱ??㣬?ǳ??????Ƶ?˼·????????ʱ?????޸ġ?
------------------ ԭʼ?ʼ? ------------------
??????: "Sun Junyi"notifications@github.com;
????ʱ??: 2013??5??25??(??????) ????9:09
?ռ???: "fxsjy/jieba"jieba@noreply.github.com;
????: "lizhengfu"405574395@qq.com;
????: Re: [jieba] ???ͷִʵĹؼ?????ȡ (#54)
@youyudehexie , ????˵stopwords?? ???ڻ?û???ṩ?ӿڣ???????????һ??Ч????https://github.com/fxsjy/jieba/blob/master/test/extract_tags.py
??
Reply to this email directly or view it on GitHub.
```
```text
 @youyudehexie ，乱码？
```
```text
 THX，停止词 和 过滤，因为我发现用这段代码发现会提取到数字和长度大于1的标点，自己简单重写了一下，我的疑问是，提取的时候，有数字和部分标点，是出于设计的思路，还是暂时不想修改。
```
```text
 @youyudehexie , 的确是没有过滤数字和长度大于1的标点。 有些时候数字也会成为关键词吧，比如911。标点的确是应该过滤，你是怎么修改的？可以给我发一个pull request。
这是我目前的实现，其实很简单：
https://github.com/fxsjy/jieba/blob/master/jieba/analyse/__init__.py
```
```text
 @fxsjy 不知道你现在做了标点没有，
刚刚感觉标点不是个问题。。。直接替换成空格就好，长度大于2的空格再替换成长度为1的空格。
但是试了一下发现用正则表达式不对，\D连汉字都匹配上了。。。
最近没时间做一个分支T^T
查了一下
正则表达式的[[:punct:]]可以匹配标点符号，我只在编辑器里试了一下。。。re模块好像要调整一下什么的，，
```
```text
 @zjjott ,[[:punct:]] 不能搞定汉语标点符号吧？
```
```text
 那我继续测试一下。。。。
希望早日可以写个pull。。。
2013/6/14 Sun Junyi notifications@github.com

@zjjott https://github.com/zjjott ,[[:punct:]] 不能搞定汉语标点符号吧？
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/54#issuecomment-19439691
.


朱涧江
微笑的猪头，帅气非凡，23333
果壳网，科技有意思哟，23333
www.guokr.com
```
```text
 No description provided.
```
```text
 你可能用yield過濾。比如說：
def 過濾(input_text):
    for token in jieba.cut(input_text):
        if token not in ",.?;'[]()`~!@#$%^&*/+_-=<>{}:，。？！·；：‘“、\"":
            yield token
```
## 296.英文句号的错误
Oct 29, 2015
```text
 在我的自定义词典中有这样一个专有名词"ens"，在分词时却将“license”强行分成了“lic/ens/e”。这个要怎么处理这种情况？
```
```text
 在词典里加上“license”，再给个比“ens”较高的数
```
```text
 我觉得jieba应该尊重英文词的边界,当自定义词典里起始或结束是英文字母,例如:



jieba.add_word('w底')
jieba.lcut("太好了w底出现了")  # 这里"w底"的w不是英文词的连续,可以切分
['太好了', 'w底', '出现', '了']
jieba.lcut("wow底出现了") # 这里"w底"的w是英文词的连续,不应该切分
['wo', 'w底', '出现', '了']
```
```text
 @chunsheng-chen 那是不是得有个字典存英文单词，感觉会很大哦
```
```text
 不知道具体细节，但我猜测jieba对英文词的分解是基于类似"[a-zA-Z0-9]*"的模式，所以不需要英文字典，例如:



jieba.lcut("this is a 1test1-abc2! call 911")
['this', ' ', 'is', ' ', 'a', ' ', '1test1', '-', 'abc2', '!', ' ', 'call', ' ', '911']



如果能尊重英文词的自然分割方式，就不会出现上面的情况了: license是一个完整的词，wow是一个完整的词。
```
```text
 自定义词典中部分词含有日文的假名，但是分词出来好像全部无视了。
```
```text
 我参看了ictclass关于词性标注的介绍[1]：标点符号都应该标注为以“w”开头的字符串，而jieba目前将它们标记为“x”，即认为是“字符串”。
代码：
list(pseg.cut(u"今天的任务有四项：写程序、看电影和跑步。"))
结果：
[今天/t, 的/uj, 任务/n, 有/v, 四项/m, ：/x, 写/v, 程序/n, 、/x, 看/v, 电影/n, 和/c, 跑步/n, 。/x]
请问这是bug么？
[1] http://ictclas.org/docs/ICTPOS3.0%E6%B1%89%E8%AF%AD%E8%AF%8D%E6%80%A7%E6%A0%87%E8%AE%B0%E9%9B%86.doc
```
```text
 另外，jieba还使用了标签“j”，而“j”并没有出现在ictclas的标准中。
```
```text
 请问uj中的j表示什么含义？
```
```text
 我在windows server 2008的服务器上pip install jieba（jieba-0.32），出现了如下错误：
Installing collected packages: jieba
Running setup.py install for jiaba
Sorry: MemoryError: ( )
Sorry: MemoryError: ( )
Successfully installed jieba
Cleaning up......
这样带来的问题就是，当我用django-haystack + whoosh + jieba的时，用rebuild_index，会出现MemoryError ，而无法生成索引。
我的是python 2.7，希望得到解决，谢谢
```
```text
 实测原来是服务器内存不够，升级服务器内存之后，上述问题均得到了解决
```
```text
 当文中出现英文句号时会被当作小数点处理，分词和相应的词性标注会出现错误。
如"1.hello.2.word."
另外“了解Java7.0.”也会出现分词错误。
```
## 297.AttributeError: 'module' object has no attribute 'cut'
Sep 24, 2015
```text
 如題，TFIDF算法一般來說topK這個參數要取多少比較合適?我目前使用的題材是金庸的小說，目前正在切神鵰俠侶這本。
```
```text
 遇到了一个奇怪的情况，我在Ubuntu16.04下通过pip安装了jieba0.39，可以正常使用，但后来可能又装了一些第三方的库，然后jieba.lcut就报错了，显示没有这个方法，由于后来装的第三方库比较多，正在生产环境上使用，不能直接卸载，感觉这个问题是因为某个第三方库中引入了jieba旧版本的代码，以至于我现在只要使用jieba，用的旧的代码，但通过pip list查询，jieba的版本还是0.39，请问应该怎样解决？谢谢~
```
```text
 我也遇到这个问题了，我是因为同时装了jieba3k，换成jieba就可以了，jieba.lcut只有在jieba中可以使用
```
```text
 No description provided.
```
```text
 用的pip3 install 安装的
后来又把github源码保存在文件夹下执行
但总是报错
mac系统
```
```text
 有出错具体位置的，不然没法定位
```
```text
 嗯 后来把源码放到文件夹里就好了 不知道为什么pip3 install不行
```
```text
 pip3 install jieba3k
again 请贴出 bug 出错详情而不是简述
```
```text
 根据 README 的说明，应当是 pip3 install jieba 而不是 jieba3k 吧……
```
```text
 好奇怪
```
```text
 file name should not be jieba.py, you can't import yourself(jieba.py)
```
```text
 i know，thank you for your reply
            On 02/07/2017 18:26, luotao wrote: file name should not be jieba.py, you can't import yourself(jieba.py)

—You are receiving this because you commented.Reply to this email directly, view it on GitHub, or mute the thread.








{"api_version":"1.0","publisher":{"api_key":"05dde50f1d1a384dd78767c55493e4bb","name":"GitHub"},"entity":{"external_key":"github/fxsjy/jieba","title":"fxsjy/jieba","subtitle":"GitHub repository","main_image_url":"https://cloud.githubusercontent.com/assets/143418/17495839/a5054eac-5d88-11e6-95fc-7290892c7bb5.png","avatar_image_url":"https://cloud.githubusercontent.com/assets/143418/15842166/7c72db34-2c0b-11e6-9aed-b52498112777.png","action":{"name":"Open in GitHub","url":"https://github.com/fxsjy/jieba"}},"updates":{"snippets":[{"icon":"PERSON","message":"@wall2flower in #299: file name should not be `jieba.py`, you can't import yourself(`jieba.py`)"}],"action":{"name":"View Issue","url":"#299 (comment)"}}}
```
## 298.带逗号的词语切分
Sep 7, 2015
```text
 Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/t8/p9fyd38n58v35j4fxrndsyyw0000gn/T/jieba.cache
Loading model cost 0.357 seconds.
Prefix dict has been built succesfully.
```
```text
 jieba 能用于elasticsearch吗？
```
```text
 我用淘宝一个搜索词的所有商品名字组成一个字符串，交给jieba 切出来的都是一个个的字。用正常的句子试了试也是一样。用的是pip install jieba下载的
```
```text
 请问有没有带标点符号的词语切分方法，比如我需要保留逗号切割“智者千慮，必有一失”，有没有可能完整保留逗号？
已经试过自定义词典和add单个词语，似乎效果并不明显。
请问大家有没有什么好的想法？
```
```text
 可以考虑一下最大匹配算法
```
```text
 类似的问题还有，如何保持空格分隔的词，比如"Steve Jobs".
```
## 299.IP地址的問題
Aug 30, 2015
```text
 java版本的分词Ansj和HanLP都支持用自己的语料再训练，为什么结巴分词这么久也不支持？
```
```text
 No description provided.
```
```text
 我只能告诉你这玩意看得懂空格和标点……（
```
```text
 用 https://github.com/moses-smt/mosesdecoder/blob/master/scripts/tokenizer/tokenizer.perl
```
```text
 楼主，一直以来用你的jieba分词，感谢你无私的分享！
近来发现一个问题：
for i in jieba.analyse.extract_tags('上海一日游攻略'):
print i
输出结果：
攻略
上海
我看到idf.txt中‘一日游’与‘攻略’这个词的IDF相差无几，为什么没‘一日游’这个词呢？
不太明白其中原理。
```
```text
 @bobo1732 , 我这里没有复现此bug，请问你的版本是多少？
```
```text
 @fxsjy 我是用的0.35版本
```
```text
 文本中有IP地址，如127.0.0.1，已經用suggest_freq添加到詞典，仍然無法分出。測試代碼：
print list(jieba.cut('127.0.0.1 is ip address'))
print jieba.suggest_freq('127.0.0.1', True)
print list(jieba.cut('127.0.0.1 is ip address'))

輸出：
[u'127.0', u'.', u'0.1', u' ', u'is', u' ', u'ip', u' ', u'address']
4
[u'127.0', u'.', u'0.1', u' ', u'is', u' ', u'ip', u' ', u'address']
```
## 300.搜索问题
Aug 26, 2015
```text
 No description provided.
```
```text
 我是用正则表达式处理的，new_sentence = re.sub(r'[^\u4e00-\u9fa5]', ' ', old_sentence) 然后再进行分词的, \u4e00-\u9fa5这个是utf-8中，中文编码的范围
```
```text
 @cbzhuang  非常谢谢你的回复！ 我用了这个，不知道可对。#169
```
```text
 Actually, CJK characters are encoded together so there's no critical range for Chinese characters. A punctuation dict could be used to do the filtering.
```
```text
 db中有数据条“酒”，“白酒”，"红酒"
搜索“酒" ,只能出现一条记录”酒"
希望能同时出现“白酒”,"红酒"
```
```text
 这种问题不归这里管吧
```
## 301.期待Golang版的！
Aug 25, 2015
```text
 Please add the LICENSE file to the MANIFEST.in in both jieba and jieba3k so that it's included in the source distribution.
```
```text
 之前看过你写的goseg，不过好像很久没更新了，期待Golang版的jieba
```
```text
 早就有了
https://github.com/yanyiwu/gojieba
```
## 302.不能删除自定义word
Aug 24, 2015
```text
 用户自定义词典定义了中文和英文结合的词 ，分词没用效果
比如词典里定义了一个a字群
分词好像没有效果
```
```text
 import jieba
a=jieba.cut("阿里妈妈1");print(" ".join(a))
Building prefix dict from /usr/lib/python2.6/site-packages/jieba-0.37-py2.6.egg/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.256 seconds.
Prefix dict has been built succesfully.
阿里 妈妈 1
jieba.add_word('阿里妈妈1',100,'nz')
a=jieba.cut("阿里妈妈1");print(" ".join(a))
阿里妈妈1
jieba.del_word('阿里妈妈1')
a=jieba.cut("阿里妈妈1");print(" ".join(a))
阿里妈妈1
```
```text
 这个问题有解决吗?
```
## 303.hadoop load userdict
Aug 13, 2015
```text
 No description provided.
```
```text
 在Hadoop上怎么加载自定义字典啊
```
## 304.IOError occured when jieba submitted to spark cluster
Aug 10, 2015
```text
 根据不同的参数名称调用不同的自定义词典和stopwords来分词,自定义词典和stopwords全部初始化加载到内存中,实际分词的时候根据参数来替换dict和stopwords的变量值,这样能实现吗?
```
```text
 I use spark cluster to seg word and submit the program with jieba to the cluster as following:
/home/hadoop/spark-1.2.0-bin-hadoop2.4/bin/spark-submit --master spark://hz0124:7077 --driver-memory 1g --executor-memory 5g --total-executor-cores 60 --py-files ../sparkcode/dependency/dependency.zip seg
word_test.py
the jieba is zipped as the file ../sparkcode/dependency/dependency.zip
then an error occured:

Anyone meets this problem? how to deal with it?
```
```text
 Using zip will cause problems to read the dictionary because it uses regular file IO.
You can unzip the dependencies, or put the dict.txt elsewhere (not in a zip) and then manually initialize jieba with jieba.initialize('/path/to/dictionary').
```
## 305.《》标题符号包围的电影名，电视名能控制直接拆分成一个词吗
Jul 29, 2015
```text
 如题，多谢！
```
```text
 No description provided.
```
```text
 用正则直接匹配可不可以？
```
```text
 正则是可以，如果词库中包含电影电视名 估计能直接拆分出来
```
```text
 用 jieba.add_word
```
## 306.tfidf
Jul 29, 2015
```text
 你好，在做基于情感词典的情绪分析，结果发现，结巴分词，会把很多词合在一起。比如：你真笨；
我想要的结果是：你，真，笨，
但是返回的却是：你，真笨
还有：你太笨，返回的也是：你，太笨
感觉很不科学呐。真和太明明是一种程度词，怎么和笨合成一个词了。
同时调研了其它多个分词开源工具，发现只有你们会合在一起。
```
```text
 是不是只有自定义的词比你们已有的词长才会有效，而比你们的词短的就不生效？
```
```text
 你可以试试lcut(seq,HMM=false)

Brent

On Aug 7, 2017, at 18:55, jiangchao123 <notifications@github.com<mailto:notifications@github.com>> wrote:


是不是只有自定义的词比你们已有的词长才会有效，而比你们的词短的就不生效？

—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub<#503 (comment)>, or mute the thread<https://github.com/notifications/unsubscribe-auth/AM-fgZNs2YqKUicyDBIwL5pGAZrIcKnbks5sVu07gaJpZM4OvHe9>.
```
```text
 试了一下，这样在有些词是可以的，有些依然不行，并且如果HMM设为FALSE的话，好多人名都会被切开了，比如：小明被切成了小，明；李小福被切成了李，小，福；
找代驾我想分成找，代驾；但是HMM设为True为找代驾；HMM设为False为找，代，驾
看来是不能两全？？？
```
```text
 想请教tfidf部分是如何进行分词的？能自定义分词字典么，自定义删除一些词汇
tags = jieba.analyse.extract_tags(content, topK=topK, withWeight=withWeight)
```
```text
 呃 后两个问题文档中就有解答
```
## 307.jieba0.36，英文单词会被切分
Jul 8, 2015
```text
 Command "/usr/bin/python2 -u -c "import setuptools, tokenize;file='/tmp/pip-build-6jExFt/jieba/setup.py';f=getattr(tokenize, 'open', open)(file);code=f.read().replace('\r\n', '\n');f.close();exec(compile(code, file, 'exec'))" install --record /tmp/pip-xAfrwY-record/install-record.txt --single-version-externally-managed --compile" failed with error code -9 in /tmp/pip-build-6jExFt/jieba/
```
```text
 Can you provide full error message?
```
```text
 请问dict.txt 词库是否都来自搜狗的2006版免费词库? #24
```
```text
 如果词典中有个关键词是ck，英文单词lock会被切分为lo、ck。
```
## 308.导入词典是报错
Jul 3, 2015
```text
 ：将目标文本按行分隔后，把各行文本分配到多个python进程并行分词，然后归并结果，从而获得分词速度的可观提升
基于python自带的multiprocessing模块，目前暂不支持windows
```
```text
 You can try this. It works in Windows.
from path import Path
from multiprocessing import Pool
import argparse
import time

LINE_PER_CORE = 5000
NUM_CORE = 30
FLOOR_COUNT = 10
CEIL_COUNT = 200

import jieba


def process_one(_in):
    r_list = []
    for l in _in:
        new_l = ' '.join(jieba.cut(l))
        r_list.append(new_l.strip())
    return r_list


def do(l_list, writer):
    pool = Pool(NUM_CORE)
    r_list=pool.map(process_one,[l_list[it:it+LINE_PER_CORE] for it in range(0,len(l_list),LINE_PER_CORE)])
    pool.close()
    pool.join()
    for lr in r_list:
        for line in lr:
            writer.write(line + '\n')

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("-i","--input", help="input folder", default=".")
    parser.add_argument("-o", "--output", help="output folder", default="w_process")
    parser.add_argument("--LINE_PER_CORE", help="# lines per core", type=int, default=20000)
    parser.add_argument("--NUM_CORE", help="# of cores", type=int, default=30)
    parser.add_argument("--coding", type=str, default="utf-8")

    args = parser.parse_args()
    print("Args :", args)
    input_folder = args.input
    output_folder = args.output
    LINE_PER_CORE = args.LINE_PER_CORE
    NUM_CORE = args.NUM_CORE
    coding = args.coding

    if not Path(output_folder).exists():
        Path(output_folder).mkdir()
    for f in Path(input_folder).files('*.txt'):
        print(f.basename(), time.strftime('%Y-%m-%d %X', time.localtime()))
        with open(output_folder + '/%s.output.txt' % (f.namebase,),'w', encoding='utf-8') as f_out:
            with open(f.abspath(),'r', encoding='utf-8') as f_in:
                l_list=[]
                all_dict = {}
                for l in f_in:
                    if len(l_list)<NUM_CORE*LINE_PER_CORE:
                        l_list.append(l)
                    else:
                        do(l_list, f_out)
                        print(f.basename(), time.strftime('%Y-%m-%d %X', time.localtime()))
                        l_list=[]
                if len(l_list)>0:
                    do(l_list, f_out)
    print(time.strftime('%Y-%m-%d %X', time.localtime()))
```
```text
 jieba.load_userdict('entity_dic.txt')
Traceback (most recent call last):
File "", line 1, in 
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 119, in wrapped
return fn(_args, *_kwargs)
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 334, in load_userdict
logger.debug('%s at line %s %s' % (f_name, lineno, line))
NameError: global name 'f_name' is not defined
我是用pip安装的。。。
```
```text
 应该是 #257 的问题，还是升级版本吧。
```
```text
 现在版本是0.36
```
```text
 这个问题出现在 0.36 整个版本，你可以自己改正：
jieba/__init__.py第 334 行：
logger.debug('%s at line %s %s' % (f.name, line_no, line))
```
## 309.第四部分，词性标注的例子出现错误 'pair' object is not iterable
Jul 3, 2015
```text
 No description provided.
```
```text
 No description provided.
```
```text
 版本 <0.37 中 pair 对象不是 iterable，即不能像 tuple 一样直接拆开来。之前的测试代码相关部分是：
for w in words:
    print('%s %s' % (w.word, w.flag))
```
```text
 恩，已经解决，谢谢！！前后例子不太一致，导入词典的那个例子里用的是你上面写的例子。。。主页上用的是tuple
```
```text
 导入词典的例子在哪里
```
## 310.jieba有官方的QQ群之类的吗？
Jul 2, 2015
```text
 能自定义词性吗？
将一部分特殊词放在自定义词典里，标注自定义词性(例如:unc)，jieba.cut后得到特殊词的词性？
例如：先生，女士，经理，标注为称谓词
```
```text
 明显可以
你可以把下面的加到自定义字典里
先生 w 1000
女士 w 1000
经理 w 1000

在这里
单行字典的结构是
word pos freq
词 词性 词频
```
```text
 这个格式正确吗？
word freq pos?
还有就是会出现自定义词典后，词的词性变为x
```
```text
 word freq pos 是这个格式啊
```
```text
 jieba有官方的QQ群之类的吗？还有jieba有考虑文本摘要的问题么，最近在做这方面的。
```
## 311.能否对已经分词的文本单独进行词性标注呢
Jun 27, 2015
```text
 请问打开新词发现后，发现的新词可以导出吗？
```
```text
 同问，pynlpir有一个get_newword用法，可以print出新词，结巴有没有类似功能？
```
```text
 能否对已经分词的文本单独进行词性标注呢
```
```text
 照 jieba 的做法实现一个：
import re

re_skip_internal = re.compile("(\r\n|\s)")
re_num = re.compile("[\.0-9]+")
re_eng = re.compile("[a-zA-Z0-9]+")

class POSSimpleTagger:

    def __init__(self, dictfile):
        self.word_tag_tab = {}
        with open(dictfile, "rb") as f:
            for lineno, line in enumerate(f, 1):
                try:
                    line = line.strip().decode("utf-8")
                    if not line:
                        continue
                    word, _, tag = line.split(" ")
                    self.word_tag_tab[word] = tag
                except Exception:
                    raise ValueError(
                        'invalid POS dictionary entry in %s at Line %s: %s' % (f_name, lineno, line))

    def tagpos(self, words, HMM=True):
        for w in words:
            if re_skip_internal.match(w):
                yield (w, 'x')
            elif re_num.match(w):
                yield (w, 'm')
            elif re_eng.match(w):
                yield (w, 'eng')
            else:
                yield (w, self.word_tag_tab.get(w, 'x'))
注意字符串要用 unicode(py2), str(py3)。
```
## 312.新词发现
Jun 18, 2015
```text
 met a problem when using jython+jieba
jython version :jython-standalone 2.7.1
jieba :0.38
Exception in thread "main" java.lang.ExceptionInInitializerError Caused by: Traceback (most recent call last): File "jiebademo.py", line 7, in <module> import jieba File "/usr/local/lib/python2.7/site-packages/jieba/__init__.py", line 16, in <module> from . import finalseg File "/usr/local/lib/python2.7/site-packages/jieba/finalseg/__init__.py", line 30, in <module> start_P, trans_P, emit_P = load_model() File "/usr/local/lib/python2.7/site-packages/jieba/finalseg/__init__.py", line 24, in load_model start_p = pickle.load(get_module_res("finalseg", PROB_START_P)) File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 1378, in load File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 858, in load File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 858, in load File "/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py", line 966, in load_string ValueError: insecure string pickle
any help??
```
```text
 这个需要你把
/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py
line 966
上下文贴出来看看
if you can show the context of
/opt/software/apache-maven-3.3.9/m3/repos/org/python/jython-standalone/2.7.1/jython-standalone-2.7.1.jar/Lib/pickle.py
line 966
maybe i can help you
```
```text
 请问如何使用新词发现功能？以您举得例子来说，"他来到了网易杭研大厦"，分词后确实出现了“杭研”，但我怎么判断“杭研”是个新词？
```
```text
 同问，怎么样爬取，学习新词？

Regards,
Leon Lee
(86)1382-510-1940
Skype:lt.leon0519
2015-06-18 15:39 GMT+08:00 maybeluo notifications@github.com:

请问如何使用新词发现功能？以您举得例子来说，"他来到了网易杭研大厦"，分词后确实出现了“杭研”，但我怎么判断“杭研”是个新词？
—
Reply to this email directly or view it on GitHub
#273.
```
```text
 【新词】这个概念吧，基本上就是说词典里面没有的就是新词，如果词典里面有，就不是新词。。。。
```
```text
 jieba是否有接口获取发现的新词列表？或者有接口获取已有的词典用来判断哪些是新词？
```
```text
 确实是词典里面没有就是新词，但是在词典很大的情况下一个个搜索分出来的词是不是“新词”也是不现实的。
```
## 313.使用阿里云，用virtualenv和pip竟然无法顺利安装jieba
Jun 11, 2015
```text
 某些新词或不常用词拆开和合并的分词质量比较？
未来会不会有这种功能的API，方便定量研究？
```
```text
 
```
```text
 返回值是列表的形式，如何用jieba的词性标注给列表中的每个关键词标注词性。
```
```text
 参数控制 withFlag=True
```
```text
 输入： hello world
输出：
……        start           end
hello      0                5
world     5                10
这个offset位置标识其实忽略了原文中 hello 和 world 中的空格位置。这样在做高亮显示场景时是错误的。
```
```text
 怪怪
byte-compiling /opt/venv27/lib/python2.7/site-packages/jieba/posseg/prob_emit.py to prob_emit.pyc
Killed
```
```text
 这是一个很大的数据文件，估计内存很小不够编译成 pyc 文件被杀了。你有多少可用内存？
```
```text
 多谢多谢，怪不得，我的机器上只有300多m的可用内存，估计是不够用了
```
```text
 不过你可以直接复制文件进去，现在 300m 应该够载入标准词典运行 jieba 的。这点内存应该也够编译啊，除非是 pypy。
```
```text
 在aws上，小于1G内存都编译不了。说的就是那个老版最小的实例
```
```text
 好像还是这样，我 pip install jieba --log log.txt ，log 最后一句还是：Killed ，我 ram 应该是够得都没怎么用。
```
## 314.ValueError: math domain error
May 22, 2015
```text
 No description provided.
```
```text
 在基于TF-IDF进行特征提取时，因为文本背景是某一具体行业，不适合使用通用的IDF语料库，我觉得应该使用自定义的基于该行业背景的IDF语料库。请问如何生成自定义IDF语料库呢？
我现在有的数据是几十万个该行业的文档，初步想法是：对每个文档分词去重，把所有文档分词结果汇集去重后形成一个分词集，然后对于分词集里的每一个词语w，按**f=log(该行业文档总数/(含有w的文档数量+1))**公式求出词语w的IDF值f，最后txt文件里每一行放一个(w, f)
是这样吗？谢啦~
```
```text
 还有2个问题：假设通用IDF语料库里有A B C三个词语及其idf值，我自定义IDF语料库里有A B D及其idf值，那么请问，在添加自定义的IDF语料库后：

自定义IDF语料库里的A和B及其相应idf值就直接覆盖通用IDF语料库里的A和B吧？
通用IDF语料库里原先的C及其idf值，现在还有吗？

（其实问题只有就1个：添加自定义IDF语料库后，是整个文件替换，还是说只有那些重复的词语才被替换？）
```
```text
 求助求助求助，没有朋友知道吗？？？
```
```text
 我也想做一个词库,满足自己的需求,自带的字库里有很多类似一一二/一一分/一三六八之类意义不大的词
```
```text
 line 是单个文档
all_dict = {}
for line in lines:
    temp_dict = {}
    total += 1
    cut_line = jieba.cut(line, cut_all=False)
    for word in cut_line:
        temp_dict[word] = 1
    for key in temp_dict:
        num = all_dict.get(key, 0)
        all_dict[key] = num + 1
for key in all_dict:
    w = key.encode('utf-8')
    p = '%.10f' % (math.log10(total/(all_dict[key] + 1)))
```
```text
 @M2shad0w 非常感谢！还有一个问题：
假设通用IDF语料库里有A B C三个词语及其idf值，我自定义的IDF语料库里有A B D及其idf值，那么请问，在添加自定义的IDF语料库后：

自定义IDF语料库里的A和B及其相应idf值就直接覆盖通用IDF语料库里的A和B吧？
通用IDF语料库里原先的C及其idf值，现在还有吗？
```
```text
 @siberiawolf61
我看了一下 结巴库中 load idf path 的代码
https://github.com/fxsjy/jieba/blob/master/jieba/analyse/tfidf.py#L65
class TFIDF(KeywordExtractor):

    def __init__(self, idf_path=None):
        self.tokenizer = jieba.dt
        self.postokenizer = jieba.posseg.dt
        self.stop_words = self.STOP_WORDS.copy()
        self.idf_loader = IDFLoader(idf_path or DEFAULT_IDF)
        self.idf_freq, self.median_idf = self.idf_loader.get_idf()

...

self.idf_loader = IDFLoader(idf_path or DEFAULT_IDF)

应该是覆盖了，c值的 idf 也没有了
```
```text
 @M2shad0w 好的，谢谢啊！
```
```text
 感谢！
```
```text
 运行test目录下的demo.py时出现了
1. 分词
Building prefix dict from /usr/local/lib/python2.7/dist-packages/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.000440835952759 seconds.
Prefix dict has been built succesfully.
Full Mode: 我/ 来/ 到/ 北/ 京/ 清/ 华/ 大/ 学
Traceback (most recent call last):
File "demo.py", line 22, in 
print("Default Mode: " + "/ ".join(seg_list))  # 默认模式
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 274, in cut
for word in cut_block(blk):
File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 199, in cut_DAG
calc(sentence, DAG, route=route)
File "/usr/local/lib/python2.7/dist-packages/jieba/__init.py", line 144, in calc
logtotal = log(total)
ValueError: math domain error
这是因为什么？
这是用python2还是python3写的？
```
```text
 词典好像没成功加载，删掉 /tmp/jieba.cache，下载最新版试试。
最新版 Python 2/3 均兼容
```
## 315.自带词典中词与自定义词曲中词相同时，出现的问题！！
May 22, 2015
```text
 每一次提取一句話的關鍵詞的速度平均都在6s左右如何優化這個過程
```
```text
 写了一个程序，使用py2exe制作可执行文件，但是制作完成后，点击执行文件，出现这样的错误，IOError: [Errno 2] No such file or directory: dist\library.zip\jieba\finalseg\prob_start.py'，而且手动添加这个文件，错误依旧，希望能够得到你的解答，谢谢。
```
```text
 @fxsjy 你好，感谢你给我们这么好用的工具！！
在使用时，发现一个小问题，不知道怎么去解释。
描述：
jieba自带词库(dict.txt)中有'云南'，我在自定义词典(user_dict.txt)中也加入'云南'。导入自定义词典，然后分词结果就变成如下这样了：



for i in pg.cut('云南旅游之地质公司'):
print i.word,i.flag



云 ns
南 ns
旅游 vn
之 u
地质 n
公司 n
即 '云南‘被拆成单字了。
```
```text
 @bobo1732 , 请问你用的版本是多少？
```
```text
 @fxsjy 版本是:0.36.2
PS:好速度！！
```
```text
 @fxsjy 貌似我已经发现的类似问题，#210
“因为add_word这个实现没有去更改已经加载了的词的概率，而total已经发生变化了。”
是这个原因吗？
```
```text
 上面的#210问题应该已经解决。最新开发版无论有没有加“云南”自定义词，都分成：
云南旅游 nz
之 u
地质 n
公司 n

删掉“云南旅游”，无论有没有加“云南”自定义词，都分成：
云南 ns
旅游 vn
之 u
地质 n
公司 n

不知道是什么问题
```
```text
 @gumblex 问题是：在自定义词典中加入‘云南’后，‘云南’被拆成‘云’、‘南’。
你用的是哪个版本？
```
```text
 GitHub 上的，我直接用 add_word 测试，与自定义词典等效。可能是自定义词典中其他词的频率影响了“云南”的分词。
```
```text
 @gumblex 可以肯定：不是自定义词典中词频影响问题.
因为我只要将自定义词典中‘云南’这个词删除后，一切都好了。。
```
```text
 我测试下来，用默认dict.txt，“云南 1000000”能正常分出“云南 ns”；用dict.txt.small，能正常分出“云南 ns”，“云南 1”就分出了“云 ns/南 ns”。所以我觉得问题出在词典上。你有没有手动指定词频？
```
```text
 @gumblex 哦，我自定义词典中的词的词频就是这个词的长度，云南只有 2
```
```text
 那真不幸，能分出“云南”的最低词频是3。请参考 suggest_freq 函数。0.36 版及以上支持自动计算词频，在自定义字典里现在不必（瞎）指定词频了。
```
```text
 @gumblex 哦，基本明白了，谢谢！
```
## 316.词性标注的问题
May 19, 2015
```text
 免费分享，造损免责。
打开默认词典（根目录）或自定义词典，把所有用来间隔词频和词性的空格间隔符改成@@
（选用@@是因为一般关键词里遇到这个分隔符的几率比较小吧）
继续，打开jieba根目录下init.py

搜索
re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._]+)", re.U)
改成
re_han_default = re.compile("(.+)", re.U)



搜索
re_userdict = re.compile('^(.+?)( [0-9]+)?( [a-z]+)?$', re.U)
改成
re_userdict = re.compile('^(.+?)(\u0040\u0040[0-9]+)?(\u0040\u0040[a-z]+)?$', re.U)



搜索
word, freq = line.split(' ')[:2]
改成
word, freq = line.split('\u0040\u0040')[:2]



补充：若用的全模式继续改。
搜索
re_han_cut_all = re.compile("([\u4E00-\u9FD5]+)", re.U)
改成
re_han_cut_all = re.compile("(.+)", re.U)
```
```text
 您好，我正好碰见这个问题，我想把 迈克尔·乔丹 不要拆开，试了您的方法，
将默认的dict改成了如下：
AT&T@@3@@nz
B超@@3@@n
c#@@3@@nz
C#@@3@@nz
c++@@3@@nz
将user_dict变成了如下：
小皇帝
科比布莱恩特
迈克尔·乔丹
init.py也按照您的方法改了，没用全模式所以没改最后一个
加载报错：
ValueError: invalid POS dictionary entry in C:\Tools\anaconda\lib\site-packages\jieba\dict_gai.txt at Line 1: AT&T@@3@@nz
请您指教！感激不尽！
```
```text
 @xusong123 我试了一下，没有问题啊，是不是用用了记事本打开txt导致加上了bom头。建议用notepad++创建编辑文本。
```
```text
 您好，我试了试，还是不行，方便的话我把这个贴上来给您看看
dict_gai3.txt
我加载的是dict_gai3.txt
```
```text
 @xusong123  你试一下，不用我的方法，然后把你字典里的点·删掉，然后看看报错不。如果也报错，就证明你后来新加的词，编码上有问题。
```
```text
 谢谢，已经解决，最好用python3
```
```text
 @xusong123 请问怎么解决的？我也出现同样的错误了，谢谢
```
```text
 我也遇到了这个问题。通常是/tmp/jieba.cache未即时更新引起的。
解决方法：删除jieba.cache，把默认字典（dict.txt）中的空格替换为@@即可。
@wonderfreer
```
```text
 我需要加载千万级的词典，词典才300M，但是加载后内存高达5G。
self.FREQ[word] = freq
self.total += freq
if tag:
self.user_word_tag_tab[word] = tag
for ch in xrange(len(word)):
wfrag = word[:ch + 1]
if wfrag not in self.FREQ:
self.FREQ[wfrag] =
FREQ
user_word_tag_tab
用了多个字典去存储数据，是不是浪费内存啊，这里可以合并成一个词典吧
```
```text
 @zhangzhenhu ， 可以试一试，发个PR吧 :-)
```
```text
 @fxsjy 128m内存直接跑不起来，256的才勉强能跑。
我用的bae。
```
```text
 是的，我也遇到一样的问题。加载词典太耗内存了！
```
```text
 请教一个词性标注的问题
现在的jieba.posseg.cut标注的词性w.flag，是不是都是从字典里面取出来的？
但是字典里面只定义了一种词性，而很多词是多词性的。
例如'连'这个词，在金山词霸网站可以查到，'连'有多重词性
http://www.iciba.com/%E8%BF%9E

本义：（动）相连；连接。（副）连续；接续；继续（表示前后没有断开）：～演三十场戏。
本义：（动）相连；连接。（介）包括在内：～我三个人｜～根拔。
本义：（动）相连；连接。（名）军队的编制单位；由若干排组成。
本义：（动）相连；连接。（Lián）姓。

ltp是哈工大在线工具分析结果：  http://www.ltp-cloud.com/demo/
例句一：连小学生都会。
ltp:   [连/u  小学生/n  都/d  会/v  。/wp]
jieba: [连/nr 小学生/nr 都/d 会/v 。/x]       -> 字典： 连 23315 nr
例句二：这是一个连的兵力。
ltp:   [这/r 是/v  一个/m  连/n  的/u  兵力/n  。/wp]
jieba: [这/r 是/v  一个/m  连的/d 兵力/n 。/x]
【问题】怎样定义多性词？比如'连'在第一个句子里应该是u，第二个句子里应该是n
谢谢。
```
```text
 目前jieba的POS Tagging，基于词库分词后，对词库中存在的词，直接取用词典中的词性（第三列为词性）；对于未登录词，再用HMM序列标注来同时完成新词发现和词性标注。
对于多性词，jieba的方案比较差，而中文中多性词约占了23.6%，比例还是很大的。
哈工大的LTP并不死绑词典中的词性，而考虑了上下文，并针对未登录词做了不少优化，相较之下好很多。
```
## 317.lcut不可用
May 18, 2015
```text
 `



seg_list = jieba.cut("他来到了网易杭研大厦")  # 默认是精确模式
print(", ".join(seg_list))
他, 来到, 了, 网易, 杭研, 大厦`



但是在分词过程中判断词性
`



words = pseg.cut("来到了网易杭研大厦")
for word, flag in words:
...  print('%s %s' % (word, flag))`



依旧会把“杭研”分为两个词  杭 j 研 vn  ；
如何判断“杭研”是否为新词？
```
```text
 可以參考我修正的錯誤 #670
```
```text
 用demo里的程序就没有问题，自己新建一个python文件，复制里面的import，执行时总是提示ImportError: No module named analyse
什么原因
C:\Python27\Lib\site-packages\jieba-0.38-py2.7.egg-info
jieba-0.38-py2.7.egg-info\jieba
都存在
```
```text
 应该是包结构的问题，看看你init.py是否正确，试试不用没有包结构的情况
```
```text
 很多专有词语中间是有空格的
例如：OpenGL ES
词典分隔符用空格不太合理，建议改成\t
```
```text
 因为多义词的原因，建议支持词典中支持多词性，方便后续设计更复杂的算法
```
```text
 结巴分词挺好用的，但是我想直接用lcut生成一个list时出现里下面的问题：
Traceback (most recent call last):
File "test.py", line 21, in 
seg_list = jieba.lcut(L[3])
AttributeError: 'module' object has no attribute 'lcut'
编译的时候报错没有lcut是为什么？
部分代码如下：
import jieba
with open('/home/arnold-hu/dataset/news_data_200line.txt','r') as f:
line = f.readline()
L = line.split('\t')
seg_list = jieba.lcut(L[3])
```
```text
 你是通过pip install安装的？这个模块似乎目前得用github源代码安装才能用上
```
```text
 用的easy_install，因为用源码安装比较麻烦，这个可以写入readme文档里说明一下吧，虽然lcut貌似用的不多
```
```text
 这是不久之前更新的功能，还没在PyPI上更新。
```
## 318.求助
May 16, 2015
```text
 比如：“我来到北京清华大学”  分成： 我 / 来 / 到 / 北 / 京 / 清 / 华 / 大 / 学
```
```text
 这种事儿也用分词么？

2017-11-21 20:00 GMT+08:00 jasmineol <notifications@github.com>:
…
 比如：“我来到北京清华大学” 分成： 我 / 来 / 到 / 北 / 京 / 清 / 华 / 大 / 学

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#552>, or mute the thread
 <https://github.com/notifications/unsubscribe-auth/ABD5_ejuZPwUklkvnwSV_5J3028wkVO3ks5s4rtUgaJpZM4QluiD>
 .
```
```text
 这个网上随便一搜都是分字的
比如把中文字分隔转为list
```
```text
 补充一下，c++如何使用jieba分字，涉及到gbk，utf8这种编码的转换，所以想看下jieba有没有现成的
```
```text
 这种直接获取编码就是一个词吧，之前看过汉子转拼音。
```
```text
 设置一个空词典，关闭HMM新词联想，分出来全是一个个字的
```
```text
 哪位大牛能教教我，谢谢
```
```text
 把jieba.cache的位置由临时目录改为当前目录就可以了，我周末刚刚修改并且成功部署了：fay@e094dac#L0L65
```
```text
 import jieba.analyse
File "/usr/local/lib/python2.7/dist-packages/jieba/analyse/init.py", line 7, in 
from .textrank import textrank
File "/usr/local/lib/python2.7/dist-packages/jieba/analyse/textrank.py", line 6, in 
import collections
File "/usr/lib/python2.7/collections.py", line 10, in 
from keyword import iskeyword as _iskeyword
ImportError: cannot import name iskeyword
```
```text
 参见 #207
```
## 319.怎样将某个词必然与其他的词分隔开？
May 15, 2015
```text
 打扰，请问这个工具用的算法和Stanford Word Segmenter用的算法有什么主要区别吗？
以及，如果使用了这个工具，需要引用你们的论文吗？（或者说这个工具有基于某篇论文吗？）——我也可以直接引用github链接。。。
```
```text
 自定义词库：
*ST舜船
代码：
title_list = jieba.cut("*ST舜船")
print "/ ".join(title_list)
分词结果：
*/ ST/ 舜船
版本0.38    查看Changelog：4. 修复load_userdict加载用户词典不能识别含有空格等特殊字符的问题， by @gumblex;
```
```text
 我也遇到这个问题
```
```text
 同遇到这个问题。
# 结巴自定义词典不支持特殊符号
import jieba
jieba.add_word("<NUM>")
jieba.add_word('<ENG>')
jieba.add_word('<UNK>')
s = "我今天买了<NUM>块钱的<UNK>，<ENG>还不错"
print(jieba.lcut(s))
实际输出
['我', '今天', '买', '了', '<', 'NUM', '>', '块钱', '的', '<', 'UNK', '>', '，', '<', 'ENG', '>', '还', '不错']

期望输出
['我', '今天', '买', '了', '<NUM>', '块钱', '的', '<UNK>', '，', '<ENG'>', '还', '不错']
```
```text
 比如我想要光彩这个词被单独分开，但是当出现光彩照人时，它们就被连在了一起，我希望将光彩固定与其他词分开
```
```text
 @CrazymageLee 有两种方式：
1、jieba.suggest_freq(('光彩','照人'),tune=True)
2、删除‘光彩照人'这个词, jieba.del_word('光彩照人')
```
```text
 @bobo1732 这种方法只适用于光彩照人这一种情况，光彩有时候会与其他的词分在一起，我是希望某个词永远与其他所有词都分开。
```
```text
 当然，简单粗暴的方法是手动指定「光彩」一个特别高的词频，100000之类。
```
## 320.对于jieba分词的一个疑问
May 13, 2015
```text
 hi,
我在jieba.dict.utf8这个字典里面等26556行发现下面的内容
今天上午 3 nr
今天下午 3 nr
nr不是人名的标注吗？ 为什么这俩的标注是nr？
```
```text
 终止时刻的可能状态列表如果为E或者S表示它是一个更合适的候选者，除非没有E和S状态，才考虑B和M，看了下char_state_tab，没有e,s状态的字很少。
另外jieba.posseg.viterbi.py19行应该注释掉吗?
```
```text
 @lurga , 对于第一个问题，需要多测试才能给出准确答复。
对于第二个问题，非常感谢你，19行本来是我对viterbi算法做的一个剪枝策略，应该是忘了注释。
剪枝策略应用后，分词速度会更快，准确度会下降多少我也没来得及测试。
你有空的话，可以帮忙看看19行的策略和20行的策略，综合起来，哪个更好。谢谢。
```
```text
 @fxsjy 我不太懂python，用scala改了个版本，由于概率文件是一样的，结果应该一样，以下测试都直接使用viterbi来分词，没有用词典。
1、第一个问题我做了下面的测试，优先选择E,S状态的效果要更好一点，当然这只是一个个例，没办法证明它的效果始终更优。
测试语句是"张三说的确实在理"
终止状态优先选择E,S：('B', 'nrfg'), ('M', 'nrfg'), ('E', 'nrfg'), ('S', 'uj'), ('B', 'ad'), ('E', 'ad'), ('B', 'v'), ('E', 'v')
不判断终止状态：('B', 'nrfg'), ('M', 'nrfg'), ('E', 'nrfg'), ('S', 'uj'), ('B', 'ad'), ('E', 'ad'), ('S', 'p'), ('B', 'n')
2、第二个问题我做了下面的测试，剪枝策略对结果影响很大。
测试语句是 “中国商务部”
top4：('M', 'nr'), ('M', 'nr'), ('M', 'nr'), ('M', 'nr'), ('M', 'nr')
top10：('E', 'nrfg'), ('S', 'n'), ('B', 'j'), ('M', 'j'), ('M', 'j')
top20：('E', 'nrfg'), ('S', 'n'), ('B', 'j'), ('M', 'j'), ('M', 'j')
top30：('M', 'nrfg'), ('E', 'nrfg'), ('B', 'j'), ('M', 'j'), ('M', 'j')
top40：('B', 's'), ('E', 's'), ('B', 'nt'), ('M', 'nt'), ('E', 'nt')
top50：('B', 'n'), ('E', 'n'), ('B', 'nt'), ('M', 'nt'), ('E', 'nt')
top60：('B', 'nt'), ('M', 'nt'), ('M', 'nt'), ('M', 'nt'), ('E', 'nt')
大致取到前60个状态才跟不剪枝的结果保持一致。对t-1时刻的状态做出限定，极大可能导致最后结果出现严重误差，数学如何解释我也不清楚，盼答复。
```
```text
 为什么在CUT里面使用的HMM是BMES 这种模型。而不是基于词性的分词？
我感觉基于词性的效果可能比BMES好一些。
是使用过传统的词性的HMM效果不好，还是其他原因。在这个地方有点困惑希望作者能解惑。
如果使用过传统的词性效果不太好，我就不去尝试了。目前觉得BMES 的分词能力可能达不到需要的精度。
```
## 321.想拆除带有特殊符号的词
Apr 30, 2015
```text
 No description provided.
```
```text
 用 del_word('男性') 直接将“男性”这个词删掉？
```
```text
 for i in jieba.analyse.extract_tags('我爱上海的水'):
print i
这个例子里面只能输出上海
我，爱，水，都没了？ 请问为什么呀
```
```text
 seg_list = jieba.cut('我爱上海的水',cut_all=True)
print("Full Mode: " + "/ ".join(seg_list))  # 全模式
```
```text
 我想把“尿β２微球蛋白值”或者“PGI/PG”作为一个词拆分，我用了add_word和suggest_freq方法，但是仍然被拆开了，这种词的拆分能实现吗？
```
```text
 因为你的词里面 β２ 和 / 不在中文正则匹配中，所以即使你添加了这两个词，在读取字典之前，就已经被拆分了：
re_han_default = re.compile("([\u4E00-\u9FA5a-zA-Z0-9+#&\._]+)", re.U)
re_han_cut_all = re.compile("([\u4E00-\u9FA5]+)", re.U)

你可以根据自己的需要来修改 re_han_default，例如：
>>> import re
>>> import jieba
>>> jieba.re_han_default = re.compile(r'([\u4e00-\u9fa5a-zA-Z0-9+#&\._/β２]+)', re.UNICODE)
>>> jieba.add_word("尿β２微球蛋白值")
>>> list(jieba.cut(u"尿β２微球蛋白值"))
['尿β２微球蛋白值']
```
```text
 我在linux和windows下的测试结果都如下
`



import jieba
import re
jieba.re_han_default = re.compile(r'([\u4e00-\u9fa5a-zA-Z0-9+#&._/β２]+)', re.UNICODE)
jieba.add_word("尿β２微球蛋白值")
list(jieba.cut(u"尿β２微球蛋白值"))
[u'\u5c3f', u'\u03b2', u'\uff12', u'\u5fae', u'\u7403', u'\u86cb', u'\u767d', u'\u503c']
`
版本的问题吗?
@wangbin
```
## 322.想拆出一个带有中文和英文的词
Apr 29, 2015
```text
 如何验证分词的准确性
请问用jieba分完词后，如何去验证分词的准确性？
```
```text
 Same question. What is the training data? 人民日报语料吗？
```
```text
 Same question. What is the training data? 人民日报语料吗？
不是人民日报的，是网上爬取的一些评论数据，最近看了数学之美的分词，有一些启发
```
```text
 一些不字开头的词语在短句中，为啥被extract_tags后"不"字就不见了，将"不"与后面词语分开，这样提取的词的意思都变了，不合理吧
比如

不合适
不适合
不太合适
不太适合
不满意
不想买
不想要
```
```text
 在用词性标注的时候如以下句子：
前港督衛奕信在八八年十月宣布成立中央政策研究組
词性标注后变成
前/f 港督/n 衛/v 衛奕/z 信/n 在/p 八八年/m 十月/t 宣布/v 成立/v 中央/n 政策/n 研究組/n
多出一个衛字
单纯的用分词是OK的，不知道是不是我词性标注的选取的词性标注方法有问题。
我用的是官网上的样例：



import jieba.posseg as pseg
words = pseg.cut("前港督衛奕信在八八年十月宣布成立中央政策研究組")
for w in words:
...    print w.word, w.flag
```
```text
 @elfcool ， 多谢反馈，这是个bug，我刚刚修复了一下：9d4ac26
你有空pull下来看看问题还存不存在。
```
```text
 我想把“超敏C反应蛋白”作为一个词拆分，我用了add_word和suggest_freq方法，但是仍然被拆开了，这种词的拆分能实现吗？
```
```text
 我测试 add_word 是可以的啊？不知道是不是我理解的不对？
>>> list(jieba.cut("超敏C反应蛋白是什么？"))
['超敏', 'C', '反应', '蛋白', '是', '什么', '？']
>>> jieba.add_word("超敏C反应蛋白")
>>> list(jieba.cut("超敏C反应蛋白是什么？"))
['超敏C反应蛋白', '是', '什么', '？']
```
```text
 你理解的是对的，真心感谢，因为这类的词比较多，我是用python读文件的形式挨个添加的，不知道是不是这个出了问题导致的
```
```text
 您好，谢谢您的回答，但是我刚才试了一下，我这里还是不行，不知道你是用的什么版本，我这个是linux系统，python环境，O(∩_∩)O~
张柳-云平台研发部
发件人： wangbin
发送时间： 2015-04-29 10:42
收件人： fxsjy/jieba
抄送： liuliu818
主题： Re: [jieba] 想拆出一个带有中文和英文的词 (#253)
我测试 add_word 是可以的啊？不知道是不是我理解的不对？



list(jieba.cut("超敏C反应蛋白是什么？"))
['超敏', 'C', '反应', '蛋白', '是', '什么', '？']
jieba.add_word("超敏C反应蛋白")
list(jieba.cut("超敏C反应蛋白是什么？"))
['超敏C反应蛋白', '是', '什么', '？']



—
Reply to this email directly or view it on GitHub.
Confidentiality Notice: The information contained in this e-mail and any accompanying attachment(s)
is intended only for the use of the intended recipient and may be confidential and/or privileged of
Neusoft Corporation, its subsidiaries and/or its affiliates. If any reader of this communication is
not the intended recipient, unauthorized use, forwarding, printing,  storing, disclosure or copying
is strictly prohibited, and may be unlawful.If you have received this communication in error,please
immediately notify the sender by return e-mail, and delete the original message and all copies from
your system. Thank you.
```
```text
 我也是 Linux 系统， Python 3.4.3， jieba 0.36.2
```
## 323.带有词性标注的分词效率是多少？
Apr 15, 2015
```text
 在自己的代码中用suggest_freq(word)操作确保特殊词不被拆分, 但同时代码其他部分引用的外部模块也同时引用了jieba, suggest_freq(word)的操作会影响外部模块分词操作(并不是希望看到的). 对此是否提供清除suggest_freq的操作?
```
```text
 import importlib
importlib.reload(jieba) is okey...
```
```text
 by using importlib.reload method, you may encounter this scenario in iPython/jupyter:
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
...
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.

Once more you reload jieba module, there is one more line output in each section when jieba.dt.initialize() is implemented.
I'm not sure if there's any influence to other codes.
```
```text
 在cut句子以後，如果變更字典集，不會更新到cache，有什麼方法可以在initialize還可以重新建立一次cache嗎？
ex.
jieba.cut("我的靴子裡有蛇") #我/ 的/ 靴子/ 裡有/ 蛇
jieba.add_word("我的靴子")
jieba.cut("我的靴子裡有蛇") #還是... 我/ 的/ 靴子/ 裡有/ 蛇
```
```text
 import jieba
import jieba.posseg as pseg
s = "挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好"
seg = jieba.cut(s, cut_all=False)
挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好






pos=pseg.cut(s)
挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好/z



这个怎么处理下，比较好呢？
```
```text
 带有词性标注的分词效率是多少？
```
## 324.How to use jieba in hadoop?
Mar 11, 2015
```text
 根据不同的参数名称调用不同的自定义词典和stopwords来分词,自定义词典和stopwords全部初始化加载到内存中,实际分词的时候根据参数来替换dict和stopwords的变量值,这样能实现吗?
```
```text
 I use spark cluster to seg word and submit the program with jieba to the cluster as following:
/home/hadoop/spark-1.2.0-bin-hadoop2.4/bin/spark-submit --master spark://hz0124:7077 --driver-memory 1g --executor-memory 5g --total-executor-cores 60 --py-files ../sparkcode/dependency/dependency.zip seg
word_test.py
the jieba is zipped as the file ../sparkcode/dependency/dependency.zip
then an error occured:

Anyone meets this problem? how to deal with it?
```
```text
 Using zip will cause problems to read the dictionary because it uses regular file IO.
You can unzip the dependencies, or put the dict.txt elsewhere (not in a zip) and then manually initialize jieba with jieba.initialize('/path/to/dictionary').
```
```text
 发现在cut_all模式下，字符串中的特殊字符不能得到很好的返回，查看代码发现这两种模式下的正则表达式设定是不同的，想请教一下，为何两种模式下，设定的正则是不同的
测试结果如图所示：
精确模式

全分割模式
```
```text
 同遇到该问题，怎么回事
```
```text
 RT
```
```text
 Maybe a Java version of jieba.
作者：piaolingxue 地址：https://github.com/huaban/jieba-analysis
```
## 325.新版本无法载入自定义词库
Mar 10, 2015
```text
 Hi I'm using Jieba with Java and I have a question.
My purpose to use Jieba is getting words from sentences and add mean per word. But sometimes the result is not what I expected.
sentence : 我要九个
result : 我要 / 九个
exptected : 我 / 要 / 九 / 个
sentence : 吃三丸
result : "吃三丸"
expected : 吃 / 三 / 丸
Because I have a Chinese English dictionary like this.
我 : I
要 : want
九 : nine
个 :  individual; measure word for people or objects
But for current result I have to have all possible cases in the dictionary.
ex）一个， 两个，三个，四个，六个，七个，八个，九个，是个。。。。
Thanks in advance.
```
```text
 代码如下：
#encoding=utf-8
import jieba
import jieba.posseg as posseg

jieba.load_userdict("/data/paper/keywords/user_dict.txt")

words = posseg.cut("在现代社会中，免洗消毒液的使用越来越普遍，消毒液中所含的肉豆蔻酸异丙酯、丙二醇和乙醇等成分可增强皮肤的渗透性。")

for w in words:
    print('%s %s' % (w.word, w.flag))

报错
Building prefix dict from /usr/lib/python2.7/site-packages/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.296606063843 seconds.
Prefix dict has been built succesfully.
Traceback (most recent call last):
File "./jieba-test.py", line 5, in 
jieba.load_userdict("/data/paper/keywords/user_dict.txt")
File "/usr/lib/python2.7/site-packages/jieba/init.py", line 119, in wrapped
return fn(_args, *_kwargs)
File "/usr/lib/python2.7/site-packages/jieba/init.py", line 312, in load_userdict
add_word(_tup)
File "/usr/lib/python2.7/site-packages/jieba/init.py", line 116, in wrapped
return fn(_args, **kwargs)
TypeError: add_word() takes at most 3 arguments (4 given)
之前还工作的好好的。user_dict.txt 类似于这样的：
酞 10 n
 酞酐 20 n
 酞磺胺噻唑 50 n
 酞磺醋胺 40 n
 酞基 20 n
 酞己炔酯 40 n
 酞腈 20 n
 酞菁 20 n
 酞菁二(吡啶)铁(II)络合物 150 n
 酞菁镉 30 n
 酞菁蓝 30 n
 酞菁锂 30 n
 酞菁镁 30 n
 酞菁染料 40 n
 酞菁素艳蓝 50 n
 酞菁锡(II) 70 n
 酞菁亚铁 40 n
 酞菁铟 30 n
 酞菁银 30 n
 酞类染料 40 n
 酞嗪 20 n
 酞醛 20 n
 酞酸 20 n
 酞酸丁苄酯 50 n
 酞酸二苯酯 50 n
 酞酸二丙酯 50 n
 酞酸二环己酯 60 n
 酞酸二甲酯 50 n
 酞酸二壬酯 50 n
 酞酸二辛酯 50 n
 酞酸二乙酯 50 n
 酞酸二异癸酯 60 n
 酞酮酸 30 n
 酞醯胺醯基 50 n
 酞醯基 30 n
 酞醯亚胺基 50 n
 酞酰胺 30 n
 酞酰氯 30 n
 酞亚胺 30 n

其中优先级是按词语长度用脚本自动生成的，给予长词更高的优先级。
```
```text
 我发现了。似乎是自定义词库本身的问题
```
## 326.word segementation of "省"
Mar 6, 2015
```text
 结巴的字典是几年前的，不知道能否加入一些新的词重新训练一份？
使用了自定义词典感觉效果不太好。
```
```text
 By theory, jieba is not a trainable model i'm afraid
```
```text
 可以训练啊，你拿别的人工或机器分词/标注的语料，统计一下词频就行。
```
```text
 @gumblex 你好 能详细说说这个过程吗, 我想做一个自己的关键字提取的idf.txt, 但是不知道如何下手, idf.txt中的第二列, 那一串数字, 不知道怎么得出  比如
劳动防护 13.900677652
勞動防護 13.900677652
生化学 13.900677652
```
```text
 Hi,
I am using Jieba in python3, and having the following problem:
for the sentence with word "省", e.g.  "丰田太省了", the results are: 丰田/ 太省/ 了.
However, it should be expected as: 丰田/ 太/ 省/ 了.  it happens also to "最省".
I checked the jieba.dict, and do not find the word "太省" in the dictionary. Could you check if it is an issue in the code? Thank you.
Best Regards,
Qing
```
```text
 >>> tuple(jieba.cut('丰田太省了'))
('丰田', '太省', '了')
>>> tuple(jieba.cut('丰田太省了',HMM=0))
('丰田', '太', '省', '了')

It's a 'new word' found by the HMM model because the model thinks '太省' is probably a word. So is the case of '最'. If you want to get the result only according to the dictionary, set HMM=False.
```
```text
 yes, it works great after set HMM=False. thanks!
```
## 327.jieba.posseg.cut 分词问题
Feb 12, 2015
```text
 把痘痘加入自定词典后，也无法把长痘痘中的长和痘痘分开。各位大佬有什么办法嘛？
```
```text
 jieba.possseg.cut的分词不能支持和jieba.cut一样的可选cut_all参数，这样两种模式下分词结果不一样。
```
## 328.windows下加载自定义词典的bug
Feb 11, 2015
```text
 代码入下：
content="环境设施都很好，服务周到！"
aspect_content_array = jieba.analyse.extract_tags(content, topK=50)
print ",".join(aspect_content_array)
抽取结果：
服务周到,设施,环境
我想得到“好”这个关键词，但是抽取不出来
另外能否在抽关键词的时候，把“服务周到” 抽取成 “服务”和“周到”两个词
```
```text
 刚试了下自定义词典，里边加入这两行：
打底裤 10000
打底衫 10000
发现这两个词仍然分开了。debug之后发现这些词并没有加载。问题在这里：
jieba/init.py:
def load_userdict(f):
......
if not line.rstrip():
continue
tup = line.split(" ")
word, freq = tup[0], tup[1]
if freq.isdigit() is False:
continue
windows下userdict结尾是"\r\n", 因为我的词典，词频“1000”后面立刻是"\r\n", "\n"被split吃掉了，"\r"还在line结尾。所以tup里，freq="1000\r"，判断 isdigit()为False，这几个词都没被加载。
怀疑前面的
if not line.rstrip():
continue
是不是应该是
line = line.rstrip()
if not line:
continue
这么改之后，词是加载了（在load_userdict()里用print看了），而且“打底衫”分好了，但“打底裤”仍然分成了“打/底裤”。我在词频后面加两个零，结果还是一样。那是什么问题呢？
```
```text
 试试 Git 里的新版本
```
```text
 @gumblex 找到原因了。。原来是我画蛇添足给文件加了个BOM头。老版本并没有删掉这个头。现在好了 :)
```
```text
 @gumblex  java版的结巴， 支持添加自定义词典吗？
```
```text
 @myccc456 Jython 应该是支持的。如果你说的是 这个 Java 版本，根据 README 也应该是支持的。
```
```text
 @gumblex   谢谢， 试了下，通过如下代码即可自定义添加词典 java版
    WordDictionary dictAdd = WordDictionary.getInstance();

    File file = new File("D:/jieba-analysis/conf/user.dict");
    dictAdd.loadUserDict(file);
```
## 329.解码的bug
Feb 11, 2015
```text
 First of all thank you for your work, you did so far!
I have a question for this specific line https://github.com/fxsjy/jieba/blob/cb0de2973b2fafaa67a0245a14206d8be70db515/jieba/posseg/init.py#L17. Why do you use this specific range of unicode literals? For my specific case(russian text) your app not splitting words which is not good at all.
For example:
>>> re_han_internal = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._]+)")
>>> re_han_internal.split("""Простой и безопасный способ делиться терминалом: обзор инструмента warp""")
['Простой и безопасный способ делиться терминалом: обзор инструмента ', 'warp', '']

Increasing unicode range helps:
>>> re_han_internal = re.compile("([\u0041-\u9FD5a-zA-Z0-9+#&\._]+)")
>>> re_han_internal.split("""Простой и безопасный способ делиться терминалом: обзор инструмента warp""")
['', 'Простой', ' ', 'и', ' ', 'безопасный', ' ', 'способ', ' ', 'делиться', ' ', 'терминалом', ': ', 'обзор', ' ', 'инструмента', ' ', 'warp', '']
```
```text
 it's based on the dict which is ONLY for chinese by default, u wanna use the program, u should add a special russian dict.
```
```text
 这个简单的测试程序，会产生exception：
import jieba
s = "女装"
a = jieba.cut(s, cut_all=False)
print "/".join(a)
错误是：
UnicodeEncodeError: 'gbk' codec can't encode character u'\u016e' in position 0: illegal multibyte sequence
我debug发现原因是 "女装" 用utf-8解码时并不会产生exception，但其实解码后的东西是乱七八糟的：u'\u016e\u05f0' 后面用这个字符串去分词了。我不知道为什么这个词这么特殊。
一个解决办法是提供一个参数，人工指定编码，这样就肯定不会错了。
```
```text
 在mac os下用Python 2.7.9测试没有出错，返回结果还是“女装”这个词。
```
```text
 @alexwwang 我的.py是GBK编码的。你的脚本头加上# -- coding: gb18030 -- 试试？
```
```text
 s 是 str 类型，gb18030 编码（根据源代码的编码而定）；你的命令行环境是 'gbk' 编码的。
因此，在分词时，由于 s 是 str 类型，程序先进行尝试 utf-8 解码（解码错误再尝试 gbk），结果！解码成功，为Ůװ两个字符，尝试 print 的分词结果为 Ů/װ。但由于 gbk 无法编码这两个字符（码表里没有），只好异常。
因此，要避免出这种状况：

s = u"女装"
尽量在源代码中使用 UTF-8 编码
```
```text
 @gumblex 谢谢。我上面的代码只是PoC，实际上是从一个GBK文件逐句读出然后分词，产生的异常。所以不能用您建议的办法来避免了。目前我的做法是先decode再传给jieba. 不过因为jieba号称不管编码是GBK或者unicode都可以自动支持，所以我觉得这个还应该是个bug。
```
```text
 做文本处理的该用 Unicode 类型就用 Unicode 类型，毕竟分词针对的是字符，不是用来负责判断这到底是 GBK 还是 Unicode 编码。
```
## 330.分词改进
Feb 9, 2015
```text
 如何将类似“abc123456”的文本切分为“abc”和“123456”？默认设置是切分不开的。
```
```text
 给代码里的几处re_eng戳个洞，分出来一类数字看看？
```
```text
 谢谢啦！
```
```text
 以下句子分词可改进：
句子：今天天气不错
jieba result: 今天天气    不错
期望结果：今天    天气     不错
句子：如果放到post中将出错。
jieba result: 如果 放到 post 中将   出错
期望结果：如果 放到 post 中    将    出错
```
```text
 fxsjy closed this in #351 on 16 Mar
```
## 331.「台中」總是被切成「台」「中」
Feb 8, 2015
```text
 def setLogLevel(log_level):
global logger
default_logger.setLevel(log_level)
logger 是否应该是 default_logger?
```
```text
 jieba.__version__ = 0.34
字典檔確定有「台中」這個詞。
不論用何種模式，「台中」都會切成「台」「中」，但是正確應該不會被切開。
「台北」、「台南」就不會被切開，這樣是正確的。


請問這有問題有辦法解決嗎？？
```
```text
 可以把台中加入自定义词典jieba.load_userdict
```
```text
 好，我試試看，但是我好奇的是，「台中」這個字詞已經在字典裡面了，為何還是斷詞還是斷成「台」「中」？？
```
```text
 >>> jieba.FREQ['台中']
3
>>> jieba.FREQ['台']
16964
>>> jieba.FREQ['中']
243191
>>> jieba.total
60101967
3/60101967 < 16964×243191/601019672
即，P(台中) < P(台)×P(中)，台中词频不够导致概率较低
在自定义词典调高词频即可，或者add_word用更高频率覆盖“台中”
（这里是>16964×243191/60101967，即≥69）
```
```text
 我发现jieba会将“不合理”看做是一个词，而不是“不/合理”。
另外，基于@gumblex 给出的计算，应该是可以将“不”和“合理”分开的，如下：
jieba.get_FREQ(u'不合理') # 470
jieba.get_FREQ(u'不') # 360331
jieba.get_FREQ(u'合理') # 3870
词典中共有584429个词（dict.txt.big），计算：
470/584429 < 360331*3870/584429^2
如果使用user dict，当把“合理”的词频调高后，可以分开。
```
```text
 这个计算的分母是总频数(jieba.dt.total = 60101967)不是总词数，这样计算下来「不合理」确实高于其他分词。
你可以手动拆分：jieba.suggest_freq(('不 ','合理'), True)，或直接删除 del_word。
```
```text
 @gumblex ：
明白了。另外一个问题是：
1）在当前程序中调用suggest_freq，会影响后续的使用么？
2）del_word只是对当前的分词起作用么？是否并没有将其在dict.txt中删除？
```
```text
 @gumblex :
厚着脸皮再问一个，POS标注中的分词结果是否和jieba.cut相同？
```
```text
 suggest_freq 和 del_word 等仅对当前分词器实例 (jieba.Tokenizer) 起作用（旧版是仅对全局词典起作用），不会修改本地词典，相当于改了几个变量。
POS标注中的分词结果不一定相同，见 #212。
```
```text
 怎么让 POS标注 支持 suggest_freq 和 del_word ?
```
## 332.請問如何簡單取出新詞?
Jan 27, 2015
```text
 在使用 load_userdict() 时，并不会主动关闭文件，导致程序警告
ResourceWarning: unclosed file <_io.BufferedReader name='dict.txt'>

不知道没有关闭文件的意义是什么？还有后续操作嘛？
```
```text
 jieba 的词典是用空格作为词，词频，词性之间的分隔符的，但是当要在句子中添加“CH125 Spacy”这种词的时候就会识别不了。
如果词典使用“\t”是不是会相对较好？
```
```text
 @rockyzhengwu, @fxsjy 您好, 請問您有解決嗎?我也想要自訂字典中加有空白的英文組合字，但都識別不出來，　都會拆成好幾個字...
```
```text
 我知道dict.txt是jieba的詞庫，請問當程式用jieba lib分析一篇文章時，
如何列出該篇文章不在dict.txt的所有關鍵詞?
底下是我想到的簡單作法:
是先用cut把文章做分詞之後，然後再一一與dict.txt的詞比較，不知道有沒有其他方式?
```
```text
 我知道dict.txt是jieba的词库，请问当程式用jieba的lib分析一篇文章时，
如何列出该篇文章不在dict.txt的所有关键词？
底下是我想到的简单作法：
是先用cut把文章做分詞之後，然後再一一與dict.txt的詞比較，不知道有沒有其他方式?
```
## 333.新增自定义词典后分词不准确了
Jan 21, 2015
```text
 举例：
sen = u"我钟爱法国"
我/r 钟爱/nr 法国/ns
sen = u"我鍾愛法國 "
我/r 鍾/nr 愛/v 法國/ns
有什么方法可以改善么？比如可以把简体字的词库改成繁体字词库么？
```
```text
 改用dict.big.txt的繁體字典後, 並參考我修正的錯誤 #670
```
```text
 import jieba as jb
import jieba.posseg as jp
if name == 'main':
r = "我来自中华人民共和国。"
d = jb.cut_for_search(r);
d1 = jp.cut(" ".join(d));
像这样写两次吗？
```
```text
 搜索引擎模式貌似把词拆碎了，再标注词性的用处不大了吧？
```
```text
 我也想问，主要是过滤标点……
```
```text
 import jieba.posseg as pseg
import jieba
加载用户词典
jieba.load_userdict("C:\Python27\Lib\site-packages\jieba\new_dict.txt")
words = pseg.cut("女性有月经暴力作用后腹痛")
w_list=[]
print type(words)
for w in words:
#if w.flag.startswith('n') or w.flag.startswith('v') :
print w.word,w.flag
加载后自定义词典后，居然把“月经”分成“月”和“经”两个词了。
在自定义的词典中月经词如下：
月经 5 n
```
```text
 >>> jieba.FREQ[u'月经']
742
>>> jieba.FREQ[u'月']
110207
>>> jieba.FREQ[u'经']
21042
>>> jieba.total
60101967

P(月经) > P(月)P(经) > P(月经)new
742/60101967  > 110207×21042/601019672 > 5/60101967
添加词汇/载入用户词库会覆盖原有词频，而不是相加。已验证这和 #210 的 bug 无关。
```
```text
 @gumblex , 多谢！！明白了
```
```text
 @gumblex ，0.38版本，jieba.total 编译报错？
```
```text
 @cavonchen 改成 jieba.dt.total
```
```text
 好的，谢谢！~
```
## 334.用 pycharm 找不到 cut 怎么回事啊
Jan 19, 2015
```text
 应该可以通过 添加 用户词典，然后判断词是否在 词典里吧？
但是添加用的是jieba.load_userdict
所以如果我要用if word in 这个dict叫什么呢？。。。。
```
```text
 添加了一个dict但是一直乱码啊，好奇怪啊，utf-8格式，别的用同样方法添加的停用词都不会乱码。一旦我想要匹配  输出词和dict就会输出乱码
```
```text
 用这个，定义主词库，jieba.set_dictionary('dict.txt.big')
自定义的字典，用Notepad++创建，别用win的记事本
```
```text
 这个是添加用户词典里的词，从而可以被识别和区分，但是并不是只输出词典里有的词吧？


Sent from Mail Master
在2016年12月20日 19:15，mycrystalgirl 写道:

用这个，定义主词库，jieba.set_dictionary('dict.txt.big')
自定义的字典，用Notepad++创建，别用win的记事本

—
You are receiving this because you authored the thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 你都已经定义了主词典，当然所有分词结果都是都主词典抽取的，所以只输出主词典有的词。
记住几点
1.记得别用关键词提取，而是用分词功能jieba.cut
2.关闭HMM
即便这样，还是匹配出字符串里任何的英文和数字字符串，所以
3.修改init.py把里面的eng正则破坏掉，我另一个帖子讲的
我也是刚刚用结巴分词，需求跟你的一样，才摸索出的。
```
```text
 你说需要破坏init.py
“”修改init.py把里面的eng正则破坏掉，我另一个帖子讲的“”
能不能把你的帖子地址分享一下。我也遇到了非常类似的问题。谢谢
```
```text
 HMM=false
```
```text
 为什么使用 Pycharm 调用 jieba 时候，cut 函数找不到？？？？
输入
import jieba
jieba.cut("我来到北京清华大学")
输出
AttributeError: 'module' object has no attribute 'cut'

更奇怪的是，输入
import jieba
print("我来到北京清华大学")
输出两行这个字
我来到北京清华大学
我来到北京清华大学

这些情况单独用 python 运行没问题。
请问是怎么回事
```
```text
 Pycharm 4.5.3 pro能出cut的呀！
```
```text
 不要将运行的文件名命名为jieba.py，自己撸自己当然出错了
```
```text
 哈哈，我也犯了同样的错误，不要将运行的文件名命名为jieba.py！！！
```
## 335.词性标注问题
Jan 16, 2015
```text
 部分代码：
sentence = getText('三国演义.txt')
keywords = jieba.analyse.extract_tags(sentence, topK=20, withWeight=True, allowPOS=('ns'))
按地名统计前20，出来的结果是：
将军 0.11860080694708297
丞相 0.0921638281251496
主公 0.0715118726461463
军士 0.061001657055284146
商议 0.05922492016471381
云长 0.054936383725351035
军马 0.05441077997191448
大喜 0.05001738396254889
后主 0.04707253654544958
先主 0.04402888433316952
都督 0.04296257804954991
众将 0.04167988860499132
天下 0.03893417374252234
陛下 0.03882608055724891
太守 0.035023324694504726
人马 0.03359843706149093
城上 0.03265662810340311
天子 0.03212681223959663
后人 0.03164288695617565
众官 0.03082541531255808
这是怎么统计出来的地名？完全不对啊！
```
```text
 No description provided.
```
```text
 在线演示还有人在维护吗？
```
```text
 我数据里的“有点”全部被标注成“/n”了。
服务/vn 有点/n 情调/n
这是怎么了？要怎么处理下才行？
谢谢。
```
```text
 @followYourHeart , 为了兼顾性能， dict.txt中的词只有一种词性，只有对于dict.txt中没有的词才会调用HMM去识别。你可以删掉dict.txt中的这行：有点 3706 n ，效果：
服务 / vn ,  有 / v ,  点 / q ,  情调 / n ,
或者，你改一下dict.txt中它的词性。
```
## 336.如何只处理中文？
Jan 14, 2015
```text
 你好，在做基于情感词典的情绪分析，结果发现，结巴分词，会把很多词合在一起。比如：你真笨；
我想要的结果是：你，真，笨，
但是返回的却是：你，真笨
还有：你太笨，返回的也是：你，太笨
感觉很不科学呐。真和太明明是一种程度词，怎么和笨合成一个词了。
同时调研了其它多个分词开源工具，发现只有你们会合在一起。
```
```text
 是不是只有自定义的词比你们已有的词长才会有效，而比你们的词短的就不生效？
```
```text
 你可以试试lcut(seq,HMM=false)

Brent

On Aug 7, 2017, at 18:55, jiangchao123 <notifications@github.com<mailto:notifications@github.com>> wrote:


是不是只有自定义的词比你们已有的词长才会有效，而比你们的词短的就不生效？

—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub<#503 (comment)>, or mute the thread<https://github.com/notifications/unsubscribe-auth/AM-fgZNs2YqKUicyDBIwL5pGAZrIcKnbks5sVu07gaJpZM4OvHe9>.
```
```text
 试了一下，这样在有些词是可以的，有些依然不行，并且如果HMM设为FALSE的话，好多人名都会被切开了，比如：小明被切成了小，明；李小福被切成了李，小，福；
找代驾我想分成找，代驾；但是HMM设为True为找代驾；HMM设为False为找，代，驾
看来是不能两全？？？
```
```text
 想请教tfidf部分是如何进行分词的？能自定义分词字典么，自定义删除一些词汇
tags = jieba.analyse.extract_tags(content, topK=topK, withWeight=withWeight)
```
```text
 呃 后两个问题文档中就有解答
```
```text
 我有一段文本，其中有很多英文以及标点符号，希望只对其中的中文进行分词处理，除了用 stop_words ,还有什么好的办法吗？
```
```text
 先用正则把非中文的字符去除
```
```text
 可以试试和这个类似的正则表达式 http://git.io/piM6rQ
```
## 337.分词结果很奇怪
Jan 11, 2015
```text
 在自己的代码中用suggest_freq(word)操作确保特殊词不被拆分, 但同时代码其他部分引用的外部模块也同时引用了jieba, suggest_freq(word)的操作会影响外部模块分词操作(并不是希望看到的). 对此是否提供清除suggest_freq的操作?
```
```text
 import importlib
importlib.reload(jieba) is okey...
```
```text
 by using importlib.reload method, you may encounter this scenario in iPython/jupyter:
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Building prefix dict from the default dictionary ...
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Loading model from cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
Dumping model to file cache /tmp/jieba.cache
...
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Loading model cost 2.636 seconds.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.
Prefix dict has been built succesfully.

Once more you reload jieba module, there is one more line output in each section when jieba.dt.initialize() is implemented.
I'm not sure if there's any influence to other codes.
```
```text
 在cut句子以後，如果變更字典集，不會更新到cache，有什麼方法可以在initialize還可以重新建立一次cache嗎？
ex.
jieba.cut("我的靴子裡有蛇") #我/ 的/ 靴子/ 裡有/ 蛇
jieba.add_word("我的靴子")
jieba.cut("我的靴子裡有蛇") #還是... 我/ 的/ 靴子/ 裡有/ 蛇
```
```text
 import jieba
import jieba.posseg as pseg
s = "挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好"
seg = jieba.cut(s, cut_all=False)
挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好 挺 好






pos=pseg.cut(s)
挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好挺好/z



这个怎么处理下，比较好呢？
```
## 338.extract_tags结果会遗漏分词
Jan 8, 2015
```text
 java版本的分词Ansj和HanLP都支持用自己的语料再训练，为什么结巴分词这么久也不支持？
```
```text
 No description provided.
```
```text
 我只能告诉你这玩意看得懂空格和标点……（
```
```text
 用 https://github.com/moses-smt/mosesdecoder/blob/master/scripts/tokenizer/tokenizer.perl
```
```text
 楼主，一直以来用你的jieba分词，感谢你无私的分享！
近来发现一个问题：
for i in jieba.analyse.extract_tags('上海一日游攻略'):
print i
输出结果：
攻略
上海
我看到idf.txt中‘一日游’与‘攻略’这个词的IDF相差无几，为什么没‘一日游’这个词呢？
不太明白其中原理。
```
```text
 @bobo1732 , 我这里没有复现此bug，请问你的版本是多少？
```
```text
 @fxsjy 我是用的0.35版本
```
## 339.ChineseAnalyzer的問題
Dec 24, 2014
```text
 关键词提取无论如何总能得到一些结果，在语料非常多的时候，人工标记不太具备可行性，那么有哪些推荐的算法用于评估提取内容的准确性和相关性呢？
```
```text
 长标点符号应该被分为一个词。
例如：
longpunct = set(u'——', #EM dash U+2014, 常用
u'--', #半角-
u'－－', #U+FF0D, 全角－
u'――', #Horizontal bar, U+2015
u'──', #Box Drawing Light Horizontal, U+2500
u'……',
u'......', # '.'*6
u'...', # '.'*3
) #etc.
```
```text
 么么哒~
```
```text
 想請問:
我使用ChineseAnalyzer for Whoosh 搜索引擎，當writer.add_document裡面的content字不多的時候，設定的keyword會跟content完全符合才會被列出來; 但當content字多到一定程度的時候(我用bog文章)，content的字只要有部分和keyword一樣就會被列出來。不知道該怎麼處理? 是預設的模式有變化嗎? 看了Whoosh說明文件也不太懂，謝謝。
```
```text
 简单改下：
class ChineseTokenizer(Tokenizer):
应该就可以实现。
if len(context) < MIN_CONTEXT:
token = Token()
token.text = context
token.pos = 0
...
yield token
else:
words = jieba.tokenize(text, mode="search")
token = Token()
for (w,start_pos,stop_pos) in words:
if not accepted_chars.match(w) and len(w)<=1:
continue
token.original = token.text = w
token.pos = start_pos
token.startchar = start_pos
token.endchar = stop_pos
yield token
```
```text
 想到可以沿用原來的程式，最後面把print hit.highlights("content")的結果中"<.....>XXX"的XXX抓出來和keyword比對，如果一樣就是真的一樣。只是這樣就跟原來ChineseAnalyzer的效果不大一樣??
```
## 340.“我们中出了一个叛徒”分词结果不一致
Dec 20, 2014
```text
 No description provided.
```
```text
 在Hadoop上怎么加载自定义字典啊
```
```text
 例如：
中共中央/总书记/，/国家/主席/，/中央军委/主席
合并成三个词：中共中央总书记/国家主席/中央军委主席
或者
一带/一路
合并成：一带一路
```
```text
 同问,在关键词抽取这个场景下,太细粒度的分词后抽取到的关键词,不一定是最准确的,这种组合起来的分词更符合关键词提取需求,NLPIR分词后会发现行业新词,因此抽取效果上感觉更好
```
```text
 使用精确分词和搜索引擎分词得到的结果均为：

我们/ 中出/ 了/ 一个/ 叛徒

而使用词性分词得到结果：

我们/r 中/f 出/v 了/ul 一个/m 叛徒/n
```
```text
 23333
词库里竟然有 中出 这个词。。。我觉得“中出”的出现频率应该不会高于“出了”，为何没有分成功呢。
```
```text
 @shanzi 哈哈哈哈
```
```text
 哈哈哈
```
```text
 哈哈哈
```
```text
 ('枪杆子', '中', '出', '政权')
(枪杆子/n, 中/f, 出/v, 政权/n)

毕竟“中出 叛徒”比“中出 政权”的可能性大
```
```text
 @tonghuix , 分词效果不一致的原因是因为：精确分词中用的HMM模型只有四种中状态（B,E,M,S)，而词性分词中HMM模型的状态是(B,E,M,S)与各种词性的笛卡尔积。
```
```text
 感觉这句话念着好别扭。呵呵…
```
```text
 @fxsjy 感谢详解！那么也就是说，这种分词不一致的情况是不是应该还是很普遍呢？但我看在线演示的例子里 “工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作” 、”我们买了一个美的空调“ 等等句子的分词结果是正确的
```
```text
 @tonghuix , sorry，当前的算法还不够智能不能cover所有的case。 最简单的方法就是你把这种bad case按照你的意图分好，然后让它“强化”学习，得到新的词频文件和HMM模型文件。不过这种调优往往是按下葫芦起了瓢。
```
```text
 @fxsjy 确实如你所说。有没有一种类似敏感词库这种机制来匹配类似情况呢？也就是当发生bad case的时候启用强制分词，当然这种方法比较dirty
```
```text
 jieba.cut('我们中出了一个叛徒', HMM=False) 没有问题，说明“中出”是 HMM 新词发现造出来的。
```
```text
 那也是因为作者训练语料选得好（邪恶……）
B '\u4e2d': -4.596743315282086,
E '\u4e2d': -5.476938898829366,
M '\u4e2d': -5.253194167156688,
S '\u4e2d': -4.81355762044073,
B '\u51fa': -5.448380394833719,
E '\u51fa': -5.556879470894473,
M '\u51fa': -7.232833997741765,
S '\u51fa': -5.854838121207067,
```
## 341.关于自定义词典词频问题等
Dec 19, 2014
```text
 jieba.lcut('2017年10月5日或2017-10-03或12:21和12点30分还有十二点三十分')
分出来
['2017年', '10月', '5日', '或', '2017', '-', '10', '-', '03', '或', '12', ':', '21', '和', '12点', '30分', '还有', '十二点', '三十分']
如何分成
['2017年10月5日', '或', '2017-10-3', '或', '12:21', '和', '12点30分', '还有', '十二点三十分']
```
```text
 把这些词汇加入到词典中

发自我的vivo智能手机

sugarZ <notifications@github.com>编写：
…
jieba.lcut('2017年10月5日或2017-10-03或12:21和12点30分还有十二点三十分')
分出来
['2017年', '10月', '5日', '或', '2017', '-', '10', '-', '03', '或', '12', ':', '21', '和', '12点', '30分', '还有', '十二点', '三十分']
如何分成
['2017年10月5日', '或', '2017-10-3', '或', '12:21', '和', '12点30分', '还有', '十二点三十分']
—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 之前也想到了,如果固定的几个还可以,但是非常大量的日期时间格式,没法全部加到字典,这方法太笨了.
```
```text
 @sugarZ，请问lcut可以将“2017年”分词到一起吗，貌似我这边是“2017”，“年”
```
```text
 @JiaWenqi 好像不行,我是加的自定义词典,把最近的年份都添加了进去,还有1-12月0-24点(钟),但是如果把时间也加上就太麻烦了,如果jieba支持特定格式分词配置就好了.
```
```text
 目前想到一个方案，将待分词文本用时间正则进行分割后分段进行分词
```
```text
 在词典中添加a/d,d/a分词的结果仍然是三部分，不能形成一个词，
权重调了也没有用，估计是特殊字符的问题。
```
```text
 与这个类似：#363
```
```text
 请教各位：
1，自定义词库的词频该如何设置，分词效果会比较好？
我的自定义词库有3w8的词量，只有词没有词频，我就统一设了个词频。起初设置词频均为1，效果不好，很多词都分不出来。后来统一增加词频到100,200,300,500,1000,2000,3000，当词频在1000下时效果随着词频增加变好，而1000以上效果差距不明显。但是我看前面的话题您有回答说词频一般不用太高，3-5就差不多了。而dict.txt这个里词频高的词也不是特别多的。我该怎么设置好呢？
2，另外，能否解释下jieba.cut，finalseg.cut，posseg.cut这几个接口对分词效果的区别？
非常感谢！
```
```text
 建议实际统计一下未分词的语料中此自定义词库的词频，若不合适可乘上一个系数试试
区别在：

jieba.cut 用了词典词频 + finalseg 中的 HMM 模型（可选择不用），准确率高
finalseg.cut 只用 HMM 模型，准确率不高，内存占用少
posseg.cut 与 jieba.cut 相似，用于分析词性（Part of Speech）
```
```text
 不好意思，關於自定義辭庫我也有相關的問題，但是類型不太一樣。
我設了一個userdict.txt並用jieba.load_userdict(file_name)，將其導入。其後是使用jieba.cut分詞。
一開始自定義的詞我只設一個詞，凱特琳，詞頻為3，此時可以將該詞分出來。
接著我將userdict.txt擴充至442個詞，詞頻也皆為3，此時卻無法將"凱特琳"分出來，結果變為"凱特"+"琳"。其他詞也無法正確分詞。
而後將"凱特琳"的詞頻逐漸從5,10,100,調高至1000都無法將"凱特琳"分出來。
最後我是將整個userdict.txt裡的內容複製到dict.txt裡面，才有辦法將所有詞分出來。
之後我又測試，將dict.txt還原，將userdict.txt把所有詞去掉只留下"凱特琳"，此時又可以分出來"凱特琳"了。
可能的問題會出在哪呢？
```
```text
 add_word 中修改词频FREQ[word] = log(freq / total) 这样只给单字打补丁又不增加 total 的算法肯定有问题。应该调整所有词汇的频率，或者干脆只记录频数，即时计算频率？（需要测试）
```
```text
 @2153030   我也碰到这种自定义词典添加后反而不能正确切词的问题。难道这是HMM的局限？
```
```text
 这应该是词频算法问题吧，如之前所说。同 #222   @fxsjy
```
```text
 @gumblex , 对，因为add_word这个实现没有去更改已经加载了的词的概率，而total已经发生变化了。
```
```text
 请教下一下，2台电脑：Mac和Ubuntu
jieba均为0.35，使用同样的用户词典
测试代码完全一样
但是对于「石墨烯」这个词，mac上可分出「石墨烯」而在ubuntu中分出「石墨」「烯」
这是为什么呢？
补充：在Mac中删掉jieba.cache，终于。。也分不出了。。
```
## 342.如何从零开始建立一个中文词典
Dec 17, 2014
```text
 jieba.posseg.cut是可以分割词并得到tag类型，但是却把百分比分割成两个部分，tag都是'x'，jieba.cut就能准确分成"2%"，如果需要tag还需要百分比，应该怎么写
```
```text
 同样遇到了这个问题
```
```text
 同问
```
```text
 已解决，见https://blog.csdn.net/liuhongyue/article/details/80498195
```
```text
 似乎没有添加自己的停用词可能，是不是可以考虑添加
```
```text
 结巴分词中 自定义词典里面的新词的tfidf是怎么算的呢？
```
```text
 自问自答吧 ， 没有出现的词，idf是 所有词idf的平均是11左右。 tf就是这个句子中这个新词的词频。
详细分析：http://www.cnblogs.com/zle1992/p/8822832.html
```
```text
 自问自答吧 ， 没有出现的词，idf是 所有词idf的平均是11左右。 tf就是这个句子中这个新词的词频。
详细分析：http://www.cnblogs.com/zle1992/p/8822832.html

如果是平均值，为什么不是标准平均值，而是有差异呢？
```
```text
 在自己有raw语料的情况下, 有一组单词(没有词频, 主要是想通过获得这些单词的词频去做jieba分词), 如何统计词频会比较合理有效呢?
```
```text
 人工标注或者用其他准确度高的分词软件分好。 然后，统计一下。
```
```text
 链接文章里的新词发现算法也许对你有帮助。http://www.matrix67.com/blog/archives/5044
2014-12-18 9:53 GMT+08:00 Sun Junyi notifications@github.com:

人工标注或者用其他准确度高的分词软件分好。 然后，统计一下。
—
Reply to this email directly or view it on GitHub
#209 (comment).
```
```text
 @fxsjy 能不能公开建立词库、finalseg 和 posseg 概率模型的源码？这对于建立自定义模型很有帮助。
```
## 343.分词部分报错
Dec 9, 2014
```text
 应该可以通过 添加 用户词典，然后判断词是否在 词典里吧？
但是添加用的是jieba.load_userdict
所以如果我要用if word in 这个dict叫什么呢？。。。。
```
```text
 添加了一个dict但是一直乱码啊，好奇怪啊，utf-8格式，别的用同样方法添加的停用词都不会乱码。一旦我想要匹配  输出词和dict就会输出乱码
```
```text
 用这个，定义主词库，jieba.set_dictionary('dict.txt.big')
自定义的字典，用Notepad++创建，别用win的记事本
```
```text
 这个是添加用户词典里的词，从而可以被识别和区分，但是并不是只输出词典里有的词吧？


Sent from Mail Master
在2016年12月20日 19:15，mycrystalgirl 写道:

用这个，定义主词库，jieba.set_dictionary('dict.txt.big')
自定义的字典，用Notepad++创建，别用win的记事本

—
You are receiving this because you authored the thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 你都已经定义了主词典，当然所有分词结果都是都主词典抽取的，所以只输出主词典有的词。
记住几点
1.记得别用关键词提取，而是用分词功能jieba.cut
2.关闭HMM
即便这样，还是匹配出字符串里任何的英文和数字字符串，所以
3.修改init.py把里面的eng正则破坏掉，我另一个帖子讲的
我也是刚刚用结巴分词，需求跟你的一样，才摸索出的。
```
```text
 你说需要破坏init.py
“”修改init.py把里面的eng正则破坏掉，我另一个帖子讲的“”
能不能把你的帖子地址分享一下。我也遇到了非常类似的问题。谢谢
```
```text
 HMM=false
```
```text
 为什么使用 Pycharm 调用 jieba 时候，cut 函数找不到？？？？
输入
import jieba
jieba.cut("我来到北京清华大学")
输出
AttributeError: 'module' object has no attribute 'cut'

更奇怪的是，输入
import jieba
print("我来到北京清华大学")
输出两行这个字
我来到北京清华大学
我来到北京清华大学

这些情况单独用 python 运行没问题。
请问是怎么回事
```
```text
 Pycharm 4.5.3 pro能出cut的呀！
```
```text
 不要将运行的文件名命名为jieba.py，自己撸自己当然出错了
```
```text
 哈哈，我也犯了同样的错误，不要将运行的文件名命名为jieba.py！！！
```
```text
 一个很奇怪的问题
我的代码如下：
import sys
reload(sys)
sys.setdefaultencoding("utf-8")
sys.path.append("../")
import jieba
jieba.setLogLevel(60)
jieba.load_userdict("dict.txt")
import jieba.analyse
test_sent = sys.stdin.read()
tags = jieba.analyse.extract_tags(test_sent, topK=6)
print ",".join(tags)
保存为 test.py
在 用户根目录执行的时候 执行成功，成功提取出结果
我将test.py 和 dict.txt 复制到 /data/www/目录下执行 就会报错
Traceback (most recent call last):
File "./keyword.py", line 9, in 
import jieba.analyse
File "/usr/local/lib/python2.7/dist-packages/jieba/analyse/init.py", line 9, in 
from textrank import textrank
File "/usr/local/lib/python2.7/dist-packages/jieba/analyse/textrank.py", line 5, in 
import collections
File "/usr/lib/python2.7/collections.py", line 10, in 
from keyword import iskeyword as _iskeyword
File "/data/www/spider/keyword.py", line 12, in 
tags = jieba.analyse.extract_tags(test_sent, topK=6)
AttributeError: 'module' object has no attribute 'analyse'
我安装jieba的时候 使用 pip安装
pip install jieba
真的搞不懂原因了，求解
```
```text
 你的 keyword.py (原 test.py) 与标准库重名了
./keyword.py 引用 jieba.analyse 引用 textrank 引用 collections 引用 keyword，然后本目录的模块优先搜索，就用你的 keyword.py 了。
解决方法：改名
```
```text
 谢了我之前的时候 经过测试发现确实是这个问题。。非常感谢
```
## 344.无法识别userdict中新增标点符号的词性
Dec 8, 2014
```text
 部分代码：
sentence = getText('三国演义.txt')
keywords = jieba.analyse.extract_tags(sentence, topK=20, withWeight=True, allowPOS=('ns'))
按地名统计前20，出来的结果是：
将军 0.11860080694708297
丞相 0.0921638281251496
主公 0.0715118726461463
军士 0.061001657055284146
商议 0.05922492016471381
云长 0.054936383725351035
军马 0.05441077997191448
大喜 0.05001738396254889
后主 0.04707253654544958
先主 0.04402888433316952
都督 0.04296257804954991
众将 0.04167988860499132
天下 0.03893417374252234
陛下 0.03882608055724891
太守 0.035023324694504726
人马 0.03359843706149093
城上 0.03265662810340311
天子 0.03212681223959663
后人 0.03164288695617565
众官 0.03082541531255808
这是怎么统计出来的地名？完全不对啊！
```
```text
 No description provided.
```
```text
 在线演示还有人在维护吗？
```
```text
 我数据里的“有点”全部被标注成“/n”了。
服务/vn 有点/n 情调/n
这是怎么了？要怎么处理下才行？
谢谢。
```
```text
 @followYourHeart , 为了兼顾性能， dict.txt中的词只有一种词性，只有对于dict.txt中没有的词才会调用HMM去识别。你可以删掉dict.txt中的这行：有点 3706 n ，效果：
服务 / vn ,  有 / v ,  点 / q ,  情调 / n ,
或者，你改一下dict.txt中它的词性。
```
```text
 自定义词典中加入以下标点符号的词性标注：
! 50000 w
" 50000 w
但是jieba无法正确分出词性，统一被识别为非语素x



words = pseg.cut('南京"中山陵"位于紫金山!')
for w in words: print w.word,w.flag
...
南京 ns
" x
中山陵 ns
" x
位于 v
紫金山 nr
! x
```
## 345.jieba支持文本断句吗？
Oct 24, 2014
```text
 jieba.analyse.textrank()方法提取英文文本的关键词没有输出数据，但是jieba.analyse.tfidf()方法则可以处理英文文本。中文文本两个都可以处理。so ,问题是，，，
```
```text
 你看看源代码就知道了，tokenizer和postokenizer的区别。
```
```text
 re_han_default = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&._]+)", re.U)
既然对C++, C#多了特判, 为什么不加入 objective-c特判
```
```text
 成功安装，但是报错信息Failed building wheel for jieba，可能是这个错误的原因，jieba模块不能正常运行
```
```text
 给详细日志啊
```
```text
 D:\learn-python>pip install jieba
Collecting jieba
Using cached jieba-0.37.zip
Building wheels for collected packages: jieba
Running setup.py bdist_wheel for jieba
Complete output from command c:\users\administrator\appdata\local\programs\python\python35-32\python.exe -c "import setuptools;file='C:\Users\ADMINI1\AppData\Local\Temp\pip-build-ibpo4da7\j
ieba\setup.py';exec(compile(open(file).read().replace('\r\n', '\n'), file, 'exec'))" bdist_wheel -d C:\Users\ADMINI1\AppData\Local\Temp\tmpf6wrfjs2pip-wheel-:
Traceback (most recent call last):
File "", line 1, in 
UnicodeDecodeError: 'gbk' codec can't decode byte 0x80 in position 100: illegal multibyte sequence

Failed building wheel for jieba
Failed to build jieba
Installing collected packages: jieba
Running setup.py install for jieba
Successfully installed jieba-0.37
给这个你看行不行
```
```text
 感觉是 pip 的锅：终端是 GBK 编码，setup.py 是 UTF-8，然后读入直接运行
你的运行环境是什么？（我打算去报 bug）
```
```text
 windows7 64位系统
D:\learn-python>python --version
Python 3.5.0
```
```text
 这个现在有解决方案吗
```
```text
 按照给定的标点符号规则，将文本断句切割
例如：
我是大哥,他是二哥,三哥没有.
按照逗号断句
返回结果是
我是大哥,
他是二哥,
三哥没有.
这样的，jeiba支持吗？
```
```text
 你其实可以自己训练HMM模型来断句。我本来想以jieba为基础做一个，可是这方面需求好像不大，就懒得做了。
```
```text
 jieba能做文本相似度的处理吗？
```
```text
 可以提取关键词，文本相似度可用别的库。话说你想干什么？
```
```text
 文本相似度有哪些好的库？
做毕设
```
```text
 这要看你的具体应用了，去看一下各种文本或数据相似度算法的介绍和比较，再做决定。每种成熟的算法基本上都有对应的Python库。
```
```text
 我基本是先分句再喂给jieba,用的re.split().
```
```text
 @gumblex 如果用HMM 該怎麼做? 就像分詞一樣的做法嗎?
```
## 346.有过滤过滤功能吗？
Aug 23, 2014
```text
 我在自定义词典里加入了带有罗马数字的词，如“纽约Ⅰ线”，并且为其赋予很高的词频，但是总是被分隔开。请问我要如何才能强制把“纽约Ⅰ线”这样的自定义词完整切分出来？？？
词典相关内容如下：
纽约Ⅰ线 100000 ns
python版本3.6.2
代码大致如下：
line = "并明确控制纽约Ⅰ线功率"
jieba.load_userdict("dict.txt")
print(jieba.get_FREQ("纽约Ⅰ线"))
print(jieba.get_FREQ("纽约"))
print(jieba.get_FREQ("Ⅰ"))
print(jieba.get_FREQ("线"))
print(jieba.get_FREQ("Ⅰ线"))
words = posseg.cut(line)
for word in words:
print(word)
输出结果如下：
100000
1758
None
7688
None
并/c
明确/ad
控制/updown
纽约/ns
Ⅰ/x
线/n
功率/n
```
```text
 In [39]: from jieba.analyse import ChineseAnalyzer
ImportError                               Traceback (most recent call last)
 in ()
----> 1 from jieba.analyse import ChineseAnalyzer
ImportError: cannot import name ChineseAnalyzer
请问是什么原因？
```
```text
 看看ChineseTokenizer这个类依赖的库是不是没install，比如whoosh这个库，pip install whoosh一下
```
```text
 在统计词与邻近词的连通数的时候，为什么只统计后面的词，按道理说前后的词都要统计的吧？
```
```text
 兄弟，我想问问结巴分词添加了关键词功能后，一些词如：“一些”、“大概”、“已经”...能不能过滤掉？
```
```text
 你可以自行增加 stop word 字典，就可以濾掉這些詞，詳見範例 https://github.com/fxsjy/jieba/blob/master/test/extract_tags_stop_words.py
```
```text
 兄弟，终于回我了，非常感谢！已添加 notifications@github.com 为联系人，发邮件比较容易联系你吗？
```
```text
 @agetoo notifications@github.com 為系統信，還是直接在 github 上討論就會看到了～
```
## 347.add_word(word, freq, tag=None)能增加对字符串的支持吗？
Aug 14, 2014
```text
 需求有些特别：要解决的问题的起点并不是文章。而是经过正确分词后的list。有没有什么办法在不经过jieba分词功能的情况下，直接使用关键词提取的功能（调用关键词提取方法之前，传入自定义的分词结果－list格式）
```
```text
 你是不是问题没描述明白？
你自定义的分词结果是什么东西分出来的？如果是jieba分出来的，那就是可以用自定义词典。
如果是其他方式分出来的，可以A+B。合并列表。
```
```text
 字典是中文分词的核心数据。分词的好坏，往往由字典的正确率和是否与时俱进决定。设想中，应该有这么一个机制，允许Jieba的使用者一起来贡献字典的更新。
1、比如，让Jieba的垂直行业用户，可以把针对他们的行业已经tain过的字典上传上来，分享给大家
2、再比如，Jieba自己的字典，是否可以通过对于当前一段时间互联网的最新热词/短语的流行潮流进行自我更新？
```
```text
 单开一个 repo，专门存放分享上来的字典？
通过 pull request 来分享、更新
不过这样要有人来维护这个 repo，要确保字典质量，要花精力的啊
```
```text
 @anjianshi 众筹词典？
```
```text
 众筹是大家出钱然后由某个人 / 组织负责提供、维护字典吗？不知道好不好操作 :)
```
```text
 导入自定义词库可以满足个人需求。
```
```text
 抽取句子中的序列规则，因为词性标注得太细的原因，很多规则冗余，请问可以修改源码的哪一部分？或者直接有这个功能么？
```
```text
 必须传unicode太逆天了，至少传字符串时给个excepiton会比较友好
```
## 348.吴国忠臣伍子胥的切词结果是吴国忠/ 臣/ 伍子胥
Jul 31, 2014
```text
 代码入下：
content="环境设施都很好，服务周到！"
aspect_content_array = jieba.analyse.extract_tags(content, topK=50)
print ",".join(aspect_content_array)
抽取结果：
服务周到,设施,环境
我想得到“好”这个关键词，但是抽取不出来
另外能否在抽关键词的时候，把“服务周到” 抽取成 “服务”和“周到”两个词
```
```text
 刚试了下自定义词典，里边加入这两行：
打底裤 10000
打底衫 10000
发现这两个词仍然分开了。debug之后发现这些词并没有加载。问题在这里：
jieba/init.py:
def load_userdict(f):
......
if not line.rstrip():
continue
tup = line.split(" ")
word, freq = tup[0], tup[1]
if freq.isdigit() is False:
continue
windows下userdict结尾是"\r\n", 因为我的词典，词频“1000”后面立刻是"\r\n", "\n"被split吃掉了，"\r"还在line结尾。所以tup里，freq="1000\r"，判断 isdigit()为False，这几个词都没被加载。
怀疑前面的
if not line.rstrip():
continue
是不是应该是
line = line.rstrip()
if not line:
continue
这么改之后，词是加载了（在load_userdict()里用print看了），而且“打底衫”分好了，但“打底裤”仍然分成了“打/底裤”。我在词频后面加两个零，结果还是一样。那是什么问题呢？
```
```text
 试试 Git 里的新版本
```
```text
 @gumblex 找到原因了。。原来是我画蛇添足给文件加了个BOM头。老版本并没有删掉这个头。现在好了 :)
```
```text
 @gumblex  java版的结巴， 支持添加自定义词典吗？
```
```text
 @myccc456 Jython 应该是支持的。如果你说的是 这个 Java 版本，根据 README 也应该是支持的。
```
```text
 @gumblex   谢谢， 试了下，通过如下代码即可自定义添加词典 java版
    WordDictionary dictAdd = WordDictionary.getInstance();

    File file = new File("D:/jieba-analysis/conf/user.dict");
    dictAdd.loadUserDict(file);
```
```text
 Thank you for visiting AppFog
The final sunset date for AppFog v1 is March 15, 2016, (located here at appfog.com). After March 15, 2016 AppFog v1 will no longer be available.
AppFog v2, located at ctl.io/appfog, has replaced AppFog v1.
For more information on this migration and needed action for AppFog v1 users, please see our original communication and migration guides.
```
```text
 切词为 吴国忠/ 臣/ 伍子胥
吴国 174 ns
吴国忠14 nr
臣 5666 n
忠臣 316 n
是因为P(吴国)_P(忠臣)<P(吴国忠)_P(臣)么
```
```text
 是的
```
## 349.用户自定义字典内容支持正则吗？
Jul 28, 2014
```text
 textrank中的rank()方法参考了pagerank的思路，并且在程序中iterate 10次。但是对于jieba案例，可否考虑把阻尼因子设定为1，基于这一点，可以推断出每一个节点的weight就是这个点的连线数，没有必要iterate 10反复迭代了。修改方法如下：
    for n, out in self.graph.items():
        ws[n] = wsdef
        outSum[n] = sum((e[2] for e in out), 0.0)

    # this line for build stable iteration
    # sorted_keys = sorted(self.graph.keys())
    # for x in xrange(10):  # 10 iters
    #     for n in sorted_keys:
    #         s = 0
    #         for e in self.graph[n]:
    #             s += e[2] / outSum[e[1]] * ws[e[1]]
    #         ws[n] = (1 - self.d) + self.d * s
    ws=outSum
    (min_rank, max_rank) = (sys.float_info[0], sys.float_info[3])
```
```text
 你好，我最近在看结巴分词的算法，看到在判断切分组合的时候，利用的是动态规划来求出最大组合概率，我觉得这个方法挺好。
不过，我有一个想法，不知道这样想对不对，还请指教。
我是这样想的，因为之前已经形成了切分词图，我在想可不可以利用图论里的最短路径的相关算法来做判断。但是最短路径求得是权值最短，如果把概率作为权值的话，这样直接算就算的是最小概率了。之前我看到过一篇论文有提出一个解决这种问题的办法，利用最短路径算法来求最大权值。
需要对权值做一定的转换。设边的权值为f（也就是词的频率或概率），取f的自然底数的对数的相反数，-ln(f)，作为新的权值w，即w=-ln(f)，利用这个权值w来求最短路径。
可以看出：
1，f越大，w越小
2，min(w1+w2+w3+...+wi)对应的是max(ln(f1_f2_f3_...fi))，也就是说求出有新权值w的图的最短路径后，相应的那个路径的组合概率 P=f1_f2_f3..._fi 是最大的，也就是我们要找的
3，由于f本来就是概率，所以0<f<1，因此w是非负数
所以可以通过求最短路径的方法来得到最大组合概率。
```
```text
 我感觉和我用的算法没有本质不同啊，最开始的时候我是用的概率相乘求最大概率路径，后来为了避免下溢，改为对log(p)求和来算最大概率路径。另外，由于是有向无环图，在有向无环图上求最长路径本来就是可以用动态规划的，不需要在转化为求最短路径问题吧。
```
```text
 import jieba
import jieba.posseg as pseg
words=pseg.cut("又跛又啞")
for w in words:
print w.word,w.flag



输出：
Building Trie..., from C:\Python27\lib\site-packages\jieba\dict.txt
loading model from cache c:\users\test\appdata\local\temp\jieba.cache
loading model cost 1.33399987221 seconds.
Trie has been built succesfully.
又 d
又 d
啞 v
“跛” 字不见了..................................
```
```text
 @elfcool , 多谢提醒，已经修复。麻烦你pull一下，看是否还有问题。
```
```text
 Hi Sun,
已经好了，万分感谢。
祝新年快乐~

Yanlei Deng
East China Normal University
Department of Computer Science
在 2014-01-28 14:07:09，"Sun Junyi" notifications@github.com 写道：
@elfcool , 多谢提醒，已经修复。麻烦你pull一下，看是否还有问题。
—
Reply to this email directly or view it on GitHub.
```
```text
 在用户自定义字典的时候，内容支持正则吗？
```
```text
 @isafe , 目前还不支持。有开源分词组件支持正则词典的？
```
```text
 好像不支持，有个问题想请教下，我用 jieba提取关键字，提取的内容中是中英文混合的
比如 sentence = “beijing,北京欢迎你”
jieba.analyse.extract_tags(sentence,1),比如这样出来的可能是“北京” 但是我想要第一个出来的是欢迎 怎么调整？
2014-07-29 19:02 GMT+08:00 Sun Junyi notifications@github.com:

@isafe https://github.com/isafe , 目前还不支持。有开源分词组件支持正则词典的？
—
Reply to this email directly or view it on GitHub
#172 (comment).
```
```text
 @isafe 這樣你要調整 idf.txt 裡面歡迎的權重值，目前最新版的 jieba 可以切換 idf 語料庫，這樣你就可以調整成你想要的權重值～ 不過建議 idf 的權重值應該還是要自己蒐集足夠量的文本之後計算出每個詞的 idf 權重值會比較客觀
```
```text
 @fxsjy  你说的开源分词组建是哪个？
```
```text
 我倒是也有类似的需求，比如凌晨（5:00-6:00）这样的想分成一个词。
```
## 350.关于标点符号
Jul 15, 2014
```text
 例如：
[笑CRY]  500
将上述整体加入词库中，并进行切分。
目前看，做不到。
是需要做哪些调整吗？词典或代码？
```
```text
 补充：上例中的500是词频。
```
```text
 我改了一下，目前支持词库中的符号和空格匹配了 https://github.com/WalkerWang731/jieba
```
```text
 你好，我很想知道，jieba如何去除标点符号
```
```text
 纯Python实现，无需jieba
punct = set(u''':!),.:;?]}¢'"、。〉》」』】〕〗〞︰︱︳﹐､﹒
﹔﹕﹖﹗﹚﹜﹞！），．：；？｜｝︴︶︸︺︼︾﹀﹂﹄﹏､～￠
々‖•·ˇˉ―--′’”([{£¥'"‵〈《「『【〔〖（［｛￡￥〝︵︷︹︻
︽︿﹁﹃﹙﹛﹝（｛“‘-—_…''')
# 对str/unicode
filterpunt = lambda s: ''.join(filter(lambda x: x not in punct, s))
# 对list
filterpuntl = lambda l: list(filter(lambda x: x not in punct, l))
```
```text
 楼上的回答看上去好酷炫。
```
```text
 吊到无法直视。谢了！
------------------ 原始邮件 ------------------
发件人: "Yanyi Wu";notifications@github.com;
发送时间: 2014年7月19日(星期六) 下午5:46
收件人: "fxsjy/jieba"jieba@noreply.github.com;
抄送: "发如雪止"734133872@qq.com;
主题: Re: [jieba] 关于标点符号 (#169)
楼上的回答看上去好酷炫。
—
Reply to this email directly or view it on GitHub.
```
```text
 我覺得可以幫 jieba 增加一個 jieba.trimPunct(content) 的 method 讓有需要的人可以使用
```
```text
 不好意思，想請問一下類似的問題
由於我目前在處理網頁資料，爬取下來的內容會有一些雜訊，類似
&nbsp;和&gt;

我該如何將其濾掉呢？
```
```text
 不好意思，现在才回复，我觉得你可以先做一遍文本过滤再用jieba分词。
可以先把里面的标点符号过滤掉。
acemoon0301@163.com
发件人： 2153030
发送时间： 2014-12-25 02:06
收件人： fxsjy/jieba
抄送： AcemoonMa
主题： Re: [jieba] 关于标点符号 (#169)
不好意思，想請問一下類似的問題
由於我目前在處理網頁資料，爬取下來的內容會有" "、">" 等雜訊
我該如何將其濾掉呢？
—
Reply to this email directly or view it on GitHub.
```
```text
 @AcemoonMa 不好意思，因為抓取的是html裡的內容，裡面的" "、">"是以字串呈現
空白是 "&+n+b+s+p+;" 將加號都去除，然而用gumblex大的方法會將所有的"&","n","b","s","p",";"都濾掉
昨天我找到了一個方法，可以濾掉字串裡的字串，應該也可以運用在分詞前的濾掉停用字
http://stackoverflow.com/questions/6116978/python-replace-multiple-strings
import re

rep = {"condition1": "", "condition2": "text"} # define desired replacements here

# use these three lines to do the replacement
rep = dict((re.escape(k), v) for k, v in rep.iteritems())
pattern = re.compile("|".join(rep.keys()))
text = pattern.sub(lambda m: rep[re.escape(m.group(0))], text)
For example:
>>> pattern.sub(lambda m: rep[re.escape(m.group(0))], "(condition1) and --condition2--")
'() and --text--'
還沒有實際測試，測試過後跟大家回報
```
```text
 @2153030 HTML 的话最好用 BeautifulSoup 之类的解析库预处理提取字符串（如它的 stripped_strings），再进行分词等自然语言方面的操作。
```
```text
 如果是html解析有很多库都可以提取你需要的text,tag,attribute这些数据啊。
至于网页本身的架构也是可以获取完整的。
例如lxml,beautifulsoap以及python自带的库。
如果你获取后的数据中仍有这些字符，你可以考虑直接写一个字符集合，然后用最基础的循环过滤出来啊。
或者直接用unicode编码过滤，把除了中文，英文，数字以外的都过滤掉就可以了。
acemoon0301@163.com
发件人： 2153030
发送时间： 2014-12-26 15:36
收件人： fxsjy/jieba
抄送： AcemoonMa
主题： Re: [jieba] 关于标点符号 (#169)
@AcemoonMa 不好意思，因為抓取的是html裡的內容，裡面的" "、">"是以字串呈現
空白是 "&+n+b+s+p+;" 將加號都去除，然而用gumblex大的方法會將所有的"&","n","b","s","p",";"都濾掉
昨天我找到了一個方法，可以濾掉字串裡的字串，應該也可以運用在分詞前的濾掉停用字
http://stackoverflow.com/questions/6116978/python-replace-multiple-strings
import re
rep = {"condition1": "", "condition2": "text"} # define desired replacements here
use these three lines to do the replacement
rep = dict((re.escape(k), v) for k, v in rep.iteritems())
pattern = re.compile("|".join(rep.keys()))
text = pattern.sub(lambda m: rep[re.escape(m.group(0))], text)
For example:



pattern.sub(lambda m: rep[re.escape(m.group(0))], "(condition1) and --condition2--")
'() and --text--'
—
Reply to this email directly or view it on GitHub.
```
```text
 你如果有常用的社交或者通讯软件，你可以发软件名和ID，我加你好友，共同探讨，邮件不太方便。
acemoon0301@163.com
发件人： 2153030
发送时间： 2014-12-26 15:36
收件人： fxsjy/jieba
抄送： AcemoonMa
主题： Re: [jieba] 关于标点符号 (#169)
@AcemoonMa 不好意思，因為抓取的是html裡的內容，裡面的" "、">"是以字串呈現
空白是 "&+n+b+s+p+;" 將加號都去除，然而用gumblex大的方法會將所有的"&","n","b","s","p",";"都濾掉
昨天我找到了一個方法，可以濾掉字串裡的字串，應該也可以運用在分詞前的濾掉停用字
http://stackoverflow.com/questions/6116978/python-replace-multiple-strings
import re
rep = {"condition1": "", "condition2": "text"} # define desired replacements here
use these three lines to do the replacement
rep = dict((re.escape(k), v) for k, v in rep.iteritems())
pattern = re.compile("|".join(rep.keys()))
text = pattern.sub(lambda m: rep[re.escape(m.group(0))], text)
For example:



pattern.sub(lambda m: rep[re.escape(m.group(0))], "(condition1) and --condition2--")
'() and --text--'
—
Reply to this email directly or view it on GitHub.
```
```text
 @gumblex @AcemoonMa 謝謝兩位大大的提示
這樣看起來似乎讓實作方便些，我去研究一下BeutifulSoup
@AcemoonMa 郵件已發
```
```text
 楼主 python菜鸟问个简单但急需回答的问题  请问你这个怎么用呢？能写个例子吗？
```
```text
 我想我知道啦  多谢楼主。
```
```text
 import re

def delete_punctuation(text):
    """删除标点符号"""
    text = re.sub(r'[^0-9A-Za-z\u4E00-\u9FFF]+', ' ', text)
    return text

这个满足你的要求吗？
```
## 351.对长标点符号的处理
Jun 22, 2014
```text
 关键词提取无论如何总能得到一些结果，在语料非常多的时候，人工标记不太具备可行性，那么有哪些推荐的算法用于评估提取内容的准确性和相关性呢？
```
```text
 长标点符号应该被分为一个词。
例如：
longpunct = set(u'——', #EM dash U+2014, 常用
u'--', #半角-
u'－－', #U+FF0D, 全角－
u'――', #Horizontal bar, U+2015
u'──', #Box Drawing Light Horizontal, U+2500
u'……',
u'......', # '.'*6
u'...', # '.'*3
) #etc.
```
## 352.生成关键词，里面有一些符号如何处理，如()空格等一些符号
Jun 13, 2014
```text
 No description provided.
```
```text
 https://www.cnblogs.com/adienhsuan/p/5674033.html，可以参考一下
```
```text
 我也遇到这个问题了，求告知一个比较全的词性介绍。
```
```text
 找到一个词性介绍，见jieba
```
```text
 官方应该添加一份词性列表到文档中
```
```text
 同问
```
```text
 应该和这个一样
http://ltp.readthedocs.io/zh_CN/latest/appendix.html#id3
```
```text
 生成关键词，里面有一些符号如何处理，如()空格等一些符号
```
```text
 例如：全角符号，在全模式下不会出现。但是在精确模式下就会出现。
```
```text
 你是想要清掉標點符號嗎？ 像這篇？ #169
```
## 353.是否会出PHP版本的接口？
May 15, 2014
```text
 No description provided.
```
```text
 需求非常大
```
```text
 自己搞个thrift之类的service... 转一下... = =
```
```text
 ls +1
```
```text
 +1
```
```text
 虽然有做广告的嫌疑，但是还是发一下把：
jieba的c++版本有提供http接口供php使用啊：https://github.com/aszxqw/cppjieba
```
```text
 @raykwok @tonicbupt @tonyseek @binaryoung @yanyiwu @fxsjy 我這個週末心血來潮翻譯了 jieba 的 PHP 版本，fukuball/jieba-php，目前是從 python 的 0.16 版翻譯過來的，未來再慢慢往上升級，效能也需要再改善，請有興趣的開發者一起加入開發！
```
```text
 @fukuball 非常感谢
```
## 354.jieba的词性标注结果与ictclas标准不一致
May 7, 2014
```text
 在我的自定义词典中有这样一个专有名词"ens"，在分词时却将“license”强行分成了“lic/ens/e”。这个要怎么处理这种情况？
```
```text
 在词典里加上“license”，再给个比“ens”较高的数
```
```text
 我觉得jieba应该尊重英文词的边界,当自定义词典里起始或结束是英文字母,例如:



jieba.add_word('w底')
jieba.lcut("太好了w底出现了")  # 这里"w底"的w不是英文词的连续,可以切分
['太好了', 'w底', '出现', '了']
jieba.lcut("wow底出现了") # 这里"w底"的w是英文词的连续,不应该切分
['wo', 'w底', '出现', '了']
```
```text
 @chunsheng-chen 那是不是得有个字典存英文单词，感觉会很大哦
```
```text
 不知道具体细节，但我猜测jieba对英文词的分解是基于类似"[a-zA-Z0-9]*"的模式，所以不需要英文字典，例如:



jieba.lcut("this is a 1test1-abc2! call 911")
['this', ' ', 'is', ' ', 'a', ' ', '1test1', '-', 'abc2', '!', ' ', 'call', ' ', '911']



如果能尊重英文词的自然分割方式，就不会出现上面的情况了: license是一个完整的词，wow是一个完整的词。
```
```text
 自定义词典中部分词含有日文的假名，但是分词出来好像全部无视了。
```
```text
 我参看了ictclass关于词性标注的介绍[1]：标点符号都应该标注为以“w”开头的字符串，而jieba目前将它们标记为“x”，即认为是“字符串”。
代码：
list(pseg.cut(u"今天的任务有四项：写程序、看电影和跑步。"))
结果：
[今天/t, 的/uj, 任务/n, 有/v, 四项/m, ：/x, 写/v, 程序/n, 、/x, 看/v, 电影/n, 和/c, 跑步/n, 。/x]
请问这是bug么？
[1] http://ictclas.org/docs/ICTPOS3.0%E6%B1%89%E8%AF%AD%E8%AF%8D%E6%80%A7%E6%A0%87%E8%AE%B0%E9%9B%86.doc
```
```text
 另外，jieba还使用了标签“j”，而“j”并没有出现在ictclas的标准中。
```
```text
 请问uj中的j表示什么含义？
```
## 355.结巴词库不支持拼音很遗憾
May 5, 2014
```text
 No description provided.
```
```text
 我是用正则表达式处理的，new_sentence = re.sub(r'[^\u4e00-\u9fa5]', ' ', old_sentence) 然后再进行分词的, \u4e00-\u9fa5这个是utf-8中，中文编码的范围
```
```text
 @cbzhuang  非常谢谢你的回复！ 我用了这个，不知道可对。#169
```
```text
 Actually, CJK characters are encoded together so there's no critical range for Chinese characters. A punctuation dict could be used to do the filtering.
```
```text
 db中有数据条“酒”，“白酒”，"红酒"
搜索“酒" ,只能出现一条记录”酒"
希望能同时出现“白酒”,"红酒"
```
```text
 这种问题不归这里管吧
```
```text
 rt，其实有时候拼音纠错还是很大作用的，不支持拼音就意味着不支持同音字，非常希望能加入这个功能
```
## 356.增量导入用户自定义词典
Apr 23, 2014
```text
 jieba.add_word("β细胞", freq=True, tag='n')
a = '揭晓！胰岛β细胞功能可以恢复吗?'
WORD = jieba.lcut(a,cut_all=False)
print ([word for word in WORD])
['揭晓', '！', '胰岛', 'β', '细胞', '功能', '可以', '恢复', '吗', '?']
```
```text
 freq的值设置高一点？
```
```text
 freq的值设置高一点？

谢谢，查到了 好像是和特殊字符有关 修改了源代码
```
```text
 jieba是否支持增量导入用户自定义词典？
即在调用load_userdict()之后，再增量地调用某个函数增加用户自定义词？
```
## 357.是否有和Haystack集成的官方说明？
Apr 13, 2014
```text
 class Tokenizer(object):
    ..
    def suggest_freq(self, segment, tune=False):
        ..
        if tune:
            add_word(word, freq)
此处,add_word(word, freq)应为self.add_word(word, freq),否则更新的是默认分词器jieba.dt,而不是用户的new_dt = Tokenizer().
```
```text
 Every time I reboot the Jieba the setting I added using jieba.add_word(word, freq=None, tag=None) will be washed away. Can I keep it somewhere making the setting permanent?
```
```text
 You can write a custom dictionary, and then load the dictionary at start.
```
```text
 当前需要hack Haystack，有点丑：
http://ashin.sinaapp.com/article/119/
http://blog.csdn.net/wenxuansoft/article/details/8170714
```
## 358.词性标注的viterbi的line28和line29的判断条件一致
Apr 4, 2014
```text
 用户自定义词典定义了中文和英文结合的词 ，分词没用效果
比如词典里定义了一个a字群
分词好像没有效果
```
```text
 import jieba
a=jieba.cut("阿里妈妈1");print(" ".join(a))
Building prefix dict from /usr/lib/python2.6/site-packages/jieba-0.37-py2.6.egg/jieba/dict.txt ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.256 seconds.
Prefix dict has been built succesfully.
阿里 妈妈 1
jieba.add_word('阿里妈妈1',100,'nz')
a=jieba.cut("阿里妈妈1");print(" ".join(a))
阿里妈妈1
jieba.del_word('阿里妈妈1')
a=jieba.cut("阿里妈妈1");print(" ".join(a))
阿里妈妈1
```
```text
 这个问题有解决吗?
```
```text
 为什么这两行的判断条件一致呢？
if len(obs_states)==0: obs_states = prev_states_expect_next
if len(obs_states)==0: obs_states = all_states
```
## 359.golang jieba
Apr 4, 2014
```text
 No description provided.
```
```text
 用golang 实现了一下jieba, 在 https://github.com/pastebt/gieba , 欢迎指教
```
## 360.AT&T和T-shirt这类词无法识别，也无法通过加字典来解决
Apr 3, 2014
```text
 sent = "基础上的"
我想把它分词成 基础上/的 。
但是不切分“上的”。
应该怎么实现这样的功能呢？
类似的词还有 “既要有” “越要有”
想分成 既要/有 越要/有
```
```text
 drop word list





来自钉钉专属商务邮箱------------------------------------------------------------------
发件人：OrangeIUH<notifications@github.com>
日　期：2018年10月27日 15:45:09
收件人：fxsjy/jieba<jieba@noreply.github.com>
抄　送：Subscribed<subscribed@noreply.github.com>
主　题：[fxsjy/jieba] 特殊分词控制 (#680)

sent = "基础上的"
 我想把它分词成 基础上/的 。
 但是不切分“上的”。
 应该怎么实现这样的功能呢？
 类似的词还有 “既要有” “越要有”
 想分成 既要/有 越要/有
—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub, or mute the thread.
```
```text
 @fxsjy @tuang
不知道是否有好的解决方法？
```
```text
 目前除了加自定义词典，还不能自动识别。我先做个记号，纳入开发计划。
```
```text
 AT&T是可以的啊
```
```text
 = =
```
```text
 T恤也可以啊。http://jiebademo.ap01.aws.af.cm/
```
```text
 嗯嗯，对的。打开方式问题。感谢耐心指正！
```
```text
 @fxsjy T-shirt 和 google+ 貌似还识别不了
```
## 361.pip安装jieba，出现MemoryError
Apr 1, 2014
```text
 在我的自定义词典中有这样一个专有名词"ens"，在分词时却将“license”强行分成了“lic/ens/e”。这个要怎么处理这种情况？
```
```text
 在词典里加上“license”，再给个比“ens”较高的数
```
```text
 我觉得jieba应该尊重英文词的边界,当自定义词典里起始或结束是英文字母,例如:



jieba.add_word('w底')
jieba.lcut("太好了w底出现了")  # 这里"w底"的w不是英文词的连续,可以切分
['太好了', 'w底', '出现', '了']
jieba.lcut("wow底出现了") # 这里"w底"的w是英文词的连续,不应该切分
['wo', 'w底', '出现', '了']
```
```text
 @chunsheng-chen 那是不是得有个字典存英文单词，感觉会很大哦
```
```text
 不知道具体细节，但我猜测jieba对英文词的分解是基于类似"[a-zA-Z0-9]*"的模式，所以不需要英文字典，例如:



jieba.lcut("this is a 1test1-abc2! call 911")
['this', ' ', 'is', ' ', 'a', ' ', '1test1', '-', 'abc2', '!', ' ', 'call', ' ', '911']



如果能尊重英文词的自然分割方式，就不会出现上面的情况了: license是一个完整的词，wow是一个完整的词。
```
```text
 自定义词典中部分词含有日文的假名，但是分词出来好像全部无视了。
```
```text
 我参看了ictclass关于词性标注的介绍[1]：标点符号都应该标注为以“w”开头的字符串，而jieba目前将它们标记为“x”，即认为是“字符串”。
代码：
list(pseg.cut(u"今天的任务有四项：写程序、看电影和跑步。"))
结果：
[今天/t, 的/uj, 任务/n, 有/v, 四项/m, ：/x, 写/v, 程序/n, 、/x, 看/v, 电影/n, 和/c, 跑步/n, 。/x]
请问这是bug么？
[1] http://ictclas.org/docs/ICTPOS3.0%E6%B1%89%E8%AF%AD%E8%AF%8D%E6%80%A7%E6%A0%87%E8%AE%B0%E9%9B%86.doc
```
```text
 另外，jieba还使用了标签“j”，而“j”并没有出现在ictclas的标准中。
```
```text
 请问uj中的j表示什么含义？
```
```text
 我在windows server 2008的服务器上pip install jieba（jieba-0.32），出现了如下错误：
Installing collected packages: jieba
Running setup.py install for jiaba
Sorry: MemoryError: ( )
Sorry: MemoryError: ( )
Successfully installed jieba
Cleaning up......
这样带来的问题就是，当我用django-haystack + whoosh + jieba的时，用rebuild_index，会出现MemoryError ，而无法生成索引。
我的是python 2.7，希望得到解决，谢谢
```
```text
 实测原来是服务器内存不够，升级服务器内存之后，上述问题均得到了解决
```
## 362.4英寸，7.5ml，这种词是否有办法辨识？
Mar 30, 2014
```text
 上图是我的数据前一部分，我的目的是对 titles 一列进行分词，分词的代码如下。现在遇到的问题是AttributeError: 'int' object has no attribute 'decode'，所以我认为是 titles 中有 int 所致，所以添加了一个判断条件，但是代码执行的结果依旧是报之前的错。请问这是什么原因？
def jiebait(text):
    seglist = jieba.cut(text, cut_all = True)
    fenci = []
    for word in seglist:
        if (not isinstance(word, int)) and (len(word) >= 2):
            fenci.append(word)
    # 如用搜索引擎模式：
    #seglist = jieba.cut_for_search(text)
    return ' '.join(fenci)
```
```text
 def jiebait(text):
    seglist = [str(w) for w in jieba.cut(text, cut_all = True)]
    fenci = []
    for word in seglist:
        if (not isinstance(word, int)) and (len(word) >= 2):
            fenci.append(word)
    # 如用搜索引擎模式：
    #seglist = jieba.cut_for_search(text)
    return ' '.join(fenci)
```
```text
 AttributeError                            Traceback (most recent call last)
 in ()
2     result_line = ""
3     # segment
----> 4     seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]
5     # remove special character
6     temp = re.sub("[\s+.!_,$%^(+"')]+|[+——()?【】“”！，。？、~@#￥%……&（）]+", "",("/".join(seg_list)))
 in (.0)
2     result_line = ""
3     # segment
----> 4     seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]
5     # remove special character
6     temp = re.sub("[\s+.!_,$%^(+"')]+|[+——()?【】“”！，。？、~@#￥%……&（）]+", "",("/".join(seg_list)))
D:\Anaconda\lib\site-packages\jieba_init_.py in cut(self, sentence, cut_all, HMM)
280             - HMM: Whether to use the Hidden Markov Model.
281         '''
--> 282         sentence = strdecode(sentence)
283
284         if cut_all:
D:\Anaconda\lib\site-packages\jieba_compat.py in strdecode(sentence)
35     if not isinstance(sentence, text_type):
36         try:
---> 37             sentence = sentence.decode('utf-8')
38         except UnicodeDecodeError:
39             sentence = sentence.decode('gbk', 'ignore')
AttributeError: 'int' object has no attribute 'decode'
```
```text
 你的 text 本身是不是 int
```
```text
 是的，本身是int，但是seg_list = [str(w) for w in jieba.cut(line, cut_all = False)]这一步，应该都转成string了
```
```text
 解决了 我这边数据转码的问题 不好意思。。。感谢
```
```text
 如题
```
```text
 同問，請教 @fxsjy
```
```text
 同問，請教 @fxsjy
```
```text
 请问有人解决了吗？
```
```text
 我改了一下，目前支持词库中的符号和空格匹配了 https://github.com/WalkerWang731/jieba
```
```text
 在应用时，往往被切分为4, 英寸，7.5, ml
```
```text
 当前我希望获得这种连词的时候，是直接判断量词前是否有数词。有数词自动连接。
但我觉得这是个dirty trick，需要原生支持。
```
```text
 对于这种量词我所呆过公司的做法一般是将这类词当成特殊词语对待，就像品牌有专门的品牌词典和品牌同义词等。所以量词也被当成一个特殊的环节来用特殊的方法对待。一般是用词典和正则。
个人愚见，期待更好的回答。
发自我的 iPhone

在 31 Mar 2014，01:13，"geekan, FSE(Full StackOverflow Engineer)" notifications@github.com 写道：
当前我希望获得这种连词的时候，是直接判断量词前是否有数词。有数词自动连接。
但我觉得这是个dirty trick，需要原生支持。
—
Reply to this email directly or view it on GitHub.
```
```text
 这种量词的可能不是很常见，有一类是AT&T、T恤等常见的品牌或者是中英标点符号共同成词，尽管userdict里有这个词，也无法正确分割。
```
```text
 @aszxqw 正则是对的，但会加倍扫描次数
我猜合并到库里会有比较好的性能
```
```text
 @tuang 这是另外一个问题，我也遇到了。
你开一个issue？
```
```text
 @tuang 我开了一个issue，你看需要怎么补充一下
```
```text
 @geekan新开的issue在哪？
```
```text
 @tuang #145
```
## 363.结巴分词去重
Mar 24, 2014
```text
 结巴的字典是几年前的，不知道能否加入一些新的词重新训练一份？
使用了自定义词典感觉效果不太好。
```
```text
 By theory, jieba is not a trainable model i'm afraid
```
```text
 可以训练啊，你拿别的人工或机器分词/标注的语料，统计一下词频就行。
```
```text
 @gumblex 你好 能详细说说这个过程吗, 我想做一个自己的关键字提取的idf.txt, 但是不知道如何下手, idf.txt中的第二列, 那一串数字, 不知道怎么得出  比如
劳动防护 13.900677652
勞動防護 13.900677652
生化学 13.900677652
```
```text
 Hi,
I am using Jieba in python3, and having the following problem:
for the sentence with word "省", e.g.  "丰田太省了", the results are: 丰田/ 太省/ 了.
However, it should be expected as: 丰田/ 太/ 省/ 了.  it happens also to "最省".
I checked the jieba.dict, and do not find the word "太省" in the dictionary. Could you check if it is an issue in the code? Thank you.
Best Regards,
Qing
```
```text
 >>> tuple(jieba.cut('丰田太省了'))
('丰田', '太省', '了')
>>> tuple(jieba.cut('丰田太省了',HMM=0))
('丰田', '太', '省', '了')

It's a 'new word' found by the HMM model because the model thinks '太省' is probably a word. So is the case of '最'. If you want to get the result only according to the dictionary, set HMM=False.
```
```text
 yes, it works great after set HMM=False. thanks!
```
```text
 请问，现在有去除相同分词的功能吗？是想用结巴来跑一个自己的中文字典，现在的重复词比较多，通过关键词也没办法很好控制，多谢。
```
## 364.词性标注多字
Mar 1, 2014
```text
 如何验证分词的准确性
请问用jieba分完词后，如何去验证分词的准确性？
```
```text
 Same question. What is the training data? 人民日报语料吗？
```
```text
 Same question. What is the training data? 人民日报语料吗？
不是人民日报的，是网上爬取的一些评论数据，最近看了数学之美的分词，有一些启发
```
```text
 一些不字开头的词语在短句中，为啥被extract_tags后"不"字就不见了，将"不"与后面词语分开，这样提取的词的意思都变了，不合理吧
比如

不合适
不适合
不太合适
不太适合
不满意
不想买
不想要
```
```text
 在用词性标注的时候如以下句子：
前港督衛奕信在八八年十月宣布成立中央政策研究組
词性标注后变成
前/f 港督/n 衛/v 衛奕/z 信/n 在/p 八八年/m 十月/t 宣布/v 成立/v 中央/n 政策/n 研究組/n
多出一个衛字
单纯的用分词是OK的，不知道是不是我词性标注的选取的词性标注方法有问题。
我用的是官网上的样例：



import jieba.posseg as pseg
words = pseg.cut("前港督衛奕信在八八年十月宣布成立中央政策研究組")
for w in words:
...    print w.word, w.flag
```
```text
 @elfcool ， 多谢反馈，这是个bug，我刚刚修复了一下：9d4ac26
你有空pull下来看看问题还存不存在。
```
## 365.头发长长了，变成了长头发。
Feb 8, 2014
```text
 比如"中断" "满格" 等会莫名奇妙的识别为ns,ns词性是地名词性，这也太假了
```
```text
 No description provided.
```
```text
 我在维护一个加速版的jieba，有issue和pr欢迎
https://github.com/deepcs233/jieba_fast
```
```text
 您好, 我想在我的论文中引用您的分词器, 请问我可以直接引用github的链接吗? 或者您有发表的作品我需要一并引用? 十分感谢!
```
```text
 提交一个小bug
demo的分词结果是：
头发/n 长长/n 了/ul ，/x 变成/v 了长/v 头发/n 。/x
结果好像不是太好。
请问是否能用自定义词库或者什么方便纠正？
```
## 366.词性标注后，少字
Jan 26, 2014
```text
 textrank中的rank()方法参考了pagerank的思路，并且在程序中iterate 10次。但是对于jieba案例，可否考虑把阻尼因子设定为1，基于这一点，可以推断出每一个节点的weight就是这个点的连线数，没有必要iterate 10反复迭代了。修改方法如下：
    for n, out in self.graph.items():
        ws[n] = wsdef
        outSum[n] = sum((e[2] for e in out), 0.0)

    # this line for build stable iteration
    # sorted_keys = sorted(self.graph.keys())
    # for x in xrange(10):  # 10 iters
    #     for n in sorted_keys:
    #         s = 0
    #         for e in self.graph[n]:
    #             s += e[2] / outSum[e[1]] * ws[e[1]]
    #         ws[n] = (1 - self.d) + self.d * s
    ws=outSum
    (min_rank, max_rank) = (sys.float_info[0], sys.float_info[3])
```
```text
 你好，我最近在看结巴分词的算法，看到在判断切分组合的时候，利用的是动态规划来求出最大组合概率，我觉得这个方法挺好。
不过，我有一个想法，不知道这样想对不对，还请指教。
我是这样想的，因为之前已经形成了切分词图，我在想可不可以利用图论里的最短路径的相关算法来做判断。但是最短路径求得是权值最短，如果把概率作为权值的话，这样直接算就算的是最小概率了。之前我看到过一篇论文有提出一个解决这种问题的办法，利用最短路径算法来求最大权值。
需要对权值做一定的转换。设边的权值为f（也就是词的频率或概率），取f的自然底数的对数的相反数，-ln(f)，作为新的权值w，即w=-ln(f)，利用这个权值w来求最短路径。
可以看出：
1，f越大，w越小
2，min(w1+w2+w3+...+wi)对应的是max(ln(f1_f2_f3_...fi))，也就是说求出有新权值w的图的最短路径后，相应的那个路径的组合概率 P=f1_f2_f3..._fi 是最大的，也就是我们要找的
3，由于f本来就是概率，所以0<f<1，因此w是非负数
所以可以通过求最短路径的方法来得到最大组合概率。
```
```text
 我感觉和我用的算法没有本质不同啊，最开始的时候我是用的概率相乘求最大概率路径，后来为了避免下溢，改为对log(p)求和来算最大概率路径。另外，由于是有向无环图，在有向无环图上求最长路径本来就是可以用动态规划的，不需要在转化为求最短路径问题吧。
```
```text
 import jieba
import jieba.posseg as pseg
words=pseg.cut("又跛又啞")
for w in words:
print w.word,w.flag



输出：
Building Trie..., from C:\Python27\lib\site-packages\jieba\dict.txt
loading model from cache c:\users\test\appdata\local\temp\jieba.cache
loading model cost 1.33399987221 seconds.
Trie has been built succesfully.
又 d
又 d
啞 v
“跛” 字不见了..................................
```
```text
 @elfcool , 多谢提醒，已经修复。麻烦你pull一下，看是否还有问题。
```
```text
 Hi Sun,
已经好了，万分感谢。
祝新年快乐~

Yanlei Deng
East China Normal University
Department of Computer Science
在 2014-01-28 14:07:09，"Sun Junyi" notifications@github.com 写道：
@elfcool , 多谢提醒，已经修复。麻烦你pull一下，看是否还有问题。
—
Reply to this email directly or view it on GitHub.
```
## 367.关于字典
Jan 15, 2014
```text
 No description provided.
```
```text
 需求非常大
```
```text
 自己搞个thrift之类的service... 转一下... = =
```
```text
 ls +1
```
```text
 +1
```
```text
 虽然有做广告的嫌疑，但是还是发一下把：
jieba的c++版本有提供http接口供php使用啊：https://github.com/aszxqw/cppjieba
```
```text
 @raykwok @tonicbupt @tonyseek @binaryoung @yanyiwu @fxsjy 我這個週末心血來潮翻譯了 jieba 的 PHP 版本，fukuball/jieba-php，目前是從 python 的 0.16 版翻譯過來的，未來再慢慢往上升級，效能也需要再改善，請有興趣的開發者一起加入開發！
```
```text
 @fukuball 非常感谢
```
```text
 建议为jieba.posseg增加cut_all选项，允许采用全模式分词状态下再进行词性标注以得到更全的结果
```
```text
 我也需要啊，现在还没有这个api？
```
```text
 首先非常感谢你开源了一个很好用的工具。
最近我在我的网站上开始使用结巴分词来提取我们网站上的每篇内容的 3 个标签，大部分时候效果很好，但是部分主题里出现了一些奇怪的结果（主题正文右下），比如：
http://www.v2ex.com/t/97153 （上换）
http://www.v2ex.com/t/97151 （想省）
http://www.v2ex.com/t/97140 （就够）
所以我在想，如果我把我们网站上的所有内容导出，是否有可能经由你的工具生成一个效果更好的 dict.txt？
谢谢。
```
```text
 jieba的extract_tags使用的分词模式是默认模式，即带有HMM新词发现的模式
感觉其实这样似乎有点不妥，在生产实践中，还是更偏向于信赖词典，因为新词发现容易产生垃圾词（线上如果出现垃圾词往往效果令人无法容忍），相对比起来，新词没有被发现反而可以让人接受（从而再通过改善词典慢慢优化）。
所以建议楼主还是使用非默认模式的jieba.cut 应该会理想一些。
```
```text
 @livid ，你好。
jieba的关键词提取功能比较简陋，是基于最简单的tf/idf排序方式。 对于idf.txt中没有的词汇，它的默认值是取得idf.txt中所有词的idf值的median，这一点搞的比较随意。  也许可以通过降低这个默认值来fix你说的这几个case。
https://github.com/fxsjy/jieba/blob/master/jieba/analyse/__init__.py#L33
```
## 368.关于jieba分词提取权重最大的关键词的示例代码问题。
Jan 13, 2014
```text
 Traceback (most recent call last):
File "", line 1, in 
File "/opt/anaconda3/lib/python3.6/site-packages/jieba/init.py", line 119, in wrapped
return fn(*args, **kwargs)
File "/opt/anaconda3/lib/python3.6/site-packages/jieba/init.py", line 307, in load_userdict
word, freq = tup[0], tup[1]
IndexError: list index out of range
```
```text
 如图，结巴加载后不会往下执行，而是停在最后一行。
(py3env1) ➜  nlpTest python -m jieba -d cutTest.txt > cuted2.txt
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/fv/9t0ldhvx03j4ch64rv3j9qh40000gn/T/jieba.cache
Loading model cost 0.855 seconds.
Prefix dict has been built succesfully.
```
```text
 当去获取词性时，可能有些词没有词性，故会损失准确度
比如“韵达” 如果用分词，就会分出 韵达
如果用获取词性，会得到 韵 m 达  v 没有出现 韵达的词性
我觉得这样并不好，如果没有词性，可否返回一个空或者其它呢，也即尽量不要破坏分词的准度
```
```text
 使用词性标注，分词结果没有精确分词结果准确，比如在精确模式下可以识别出干锅，而在词性标注模式下变成了干和锅两个词，请问如何修正这个问题？
```
```text
 本人接触jieba跟python不久，关于返回权重最大关键词参照了给出的源码 https://github.com/fxsjy/jieba/blob/master/test/extract_tags.py 现在不出效果好像卡在了这一行“usage:    python extract_tags.py [filename] -k [top k]” 。。不太了解这些参数的含义，本人上网查阅未找到解决方法。请教一下有朋友能不能大致讲解一下这个源码的意思，最好能简单结合一下实际案例，如：
encoding=utf-8
import jieba
seg_list = jieba.cut("我来到北京清华大学",cut_all=True)
print "Full Mode:", "/ ".join(seg_list) #全模式
seg_list = jieba.cut("我来到北京清华大学",cut_all=False)
print "Default Mode:", "/ ".join(seg_list) #精确模式
seg_list = jieba.cut("他来到了网易杭研大厦") #默认是精确模式
print ", ".join(seg_list)
seg_list = jieba.cut_for_search("小明硕士毕业于中国科学院计算所，后在日本京都大学深造") #搜索引擎模式
print ", ".join(seg_list)
不胜感激！
```
## 369.大侠，jieba没有词性消歧么？
Jan 6, 2014
```text
 http://jiebademo.ap01.aws.af.cm/
not work
```
```text
 求助；
postgresql添加jieba扩展，提示找不到jieba的 module ，我的Postgresql的版本是8.4，看到github上有pg_jieba的项目，要求postgresql的版本都是9.1以上，请问Jieba对Postgresql有版本限制吗？紧急求助，谢谢！
```
```text
 大侠，你好，最近一直在研读jieba，好像没发现有词性消歧相关的处理。
真的没有么？
```
## 370.结巴分词有命名实体识别功能吗？
Jan 3, 2014
```text
 请问打开新词发现后，发现的新词可以导出吗？
```
```text
 同问，pynlpir有一个get_newword用法，可以print出新词，结巴有没有类似功能？
```
```text
 能否对已经分词的文本单独进行词性标注呢
```
```text
 照 jieba 的做法实现一个：
import re

re_skip_internal = re.compile("(\r\n|\s)")
re_num = re.compile("[\.0-9]+")
re_eng = re.compile("[a-zA-Z0-9]+")

class POSSimpleTagger:

    def __init__(self, dictfile):
        self.word_tag_tab = {}
        with open(dictfile, "rb") as f:
            for lineno, line in enumerate(f, 1):
                try:
                    line = line.strip().decode("utf-8")
                    if not line:
                        continue
                    word, _, tag = line.split(" ")
                    self.word_tag_tab[word] = tag
                except Exception:
                    raise ValueError(
                        'invalid POS dictionary entry in %s at Line %s: %s' % (f_name, lineno, line))

    def tagpos(self, words, HMM=True):
        for w in words:
            if re_skip_internal.match(w):
                yield (w, 'x')
            elif re_num.match(w):
                yield (w, 'm')
            elif re_eng.match(w):
                yield (w, 'eng')
            else:
                yield (w, self.word_tag_tab.get(w, 'x'))
注意字符串要用 unicode(py2), str(py3)。
```
```text
 rt。。。
```
## 371.使用其他字典出现的问题ValueError: too many values to unpack
Dec 24, 2013
```text
 请问：
结巴分词中用户自定义的字典受默认字典词频的影响是和单个词频的值有关么，例如我自定义词典中增加一个词“汉武帝”，词频为30 ，而默认词典中“汉”和“武帝”的词频都是100，那是不是说我这个词一定会被分成默认字典“汉/武帝”，而不是我自定义词表中的“汉武帝”，如果是这样的话，我后期自定义的字典如果批量加入，我怎么定词频的值合适，还是我需要对一类词进入默认词表中先检查一下，在定义自己的字典，不知道有没有什么好的办法，望回复，感谢！！
```
```text
 我也有这个问题，请问您解决了吗？
```
```text
 我解决了，你在用户自定义词典中不要规定词频，如下
体外药物释放
增效剂
航道整治
肌腱膜纤维肉瘤癌基因同系物A
这样就会自动生成一个合适的词频，从而可以把这个词分出来
```
```text
 @我也碰到了类似的问题，使用楼上的@grandmoi 的方法不写词频和词性会报下面这个错误
IndexError: list index out of range
不写词性，词频写多高都分不出词
难不成真的要自己一个个加到字典里……
```
```text
 @grandmoi 我也有这个问题。我的自定义词典中定义了‘国债’，没有设置词频。然后输入‘中国债’，想分出来‘中’、‘国债’。但是实际输出还是‘中国’、‘债’
```
```text
 我按照规范建立了自己的字典，因为业务需要，我创建了两个字典，分别名字为a.txt ，b.txt 这两个文件分表包括了两个数据表格里面的数据，自己测试了一下不能如果采用以下的方法加载两个文件，好像不能同时生效


import jieba
jieba.load_userdict('a.txt')
jieba.load_userdict('a.txt')


请问结巴能否同时加载两个用户字典呢？？
```
```text
 今天又遇到了问题，我在我的flask web app 中使用了jieba的加载自定义字典功能，然后用下面的命令启动


gunicorn -w 4 -p gevent -b 0.0.0.0:9999 --reload run:app
发现jieba连续不断的吐出下面的提示，我觉得应该是gunicorn开启了多个线程导致了这个问题，我想请教下，该如何解决？


loading model from cache /tmp/jieba.cache
loading model cost 2.44023799896 seconds.
Trie has been built succesfully.
[2016-09-29 17:05:37 +0000] [32528] [INFO] Booting worker with pid: 32528
Building Trie..., from /root/py27/lib/python2.7/site-packages/jieba/dict.txt
loading model from cache /tmp/jieba.cache
loading model cost 2.28571200371 seconds.
Trie has been built succesfully.
[2016-09-29 17:06:06 +0000] [32556] [INFO] Booting worker with pid: 32556
Building Trie..., from /root/py27/lib/python2.7/site-packages/jieba/dict.txt
loading model from cache /tmp/jieba.cache
loading model cost 2.27150511742 seconds.
Trie has been built succesfully.
[2016-09-29 17:06:10 +0000] [32560] [INFO] Booting worker with pid: 32560
Building Trie..., from /root/py27/lib/python2.7/site-packages/jieba/dict.txt
loading model from cache /tmp/jieba.cache
```
```text
 gunicorn会fork多个进程，但是jieba是lazy加载词典的。你可以在import jieba后，调用一下jieba.initialize()。 这样就不会多次加载了。
```
```text
 同样也是jieba load_dict的问题，我发现我自己在词典中添加了一个词并设定了参数比如：萌萌哒 50 a，但是使用posseg分词的结果却是 萌萌哒 x，这是版本问题还是其他设定的问题？
```
```text
 @fxsjy
具体的代码用到了这几个部分
import jieba
jieba.initialize()
import os
if os.path.exists('cbi360.txt'):
jieba.load_userdict('cbi360.txt')
import jieba.posseg as peg
其中 cbi360.txt是我的自己的字典，而且我还用刀了jieba.posseg 的方法，请问这个具体的顺序是怎么样的啊？
```
```text
 你好，我使用其他字典，字典是根据你的格式来改的！
```
```text
 你是用空格分隔的还是用tab分隔的？我用tab分隔也出现了这个问题
```
```text
 还有，总共三个字段，虽然文档上面说最后一个词性可以省略，但是我看源码，貌似不能省略，要不然也会报错（我用Python3）
```
```text
 我是修改源码来解决的...
修改jieba: /usr/lib/python2.7/site-packages/jieba/analyse/init.py 的get_idf(abs_path), 在第29行加上:
if lines and not lines[-1]:                                     
    lines.pop(-1)

最后是这个样子:
  lines = content.split('\n')
  if lines and not lines[-1]:
      lines.pop(-1)
  for line in lines:
      word,freq = line.split(' ')
      idf_freq[word] = float(freq)
```
## 372.jieba分词的学习能力如何体现
Dec 12, 2013
```text
 No description provided.
```
```text
 你没看代码吗？
```
```text
 哎，作者就是不爱写注释，哎。
```
```text
 @ilqxejraha 谢谢您，可能是我表达有误，我是想问词汇库里面的词汇和词频是人工统计的吗？还是通过其它的方法。
```
```text
 统计过来的。最后得到了就是这么一个模型。
```
```text
 你在源码中看到词频的使用了吗？
```
```text
 具体的统计词频，词频的作用可能体现在，一个词存在多个意思。
比如英语中，经常有一个词会有很多个意思。
比如出现一个词，并且，这个词有很多种解释，这时候词频可能会对词意的选择有一定帮助。
具体的其他算法hmm的我还没看。
```
```text
 @KevinDotW 据说是基于人民日报的语料库，我也想知道怎么才能创建自己的词典
```
```text
 词性标注是否可以指定分词模式；
想指定搜索模式
谢谢。
```
```text
 比如微博的表情文字是   [泪]
我想拿到这个表情统计个数，请问可以怎么做？
```
```text
 想要有个词语与词语的关系，用来做分析？或者可以一同开发？请指点下，或我留个邮箱？
```
```text
 词语与词语间关系 是基于词性还是完全不考虑词性呢
```
```text
 基于的，比如钱包，近义词是钱夹，属于：皮革制品，纺织品，配件。。。
```
```text
 @c2h2 几乎不可能的，即使有准确率也不高，因为这问题我觉得还是主要靠人工review维护一个同义词词典。虽然这个应用的需求很强烈。
```
```text
 @c2h2 ， 试一试Google的word2vec ?
```
```text
 各位，新接触这块领域，多多指教。
对于jieba的学习新词的能力，是会自动讲本次分词过程中的新词自动加入默认的词典，还是其他方式体现。
谢谢。
Kane
```
```text
 刚详细看了下几个问题列表，对https://github.com/fxsjy/jieba/issues/7中提到的回复也看了下，
基本知道jieba的新词学习能力跟hmm有关系，也试了下finalseg的新词学习功能：
姚晨和老凌离婚了
张绍刚发道歉信网友不认可
结果：
/ 姚晨/ 和/ 老凌/ 离婚/ 了/
/ 张绍/ 刚发/ 道歉信/ 网友/ 不/ 认可/
基本知道jieba默认分词是打开了新词学习的，想问下：
1、姚晨这个例子，是否可以讲学习到的新词自动补充到默认词典中？
1、第二个错误，需要自己补充词典来解决这个错误么？
望不吝赐教，谢谢。
```
```text
 @xmkane , 现在结巴分词并不能很好地处理一些歧义case，解决办法暂时只有加词典条目。 自定义词典如何添加可以参考wiki，另外git repository中的最新版本也支持调用add_word加词条   #122 。
另外，你举的【张绍刚发道歉信网友不认可】这个例子可以用jieba分词子模块posseg来试一试，它的新词识别能力比较强，但是速度要慢一些。
张绍刚/nr 发/v 道歉信/vn 网友/n 不/d 认可/v
http://jiebademo.ap01.aws.af.cm/  （选择“显示词性”）
```
```text
 @fxsjy ,我最近也在jieba，感觉很好，但也不清楚jieba的新词发现机制，简单测试了一下，貌似新词发现能力和文本长度没有关系？我以前了解过基于统计的新词发现方法，比如考虑词的内聚性和自由度，这就文本越多分词越准确，但我感觉jieba使用的不是这个原理？可以稍微说一下jieba的新词发现机制么？非常感谢呢～～
```
## 373.如何控制精确分词的粒度，尽可能小
Dec 4, 2013
```text
 就是我有一个给定的关键词库，然后新来一篇文档，从词库里面找出几个词语作为这篇文档的关键词。
```
```text
 目前我就在做类似的工作，效果挺不错。
你的这种情况，可以这么做。

先将文档按句子切分成多个句子，然后计算关键词库中的每个关键词在这篇文档中的句子集合；
接着计算关键词库中两两关键词的相似性（可以用Jaccard相似性度量），这样构成了一个相似性矩阵；
接着对相似性矩阵进行特征分解，然后对特征值进行归一化；
对归一化的特征值从大到小排序，并累计求和（cumsum），选取前<=0.8的特征值对应的关键词作为这篇文档的关键词

以上是一个基本的版本，直接用，效果一般。因此，需要考虑关键词的tf-idf。我是这么做的，在上面的步骤3时，对特征值进行tfidf加权。最后实验结果很好。
以上，楼主可以试试。
```
```text
 @MacQing 非常感谢，我试一下。
```
```text
 我在Pyspark中引用jieba，用jieba.load_userdict('xxx.txx')加载了自己的词典，但是发现RDD.map操作jieba的时候，该词典没有加载成功。于是，我定义了一个jieba_instance = jieba.Tokernice()，通过传参的方式将定义个一个jieba实例传递到RDD的Map操作函数里面。操作如下:
ss = rdd_data.map(lambda x:judge_disease(x,column,disease_connect,table_name,jieba_instance ))
judge_disease方法如下：
def judge_disease(x,column,disease_connect,table_name,jieba_instance ):
    arr = []
    if not column:
        return (x,None,table_name)
    seg = jieba_instance .cut(x[column.encode('utf-8')])
    for item in seg:
        item = item.encode('utf-8')
        if disease_connect.has_key(item):
            arr.append(unicode(item,'utf-8'))
    arr = list(set(arr))
    res = (x,"$$".join(arr).split("$$"),table_name)

    return res

报如下错误：
‘UnpicklingError: NEWOBJ class argument has NULL tp_new’
求助：首先，我该如何通过传参的方式，将一个jieba实例传入函数？
另外，有碰到相关问题的其它解决方法吗？
```
```text
 机器运行环境：
jieba + python2.7+spark1.6
在用jieba load user dict的时候，发现该dict加载后是以python 字典的数据类型存在于spark的driver里面，但是worker进程无法共享这段内存。同时，我也发现jieba分词调用的源数据是jieba.cache和FREQ（dictionary）。cache文件当set_dictionary的时候，是生成jieba.xxxxxxxxxxxx.cache，该文件在worker进程也无法访问。所以，spark里面，运用常规方法无法访问jieba自己的词典。
所以，特提供自己的意见，看能否在版本中进行修改。
1.可以直接加载python字典类型的函数，这样，我可以将该字典在spark里面共享，然后在spark的worker进程独立加载。
2.增加一种模式，使得在类似spark这种方式里面，可以直接选择如何生成jieba.cache，比如，可以直接将add_word或者load_userdict新增的词，放到该jieba.cache里面。
以上是我个人的浅见，希望继续交流。
```
```text
 不用传jieba实例，在judge_disease中初始化jieba实例呢？
```
```text
 @snowlord 那相当于每次都得新建jieba实例，对spark的map操作来说，太耗费资源和时间。
```
```text
 请问，楼主解决了吗？怎么解决的呢？ @liaicheng
```
```text
 把默认的dict.txt替换成自己的dict.txt。
```
```text
 我也尝试了这么做，但是dict.txt文件需要 词频 吧？ 我测试的，如果不添加词频，报错。
那么 对于自己词典中的词，没有词频，楼主添加的默认的吗？ 比如 3？ 还是怎么解决的呢？ @liaicheng
```
```text
 没有用过spark，估计是不是由于全局变量引起的一些问题？
建议用jieba.Tokenizer 得到一个分词器实例，然后再调用相关方法。
jieba.Tokenizer(dictionary=DEFAULT_DICT) 新建自定义分词器，可用于同时使用不同词典。jieba.dt 为默认分词器，所有全局分词相关函数都是该分词器的映射。
例子：https://github.com/fxsjy/jieba/blob/0243d568e9421ab7d3c75f49e9adfc230810e0a3/test/test_lock.py
```
```text
 @fxsjy Spark中使用tokr.cut()会报can't pickle thread.lock objects错误
@liaicheng 把默认的dict.txt替换成自己的dict.txt，你的意思是jieba.set_dictionary("dict\dict.txt")
jieba.initialize() 这样？但是我试了还是不行，每次在map中使用jieba时，还是会执行Building prefix dict from the default dictionary ……
```
```text
 请问我是将jieba分词包装到worker上了，这么做的话会有什么问题吗？效率会低吗？？然后用udf函数去解决这个问题
split_precision = udf(lambda value:[token for token in (jieba.cut(value.encode('utf-8'), cut_all=False))],ArrayType(StringType()))
split_result = df.withColumn("tokens"+'_'+column, split_precision(col(column)))
最后我想问一个问题，如果我想增加某个词的权重以让该词能够分出来的话，我在driver的程序里写了jieba.add_word(df_dict.iloc[i][0],freq=10000)
但最后没有将词分出来，猜想是没有将词加载到worker的词库里，所以没有将词分出来。。那么该怎么去做呢？
```
```text
 我也遇到這個問題，我的解法是：

把 /usr/local/lib/python2.7/dist-packages/jieba/ 裡面的 dict.txt 替換成自己的字典
把 /tmp/ 裡面的  jieba.cache 清掉

然後就可以了，參考看看
```
```text
 有解决方案了吗@liaicheng
```
```text
 我把版本为调低到0.36就不报‘UnpicklingError: NEWOBJ class argument has NULL tp_new’ 了
```
```text
 同样遇到pyspark加载用户自定义词典，未生效，有什么解决方案吗？
```
```text
 16年的bug，到现在都没解决么
```
```text
 我也遇到這個問題，我的解法是：

把 /usr/local/lib/python2.7/dist-packages/jieba/ 裡面的 dict.txt 替換成自己的字典
把 /tmp/ 裡面的  jieba.cache 清掉

然後就可以了，參考看看

这个cache需要去每一台机器上清理么？
```
```text
 问题：jieba分词器spark分布式环境中加载词典失效(不会报错，但是不会加载词典分出想要的词，这个也是我用了一年的pyspark+jieba之后才突然发现自己用错了一年)
尝试过的方法：

修改默认词典dict.txt。
失败，测试发现jieba并不加载自己的默认词典！！！
将jieba作为参数传入柯里化的udf。
失败，jieba中有线程锁，无法序列化。
新建一个jieba的Tokenizer实例，作为参数传入udf或者作为全局变量。
失败原因同2
使用.rdd.mapPartiton算子, 在每个partiton中加载词典。
失败原因同2
在udf中通过jieba.dt.initialized判断是否需要加载词典
成功。
原因：通过本地测试（这样executor日志也会痛driver日志一起打印出来），
发现，结巴会在每个partiton中分别初始化一个默认分词器。
jieba采用延迟加载机制，在没有使用jieba去分词或加载词典时，
jieba中的默认分词器，jieba.dt不会初始化。
jieba.dt.initialized属性在初始化后会从False变为True，
所以可以依据这个来判断jieba是否初始化。从而决定是否加载词典。

import jieba
from pyspark.sql import SparkSession, functions as F
from pyspark.sql.types import ArrayType, StringType

spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

df = spark.createDataFrame([
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
], ["id", "content"])
df.show()

jieba.load_userdict(
    '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')


def _wrong_tokenize(content):
    """错误的分词函数。此udf不会加载词典，"不用加水光",
    其中的"水光"会被切开，结果为[不用, 加水, 光, ？]

    :param content: {str} 要分词的句子
    :return: list[word, word, ...]
    """
    return [i for i in jieba.cut(content)]


wrong_tokenize = F.udf(_wrong_tokenize, ArrayType(StringType()))
df.select(wrong_tokenize('content').alias('words')).show()
print('df 的partition 数为：', df.rdd.getNumPartitions())
输出如下，“水光”一词被切开了，说明加载词典失效了。但是惊喜地发现jieba在每个partiton中只初始化了一次。
+---+--------+
| id| content|
+---+--------+
|  0|  不用加水光？|
|  0|  不用加水光？|
|  0|  不用加水光？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
+---+--------+

[Stage 2:>                                                          (0 + 1) / 1]Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Loading model cost 1.588 seconds.
Prefix dict has been built succesfully.
[Stage 3:>                                                          (0 + 3) / 3]Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Loading model cost 5.942 seconds.
Prefix dict has been built succesfully.
[Stage 3:===================>                                       (1 + 2) / 3]Loading model cost 6.045 seconds.
Prefix dict has been built succesfully.
[Stage 3:=======================================>                   (2 + 1) / 3]Loading model cost 6.320 seconds.
Prefix dict has been built succesfully.
+--------------------+
|               words|
+--------------------+
|      [不用, 加水, 光, ？]|
|      [不用, 加水, 光, ？]|
|      [不用, 加水, 光, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|      [单纯, 水, 光, 多少]|
|      [单纯, 水, 光, 多少]|
|      [单纯, 水, 光, 多少]|
+--------------------+
df 的partition 数为： 4
import jieba
from pyspark.sql import SparkSession, functions as F
from pyspark.sql.types import ArrayType, StringType

spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

df = spark.createDataFrame([
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
], ["id", "content"])
df.show()


def _tokenize(content):
    """此udf在每个partiton中加载一次词典，
    "不用加水光", 其中的"水光"会被保留，结果为[不用, 加, 水光, ？]

    :param content: {str} 要分词的内容
    :return: list[word, word, ...]
    """
    if not jieba.dt.initialized:
        # 词典中有"水光"这个词
        jieba.load_userdict(
            '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')
    return [i for i in jieba.cut(content)]


tokenize = F.udf(_tokenize, ArrayType(StringType()))
df.select(tokenize('content').alias('words')).show()
print('df 的partition 数为：', df.rdd.getNumPartitions())

+---+--------+
| id| content|
+---+--------+
|  0|  不用加水光？|
|  0|  不用加水光？|
|  0|  不用加水光？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
+---+--------+

Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
[Stage 2:>                                                          (0 + 1) / 1]Loading model cost 0.804 seconds.
Prefix dict has been built succesfully.
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
[Stage 3:>                                                          (0 + 3) / 3]Loading model cost 1.204 seconds.
Prefix dict has been built succesfully.
Loading model cost 1.200 seconds.
Prefix dict has been built succesfully.
Loading model cost 1.224 seconds.
Prefix dict has been built succesfully.
+------------------+
|             words|
+------------------+
|    [不用, 加, 水光, ？]|
|    [不用, 加, 水光, ？]|
|    [不用, 加, 水光, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|      [单纯, 水光, 多少]|
|      [单纯, 水光, 多少]|
|      [单纯, 水光, 多少]|
+------------------+

df 的partition 数为： 4
@fxsjy @liaicheng 至此，这个issue应该可以关闭了，jieba没有问题，saprk也没有问题，但是spark+jieba就产生了问题。建议作者大大把下面这个udf的写法，添加到官方的readme中，毕竟这个bug不会报错，难以感知，有很多想我一样的小白，错误地使用了一年的pyspark+jieba却没有发现问题。
def _tokenize(content):
    """此udf在每个partiton中加载一次词典"""
    if not jieba.dt.initialized:
        jieba.load_userdict('user_dict.txt')
    return [i for i in jieba.cut(content)]


tokenize = F.udf(_tokenize, ArrayType(StringType()))
```
```text
 This is an awesome answer. You could write a sample program and request a
merge.

Actually this is referring to understanding of spark behavior and your
sample is valuablly helpful.
…
On Wed, Dec 5, 2018, 10:27 Chant00 ***@***.*** wrote:

 问题：jieba分词器spark分布式环境中加载词典失效(不会报错，但是不会加载词典分出想要的词，这个也是我用了一年的pyspark+jieba之后才突然发现自己用错了一年)
 尝试过的方法：

    1. 修改默认词典dict.txt。
    失败，测试发现jieba并不加载自己的默认词典！！！
    2. 将jieba作为参数传入柯里化的udf。
    失败，jieba中有线程锁，无法序列化。
    3. 新建一个jieba的Tokenizer实例，作为参数传入udf或者作为全局变量。
    失败原因同2
    4. 使用.rdd.mapPartiton算子, 在每个partiton中加载词典。
    失败原因同2
    5. 在udf中通过jieba.dt.initialized判断是否需要加载词典
    成功。
    原因：通过本地测试（这样executor日志也会痛driver日志一起打印出来），
    发现，结巴会在每个partiton中分别初始化一个默认分词器。
    jieba采用延迟加载机制，在没有使用jieba去分词或加载词典时，
    jieba中的默认分词器，jieba.dt不会初始化。
    jieba.dt.initialized属性在初始化后会从False变为True，
    所以可以依据这个来判断jieba是否初始化。从而决定是否加载词典。

 import jiebafrom pyspark.sql import SparkSession, functions as Ffrom pyspark.sql.types import ArrayType, StringType

 spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

 df = spark.createDataFrame([
     (0, '不用加水光？'),
     (0, '不用加水光？'),
     (0, '不用加水光？'),
     (1, '加水光多少钱啊？'),
     (1, '加水光多少钱啊？'),
     (1, '加水光多少钱啊？'),
     (2, '单纯水光多少'),
     (2, '单纯水光多少'),
     (2, '单纯水光多少'),
 ], ["id", "content"])
 df.show()

 jieba.load_userdict(
     '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')

 def _wrong_tokenize(content):
     """错误的分词函数。此udf不会加载词典，"不用加水光",    其中的"水光"会被切开，结果为[不用, 加水, 光, ？]    :param content: {str} 要分词的句子    :return: list[word, word, ...]    """
     return [i for i in jieba.cut(content)]


 wrong_tokenize = F.udf(_wrong_tokenize, ArrayType(StringType()))
 df.select(wrong_tokenize('content').alias('words')).show()print('df 的partition 数为：', df.rdd.getNumPartitions())

 输出如下，“水光”一词被切开了，说明加载词典失效了。但是惊喜地发现jieba在每个partiton中只初始化了一次。
 +---+--------+| id| content|
 +---+--------+|  0|  不用加水光？||  0|  不用加水光？||  0|  不用加水光？||  1|加水光多少钱啊？||  1|加水光多少钱啊？||  1|加水光多少钱啊？||  2|  单纯水光多少||  2|  单纯水光多少||  2|  单纯水光多少|
 +---+--------+

 [Stage 2:>                                                          (0 + 1) / 1]Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Loading model cost 1.588 seconds.
 Prefix dict has been built succesfully.
 [Stage 3:>                                                          (0 + 3) / 3]Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Loading model cost 5.942 seconds.
 Prefix dict has been built succesfully.
 [Stage 3:===================>                                       (1 + 2) / 3]Loading model cost 6.045 seconds.
 Prefix dict has been built succesfully.
 [Stage 3:=======================================>                   (2 + 1) / 3]Loading model cost 6.320 seconds.
 Prefix dict has been built succesfully.
 +--------------------+|               words|
 +--------------------+|      [不用, 加水, 光, ？]||      [不用, 加水, 光, ？]||      [不用, 加水, 光, ？]||[加水, 光, 多少, 钱, 啊, ？]||[加水, 光, 多少, 钱, 啊, ？]||[加水, 光, 多少, 钱, 啊, ？]||      [单纯, 水, 光, 多少]||      [单纯, 水, 光, 多少]||      [单纯, 水, 光, 多少]|
 +--------------------+
 df 的partition 数为： 4

 import jiebafrom pyspark.sql import SparkSession, functions as Ffrom pyspark.sql.types import ArrayType, StringType

 spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

 df = spark.createDataFrame([
     (0, '不用加水光？'),
     (0, '不用加水光？'),
     (0, '不用加水光？'),
     (1, '加水光多少钱啊？'),
     (1, '加水光多少钱啊？'),
     (1, '加水光多少钱啊？'),
     (2, '单纯水光多少'),
     (2, '单纯水光多少'),
     (2, '单纯水光多少'),
 ], ["id", "content"])
 df.show()

 def _tokenize(content):
     """此udf在每个partiton中加载一次词典，    "不用加水光", 其中的"水光"会被保留，结果为[不用, 加, 水光, ？]    :param content: {str} 要分词的内容    :return: list[word, word, ...]    """
     if not jieba.dt.initialized:
         # 词典中有"水光"这个词
         jieba.load_userdict(
             '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')
     return [i for i in jieba.cut(content)]


 tokenize = F.udf(_tokenize, ArrayType(StringType()))
 df.select(tokenize('content').alias('words')).show()print('df 的partition 数为：', df.rdd.getNumPartitions())

 +---+--------+| id| content|
 +---+--------+|  0|  不用加水光？||  0|  不用加水光？||  0|  不用加水光？||  1|加水光多少钱啊？||  1|加水光多少钱啊？||  1|加水光多少钱啊？||  2|  单纯水光多少||  2|  单纯水光多少||  2|  单纯水光多少|
 +---+--------+

 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 [Stage 2:>                                                          (0 + 1) / 1]Loading model cost 0.804 seconds.
 Prefix dict has been built succesfully.
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 Building prefix dict from the default dictionary ...
 Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
 [Stage 3:>                                                          (0 + 3) / 3]Loading model cost 1.204 seconds.
 Prefix dict has been built succesfully.
 Loading model cost 1.200 seconds.
 Prefix dict has been built succesfully.
 Loading model cost 1.224 seconds.
 Prefix dict has been built succesfully.
 +------------------+|             words|
 +------------------+|    [不用, 加, 水光, ？]||    [不用, 加, 水光, ？]||    [不用, 加, 水光, ？]||[加, 水光, 多少钱, 啊, ？]||[加, 水光, 多少钱, 啊, ？]||[加, 水光, 多少钱, 啊, ？]||      [单纯, 水光, 多少]||      [单纯, 水光, 多少]||      [单纯, 水光, 多少]|
 +------------------+

 df 的partition 数为： 4

 @fxsjy <https://github.com/fxsjy> @liaicheng
 <https://github.com/liaicheng>
 至此，这个issue应该可以关闭了，jieba没有问题，saprk也没有问题，但是spark+jieba就产生了问题。建议作者大大把下面这个udf的写法，添加到官方的readme中，毕竟这个bug不会报错，难以感知，有很多想我一样的小白，错误地使用了一年的pyspark+jieba却没有发现问题。

 def _tokenize(content):
     """此udf在每个partiton中加载一次词典"""
     if not jieba.dt.initialized:
         jieba.load_userdict('user_dict.txt')
     return [i for i in jieba.cut(content)]


 tokenize = F.udf(_tokenize, ArrayType(StringType()))

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#387 (comment)>, or mute
 the thread
 <https://github.com/notifications/unsubscribe-auth/AA0SqpifJt8gBxevt5TSMZ-cvH2UuJz7ks5u1y8PgaJpZM4JkN9I>
 .
```
```text
 问题：jieba分词器spark分布式环境中加载词典失效(不会报错，但是不会加载词典分出想要的词，这个也是我用了一年的pyspark+jieba之后才突然发现自己用错了一年)
尝试过的方法：

修改默认词典dict.txt。
失败，测试发现jieba并不加载自己的默认词典！！！
将jieba作为参数传入柯里化的udf。
失败，jieba中有线程锁，无法序列化。
新建一个jieba的Tokenizer实例，作为参数传入udf或者作为全局变量。
失败原因同2
使用.rdd.mapPartiton算子, 在每个partiton中加载词典。
失败原因同2
在udf中通过jieba.dt.initialized判断是否需要加载词典
成功。
原因：通过本地测试（这样executor日志也会痛driver日志一起打印出来），
发现，结巴会在每个partiton中分别初始化一个默认分词器。
jieba采用延迟加载机制，在没有使用jieba去分词或加载词典时，
jieba中的默认分词器，jieba.dt不会初始化。
jieba.dt.initialized属性在初始化后会从False变为True，
所以可以依据这个来判断jieba是否初始化。从而决定是否加载词典。

import jieba
from pyspark.sql import SparkSession, functions as F
from pyspark.sql.types import ArrayType, StringType

spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

df = spark.createDataFrame([
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
], ["id", "content"])
df.show()

jieba.load_userdict(
    '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')


def _wrong_tokenize(content):
    """错误的分词函数。此udf不会加载词典，"不用加水光",
    其中的"水光"会被切开，结果为[不用, 加水, 光, ？]

    :param content: {str} 要分词的句子
    :return: list[word, word, ...]
    """
    return [i for i in jieba.cut(content)]


wrong_tokenize = F.udf(_wrong_tokenize, ArrayType(StringType()))
df.select(wrong_tokenize('content').alias('words')).show()
print('df 的partition 数为：', df.rdd.getNumPartitions())
输出如下，“水光”一词被切开了，说明加载词典失效了。但是惊喜地发现jieba在每个partiton中只初始化了一次。
+---+--------+
| id| content|
+---+--------+
|  0|  不用加水光？|
|  0|  不用加水光？|
|  0|  不用加水光？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
+---+--------+

[Stage 2:>                                                          (0 + 1) / 1]Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Loading model cost 1.588 seconds.
Prefix dict has been built succesfully.
[Stage 3:>                                                          (0 + 3) / 3]Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Loading model cost 5.942 seconds.
Prefix dict has been built succesfully.
[Stage 3:===================>                                       (1 + 2) / 3]Loading model cost 6.045 seconds.
Prefix dict has been built succesfully.
[Stage 3:=======================================>                   (2 + 1) / 3]Loading model cost 6.320 seconds.
Prefix dict has been built succesfully.
+--------------------+
|               words|
+--------------------+
|      [不用, 加水, 光, ？]|
|      [不用, 加水, 光, ？]|
|      [不用, 加水, 光, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|[加水, 光, 多少, 钱, 啊, ？]|
|      [单纯, 水, 光, 多少]|
|      [单纯, 水, 光, 多少]|
|      [单纯, 水, 光, 多少]|
+--------------------+
df 的partition 数为： 4
import jieba
from pyspark.sql import SparkSession, functions as F
from pyspark.sql.types import ArrayType, StringType

spark = SparkSession.builder.appName("word2vec_cluster").getOrCreate()

df = spark.createDataFrame([
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (0, '不用加水光？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (1, '加水光多少钱啊？'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
    (2, '单纯水光多少'),
], ["id", "content"])
df.show()


def _tokenize(content):
    """此udf在每个partiton中加载一次词典，
    "不用加水光", 其中的"水光"会被保留，结果为[不用, 加, 水光, ？]

    :param content: {str} 要分词的内容
    :return: list[word, word, ...]
    """
    if not jieba.dt.initialized:
        # 词典中有"水光"这个词
        jieba.load_userdict(
            '/Users/chant/sy/dwh/utils/static/soyoung.item.jieba.dic')
    return [i for i in jieba.cut(content)]


tokenize = F.udf(_tokenize, ArrayType(StringType()))
df.select(tokenize('content').alias('words')).show()
print('df 的partition 数为：', df.rdd.getNumPartitions())
+---+--------+
| id| content|
+---+--------+
|  0|  不用加水光？|
|  0|  不用加水光？|
|  0|  不用加水光？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  1|加水光多少钱啊？|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
|  2|  单纯水光多少|
+---+--------+

Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
[Stage 2:>                                                          (0 + 1) / 1]Loading model cost 0.804 seconds.
Prefix dict has been built succesfully.
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/px/cyksvc795bb2qr4rg29s4y5m0000gn/T/jieba.cache
[Stage 3:>                                                          (0 + 3) / 3]Loading model cost 1.204 seconds.
Prefix dict has been built succesfully.
Loading model cost 1.200 seconds.
Prefix dict has been built succesfully.
Loading model cost 1.224 seconds.
Prefix dict has been built succesfully.
+------------------+
|             words|
+------------------+
|    [不用, 加, 水光, ？]|
|    [不用, 加, 水光, ？]|
|    [不用, 加, 水光, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|[加, 水光, 多少钱, 啊, ？]|
|      [单纯, 水光, 多少]|
|      [单纯, 水光, 多少]|
|      [单纯, 水光, 多少]|
+------------------+

df 的partition 数为： 4
@fxsjy @liaicheng 至此，这个issue应该可以关闭了，jieba没有问题，saprk也没有问题，但是spark+jieba就产生了问题。建议作者大大把下面这个udf的写法，添加到官方的readme中，毕竟这个bug不会报错，难以感知，有很多想我一样的小白，错误地使用了一年的pyspark+jieba却没有发现问题。
def _tokenize(content):
    """此udf在每个partiton中加载一次词典"""
    if not jieba.dt.initialized:
        jieba.load_userdict('user_dict.txt')
    return [i for i in jieba.cut(content)]


tokenize = F.udf(_tokenize, ArrayType(StringType()))

想请问你这个是分布式集群环境做的吗？
```
```text
 @Chant00
用此方法成功了~~感谢！
sc.addPyFile('file_path/userdict.txt') 
def seg_sentence(sentence, stopwords): 
sentence = re.sub("[A-Za-z0-9\!\%\[\]\,\。\@]", "", sentence) 
if not jieba.dt.initialized:
 jieba.load_userdict("userdict.txt") 
sentence_seged = jieba.cut(sentence.strip()) 
outstr = [] 
for word in sentence_seged:
 if word not in stopwords: 
if (word != '\t') or (word != "，"): 
outstr.append(word) 
return outstr
```
```text
 如何控制精确分词的粒度，尽可能小
```
## 374.请教一下jieba的词典的格式是怎样的？
Dec 2, 2013
```text
 我尝试对恒大相关的文章进行分词，结果错误率奇高，结果如下
'''
按照原计划，凯赛尔在西班牙学习三年后就将回国，但对于志向高远的“凯撒”来说，他更希望凭借未来三年的表现能留在西班牙继续深造，从而拉开自己的职业生涯，“**/随恒大/足校在西班牙学习三年后，我希望能留在这里，并开启自己的职业生涯
保监会再发狠/招恒大/人寿被禁止投资股票
这些/对恒大/概念股的影响有多大
然后在10月底三季报披露散户看/到恒大/**买了4.95%是不是要举牌了，许老板就把货都卖给你们了
'''
我将自定义词典中”恒大“的词频调整到了10000都没有任何变化，尝试jieba.suggest_freq('恒大', True)也没有用；尝试将HMM关掉，结果恒大这个词会被一直分成”恒/大“。
```
```text
 能不能再切之前，直接添加：
jieba.add_word("恒大")
```
```text
 我想把只有在指定的字典有的关键词提出来。
import jieba
jieba.initialize()
jieba.set_dictionary('mydic.txt')
但是出来的词还是包含一堆不在字典中的。请问我该怎么做？谢谢
```
```text
 如果 jieba.analyse.extract_tags 能只返回指定字典里的就更好了
```
```text
 http://jiebademo.ap01.aws.af.cm/ 打不开，console里提示net::ERR_NAME_NOT_RESOLVED
```
```text
 想找一些词典做些训练用，不知道有人使用过没？
另外网上能找到类似的东西吗？例如名词，形容词词表一类的。
```
```text
 参看搜狗的2006年的输入法词库
```
## 375.关于CRF分词
Nov 21, 2013
```text
 wiki里“对Python中文分词模块结巴分词算法过程的理解和分析”跳转到黄网。
```
```text
 哈哈哈哈哈哈哈

2017-11-02 11:37 GMT+08:00 kercker <notifications@github.com>:
…
 wiki里“对Python中文分词模块结巴分词算法过程的理解和分析”跳转到黄网。

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#541>, or mute the thread
 <https://github.com/notifications/unsubscribe-auth/ABD5_fZrOpJSHG1DaFGVx-iT7X4t9A95ks5syTkJgaJpZM4QPIGV>
 .
```
```text
 6666666666老铁网站发出来呀
```
```text
 又骗我点击增加csdn浏览量。。。。。。。。。。。。。。。。。
```
```text
 我看到CRF分词算法的介绍，http://blog.csdn.net/ifengle/article/details/3849852
感觉还行。不知道jieba分词是怎么看CRF的？或者已经用了CRF?
```
```text
 从效果上看，CRF可以有更好的切分效果。百度已经在工程上应用了。
```
```text
 @chenweican 真的吗？感觉crf资料不是很多啊。
```
## 376.可以加入HMM模型tran的部分嘛？
Nov 20, 2013
```text
 如果我输入 “15路公交” 或“150平米以上的房子”
现在的切割是 15/公交 150/平米/以上/的/房子
我自定义词典里词频调成 0 或100都不行，谢谢
```
```text
 同问
能不能加入正则规则？
```
```text
 能否考虑，先切分再用某种方法合并？
```
```text
 嗯嗯...
```
```text
 您解决了么？
```
## 377.分词时能同时去掉停用词么？
Nov 16, 2013
```text
 “车、门、窗” 这类单字词无法抽取，这个可以加一个可选参数来来允许单字关键字呢？
```
```text
 结巴分词好像有停用字典，我在调用的时候会自动去除停用词么？如果不能，应该怎么调用？
```
```text
 我觉得更应该关心怎么使用结巴的停用词表
```
```text
 @DonMillion @chenweican 似乎已有类似的Issue了，请移步#77
```
## 378.求教搜索引擎模式下如何实现词性标注
Nov 13, 2013
```text
 举例：
sen = u"我钟爱法国"
我/r 钟爱/nr 法国/ns
sen = u"我鍾愛法國 "
我/r 鍾/nr 愛/v 法國/ns
有什么方法可以改善么？比如可以把简体字的词库改成繁体字词库么？
```
```text
 改用dict.big.txt的繁體字典後, 並參考我修正的錯誤 #670
```
```text
 import jieba as jb
import jieba.posseg as jp
if name == 'main':
r = "我来自中华人民共和国。"
d = jb.cut_for_search(r);
d1 = jp.cut(" ".join(d));
像这样写两次吗？
```
```text
 搜索引擎模式貌似把词拆碎了，再标注词性的用处不大了吧？
```
```text
 我也想问，主要是过滤标点……
```
## 379.UnicodeEncodeError,源码只decode?
Nov 1, 2013
```text
 比如要分词的字符串：http://192.168.0.121/login.view
因为：re_num = re.compile("[.0-9]+")这样写，
分词的结果为：
login eng
. m
view eng
"."被判断成数字，建议修改一下这个正则
```
```text
 这个“.”  很有必要被当做数字，信我。
```
```text
 查看源码,代码中会对文本进行decode,最终生成的分词结果并没有encode回来?
如果后续用其他工具对分词结果进行处理,会出UnicodeEncodeError,如果对分词结果手动encode就没问题.
猜想decode是为了把一个字当做一个字符处理,防止一个字的长度大于1.但是切词完之后不encode回来算是隐性的bug?
```
## 380.jieba分词 py2exe
Oct 23, 2013
```text
 每一次提取一句話的關鍵詞的速度平均都在6s左右如何優化這個過程
```
```text
 写了一个程序，使用py2exe制作可执行文件，但是制作完成后，点击执行文件，出现这样的错误，IOError: [Errno 2] No such file or directory: dist\library.zip\jieba\finalseg\prob_start.py'，而且手动添加这个文件，错误依旧，希望能够得到你的解答，谢谢。
```
## 381.结巴分词出现type 'exceptions.IOError'
Oct 16, 2013
```text
 无论是否翻墙，均无法访问，也ping不到
```
```text
 @Pzoom522 这个链接已经崩了很久了。
```
```text
 这个分词里面多少只有一种词性吗。
```
```text
 逻辑上 第一个多少应该分为 多 和 少
```
```text
 @sidealice , 这个难度挺高的。先mark一下。
```
```text
 这个好难哈哈
```
```text
 这种问题估计本身就有问题：比如我在你那句接一句。你说的多少是多少。只有说得人才知道哪个是多少，哪个是多+少了，甚至乎自己都不知道。这是语言表达本身就不够准确规范，让人都猜不透的，估计就不是难的问题了吧？
```
```text
 这个属于语意消除歧义，以现阶段的计算能力，是很难的。
```
```text
 <type 'exceptions.IOError'> [Errno 5] Input/output error   for term in terms:\n', '  File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 255, in cut\n    for word in cut_block(blk):\n', '  File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 195, in cut_DAG\n    DAG = get_DAG(sentence)\n', '  File "/usr/local/lib/python2.7/dist-packages/jieba/__init.py", line 115, in wrapped\n    initialize(DICTIONARY)\n', '  File "/usr/local/lib/python2.7/dist-packages/jieba/init.py", line 65, in initialize\n    print >> sys.stderr, "Building Trie..., from " + abs_path\n'
难道是由于后台运行后，console退出了，导致print出差？
```
## 382.posseg中viterbi需要在终止时刻优先选取为ES的状态吗？
Oct 10, 2013
```text
 hi,
我在jieba.dict.utf8这个字典里面等26556行发现下面的内容
今天上午 3 nr
今天下午 3 nr
nr不是人名的标注吗？ 为什么这俩的标注是nr？
```
```text
 终止时刻的可能状态列表如果为E或者S表示它是一个更合适的候选者，除非没有E和S状态，才考虑B和M，看了下char_state_tab，没有e,s状态的字很少。
另外jieba.posseg.viterbi.py19行应该注释掉吗?
```
```text
 @lurga , 对于第一个问题，需要多测试才能给出准确答复。
对于第二个问题，非常感谢你，19行本来是我对viterbi算法做的一个剪枝策略，应该是忘了注释。
剪枝策略应用后，分词速度会更快，准确度会下降多少我也没来得及测试。
你有空的话，可以帮忙看看19行的策略和20行的策略，综合起来，哪个更好。谢谢。
```
```text
 @fxsjy 我不太懂python，用scala改了个版本，由于概率文件是一样的，结果应该一样，以下测试都直接使用viterbi来分词，没有用词典。
1、第一个问题我做了下面的测试，优先选择E,S状态的效果要更好一点，当然这只是一个个例，没办法证明它的效果始终更优。
测试语句是"张三说的确实在理"
终止状态优先选择E,S：('B', 'nrfg'), ('M', 'nrfg'), ('E', 'nrfg'), ('S', 'uj'), ('B', 'ad'), ('E', 'ad'), ('B', 'v'), ('E', 'v')
不判断终止状态：('B', 'nrfg'), ('M', 'nrfg'), ('E', 'nrfg'), ('S', 'uj'), ('B', 'ad'), ('E', 'ad'), ('S', 'p'), ('B', 'n')
2、第二个问题我做了下面的测试，剪枝策略对结果影响很大。
测试语句是 “中国商务部”
top4：('M', 'nr'), ('M', 'nr'), ('M', 'nr'), ('M', 'nr'), ('M', 'nr')
top10：('E', 'nrfg'), ('S', 'n'), ('B', 'j'), ('M', 'j'), ('M', 'j')
top20：('E', 'nrfg'), ('S', 'n'), ('B', 'j'), ('M', 'j'), ('M', 'j')
top30：('M', 'nrfg'), ('E', 'nrfg'), ('B', 'j'), ('M', 'j'), ('M', 'j')
top40：('B', 's'), ('E', 's'), ('B', 'nt'), ('M', 'nt'), ('E', 'nt')
top50：('B', 'n'), ('E', 'n'), ('B', 'nt'), ('M', 'nt'), ('E', 'nt')
top60：('B', 'nt'), ('M', 'nt'), ('M', 'nt'), ('M', 'nt'), ('E', 'nt')
大致取到前60个状态才跟不剪枝的结果保持一致。对t-1时刻的状态做出限定，极大可能导致最后结果出现严重误差，数学如何解释我也不清楚，盼答复。
```
## 383.中文名支持
Oct 8, 2013
```text
 “我很帅”这句话词性标注时，“帅”这个词标注为了形容词a，但是在“我非常帅”中“帅”这个词标注为了nr。词典中的“帅”词性为nr。这是什么原因啊，我刚开始使用词性标注工具，原理还不太懂。麻烦懂的解答一下。
```
```text
 Jieba对中文名的支持不是很好啊,例如：
中国商务部国际贸易谈判副代表俞建华表示，2001年在上海，中方成功举办了APEC第九次领导人非正式会议。时隔13年，中方非常荣幸再次成为APEC东道主。中国政府高度重视这一盛事，目前各项准备工作已全面展开。
分词为
中国/ns 商务部/nt 国际/n 贸易谈判/n 副/b 代表/n 俞/zg 建华/nz 表示/v ，/x 2001/m 年/m 在/p 上海/ns ，/x 中方/f 成功/a 举办/v 了/ul APEC/eng 第九次/m 领导人/n 非正式/b 会议/n 。/x 时隔/n 13/m 年/m ，/x 中方/f 非常/d 荣幸/nr 再次/d 成为/v APEC/eng 东道主/nr 。/x 中国政府/nt 高度重视/l 这/r 一/m 盛事/n ，/x 目前/t 各项/r 准备/v 工作/vn 已/d 全面/n 展开/v 。/x
其中人名“俞建华”被分成两个词了"俞/zg 建华/nz ”,有计划改进吗？
```
```text
 如果只用posseg中的viterbi方法进行分词，“中国商务部国际贸易谈判副代表俞建华表示” 被分得非常完美
中国商务部/nt, 国际/n, 贸易/vn, 谈判/vn, 副/b, 代表/n, 俞建华/nr, 表示/v
```
## 384.GAE的支持
Oct 2, 2013
```text
 如果我想自己產生詞典, 我的資料只有
结巴中文分词
這句話
那我的詞典應該是這樣嗎
结 1
巴 1
中 1
文 1
分 1
词 1
结巴 1
巴中 1
中文 1
文分 1
分词 1
结巴中 1
.
.
.
```
```text
 请问jieba可放上GAE吗？谢谢。
```
```text
 不是什么问题，主要是把第一次生成的temp文件改一改。
On Wed, Oct 2, 2013 at 11:14 PM, vaemon notifications@github.com wrote:

请问jieba可放上GAE吗？谢谢。
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/109
.
```
```text
 I think what this issue is saying is "Can Jieba be used on Google App Engine?" and jannson's response is "well, creating the temp files doesn't work".
That's also what I'm finding:
Traceback (most recent call last):
  File "/base/data/home/apps/s~dokibo-dict/20161021t175549.396505411313804719/lib/jieba/__init__.py", line 123, in initialize
    self.tmp_dir or tempfile.gettempdir(), cache_file)
  File "/base/data/home/runtimes/python27/python27_dist/lib/python2.7/tempfile.py", line 45, in PlaceHolder
    raise NotImplementedError("Only tempfile.TemporaryFile is available for use")
NotImplementedError: Only tempfile.TemporaryFile is available for use

What approach could be used to fix this?
```
```text
 @astromme
Maybe you can make a patch to gae....tempfile.py
like this :  https://github.com/liantian-cn/jieba-gae/blob/master/lib/tempfile_gae.py
Then just replace import tempfile for import tempfile_gae as tempfile , in the error of the file.
like this : https://github.com/liantian-cn/jieba-gae/blob/master/lib/jieba/__init__.py
It looks good now : https://jieba.liantian.me/
note:
If you replace tempfile.py directly, do not use tempfile_gae.py, the development service is not available.
```
## 385.获取词性时，会缺失分词准确度
Sep 29, 2013
```text
 Traceback (most recent call last):
File "", line 1, in 
File "/opt/anaconda3/lib/python3.6/site-packages/jieba/init.py", line 119, in wrapped
return fn(*args, **kwargs)
File "/opt/anaconda3/lib/python3.6/site-packages/jieba/init.py", line 307, in load_userdict
word, freq = tup[0], tup[1]
IndexError: list index out of range
```
```text
 如图，结巴加载后不会往下执行，而是停在最后一行。
(py3env1) ➜  nlpTest python -m jieba -d cutTest.txt > cuted2.txt
Building prefix dict from the default dictionary ...
Loading model from cache /var/folders/fv/9t0ldhvx03j4ch64rv3j9qh40000gn/T/jieba.cache
Loading model cost 0.855 seconds.
Prefix dict has been built succesfully.
```
```text
 当去获取词性时，可能有些词没有词性，故会损失准确度
比如“韵达” 如果用分词，就会分出 韵达
如果用获取词性，会得到 韵 m 达  v 没有出现 韵达的词性
我觉得这样并不好，如果没有词性，可否返回一个空或者其它呢，也即尽量不要破坏分词的准度
```
```text
 使用词性标注，分词结果没有精确分词结果准确，比如在精确模式下可以识别出干锅，而在词性标注模式下变成了干和锅两个词，请问如何修正这个问题？
```
## 386.怎么用结巴分词提取关键词和各自的频数啊？
Sep 5, 2013
```text
 “车、门、窗” 这类单字词无法抽取，这个可以加一个可选参数来来允许单字关键字呢？
```
```text
 结巴分词好像有停用字典，我在调用的时候会自动去除停用词么？如果不能，应该怎么调用？
```
```text
 我觉得更应该关心怎么使用结巴的停用词表
```
```text
 @DonMillion @chenweican 似乎已有类似的Issue了，请移步#77
```
```text
 现在只能提取关键词 仍然不能提取频数.....
```
```text
 @fxsjy 楼主可以开放统计频数不？这个接口很重要啊!
```
## 387.请问，为什么不用C/C++写个呢？结巴的分词效果挺好的。
Sep 4, 2013
```text
 比如要分词的字符串：http://192.168.0.121/login.view
因为：re_num = re.compile("[.0-9]+")这样写，
分词的结果为：
login eng
. m
view eng
"."被判断成数字，建议修改一下这个正则
```
```text
 这个“.”  很有必要被当做数字，信我。
```
```text
 查看源码,代码中会对文本进行decode,最终生成的分词结果并没有encode回来?
如果后续用其他工具对分词结果进行处理,会出UnicodeEncodeError,如果对分词结果手动encode就没问题.
猜想decode是为了把一个字当做一个字符处理,防止一个字的长度大于1.但是切词完之后不encode回来算是隐性的bug?
```
```text
 No description provided.
```
```text
 @54xiaosu 您好。
我之前仿照楼主大神的算法写了一个，https://github.com/aszxqw/cppjieba
经线上测试，挺稳定。效果蛮好。不知道是否对您有所帮助 :)
```
```text
 @aszxqw ，线上使用了啊，方便透漏是哪家网站吗？
```
```text
 @fxsjy 已私聊。
```
## 388.有词库意义的库吗？类似wordnet的功能？除了词性，还有近义词，和上下位？
Aug 28, 2013
```text
 No description provided.
```
```text
 你没看代码吗？
```
```text
 哎，作者就是不爱写注释，哎。
```
```text
 @ilqxejraha 谢谢您，可能是我表达有误，我是想问词汇库里面的词汇和词频是人工统计的吗？还是通过其它的方法。
```
```text
 统计过来的。最后得到了就是这么一个模型。
```
```text
 你在源码中看到词频的使用了吗？
```
```text
 具体的统计词频，词频的作用可能体现在，一个词存在多个意思。
比如英语中，经常有一个词会有很多个意思。
比如出现一个词，并且，这个词有很多种解释，这时候词频可能会对词意的选择有一定帮助。
具体的其他算法hmm的我还没看。
```
```text
 @KevinDotW 据说是基于人民日报的语料库，我也想知道怎么才能创建自己的词典
```
```text
 词性标注是否可以指定分词模式；
想指定搜索模式
谢谢。
```
```text
 比如微博的表情文字是   [泪]
我想拿到这个表情统计个数，请问可以怎么做？
```
```text
 想要有个词语与词语的关系，用来做分析？或者可以一同开发？请指点下，或我留个邮箱？
```
```text
 词语与词语间关系 是基于词性还是完全不考虑词性呢
```
```text
 基于的，比如钱包，近义词是钱夹，属于：皮革制品，纺织品，配件。。。
```
```text
 @c2h2 几乎不可能的，即使有准确率也不高，因为这问题我觉得还是主要靠人工review维护一个同义词词典。虽然这个应用的需求很强烈。
```
```text
 @c2h2 ， 试一试Google的word2vec ?
```
## 389.TOPK排序后结果没有调用自定义词典的分词结果？
Aug 14, 2013
```text
 ['a', ' ', 'b', ' ', 'c']
出来空格。。
```
```text
 基于 TextRank 算法的关键词抽取
jieba.analyse.textrank(sentence, topK=20, withWeight=False, allowPOS=('ns', 'n', 'vn', 'v')) 直接使用，接口相同，注意默认过滤词性。
jieba.analyse.TextRank() 新建自定义 TextRank 实例
算法论文： TextRank: Bringing Order into Texts
基本思想:
将待抽取关键词的文本进行分词
以固定窗口大小(默认为5，通过span属性调整)，词之间的共现关系，构建图
计算图中节点的PageRank，注意是无向带权图
我在使用textrank的时候想调整滑动窗口为2,上面说窗口大小默认为5,但找不到span在哪里,求帮助!
```
```text
 搞掂了,原来要新建一个对象,然后改属性,例如:
text = jieba.analyse.TextRank()
text.span = 2
```
```text
 我有一批关键词，处理步骤如下：
1、先将关键词根据自定义词典分词
2、将分词后的结果topk排序
3、把排序后的结果排除一些无意义的词
现在的问题是发现topk排序，调用的分词好像没有调用自定义词典，也不知道是不是我搞错了，求作者出来解释下topk是怎么处理的，谢谢～～～
疑问：
1、怎么样可以topk排序调用分词后的结果呢
2、topk排序计算的tf-idf值可以调用出来么？计算的文档数量是多少？
```
```text
 @lbw1215 ，没太看明白。贴点代码吧
```
```text
 @fxsjy
雅思必备词汇有多少     雅思|必|备|词汇|有|多少   雅思|必备|词汇|有|多少  雅思|必备|词汇
一共4列，第一列是原词，第二列是分词后的结果，第三列是对分词后的结果进行topk排序，第四列是对排序后的结果排除某些无意义的词
现在的问题是发现第三列处理的数据，好像不能调用第二列处理后的数据，不知道topk的调用机制是怎么样呢？
import sys
reload(sys)
sys.setdefaultencoding('utf-8')
sys.path.append('../')
import jieba
jieba.initialize()
jieba.load_userdict("userdict.txt")   #自定义词典
import jieba.posseg as pseg
import jieba.analyse
import re
f = open('output.txt','wb')
f_ex = open('words.txt','rb')   #排除词，一行一个单词
words = [line.strip() for line in f_ex.readlines()]
for eachline in open('input.txt','rb'):
i = 0
line = eachline.strip().decode('gbk')      #str
f.write(line.encode('utf-8')+'\t')


cut_result = list(pseg.cut(line))                 #分词
for w in cut_result:
    f.write(w.word.encode('utf-8')+'|')
f.seek(-1,1)
f.write('\t')

tags = jieba.analyse.extract_tags(line, topK=20)    #TOPK排序
for w in tags:
    f.write(w.encode('utf-8')+'|')
f.seek(-1,1)
f.write('\t')

for cut_w in tags:
    for w in words:                           #去除指定关键词
        #cut_w.word = re.sub(w.decode('gbk'),'',cut_w.word)
        if w.decode('gbk') == cut_w :
            cut_w = ''
    if cut_w != '':
        f.write(cut_w.encode('utf-8')+'|')
f.seek(-1,1)
f.write('\t')
f.write('\r\n')

f.close()
f_ex.close()
```
## 390.出现 Cannot allocate memory  错误
Aug 14, 2013
```text
 First of all thank you for your work, you did so far!
I have a question for this specific line https://github.com/fxsjy/jieba/blob/cb0de2973b2fafaa67a0245a14206d8be70db515/jieba/posseg/init.py#L17. Why do you use this specific range of unicode literals? For my specific case(russian text) your app not splitting words which is not good at all.
For example:
>>> re_han_internal = re.compile("([\u4E00-\u9FD5a-zA-Z0-9+#&\._]+)")
>>> re_han_internal.split("""Простой и безопасный способ делиться терминалом: обзор инструмента warp""")
['Простой и безопасный способ делиться терминалом: обзор инструмента ', 'warp', '']

Increasing unicode range helps:
>>> re_han_internal = re.compile("([\u0041-\u9FD5a-zA-Z0-9+#&\._]+)")
>>> re_han_internal.split("""Простой и безопасный способ делиться терминалом: обзор инструмента warp""")
['', 'Простой', ' ', 'и', ' ', 'безопасный', ' ', 'способ', ' ', 'делиться', ' ', 'терминалом', ': ', 'обзор', ' ', 'инструмента', ' ', 'warp', '']
```
```text
 it's based on the dict which is ONLY for chinese by default, u wanna use the program, u should add a special russian dict.
```
```text
 这个简单的测试程序，会产生exception：
import jieba
s = "女装"
a = jieba.cut(s, cut_all=False)
print "/".join(a)
错误是：
UnicodeEncodeError: 'gbk' codec can't encode character u'\u016e' in position 0: illegal multibyte sequence
我debug发现原因是 "女装" 用utf-8解码时并不会产生exception，但其实解码后的东西是乱七八糟的：u'\u016e\u05f0' 后面用这个字符串去分词了。我不知道为什么这个词这么特殊。
一个解决办法是提供一个参数，人工指定编码，这样就肯定不会错了。
```
```text
 在mac os下用Python 2.7.9测试没有出错，返回结果还是“女装”这个词。
```
```text
 @alexwwang 我的.py是GBK编码的。你的脚本头加上# -- coding: gb18030 -- 试试？
```
```text
 s 是 str 类型，gb18030 编码（根据源代码的编码而定）；你的命令行环境是 'gbk' 编码的。
因此，在分词时，由于 s 是 str 类型，程序先进行尝试 utf-8 解码（解码错误再尝试 gbk），结果！解码成功，为Ůװ两个字符，尝试 print 的分词结果为 Ů/װ。但由于 gbk 无法编码这两个字符（码表里没有），只好异常。
因此，要避免出这种状况：

s = u"女装"
尽量在源代码中使用 UTF-8 编码
```
```text
 @gumblex 谢谢。我上面的代码只是PoC，实际上是从一个GBK文件逐句读出然后分词，产生的异常。所以不能用您建议的办法来避免了。目前我的做法是先decode再传给jieba. 不过因为jieba号称不管编码是GBK或者unicode都可以自动支持，所以我觉得这个还应该是个bug。
```
```text
 做文本处理的该用 Unicode 类型就用 Unicode 类型，毕竟分词针对的是字符，不是用来负责判断这到底是 GBK 还是 Unicode 编码。
```
```text
 做个应用，抓取网页里中文，然后统计出现数量较多的关键词。content是网页的html代码。运行会出现Cannot allocate memory，free-m 看内存被使用290多m（vps是512内存，基本上就剩290多m），请问如何优化下代码使其可以正常运行？match结果大概1200多个词，每个词不超过5个字。以下是代码
count = {}
match=re.findall(r'[\x80-\xff]+',content)
ss=','.join(match)
seg_list = jieba.cut_for_search(ss)
for c in seg_list:
if c.encode('utf-8') in count:
count[c.encode('utf-8')] = count[c.encode('utf-8')] + 1
else:
count[c.encode('utf-8')]=1
return sorted(count.iteritems(),key=lambda x:(x[1],x[0]),reverse=True)[0:10]
```
```text
 @pydotnet , 用一个小一点的字典试试：https://github.com/fxsjy/jieba/blob/master/extra_dict/dict.txt.small
替换掉dict.txt
```
```text
 @pydotnet ， 可参考： #17
```
## 391.jieba0.31与whoosh2.5.1集成过程中，出现了KeyError:"Can't store a null key ()"
Aug 9, 2013
```text
 No description provided.
```
```text
 英文的你要中文分词怎么做，加入词典看看
```
```text
 同样的问题 可以用同义词 但是试了不好用
```
```text
 这里可以将Python ， C++分出来。
```
```text
 建议提供 debug 选项控制类似下面这样的输出
print >> sys.stderr, f_name, ' at line', lineno, line
```
```text
 @shuge ， 什么样的控制？是控制让这些日志不显示，还是说输出到其他地方？
```
```text
 def foo(..., debug=False, log_file=sys.stderr):
    if debug:
        print >> log_file, f_name, ' at line', lineno, line
    ...

您觉得这样会不会更好些？
```
```text
 @shuge , 好的，谢谢你的建议。你要是能给我发pull request就更好了:-)
```
```text
 稍后吧，我上周临时用 pymmseg + AdvancdLangConv 解决了
jieba 是一个很有意思的项目
```
```text
 这个问题是出现在“全模式”下（cut_all=True），而在“精确模式”下则没有上面的问题。
我在全模式下输出分词结果，发现list列表前后多了两个空值。
```
```text
 Fixed. please check:   90ab511
```
```text
 @jaffwu , 另外，jieba已经内置了和whoosh的集成功能。 请看：https://github.com/fxsjy/jieba/blob/master/test/test_whoosh.py
```
## 392.def get_DAG(sentence) method question
Aug 9, 2013
```text
 【词频省略时使用自动计算的能保证分出该词的词频】这句话不太明白。
词频省略的时候，要如何自动计算词频呢（详细），谢谢~
```
```text
 Hi  抱歉咨询一个简单的问题， 自定义词典更新后，缓存文件jieba.cache是否会更新?
```
```text
 我发现没有更新，好郁闷！不知道如何更新。
```
```text
 对--"我来到北京清华大学" 分词，
用 get_DAG(sentence)方法，返回的DAG为：
dict: {0: [0], 1: [1, 2], 2: [2], 3: [3, 4], 4: [4], 5: [5, 6, 8], 6: [6, 7], 7: [7, 8], 8: [8]}
而5为清：在dict.txt中做trie树后，查找结果应该为 清华 ，清大 为---5：[5,6,7]
个人认为方法那个迭代有问题，谢谢！
```
```text
 @gucasbrg, 有什么问题？是闭区间。
5: [5, 6, 8]
55：清
56：清华
5~8：清华大学
```
## 393.关于算法
Aug 7, 2013
```text
 No description provided.
```
```text
 No description provided.
```
```text
 版本 <0.37 中 pair 对象不是 iterable，即不能像 tuple 一样直接拆开来。之前的测试代码相关部分是：
for w in words:
    print('%s %s' % (w.word, w.flag))
```
```text
 恩，已经解决，谢谢！！前后例子不太一致，导入词典的那个例子里用的是你上面写的例子。。。主页上用的是tuple
```
```text
 导入词典的例子在哪里
```
```text
 Hi，你有算法详细一点的介绍文档么？github里的algorithm 写得太简单了。
```
```text
 久远的帖子了，今天闲着没事，翻了过去的记录，给你发一些链接：
http://book.51cto.com/art/201106/269050.htm ——概率语言模型的分词方法原理介绍（HMM模型的好像也有介绍）
http://www.matrix67.com/blog/archives/4212 ——漫话中文自动分词和语义识别（上）：中文分词算法
http://www.isnowfy.com/python-chinese-segmentation/ ——最简单的概率分词实现
http://www.oschina.net/code/snippet_1180874_24398 ——最简单的概率分词实现（c++版）
求介绍文档没太多意思，重要的还是看代码，上面的链接我也并没有完全看完:)
```
## 394.作者能否分享语料分析的源码？
Aug 7, 2013
```text
 把痘痘加入自定词典后，也无法把长痘痘中的长和痘痘分开。各位大佬有什么办法嘛？
```
```text
 jieba.possseg.cut的分词不能支持和jieba.cut一样的可选cut_all参数，这样两种模式下分词结果不一样。
```
```text
 繁体文本分词的时候用的是官方提供的dict.txt.big。可以看到里面“清洁机”和“清潔機”的词频、词性都是相同的。但是切分同一个、仅是简繁体不同的句子时：
“这台清洁机引出一条胶管。”
“這台清潔機引出一條膠管。”
“清洁机”和“清潔機”得出的词性却不同，而且“清潔機”的词性是“x”，即“非语素字”。不止一个词有这样的情况。
```
```text
 我也遇到這個問題，而且是默認字典裡面的詞才會有這樣的問題
如 “攝影機” 在dict.txt.big是n，但是切出來的詞性卻是x。
```
```text
 也就是分析和统计语料生成词典文件、idf.txt文件的源代码，非常期待~~
```
```text
 @lewsn2008 你看看作者其他的项目，已经分享了
```
```text
 @piaolingxue thanks~
```
```text
 @piaolingxue 是按个项目啊？能给个项目名不？
```
## 395.能加上分出时间和数字序列吗
Aug 4, 2013
```text
 如果我想自己產生詞典, 我的資料只有
结巴中文分词
這句話
那我的詞典應該是這樣嗎
结 1
巴 1
中 1
文 1
分 1
词 1
结巴 1
巴中 1
中文 1
文分 1
分词 1
结巴中 1
.
.
.
```
```text
 请问jieba可放上GAE吗？谢谢。
```
```text
 不是什么问题，主要是把第一次生成的temp文件改一改。
On Wed, Oct 2, 2013 at 11:14 PM, vaemon notifications@github.com wrote:

请问jieba可放上GAE吗？谢谢。
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/109
.
```
```text
 I think what this issue is saying is "Can Jieba be used on Google App Engine?" and jannson's response is "well, creating the temp files doesn't work".
That's also what I'm finding:
Traceback (most recent call last):
  File "/base/data/home/apps/s~dokibo-dict/20161021t175549.396505411313804719/lib/jieba/__init__.py", line 123, in initialize
    self.tmp_dir or tempfile.gettempdir(), cache_file)
  File "/base/data/home/runtimes/python27/python27_dist/lib/python2.7/tempfile.py", line 45, in PlaceHolder
    raise NotImplementedError("Only tempfile.TemporaryFile is available for use")
NotImplementedError: Only tempfile.TemporaryFile is available for use

What approach could be used to fix this?
```
```text
 @astromme
Maybe you can make a patch to gae....tempfile.py
like this :  https://github.com/liantian-cn/jieba-gae/blob/master/lib/tempfile_gae.py
Then just replace import tempfile for import tempfile_gae as tempfile , in the error of the file.
like this : https://github.com/liantian-cn/jieba-gae/blob/master/lib/jieba/__init__.py
It looks good now : https://jieba.liantian.me/
note:
If you replace tempfile.py directly, do not use tempfile_gae.py, the development service is not available.
```
```text
 像7月5日或四四五二或两万三千零二十这样的要怎么分？
```
```text
 @sunsol ， 现在还不行啊，有什么思路吗？
```
```text
 日期貌似可以蛮力法来解决。
```
```text
 词库支持表达式觉得能 就爽了
```
## 396.请问大家都是如何利用这个工具的呢？
Jul 30, 2013
```text
 def setLogLevel(log_level):
global logger
default_logger.setLevel(log_level)
logger 是否应该是 default_logger?
```
```text
 jieba.__version__ = 0.34
字典檔確定有「台中」這個詞。
不論用何種模式，「台中」都會切成「台」「中」，但是正確應該不會被切開。
「台北」、「台南」就不會被切開，這樣是正確的。


請問這有問題有辦法解決嗎？？
```
```text
 可以把台中加入自定义词典jieba.load_userdict
```
```text
 好，我試試看，但是我好奇的是，「台中」這個字詞已經在字典裡面了，為何還是斷詞還是斷成「台」「中」？？
```
```text
 >>> jieba.FREQ['台中']
3
>>> jieba.FREQ['台']
16964
>>> jieba.FREQ['中']
243191
>>> jieba.total
60101967
3/60101967 < 16964×243191/601019672
即，P(台中) < P(台)×P(中)，台中词频不够导致概率较低
在自定义词典调高词频即可，或者add_word用更高频率覆盖“台中”
（这里是>16964×243191/60101967，即≥69）
```
```text
 我发现jieba会将“不合理”看做是一个词，而不是“不/合理”。
另外，基于@gumblex 给出的计算，应该是可以将“不”和“合理”分开的，如下：
jieba.get_FREQ(u'不合理') # 470
jieba.get_FREQ(u'不') # 360331
jieba.get_FREQ(u'合理') # 3870
词典中共有584429个词（dict.txt.big），计算：
470/584429 < 360331*3870/584429^2
如果使用user dict，当把“合理”的词频调高后，可以分开。
```
```text
 这个计算的分母是总频数(jieba.dt.total = 60101967)不是总词数，这样计算下来「不合理」确实高于其他分词。
你可以手动拆分：jieba.suggest_freq(('不 ','合理'), True)，或直接删除 del_word。
```
```text
 @gumblex ：
明白了。另外一个问题是：
1）在当前程序中调用suggest_freq，会影响后续的使用么？
2）del_word只是对当前的分词起作用么？是否并没有将其在dict.txt中删除？
```
```text
 @gumblex :
厚着脸皮再问一个，POS标注中的分词结果是否和jieba.cut相同？
```
```text
 suggest_freq 和 del_word 等仅对当前分词器实例 (jieba.Tokenizer) 起作用（旧版是仅对全局词典起作用），不会修改本地词典，相当于改了几个变量。
POS标注中的分词结果不一定相同，见 #212。
```
```text
 怎么让 POS标注 支持 suggest_freq 和 del_word ?
```
```text
 我在思考如何将爬虫抓取的文本结合jieba工具进行文本分析。
在“终端”中一段一段地进行分词不大方便，不知道有没有同学做过这个整合？如果没有，请告诉我大概需要些什么，我想试着做一下（我目前只会matlab，哈哈都不好意思说了）
谢谢～
```
```text
 我用jieba来分析用户的评论，然后提取关键评论短语。
```
```text
 @linkerlin yes~当有大量的用户评论需要统一处理的时候，要怎么做呢？
```
```text
 @zihaolucky 另写一个脚本读取评论，调用jieba分词，作分析，输出结果，用matlab不知道怎么搞，用python很好搞。
```
```text
 我也是用python写个脚本，分析微博内容数据，获取高频词汇
在 2013年7月31日上午9:33，maxint64 notifications@github.com写道：

@zihaolucky https://github.com/zihaolucky另写一个脚本读取评论，调用jieba分词，作分析，输出结果，用matlab不知道怎么搞，用python很好搞。
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/86#issuecomment-21835061
.



Regards
Bojan Liu
My Blog: http://newliu.com
```
```text
 @maxint64 嗯～我是抓取知乎上的内容进行文本分析。抓取工具用的是GoSeeker，不知道两个要怎么配合起来才好。因为到了后面，如果抓取内容多的话，手工一个个来分词就不行了。
maxint64 是不是把文本文件读进来，然后返回结果也形成文本文件？我是想这样做，可以指点一下吗？
```
```text
 @zihaolucky , 这个例子就是从一个文件读入文本，然后分词后输出到另外一个文件。 https://github.com/fxsjy/jieba/blob/master/test/test_file.py
```
```text
 GoSeeker 是什么？
没有Google到。。。
```
```text
 文本分析需要分析哪些内容？
爬虫是肯定需要的了，推荐你一个python写的爬虫:scrapy
2013/7/30 Zihao Zheng notifications@github.com

我在思考如何将爬虫抓取的文本结合jieba工具进行文本分析。
在“终端”中一段一段地进行分词不大方便，不知道有没有同学做过这个整合？如果没有，请告诉我大概需要些什么，我想试着做一下（我目前只会matlab，哈哈都不好意思说了）
谢谢～
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/86
.
```
```text
 @zihaolucky
没用过GoSeeker，不过基本就是你说的那样做。
也不一定非要是文本文件，什么格式方便你做分析就用什么格式。
```
```text
 @linkerlin 嗷、是GooSeeker，一个Firefox的插件。
```
```text
 @jannson thanks~似乎是很强大的一个工具。我会静下心来好好看代码的了，老是抗拒也不是个问题啊哈。
```
```text
 @fxsjy 收到！学习一下，谢谢～
```
```text
 我是对nlp刚研究的新手  想先从一段文字中进行分词 从而提取他的教育经历和工作经历.
目前看来有些效果挺好 有些还一般 ,  应该是组织机构的库还不好.
在 Jul 31, 2013，9:52，BojanLiu notifications@github.com 写道：

我也是用python写个脚本，分析微博内容数据，获取高频词汇
在 2013年7月31日上午9:33，maxint64 notifications@github.com写道：

@zihaolucky https://github.com/zihaolucky另写一个脚本读取评论，调用jieba分词，作分析，输出结果，用matlab不知道怎么搞，用python很好搞。
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/86#issuecomment-21835061
.



Regards
Bojan Liu
My Blog: http://newliu.com
—
Reply to this email directly or view it on GitHub.
```
```text
 看来这个工具方便了很多人呐～
```
```text
 @fxsjy 我成功啦～
不过似乎有点问题。我把这条新闻存到txt文件中

分词结果是这样的：

当我再一次 run test_file.py a.txt 出现这样：

主要是最后一段的问题
```
```text
 @zihaolucky , 你的问题解决了么？ 我没太看懂截图要表示的意思。
```
```text
 @fxsjy 抱歉我没有表达清楚。
我有这样一则新闻：

我打开'终端“，第一次run test_file.py a.txt的时候得到了log文件的结果是这样的：

随后紧接着再执行一次，才得到了正确的结果：

不知道这是什么问题呢？我必须得执行run test_file.py a.txt两次才能完整地分词。
这个是我的执行记录：
```
```text
 @zihaolucky ,  可能是文件没有close导致的。你看看利用的test_file.py和这一个一样么？https://github.com/fxsjy/jieba/blob/master/test/test_file.py
```
```text
 import urllib2
import sys,time
import sys
sys.path.append("../")
import jieba
jieba.initialize()

url = sys.argv[1]
content = open(url,"rb").read()
t1 = time.time()
words = list(jieba.cut(content))

t2 = time.time()
tm_cost = t2-t1

log_f = open("1.log","wb")
for w in words:
    print >> log_f, w.encode("utf-8"), "/" ,
print 'cost',tm_cost
print 'speed' , len(content)/tm_cost, " bytes/second"

我的错..我是不是把这个tets_file.py的文件替换成现在这个就可以了？
嗯，确认可以了～谢谢你啊
```
## 397.关于idf.txt
Jul 30, 2013
```text
 你好，在做基于情感词典的情绪分析，结果发现，结巴分词，会把很多词合在一起。比如：你真笨；
我想要的结果是：你，真，笨，
但是返回的却是：你，真笨
还有：你太笨，返回的也是：你，太笨
感觉很不科学呐。真和太明明是一种程度词，怎么和笨合成一个词了。
同时调研了其它多个分词开源工具，发现只有你们会合在一起。
```
```text
 是不是只有自定义的词比你们已有的词长才会有效，而比你们的词短的就不生效？
```
```text
 你可以试试lcut(seq,HMM=false)

Brent

On Aug 7, 2017, at 18:55, jiangchao123 <notifications@github.com<mailto:notifications@github.com>> wrote:


是不是只有自定义的词比你们已有的词长才会有效，而比你们的词短的就不生效？

—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub<#503 (comment)>, or mute the thread<https://github.com/notifications/unsubscribe-auth/AM-fgZNs2YqKUicyDBIwL5pGAZrIcKnbks5sVu07gaJpZM4OvHe9>.
```
```text
 试了一下，这样在有些词是可以的，有些依然不行，并且如果HMM设为FALSE的话，好多人名都会被切开了，比如：小明被切成了小，明；李小福被切成了李，小，福；
找代驾我想分成找，代驾；但是HMM设为True为找代驾；HMM设为False为找，代，驾
看来是不能两全？？？
```
```text
 想请教tfidf部分是如何进行分词的？能自定义分词字典么，自定义删除一些词汇
tags = jieba.analyse.extract_tags(content, topK=topK, withWeight=withWeight)
```
```text
 呃 后两个问题文档中就有解答
```
```text
 我有一段文本，其中有很多英文以及标点符号，希望只对其中的中文进行分词处理，除了用 stop_words ,还有什么好的办法吗？
```
```text
 先用正则把非中文的字符去除
```
```text
 可以试试和这个类似的正则表达式 http://git.io/piM6rQ
```
```text
 貌似代码里面没有从语料生成idf.txt的代码，是否遗漏？
```
```text
 词典和idf.txt都是坐着事先对语料进行训练和分析得到的，不包含在这个工程里面。
不过我也很想看到语料分析的代码，期待作者共享！
```
```text
 如果大家对这个有兴趣，我愿意分享给大家。
当时觉得idf.txt里面的统计不大好，所以自己想办法生成了一份。还有词料统计，用了最大熵的思路，生成了一分语料统计。主要是为了发现新词的。
2013/8/7 lewsn2008 notifications@github.com

词典和idf.txt都是坐着事先对语料进行训练和分析得到的，不包含在这个工程里面。
不过我也很想看到语料分析的代码，期待作者共享！
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/87#issuecomment-22226557
.
```
```text
 @linkerlin , @lewsn2008 , 我找到了之前写的生成idf.txt的脚本，基本思路是对一些小说报刊语料进行分词，然后以段落为单位，统计idf.
import jieba
import math
import sys
import re
re_han = re.compile(ur"([\u4E00-\u9FA5]+)")

d={}
total = 0
for line in open("yuliao_onlyseg.txt",'rb'):
    sentence = line.decode('utf-8').strip()
    words = set(jieba.cut(sentence))
    for w in words:
        if w in jieba.FREQ:
            d[w]=d.get(w,0.0) + 1.0
    total+=1
    if total%10000==0:
        print >>sys.stderr,'sentence count', total

new_d = [(k,math.log(v/total)*-1 ) for k,v in d.iteritems()]

for k,v in new_d:
    print k.encode('utf-8'),v
```
```text
 @fxsjy 非常感谢，作者真是无私的大牛啊，膜拜！
```
```text
 谢谢！
没看明白为何要 math.log(v/total)*-1
```
```text
 @linkerlin , 也可以math.log(total/v)
```
```text
 求语料数据，程序中的yuliao_onlyseg.txt还有吗？想跑一下学习学习，thks
```
```text
 一个新的分词库：https://github.com/jannson/yaha ，仅与大家交流学习：

提供了解决结巴分词库的姓名识别，后缀名识别，使用正则表达式等问题的思路；
同时对提取关键字，ChineseAnalyzer进行了小小的优化；
附加了 最大熵算法生成新词，自动摘要，比较两文本的相似度算法的实现。

产生这个分词库的原因，是因为在我的一个小小的爬虫，搜索引擎上使用结巴分词库之后，发现了一些小问题优化之后形成的，本来想直接修改结巴代码并提交，但是有一些设计思路区别较大才弄的新的分词库，不是对结巴作者的不敬。
最后感谢结巴库作者，里面的字典以及一些代码思路来自于结巴库。也希望大家以后能有更多交流。
On Tue, Aug 20, 2013 at 5:40 PM, lewsn2008 notifications@github.com wrote:

求语料数据，程序中的yuliao_onlyseg.txt还有吗？想跑一下学习学习，thks
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/87#issuecomment-22933170
.
```
```text
 @jannson , 已关注。
```
```text
 @lewsn2008 , 这个文件有200多MB，怎么发给你？
```
```text
 能否发个dropbox链接，我也想下载一份。
wuyanyi09@gmail.com
发件人： Sun Junyi
发送时间： 2013-08-21 17:13
收件人： fxsjy/jieba
主题： Re: [jieba] 关于idf.txt (#87)
@lewsn2008 , 这个文件有200多MB，怎么发给你？
—
Reply to this email directly or view it on GitHub.
```
```text
 @aszxqw , @lewsn2008 ，试用了一下百度云盘，分享地址：http://pan.baidu.com/share/link?shareid=4094310849&uk=1124369080
```
```text
 if w in jieba.FREQ:
        d[w]=d.get(w,0.0) + 1.0

jieba.FREQ现在已经不存在了，请问现在应该如何写if w in   ?
```
```text
 Try `jieba.dt.FREQ`


On Feb 24, 2018 17:55, "xunyl" <notifications@github.com> wrote:

    if w in jieba.FREQ:
        d[w]=d.get(w,0.0) + 1.0

jieba.FREQ现在已经不存在了，请问现在应该如何写if w in ?

—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub
<#87 (comment)>, or mute
the thread
<https://github.com/notifications/unsubscribe-auth/AA0SqiLtUKC9m_qw6PFto87F-Vh10uWwks5tX9x8gaJpZM4A2-YO>
.
```
## 398.支持PyPy
Jul 27, 2013
```text
 import jieba

jieba.suggest_freq("{UM}", tune=True)
jieba.lcut("{UM}卡上写的地址就是那个", HMM=False)
# ['{', 'UM', '}', '卡', '上', '写', '的', '地址', '就是', '那个']
jieba.suggest_freq("卡上", tune=True)
jieba.lcut("{UM}卡上写的地址就是那个", HMM=False)
# ['{', 'UM', '}', '卡上', '写', '的', '地址', '就是', '那个']
suggest_freq只能处理全部为中文的情况, 如果希望把特殊符号识别成一个词就会出错.
```
```text
 请问下同义词该怎么处理呢？
如：
北京大学,北大,pku
清华大学,清华,Tsinghua University
```
```text
 具体指的是prob_*.py文件里面的数据
是不是对训练的样本的每个字都分别贴上‘BMES’的标签？然后再统计出三个概率表？
```
```text
 模型的数据是如何生成的？ #7
```
```text
 目前jieba的大部分功能可以在PyPy环境运行。
希望能作为官方的标准环境，以及性能数据用PyPy环境为准。
:)
可以大大增强性能数据的说服力。
```
```text
 @linkerlin , 目前我还没有发现结巴分词在PyPy下的兼容性问题。至于性能方面，pypy的性能并不是总是比CPython好。
```
```text
 这个要看数据量大小，我这里的现象是，可以减少一半的耗时（小时级别的任务）。
```
## 399.合并数词和量词
Jul 16, 2013
```text
 如题，多谢！
```
```text
 No description provided.
```
```text
 用正则直接匹配可不可以？
```
```text
 正则是可以，如果词库中包含电影电视名 估计能直接拆分出来
```
```text
 用 jieba.add_word
```
```text
 This simple sentence:
今天下午三点提醒我开会,
when I am using posseg but pos this sentence, got this result:
[pair('今天下午', 'nr'), pair('三点', 'm'), pair('提醒', 'v'), pair('我', 'r'), pair('开会', 'v')]

Obviously, this not we want, 今天下午 should be 't' rather than 'nr'.
Even I am using user_dict to separate 今天 and 下午， no result.
```
```text
 这个jieba词性烂的很，不建议用，应该是训练数据问题
```
```text
 当然跟分词效果也有很大关系，你这个明显是分词不对，应该是今天/下午
```
```text
 真的有这么烂？？
```
```text
 合并数词和量词
数量词 分开之后基本没有意义.
```
```text
 @luw2007 ， 可以考虑。不过只能在posseg子模块中做了，jieba.cut被没有识别词性。
```
```text
 @fxsjy 如果不借助结巴分词，怎么解决它比较好？
我的想法是直接用正则扫，效率很低。还请提点 :)
```
```text
 现在量词直接被放到数词里面了 有时候会导致歧义
比如杯子
水晶杯 水泥杯 冷藏杯 等 直接会被解析成 n + q ..
```
## 400.请问，jieba有关于停用词的处理吗
Jul 13, 2013
```text
 我在加载了自定义词典之前enable_parallel，发现分词结果并没有使用词典。如果在加载自定义词典之后再enable_parallel则能按照预期结果分词。
#encoding=utf-8
import sys
import jieba
jieba.enable_parallel(8)
jieba.load_userdict('user.dict')

words = jieba.cut("热水器加热时间太长")
print('/'.join(words))

结果是 “热水器/加热/时间/太/长”
#encoding=utf-8
import sys
import jieba
jieba.load_userdict('user.dict') 
jieba.enable_parallel(8)      # 放在load_userdict后面了

words = jieba.cut("热水器加热时间太长")
print('/'.join(words))

结果是*“热水器/加热时间/太长”*
词典里是有“加热时间”，所以可以猜测第一个代码分词前没有载入自定义词典，可能是因为enable_parallel放load_userdict前面的。这不一致的情况是不是多线程读入词典的bug？
```
```text
 如题
```
```text
 @gausszh ，jieba.cut没有对停用词的处理。 jieba.analyse.ChineseAnalyzer有停用词的处理：https://github.com/fxsjy/jieba/blob/master/jieba/analyse/analyzer.py
```
```text
 只能通过这样的方法加停用词么？如果停用词表过大怎么办呢？可不可以通过自定义词典给超高词频，来削弱停用词以达到过滤的效果？
```
```text
 @Jaybeka 我做了一些尝试，请你参考！
你需要到对应的文件夹里找到analyzer.py并做相应的修改即可，其中stopword.txt是用户定义的停用词表
import codecs

if os.path.exists(r"/Library/Python/2.7/site-packages/jieba/analyse/stopword.txt"):
    print "Using your own stopwords.\n"
    STOP_WORDS = frozenset(( line.rstrip() for line in codecs.open('/Library/Python/2.7/site-packages/jieba/analyse/stopword.txt','r','utf-8') ))
else:
    print "Using jieba's stopwords.\n"
    STOP_WORDS = frozenset(('a', 'an', 'and', 'are', 'as', 'at', 'be', 'by', 'can',
                        'for', 'from', 'have', 'if', 'in', 'is', 'it', 'may',
                        'not', 'of', 'on', 'or', 'tbd', 'that', 'the', 'this',
                        'to', 'us', 'we', 'when', 'will', 'with', 'yet',
                        'you', 'your',u'的',u'了',u'和'))
```
```text
 @fxsjy 我在实验效果的时候看到这样的情况：

我想应该是analyzer此时使用了搜索引擎模式，我最近在做文本分析，想要把结果中“人民大会堂”去掉（即不含“人民”）；请问需要如何做呢？
```
```text
 @zihaolucky
谢谢你的建议，我现在在研究关键词提取的问题，而原作者@fxsjy 好像没有把analyzer和analyse的停词表统一，所以我对analyse/init.py进行了修改，加了支持导入自定义停词表和自定义语料库（用于得到tf-idf），修改后的文件如下
https://github.com/Jaybeka/jieba/blob/master/jieba/analyse/__init__.py
因为python新手，可能写得比较复杂，还请多多指教。‘
但我现在遇到的问题是用微博爬取的数据，关注了什么值得买后，关键词前N个一定是与之相关的。
我不知道如果调低其tf-idf，是算法问题？还是语料库不够大？
测试程序如下，备选文本还有log.txt，进行相应替换就好了
https://github.com/Jaybeka/jieba/blob/master/test/test_stopword_tfidf.py
```
```text
 @Jaybeka 关键词提取的问题要请 @fxsjy 同学来解答了。如果你可以把例子截图放上来应该就更清楚了^^;
```
```text
 @fxsjy ，如何升级到最新版的jieba
```
```text
 @just4thu , 最简单的方法是 pip install -U jieba
```
```text
 @fxsjy ，thx very much
```
```text
 @fxsjy 請問如何在斷詞之後使用停用詞處理??
```
```text
 修改https://github.com/fxsjy/jieba/blob/master/jieba/analyse/analyzer.py这个链接给出的对应文件中ChineseTokenizer类就可以了，把mode改成default
```
```text
 @gausszh ，jieba.cut没有对停用词的处理。 jieba.analyse.ChineseAnalyzer有停用词的处理：https://github.com/fxsjy/jieba/blob/master/jieba/analyse/analyzer.py

请问为什么cut/cut_for_search 没有停用词的支持呢？我发现用jieba.analyse.extract_tags 抽取关键词虽然支持停用词字典，但原本就会把一些较不重要的词过滤掉，而被过滤掉的词可能对搜索会有用，比如谁的动物的脚谁是最长的返回['最长', '动物']，而'脚' 对于搜索引擎来说其实也是一个关键词。
或者我直接将停用词加载到一个集合中，自己在cut/cut_for_search后过滤掉？
```
```text
 单个字的词结巴分不出的，所以脚字不会出
```
```text
 @ShenDezhou 其实我想问为什么cut/cut_for_search 中不能加入停用词？
```
```text
 @morefreeze jieba做的是中文分词，先用(\r\n|\s)正则分出来句子后，再进行前缀匹配或者用HMM从概率分布图上计算最短路径，在最新版本引入的ChineseAnalyzer模块作者用了whoosh模块的StopAnalyzer来剔除分词后出现的停用词。如果你想在分词前剔除停用词，可以在修改re_skip_default和re_skip_cut_all正则，但可能会影响jieba分词的F1Measure结果。
```
```text
 @ShenDezhou 感谢回复！
```
## 401.夏天能穿多少就穿多少 冬天能穿多少穿多少
Jun 28, 2013
```text
 无论是否翻墙，均无法访问，也ping不到
```
```text
 @Pzoom522 这个链接已经崩了很久了。
```
```text
 这个分词里面多少只有一种词性吗。
```
```text
 逻辑上 第一个多少应该分为 多 和 少
```
```text
 @sidealice , 这个难度挺高的。先mark一下。
```
```text
 这个好难哈哈
```
```text
 这种问题估计本身就有问题：比如我在你那句接一句。你说的多少是多少。只有说得人才知道哪个是多少，哪个是多+少了，甚至乎自己都不知道。这是语言表达本身就不够准确规范，让人都猜不透的，估计就不是难的问题了吧？
```
```text
 这个属于语意消除歧义，以现阶段的计算能力，是很难的。
```
## 402.用户字典
May 28, 2013
```text
 部分地名词性为nr，比如“孙吴县”。建议按照民政部最新的行政区划表更新词库。
2018行政区划代码
```
```text
 如何解决向字典添加“leap motion”这种中间带空格的关键词
```
```text
 @microcao , 目前不行，得改代码。 词典里面的分割符就是空格，此外还要改一处代码里面的正则表达式，把空格也作为成词的合法字符。
```
```text
 @fxsjy 能帮忙搞一下带空格的英文单词的实现吗？感谢！
```
```text
 同问。@fxsjy
```
```text
 @geekan 这会带来另一个问题。如果“leap motion”作为一个词的话，那么就要考虑：搜索引擎模式下leap和motion怎样切分出来？
```
```text
 @anderscui 不了解这里的逻辑，不过这个看起来像是一个通用的问题：

长短词同时存在时，搜索引擎模式能否全切分？我觉得是必须要全切分的。
```
```text
 @geekan 对，所以我在考虑，能否先按照现有方式切分，然后根据一个配置文件将包含空格的词merge，而不是把这样的词直接加入自定义词典。
```
```text
 @anderscui 不是很明白，为啥要单独定义配置文件呢？
是因为搜索引擎模式对空格有特殊处理吗？
```
## 403.结巴分词的关键词提取
May 25, 2013
```text
 您好，每一次我用jieba进行分词的时候，都会有
Building prefix dict from the default dictionary ...
Loading model from cache /tmp/jieba.cache
Loading model cost 0.128 seconds.
Prefix dict has been built succesfully.

这样的提示。怎么去除这些提示呢？
```
```text
 找到jieba库的__init__.py， 大约28~30行左右，
log_console = logging.StreamHandler(sys.stderr)
default_logger = logging.getLogger(__name__)
default_logger.setLevel(logging.DEBUG)
default_logger.addHandler(log_console)

将default_logger.setLevel(logging.DEBUG) 改为 default_logger.setLevel(logging.INFO) 试试。
```
```text
 @oisc it shows error:
Traceback (most recent call last):
  File "xxx.py", line 29, in <module>
    jieba_logger.handlers.clear()
AttributeError: 'list' object has no attribute 'clear'
```
```text
 @IvyGongoogle
try del jieba_logger.handlers[:] if you are using Python2
```
```text
 @aqua7regia it works. thank you~~
```
```text
 jieba.setLogLevel(20)
或
import logging
jieba.setLogLevel(logging.INFO)
```
```text
 @oisc  sorry, What do you mean for del jieba_logger.handlers[:]?
```
```text
 @gaoyangthu it works, but It seems to be jieba.setLogLevel(logging.INFO)
```
```text
 @IvyGongoogle You're right!
setLogLevel(0) equals setLogLevel(logging.NOTSET)
setLogLevel(10) equals setLogLevel(logging.DEBUG)
setLogLevel(20) equals setLogLevel(logging.INFO)
setLogLevel(30) equals setLogLevel(logging.WARNING)
setLogLevel(40) equals setLogLevel(logging.ERROR)
setLogLevel(50) equals setLogLevel(logging.CRITICAL)
```
```text
 @gaoyangthu  yes, thank you very much~~
```
```text
 提取的内容，方便排除部分标点，英文和数字么，需要调整一下吗？
```
```text
 你可以看看test.py的样本和结果
lizhengfu notifications@github.com编写：

提取的内容，方便排除部分标点，英文和数字么，需要调整一下吗？
—
Reply to this email directly or view it on GitHub.￼
```
```text
 @youyudehexie , 你是说stopwords？ 现在还没有提供接口，你可以先试一试效果：https://github.com/fxsjy/jieba/blob/master/test/extract_tags.py
```
```text
 THX??ֹͣ?? ?? ???ˣ???Ϊ?ҷ????????δ??뷢?ֻ???ȡ?????ֺͳ??ȴ???1?ı??㣬?Լ???????д??һ?£??ҵ??????ǣ???ȡ??ʱ?????????ֺͲ??ֱ??㣬?ǳ??????Ƶ?˼·????????ʱ?????޸ġ?
------------------ ԭʼ?ʼ? ------------------
??????: "Sun Junyi"notifications@github.com;
????ʱ??: 2013??5??25??(??????) ????9:09
?ռ???: "fxsjy/jieba"jieba@noreply.github.com;
????: "lizhengfu"405574395@qq.com;
????: Re: [jieba] ???ͷִʵĹؼ?????ȡ (#54)
@youyudehexie , ????˵stopwords?? ???ڻ?û???ṩ?ӿڣ???????????һ??Ч????https://github.com/fxsjy/jieba/blob/master/test/extract_tags.py
??
Reply to this email directly or view it on GitHub.
```
```text
 @youyudehexie ，乱码？
```
```text
 THX，停止词 和 过滤，因为我发现用这段代码发现会提取到数字和长度大于1的标点，自己简单重写了一下，我的疑问是，提取的时候，有数字和部分标点，是出于设计的思路，还是暂时不想修改。
```
```text
 @youyudehexie , 的确是没有过滤数字和长度大于1的标点。 有些时候数字也会成为关键词吧，比如911。标点的确是应该过滤，你是怎么修改的？可以给我发一个pull request。
这是我目前的实现，其实很简单：
https://github.com/fxsjy/jieba/blob/master/jieba/analyse/__init__.py
```
```text
 @fxsjy 不知道你现在做了标点没有，
刚刚感觉标点不是个问题。。。直接替换成空格就好，长度大于2的空格再替换成长度为1的空格。
但是试了一下发现用正则表达式不对，\D连汉字都匹配上了。。。
最近没时间做一个分支T^T
查了一下
正则表达式的[[:punct:]]可以匹配标点符号，我只在编辑器里试了一下。。。re模块好像要调整一下什么的，，
```
```text
 @zjjott ,[[:punct:]] 不能搞定汉语标点符号吧？
```
```text
 那我继续测试一下。。。。
希望早日可以写个pull。。。
2013/6/14 Sun Junyi notifications@github.com

@zjjott https://github.com/zjjott ,[[:punct:]] 不能搞定汉语标点符号吧？
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/54#issuecomment-19439691
.


朱涧江
微笑的猪头，帅气非凡，23333
果壳网，科技有意思哟，23333
www.guokr.com
```
## 404.关于利用动态规划求解最大概率路径
May 14, 2013
```text
 textrank中的rank()方法参考了pagerank的思路，并且在程序中iterate 10次。但是对于jieba案例，可否考虑把阻尼因子设定为1，基于这一点，可以推断出每一个节点的weight就是这个点的连线数，没有必要iterate 10反复迭代了。修改方法如下：
    for n, out in self.graph.items():
        ws[n] = wsdef
        outSum[n] = sum((e[2] for e in out), 0.0)

    # this line for build stable iteration
    # sorted_keys = sorted(self.graph.keys())
    # for x in xrange(10):  # 10 iters
    #     for n in sorted_keys:
    #         s = 0
    #         for e in self.graph[n]:
    #             s += e[2] / outSum[e[1]] * ws[e[1]]
    #         ws[n] = (1 - self.d) + self.d * s
    ws=outSum
    (min_rank, max_rank) = (sys.float_info[0], sys.float_info[3])
```
```text
 你好，我最近在看结巴分词的算法，看到在判断切分组合的时候，利用的是动态规划来求出最大组合概率，我觉得这个方法挺好。
不过，我有一个想法，不知道这样想对不对，还请指教。
我是这样想的，因为之前已经形成了切分词图，我在想可不可以利用图论里的最短路径的相关算法来做判断。但是最短路径求得是权值最短，如果把概率作为权值的话，这样直接算就算的是最小概率了。之前我看到过一篇论文有提出一个解决这种问题的办法，利用最短路径算法来求最大权值。
需要对权值做一定的转换。设边的权值为f（也就是词的频率或概率），取f的自然底数的对数的相反数，-ln(f)，作为新的权值w，即w=-ln(f)，利用这个权值w来求最短路径。
可以看出：
1，f越大，w越小
2，min(w1+w2+w3+...+wi)对应的是max(ln(f1_f2_f3_...fi))，也就是说求出有新权值w的图的最短路径后，相应的那个路径的组合概率 P=f1_f2_f3..._fi 是最大的，也就是我们要找的
3，由于f本来就是概率，所以0<f<1，因此w是非负数
所以可以通过求最短路径的方法来得到最大组合概率。
```
```text
 我感觉和我用的算法没有本质不同啊，最开始的时候我是用的概率相乘求最大概率路径，后来为了避免下溢，改为对log(p)求和来算最大概率路径。另外，由于是有向无环图，在有向无环图上求最长路径本来就是可以用动态规划的，不需要在转化为求最短路径问题吧。
```
## 405.提供选项控制输出
May 3, 2013
```text
 No description provided.
```
```text
 英文的你要中文分词怎么做，加入词典看看
```
```text
 同样的问题 可以用同义词 但是试了不好用
```
```text
 这里可以将Python ， C++分出来。
```
```text
 建议提供 debug 选项控制类似下面这样的输出
print >> sys.stderr, f_name, ' at line', lineno, line
```
```text
 @shuge ， 什么样的控制？是控制让这些日志不显示，还是说输出到其他地方？
```
```text
 def foo(..., debug=False, log_file=sys.stderr):
    if debug:
        print >> log_file, f_name, ' at line', lineno, line
    ...

您觉得这样会不会更好些？
```
```text
 @shuge , 好的，谢谢你的建议。你要是能给我发pull request就更好了:-)
```
```text
 稍后吧，我上周临时用 pymmseg + AdvancdLangConv 解决了
jieba 是一个很有意思的项目
```
## 406.你好~ 在打包程序的时候出现了问题
Apr 10, 2013
```text
 例如（￣▽￣），(｀・ω・´)等颜文字表情。
在处理的时候，即使添加 add_word 或者 load_userdict ，结果都会被分开成：
（   ￣   ▽   ￣   ）
```
```text
 你好我在我的一个小项目里使用了结巴分词，首先对您的努力表示感谢
我在使用py2exe打包我的程序的时候，其他的模块都没有出现问题，打包能够成功，但在运行的时候一直显示
IOError: [Errno 2] No such file or directory: 'E:\python\links\romeo\dist\l
ibrary.zip\jieba\finalseg\prob_start.py'
我单独写了一个.py文件，只有一句话：
import jieba
seg=jieba.cut('hello jieba')
打包，但运行的时候也是上面的那个结果。

我不知道您是否打包过使用了结巴的程序，如果有，您遇到过这样的问题吗？
或者您知道问题可能出现在哪吗？
谢谢啦~
```
```text
 @lisnb , 我没有用过py2exe。看错误信息，'E:\python\links\romeo\dist\l
ibrary.zip\jieba\finalseg\prob_start.py' 这个py文件被打包在zip里面了，所以读不出来。等会我装个py2exe试一试
```
```text
 @fxsjy  嗯，非常感谢您的回复
您试过之后应该就会看到，所有的依赖项都会被打包在一个zip文件里（如果没有其他设置的话）
也有可能是py2exe的问题，那么您用过其他的什么打包的工具打包过使用了结巴的程序吗？
```
```text
 @lisnb ，试了一下，确实不行。 路径问题不好处理啊，数据文件应该怎么放？
```
```text
 @fxsjy
网上是这么说的，不知道你问的是不是这个
from distutils.core import setup
import glob
import py2exe
setup(console=["myscript.py"],
data_files=[("bitmaps",
["bm/large.gif", "bm/small.gif"]),
("fonts",
glob.glob("fonts*.fnt"))],
)
说明：data_files选项将创建一个子目录dist\bitmaps，其中包含两个.gif文件；一个子目录dist\fonts，其中包含了所有的.fnt文件。
我在想的是，好像分词的字典没有被收进包里，为什么代码文件也没有呢，好奇怪。我以前也没有打包过，我也研究一下...
```
```text
 问题还在，我尝试手工把prob_start.py插入到zip包中，可这也不是解决方法，而且不管用。
```
```text
 @yangboz , 这个问题解决了，你git pull更新一下。 我用cxfree试验成功了。
测试程序： hello.py
#encoding=utf-8
import jieba
jieba.set_dictionary("./dict.txt")
jieba.initialize()
s="我研究生命起源。"
print " ".join(jieba.cut(s))
dict.txt 放在hello.py 同级目录。
```
```text
 pull下来安装最新的jieba 代码后，引用“jieba.set_dictionary("./dict.txt")jieba.initialize()” 路径错误信息是：jieba\posseg..dict.txt;
不用“jieba.set_dictionary("./dict.txt")jieba.initialize()” 后的错误信息是：jieba\posseg\dict.txt;
```
```text
 @yangboz , sorry，昨天的更新没有解决posseg读词典的路径问题，麻烦你再pull一下试一试。
下面是我的测试程序：
#encoding=utf-8
import jieba
jieba.set_dictionary("./dict.txt") #相对路径
#jieba.set_dictionary("c:/tmp/dict.txt")  #也支持绝对路径
jieba.initialize()

from jieba import posseg

s="我研究生命起源。"
print " ".join(jieba.cut(s))
for w in posseg.cut(s):
    print w.word,w.flag
```
```text
 pull到最新代码，最后还有提示提示 File "genericpath.pyc", line 54,in getmtime
WindowError: [Error 3] " "xxx\dist\library.zip\jieba\dict.txt"
我用的是 py2exe，和cxtree应该差不多。
```
```text
 @yangboz , 必须调用jieba.set_dictionary明确指出词典的路径，否则会在默认位置寻找dict.txt。但是被py2exe打包之后就找不到了。
```
```text
 pull到最新的，代码有看到有变化，pseg_cut 有过滤一些空白词汇了，但是最后打包成exe的时候，同时“明确指出词典的路径”后，原来utf-8编码到exe运行的时候成了“cp936”了，还不确定是py2exe的问题。
```
```text
 车库咖啡会场现场解决了:-)
总结如下：
str()操作中文字符+操作，python转换尝试寻找系统默认的编码;
dict定义中文需要"\u";
文件io输出也需要"encode('utf-8)", py2exe中文问题解决.
```
```text
 @yangboz ， 昨晚很有意思，大牛们在上面讲，咱俩在调程序。
```
```text
 你好，这个问题解决了吗？我也碰到相同的问题，很懊恼，jieba官网上说，支持py2exe的，但是制作后，还是不行
```
```text
 在上次来源大会现场作者帮助下，解决了，麻烦你贴下你的错误信息！
Send from Yangbo's iPhone.

On 2013年10月23日, at 15:17, AlgorithmFan notifications@github.com wrote:
你好，这个问题解决了吗？我也碰到相同的问题，很懊恼，jieba官网上说，支持py2exe的，但是制作后，还是不行
—
Reply to this email directly or view it on GitHub.
```
```text
 你好，按照你上面说的方法，加上
jieba.set_dictionary("dict.txt")
jieba.initialize()
仍然出现下面的错误。
Traceback (most recent call last):
File "mainWidget.py", line 6, in 
File "docKeyword.pyc", line 13, in 
File "jieba\posseg__init__.pyc", line 61, in 
File "jieba\posseg__init__.pyc", line 20, in load_model
IOError: [Errno 2] No such file or directory: 'E:\crawWenkuBC\dist\library.zip\jieba\dict.txt'
```
```text
 看错误信息还是这个dict.txt的路径问题，这个文档一直说的很晦涩，你先参考下我这打包程序中的路径 https://github.com/yangboz/hairy-avenger/tree/master/KingOfProgrammer/src/Q2
```
```text
 谢谢，已经解决，犯了低级错误，应该将
jieba.set_dictionary("dict.txt")
jieba.initialize()
直接放在import jieba后面，然后再进行import jieba.posseg as pseg,这样才可以在程序中修改dict.txt的路径，再次表示感谢。
```
```text
 Building prefix dict from C:\Users\game\Desktop\qt\test\dist\dict.txt ...
Loading model from cache c:\users\game\appdata\local\temp\jieba.ua7197607c9829a7854ca3e54b4005544.cache
Loading model cost 0.355 seconds.
Prefix dict has been built succesfully.
Traceback (most recent call last):
File "test.py", line 7, in 
from jieba import analyse
File "jieba\analyse_init_.pyo", line 9, in 
File "jieba\analyse\tfidf.pyo", line 65, in init
File "jieba\analyse\tfidf.pyo", line 42, in init
File "jieba\analyse\tfidf.pyo", line 47, in set_new_path
IOError: [Errno 2] No such file or directory: 'C:\Users\game\Desktop\qt\test\dist\library.zip\jieba\analyse\idf.txt
idf.txt还是不行
`# -- coding: utf-8 --
import sys
sys.path.append("../")
import jieba
jieba.set_dictionary("./dict.txt")
jieba.initialize()
from jieba import analyse
jieba.analyse.set_idf_path("./idf.txt")
if name == 'main':
seg_list = jieba.cut("我来到北京清华大学", cut_all=False)
print("Default Mode: " + "/ ".join(seg_list))  # 默认模式
s = "此外，公司拟对全资子公司吉林欧亚置业有限公司增资4.3亿元，增资后，吉林欧亚置业注册资本由7000万元增加到5亿元。吉林欧亚置业主要经营范围为房地产开发及百货零售等业务。目前在建吉林欧亚城市商业综合体项目。2013年，实现营业收入0万元，实现净利润-139.13万元。"
for x, w in jieba.analyse.extract_tags(s, withWeight=True):
    print('%s %s' % (x, w))`

很简单的代码
```
```text
 还是不行哈
```
## 407.处理词典时的建议
Mar 28, 2013
```text
 通过 jieba.set_dictionary("."+os.sep+"dict.txt") 指定dict.txt 文件位于一个带中文的父目录时，使用jieba.initialize() 进行初始化报错，报错为：
Traceback (most recent call last):
File "", line 26, in 
File "F:\Qiyi\build\getcalldatafrom_txt\out00-PYZ.pyz\jieba", line 111, in initialize
UnicodeDecodeError: 'ascii' codec can't decode byte 0xcc in position 3: ordinal not in range(128)
定位于语句： jieba_init.py_ 中的 111行： default_logger.debug("Building prefix dict from %s ..." % (abs_path or 'the default dictionary'))
将程序至于全英文目录时，运行正常。
```
```text
 我在使用jieba分词时，遇到了如下情况：
待分词句子：奥布瑞·普拉扎（Aubrey Plaza），1984年6月26日出生于美国特拉华州威尔明顿，演员。
分词结果：
jieba.cut:
奥/布/瑞/·/普拉/扎/（/Aubrey/ /Plaza/）/，/1984/年/6/月/26/日出/生于/美国/特拉华州/威尔/明/顿/，/演员
posseg.cut:
奥 nr 布 nr 瑞 ns · x 普拉 nrt 扎 v （ x Aubrey eng   x Plaza eng ） x ， x 1984 eng 年 m 6 eng 月 m 26 eng 日出 v 生于 v 美国 ns 特拉华州 ns 威尔 nrt 明 a 顿 q ， x 演员 n
出生这个词一直无法正确分出来，我发现词典中已经包含了该词：出生 1979 v，而日出频率低于出生：日出 394 v；然后我尝试自己再一次将“出生”添加到词典，以及关闭HMM，均没有作用，请问这是什么问题呢，谢谢！
```
```text
 找到了原因，日出的频率是394，出生频率是1979，生于的频率是4690，导致分为日出/生于
解决办法：
jieba.del_word('日出')
jieba.add_word('出生于')
jieba.add_word('日出')
```
```text
 你好。
在 /jieda/init.py 的 gen_trie(f_name) 函数中，目前只是简单地读出所有文件内容，然后以 \n 切分，然后再对每一行切片。
这样会有一个问题，当碰到一个空行时，程序就会出错（split(' ') 之后的三个变量赋值会不匹配），并且，这个错误处理起来还比较麻烦（事实上我改了一下默认的字典文件后就出错了，我也找不出字典哪里有问题）。
我建议是不是对空行处理一下，比如：
def gen_trie(f_name):
    lfreq = {}
    trie = {}
    ltotal = 0.0

    with open(f_name, 'rb') as f:
        for l in f:
            l = l.decode('utf-8').strip()
            if not l:
                continue
            word, freq, _ = l.split(' ')
            freq = float(freq)
            lfreq[word] = freq
            ltotal+=freq
            p = trie
            for c in word:
                if not c in p:
                    p[c] ={}
                p = p[c]
            p['']='' #ending flag

    return trie, lfreq,ltotal
```
```text
 嗯，可以考虑。不过现在加载词典已经很慢了，我需要测试加上这些条件判断后，加载速度会不会下降。
```
```text
 默认词典直接写成.py文件会不会快一些
```
```text
 @xluer ，现在的机制是第一次从文本加载，然后会用marshal.dump把词典序列化到临时磁盘文件。以后再启动时，会比较dump文件和词典文本文件的时间戳，如果dump文件更新，就不会读取文本文件了。
为什么用文本文件？-- 最初的想法是便于用于修改词典文件。
```
## 408.能整合到sphinx中么？
Mar 17, 2013
```text
 如果我输入 “15路公交” 或“150平米以上的房子”
现在的切割是 15/公交 150/平米/以上/的/房子
我自定义词典里词频调成 0 或100都不行，谢谢
```
```text
 同问
能不能加入正则规则？
```
```text
 能否考虑，先切分再用某种方法合并？
```
```text
 嗯嗯...
```
```text
 您解决了么？
```
```text
 大工程啊 现在没有好用的中文引擎
```
```text
 我也想问这个问题，Coreseek是基于sphinx的。。
```
```text
 那几个分词太让人纠结了
```
```text
 @fxsjy 老大考虑下
```
## 409.jieba 能否部署到新浪SAE中？如何操作？
Mar 5, 2013
```text
 比如：“我来到北京清华大学”  分成： 我 / 来 / 到 / 北 / 京 / 清 / 华 / 大 / 学
```
```text
 这种事儿也用分词么？

2017-11-21 20:00 GMT+08:00 jasmineol <notifications@github.com>:
…
 比如：“我来到北京清华大学” 分成： 我 / 来 / 到 / 北 / 京 / 清 / 华 / 大 / 学

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#552>, or mute the thread
 <https://github.com/notifications/unsubscribe-auth/ABD5_ejuZPwUklkvnwSV_5J3028wkVO3ks5s4rtUgaJpZM4QluiD>
 .
```
```text
 这个网上随便一搜都是分字的
比如把中文字分隔转为list
```
```text
 补充一下，c++如何使用jieba分字，涉及到gbk，utf8这种编码的转换，所以想看下jieba有没有现成的
```
```text
 这种直接获取编码就是一个词吧，之前看过汉子转拼音。
```
```text
 设置一个空词典，关闭HMM新词联想，分出来全是一个个字的
```
```text
 哪位大牛能教教我，谢谢
```
```text
 把jieba.cache的位置由临时目录改为当前目录就可以了，我周末刚刚修改并且成功部署了：fay@e094dac#L0L65
```
## 410.能否为cut_for_search方法添加上词性标注？
Mar 4, 2013
```text
 刚开始导入库时设定了自定义词典:
import jieba
jieba.load_userdict(
    "/home/weiwu/share/deep_learning/data/dict/finance_dict.txt")


如果程序中途想更换另一个词典，请问该怎么做？
```
```text
 reload(jieba)?
```
```text
 现在cut方法有了词性标注，但cut_for_search方法好像没有词性标注。我现在在用Jieba做搜索引擎，想要个词性标注，以便获取不同的词性，不知道能否添加上？另外，能否有份关于词性说明的文档，即说明每个词性符号代表着哪种词性？
```
```text
 我也觉得应该返回词性标注，因为有时词性是很重要的。最好再加个方法来获取单个词的词性。 @xgfone 要找的词性意思，可以去 中科院的分词网站去下载 《ICTPOS3.0汉语词性标记集》我已经下载了一份，很是全面。
```
## 411.关键词提取质量很差
Feb 18, 2013
```text
 文档中写到，词频省略时使用自动计算能保证分出该词的词频。
例如，“吉林省延边朝鲜族自治州”，期望分词为“吉林省”和“延边朝鲜族自治州”。添加自定义词典并省略词频，并没有成功分出期望的结果。
In [126]: jieba.lcut('吉林省延边朝鲜族自治州',  HMM=False)
Out[126]: ['吉林省延边朝鲜族自治州']
请指点！
```
```text
 jieba.del_word('吉林省延边朝鲜族自治州')
```
```text
 @Silencezjl 这样一个一个del_word也不是办法...似乎得完全自定义词典
```
```text
 使用jieba提取readme内容中的关键词，返回的结果是：
jieba,the,cut,to,list
这个结果中是否可以将the，to剔除掉，有没有什么选项？
```
```text
 @yukaizhao, 主要是idf.txt中只有中文词汇。我会加强这一块，需要分析英文语料。
```
```text
 @fxsjy 我也碰到这个问题，但我试着添加了英文的idf数据，但是仍然有the,to这样的英文出现，能否添加stop words？
```
```text
 @fay, @yukaizhao , 最新版本已经对关键词提取功能做了一些改进，也加了一些英文的stop words。 抽取效果Demo: http://jiebademo.ap01.aws.af.cm/extract
```
```text
 试了一下，效果好多了，正好有件事儿请教，现在的master版本是稳定版本吗，我想换一个小点的词典，更新词典后发现和以前的jieba版本不兼容。
```
```text
 @yukaizhao , master版本不是稳定版，但是应该没有兼容性的问题吧。出错信息是什么？
```
```text
 确实是词典的格式不一样，新的词典添加词性，老词典没有词性
老词典：
一万三千五百斤 4
新词典：
她 134035 r
请问哪儿可以下载最新的稳定版呀？
```
```text
 https://pypi.python.org/pypi/jieba
```
```text
 我自己在网上爬了1500个网页评论，训练了一个idf.txt，然后读取前面的名词，也就是字典里有 "n"的词，效果还很不错。有需要我可以给你。
On Thu, Aug 22, 2013 at 1:26 PM, Zoey Young notifications@github.comwrote:

抽取关键词中含...
—
Reply to this email directly or view it on GitHubhttps://github.com//issues/27#issuecomment-23069212
.
```
```text
 @jannson 求教怎么训练idf
```
## 412.关于jieba的分词效果
Jan 16, 2013
```text
 词性 eng 是啥？   为什么官方没有词性对照表？  好纠结， 上网查了资料也找不到eng是啥
```
```text
 http://blog.csdn.net/syani/article/details/52276282
```
```text
 eng是英语的意思。
我也没找到官方的词性对照表。楼上那个词性对照表是很久之前的东西了，现在使用的词性集又加入了一些新的词性。
```
```text
 词性参考ICTCLAS呀：
POS = {
    "n": {  # 1. 名词  (1个一类，7个二类，5个三类)
        "n": "名词",
        "nr": "人名",
        "nr1": "汉语姓氏",
        "nr2": "汉语名字",
        "nrj": "日语人名",
        "nrf": "音译人名",
        "ns": "地名",
        "nsf": "音译地名",
        "nt": "机构团体名",
        "nz": "其它专名",
        "nl": "名词性惯用语",
        "ng": "名词性语素"
    },
    "t": {  # 2. 时间词(1个一类，1个二类)
        "t": "时间词",
        "tg": "时间词性语素"
    },
    "s": {  # 3. 处所词(1个一类)
        "s": "处所词"
    },
    "f": {  # 4. 方位词(1个一类)
        "f": "方位词"
    },
    "v": {  # 5. 动词(1个一类，9个二类)
        "v": "动词",
        "vd": "副动词",
        "vn": "名动词",
        "vshi": "动词“是”",
        "vyou": "动词“有”",
        "vf": "趋向动词",
        "vx": "形式动词",
        "vi": "不及物动词（内动词）",
        "vl": "动词性惯用语",
        "vg": "动词性语素"
    },
    "a": {  # 6. 形容词(1个一类，4个二类)
        "a": "形容词",
        "ad": "副形词",
        "an": "名形词",
        "ag": "形容词性语素",
        "al": "形容词性惯用语"
    },
    "b": {  # 7. 区别词(1个一类，2个二类)
        "b": "区别词",
        "bl": "区别词性惯用语"
    },
    "z": {  # 8. 状态词(1个一类)
        "z": "状态词"
    },
    "r": {  # 9. 代词(1个一类，4个二类，6个三类)
        "r": "代词",
        "rr": "人称代词",
        "rz": "指示代词",
        "rzt": "时间指示代词",
        "rzs": "处所指示代词",
        "rzv": "谓词性指示代词",
        "ry": "疑问代词",
        "ryt": "时间疑问代词",
        "rys": "处所疑问代词",
        "ryv": "谓词性疑问代词",
        "rg": "代词性语素"
    },
    "m": {  # 10. 数词(1个一类，1个二类)
        "m": "数词",
        "mq": "数量词"
    },
    "q": {  # 11. 量词(1个一类，2个二类)
        "q": "量词",
        "qv": "动量词",
        "qt": "时量词"
    },
    "d": {  # 12. 副词(1个一类)
        "d": "副词"
    },
    "p": {  # 13. 介词(1个一类，2个二类)
        "p": "介词",
        "pba": "介词“把”",
        "pbei": "介词“被”"
    },
    "c": {  # 14. 连词(1个一类，1个二类)
        "c": "连词",
        "cc": "并列连词"
    },
    "u": {  # 15. 助词(1个一类，15个二类)
        "u": "助词",
        "uzhe": "着",
        "ule": "了 喽",
        "uguo": "过",
        "ude1": "的 底",
        "ude2": "地",
        "ude3": "得",
        "usuo": "所",
        "udeng": "等 等等 云云",
        "uyy": "一样 一般 似的 般",
        "udh": "的话",
        "uls": "来讲 来说 而言 说来",
        "uzhi": "之",
        "ulian": "连 "  # （“连小学生都会”）
    },
    "e": {  # 16. 叹词(1个一类)
        "e": "叹词"
    },
    "y": {  # 17. 语气词(1个一类)
        "y": "语气词(delete yg)"
    },
    "o": {  # 18. 拟声词(1个一类)
        "o": "拟声词"
    },
    "h": {  # 19. 前缀(1个一类)
        "h": "前缀"
    },
    "k": {  # 20. 后缀(1个一类)
        "k": "后缀"
    },
    "x": {  # 21. 字符串(1个一类，2个二类)
        "x": "字符串",
        "xx": "非语素字",
        "xu": "网址URL"
    },
    "w": {   # 22. 标点符号(1个一类，16个二类)
        "w": "标点符号",
        "wkz": "左括号",  # （ 〔  ［  ｛  《 【  〖 〈   半角：( [ { <
        "wky": "右括号",  # ） 〕  ］ ｝ 》  】 〗 〉 半角： ) ] { >
        "wyz": "全角左引号",  # “ ‘ 『
        "wyy": "全角右引号",  # ” ’ 』
        "wj": "全角句号",  # 。
        "ww": "问号",  # 全角：？ 半角：?
        "wt": "叹号",  # 全角：！ 半角：!
        "wd": "逗号",  # 全角：， 半角：,
        "wf": "分号",  # 全角：； 半角： ;
        "wn": "顿号",  # 全角：、
        "wm": "冒号",  # 全角：： 半角： :
        "ws": "省略号",  # 全角：……  …
        "wp": "破折号",  # 全角：——   －－   ——－   半角：---  ----
        "wb": "百分号千分号",  # 全角：％ ‰   半角：%
        "wh": "单位符号"  # 全角：￥ ＄ ￡  °  ℃  半角：$
    }
}
```
```text
 see: https://gist.github.com/luw2007/6016931
same as: #453
```
```text
 首先感谢 @fxsjy 做了这个项目。
请问 @fxsjy ，jieba分词的效果如何呢，跟ICTCLASS相比呢？我是说是否做过相关的评测。
我测试了一些句子，感觉效果还可以。
```
```text
 @hitalex , 目前还没有进行学术上的规范化测试，因为不知道去哪里找这些benchmark数据
```
## 413.添加对 whoosh 的接口
Nov 6, 2012
```text
 作为一个大四学生，刚刚学过python，想看看一些成熟的项目具体代码，透彻地进行分析，不知道jieba适不适合呢？
```
```text
 個人感覺挺好的, http://www.cnblogs.com/zhbzz2007/tag/Natural%20language%20processing/
```
```text
 形容词某些情况下变成了名词了 如下所示
# 新 a
>>> pseg.lcut("新材料装备制造业的循环产业体系")
[pair('新', 'a'), pair('材料', 'n'), pair('装备', 'n'), pair('制造业', 'n'), pair('的', 'uj'), pair('循环', 'vn'), pair('产业', 'n'), pair('体系', 'n')]
# 新 n
>>> pseg.lcut("青海格尔木工业园镁锂新材料产业基地")
[pair('青海', 'ns'), pair('格尔木', 'nr'), pair('工业园', 'n'), pair('镁锂', 'nz'), pair('新', 'n'), pair('材料', 'n'), pair('产业基地', 'n')]


# 大 a
>>> pseg.lcut("大工业自由")
[pair('大', 'a'), pair('工业', 'n'), pair('自由', 'a')]
# 大 n
>>> pseg.lcut("7个国家的大公司中")
[pair('7', 'm'), pair('个', 'm'), pair('国家', 'n'), pair('的', 'uj'), pair('大', 'n'), pair('公司', 'n'), pair('中', 'f')]

有什么办法可以使其保持一致吗？
```
```text
 参考hanlp
```
```text
 可以參考我修正的錯誤 #670
測試完 您上述的語句都正常
(可能跟我修正的錯誤有關係)
#'新', 'a'
>>> pseg.lcut("新材料装备制造业的循环产业体系")
[pair('新', 'a'), pair('材料', 'n'), pair('装备', 'n'), pair('制造业', 'n'), pair('的', 'uj'), pair('循环', 'vn'), pair('产业', 'n'), pair('体系', 'n')]
#新,'n' -> #'新', 'a'
>>> pseg.lcut("青海格尔木工业园镁锂新材料产业基地")
[pair('青海', 'ns'), pair('格尔木', 'nr'), pair('工业园', 'n'), pair('镁', 'n'), pair('锂', 'n'), pair('新', 'a'), pair('材料', 'n'), pair('产业基地', 'n')]

#'大', 'a'
pseg.lcut("大工业自由")
[pair('大', 'a'), pair('工业', 'n'), pair('自由', 'a')]

#'大', 'n' -> #'大', 'a'
pseg.lcut("7个国家的大公司中")
[pair('7', 'm'), pair('个', 'm'), pair('国家', 'n'), pair('的', 'uj'), pair('大', 'a'), pair('公司', 'n'), pair('中', 'f')]
```
```text
 最近研究 中文分词，准备自己做一个，采用双向匹配分词和HMM处理未登录词、削歧义。
不过看了jieba，感觉细节已经做了很多。
准备实现 whoosh 的分词接口，就用在下一个项目中。
不知道能不能提供些 jieba设计方面的资料
```
```text
 可以提供源码 ;)
```
```text
 @terminus318 , 最近工作较忙，还未研究whoosh。 在网上搜索时发现有人对whoosh和jieba做了集成。先mark一下： http://blog.csdn.net/wenxuansoft/article/details/8170714
```
```text
 http://www.verydemo.com/demo_c230_i1903.html
```
```text
 @terminus318 , @oldcai , 结巴0.30版已经添加了用于Whoosh的分词接口：ChineseAnalyzer。
用法：https://github.com/fxsjy/jieba/blob/master/test/test_whoosh.py
```
```text
 哈哈，感谢。
另：工信处女干事好忙，每次测试都要请她过来。
```
## 414.能提供简单的说明文档吗？
Oct 18, 2012
```text
 Please add the LICENSE file to the MANIFEST.in in both jieba and jieba3k so that it's included in the source distribution.
```
```text
 之前看过你写的goseg，不过好像很久没更新了，期待Golang版的jieba
```
```text
 早就有了
https://github.com/yanyiwu/gojieba
```
```text
 看了一下结果，感觉挺好的，想研究一下整个过程，能否提供一个简单的文档说明，代码都没有注释，python菜鸟请教。
```
```text
 争取在一个月之类完善设计文档
```
```text
 期待啊，想认真学习一下楼主的代码
```
```text
 我最近正在研究爬虫来抓取数据，需要中文分词系统，也想研究一下分词系统。ICTCLAS中文分词系统竟然没有给出Python的接口，让人很无语！！！
```
```text
 文档还未完成啊~~~~，自罚一杯
```
```text
 今天发现一个网友写的分析文档，很感动，此文基本上揭示了结巴分词的原理：http://ddtcms.com/blog/archive/2013/2/4/69/jieba-fenci-suanfa-lijie/
```
```text
 多谢，楼主的什么时候出来呢？呵呵
```
```text
 @kingllon, 公司有很多码要砌，github的业务荒废有月余了，惭愧。大家可以先看ddtcms写的那个:-)
```
```text
 粗略看了一遍，的却写总结的很好。细节需要精度，多谢提供资料。
```
```text
 感谢作者的无私分享。不过好像已经打不开链接了。不知能不能重新分享一下
```
```text
 对呀，分析文档的网页打不开了，好像是服务器不行了，不知道有没有新的链接
```
```text
 分析文档的网页打不开，求新地址 ~~~谢谢谢谢谢谢
```
## 415.通过关键词获取其同义词
Oct 18, 2012
```text
 如圖所示

我是用簡體中文維基百科的語料，經過處理之後把符號換成 space，再將整個文本當成一個當成一個 string
餵進去。
下面是我的代碼
import time

start_time = time.time()

jieba.enable_parallel(16)
jieba_words = {word for word in jieba.cut_for_search(text_processed, HMM=True) if len(word) > 1 and len(word) <= 7}

print("Segmentizing took {} seconds.".format(time.time() - start_time))

謝謝幫忙！
```
```text
 一直在等待一个Ruby版本的jieba，求大神来开发一个！
```
```text
 No description provided.
```
```text
 比如：
高级Java工程师 100 position
高级Java工程师 110 certificate
词频是乱填的，两者来自不同的user_dict,会产生覆盖的情况
```
```text
 请问jieba分词有提供关键词词典吗？（用来获取该关键词的同义词）谢谢！
```
```text
 同求
```
