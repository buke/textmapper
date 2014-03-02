/**
 * Copyright 2002-2014 Evgeny Gryaznov
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
package org.textmapper.templates.types.ast;

import org.textmapper.templates.types.TypesTree.TextSource;

public class AstMultiplicity extends AstNode {

	private final Integer lo;
	private final boolean hasNoUpperBound;
	private final Integer hi;

	public AstMultiplicity(Integer lo, boolean hasNoUpperBound, Integer hi, TextSource source, int line, int offset, int endoffset) {
		super(source, line, offset, endoffset);
		this.lo = lo;
		this.hasNoUpperBound = hasNoUpperBound;
		this.hi = hi;
	}

	public Integer getLo() {
		return lo;
	}

	public boolean getHasNoUpperBound() {
		return hasNoUpperBound;
	}

	public Integer getHi() {
		return hi;
	}

	public void accept(AstVisitor v) {
		v.visit(this);
	}
}
