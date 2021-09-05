package pattern

import (
	"awesomeProject/pattern/decoration"
	"awesomeProject/pattern/generation"
	"awesomeProject/pattern/generic"
	"awesomeProject/pattern/ioc"
	"awesomeProject/pattern/map_reduce"
	"awesomeProject/pattern/pipeline"
	"awesomeProject/pattern/visitor"
)

func Entry() {
	//errors.ErrorsEntry()
	//functional_options.FunctionalOptionsEntry()
	ioc.IocEntry()
	map_reduce.MapReduceEntry()
	generation.GenerationEntry()
	decoration.DecorationEntry()
	pipeline.PipelineEntry()
	visitor.VisitorEntry()
	generic.GenericEntry()
}
