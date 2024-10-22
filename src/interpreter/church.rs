use crate::interpreter::Expr;

fn numeral(n: i32) -> Expr {
    let mut body = Expr::Variable(String::from("x"));
    for _ in 0..n {
        body = Expr::Application(Box::from(Expr::Variable(String::from("f"))), Box::from(body));
    }
    Expr::Abstraction(String::from("f"), Box::from(Expr::Abstraction(String::from("x"), Box::from(body))))
}

pub fn predefined(name: &str) -> Option<Expr> {
    if let Ok(n) = name.parse::<i32>() {
        Some(numeral(n))
    } else {
        None
    }
}
