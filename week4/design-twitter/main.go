/*
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-twitter
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "github.com/vrealzhou/geekbang_algorithm/internal/data"

type Twitter struct {
	tweets  map[int]*tweet
	follows map[int]set // key follower, value followee
	count   int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		tweets:  make(map[int]*tweet),
		follows: make(map[int]set),
	}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.checkNewUser(userId)
	t.count++
	latest := &tweet{
		id:   tweetId,
		ts:   t.count,
		Next: t.tweets[userId],
	}
	t.tweets[userId] = latest
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	t.checkNewUser(userId)
	// fmt.Printf("***GetNewsFeed user %d, %v\n", userId, this.follows[userId])
	result := make([]int, 10)
	userCount := len(t.follows[userId]) + 1
	// init binary heap
	h := data.NewHeap(userCount, func(source, target interface{}) int {
		return source.(*tweet).ts - target.(*tweet).ts
	})
	if t.tweets[userId].Next != nil {
		h.Add(t.tweets[userId])
	}
	for followeeId := range t.follows[userId] {
		item := t.tweets[followeeId]
		if item.Next != nil {
			h.Add(item)
		}
	}
	i := 0
	for ; i < 10; i++ {
		top := h.Pop()
		if top == nil {
			break
		}
		item := top.(*tweet)
		result[i] = item.id
		if item.Next != nil && item.Next.Next != nil {
			h.Add(item.Next)
		}
	}
	return result[:i]
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	t.checkNewUser(followerId)
	t.checkNewUser(followeeId)
	t.follows[followerId].Add(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	t.checkNewUser(followerId)
	t.checkNewUser(followeeId)
	t.follows[followerId].Remove(followeeId)
}

func (t *Twitter) checkNewUser(userId int) {
	if _, ok := t.tweets[userId]; !ok {
		t.tweets[userId] = &tweet{
			id: -1,
		}
		t.follows[userId] = make(set)
	}
}

type tweet struct {
	id   int
	ts   int
	Next *tweet
}

type set map[int]struct{}

func (s set) Add(value int) {
	s[value] = struct{}{}
}
func (s set) Remove(value int) {
	delete(s, value)
}
