use clap::Parser;

mod interpreter;

#[derive(Parser, Debug)]
#[command(about, long_about = None)]
struct Args {
    path: String,
}

fn main() {
    let args = Args::parse();

    let path = args.path;
    let expr = interpreter::parser::parse_file(path.as_str())
        .unwrap_or_else(|err| panic!("{}", err));
    let value = expr.reduce();
    println!("{}", value);
}
