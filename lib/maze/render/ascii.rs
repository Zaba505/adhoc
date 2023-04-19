// Copyright (c) 2023 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

use std::io;
use std::iter;

use crate::render::Renderer as RendererTrait;
use crate::Walls;

pub struct Renderer {}

impl<T: crate::Cell> RendererTrait<T> for Renderer {
    fn render(
        &self,
        w: &mut impl io::Write,
        maze: &impl crate::Maze<T>,
    ) -> Result<(), Box<dyn std::error::Error>> {
        let width = maze.width();

        iter::repeat(0..width)
            .take(maze.height())
            .enumerate()
            .map(|(row, cols)| {
                cols.map(|col| maze.at(row, col)).enumerate().fold(
                    vec![""; width * 9 + 3],
                    |mut acc, (idx, cell)| {
                        let walls = cell.walls();

                        // Top wall
                        acc[3 * idx] = "1";
                        acc[3 * idx + 1] = if walls.contains(Walls::T) { "1" } else { "0" };
                        acc[3 * idx + 2] = "1";

                        // Middle of left and right walls
                        acc[3 * idx + (3 * width + 1)] =
                            if walls.contains(Walls::L) { "1" } else { "0" };
                        acc[3 * idx + (3 * width + 1) + 1] = "0";
                        acc[3 * idx + (3 * width + 1) + 2] =
                            if walls.contains(Walls::R) { "1" } else { "0" };

                        // Bottom wall
                        acc[3 * idx + 2 * (3 * width + 1)] = "1";
                        acc[3 * idx + 2 * (3 * width + 1) + 1] =
                            if walls.contains(Walls::B) { "1" } else { "0" };
                        acc[3 * idx + 2 * (3 * width + 1) + 2] = "1";

                        if idx == width - 1 {
                            acc[3 * width] = "\n";
                            acc[3 * width + 3 * width + 1] = "\n";
                            acc[3 * width + 2 * (3 * width + 1)] = "\n";
                        }

                        acc
                    },
                )
            })
            .try_for_each(|row| write!(w, "{}", row.join("")))
            .map_err(|e| Box::new(e) as Box<dyn std::error::Error>)
    }
}

pub fn render<T: crate::Cell>(
    w: &mut impl io::Write,
    maze: &impl crate::Maze<T>,
) -> Result<(), Box<dyn std::error::Error>> {
    let r = Renderer {};
    r.render(w, maze)
}

#[cfg(test)]
mod will_render {
    use super::*;

    struct Cell(Walls);

    impl crate::Cell for Cell {
        fn walls(&self) -> Walls {
            self.0
        }
    }

    #[test]
    fn if_the_maze_is_1_by_1() {
        let mut out = Vec::new();
        let maze: crate::Constant<Cell, 1, 1> = crate::Constant {
            cells: [[Cell(Walls::LRTB)]],
        };
        render(&mut out, &maze).unwrap();

        let expected = "111
101
111
";
        let actual = std::str::from_utf8(&out).unwrap();
        assert_eq!(expected, actual);
    }

    #[test]
    fn if_the_maze_is_2_by_2() {
        let mut out = Vec::new();
        let maze: crate::Constant<Cell, 2, 2> = crate::Constant {
            cells: [
                [Cell(Walls::LT), Cell(Walls::RT)],
                [Cell(Walls::LB), Cell(Walls::RB)],
            ],
        };
        render(&mut out, &maze).unwrap();

        let expected = "111111
100001
101101
101101
100001
111111
";
        let actual = std::str::from_utf8(&out).unwrap();
        assert_eq!(expected, actual);
    }

    #[test]
    fn if_the_maze_is_3_by_3() {
        let mut out = Vec::new();
        let maze: crate::Constant<Cell, 3, 3> = crate::Constant {
            cells: [
                [Cell(Walls::LTB), Cell(Walls::T), Cell(Walls::RTB)],
                [Cell(Walls::LRT), Cell(Walls::LR), Cell(Walls::LRT)],
                [Cell(Walls::LB), Cell(Walls::B), Cell(Walls::RB)],
            ],
        };
        render(&mut out, &maze).unwrap();

        let expected = "111111111
100000001
111101111
111101111
101101101
101101101
101101101
100000001
111111111
";
        let actual = std::str::from_utf8(&out).unwrap();
        assert_eq!(expected, actual);
    }
}
