use std::fmt;
use parser::Expr;

pub mod parser;

#[derive(Debug)]
pub enum InterpretError {
    UndefinedVariable(String),
}

type InterpretResult<T> = Result<T, InterpretError>;

impl fmt::Display for InterpretError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            InterpretError::UndefinedVariable(name) =>
                write!(f, "Undefined variable {}", name),
        }
    }
}

impl Expr {
    pub fn beta(self, name: &str, expr: &Expr) -> Expr {
        match self {
            Expr::Application(e1, e2) => Expr::Application(
                Box::from(e1.beta(name, expr)),
                Box::from(e2.beta(name, expr)),
            ),
            Expr::Abstraction(n, e1) => Expr::Abstraction(
                n.to_string(),
                if n == name {
                    Box::from(expr.clone())
                } else {
                    Box::from(e1.beta(name, expr))
                }
            ),
            Expr::Variable(n) => {
                if n == name {
                    expr.clone()
                } else {
                    Expr::Variable(n.to_string())
                }
            }
        }
    }

}

pub fn interpret(expr: Expr) -> InterpretResult<Expr> {
    match expr {
        Expr::Abstraction(name, expr) => Ok(Expr::Abstraction(name, expr)),
        Expr::Application(expr, operand) => {
            let abstraction = interpret(*expr)?;
            if let Expr::Abstraction(name, expr) = abstraction {
                return interpret(expr.beta(&name, &operand));
            }
            panic!();
        },
        Expr::Variable(name) => Err(InterpretError::UndefinedVariable(name))
    }
}
