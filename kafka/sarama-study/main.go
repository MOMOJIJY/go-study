package main

const (
	topic = "demo-topic"
	broker = "127.0.0.1:9092"
	group = "my-group"
)

var (
	mode    string
	action  string
	ktopic  string
	kconfig string
)

func main() {
	idempotence()


	//fmode := flag.String("mode", "producer", "kafka mode")
	//faction := flag.String("action", "list", "kafka action")
	//ftopic := flag.String("topic", "", "kafka topic")
	//flag.Parse()
	//mode = *fmode
	//action = *faction
	//ktopic = *ftopic
	//
	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	//
	//switch mode {
	//case "producer":
	//	go asyncProduce()
	//case "consumer":
	//	//go asyncConsume()
	//	go consumeGroup()
	//case "metadata":
	//	metadata()
	//case "admin":
	//	NewAdmin()
	//default:
	//	panic("invalid mode")
	//}
	//
	//<-ch
}

