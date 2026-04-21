pub struct Solution;

impl Solution {
    pub fn merge_alternately(word1: String, word2: String) -> String {

        let (mut i, mut j) = (0, 0);
        let (word1_len, word2_len) = (word1.len(), word2.len());
        
        let mut result_string: Vec<char> = Vec::with_capacity(word1_len + word2_len);


        while i < word1_len || j < word2_len {
            if i < word1_len {
                result_string.push(word1.chars().nth(i).unwrap());
                i += 1;
            }
            if j < word2_len {
                result_string.push(word2.chars().nth(j).unwrap());
                j += 1;
            }
        }

        return result_string.into_iter().collect();
    }
}
