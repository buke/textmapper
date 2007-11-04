package net.sf.lapg.templates.test.cases;

import java.util.Hashtable;

import junit.framework.Assert;
import net.sf.lapg.templates.api.IEvaluationEnvironment;
import net.sf.lapg.templates.test.TemplateTestCase;

public class LoopTest extends TemplateTestCase {

	public void testForEach() {
		Hashtable<String,String[]> h = new Hashtable<String,String[]>();
		h.put("list", new String[] { "a", "b" });
		IEvaluationEnvironment env = new TestEnvironment(getClass().getClassLoader(), "net/sf/lapg/templates/test/ltp");

		// test 1
		String q = env.executeTemplate("loop.loop1", h, null);
		Assert.assertEquals("Hmm: \r\n\r\n0: a\r\n1: b\r\n\r\n", q);

		// test 2
		q = env.executeTemplate("loop.loop2", h, null);
		Assert.assertEquals("\r\nHmm: \r\n\r\n0: a\r\n1: b\r\n\r\n", q);

		// test 3
		h.put("list", new String[] {});
		q = env.executeTemplate("loop.loop2", h, null);
		Assert.assertEquals("\r\nHmm: \r\n\r\n\r\n", q);
	}

	public void testCall() {
		Hashtable<String,String[]> h = new Hashtable<String,String[]>();
		h.put("list", new String[] { "a", "b" });
		IEvaluationEnvironment env = new TestEnvironment(getClass().getClassLoader(), "net/sf/lapg/templates/test/ltp");

		// test 1
		String q = env.executeTemplate("format.callTempl", h, null);
		Assert.assertEquals("\r\nstatic int a[] {\r\n	0,\r\n1,\r\n2,\r\n3\r\n};\r\n\r\n", q);

		// test 2
		q = env.executeTemplate("format.useformat", h, null);
		Assert.assertEquals("\r\nstatic int a[] {\r\n	1,2,aa,4,5,\n6,7,8,9,10,\n11,12,13,14,\n15,16,17,19,\n20,21,22,23,\n24,25\r\n};\r\n\r\n", q);
	}

	public void testArithm() {
		IEvaluationEnvironment env = new TestEnvironment(getClass().getClassLoader(), "net/sf/lapg/templates/test/ltp");

		// test 1
		String q = env.executeTemplate("loop.arithm1", new Object(), null);
		Assert.assertEquals("\r\n2 = 2\r\n\r\n", q);

		// test 2
		q = env.executeTemplate("loop.arithm2", new Object(), null);
		Assert.assertEquals("\r\n10 = 10\r\n\r\n", q);

		// test 3
		q = env.executeTemplate("loop.arithm3", new Object(), null);
		Assert.assertEquals("\r\n-1 = -1\r\n\r\n", q);

		// test 4
		q = env.executeTemplate("loop.arithm4", new Object(), null);
		Assert.assertEquals("2 = 2\r\n", q);

		// test 5
		q = env.executeTemplate("loop.arithm5", new Object(), null);
		Assert.assertEquals("true false true true false -2\r\n", q);

		// test 6
		q = env.executeTemplate("loop.arithm6", new Object(), null);
		Assert.assertEquals("uh: lite1\r\noh: okey\r\n", q);
	}

	public void testAssert() {
		Hashtable<String,String[]> h = new Hashtable<String,String[]>();
		h.put("list", new String[] { "w1", "w2" });
		TestEnvironment env = new TestEnvironment(getClass().getClassLoader(), "net/sf/lapg/templates/test/ltp");

		// test 1
		env.addErrors("Evaluation of `l` failed for java.util.Hashtable: null");
		env.addErrors("Assertion `list.length > 5` failed for java.util.Hashtable");
		env.addErrors("Assertion `list[1] == 'w4'` failed for java.util.Hashtable");
		String q = env.executeTemplate("format.assertit", h, null);
		env.assertEmptyErrors();
	}
}
