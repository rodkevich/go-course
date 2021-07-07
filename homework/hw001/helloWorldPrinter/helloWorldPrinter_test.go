package helloWorldPrinter

import (
	"fmt"
	"github.com/kyokomi/emoji"
	"testing"
)

//const LIMIT int = 1000

//func benchmarkAddEmoji(w  func(a ...interface{}) (int, error), b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < LIMIT; j++ {
//			AddEmoji(`:relaxed:`, w)
//		}
//	}
//	b.ReportAllocs()
//}

func benchmarkAddEmoji(w func(a ...interface{}) (int, error), b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddEmoji(`:relaxed:`, w)
	}
	//b.ReportAllocs()
}

func BenchmarkAddEmoji1(b *testing.B) {
	benchmarkAddEmoji(emoji.Println, b)
}

func BenchmarkAddEmoji2(b *testing.B) {
	benchmarkAddEmoji(fmt.Println, b)
}

//func BenchmarkAddEmoji2(b *testing.B) {
//	var w = emoji.Println
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < LIMIT; j++ {
//			AddEmoji2(`:relaxed:`, w)
//		}
//	}
//	b.ReportAllocs()
//}

//func BenchmarkExamples(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < LIMIT; j++ {
//			fmt.Println("Hello World Emoji!")
//			//emoji.Println(":beer: Beer!!!")
//			//pizzaMessage := emoji.Sprint("I like a :pizza: and :sushi:!!")
//			//fmt.Println(pizzaMessage)
//		}
//	}
//	b.ReportAllocs()
//}
