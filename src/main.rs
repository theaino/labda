pub mod interpreter;

fn main() {
    let expr = interpreter::parser::parse("a b c d").unwrap();
    println!("{:?}", expr);
}
