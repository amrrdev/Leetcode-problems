1type Foo struct {
2	channelOne chan struct{}
3	channelTwo chan struct{}
4}
5
6func NewFoo() *Foo {
7	return &Foo{
8		channelOne: make(chan struct{}),
9		channelTwo: make(chan struct{}),
10	}
11}
12
13func (f *Foo) First(printFirst func()) {
14	printFirst()
15	f.channelOne <- struct{}{}
16}
17
18func (f *Foo) Second(printSecond func()) {
19	<-f.channelOne
20	printSecond()
21	f.channelTwo <- struct{}{}
22
23}
24
25func (f *Foo) Third(printThird func()) {
26	<-f.channelTwo
27	printThird()
28}
29
30