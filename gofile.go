package main

import (
	"errors"
)

type Observer interface {
	Update(string)
}

type Subject interface {
	Attach(o Observer) (bool, error)
	Detach(o Observer) (bool, error)
	Notify() (bool, error)
}

type NewsObserver struct {
	name string
}

func (s *NewsObserver) Update(update string) {
	println(s.name + ": " + update + "\n")
}

type NewsMonitor struct {
	reporter string
	news     string

	observers []Observer
}

func (n *NewsMonitor) Attach(o Observer) (bool, error) {
	for _, observer := range n.observers {
		if observer == o {
			return false, errors.New("Observer already exists")
		}
	}

	n.observers = append(n.observers, o)

	return true, nil
}

func (n *NewsMonitor) Detach(o Observer) (bool, error) {

	for i, observer := range n.observers {
		if observer == o {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Observer not found")
}

func (n *NewsMonitor) String() string {
	return "Current new from " + n.reporter + ": " + n.news
}

func (n *NewsMonitor) Notify() (bool, error) {
	for _, observer := range n.observers {
		observer.Update(n.String())
	}
	return true, nil
}

func (n *NewsMonitor) SetNews(reporter string, news string) {
	n.reporter = reporter
	n.news = news
	n.Notify()
}

func main() {
	newsMonitor := &NewsMonitor{
		reporter: "BBC",
		news:     "The new king of England is Charles III",
	}

	observer1 := &NewsObserver{
		name: "Daniyar",
	}

	observer2 := &NewsObserver{
		name: "Chapagan",
	}

	newsMonitor.Attach(observer1) //subscribed on news
	newsMonitor.Attach(observer2)

	newsMonitor.Notify()

	newsMonitor.SetNews("Yandex Zen", "Adobe has bought Figma")

	newsMonitor.Detach(observer2) //removed subscription so only Daniar will be notified

	newsMonitor.Notify()
}
