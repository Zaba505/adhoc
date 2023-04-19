// Copyright (c) 2023 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

pub mod ascii;

use std::io;

pub trait Renderer<T: crate::Cell> {
    fn render(
        &self,
        w: &mut impl io::Write,
        maze: &impl crate::Maze<T>,
    ) -> Result<(), Box<dyn std::error::Error>>;
}
