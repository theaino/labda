use parser::Expr;

pub mod parser;

impl Expr {
    pub fn substitute(self, name: &str, expr: &Expr) -> Expr {
        match self {
            Expr::Application(e1, e2) => Expr::Application(
                Box::from(e1.substitute(name, expr)),
                Box::from(e2.substitute(name, expr)),
            ),
            Expr::Abstraction(n, e1) => Expr::Abstraction(
                n.to_string(),
                if n == name {
                    Box::from(e1)
                } else {
                    Box::from(e1.substitute(name, expr))
                }
            ),
            Expr::Variable(n) => {
                if n == name {
                    expr.clone()
                } else {
                    Expr::Variable(n)
                }
            }
        }
    }

    pub fn reduce(self) -> Expr {
        match self {
            Expr::Abstraction(name, expr) => Expr::Abstraction(
                name,
                Box::from(expr)
            ),
            Expr::Application(expr, operand) => {
                match expr.reduce() {
                    Expr::Abstraction(name, body) => {
                        body.substitute(&name, &operand).reduce()
                    },
                    v => Expr::Application(
                        Box::from(v),
                        Box::from(operand.reduce())
                    )
                }
            },
            v => v,
        }
    }
}

