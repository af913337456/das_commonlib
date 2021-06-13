package kafka

import (
	"github.com/DeAccountSystems/das_commonlib/common"
	"github.com/Shopify/sarama"
	"time"
)

/**
 * Copyright (C), 2019-2019
 * FileName: kafka
 * Author:   LinGuanHong
 * Date:     2019-10-18 14:47
 * Description:
 */

const kafkaLostConnectErrMsg = "client has run out of available brokers to talk to"
type RetryConfig struct {
	RetryTime   int
	DelayTime time.Duration
}
type KafkaMessageQueueProducer struct {
	SyncProducer *sarama.SyncProducer
	ReqTryCfg   RetryConfig
}

func NewDefaultKafkaMessageQueueProducer(brokersAddress []string, reqTryCfg RetryConfig) (*KafkaMessageQueueProducer, error) {
	config := sarama.NewConfig()
	config.Producer.MaxMessageBytes = 100000000
	config.Admin.Timeout = 20 * time.Second
	config.Net.DialTimeout = 10 * time.Second
	config.Net.ReadTimeout = 15 * time.Second
	config.Net.WriteTimeout = 15 * time.Second
	config.ChannelBufferSize = 512
	config.Producer.Flush.Frequency = 50 * time.Millisecond
	config.Producer.Flush.Messages = 200
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokersAddress, config)
	if err != nil {
		return nil, err
	}
	msgQueue := &KafkaMessageQueueProducer{
		SyncProducer: &producer,
		ReqTryCfg:    reqTryCfg,
	}
	return msgQueue, nil
}

// single send
func (k *KafkaMessageQueueProducer) SendJsonMessage(topic string, jsonByte []byte) error {
	msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.ByteEncoder(jsonByte)}
	_, err := common.RetryReq(k.ReqTryCfg.RetryTime, k.ReqTryCfg.DelayTime, func() (interface{}, error) {
		if k.SyncProducer == nil {
			return nil, nil
		}
		if _, _, err := (*k.SyncProducer).SendMessage(msg); err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}

// batch send
func (k *KafkaMessageQueueProducer) SendJsonMessages(topic string, jsonBytes ...[]byte) error {
	size := len(jsonBytes)
	msgs := make([]*sarama.ProducerMessage, 0, size)
	for i := 0; i < size; i++ {
		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.ByteEncoder(jsonBytes[i])}
		msgs = append(msgs, msg)
	}
	_, err := common.RetryReq(k.ReqTryCfg.RetryTime, k.ReqTryCfg.DelayTime, func() (interface{}, error) {
		if err := (*k.SyncProducer).SendMessages(msgs); err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}

func (k *KafkaMessageQueueProducer) CloseKafkaProducer() {
	if (*k.SyncProducer) != nil {
		_ = (*k.SyncProducer).Close()
		*k.SyncProducer = nil
	}
}

func (k *KafkaMessageQueueProducer) Close() {
	k.CloseKafkaProducer()
}