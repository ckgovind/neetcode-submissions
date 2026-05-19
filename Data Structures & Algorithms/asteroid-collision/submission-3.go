func asteroidCollision(asteroids []int) []int {
    stack := []int{}

    for _, ast := range asteroids {
        survived := true

        for survived && ast < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
            top := stack[len(stack)-1]
            if top > -ast {
                survived = false // incoming left-mover destroyed
            } else if top == -ast {
                stack = stack[:len(stack)-1] // both destroyed
                survived = false
            } else {
                stack = stack[:len(stack)-1] // right-mover destroyed, keep going
            }
        }

        if survived {
            stack = append(stack, ast)
        }
    }

    return stack
}
