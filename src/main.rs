use std::env;

pub mod interpreter;

fn main() {
    let args: Vec<String> = env::args().collect();
    let path = match args.get(1) {
        Some(x) => x,
        None => panic!("Usage: <file>")
    };
    let expr = interpreter::parser::parse_file(path.as_str())
        .unwrap_or_else(|err| panic!("{}", err));
    let value = expr.reduce().reduce();
    println!("{}", value);
}
