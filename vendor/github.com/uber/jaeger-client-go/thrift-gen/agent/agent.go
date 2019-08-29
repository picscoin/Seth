// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package agent

import (
	"bytes"
	"fmt"
	"github.com/uber/jaeger-client-go/thrift"
	"github.com/uber/jaeger-client-go/thrift-gen/jaeger"
	"github.com/uber/jaeger-client-go/thrift-gen/zipkincore"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = jaeger.GoUnusedProtection__
var _ = zipkincore.GoUnusedProtection__

type Agent interface {
	// Parameters:
	//  - Spans
	EmitZipkinBatch(spans []*zipkincore.Span) (err error)
	// Parameters:
	//  - Batch
	EmitBatch(batch *jaeger.Batch) (err error)
}

type AgentClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewAgentClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AgentClient {
	return &AgentClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewAgentClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AgentClient {
	return &AgentClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Spans
func (p *AgentClient) EmitZipkinBatch(spans []*zipkincore.Span) (err error) {
	if err = p.sendEmitZipkinBatch(spans); err != nil {
		return
	}
	return
}

func (p *AgentClient) sendEmitZipkinBatch(spans []*zipkincore.Span) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("emitZipkinBatch", thrift.ONEWAY, p.SeqId); err != nil {
		return
	}
	args := AgentEmitZipkinBatchArgs{
		Spans: spans,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

// Parameters:
//  - Batch
func (p *AgentClient) EmitBatch(batch *jaeger.Batch) (err error) {
	if err = p.sendEmitBatch(batch); err != nil {
		return
	}
	return
}

func (p *AgentClient) sendEmitBatch(batch *jaeger.Batch) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("emitBatch", thrift.ONEWAY, p.SeqId); err != nil {
		return
	}
	args := AgentEmitBatchArgs{
		Batch: batch,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

type AgentProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Agent
}

func (p *AgentProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AgentProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AgentProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAgentProcessor(handler Agent) *AgentProcessor {

	self0 := &AgentProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self0.processorMap["emitZipkinBatch"] = &agentProcessorEmitZipkinBatch{handler: handler}
	self0.processorMap["emitBatch"] = &agentProcessorEmitBatch{handler: handler}
	return self0
}

func (p *AgentProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x1 := thrift.NewTApplicationException(thrift.UNKNOWN_MSEVOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x1.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x1

}

type agentProcessorEmitZipkinBatch struct {
	handler Agent
}

func (p *agentProcessorEmitZipkinBatch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AgentEmitZipkinBatchArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	if err2 = p.handler.EmitZipkinBatch(args.Spans); err2 != nil {
		return true, err2
	}
	return true, nil
}

type agentProcessorEmitBatch struct {
	handler Agent
}

func (p *agentProcessorEmitBatch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AgentEmitBatchArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	if err2 = p.handler.EmitBatch(args.Batch); err2 != nil {
		return true, err2
	}
	return true, nil
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Spans
type AgentEmitZipkinBatchArgs struct {
	Spans []*zipkincore.Span `thrift:"spans,1" json:"spans"`
}

func NewAgentEmitZipkinBatchArgs() *AgentEmitZipkinBatchArgs {
	return &AgentEmitZipkinBatchArgs{}
}

func (p *AgentEmitZipkinBatchArgs) GetSpans() []*zipkincore.Span {
	return p.Spans
}
func (p *AgentEmitZipkinBatchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AgentEmitZipkinBatchArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*zipkincore.Span, 0, size)
	p.Spans = tSlice
	for i := 0; i < size; i++ {
		_elem2 := &zipkincore.Span{}
		if err := _elem2.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem2), err)
		}
		p.Spans = append(p.Spans, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *AgentEmitZipkinBatchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("emitZipkinBatch_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AgentEmitZipkinBatchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("spans", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:spans: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Spans)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Spans {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:spans: ", p), err)
	}
	return err
}

func (p *AgentEmitZipkinBatchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AgentEmitZipkinBatchArgs(%+v)", *p)
}

// Attributes:
//  - Batch
type AgentEmitBatchArgs struct {
	Batch *jaeger.Batch `thrift:"batch,1" json:"batch"`
}

func NewAgentEmitBatchArgs() *AgentEmitBatchArgs {
	return &AgentEmitBatchArgs{}
}

var AgentEmitBatchArgs_Batch_DEFAULT *jaeger.Batch

func (p *AgentEmitBatchArgs) GetBatch() *jaeger.Batch {
	if !p.IsSetBatch() {
		return AgentEmitBatchArgs_Batch_DEFAULT
	}
	return p.Batch
}
func (p *AgentEmitBatchArgs) IsSetBatch() bool {
	return p.Batch != nil
}

func (p *AgentEmitBatchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AgentEmitBatchArgs) readField1(iprot thrift.TProtocol) error {
	p.Batch = &jaeger.Batch{}
	if err := p.Batch.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Batch), err)
	}
	return nil
}

func (p *AgentEmitBatchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("emitBatch_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AgentEmitBatchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("batch", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:batch: ", p), err)
	}
	if err := p.Batch.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Batch), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:batch: ", p), err)
	}
	return err
}

func (p *AgentEmitBatchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AgentEmitBatchArgs(%+v)", *p)
}
