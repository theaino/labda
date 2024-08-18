use pest_derive::Parser;
use pest::Parser;
use pest::iterators::Pair;

#[derive(Debug)]
pub enum ParseError {
    ParseError(String),
    InvalidRule(Rule),
    RuleExpected(Rule)
}

type ParseResult<T> = Result<T, ParseError>;

#[derive(Debug)]
pub enum Expr {
    Application(Box<Expr>, Box<Expr>),
    Abstraction(String, Box<Expr>),
    Variable(String)
}

#[derive(Parser)]
#[grammar = "labda.pest"]
struct LabdaParser;

pub fn parse(input: &str) -> ParseResult<Expr> {
    let pairs = match LabdaParser::parse(Rule::expression, input) {
        Ok(x) => Ok(x),
        Err(err) => Err(ParseError::ParseError(err.to_string()))
    }?;
    Expr::from(match pairs.peek() {
        Some(x) => x,
        None => return Err(ParseError::RuleExpected(Rule::expression))
    })
}

impl Expr {
    fn from(pair: Pair<Rule>) -> ParseResult<Expr> {
        match pair.as_rule() {
            Rule::variable => Ok(Expr::Variable(pair.as_str().to_string())),
            Rule::application => Expr::from_application(pair),
            Rule::abstraction => Expr::from_abstraction(pair),
            rule => Err(ParseError::InvalidRule(rule))
        }
    }

    fn from_application(pair: Pair<Rule>) -> ParseResult<Expr> {
        let mut pairs = pair.clone().into_inner();
        let abstraction = match pairs.next() {
            Some(x) => x,
            None => return Err(ParseError::RuleExpected(Rule::expression))
        };
        let operand = match pairs.next() {
            Some(x) => x,
            None => return Err(ParseError::RuleExpected(Rule::expression))
        };
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
        let variable = match pairs.next() {
            Some(x) => x.as_str(),
            None => return Err(ParseError::RuleExpected(Rule::variable))
        };
        let expr = match pairs.next() {
            Some(x) => x,
            None => return Err(ParseError::RuleExpected(Rule::expression))
        };
        Ok(Expr::Abstraction(
                variable.to_string(),
                Box::from(Expr::from(expr)?)
        ))
    }
}

