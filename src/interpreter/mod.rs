use std::fmt;
use parser::Expr;

pub mod parser;

#[derive(Debug)]
pub enum InterpretError {
    UndefinedVariable(String),
}

impl fmt::Display for InterpretError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            InterpretError::UndefinedVariable(name) =>
                write!(f, "Undefined variable {}", name),
        }
    }
}

type InterpretResult<T> = Result<T, InterpretError>;

#[derive(Clone, Debug)]
pub struct Abstraction(String, Expr);

pub fn interpret(expr: Expr) -> InterpretResult<Abstraction> {
    match expr {
        Expr::Abstraction(name, expr) => Ok(Abstraction(name, *expr)),
        Expr::Application(expr, operand) => {
            let abstraction = interpret(*expr)?;
            interpret(abstraction.1.beta(&abstraction.0, &operand))
        },
        Expr::Variable(name) => Err(InterpretError::UndefinedVariable(name))
    }
}
