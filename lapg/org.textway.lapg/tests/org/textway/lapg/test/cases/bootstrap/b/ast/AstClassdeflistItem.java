/**
 * Copyright 2002-2012 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textway.lapg.test.cases.bootstrap.b.ast;

import org.textway.lapg.test.cases.bootstrap.b.SampleBTree.TextSource;

public class AstClassdeflistItem extends AstNode {

	private AstClassdef classdef;
	private String identifier;

	public AstClassdeflistItem(AstClassdef classdef, String identifier, TextSource input, int start, int end) {
		super(input, start, end);
		this.classdef = classdef;
		this.identifier = identifier;
	}

	public AstClassdef getClassdef() {
		return classdef;
	}
	public String getIdentifier() {
		return identifier;
	}
	public void accept(AstVisitor v) {
		if (!v.visit(this)) {
			return;
		}

		if (classdef != null) {
			classdef.accept(v);
		}
		// TODO for identifier
	}
}
