/**
 * Copyright 2002-2011 Evgeny Gryaznov
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
package org.textway.xml;


public class XmlData extends XmlElement {

	private final char[] buffer;
	private final int start;
	private final int len;

	XmlData(char[] buffer, int start, int len) {
		this.buffer = buffer;
		this.start = start;
		this.len = len;
	}

	@Override
	public void toString(StringBuffer sb) {
		sb.append(new String(buffer, start, len));
	}

	public String getTitle() {
		return "XMLDATA";
	}
}
