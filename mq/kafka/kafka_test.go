package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"runtime"
	"testing"
	"time"
)

/**
 * Copyright (C), 2019-2020
 * FileName: lib_test
 * Author:   LinGuanHong
 * Date:     2020/12/15 11:20 上午
 * Description:
 */

func Test_kafkaConsumer(t *testing.T) {
	runtime.GOMAXPROCS(4)
	ctx, cancel := context.WithCancel(context.Background())
	consumer, err := NewDefaultKafkaMessageQueueConsumer([]string{"127.0.0.1:9092"}, "2",false, ctx)
	if err != nil {
		panic(err.Error())
	}
	go func() {
		time.Sleep(time.Second * 2)
		cancel()
		fmt.Println("start close")
		if err := consumer.Close(); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("close success")
		}
	}()
	consumer.ConsumeWithHandCommit([]string{"test"}, func(msg *sarama.ConsumerMessage) bool {
		msgStr := fmt.Sprintf("Message claimed: value = %s, timestamp = %v, topic = %s", string(msg.Value), msg.Timestamp, msg.Topic)
		fmt.Println(msgStr)
		time.Sleep(time.Second * 2)
		fmt.Println("finish one")
		return true // commit it
	})
	select {}
}
