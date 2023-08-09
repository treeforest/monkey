package ast

import "github.com/treeforest/monkey/token"

//
//  Node
//  @Description: 节点
//
type Node interface {
	TokenLiteral() string
}

//
//  Statement
//  @Description: 语句
//
type Statement interface {
	Node
	statementNode()
}

//
//  Expression
//  @Description: 表达式
//
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	//
	//  Token
	//  @Description: token.LET 词法单元
	//
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	//
	//  Token
	//  @Description: token.IDENT 词法单元
	//
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
