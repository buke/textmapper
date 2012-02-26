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
package org.textway.lapg.ant;

import java.io.File;

import org.textway.lapg.gen.LapgOptions;

import org.apache.tools.ant.BuildException;
import org.apache.tools.ant.Task;

public class Lapg extends Task {

	private final LapgOptions options = new LapgOptions();
	private File input;
	private File outputFolder;

	public void setSource(File source) {
		if (!source.exists() || !source.isFile()) {
			throw new IllegalArgumentException("Invalid source.");
		}
		input = source;
	}

	public void setDestDir(File destDir) {
		if (!destDir.exists() || !destDir.isDirectory()) {
			throw new IllegalArgumentException("Invalid directory.");
		}
		outputFolder = destDir;
	}

	public void addConfiguredOption(Option option) {
		options.getAdditionalOptions().put(option.getName(), option.getValue());
	}

	@Override
	public void execute() throws BuildException {

	}

	public static class Option {
		private String name;
		private String value;

		public String getName() {
			return name;
		}

		public void setName(String name) {
			this.name = name;
		}

		public String getValue() {
			return value;
		}

		public void setValue(String value) {
			this.value = value;
		}
	}
}
