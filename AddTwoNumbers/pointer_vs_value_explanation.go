package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func demonstratePointerVsValue() {
	// Создаём цепочку узлов
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	node1.Next = node2
	node2.Next = node3

	// Создаём указатель curr
	curr := node1

	fmt.Println("=== ДЕМОНСТРАЦИЯ: Адрес переменной vs Значение переменной ===")
	fmt.Printf("Адрес переменной curr (где она хранится): %p\n", &curr)
	fmt.Printf("Значение переменной curr (на что она указывает): %p\n", curr)
	fmt.Printf("Значение в узле: %d\n", curr.Val)
	fmt.Println()

	fmt.Println("🔸 Выполняем: curr = curr.Next")
	curr = curr.Next

	fmt.Printf("Адрес переменной curr (НЕ ИЗМЕНИЛСЯ!): %p\n", &curr)
	fmt.Printf("Значение переменной curr (ИЗМЕНИЛОСЬ!): %p\n", curr)
	fmt.Printf("Значение в узле: %d\n", curr.Val)
	fmt.Println()

	fmt.Println("🔸 Ещё раз: curr = curr.Next")
	curr = curr.Next

	fmt.Printf("Адрес переменной curr (ОПЯТЬ НЕ ИЗМЕНИЛСЯ!): %p\n", &curr)
	fmt.Printf("Значение переменной curr (ОПЯТЬ ИЗМЕНИЛОСЬ!): %p\n", curr)
	fmt.Printf("Значение в узле: %d\n", curr.Val)

	fmt.Println("\n💡 ВЫВОД:")
	fmt.Println("✅ Адрес переменной curr всегда один и тот же")
	fmt.Println("✅ Значение переменной curr (куда она указывает) меняется")
	fmt.Println("✅ Сами объекты в памяти НЕ ТРОГАЮТСЯ!")
}

func demonstrateInAlgorithm() {
	fmt.Println("\n=== В КОНТЕКСТЕ АЛГОРИТМА ===")

	dummyHead := &ListNode{Val: 0}
	curr := dummyHead

	fmt.Printf("🎯 Адрес переменной dummyHead: %p\n", &dummyHead)
	fmt.Printf("🎯 Адрес переменной curr: %p\n", &curr)
	fmt.Printf("📍 Значение dummyHead (объект): %p\n", dummyHead)
	fmt.Printf("📍 Значение curr (объект): %p\n", curr)
	fmt.Println()

	// Создаём первый узел
	curr.Next = &ListNode{Val: 7}
	fmt.Println("✅ Создали curr.Next")

	// Перемещаем curr
	fmt.Println("🔸 ПЕРЕД curr = curr.Next:")
	fmt.Printf("   Переменная curr находится по адресу: %p\n", &curr)
	fmt.Printf("   Переменная dummyHead находится по адресу: %p\n", &dummyHead)
	fmt.Printf("   curr указывает на объект: %p\n", curr)
	fmt.Printf("   dummyHead указывает на объект: %p\n", dummyHead)

	curr = curr.Next

	fmt.Println("🔸 ПОСЛЕ curr = curr.Next:")
	fmt.Printf("   Переменная curr ВСЁ ЕЩЁ по адресу: %p (НЕ ИЗМЕНИЛСЯ!)\n", &curr)
	fmt.Printf("   Переменная dummyHead ВСЁ ЕЩЁ по адресу: %p (НЕ ИЗМЕНИЛСЯ!)\n", &dummyHead)
	fmt.Printf("   curr ТЕПЕРЬ указывает на объект: %p (ИЗМЕНИЛОСЬ!)\n", curr)
	fmt.Printf("   dummyHead ВСЁ ЕЩЁ указывает на: %p (НЕ ИЗМЕНИЛОСЬ!)\n", dummyHead)

	fmt.Println("\n🧠 ГЛАВНОЕ ПОНИМАНИЕ:")
	fmt.Println("🔹 Переменные curr и dummyHead живут в фиксированных местах памяти")
	fmt.Println("🔹 Мы меняем только их ЗНАЧЕНИЯ (на какой объект они указывают)")
	fmt.Println("🔹 Объекты ListNode остаются связанными через поле Next")
	fmt.Println("🔹 dummyHead всегда может дойти до любого узла по цепочке!")
}

func main() {
	demonstratePointerVsValue()
	demonstrateInAlgorithm()
}
