package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mindguard/model"
	"mindguard/utils"
)

var (
	DB *gorm.DB
)

// 初始化数据库
func InitMysql(databaseConfig utils.DatabaseConfig) (err error) {
	dsn := databaseConfig.User + ":" + databaseConfig.Password +
		"@tcp(" + databaseConfig.Host + ":" + databaseConfig.Port + ")/" + databaseConfig.DbName + "?charset=" +
		databaseConfig.Charset + "&parseTime=" + databaseConfig.ParseTime + "&loc=" + databaseConfig.Loc
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = DB.AutoMigrate(&model.User{}, &model.Article{}, &model.Question{},
		&model.Answer{}, &model.UserRecord{}, &model.Reservation{}, &model.Evaluation{}, &model.Communication{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 初始化文章
	InitArticle()
	// 初始化老师账号
	InitUsers()
	// 初始化测试题目
	InitQuestion()
	// 初始化测试题目的答案
	InitAnswer()

	return
}

// 关闭数据库
func Close() {
	sqlDB, _ := DB.DB()
	_ = sqlDB.Close()
}

// 初始化文章
func InitArticle() {
	articles := []model.Article{
		{
			Title:       "有话好好说——克服语言暴力",
			Content:     "当今社会普遍存在一种心理障碍，叫做“讨好型人格”。\n\n讨好型人格一般有以下几种特征：\n\n1、没有自信，容易受到他人的影响，特别是注重别人对自己的评价。\n\n2、没有主见，不知道自己想要什么，在生活和工作中常常犹豫不决、不知所措。\n\n3、心思敏感细腻，能察觉到别人的情绪变化以及需求，害怕被被人忽略或者让他人不高兴。\n\n4、虽然表面友善但是内心不愿和别人太过亲近，不愿意向他人袒露自己的需求与看法。\n\n5、当他人对自己寻求帮助时，往往不太懂得拒绝别人，自己常感无奈，却又无能为力。\n\n这一类人，为人处世的原则是“牺牲自己，成全别人”。总是无条件的满足别人的要求，从不拒绝；做事之前也常常思前想后，优先去考虑别人的想法，害怕与别人起冲突……\n\n也许，看到这里你已经对号入座，找到了自己的影子。",
			Author_Name: "刘冬梅",
			Created_At:  "2020-07-17",
			ImageURL:    "https://s11.ax1x.com/2023/12/21/piTwreP.jpg",
		},
		{
			Title:       "心理咨询对焦虑症的帮助效果大吗？",
			Content:     "首先我们要明白，心理咨询它是一个需要长期坚持的过程，不是说你来个一次两次，就可以帮助你彻底解决问题，只要你在专业咨询师的帮助下，有规律，有频次的去做咨询，那么心理咨询师可以帮助你治疗焦虑症的。\n\n焦虑症的常见类型有6种：恐慌症、恐惧症、广泛性焦虑症、强迫症、社交恐惧症和创伤后压力症候群。\n\n每个类型都有它各自的诊断标准，大家普遍认为焦虑症其实是广泛性焦虑症，而它的判定标准主要有6点：\n\n1、对许多事件或活动会过度焦虑或担忧，而这些担忧必须至少在6个月内，有症状的日子比没症状的日子多；\n\n2、这些焦虑和担忧是难以控制的；\n\n3、这些焦虑和担忧，必须伴随以下6个症状中的三项以上：\n\nA. 坐立不安感到紧张心情不定容易疲劳\n\nB. 无法集中注意力脑筋一片空白\n\nC. 容易发怒\n\nD. 肌肉紧绷\n\nE. 睡眠障碍失眠和睡眠品质不佳；\n\n4、这些焦虑担忧或身体不适症状会严重影响日常生活；\n\n5、此症状不是因药物引起的；\n\n6、此症状也不能用其他疾病来解释，如恐慌症恐惧症等。\n\n一般的焦虑情绪、焦虑症是可以通过心理咨询改善的。主要方法有1，行为治疗。2，认知治疗。3，短程精神动力学心理治疗。\n\n此外，心理咨询是一个协作过程，心理咨询师和来访需要共同努力确定具体问题并开发应对焦虑的具体技能和技巧。来访可以在治疗之外练习新技能，以在可能使他们感到不舒服的情况下控制焦虑。\n\n最后，想告诉大家，焦虑症是可以治疗的。大多数患有焦虑症的患者经过几个月（或更少）个月的心理治疗后能够减轻或消除症状，并且许多患者在几次治疗后就明显感受了自己的到症状有所改善。",
			Author_Name: "章成松",
			Created_At:  "2023-10-26",
			ImageURL:    "https://s11.ax1x.com/2023/12/21/piTwBLt.jpg"},
		{
			Title:       "10个妙招赶走心理焦虑",
			Content:     "1. 充足的睡眠\n\n睡眠不足会导致严重后果：不仅影响我们的身体健康，还能造成全天焦虑和紧张。有时还会形成恶性循环，因为焦虑通常会阻碍睡眠。\n\n尤其当你感到烦躁不安时，试着制定一个7--9个小时的睡眠计划，饱饱地睡上几晚，看看白天的焦虑是不是减轻了。\n\n2. 笑一笑\n\n当工作让我们情绪低落时，迅速调整下心态，咯咯地笑几声吧。研究表明笑声能够缓解抑郁和焦虑，所以不妨从网络上找些搞笑的段子平复下紧张的神经吧。\n\n3. 简化大脑\n\n物质简化=心理简化。\n\n如果工作的地点混乱不堪，就很难放松心情，且使工作显得更加凌乱繁琐。因此花15分钟左右整理一下房间或办公桌，并养成保持事物干净的好习惯。这些可以帮助我们更理性地思考问题，也就没有焦虑的机会喽。\n\n 4. 表达感激之情\n\n研究证实常念感恩有助于减轻焦虑，尤其当我们休息充分后。因此摒弃疲惫不堪的心态，怀着感激的心态开启你的感恩之程吧。\n\n5. 吃对事物\n\n焦虑会让我们的身体乱作一团：胃口也会跟着改变。为了给身体提供所需的支持，应该选择富含维生素B和Ω-3等营养元素的食物，并配以健康的全谷物碳水化合物。\n\n研究证实维生素B与良好的精神状态有关，而Ω-3可以减少抑郁症和焦虑症。全谷物碳水化合物可以帮助调节体内五羟色胺——一种让我们“感觉良好”，并保持心态平和的神经递质——的水平。\n\n不过要注意，吃含糖量较高和加工的食品会加重焦虑症状。\n\n6. 冥想\n\n现在我们应该都知道冥想就是放松，但科学家们同时也发现冥想实际会增加大脑内的灰质——可令体内的压力减少的物质。很多专业人士都强调了冥想对焦虑、情绪和压力症状的积极作用。\n\n此外，冥想还是一种观察大脑的方法，让我们搞清楚耐人寻味的焦虑情绪到底是如何产生的。而理解大脑的思维模式有助于让我们远离那些负面情绪。\n\n7. 作一个前景板 \n\n如果未来看起来过于苍茫而可怕的话，就改变对目前现状的看法吧。花一小时制作一个前景板，有时单单只是设定具体的目标就能将我们从对未来未知的焦虑中解救出来。\n\n而对于不善手工的人来说，可以试着制作一份有趣的电子版前景图，为自己增添一些动力。在制作画板时，不妨想想“真益激必善”五字箴言：即我的想法是真实的、有益的、激励的、必要的且善良的吗？如果不是的话，赶紧摒弃掉。\n\n8. 玩起来\n\n小孩子和动物似乎天生有玩耍的能力，因为他们没有像邮箱过满这样的焦虑。直到办公室发出放假的消息，我们才必须负责任地安排自己的闲暇时光。\n\n可以提议带朋友的狗出去溜溜弯，或是帮朋友带一下午孩子，让自己的大脑放松放松，让这些无忧无虑的小朋友带动你一起玩儿吧。\n\n9. 绝对安静 \n\n计划出一段时间，让自己与外界完全隔离。先从适合自己的一小段时间开始，以便能持续下去，哪怕是很短的5分钟也行。\n\n绝对安静意味着在此期间你关掉手机，关掉电视，不看邮件，不看新闻，统统关掉不看。让别人这段时间内是联系不到你的，这样就能暂时远离忧虑。\n\n噪音过多会增加紧张程度，那就在嘈杂的日常生活中为自己留出绝对安静的片刻吧\n\n10. 提前制定计划\n\n提前抵抗焦虑的方法就是事先准备好。试着制定一份工作计划或列出待办事项，养成提高工作效率的好习惯。因此与其每天早上花十几分钟疯狂的找钥匙，倒不如养成每天回到家就把钥匙放到同一个地方的习惯。\n\n前一晚就找出要穿的衣服，装好运动背包并将其放在门口，或提前定午餐。一定要提前准备才能避免焦虑产生。",
			Author_Name: "张兰欣",
			Created_At:  "2013-10-27",
			ImageURL:    "https://s11.ax1x.com/2023/12/21/piTwsdf.jpg",
		},
		{
			Title:       "焦虑了怎么办！7种快速平复焦虑的办法！",
			Content:     "许多人都有过焦虑的体验，甚至有些人会一直被焦虑所困扰。当出现问题时，焦虑会让你更难做出决定，更不用说拿出实际行动解决问题了。焦虑会导致对每一件事的过度思虑，从而陷入“想太多”→“更焦虑”的恶性循环。\n\n要如何才能跳出这样的恶性循环呢？简单地想要停止想法是没有用的，那些想法会不断冒出来，甚至会更激烈。\n\n不过，确实有一些更加有效的技巧存在。今天我们为大家介绍从正念减压（Mindfullness Based Stress Reduction）和认知行为疗法（Cognitive-Behavioral Therapies）中提取的7种心理策略，能够有效解决焦虑问题。\n\n1．认知重构\n\n尝试把自己焦虑的想法看成是“猜测”，而不是“事实”。当你焦虑时，其实是你的头脑在尝试保护你，它希望能够预测出未来可能发生的事，从而减少如果真的事发对你造成的冲击。不过，未来“有可能”发生，并不等于未来“一定会”发生。比起一味沉浸在负面的揣测里，你需要寻找一些客观的证据(objective evidence)。可以做一张这样的表格：有多少/哪些事实的证据，能证明那些负面的结果会发生？有没有其他的证据能证明有其他可能，如发生好的结果？当你罗列完这两部分内容，也许你会发现，让你焦虑的只是许多可能中的一种，且并没有压倒性的证据。\n\n不要让自己和自己的想法绑定。把你的想法看成许多经过你大脑的数据，你有力量选择去相信其中的一些、不相信其中的另一些；而不是照单全收。我们的祖先在野外为了生存，必须对危险和威胁有着过度高的警惕，我们的大脑沿袭了这一习惯，有时一些负面想法只是一种自动的条件反射。重要的是我们应该有选择地去相信自己的想法。\n\n2. 聚焦此刻\n\n你的头脑是不是始终在重复过去？曾经有些不好的经历发生过，并不意味着现在它们一定会再次发生。问问你自己，经过上次之后，你的应对能力、你的知识储备，以及情境本身是不是发生了变化？你已经不是过去那个你了。作为一个成年人， 你比青春期或儿时有了更多的选择，你更有能力去选择和哪些人来往，更有能力去识别、并主动离开一个坏的处境。\n\n你的大脑有时候会编造一些故事： 有的故事关于“你是谁”，有的关于“你是否安全”， 还有一些关于“你是否值得爱”等等。但并不是所有这些故事都是真实的。有时候，我们的大脑因为过去负面的经历变得充满偏见。重要的问题是：你在此刻的经历究竟是什么？它是真实已经在发生的么？还是只是可能会发生的？我们因为过去经历而变得偏见的大脑，可能会把这两种情况当作一种情况来处理，给你带来非常不好的感受，我们要做的就是尽可能明确意识到这两者不是一回事。\n\n3. 给你的想法打上标签\n\n给你的想法分类打上标签，而不是直接去关注想法的内容。观察自己的想法，当你注意到你开始判断的时候，（例如当你开始判断现在的处境多好或者多坏的时候），不要急于关注想法的内容，而是给它打上标签，告诉自己“我在（无端的）判断”。如果你注意到自己在担忧，（例如当你开始担心你会失败、或经历一种失去），给它打上标签“我在担忧”。如果你开始批评自己，打上标签“批评”。这个过程能够帮助你脱离想法的内容，避免被内容困住，而让你对“我在做什么”、以及“我为什么会有这些想法”的过程更有觉知。从而你也许可以意识到，是不是有其他方式来看待现在的处境（更客观平和的方式）。\n\n4. 想得更大更远\n\n对于你当下的处境，你是不是看得太过狭隘了？你是不是只看到了负面的部分，而没有看清楚事情的全貌？焦虑会让我们只能注意到眼前的威胁，而无法从更大更远的角度思考。 眼前这个情境，真的像你的焦虑告诉你的那样重要么？5年甚至10年以后，你还会如此在意眼前的这个问题么？\n\n5.先去干点儿别的吧\n\n你在为同一件事反复纠结，找不到解决的办法，这时，如果你还死盯着这件事，你就会陷入本文一开始所说的恶性循环。因此，不妨先开始做一些其他的任务，比如期末你写不出论文的时候，可以先看点儿闲书或者做个饭。\n\n不要以为这是在浪费你解决问题的时间，往往当你重新开始面对自己的问题时，你会发现自己有了不一样的感受。\n\n6. 看看这个想法是不是真的有用\n\n在第一点中，我们已经谈到，要注意辨识我们的想法是不是真实的。然而，并不是所有真实的想法对我们都是有用的。比如你现在要应聘，你知道你所应聘的工作要在10个人中录取1个人，因此，你只有1/10的录取率。这个想法是真实的，却对你没有帮助，也许会吓到你自己，让你连申请都交不上去。\n\n记住，把注意力集中在对你有用的想法上。\n\n7. 别太担心！适当的焦虑是有好处的\n\n研究表明，一定程度的焦虑能让人表现更好。这和一种叫做“不现实的乐观主义 (unrealistic optimism)”的心理现象有关。大多数人在他们的一生中都会有这样的倾向：无视负面的信息，而偏爱那些满足我们的反馈。我们通常会不理性地忽略那些会对我们造成负面作用的信息，而非常快乐地接受对我们某种程度上有好处的信息。一个新的、还未发表的研究 (by Tali Sharot at University College London) 指出，当人们焦虑时，这种“不现实的乐观主义”就会消失。他们变得能够客观地接受信息，从而导向更好的决策。也就是说，适当水平的焦虑能帮助你客观认识你的处境，你会能够同时看到事情好的一面和坏的一面，而不是只看到其中一面，从而做出有偏见的决策。\n\n最后要告诉大家的是，你们也许还不知道：如果你经常觉得焦虑，这是你智商高的表现。2015年的一个新研究发现了这样一种相关性：焦虑水平高的人在智商测试中表现更好，尤其是说话表达方面的智力。你觉得这个研究结果是不是准确可靠呢？\n\n著名德语诗人里尔克（Rainer Maria Rilke）说：“我们必须全力以赴，同时又不抱持任何希望。……不管做什么事，都要当它是全世界最重要的一件事，但同时又知道这件事根本无关紧要。”\n\n——这或许是我们能对抗焦虑的最本质的方法。",
			Author_Name: "宫丹",
			Created_At:  "2016-06-07",
			ImageURL:    "https://s11.ax1x.com/2023/12/21/piTw0sI.jpg",
		},
		{
			Title:       "7个方法帮你缓解焦虑症症状",
			Content:     "当社会普遍存在一种心理障碍，叫做“讨好型人格”。\n\n讨好型人格一般有以下几种特征：\n\n1、没有自信，容易受到他人的影响，特别是注重别人对自己的评价。\n\n2、没有主见，不知道自己想要什么，在生活和工作中常常犹豫不决、不知所措。\n\n3、心思敏感细腻，能察觉到别人的情绪变化以及需求，害怕被被人忽略或者让他人不高兴。\n\n4、虽然表面友善但是内心不愿和别人太过亲近，不愿意向他人袒露自己的需求与看法。\n\n5、当他人对自己寻求帮助时，往往不太懂得拒绝别人，自己常感无奈，却又无能为力。\n\n这一类人，为人处世的原则是“牺牲自己，成全别人”。总是无条件的满足别人的要求，从不拒绝；做事之前也常常思前想后，优先去考虑别人的想法，害怕与别人起冲突……\n\n也许，看到这里你已经对号入座，找到了自己的影子。",
			Author_Name: "525心理网",
			Created_At:  "2013-01-24",
			ImageURL:    "https://s11.ax1x.com/2023/12/21/piTwwQA.jpg",
		},
	}

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("transiction error")
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return
	}
	// 插入数据库
	for _, article := range articles {
		err := tx.Create(&article).Error
		if err != nil {
			fmt.Println(err)
			tx.Rollback() // 错误就回滚
			return
		}
	}
	// 提交事务
	err := tx.Commit().Error
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 初始化老师账号
func InitUsers() {
	users := []model.User{
		{
			Username: "admin",
			Status:   "教师",
			Online:   "0",
			Avator:   "https://s11.ax1x.com/2024/01/06/pizA82t.jpg",
		},
		{
			Username: "admin2",
			Status:   "教师",
			Online:   "0",
			Avator:   "https://s11.ax1x.com/2024/01/06/pizA1PA.jpg",
		},
	}

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("transiction error")
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return
	}
	// 插入数据库
	for _, user := range users {
		// 设置密码
		password := "123456"
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			// 处理错误，例如打印错误日志、返回错误等
			fmt.Printf("Error hashing password for user %s: %s\n", user.Username, err)
			return
		}

		user.Password = hashedPassword
		//fmt.Println("user: ", user)
		err = tx.Create(&user).Error
		if err != nil {
			fmt.Println(err)
			tx.Rollback() // 错误就回滚
			return
		}
	}
	// 提交事务
	err := tx.Commit().Error
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 初始化测试题目
func InitQuestion() {
	questions := []model.Question{
		{
			Id:           1,
			QuestionText: "在社交场合中，你更喜欢:",
		},
		{
			Id:           2,
			QuestionText: "当面对困难时，你的第一反应是:",
		},
		{
			Id:           3,
			QuestionText: "在工作中，你更倾向于:",
		},
		{
			Id:           4,
			QuestionText: "你如何处理失败:",
		},
		{
			Id:           5,
			QuestionText: "当面对陌生环境时，你的反应是:",
		},
		{
			Id:           6,
			QuestionText: "你更喜欢:",
		},
		{
			Id:           7,
			QuestionText: "在决策时，你更看重:",
		},
		{
			Id:           8,
			QuestionText: "你更容易受到哪种类型的影响:",
		},
		{
			Id:           9,
			QuestionText: "当周围环境变得嘈杂时，你的注意力更集中在:",
		},
		{
			Id:           10,
			QuestionText: "你如何处理与他人的冲突:",
		},
	}

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("transiction error")
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return
	}
	// 插入数据库
	for _, question := range questions {
		err := tx.Create(&question).Error
		if err != nil {
			fmt.Println(err)
			tx.Rollback() // 错误就回滚
			return
		}
	}
	// 提交事务
	err := tx.Commit().Error
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 初始化测试题目的答案
func InitAnswer() {
	answers := []model.Answer{
		// 第一题
		{
			QuestionID: 1,
			AnswerText: "保持低调，观察他人",
			Score:      3,
		},
		{
			QuestionID: 1,
			AnswerText: "与几个好友深度交流",
			Score:      2,
		},
		{
			QuestionID: 1,
			AnswerText: "大声热烈地参与",
			Score:      1,
		},
		// 第二题
		{
			QuestionID: 2,
			AnswerText: "迅速采取行动",
			Score:      3,
		},
		{
			QuestionID: 2,
			AnswerText: "仔细思考并寻求建议",
			Score:      2,
		},
		{
			QuestionID: 2,
			AnswerText: "感觉困惑，需要一些时间",
			Score:      1,
		},
		// 第三题
		{
			QuestionID: 3,
			AnswerText: "领导团队",
			Score:      3,
		},
		{
			QuestionID: 3,
			AnswerText: "与同事合作",
			Score:      2,
		},
		{
			QuestionID: 3,
			AnswerText: "独自完成任务",
			Score:      1,
		},
		// 第四题
		{
			QuestionID: 4,
			AnswerText: "分析失败的原因并学习",
			Score:      3,
		},
		{
			QuestionID: 4,
			AnswerText: "立即寻找下一步行动计划",
			Score:      2,
		},
		{
			QuestionID: 4,
			AnswerText: "感到沮丧，需要一些时间来恢复",
			Score:      1,
		},
		// 第五题
		{
			QuestionID: 5,
			AnswerText: "兴奋，愿意尝试新事物",
			Score:      3,
		},
		{
			QuestionID: 5,
			AnswerText: "谨慎，先观察再决定行动",
			Score:      2,
		},
		{
			QuestionID: 5,
			AnswerText: "紧张，不太愿意冒险",
			Score:      1,
		},
		// 第六题
		{
			QuestionID: 6,
			AnswerText: "制定详细计划并按计划行事",
			Score:      3,
		},
		{
			QuestionID: 6,
			AnswerText: "灵活适应，不拘泥于计划",
			Score:      2,
		},
		{
			QuestionID: 6,
			AnswerText: "随心而行，不做太多计划",
			Score:      1,
		},
		// 第七题
		{
			QuestionID: 7,
			AnswerText: "逻辑和事实",
			Score:      3,
		},
		{
			QuestionID: 7,
			AnswerText: "他人的意见和感受",
			Score:      2,
		},
		{
			QuestionID: 7,
			AnswerText: "直觉和个人价值观",
			Score:      1,
		},
		// 第八题
		{
			QuestionID: 8,
			AnswerText: "专业领域的专家",
			Score:      3,
		},
		{
			QuestionID: 8,
			AnswerText: "亲密关系的人",
			Score:      2,
		},
		{
			QuestionID: 8,
			AnswerText: "个人信仰和原则",
			Score:      1,
		},
		// 第九题
		{
			QuestionID: 9,
			AnswerText: "与他人的交流",
			Score:      3,
		},
		{
			QuestionID: 9,
			AnswerText: "完成手头任务",
			Score:      2,
		},
		{
			QuestionID: 9,
			AnswerText: "内心思考和感受",
			Score:      1,
		},
		// 第十题
		{
			QuestionID: 10,
			AnswerText: "直接解决问题，寻找共识",
			Score:      3,
		},
		{
			QuestionID: 10,
			AnswerText: "寻求妥协，保持和谐",
			Score:      2,
		},
		{
			QuestionID: 10,
			AnswerText: "避免冲突，退避求安宁",
			Score:      1,
		},
	}

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("transiction error")
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return
	}
	// 插入数据库
	for _, answer := range answers {
		err := tx.Create(&answer).Error
		if err != nil {
			fmt.Println(err)
			tx.Rollback() // 错误就回滚
			return
		}
	}
	// 提交事务
	err := tx.Commit().Error
	if err != nil {
		fmt.Println(err)
		return
	}
}
