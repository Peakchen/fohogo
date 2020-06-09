package U_Kafka

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Shopify/sarama"
)

/*
	需要开启zookeeper和kafka进程
*/

var (
	host = "192.168.126.128:9092"
)

func TestKafka(t *testing.T) {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer()
	go consumer()
	wg.Wait()
}

func producer() (succ bool) {
	akLog.FmtPrintln("begin producer...")
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{host}, config)
	if err != nil {
		akLog.FmtPrintln("create producer err: ", err)
		return
	}
	defer producer.Close()
	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic:     "test",                      //包含了消息的主题
		Partition: int32(10),                   //
		Key:       sarama.StringEncoder("key"), //
	}
	for i := 0; i < 10; i++ {
		msg.Value = sarama.StringEncoder("this is a good test, hello kafka.")
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			akLog.FmtPrintln("Send message Fail")
			return
		}
		akLog.FmtPrintf("Partition = %d, offset=%d\n", partition, offset)
		time.Sleep(2 * time.Second)
	}
	succ = true
	return
}

func consumer() (succ bool) {
	akLog.FmtPrintln("begin consumer...")
	var (
		wg sync.WaitGroup
	)
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer([]string{host}, nil)
	if err != nil {
		akLog.FmtPrintln("create consumer err: ", err)
		return
	}
	defer consumer.Close()
	//Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions("test")
	if err != nil {
		akLog.FmtPrintln("get Partitions err: ", err)
		return
	}

	for partition := range partitionList {
		//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		//如果该分区消费者已经消费了该信息将会返回error
		//sarama.OffsetNewest:表明了为最新消息
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			akLog.FmtPrintln("get consumer Partitions err: ", err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	succ = true
	return
}
