package internal

const LogPatg = "/Users/coffee/Downloads/com.tujigu/log/"
const FilePath = "/Users/coffee/Downloads/com.tujigu/download/"
const WebPath = "https://www.tujigu.com"

var AllDoc = map[string]string{
	"https://www.tujigu.com/a/28235/": "匿名",
	"https://www.tujigu.com/a/33875/": "萝莉",
	"https://www.tujigu.com/a/34209/": "馨馨",
	"https://www.tujigu.com/a/34661/": "沐娅",
	"https://www.tujigu.com/a/34702/": "匿名",
	"https://www.tujigu.com/a/34704/": "匿名",
	"https://www.tujigu.com/a/34831/": "dodo",
	"https://www.tujigu.com/t/1178/":  "玉兔miki",
	"https://www.tujigu.com/t/1185/":  "妲己_Toxic",
	"https://www.tujigu.com/t/1194/":  "周妍希",
	"https://www.tujigu.com/t/1289/":  "黄楽然",
	"https://www.tujigu.com/t/161/":   "糯美子Mini",
	"https://www.tujigu.com/t/190/":   "易阳",
	"https://www.tujigu.com/t/2189/":  "朴善慧",
	"https://www.tujigu.com/t/2196/":  "许允美",
	"https://www.tujigu.com/t/2234/":  "张楚珊",
	"https://www.tujigu.com/t/242/":   "沈梦瑶",
	"https://www.tujigu.com/t/2693/":  "天使萌",
	"https://www.tujigu.com/t/2783/":  "大安妮",
	"https://www.tujigu.com/t/292/":   "夏小秋",
	"https://www.tujigu.com/t/293/":   "王雨纯",
	"https://www.tujigu.com/t/295/":   "朱可儿",
	"https://www.tujigu.com/t/296/":   "许诺",
	"https://www.tujigu.com/t/298/":   "夏美酱",
	"https://www.tujigu.com/t/3156/":  "周于希",
	"https://www.tujigu.com/t/3160/":  "little贝壳",
	"https://www.tujigu.com/t/3171/":  "小狐狸Sica",
	"https://www.tujigu.com/t/3261/":  "杨暖",
	"https://www.tujigu.com/t/3307/":  "宁宁",
	"https://www.tujigu.com/t/4218/":  "任莹樱",
	"https://www.tujigu.com/t/446/":   "陈潇",
	"https://www.tujigu.com/t/4476/":  "芸斐",
	"https://www.tujigu.com/t/4530/":  "姜璐",
	"https://www.tujigu.com/t/4568/":  "Lavinia肉肉",
	"https://www.tujigu.com/t/459/":   "杨晨晨",
	"https://www.tujigu.com/t/4640/":  "徐微微",
	"https://www.tujigu.com/t/465/":   "菲儿",
	"https://www.tujigu.com/t/4708/":  "米米",
	"https://www.tujigu.com/t/5034/":  "姜仁卿",
	"https://www.tujigu.com/t/5109/":  "桜桃喵",
	"https://www.tujigu.com/t/5110/":  "疯猫ss",
	"https://www.tujigu.com/t/5466/":  "克拉女神芊芊",
	"https://www.tujigu.com/t/5496/":  "星之迟迟",
	"https://www.tujigu.com/t/5497/":  "面饼仙儿",
	"https://www.tujigu.com/t/5499/":  "雯妹不讲道理",
	"https://www.tujigu.com/t/5500/":  "一小央泽",
	"https://www.tujigu.com/t/5501/":  "鬼畜瑶",
	"https://www.tujigu.com/t/5504/":  "魔物喵",
	"https://www.tujigu.com/t/5505/":  "你的负卿",
	"https://www.tujigu.com/t/5511/":  "雪琪",
	"https://www.tujigu.com/t/5513/":  "白银81",
	"https://www.tujigu.com/t/5514/":  "米线线sama",
	"https://www.tujigu.com/t/5515/":  "黑川",
	"https://www.tujigu.com/t/5529/":  "李浅浅",
	"https://www.tujigu.com/t/5533/":  "袁圆",
	"https://www.tujigu.com/t/5559/":  "爱丽丝",
	"https://www.tujigu.com/t/5609/":  "小牛奶",
	"https://www.tujigu.com/t/5613/":  "林文文",
	"https://www.tujigu.com/t/5652/":  "郭子蜜",
	"https://www.tujigu.com/t/5672/":  "奈汐酱",
	"https://www.tujigu.com/t/5674/":  "UU酱",
	"https://www.tujigu.com/t/5697/":  "秋楚楚",
	"https://www.tujigu.com/t/5712/":  "古川kagura",
	"https://www.tujigu.com/t/5714/":  "戚顾儿",
	"https://www.tujigu.com/t/5716/":  "南桃Momoko",
	"https://www.tujigu.com/t/5720/":  "从从从从鸾",
	"https://www.tujigu.com/t/5721/":  "是依酱呀",
	"https://www.tujigu.com/t/5725/":  "清纯妹子西瓜",
	"https://www.tujigu.com/t/5728/":  "雪晴Astra",
	"https://www.tujigu.com/t/5729/":  "-白烨-",
	"https://www.tujigu.com/t/5736/":  "绮太郎",
	"https://www.tujigu.com/t/5740/":  "Nyako喵子",
	"https://www.tujigu.com/t/5743/":  "蠢沫沫",
	"https://www.tujigu.com/t/5744/":  "弥音音ww",
	"https://www.tujigu.com/t/5745/":  "镜酱",
	"https://www.tujigu.com/t/5749/":  "樱落酱w",
	"https://www.tujigu.com/t/5755/":  "萌芽儿o0",
	"https://www.tujigu.com/t/5756/":  "十万珍吱伏特",
	"https://www.tujigu.com/t/5762/":  "眼酱大魔王w",
	"https://www.tujigu.com/t/5765/":  "蜜汁猫裘",
	"https://www.tujigu.com/t/5768/":  "是青水",
	"https://www.tujigu.com/t/5774/":  "抖娘-利世",
	"https://www.tujigu.com/t/5777/":  "南鸽",
	"https://www.tujigu.com/t/5794/":  "舒彤",
	"https://www.tujigu.com/t/5799/":  "西景",
	"https://www.tujigu.com/t/5800/":  "子吟",
	"https://www.tujigu.com/t/654/":   "玛鲁娜",
	"https://www.tujigu.com/t/657/":   "蛋糕Cake",
	"https://www.tujigu.com/t/663/":   "夏天Sienna",
	"https://www.tujigu.com/t/903/":   "娜露",
	"https://www.tujigu.com/t/919/":   "李雅",
	"https://www.tujigu.com/x/2/":     "网络美女",
	"https://www.tujigu.com/x/86/":    "风之领域",
	"https://www.tujigu.com/x/95/":    "喵糖映画写真",
	"https://www.tujigu.com/t/5788/":  "腿模YuHsuan",
	"https://www.tujigu.com/t/290/":   "刘娅希",
	"https://www.tujigu.com/t/5798/":  "蓓颖",
	"https://www.tujigu.com/t/150/":   "于姬",
	"https://www.tujigu.com/t/2438/":  "小热巴",
	"https://www.tujigu.com/t/5790/":  "申才恩",
	"https://www.tujigu.com/t/5472/":  "杨紫嫣",
	"https://www.tujigu.com/t/4780/":  "小尤奈",
	"https://www.tujigu.com/t/4729/":  "九月生",
	"https://www.tujigu.com/t/5553/":  "陶喜乐",
	"https://www.tujigu.com/t/5375/":  "林子欣",
	"https://www.tujigu.com/t/5177/":  "莉娜lena",
	"https://www.tujigu.com/t/455/":   "陈思琪",
	"https://www.tujigu.com/t/954/":   "芝芝Bootyw",
	"https://www.tujigu.com/t/4634/":  "美替",
	"https://www.tujigu.com/t/4242/":  "陈亦菲",
	"https://www.tujigu.com/t/5175/":  "LindaLinda",
	"https://www.tujigu.com/t/4561/":  "水花花",
	"https://www.tujigu.com/t/4441/":  "韩羽",
	"https://www.tujigu.com/t/3852/":  "Egg_尤妮丝",
	"https://www.tujigu.com/t/3243/":  "敏珺",
	"https://www.tujigu.com/t/674/":   "佟蔓",
	"https://www.tujigu.com/t/437/":   "阿朱",
	"https://www.tujigu.com/t/5506/":  "抱走莫子",
	"https://www.tujigu.com/t/5789/":  "BamBi",
	"https://www.tujigu.com/t/5754/":  "黑猫猫OvO",
	"https://www.tujigu.com/t/5925/":  "霜月shimo",
	"https://www.tujigu.com/t/3659/":  "猫九酱",
	"https://www.tujigu.com/t/5902/":  "采妮么么",
	"https://www.tujigu.com/t/5732/":  "水淼aqua",
	"https://www.tujigu.com/t/5891/":  "塔塔_Lo1iTa",
	"https://www.tujigu.com/t/5917/":  "Azami",
	"https://www.tujigu.com/t/5724/":  "晓美嫣",
}
