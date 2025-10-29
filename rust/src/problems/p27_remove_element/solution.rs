pub struct Solution;

impl Solution {
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        nums.retain(|&num| num != val);
        nums.len() as i32
    }
}



#[test]
fn test_solution_case_1() {
    let mut nums = vec![4,3,3,4];
    assert_eq!(Solution::remove_element(&mut nums, 3), 2);
    assert_eq!(nums, [4,4].to_vec());
}