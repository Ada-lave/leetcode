pub struct Solution;

impl Solution {
    pub fn max_operations(mut nums: Vec<i32>, k: i32) -> i32 {
        let mut counter = 0;
        nums.sort_unstable();
        let (mut idx1, mut idx2) = (0, nums.len() - 1);

        while idx1 < idx2 {
            let sum_of_idxs = nums[idx1] + nums[idx2];
            match sum_of_idxs.cmp(&k) {
                std::cmp::Ordering::Equal => {
                    counter += 1;
                    idx1 += 1;
                    idx2 -= 1
                }
                
                std::cmp::Ordering::Greater => {
                    idx2 -= 1
                }

                std::cmp::Ordering::Less => {
                    idx1 += 1
                }
            }
        }

        counter
    }
}
