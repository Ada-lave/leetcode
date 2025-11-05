pub struct Solution;

impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let (mut l1, mut l2) = (m -1, n - 1);
        let mut l3 = (m + n) - 1;
        while l2 >= 0 {
            if l1 >= 0 && nums1[l1 as usize] > nums2[l2 as usize] {
                nums1[l3 as usize] = nums1[l1 as usize];
                l1 -= 1;
            } else {
                nums1[l3 as usize] = nums2[l2 as usize];

                l2 -= 1;
            }
            l3 -= 1;
        }
    }
}



#[test]
fn test_solution_case_1() {
    let mut nums1 = vec![1,2,3,0,0,0];
    let m = 3;
    let mut nums2 = vec![2,5,6];
    let n = 3;
    Solution::merge(&mut nums1, m, &mut nums2, n);
    assert_eq!(nums1,vec![1,2,2,3,5,6]);
}

#[test]
fn test_solution_case_2() {
    let mut nums1 = vec![0];
    let m = 0;
    let mut nums2 = vec![1];
    let n = 1;
    Solution::merge(&mut nums1, m, &mut nums2, n);
    assert_eq!(nums1, vec![1]);
}