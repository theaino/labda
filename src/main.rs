use std::{env, fs};

pub mod interpreter;

fn main() {
    let args: Vec<String> = env::args().collect();
    let contents = fs::read_to_string(match args.get(1) {
        Some(x) => x,
        None => panic!("Usage: <file>")
    }).expect("Should have been able to read the file");
    let expr = interpreter::parser::parse(&contents)
        .unwrap_or_else(|err| panic!("{}", err));
    let value = interpreter::interpret(expr);
    println!("{}", value.unwrap());
}
