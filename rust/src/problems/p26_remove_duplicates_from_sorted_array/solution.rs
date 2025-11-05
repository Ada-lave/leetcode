pub struct Solution;

impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.is_empty() {
            return 0;
        }
        let mut k = 1;
        for n in 1..nums.len() {
            if nums[n] != nums[k - 1] {
                nums[k] = nums[n];
                k += 1;
            }
        }
        k as i32
    }
}



#[test]
fn test_solution_case_1() {
    let mut nums = vec![0,0,1,1,1,2,2,3,3,4];
    assert_eq!(Solution::remove_duplicates(&mut nums), 5);
    assert_eq!(nums, vec![0, 1, 2, 3, 4, 2, 2, 3, 3, 4])
}

