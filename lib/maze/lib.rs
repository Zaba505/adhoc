#![allow(incomplete_features)]
#![feature(generic_const_exprs)]
#![feature(iter_collect_into)]

pub mod render;

use std::time;

use bitflags::bitflags;
use futures;
use rand::Rng;

pub enum Wall {
    Left,
    Top,
    Right,
    Bottom,
}

bitflags! {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash)]
    pub struct Walls: u8 {
        const None = 0;
        const L = 0b00000001;
        const R = 0b00000010;
        const T = 0b00000100;
        const B = 0b00001000;
        const LR = Self::L.bits() | Self::R.bits();
        const LT = Self::L.bits() | Self::T.bits();
        const LB = Self::L.bits() | Self::B.bits();
        const RT = Self::R.bits() | Self::T.bits();
        const RB = Self::R.bits() | Self::B.bits();
        const TB = Self::T.bits() | Self::B.bits();
        const LRT = Self::L.bits() | Self::R.bits() | Self::T.bits();
        const LRB = Self::L.bits() | Self::R.bits() | Self::B.bits();
        const LTB = Self::L.bits() | Self::T.bits() | Self::B.bits();
        const RTB = Self::R.bits() | Self::T.bits() | Self::B.bits();
        const LRTB = Self::L.bits() | Self::R.bits() | Self::T.bits() | Self::B.bits();
    }
}

pub trait Cell {
    fn walls(&self) -> Walls;
}

pub trait Maze<T: Cell> {
    fn width(&self) -> usize;
    fn height(&self) -> usize;

    fn at(&self, i: usize, j: usize) -> &T;
}

#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash)]
pub struct ImmutableCell(Walls);

impl Cell for ImmutableCell {
    fn walls(&self) -> Walls {
        self.0
    }
}

#[derive(Clone, Debug)]
pub struct Immutable {
    width: usize,
    height: usize,
    cells: Vec<ImmutableCell>,
}

impl Immutable {
    pub fn generate(width: usize, height: usize) -> Immutable {
        let mut rng = rand::thread_rng();
        let cells = random_walls(width, height, &mut rng)
            .into_iter()
            .map(|walls| ImmutableCell(walls))
            .collect();

        Immutable {
            width,
            height,
            cells,
        }
    }
}

impl Maze<ImmutableCell> for Immutable {
    fn at(&self, i: usize, j: usize) -> &ImmutableCell {
        if i >= self.width {
            panic!("i must be less than the width of the maze")
        }
        if j >= self.height {
            panic!("j must be less than the height of the maze")
        }
        &self.cells[i * self.width + j]
    }

    fn width(&self) -> usize {
        self.width
    }

    fn height(&self) -> usize {
        self.height
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash)]
pub struct MutableCell<T>(Walls, Option<T>);

impl<T> Cell for MutableCell<T> {
    fn walls(&self) -> Walls {
        self.0
    }
}

#[derive(Clone, Debug)]
pub struct Mutable<T> {
    width: usize,
    height: usize,
    cells: Vec<MutableCell<T>>,
}

impl<T> Mutable<T> {
    pub fn generate(width: usize, height: usize) -> Mutable<T> {
        let mut rng = rand::thread_rng();
        let cells = random_walls(width, height, &mut rng)
            .into_iter()
            .map(|walls| MutableCell(walls, None))
            .collect();

        Mutable {
            width,
            height,
            cells,
        }
    }
}

impl<T> Maze<MutableCell<T>> for Mutable<T> {
    fn at(&self, i: usize, j: usize) -> &MutableCell<T> {
        if i >= self.width {
            panic!("i must be less than the width of the maze")
        }
        if j >= self.height {
            panic!("j must be less than the height of the maze")
        }
        &self.cells[i * self.width + j]
    }

    fn width(&self) -> usize {
        self.width
    }

    fn height(&self) -> usize {
        self.height
    }
}

pub struct Constant<T, const W: usize, const H: usize> {
    pub cells: [[T; W]; H],
}

impl<T: Cell, const W: usize, const H: usize> Maze<T> for Constant<T, W, H> {
    fn width(&self) -> usize {
        W
    }

