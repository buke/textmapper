/**
 * Copyright 2002-2015 Evgeny Gryaznov
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
package org.textmapper.tool.bootstrap.a;

import java.io.IOException;
import java.io.Reader;
import java.text.MessageFormat;

public class SampleALexer {

	public static class Span {
		public Object value;
		public int symbol;
		public int state;
		public int line;
		public int column;
		public int offset;
		public int endline;
		public int endcolumn;
		public int endoffset;
	}

	public interface Tokens {
		int Unavailable_ = -1;
		int eoi = 0;
		int identifier = 1;
		int _skip = 2;
		int Lclass = 3;
		int Lcurly = 4;
		int Rcurly = 5;
		int error = 6;
	}

	public interface ErrorReporter {
		void error(String message, int line, int offset, int column, int endline, int endoffset, int endcolumn);
	}

	public static final int TOKEN_SIZE = 2048;

	private Reader stream;
	final private ErrorReporter reporter;

	final private char[] data = new char[2048];
	private int datalen;
	private int tokenOffset;
	private int l;
	private int charOffset;
	private int chr;

	private int state;

	final private StringBuilder tokenBuffer = new StringBuilder(TOKEN_SIZE);

	private int tokenLine;
	private int currLine;
	private int currColumn;
	private int currOffset;

	public SampleALexer(Reader stream, ErrorReporter reporter) throws IOException {
		this.reporter = reporter;
		reset(stream);
	}

	public void reset(Reader stream) throws IOException {
		this.state = 0;
		tokenLine = currLine = 1;
		currColumn = 1;
		currOffset = 0;
		this.stream = stream;
		datalen = stream.read(data);
		l = 0;
		tokenOffset = -1;
		if (l + 1 >= datalen) {
			if (l < datalen) {
				data[0] = data[l];
				datalen = Math.max(stream.read(data, 1, data.length - 1) + 1, 1);
			} else {
				datalen = stream.read(data);
			}
			l = 0;
		}
		charOffset = l;
		chr = l < datalen ? data[l++] : -1;
		if (chr >= Character.MIN_HIGH_SURROGATE && chr <= Character.MAX_HIGH_SURROGATE && l < datalen &&
				Character.isLowSurrogate(data[l])) {
			chr = Character.toCodePoint((char) chr, data[l++]);
		}
	}

	protected void advance() throws IOException {
		if (chr == -1) return;
		currOffset += l - charOffset;
		currColumn += l - charOffset;
		if (chr == '\n') {
			currColumn = 1;
			currLine++;
		}
		if (l + 1 >= datalen) {
			if (tokenOffset >= 0) {
				tokenBuffer.append(data, tokenOffset, l - tokenOffset);
				tokenOffset = 0;
			}
			if (l < datalen) {
				data[0] = data[l];
				datalen = Math.max(stream.read(data, 1, data.length - 1) + 1, 1);
			} else {
				datalen = stream.read(data);
			}
			l = 0;
		}
		charOffset = l;
		chr = l < datalen ? data[l++] : -1;
		if (chr >= Character.MIN_HIGH_SURROGATE && chr <= Character.MAX_HIGH_SURROGATE && l < datalen &&
				Character.isLowSurrogate(data[l])) {
			chr = Character.toCodePoint((char) chr, data[l++]);
		}
	}

	public int getState() {
		return state;
	}

	public void setState(int state) {
		this.state = state;
	}

	public int getTokenLine() {
		return tokenLine;
	}

	public int getLine() {
		return currLine;
	}

	public void setLine(int currLine) {
		this.currLine = currLine;
	}

	public int getColumn() {
		return currColumn;
	}

	public void setColumn(int currColumn) {
		this.currColumn = currColumn;
	}

	public int getOffset() {
		return currOffset;
	}

	public void setOffset(int currOffset) {
		this.currOffset = currOffset;
	}

	public String tokenText() {
		return tokenBuffer.toString();
	}

	public int tokenSize() {
		return tokenBuffer.length();
	}

	private static final short tmCharClass[] = {
		1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 10, 1, 1, 10, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 1, 1, 1, 1, 1, 1,
		1, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
		8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 1, 1, 1, 1, 8,
		1, 4, 8, 2, 8, 8, 8, 8, 8, 8, 8, 8, 3, 8, 8, 8,
		8, 8, 8, 5, 8, 8, 8, 8, 8, 8, 8, 6, 1, 7, 1, 1
	};

	private static final short[] tmRuleSymbol = unpack_short(5,
		"\1\2\3\4\5");

	private static final int tmClassesCount = 11;

	private static final short[] tmGoto = unpack_vc_short(110,
		"\1\ufffe\1\uffff\1\1\3\2\1\3\1\4\1\2\1\uffff\1\5\2\ufffd\1\2\1\6\2\2\2\ufffd\2\2" +
		"\3\ufffd\4\2\2\ufffd\2\2\1\ufffd\13\ufffa\13\ufff9\12\ufffc\1\5\2\ufffd\2\2\1\7\1" +
		"\2\2\ufffd\2\2\3\ufffd\3\2\1\10\2\ufffd\2\2\3\ufffd\3\2\1\11\2\ufffd\2\2\1\ufffd" +
		"\2\ufffb\4\2\2\ufffb\2\2\1\ufffb");

	private static short[] unpack_vc_short(int size, String... st) {
		short[] res = new short[size];
		int t = 0;
		int count = 0;
		for (String s : st) {
			int slen = s.length();
			for (int i = 0; i < slen; ) {
				count = i > 0 || count == 0 ? s.charAt(i++) : count;
				if (i < slen) {
					short val = (short) s.charAt(i++);
					while (count-- > 0) res[t++] = val;
				}
			}
		}
		assert res.length == t;
		return res;
	}

	private static int mapCharacter(int chr) {
		if (chr >= 0 && chr < 128) return tmCharClass[chr];
		return chr == -1 ? 0 : 1;
	}

	public Span next() throws IOException {
		Span token = new Span();
		int state;

		tokenloop:
		do {
			token.offset = currOffset;
			tokenLine = token.line = currLine;
			token.column = currColumn;
			if (tokenBuffer.length() > TOKEN_SIZE) {
				tokenBuffer.setLength(TOKEN_SIZE);
				tokenBuffer.trimToSize();
			}
			tokenBuffer.setLength(0);
			tokenOffset = charOffset;

			for (state = this.state; state >= 0; ) {
				state = tmGoto[state * tmClassesCount + mapCharacter(chr)];
				if (state == -1 && chr == -1) {
					token.endoffset = currOffset;
					token.endline = currLine;
					token.endcolumn = currColumn;
					token.symbol = 0;
					token.value = null;
					reporter.error("Unexpected end of input reached", token.line, token.offset, token.column, token.endline, token.endoffset, token.endcolumn);
					token.offset = currOffset;
					break tokenloop;
				}
				if (state >= -1 && chr != -1) {
					currOffset += l - charOffset;
					currColumn += l - charOffset;
					if (chr == '\n') {
						currColumn = 1;
						currLine++;
					}
					if (l + 1 >= datalen) {
						tokenBuffer.append(data, tokenOffset, l - tokenOffset);
						tokenOffset = 0;
						if (l < datalen) {
							data[0] = data[l];
							datalen = Math.max(stream.read(data, 1, data.length - 1) + 1, 1);
						} else {
							datalen = stream.read(data);
						}
						l = 0;
					}
					charOffset = l;
					chr = l < datalen ? data[l++] : -1;
					if (chr >= Character.MIN_HIGH_SURROGATE && chr <= Character.MAX_HIGH_SURROGATE && l < datalen &&
							Character.isLowSurrogate(data[l])) {
						chr = Character.toCodePoint((char) chr, data[l++]);
					}
				}
			}
			token.endoffset = currOffset;
			token.endline = currLine;
			token.endcolumn = currColumn;

			if (state == -1) {
				if (charOffset > tokenOffset) {
					tokenBuffer.append(data, tokenOffset, charOffset - tokenOffset);
				}
				reporter.error(MessageFormat.format("invalid lexeme at line {0}: `{1}`, skipped", currLine, tokenText()), token.line, token.offset, token.column, token.endline, token.endoffset, token.endcolumn);
				token.symbol = -1;
				continue;
			}

			if (state == -2) {
				token.symbol = Tokens.eoi;
				token.value = null;
				break tokenloop;
			}

			if (charOffset > tokenOffset) {
				tokenBuffer.append(data, tokenOffset, charOffset - tokenOffset);
			}

			token.symbol = tmRuleSymbol[-state - 3];
			token.value = null;

		} while (token.symbol == -1 || !createToken(token, -state - 3));
		tokenOffset = -1;
		return token;
	}

	protected boolean createToken(Span token, int ruleIndex) throws IOException {
		boolean spaceToken = false;
		switch (ruleIndex) {
			case 0: // identifier: /[a-zA-Z_][a-zA-Z_0-9]*/
				{ token.value = tokenText(); }
				break;
			case 1: // _skip: /[\n\t\r ]+/
				spaceToken = true;
				break;
		}
		return !(spaceToken);
	}

	/* package */ static int[] unpack_int(int size, String... st) {
		int[] res = new int[size];
		boolean second = false;
		char first = 0;
		int t = 0;
		for (String s : st) {
			int slen = s.length();
			for (int i = 0; i < slen; i++) {
				if (second) {
					res[t++] = (s.charAt(i) << 16) + first;
				} else {
					first = s.charAt(i);
				}
				second = !second;
			}
		}
		assert !second;
		assert res.length == t;
		return res;
	}

	/* package */ static short[] unpack_short(int size, String... st) {
		short[] res = new short[size];
		int t = 0;
		for (String s : st) {
			int slen = s.length();
			for (int i = 0; i < slen; i++) {
				res[t++] = (short) s.charAt(i);
			}
		}
		assert res.length == t;
		return res;
	}
}