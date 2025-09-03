package main

func removeDuplicates(nums []int) int {
	slow := 0
	same := 1

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
			same = 1
			continue
		}
		if same == 2 {
			continue
		}
		slow++
		nums[slow] = nums[fast]
		same++
	}

	return slow + 1
}

func main() {

}
