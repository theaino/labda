use std::{path::Path, sync::OnceLock};
use clap::Parser;

use crate::interpreter::parser::Expr;

mod interpreter;

#[derive(Parser, Debug)]
#[command(about, long_about = None)]
struct Args {
    path: String,

    #[clap(long, short, action)]
    no_prelude: bool,
}

static LB_PATHS: OnceLock<Vec<String>> = OnceLock::new();

fn parent(path: String) -> String {
    String::from(match Path::new(path.as_str()).parent() {
        Some(p) => p.to_str().unwrap(),
        None => ""
    })
}

fn main() {
    let args = Args::parse();

    let path = args.path;

    LB_PATHS.get_or_init(|| {
        vec![parent(String::from(file!())), String::from(env!("CARGO_MANIFEST_DIR"))]
    });

    let mut expr = interpreter::parser::parse_file(path.as_str(), parent(path.clone()).as_str())
        .unwrap_or_else(|err| panic!("{}", err));
    if !args.no_prelude {
        let prelude = interpreter::parser::parse_file("prelude/mod.lb", "").unwrap_or_else(|err| panic!("{}", err));
        expr = Expr::Application(Box::from(prelude), Box::from(expr));
    }
    let value = expr.reduce();
    println!("{}", value);

}
