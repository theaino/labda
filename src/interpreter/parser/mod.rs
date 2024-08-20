use pest_derive::Parser;
use pest::Parser;
use pest::iterators::Pair;
use std::fmt;

#[derive(Debug)]
pub enum ParseError {
    SyntaxError(String),
}

impl fmt::Display for ParseError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            ParseError::SyntaxError(message) =>
                write!(f, "Syntax error\n{}", message),
        }
    }
}

type ParseResult<T> = Result<T, ParseError>;

#[derive(Clone, Debug)]
pub enum Expr {
    Application(Box<Expr>, Box<Expr>),
    Abstraction(String, Box<Expr>),
    Variable(String)
}

impl fmt::Display for Expr {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            Expr::Variable(name) => write!(f, "{}", name),
            Expr::Abstraction(name, e) => write!(f, "${}.{}", name, e),
            Expr::Application(e1, e2) => write!(f, "({}) ({})", e1, e2)
        }
    }
}

#[derive(Parser)]
#[grammar = "labda.pest"]
struct LabdaParser;

pub fn parse(input: &str) -> ParseResult<Expr> {
    let pairs = match LabdaParser::parse(Rule::expression, input) {
        Ok(x) => Ok(x),
        Err(err) => Err(ParseError::SyntaxError(err.to_string()))
    }?;
    Expr::from(pairs.peek().unwrap())
}

impl Expr {
    fn from(pair: Pair<Rule>) -> ParseResult<Expr> {
        match pair.as_rule() {
            Rule::variable => Ok(Expr::Variable(pair.as_str().to_string())),
            Rule::application => Expr::from_application(pair),
            Rule::abstraction => Expr::from_abstraction(pair),
            Rule::wrapper => Expr::from_wrapper(pair),
            rule => panic!("{:?}", rule)
        }
    }

    fn from_application(pair: Pair<Rule>) -> ParseResult<Expr> {
        let mut pairs = pair.clone().into_inner();
        let abstraction = pairs.next().unwrap();
        let operand = pairs.next().unwrap();
        let mut expr = Expr::Application(
                Box::from(Expr::from(abstraction)?),
                Box::from(Expr::from(operand)?)
        );
        loop {
            let operand = match pairs.next() {
                Some(x) => x,
                None => break
            };
            expr = Expr::Application(
                Box::from(expr),
                Box::from(Expr::from(operand)?)
            );
        };
        Ok(expr)
    }

    fn from_abstraction(pair: Pair<Rule>) -> ParseResult<Expr> {
        let mut pairs = pair.clone().into_inner();
        let variable = pairs.next().unwrap().as_str();
        let expr = pairs.next().unwrap();
        Ok(Expr::Abstraction(
                variable.to_string(),
                Box::from(Expr::from(expr)?)
        ))
    }
    
    fn from_wrapper(pair: Pair<Rule>) -> ParseResult<Expr> {
        let mut pairs = pair.clone().into_inner();
        let variable = pairs.next().unwrap().as_str();
        let value = pairs.next().unwrap();
        let expr = pairs.next().unwrap();
        Ok(Expr::Application(
                Box::from(Expr::Abstraction(
                        variable.to_string(),
                        Box::from(Expr::from(expr)?)
                )),
                Box::from(Expr::from(value)?)
        ))
    }
}