    fn height(&self) -> usize {
        H
    }

    fn at(&self, i: usize, j: usize) -> &T {
        &self.cells[i][j]
    }
}

fn random_walls(width: usize, height: usize, rng: &mut impl Rng) -> Vec<Walls> {
    let graph = random_maze_graph(width, height, rng);
    build_walls_from_graph(width, height, graph)
}

fn build_walls_from_graph(width: usize, height: usize, graph: Vec<(usize, usize)>) -> Vec<Walls> {
    graph
        .into_iter()
        .fold(vec![Walls::LRTB; width * height], |mut acc, edge| {
            let (u, v) = if edge.0 < edge.1 {
                (edge.0, edge.1)
            } else {
                (edge.1, edge.0)
            };

            acc[u] ^= if u == v - 1 { Walls::R } else { Walls::B };
            acc[v] ^= if u == v - 1 { Walls::L } else { Walls::T };
            acc
        })
}

fn random_maze_graph(width: usize, height: usize, rng: &mut impl Rng) -> Vec<(usize, usize)> {
    let lattice_start = time::Instant::now();
    let directed_edges = lattice(width, height, rng);
    let lattice_end = time::Instant::now();
    println!(
        "generating lattice took: {:?}",
        lattice_end.duration_since(lattice_start)
    );

    let kruskal_start = time::Instant::now();
    let graph = kruskal(width * height, directed_edges)
        .into_iter()
        .map(|(u, v, _)| (u, v))
        .collect();
    let kruskal_end = time::Instant::now();
    println!(
        "running kruskal took: {:?}",
        kruskal_end.duration_since(kruskal_start)
    );

    graph
}

fn lattice(width: usize, height: usize, rng: &mut impl Rng) -> Vec<(usize, usize, f64)> {
    (0..width * height).fold(
        Vec::with_capacity(2 * (height * (width - 1) + width * (height - 1))),
        |mut edges, i| {
            let row = i / width;
            let col = i % width;

            let u = i;
            let v1 = i + 1;
            let v2 = i + width;

            if col < width - 1 {
                edges.push((u, v1, rng.gen_range(0.0..1.0)));
                edges.push((v1, u, rng.gen_range(0.0..1.0)));
            }
            if row < height - 1 {
                edges.push((u, v2, rng.gen_range(0.0..1.0)));
                edges.push((v2, u, rng.gen_range(0.0..1.0)));
            }
            edges
        },
    )
}

fn kruskal(
    num_of_vertices: usize,
    mut edges: Vec<(usize, usize, f64)>,
) -> Vec<(usize, usize, f64)> {
    let mut sorting: Vec<usize> = Vec::with_capacity(num_of_vertices);
    let mut mst: Vec<(usize, usize, f64)> = Vec::with_capacity(num_of_vertices - 1);

    let sorting_start = time::Instant::now();
    edges.sort_unstable_by(|a, b| a.2.partial_cmp(&b.2).unwrap());
    let sorting_end = time::Instant::now();
    println!(
        "sorting edges took: {:?}",
        sorting_end.duration_since(sorting_start)
    );

    let mst_start = time::Instant::now();
    edges
        .into_iter()
        .filter(move |edge| {
            let start_async = async { sorting.iter().position(|&i| i == edge.0) };
            let end_async = async { sorting.iter().position(|&i| i == edge.1) };
            let (start, end) =
                futures::executor::block_on(async { futures::join!(start_async, end_async) });
            if start.is_none() && end.is_none() {
                sorting.push(edge.0);
                sorting.push(edge.1);
                return true;
            }
            if start.is_none() {
                sorting.insert(end.unwrap(), edge.0);
                return true;
            }
            if end.is_none() {
                sorting.push(edge.1);
                return true;
            }

            if end.unwrap() < start.unwrap() {
                return false;
            }

            true
        })
        .take(num_of_vertices - 1)
        .collect_into(&mut mst);
    let mst_end = time::Instant::now();
    println!(
        "constructing mst took: {:?}",
        mst_end.duration_since(mst_start)
    );

    mst
}
