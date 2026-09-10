Дана функция:

func Sum(nums []int) int {
    var wg sync.WaitGroup
    var mu sync.Mutex
    total := 0

    for _, n := range nums {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            mu.Lock()
            total += n
            mu.Unlock()
        }(n)
    }
    wg.Wait()
    return total
}

Напишите табличный тест TestSum.