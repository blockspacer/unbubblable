package connect

import (
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"github.com/kutuluk/unbubblable/server/protocol"
)

// SenderQueueLength определяет размер очереди сендера
const SenderQueueLength = 10

const (
	// Таймаут записи
	writeWait = 1 * time.Second
	// Таймаут чтения
	pongWait = 1 * time.Second
	// Задержка между пингами
	pingPeriod = 250 * time.Millisecond
)

// sender возвращает буферизованный канал и запускает горутину, отправляющую в сеть пришедшие
// в этот канал сообщения. Задача этой горутины - обеспечить запись в сокет из единственного
// конкурентного потока
func (c *Connect) sender() chan<- []byte {
	outbound := make(chan []byte, SenderQueueLength)
	pinger := time.NewTicker(pingPeriod)
	go func() {
		defer func() {
			pinger.Stop()
			//			c.close()
			log.Println("[sender]: сендер завершился")
		}()

		//		for message := range outbound {
		for {

			select {
			case message, ok := <-outbound:
				c.ws.SetWriteDeadline(time.Now().Add(writeWait))
				if !ok {
					// Канал закрыт
					c.ws.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}

				// Отправляем сообщение
				err := c.ws.WriteMessage(websocket.BinaryMessage, message)
				if err != nil {
					log.Println("[sender]:", err)
					//ToDo: сделать полноценную обработку ошибки
					if err == websocket.ErrCloseSent {
						// Сокет закрыт - выход из горутины
						// ToDo: нужно ли предпринимать еще какие-то действия, или коннект
						// все равно закроется в ресивере, поскольку соединение закрыто?
						return
					}
					// В этом месте сообщение теряется
					// ToDo: возможно стоит попытаться отправить сообщение снова?
					// ToDo: либо сразу закрывать соединение? Могут ли тут быть не критичные ошибки?
					break
				}
				// Увеличиваем счетчик отправленных байт
				c.Sent += len(message)

			case <-pinger.C:
				c.update()
			}
		}
	}()
	return outbound
}

// Send сериализует сообщение и отправляет его в исходящий канал
func (c *Connect) Send(msgType protocol.MessageType, msgBody []byte) {

	message := &protocol.Message{
		Type: msgType,
		Body: msgBody,
	}

	buffer, err := proto.Marshal(message)
	if err != nil {
		log.Println("[proto send]:", err)
		return
	}

	//	if c.state != StateCLOSED {
	// ToDo: возможна ситуация, когда канал уже закрыт
	// Разобраться, как обрабатывать такую ситуацию
	select {
	case c.outbound <- buffer:
	default:
		// Буфер переполнен - в этом месте сообщение теряется
		log.Println("[connect]: очередь на отправку переполнена")
	}
	//	}
}

// SendChain упаковывает данные в цепочку из одного сообщения и отправляет его
func (c *Connect) SendChain(msgType protocol.MessageType, msgBody []byte) {

	// Формируем сообщение
	message := &protocol.Message{
		Type: msgType,
		Body: msgBody,
	}

	// Упаковываем сообщение в цепочку
	//	chain := new(protocol.MessageChain)
	//	chain.Chain = append(chain.Chain, message)
	chain := &protocol.MessageChain{
		Chain: []*protocol.Message{message},
	}

	// Сериализуем сообщение протобафом
	buffer, err := proto.Marshal(chain)
	if err != nil {
		log.Println("[proto send]:", err)
		return
	}

	// Отправляем сообщение
	c.Send(protocol.MessageType_MsgChain, buffer)

}

// SendSystemMessage отправляет клиенту системное сообщение
func (c *Connect) SendSystemMessage(level int, text string) {

	message := &protocol.SystemMessage{
		Level: int32(level),
		Text:  text,
	}

	buffer, err := proto.Marshal(message)
	if err != nil {
		log.Println("[proto send]:", err)
		return
	}

	c.Send(protocol.MessageType_MsgSystemMessage, buffer)

}

// SendInfo отправляет клиенту служебную информацию
func (c *Connect) SendInfo() {

	message := &protocol.Info{
		Ping: int32(c.ping.median),
	}

	buffer, err := proto.Marshal(message)
	if err != nil {
		log.Println("[proto send]:", err)
		return
	}

	c.Send(protocol.MessageType_MsgInfo, buffer)

}