pub struct Solution;

impl Solution {
    pub fn find_degrees(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let mut output: Vec<i32> = Vec::new();
        for row in matrix {
            let mut count = 0;
            for value in row {
                count += value
            }
            output.push(count);
        }
        return output;
    }
}
