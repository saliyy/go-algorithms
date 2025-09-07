package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersWithExplanation(l1 *ListNode, l2 *ListNode) *ListNode {
	// Создаём "фиктивный" узел - это трюк для упрощения кода
	dummyHead := &ListNode{Val: 0}
	fmt.Printf("1. Создали dummyHead: %p, значение: %v\n", dummyHead, dummyHead.Val)

	// curr указывает на тот же объект что и dummyHead
	curr := dummyHead
	fmt.Printf("2. curr = dummyHead, curr указывает на: %p\n", curr)

	carry := 0
	iteration := 1

	for l1 != nil || l2 != nil || carry != 0 {
		fmt.Printf("\n--- Итерация %d ---\n", iteration)

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

		fmt.Printf("3. Сумма: %d, результат: %d\n", sum, sum%10)

		// КЛЮЧЕВОЙ МОМЕНТ: создаём новый узел и привязываем к текущему
		curr.Next = &ListNode{Val: sum % 10}
		fmt.Printf("4. Создали новый узел: %p, curr.Next теперь указывает на него\n", curr.Next)

		// ВАЖНО: curr перемещается на новый узел, но dummyHead остаётся на месте!
		fmt.Printf("5. ПЕРЕД curr = curr.Next: curr = %p\n", curr)
		curr = curr.Next
		fmt.Printf("6. ПОСЛЕ curr = curr.Next: curr = %p\n", curr)
		fmt.Printf("7. dummyHead всё ещё указывает на: %p (не изменился!)\n", dummyHead)

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		iteration++
	}

	// Возвращаем dummyHead.Next, чтобы пропустить фиктивный узел
	fmt.Printf("\n8. Возвращаем dummyHead.Next: %p\n", dummyHead.Next)
	return dummyHead.Next
}

func printList(head *ListNode) {
	fmt.Print("Список: ")
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Создаём тестовые списки: 342 + 465 = 807
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	fmt.Println("=== Объяснение работы указателей ===")
	result := addTwoNumbersWithExplanation(l1, l2)
	fmt.Println("\n=== Результат ===")
	printList(result)
}
