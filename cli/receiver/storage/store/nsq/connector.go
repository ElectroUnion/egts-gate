package nsq

/*
Плагин для работы с NSQ.
Плагин отправляет пакет в топик NSQ messaging system.

Раздел настроек, которые должны отвечають в конфиге для подключения хранилища:

servers = "host:port"
topic = "receiver"
*/

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type Connector struct {
	producer   *nsq.Producer
	config     map[string]string
}

func (c *Connector) Init(cfg map[string]string) error {
	var (
		err error
	)
	if cfg == nil {
		return fmt.Errorf("Не корректная ссылка на конфигурацию")
	}
	c.config = cfg

	fmt.Println("Connect to NSQ server:", c.config["server"])

	config := nsq.NewConfig()

	if c.producer, err = nsq.NewProducer(c.config["server"], config); err != nil {
		return fmt.Errorf("Ошибка подключения к nsq шине: %v", err)
	}
	return err
}

func (c *Connector) Save(msg interface{ ToBytes() ([]byte, error) }) error {
	if msg == nil {
		return fmt.Errorf("Не корректная ссылка на пакет")
	}

	innerPkg, err := msg.ToBytes()
	if err != nil {
		return fmt.Errorf("Ошибка сериализации  пакета: %v", err)
	}

	if err = c.producer.Publish(c.config["topic"], innerPkg); err != nil {
		return fmt.Errorf("Не удалось отправить сообщение в топик: %v", err)
	}
	return nil
}

func (c *Connector) Close() error {
	c.producer.Stop()
	return nil
}
