package net.sf.lapg.templates.api;

public interface ITemplate extends ILocatedEntity {
	String getName();
	String getPackage();
	String apply(Object context, IEvaluationEnvironment env, Object[] arguments) throws EvaluationException;
}
