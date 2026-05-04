pub struct Solution;

fn find_greatest_candy_count(candies: &Vec<i32>) -> i32 {
    if candies.len() < 1 {
        return 0;
    }

    let mut kid_with_greatest_idx = 0;

    for (i, c) in candies.iter().enumerate() {
        if c > &candies[kid_with_greatest_idx] {
            kid_with_greatest_idx = i
        }
    }

    return candies[kid_with_greatest_idx];
}

impl Solution {
    pub fn kids_with_candies(candies: Vec<i32>, extra_candies: i32) -> Vec<bool> {
        let greatest_candy_count = find_greatest_candy_count(&candies);
        let mut result= vec![false; candies.len()];
        
        for (i, c) in candies.iter().enumerate() {
            result[i] = (c + extra_candies) >= greatest_candy_count;
        }

        result
    }
}
