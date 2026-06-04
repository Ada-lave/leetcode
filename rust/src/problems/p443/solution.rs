pub struct Solution;

impl Solution {
    pub fn compress(chars: &mut Vec<char>) -> i32 {
        if chars.len() == 1 {
            return 1;
        }

        let mut write_ptr = 0;
        let mut current_char = chars[write_ptr];
        let mut counter = 0;

        for read_ptr in 0..chars.len() {
            if chars[read_ptr] != current_char {
                chars[write_ptr] = current_char;
                write_ptr += 1;

                if counter > 1 {
                    for char in counter.to_string().chars() {
                        chars[write_ptr] = char;
                        write_ptr += 1;
                    }

                    counter = 1;
                    current_char = chars[read_ptr];
                } 
            } else {
                counter += 1
            }
        }


        chars[write_ptr] = current_char;
        write_ptr += 1;


        if counter > 1 {
            for char in counter.to_string().chars() {
                chars[write_ptr] = char;
                write_ptr += 1;
            }
        } 

        return write_ptr.try_into().unwrap();
    }
}
