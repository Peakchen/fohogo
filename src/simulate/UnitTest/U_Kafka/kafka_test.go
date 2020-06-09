package U_Kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

func TestSyncKafka(t *testing.T) {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	var wg sync.WaitGroup
	wg.Add(2)
	go syncProducer()
	go consumer()
	wg.Wait()
}

func syncProducer() (succ bool) {
	akLog.FmtPrintln("begin syncProducer...")
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

func TestASyncKafka(t *testing.T) {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	var wg sync.WaitGroup
	wg.Add(2)
	go asyncProducer()
	go asyncConsumer()
	wg.Wait()
}

func asyncProducer() (succ bool) {
	akLog.FmtPrintln("begin asyncProducer...")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse                          // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy                    // Compress messages
	config.Producer.Flush.Frequency = time.Duration(10000) * time.Millisecond // Flush batches every 500ms
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	producer, err := sarama.NewAsyncProducer([]string{host}, config)
	if err != nil {
		akLog.FmtPrintln("create producer err: ", err)
		return
	}
	defer producer.AsyncClose()

	for i := 0; i < 10; i++ {
		akLog.FmtPrintln("produce msg... num: ", i)
		//构建发送的消息
		msg := &sarama.ProducerMessage{
			Topic:     "test",                      //包含了消息的主题
			Partition: int32(10),                   //
			Key:       sarama.StringEncoder("key"), //
		}
		msg.Value = sarama.StringEncoder(fmt.Sprintf("this is a good test, hello kafka, num: %v.", i))
		producer.Input() <- msg
		//time.Sleep(1 * time.Second)
	}
	succ = true
	akLog.FmtPrintln("end asyncProducer...")
	return
}

func asyncConsumer() (succ bool) {
	akLog.FmtPrintln("begin asyncConsumer...")
	// 根据给定的代理地址和配置创建一个消费者
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.CommitInterval = 10
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 2
	config.Version = sarama.V2_5_0_0
	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup([]string{host}, "test-consumer-group", config)
	if err != nil {
		akLog.FmtPrintln("create ConsumerGroup err: ", err)
		return
	}

	consumer := &AsyncConsumer{}
	go func() {
		for {
			err := client.Consume(ctx, []string{"test"}, consumer)
			if err != nil {
				akLog.FmtPrintln("client.Consume error=[%v]", err.Error())
				// 5秒后重试
				time.Sleep(time.Second * 5)
			}
		}
	}()

	// os signal
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
	cancel()
	err = client.Close()
	if err != nil {
		panic(err)
	}
	succ = true
	akLog.FmtPrintln("end asyncConsumer...")
	return
}

type AsyncConsumer struct {
}

func (consumer *AsyncConsumer) Setup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *AsyncConsumer) Cleanup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *AsyncConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		key := string(message.Key)
		val := string(message.Value)
		akLog.FmtPrintf("key:%s, val:%s.\n", key, val)
		session.MarkMessage(message, "")
	}

	return nil
}
