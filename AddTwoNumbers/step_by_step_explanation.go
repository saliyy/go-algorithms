package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printMemoryState(dummyHead, curr *ListNode, step string) {
	fmt.Printf("\n=== %s ===\n", step)
	fmt.Printf("dummyHead указывает на: %p\n", dummyHead)
	fmt.Printf("curr указывает на: %p\n", curr)

	fmt.Printf("Состояние цепочки от dummyHead:\n")
	node := dummyHead
	position := 0
	for node != nil {
		marker := ""
		if node == curr {
			marker = " <-- curr здесь"
		}
		if node == dummyHead {
			marker += " <-- dummyHead здесь"
		}
		fmt.Printf("  Позиция %d: адрес %p, значение %d%s\n", position, node, node.Val, marker)
		node = node.Next
		position++
		if position > 10 { // защита от бесконечного цикла
			break
		}
	}
}

func addTwoNumbersStepByStep(l1 *ListNode, l2 *ListNode) *ListNode {
	// ШАГ 0: Начальное состояние
	dummyHead := &ListNode{Val: 0}
	curr := dummyHead
	printMemoryState(dummyHead, curr, "ШАГ 0: Начальное состояние")

	carry := 0
	iteration := 1

	for l1 != nil || l2 != nil || carry != 0 {
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}
		sum := carry + x + y
		carry = sum / 10

		fmt.Printf("\n🔸 Итерация %d: создаём узел со значением %d\n", iteration, sum%10)

		// КРИТИЧЕСКИЙ МОМЕНТ 1: Создаём связь
		curr.Next = &ListNode{Val: sum % 10}
		fmt.Printf("✅ Выполнили: curr.Next = &ListNode{Val: %d}\n", sum%10)
		fmt.Printf("   Это означает: узел по адресу %p теперь ссылается на новый узел %p\n", curr, curr.Next)

		printMemoryState(dummyHead, curr, fmt.Sprintf("После создания узла в итерации %d", iteration))

		// КРИТИЧЕСКИЙ МОМЕНТ 2: Перемещаем curr
		fmt.Printf("\n🔸 Перемещаем curr на новый узел\n")
		fmt.Printf("   БЫЛО: curr = %p\n", curr)
		curr = curr.Next
		fmt.Printf("   СТАЛО: curr = %p\n", curr)
		fmt.Printf("   ⚠️  ВАЖНО: dummyHead остался на %p и ВСЁ ЕЩЁ МОЖЕТ ДОЙТИ до нового узла по цепочке!\n", dummyHead)

		printMemoryState(dummyHead, curr, fmt.Sprintf("После перемещения curr в итерации %d", iteration))

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		iteration++
	}

	fmt.Printf("\n🏁 ФИНАЛ:\n")
	fmt.Printf("dummyHead находится на: %p\n", dummyHead)
	fmt.Printf("curr находится на: %p\n", curr)
	fmt.Printf("dummyHead.Next (начало результата): %p\n", dummyHead.Next)
	fmt.Printf("\n💡 dummyHead может дойти до curr через цепочку Next!\n")

	return dummyHead.Next
}

func main() {
	// Простой пример: 5 + 5 = 10 (0 с переносом 1)
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5}

	result := addTwoNumbersStepByStep(l1, l2)

	fmt.Printf("\n=== ПРОВЕРКА РЕЗУЛЬТАТА ===\n")
	fmt.Print("Результирующий список: ")
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Println()
}
