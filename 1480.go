package leetcodego

func main() {

}

func runningSum(nums []int) []int {
	var output []int
	var sum int
	for _, v := range nums {
		sum = sum + v
		output = append(output, sum)
	}
	return output
}
