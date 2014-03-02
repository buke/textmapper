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
package org.textmapper.tool.test.bootstrap.a.ast;

import org.textmapper.tool.test.bootstrap.a.SampleATree.TextSource;

public abstract class AstNode implements IAstNode {
	
	protected TextSource source;
	protected int line;
	protected int offset;
	protected int column;
	protected int endline;
	protected int endoffset;
	protected int endcolumn;

	public AstNode(TextSource source, int line, int offset, int column, int endline, int endoffset, int endcolumn) {
		this.source = source;
		this.line = line;
		this.offset = offset;
		this.column = column;
		this.endline = endline;
		this.endoffset = endoffset;
		this.endcolumn = endcolumn;
	}

	public int getLine() {
		return this.line;
	}

	public int getOffset() {
		return this.offset;
	}

	public int getColumn() {
		return this.column;
	}

	public int getEndline() {
		return this.endline;
	}

	public int getEndoffset() {
		return this.endoffset;
	}

	public int getEndcolumn() {
		return this.endcolumn;
	}

	public TextSource getSource() {
		return source;
	}

	public String toString() {
		return source == null ? "" : source.getText(offset, endoffset);
	}

	//public abstract void accept(Visitor v);
}

