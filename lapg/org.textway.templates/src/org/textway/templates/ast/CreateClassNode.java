/**
 * Copyright 2002-2010 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textway.templates.ast;

import org.textway.templates.api.EvaluationContext;
import org.textway.templates.api.EvaluationException;
import org.textway.templates.api.IEvaluationStrategy;
import org.textway.templates.api.types.IClass;
import org.textway.templates.api.types.IType;
import org.textway.templates.ast.TemplatesTree.TextSource;
import org.textway.templates.types.TiExpressionBuilder;

import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;

public class CreateClassNode extends ExpressionNode {

	private final String className;
	private final Map<String,ExpressionNode> fieldInitializers;

	protected CreateClassNode(String className, Map<String, ExpressionNode> fieldInitializers, TextSource source, int offset, int endoffset) {
		super(source, offset, endoffset);
		this.className = className;
		this.fieldInitializers = fieldInitializers;
	}

	public String getClassName() {
		return className;
	}

	public Map<String, ExpressionNode> getFieldInitializers() {
		return fieldInitializers;
	}

	@Override
	public Object evaluate(final EvaluationContext context, final IEvaluationStrategy env) throws EvaluationException {
		return new TiExpressionBuilder<ExpressionNode>() {
			@Override
			public IClass resolveType(String className) {
				if(className.indexOf('.') == -1) {
					className = context.getCurrent().getPackage() + "." + className;
				}
				return env.getTypesRegistry().loadClass(className, CreateClassNode.this);
			}

			@Override
			public Object resolve(ExpressionNode expression, IType type) {
				if(expression instanceof CreateClassNode) {
					CreateClassNode newexpr = (CreateClassNode) expression;
					return convertNew(newexpr, newexpr.getClassName(), newexpr.getFieldInitializers(), type);
				}
				if(expression instanceof ListNode) {
					List<ExpressionNode> content = Arrays.asList(((ListNode) expression).getExpressions());
					return convertArray(expression, content, type);
				}
				if(expression instanceof LiteralNode) {
					Object literal = ((LiteralNode)expression).getLiteral();
					return convertLiteral(expression, literal, type);
				}

				Object result = null;
				try {
					result = env.evaluate(expression, context, true);
					if(result instanceof Boolean || result instanceof Number || result instanceof String) {
						return convertLiteral(expression, result, type);
					}

					// TODO check type
				} catch (EvaluationException e) {
					/* ignore, error is already produced */
				}
				return result;
			}

			@Override
			public void report(ExpressionNode expression, String message) {
				env.fireError(expression, message);
			}
		}.resolve(this, null);
	}

	@Override
	public void toString(StringBuilder sb) {
		sb.append("new ");
		sb.append(className);
		sb.append("(");
		int index = sb.length();
		if(fieldInitializers != null) {
			for (Entry<String, ExpressionNode> entry : fieldInitializers.entrySet()) {
				if(sb.length() > index) {
					sb.append(", ");
				}
				sb.append(entry.getKey());
				sb.append("=");
				entry.getValue().toString(sb);
			}
		}
		sb.append(")");
	}
}