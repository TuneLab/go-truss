package templates

const bbb2306HandlerMethods = `
{{ with $te := .}}
		{{range $i := .Methods}}
		// {{.Name}} implements Service.
		func (s {{ToLower $te.ServiceName}}Service) {{.Name}}(ctx context.Context, in *pb.{{GoName .RequestType.Name}}) (*pb.{{GoName .ResponseType.Name}}, error){
			var resp pb.{{GoName .ResponseType.Name}}
			resp = pb.{{GoName .ResponseType.Name}}{
				{{range $j := $i.ResponseType.Message.Fields -}}
					// {{GoName $j.Name}}:
				{{end -}}
			}
			return &resp, nil
		}
		{{end}}
{{- end}}
`

const bbb2306Handlers = `
package handlers

import (
	"golang.org/x/net/context"

	pb "{{.PBImportPath -}}"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.{{GoName .Service.Name}}Server {
	return {{ToLower .Service.Name}}Service{}
}

type {{ToLower .Service.Name}}Service struct{}

{{with $te := . }}
	{{range $i := $te.Service.Methods}}
		// {{$i.Name}} implements Service.
		func (s {{ToLower $te.Service.Name}}Service) {{$i.Name}}(ctx context.Context, in *pb.{{GoName $i.RequestType.Name}}) (*pb.{{GoName $i.ResponseType.Name}}, error){
			var resp pb.{{GoName $i.ResponseType.Name}}
			resp = pb.{{GoName $i.ResponseType.Name}}{
				{{range $j := $i.ResponseType.Message.Fields -}}
					// {{GoName $j.Name}}:
				{{end -}}
			}
			return &resp, nil
		}
	{{end}}
{{- end}}
`

const v050Handlers = `
package handlers

import (
	"context"

	pb "{{.PBImportPath -}}"
)

// {{GoName .Service.Name}}Servicer describes a service that {{ToLower .Service.Name}}s
type {{GoName .Service.Name}}Servicer interface {
	{{with $te := . }}
		{{range $i := $te.Service.Methods}}
			{{$i.Name}}(ctx context.Context, in *pb.{{GoName $i.RequestType.Name}}) (*pb.{{GoName $i.ResponseType.Name}}, error)
		{{end}}
	{{- end}}
}

// NewService returns a naïve, stateless implementation of Service.
func NewService() {{GoName .Service.Name}}Servicer {
	return {{ToLower .Service.Name}}Service{}
}

type {{ToLower .Service.Name}}Service struct{}

{{with $te := . }}
	{{range $i := $te.Service.Methods}}
		// {{$i.Name}} implements Service.
		func (s {{ToLower $te.Service.Name}}Service) {{$i.Name}}(ctx context.Context, in *pb.{{GoName $i.RequestType.Name}}) (*pb.{{GoName $i.ResponseType.Name}}, error){
			var resp pb.{{GoName $i.ResponseType.Name}}
			resp = pb.{{GoName $i.ResponseType.Name}}{
				{{range $j := $i.ResponseType.Message.Fields -}}
					// {{GoName $j.Name}}:
				{{end -}}
			}
			return &resp, nil
		}
	{{end}}
{{- end}}
`

var Handlers = map[string]map[string]string{
	"bbb2306": {
		"HandlerMethods": bbb2306HandlerMethods,
		"Handlers":       bbb2306Handlers,
	},
	"v0.5.0": {
		"HandlerMethods": bbb2306HandlerMethods,
		"Handlers":       v050Handlers,
	},
}
