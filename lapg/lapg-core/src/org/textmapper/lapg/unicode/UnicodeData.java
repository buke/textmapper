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
package org.textmapper.lapg.unicode;

import org.textmapper.lapg.api.regex.CharacterSet;
import org.textmapper.lapg.common.CharacterSetImpl;

import java.util.HashMap;
import java.util.Map;

/**
 * Gryaznov Evgeny, 4/22/12
 */
public class UnicodeData {

	private static UnicodeData INSTANCE = new UnicodeData();

	public static UnicodeData getInstance() {
		return INSTANCE;
	}

	private Map<String, String> rawData;
	private Map<String, CharacterSet> set = new HashMap<String, CharacterSet>();

	private UnicodeData() {
	}

	public CharacterSet getCharacterSet(String cl) {
		if (rawData == null) {
			rawData = new HashMap<String, String>();
			String[] properties = UnicodeDataTables.PROPERTIES;
			for (int i = 0; i < properties.length; ) {
				rawData.put(properties[i++], properties[i++]);
			}
		}
		CharacterSet result = set.get(cl);
		if (result != null || !(rawData.containsKey(cl))) {
			return result;
		}

		String data = rawData.get(cl);
		result = decode(data);
		set.put(cl, result);
		return result;
	}

	private static CharacterSet decode(String data) {
		int len = data.charAt(0);
		int[] set = new int[len];
		int index = 1;
		for (int i = 0; i < len; i++) {
			int val = data.codePointAt(index++);
			set[i] = val;
			if (val > 0xffff) {
				index++;
			}
		}
		return new CharacterSetImpl(set);
	}
}
