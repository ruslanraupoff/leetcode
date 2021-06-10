package main

type MyCalendar struct {
	calendar map[int]int
}

func Constructor() MyCalendar {
	calendar := make(map[int]int)
	return MyCalendar{calendar: calendar}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for s, e := range this.calendar {
		if (s <= start && start < e) || (start < s && s < end) {
			return false
		}
	}
	this.calendar[start] = end
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
