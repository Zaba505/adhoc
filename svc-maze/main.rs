// Copyright (c) 2023 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

use std::io;

use maze;
use maze::render::ascii;

fn main() {
    let width = 3;
    let height = 3;
    let m = maze::Immutable::generate(width, height);
    // let mut out = io::stdout();
    let mut out = Vec::with_capacity(9 * width * height + 3 * height);
    match ascii::render(&mut out, &m) {
        Ok(()) => {}
        Err(e) => println!("failed to render maze to ASCII: {}", e),
    };
}
