use pest_derive::Parser;
use pest::Parser;
use pest::iterators::Pair;
use std::{fmt, fs, path::Path};

use crate::LB_PATHS;

#[derive(Debug)]
pub enum ParseError {
    SyntaxError(String),
    FileError(String)
}

impl fmt::Display for ParseError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            ParseError::SyntaxError(message) =>
                write!(f, "Syntax error\n{}", message),
            ParseError::FileError(path) =>
                write!(f, "Failed to read file {}\n", path),
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

pub fn parse(input: &str, local_path: &str) -> ParseResult<Expr> {
    let pairs = match LabdaParser::parse(Rule::main, input) {
        Ok(x) => Ok(x),
        Err(err) => Err(ParseError::SyntaxError(err.to_string()))
    }?;
    Expr::from(pairs.peek().unwrap(), local_path)
}

pub fn parse_file(path: &str, local_path: &str) -> ParseResult<Expr> {
    let mut lb_paths = LB_PATHS.get().unwrap().clone();
    lb_paths.push(String::from(local_path));
    for lb_path in lb_paths {
        let file_path = Path::new(lb_path.as_str()).join(path);
        if let Ok(contents) = fs::read_to_string(file_path.clone()) {
            return parse(contents.as_str(), file_path.parent().unwrap().to_str().unwrap());
        }
    }
    Err(ParseError::FileError(path.to_string()))
}

impl Expr {
    fn from(pair: Pair<Rule>, local_path: &str) -> ParseResult<Expr> {
        match pair.as_rule() {
            Rule::variable => Ok(Expr::Variable(pair.as_str().to_string())),
            Rule::application => Expr::from_application(pair, local_path),
            Rule::abstraction => Expr::from_abstraction(pair, local_path),
            Rule::wrapper => Expr::from_wrapper(pair, local_path),
            Rule::opener => Expr::from_opener(pair, local_path),
            Rule::remote => Expr::from_remote(pair, local_path),
            rule => panic!("{:?}", rule)
        }
    }

    fn from_application(pair: Pair<Rule>, local_path: &str) -> ParseResult<Expr> {
        let mut pairs = pair.clone().into_inner();
        let abstraction = pairs.next().unwrap();
        let operand = pairs.next().unwrap();
        let mut expr = Expr::Application(
                Box::from(Expr::from(abstraction, local_path)?),
                Box::from(Expr::from(operand, local_path)?)
        );
        loop {
            let operand = match pairs.next() {
                Some(x) => x,
                None => break
            };
            expr = Expr::Application(
                Box::from(expr),
                Box::from(Expr::from(operand, local_path)?)
            );
        };
        Ok(expr)
    }

    fn from_abstraction(pair: Pair<Rule>, local_path: &str) -> ParseResult<Expr> {
        let mut pairs = pair.clone().into_inner();
        let variable = pairs.next().unwrap().as_str();
        let expr = pairs.next().unwrap();
        Ok(Expr::Abstraction(
                variable.to_string(),
                Box::from(Expr::from(expr, local_path)?)
        ))
    }
    
    fn from_wrapper(pair: Pair<Rule>, local_path: &str) -> ParseResult<Expr> {
        let mut pairs = pair.clone().into_inner();
        let variable = pairs.next().unwrap().as_str();
        let value = pairs.next().unwrap();
        let expr = pairs.next().unwrap();
        Ok(Expr::Application(
                Box::from(Expr::Abstraction(
                        variable.to_string(),
                        Box::from(Expr::from(expr, local_path)?)
                )),
                Box::from(Expr::from(value, local_path)?)
        ))
    }

    fn from_opener(pair: Pair<Rule>, local_path: &str) -> ParseResult<Expr> {
        let mut pairs = pair.clone().into_inner();
        let expr = pairs.next().unwrap();
        Expr::from(expr, local_path)
    }

    fn from_remote(pair: Pair<Rule>, local_path: &str) -> ParseResult<Expr> {
        let mut pairs = pair.clone().into_inner();
        let location = pairs.next().unwrap().as_str();
        parse_file(location, local_path)
    }
}

