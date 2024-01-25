package main

import (
	"context"
	"flag"
	"fmt"
	log2 "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/OpenIMSDK/tools/log"
	"github.com/openimsdk/openim-sdk-core/v3/msgtest/module"
)

func init() {
	//_ = runtime.GOMAXPROCS(7)
	InitWithFlag()
	if err := log.InitFromConfig("sdk.log", "sdk", 3,
		true, false, "./", 2, 24); err != nil {
		panic(err)
	}
}

var (
	totalOnlineUserNum    int     // 总在线用户数
	randomSender          int     // 随机发送者数
	randomReceiver        int     // 随机接收者数
	singleSamplingRate    float64 // 单聊抽样率
	GroupSenderRate       float64 // 群聊随机的发送者比例
	NotFriendMsgSenderNum int     // 非好友消息发送者数
	msgSenderNumEveryUser int     // 每个用户的消息数
	start                 int
	end                   int
	count                 int
	sendInterval          int

	//recvMsgUserNum int // 消息接收者数, 抽样账号
	isRegisterUser  bool // 是否注册用户
	onlineUsersOnly bool

	hundredThousandGroupNum int //
	tenThousandGroupNum     int
	thousandGroupNum        int
	hundredGroupNum         int
	fiftyGroupNum           int
	tenGroupNum             int
)

func InitWithFlag() {
	flag.IntVar(&totalOnlineUserNum, "o", 20000, "total online user num")
	flag.IntVar(&randomSender, "rs", 0, "random sender num")
	flag.IntVar(&randomReceiver, "rr", 0, "random receiver num")
	flag.IntVar(&start, "s", 0, "start user")

	flag.IntVar(&end, "e", 0, "end user")
	flag.Float64Var(&singleSamplingRate, "sr", 0.01, "single chat sampling rate")
	flag.Float64Var(&GroupSenderRate, "gsr", 0.1, "group chat sender rate")
	flag.IntVar(&count, "c", 200, "number of messages per user")
	flag.IntVar(&sendInterval, "i", 1000, "send message interval per user(milliseconds)")
	flag.IntVar(&NotFriendMsgSenderNum, "n", 100, "not friend msg sender num")
	flag.IntVar(&hundredThousandGroupNum, "htg", 0, "quantity of 100k user groups")
	flag.IntVar(&tenThousandGroupNum, "ttg", 0, "quantity of 10k user groups")
	flag.IntVar(&thousandGroupNum, "otg", 0, "quantity of 1k user groups")
	flag.IntVar(&hundredGroupNum, "hog", 0, "quantity of 100 user groups")
	flag.IntVar(&fiftyGroupNum, "fog", 0, "quantity of 50 user groups")
	flag.IntVar(&tenGroupNum, "teg", 0, "quantity of 10 user groups")

	flag.BoolVar(&isRegisterUser, "r", false, "register user to IM system")
	flag.BoolVar(&onlineUsersOnly, "u", false, "consider only online users")
}

func PrintQPS() {
	for {

		log.ZError(context.Background(), "QPS", nil, "qps", module.GetQPS())
		time.Sleep(time.Second * 1)
	}
}

func main() {
	flag.Parse()
	ctx := context.Background()
	fmt.Println("1111:::", onlineUsersOnly)
	log.ZWarn(ctx, "flag args", nil, "totalOnlineUserNum", totalOnlineUserNum,
		"randomSender", randomSender, "randomReceiver", randomReceiver,
		"singleSamplingRate", singleSamplingRate, "start", start, "end", end, "count", count, "sendInterval", sendInterval,
		"onlineUsersOnly", onlineUsersOnly, "isRegisterUser", isRegisterUser, "groupSenderRate", GroupSenderRate,
		"hundredThousandGroupNum", hundredThousandGroupNum, "tenThousandGroupNum", tenThousandGroupNum, "thousandGroupNum", thousandGroupNum,
		"hundredGroupNum", hundredGroupNum, "fiftyGroupNum", fiftyGroupNum, "tenGroupNum", tenGroupNum)

	go func() {
		log2.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	p := module.NewPressureTester()
	var f, r []string
	var err error
	if start != 0 {
		f, r, err = p.SelectSampleFromStarEnd(start, end, singleSamplingRate)
	} else {
		f, r, err = p.SelectSample(totalOnlineUserNum, singleSamplingRate)
	}
	if err != nil {
		log.ZError(ctx, "Sample UserID failed", err)
		return
	}
	p.CreateTestGroups(f, totalOnlineUserNum, GroupSenderRate, hundredThousandGroupNum,
		tenThousandGroupNum, thousandGroupNum, hundredGroupNum, fiftyGroupNum, tenGroupNum)

	log.ZWarn(ctx, "Sample UserID", nil, "sampleUserLength", len(r), "sampleUserID", r, "length", len(f))
	p.FormatGroupInfo(ctx)
	time.Sleep(10 * time.Second)
	//
	if isRegisterUser {
		if err := p.RegisterUsers(f, nil, nil); err != nil {
			log.ZError(ctx, "Sample UserID failed", err)
			return
		}
	}

	//go PrintQPS()
	// init users
	p.InitUserConns(f)
	log.ZWarn(ctx, "all user init connect to server success,start send message", nil, "count", count)
	if onlineUsersOnly {
		log.ZWarn(ctx, "OnlineUsersOnly do not send messages blocking the process...", nil)
		// Create a channel to receive operating system interrupt signals
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

		// Block the process until an interrupt signal is received
		<-signalChannel
		log.ZWarn(ctx, "OnlineUsersOnly do not send messages received interrupt signal. Exiting...", nil)
		return
	}

	time.Sleep(10 * time.Second)
	p.SendSingleMessages(ctx, f, p.Shuffle(f, randomSender), randomReceiver, count, time.Millisecond*time.Duration(sendInterval))
	p.SendGroupMessage(ctx, count, time.Millisecond*time.Duration(sendInterval))
	log.ZWarn(ctx, "send all message over", nil, "singleNum", p.GetSingleSendNum())
	//p.SendSingleMessagesTo(f, 20000, time.Millisecond*1)
	//p.SendMessages("fastened_user_prefix_testv3new_0", "fastened_user_prefix_testv3new_1", 100000)
	time.Sleep(1 * time.Minute)
	p.CheckMsg(ctx)

	time.Sleep(time.Hour * 60)

}
