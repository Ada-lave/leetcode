pub struct Solution;

impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        if nums.len() == 1 {
            return nums[0];
        }
        let mut candidate = nums[0];
        let mut counter = 1;
        for i in 1..nums.len() {
            if nums[i] == candidate {
                counter += 1;
            } else if counter == 0 {
                candidate = nums[i];
                counter = 1;
            } else {
                counter -= 1;
            }
        }
        candidate
    }
}

#[test]
fn test_solution_case_1() {
    let nums = vec![2,2,1,1,1,2,2];
    assert_eq!(Solution::majority_element(nums), 2);
}

#[test]
fn test_solution_case_2() {
    let nums = vec![3,2,3];
    assert_eq!(Solution::majority_element(nums), 3);
}
