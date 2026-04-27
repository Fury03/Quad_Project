package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
    count := 0
    if n <= 0 {
        return 0, fmt.Errorf("error")
    }
	for n > 1 {
        if n%2 == 0 {
            n /= 2
            count++
        } else {
            n = (n*3)+1 
            count++
        }
    }
    return count, nil
}
