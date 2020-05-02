// generated by Textmapper; DO NOT EDIT

package ast

import (
	"fmt"
	"github.com/inspirer/textmapper/tm-parsers/js"
)

func ToJsNode(n *Node) JsNode {
	switch n.Type() {
	case js.Abstract:
		return &Abstract{n}
	case js.AccessibilityModifier:
		return &AccessibilityModifier{n}
	case js.AdditiveExpr:
		return &AdditiveExpr{n}
	case js.Arguments:
		return &Arguments{n}
	case js.ArrayLiteral:
		return &ArrayLiteral{n}
	case js.ArrayPattern:
		return &ArrayPattern{n}
	case js.ArrayType:
		return &ArrayType{n}
	case js.ArrowFunc:
		return &ArrowFunc{n}
	case js.AssertsType:
		return &AssertsType{n}
	case js.AssignmentExpr:
		return &AssignmentExpr{n}
	case js.AssignmentOperator:
		return &AssignmentOperator{n}
	case js.AsyncArrowFunc:
		return &AsyncArrowFunc{n}
	case js.AsyncFunc:
		return &AsyncFunc{n}
	case js.AsyncFuncExpr:
		return &AsyncFuncExpr{n}
	case js.AsyncMethod:
		return &AsyncMethod{n}
	case js.AwaitExpr:
		return &AwaitExpr{n}
	case js.BindingRestElement:
		return &BindingRestElement{n}
	case js.BitwiseAND:
		return &BitwiseAND{n}
	case js.BitwiseOR:
		return &BitwiseOR{n}
	case js.BitwiseXOR:
		return &BitwiseXOR{n}
	case js.Block:
		return &Block{n}
	case js.Body:
		return &Body{n}
	case js.BreakStmt:
		return &BreakStmt{n}
	case js.CallExpr:
		return &CallExpr{n}
	case js.CallSignature:
		return &CallSignature{n}
	case js.Case:
		return &Case{n}
	case js.Catch:
		return &Catch{n}
	case js.Class:
		return &Class{n}
	case js.ClassBody:
		return &ClassBody{n}
	case js.ClassExpr:
		return &ClassExpr{n}
	case js.CoalesceExpr:
		return &CoalesceExpr{n}
	case js.CommaExpr:
		return &CommaExpr{n}
	case js.ComputedPropertyName:
		return &ComputedPropertyName{n}
	case js.ConciseBody:
		return &ConciseBody{n}
	case js.ConditionalExpr:
		return &ConditionalExpr{n}
	case js.ConstructSignature:
		return &ConstructSignature{n}
	case js.ConstructorType:
		return &ConstructorType{n}
	case js.ContinueStmt:
		return &ContinueStmt{n}
	case js.DebuggerStmt:
		return &DebuggerStmt{n}
	case js.Declare:
		return &Declare{n}
	case js.DecoratorCall:
		return &DecoratorCall{n}
	case js.DecoratorExpr:
		return &DecoratorExpr{n}
	case js.Default:
		return &Default{n}
	case js.DefaultParameter:
		return &DefaultParameter{n}
	case js.DoWhileStmt:
		return &DoWhileStmt{n}
	case js.ElementBinding:
		return &ElementBinding{n}
	case js.EmptyDecl:
		return &EmptyDecl{n}
	case js.EmptyStmt:
		return &EmptyStmt{n}
	case js.EqualityExpr:
		return &EqualityExpr{n}
	case js.ExponentiationExpr:
		return &ExponentiationExpr{n}
	case js.ExportClause:
		return &ExportClause{n}
	case js.ExportDecl:
		return &ExportDecl{n}
	case js.ExportDefault:
		return &ExportDefault{n}
	case js.ExportSpecifier:
		return &ExportSpecifier{n}
	case js.ExprStmt:
		return &ExprStmt{n}
	case js.Extends:
		return &Extends{n}
	case js.Finally:
		return &Finally{n}
	case js.ForBinding:
		return &ForBinding{n}
	case js.ForCondition:
		return &ForCondition{n}
	case js.ForFinalExpr:
		return &ForFinalExpr{n}
	case js.ForInStmt:
		return &ForInStmt{n}
	case js.ForInStmtWithVar:
		return &ForInStmtWithVar{n}
	case js.ForOfStmt:
		return &ForOfStmt{n}
	case js.ForOfStmtWithVar:
		return &ForOfStmtWithVar{n}
	case js.ForStmt:
		return &ForStmt{n}
	case js.ForStmtWithVar:
		return &ForStmtWithVar{n}
	case js.Function:
		return &Function{n}
	case js.FunctionExpr:
		return &FunctionExpr{n}
	case js.FunctionType:
		return &FunctionType{n}
	case js.Generator:
		return &Generator{n}
	case js.GeneratorExpr:
		return &GeneratorExpr{n}
	case js.GeneratorMethod:
		return &GeneratorMethod{n}
	case js.Getter:
		return &Getter{n}
	case js.IdentExpr:
		return &IdentExpr{n}
	case js.IfStmt:
		return &IfStmt{n}
	case js.ImportDecl:
		return &ImportDecl{n}
	case js.ImportSpecifier:
		return &ImportSpecifier{n}
	case js.ImportType:
		return &ImportType{n}
	case js.InExpr:
		return &InExpr{n}
	case js.IndexAccess:
		return &IndexAccess{n}
	case js.IndexSignature:
		return &IndexSignature{n}
	case js.IndexedAccessType:
		return &IndexedAccessType{n}
	case js.Initializer:
		return &Initializer{n}
	case js.InstanceOfExpr:
		return &InstanceOfExpr{n}
	case js.IntersectionType:
		return &IntersectionType{n}
	case js.JSXAttributeName:
		return &JSXAttributeName{n}
	case js.JSXClosingElement:
		return &JSXClosingElement{n}
	case js.JSXElement:
		return &JSXElement{n}
	case js.JSXElementName:
		return &JSXElementName{n}
	case js.JSXExpr:
		return &JSXExpr{n}
	case js.JSXLiteral:
		return &JSXLiteral{n}
	case js.JSXNormalAttribute:
		return &JSXNormalAttribute{n}
	case js.JSXOpeningElement:
		return &JSXOpeningElement{n}
	case js.JSXSelfClosingElement:
		return &JSXSelfClosingElement{n}
	case js.JSXSpreadAttribute:
		return &JSXSpreadAttribute{n}
	case js.JSXSpreadExpr:
		return &JSXSpreadExpr{n}
	case js.JSXText:
		return &JSXText{n}
	case js.KeyOfType:
		return &KeyOfType{n}
	case js.LabelIdent:
		return &LabelIdent{n}
	case js.LabelledStmt:
		return &LabelledStmt{n}
	case js.LetOrConst:
		return &LetOrConst{n}
	case js.LexicalBinding:
		return &LexicalBinding{n}
	case js.LexicalDecl:
		return &LexicalDecl{n}
	case js.Literal:
		return &Literal{n}
	case js.LiteralPropertyName:
		return &LiteralPropertyName{n}
	case js.LiteralType:
		return &LiteralType{n}
	case js.LogicalAND:
		return &LogicalAND{n}
	case js.LogicalOR:
		return &LogicalOR{n}
	case js.MappedType:
		return &MappedType{n}
	case js.MemberMethod:
		return &MemberMethod{n}
	case js.MemberVar:
		return &MemberVar{n}
	case js.Method:
		return &Method{n}
	case js.MethodSignature:
		return &MethodSignature{n}
	case js.Module:
		return &Module{n}
	case js.ModuleSpecifier:
		return &ModuleSpecifier{n}
	case js.MultiplicativeExpr:
		return &MultiplicativeExpr{n}
	case js.NameIdent:
		return &NameIdent{n}
	case js.NameSpaceImport:
		return &NameSpaceImport{n}
	case js.NamedImports:
		return &NamedImports{n}
	case js.NewExpr:
		return &NewExpr{n}
	case js.NewTarget:
		return &NewTarget{n}
	case js.NoElement:
		return &NoElement{n}
	case js.NonNullableType:
		return &NonNullableType{n}
	case js.NullableType:
		return &NullableType{n}
	case js.ObjectLiteral:
		return &ObjectLiteral{n}
	case js.ObjectMethod:
		return &ObjectMethod{n}
	case js.ObjectPattern:
		return &ObjectPattern{n}
	case js.ObjectType:
		return &ObjectType{n}
	case js.OptionalCallExpr:
		return &OptionalCallExpr{n}
	case js.OptionalIndexAccess:
		return &OptionalIndexAccess{n}
	case js.OptionalPropertyAccess:
		return &OptionalPropertyAccess{n}
	case js.OptionalTaggedTemplate:
		return &OptionalTaggedTemplate{n}
	case js.Parameters:
		return &Parameters{n}
	case js.Parenthesized:
		return &Parenthesized{n}
	case js.ParenthesizedType:
		return &ParenthesizedType{n}
	case js.PostDec:
		return &PostDec{n}
	case js.PostInc:
		return &PostInc{n}
	case js.PreDec:
		return &PreDec{n}
	case js.PreInc:
		return &PreInc{n}
	case js.PredefinedType:
		return &PredefinedType{n}
	case js.Property:
		return &Property{n}
	case js.PropertyAccess:
		return &PropertyAccess{n}
	case js.PropertyBinding:
		return &PropertyBinding{n}
	case js.PropertySignature:
		return &PropertySignature{n}
	case js.Readonly:
		return &Readonly{n}
	case js.ReadonlyType:
		return &ReadonlyType{n}
	case js.RefIdent:
		return &RefIdent{n}
	case js.Regexp:
		return &Regexp{n}
	case js.RelationalExpr:
		return &RelationalExpr{n}
	case js.RestParameter:
		return &RestParameter{n}
	case js.RestType:
		return &RestType{n}
	case js.ReturnStmt:
		return &ReturnStmt{n}
	case js.Setter:
		return &Setter{n}
	case js.ShiftExpr:
		return &ShiftExpr{n}
	case js.ShorthandProperty:
		return &ShorthandProperty{n}
	case js.SingleNameBinding:
		return &SingleNameBinding{n}
	case js.SpreadElement:
		return &SpreadElement{n}
	case js.SpreadProperty:
		return &SpreadProperty{n}
	case js.Static:
		return &Static{n}
	case js.SuperExpr:
		return &SuperExpr{n}
	case js.SwitchStmt:
		return &SwitchStmt{n}
	case js.SyntaxProblem:
		return &SyntaxProblem{n}
	case js.TaggedTemplate:
		return &TaggedTemplate{n}
	case js.TemplateLiteral:
		return &TemplateLiteral{n}
	case js.This:
		return &This{n}
	case js.ThisType:
		return &ThisType{n}
	case js.ThrowStmt:
		return &ThrowStmt{n}
	case js.TryStmt:
		return &TryStmt{n}
	case js.TsAmbientBinding:
		return &TsAmbientBinding{n}
	case js.TsAmbientClass:
		return &TsAmbientClass{n}
	case js.TsAmbientEnum:
		return &TsAmbientEnum{n}
	case js.TsAmbientExportDecl:
		return &TsAmbientExportDecl{n}
	case js.TsAmbientFunc:
		return &TsAmbientFunc{n}
	case js.TsAmbientGlobal:
		return &TsAmbientGlobal{n}
	case js.TsAmbientImportAlias:
		return &TsAmbientImportAlias{n}
	case js.TsAmbientInterface:
		return &TsAmbientInterface{n}
	case js.TsAmbientModule:
		return &TsAmbientModule{n}
	case js.TsAmbientNamespace:
		return &TsAmbientNamespace{n}
	case js.TsAmbientTypeAlias:
		return &TsAmbientTypeAlias{n}
	case js.TsAmbientVar:
		return &TsAmbientVar{n}
	case js.TsAsConstExpr:
		return &TsAsConstExpr{n}
	case js.TsAsExpr:
		return &TsAsExpr{n}
	case js.TsCastExpr:
		return &TsCastExpr{n}
	case js.TsConditional:
		return &TsConditional{n}
	case js.TsConst:
		return &TsConst{n}
	case js.TsDynamicImport:
		return &TsDynamicImport{n}
	case js.TsEnum:
		return &TsEnum{n}
	case js.TsEnumBody:
		return &TsEnumBody{n}
	case js.TsEnumMember:
		return &TsEnumMember{n}
	case js.TsExclToken:
		return &TsExclToken{n}
	case js.TsExport:
		return &TsExport{n}
	case js.TsExportAssignment:
		return &TsExportAssignment{n}
	case js.TsImplementsClause:
		return &TsImplementsClause{n}
	case js.TsImportAliasDecl:
		return &TsImportAliasDecl{n}
	case js.TsImportRequireDecl:
		return &TsImportRequireDecl{n}
	case js.TsIndexMemberDecl:
		return &TsIndexMemberDecl{n}
	case js.TsInterface:
		return &TsInterface{n}
	case js.TsInterfaceExtends:
		return &TsInterfaceExtends{n}
	case js.TsNamespace:
		return &TsNamespace{n}
	case js.TsNamespaceBody:
		return &TsNamespaceBody{n}
	case js.TsNamespaceExportDecl:
		return &TsNamespaceExportDecl{n}
	case js.TsNonNull:
		return &TsNonNull{n}
	case js.TsThisParameter:
		return &TsThisParameter{n}
	case js.TsTypeOnly:
		return &TsTypeOnly{n}
	case js.TupleType:
		return &TupleType{n}
	case js.TypeAliasDecl:
		return &TypeAliasDecl{n}
	case js.TypeAnnotation:
		return &TypeAnnotation{n}
	case js.TypeArguments:
		return &TypeArguments{n}
	case js.TypeConstraint:
		return &TypeConstraint{n}
	case js.TypeName:
		return &TypeName{n}
	case js.TypeParameter:
		return &TypeParameter{n}
	case js.TypeParameters:
		return &TypeParameters{n}
	case js.TypePredicate:
		return &TypePredicate{n}
	case js.TypeQuery:
		return &TypeQuery{n}
	case js.TypeReference:
		return &TypeReference{n}
	case js.TypeVar:
		return &TypeVar{n}
	case js.UnaryExpr:
		return &UnaryExpr{n}
	case js.UnionType:
		return &UnionType{n}
	case js.UniqueType:
		return &UniqueType{n}
	case js.Var:
		return &Var{n}
	case js.VariableDecl:
		return &VariableDecl{n}
	case js.VariableStmt:
		return &VariableStmt{n}
	case js.WhileStmt:
		return &WhileStmt{n}
	case js.WithStmt:
		return &WithStmt{n}
	case js.Yield:
		return &Yield{n}
	case js.InsertedSemicolon:
		return &InsertedSemicolon{n}
	case js.MultiLineComment, js.SingleLineComment, js.InvalidToken, js.NoSubstitutionTemplate, js.TemplateHead, js.TemplateMiddle, js.TemplateTail:
		return &Token{n}
	case js.NoType:
		return nilInstance
	}
	panic(fmt.Errorf("ast: unknown node type %v", n.Type()))
	return nil
}
