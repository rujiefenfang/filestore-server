package main

func main() {

}

//type LoadBalance struct {
//	Client []*Client
//	Size   int
//	NowId  int
//}
//
//type Client struct {
//	Name int
//}
//
//func (l *LoadBalance) Get() *Client {
//	client := l.Client[l.NowId]
//	l.NowId++
//	l.NowId %= l.Size
//	return client
//}
//func (c *Client) print() {
//	fmt.Println(c.Name)
//}
//func InitLoadBalance(size int) *LoadBalance {
//	clients := make([]*Client, size)
//	for i := 0; i < size; i++ {
//		clients[i] = &Client{
//			Name: i,
//		}
//	}
//	return &LoadBalance{
//		Client: clients,
//		Size:   size,
//		NowId:  0,
//	}
//
//}
//
//var (
//	path []string
//	res  [][]string
//)
//
//func letter(maps map[int][]string, s string) {
//	back(maps, s, 0)
//}
//
//func back(maps map[int][]string, s string, length int) {
//	if length == len(s) {
//		temp := make([]string, length)
//		copy(temp, path)
//		res = append(res, temp)
//		return
//	}
//	for i := length; i < len(s); i++ {
//		num, _ := strconv.Atoi(string(s[i]))
//		for _, v := range maps[num] {
//			path = append(path, v)
//		}
//	}
//}
