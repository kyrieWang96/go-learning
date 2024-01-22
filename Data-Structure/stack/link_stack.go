package main

import "fmt"

type LinkList struct {
	value any
	Next  *LinkList
}

type LinkListStak struct {
	length   int
	LinkList *LinkList
}

func NewLinkListStak() *LinkListStak {
	return &LinkListStak{}
}

func (s *LinkListStak) IsEmpty() bool {
	return s.LinkList == nil
}

func (s *LinkListStak) Push(value any) {
	linkList := &LinkList{
		value: value,
	}
	if s.LinkList == nil {
		s.LinkList = linkList
	} else {
		linkList.Next, s.LinkList = s.LinkList, linkList
	}

	s.length++
}

func (s *LinkListStak) Pop() any {
	if s.IsEmpty() {
		return nil
	}

	value := s.LinkList.value
	s.LinkList = s.LinkList.Next
	s.length--

	return value
}

func (s *LinkListStak) Peek() any {
	if s.IsEmpty() {
		return nil
	}

	return s.LinkList.value
}

func (s *LinkListStak) Print() {
	if s.IsEmpty() {
		return
	}

	for link := s.LinkList; link != nil; {
		fmt.Println(link.value)
		link = link.Next
	}
}
