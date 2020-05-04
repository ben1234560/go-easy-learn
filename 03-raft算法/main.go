package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// README：1.实现：每个节点都可以成为候选者，谁先成为其它节点则不能成为
// 		   2.每个节点能投一票
// 		   3.选举出一个leader节点
// 判断是否有节点先选举了
var electCheck = false
// 实现3个节点分布式一致性
type Leader struct {
	// 任期
	Term int
	// 领导编号
	LeaderId int
}

// 定义结构体，后续调用作为存储，命名需要遵循raft规范
type Raft struct {
	// 锁
	mu sync.Mutex
	// 节点编号：0.1.2节点
	me int
	// 当前任期
	currentTerm int
	// 为哪个节点投票
	voteFor int
	// 节点状态：0 Follower、1 Candidate、2Leader
	state int
	// 最后一条信息的时间
	lastMessageTime int64
	// 当前的领导节点
	currentLeader int
	// 节点直接发消息的管道
	message chan bool
	// 选举通道
	electCh chan bool
	// 心跳信号通道
	heartBeat chan bool
	// 返回心跳信号的通道
	heartbeatRe chan bool
	// 超时时间
	timeout int
}

// 定义节点个数，模拟3个节点
const raftCount = 3

func main() {
	// 循环创建节点（目前设置的是3个），一开始都是follower状态
	for i := 0; i < raftCount; i++ {
		// 创建节点
		Make(i)
	}
	// 夯住等跑完
	select {}
}

// 创建节点的方法
// Make函数名，me int传入参数及类型，*Raft返回的具体对象
func Make(me int) *Raft {
	// 构建节点
	rf := &Raft{}
	// 赋予编号：0/1/2
	rf.me = me
	// 表示还没投票
	rf.voteFor = -1
	// 目前状态follower
	rf.state = 0
	// 设置为随机数，这里先放0
	rf.timeout = 0
	// 目前没有领导
	rf.currentLeader = -1
	rf.currentTerm = 0
	// 4个管道分别初始化
	rf.electCh = make(chan bool)
	rf.message = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbeatRe = make(chan bool)
	// 这样三个节点都初始化好了，但还没有投票和选举的功能
	// 1.实现选举投票
	go rf.election()
	// 2.实现心跳检查
	go rf.sendHeartBeat()
	return rf
}

// 具体的选举函数
// 绑定结构体
func (rf *Raft) election() {
	var result bool
	// 没选出来之前一直在跑，不断选主
	for {
		// 设置随机的缓存时间
		rand.Seed(time.Now().UnixNano())
		// 产生一个150-300的值
		timeout := rand.Int63n(300-150) + 150
		// 设置最后一条信息时间
		rf.lastMessageTime = millisecond()
		// 目前节点状态是follower
		select {
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前节点状态为：", rf.state)
		}
		result = false
		// 进行选主
		for !result {
			// 当有节点为主时，别的节点不再选
			var leader = Leader{0, -1}
			result = rf.electionOne(&leader)
		}
	}
}

// 具体选出一个主
func (rf *Raft) electionOne(leader *Leader) bool {
	if leader.LeaderId > -1 && leader.LeaderId != rf.me {
		fmt.Printf("%d已是Leader, 终止%d选举\n", leader.LeaderId, rf.me)
		return true
	}
	// 定义超时，超了时间重新选
	var timeout int64 = 100
	// 定义投票的数量
	var vote int
	// 心跳信号是否开始产生
	var triggerHeartBeat bool
	last := millisecond()
	// 默认没选出来，目前是false
	success := false
	// 判断是否有节点先成为候选者
	if electCheck == true {
		fmt.Printf("已有节点先成为候选者，当前节点%d停止选举\n", rf.me)
		return true
	}
	// 更新节点为候选者，并为自己投票
	rf.mu.Lock()
	rf.state = 1
	rf.voteFor = rf.me
	rf.currentLeader = -1
	electCheck = true
	rf.mu.Unlock()
	// 开始选举
	fmt.Println("开始选举主节点")
	triggerHeartBeat = false
	// 拉选票等等
	for {
		if leader.LeaderId > -1 && leader.LeaderId != rf.me {
			fmt.Printf("%d已是Leader, 终止%d选举\n", leader.LeaderId, rf.me)
			return true
		}

		// 不断的拉票
		for i := 0; i < raftCount; i++ {
			// 拉票，判断如果遍历的节点不是自己，就拉票
			if i != rf.me {
				go func() {
					// 判断下是否有领导
					if leader.LeaderId < 0 {
						// 往管道添加数据，证明投票了
						rf.electCh <- true

					}
				}()
			}
		}
		// 统计投票
		// 判断选票数量是否大于节点个数/2，大于则成为主
		vote = 0
		for i := 0; i < raftCount; i++ {
			// 计算投票数量
			select {
			case ok := <-rf.electCh:
				// 如果接过来的数据为true，增加票数并判断是否选举成功
				if ok {
					vote++
					success = vote > raftCount/2
					if success && !triggerHeartBeat {
						// 开始心跳检查
						triggerHeartBeat = true
						// 选举成功了，更新节点为Leader状态
						rf.mu.Lock()
						rf.state = 2
						rf.currentLeader = rf.me
						leader.LeaderId = rf.me
						rf.currentTerm = rf.currentTerm + 1
						rf.mu.Unlock()
						// 发信号通知
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为了Leader")
						fmt.Println("开始发送心跳信号")
					}
				}
			}
		}
		// 最终的判断：没超时且投票数大于且leader值没问题
		if timeout+last < millisecond() || (vote >= raftCount/2 || rf.currentLeader > -1) {
			break
		} else {
			// 延时操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

// 获取当前时间的毫秒数
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 心跳函数
func (rf *Raft) sendHeartBeat() {
	// 循环操作
	for {
		select {
		case <-rf.heartBeat:
			// 若是主，则发起心跳检查
			if rf.currentLeader == rf.me {
				// 此时是主节点，给所有小弟发信息
				for i := 0; i < raftCount; i++ {
					if i != rf.me {
						go func() {
							rf.heartbeatRe <- true
						}()
					}
				}
				// 统计返回的节点个数
				// 记录有几个小弟回复了
				count := 0
				for i := 0; i < raftCount; i++ {
					select {
					case ok := <-rf.heartbeatRe:
						if ok {
							count++
							if count > raftCount/2 {
								fmt.Println("心跳检查成功")
								log.Fatal("程序结束")
							}
						}
					}
				}
			}
		}
	}
}
