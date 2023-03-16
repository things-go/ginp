package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/things-go/dyn/internal/infra"
)

const deprecationComment = "// Deprecated: Do not use."

var (
	errorsPackage  = protogen.GoImportPath("errors")
	contextPackage = protogen.GoImportPath("context")
	ginPackage     = protogen.GoImportPath("github.com/gin-gonic/gin")
	netHttpPackage = protogen.GoImportPath("net/http")
)

var methodSets = make(map[string]int)

func runProtoGen(gen *protogen.Plugin) error {
	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		generateFile(gen, f, *omitempty)
	}
	return nil
}

// generateFile generates a .gin.pb.go file.
func generateFile(gen *protogen.Plugin, file *protogen.File, omitempty bool) *protogen.GeneratedFile {
	if len(file.Services) == 0 || (omitempty && !hasHTTPRule(file.Services)) {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + ".gin.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-gin. DO NOT EDIT.")
	g.P("// versions:")
	g.P("//   - protoc-gen-go-gin ", version)
	g.P("//   - protoc            ", infra.ProtocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g, omitempty)
	return g
}

// generateFileContent generates the errors definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, omitempty bool) {
	if len(file.Services) == 0 {
		return
	}
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible.")
	g.P("var _ = ", errorsPackage.Ident("New"))
	g.P("var _ = ", contextPackage.Ident("TODO"))
	g.P("var _ = ", ginPackage.Ident("New"))
	if *useEncoding {
		g.P("var _ = ", netHttpPackage.Ident("HandleFunc"))
	}
	g.P()

	for _, service := range file.Services {
		genService(gen, file, g, service, omitempty)
	}
}

func genService(gen *protogen.Plugin, file *protogen.File,
	g *protogen.GeneratedFile, service *protogen.Service, omitempty bool) {
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	// HTTP Server.
	sd := &serviceDesc{
		ServiceType:       service.GoName,
		ServiceName:       string(service.Desc.FullName()),
		Metadata:          file.Desc.Path(),
		UseCustomResponse: *useCustomResponse,
		UseEncoding:       *useEncoding,
	}
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		rule, ok := proto.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
		if rule != nil && ok {
			for _, bind := range rule.AdditionalBindings {
				sd.Methods = append(sd.Methods, buildHTTPRule(g, method, bind))
			}
			sd.Methods = append(sd.Methods, buildHTTPRule(g, method, rule))
		} else if !omitempty {
			path := fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())
			sd.Methods = append(sd.Methods, buildMethodDesc(g, method, "POST", path))
		}
	}
	if len(sd.Methods) == 0 {
		return
	}
	err := sd.execute(g)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr,
			"\u001B[31mWARN\u001B[m: execute template failed.\n")
	}
}

func hasHTTPRule(services []*protogen.Service) bool {
	for _, service := range services {
		for _, method := range service.Methods {
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				continue
			}
			rule, ok := proto.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
			if rule != nil && ok {
				return true
			}
		}
	}
	return false
}

func buildHTTPRule(g *protogen.GeneratedFile, m *protogen.Method, rule *annotations.HttpRule) *methodDesc {
	var (
		path         string
		method       string
		body         string
		responseBody string
	)
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		path = pattern.Get
		method = http.MethodGet
	case *annotations.HttpRule_Put:
		path = pattern.Put
		method = http.MethodPut
	case *annotations.HttpRule_Post:
		path = pattern.Post
		method = http.MethodPost
	case *annotations.HttpRule_Delete:
		path = pattern.Delete
		method = http.MethodDelete
	case *annotations.HttpRule_Patch:
		path = pattern.Patch
		method = http.MethodPatch
	case *annotations.HttpRule_Custom:
		path = pattern.Custom.Path
		method = pattern.Custom.Kind
	}
	body = rule.Body
	responseBody = rule.ResponseBody
	md := buildMethodDesc(g, m, method, path)
	switch {
	case method == http.MethodGet:
		if body != "" {
			_, _ = fmt.Fprintf(os.Stderr,
				"\u001B[31mWARN\u001B[m: %s %s body should not be declared.\n", method, path)
		}
		md.HasBody = false
	case method == http.MethodDelete:
		if body != "" {
			md.HasBody = true
			if !*allowDeleteBody {
				md.HasBody = false
				_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: %s %s body should not be declared.\n", method, path)
			}
		} else {
			md.HasBody = false
		}
	case method == http.MethodPatch:
		if body != "" {
			md.HasBody = true
		} else {
			md.HasBody = false
			if !*allowEmptyPatchBody {
				_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: %s %s is does not declare a body.\n", method, path)
			}
		}
	case body == "*":
		md.HasBody = true
		md.Body = ""
	case body != "":
		md.HasBody = true
		md.Body = "." + camelCaseVars(body)
	default:
		md.HasBody = false
		_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: %s %s is does not declare a body.\n", method, path)
	}
	if responseBody == "*" {
		md.ResponseBody = ""
	} else if responseBody != "" {
		md.ResponseBody = "." + camelCaseVars(responseBody)
	}
	return md
}

func buildMethodDesc(g *protogen.GeneratedFile, m *protogen.Method, method, path string) *methodDesc {
	defer func() { methodSets[m.GoName]++ }()
	vars := buildPathVars(m, path)
	fields := m.Input.Desc.Fields()
	for _, v := range vars {
		for _, field := range strings.Split(v, ".") {
			if strings.TrimSpace(field) == "" {
				continue
			}
			if strings.Contains(field, ":") {
				field = strings.Split(field, ":")[0]
			}
			fd := fields.ByName(protoreflect.Name(field))
			if fd == nil {
				// nolint: lll
				fmt.Fprintf(os.Stderr, "\u001B[31mERROR\u001B[m: The corresponding field '%s' declaration in message could not be found in '%s'\n", v, path)
				os.Exit(2) // nolint: gocritic
			}
			switch {
			case fd.IsMap():
				fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: The field in path:'%s' shouldn't be a map.\n", v)
			case fd.IsList():
				fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: The field in path:'%s' shouldn't be a list.\n", v)
			case fd.Kind() == protoreflect.MessageKind || fd.Kind() == protoreflect.GroupKind:
				fields = fd.Message().Fields()
			}
		}
	}
	comment := m.Comments.Leading.String() + m.Comments.Trailing.String()
	if comment != "" {
		comment = "// " + m.GoName + strings.TrimPrefix(strings.TrimSuffix(comment, "\n"), "//")
	} else {
		comment = "// " + m.GoName + " ..."
	}
	return &methodDesc{
		Name:    m.GoName,
		Num:     methodSets[m.GoName],
		Request: g.QualifiedGoIdent(m.Input.GoIdent),
		Reply:   g.QualifiedGoIdent(m.Output.GoIdent),
		Comment: comment,
		Path:    transformPathParams(path),
		Method:  method,
		HasVars: len(vars) > 0,
	}
}

// transformPathParams 路由路由 {xx} --> :xx
func transformPathParams(path string) string {
	paths := strings.Split(path, "/")
	for i, p := range paths {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") || strings.HasPrefix(p, ":") {
			paths[i] = ":" + p[1:len(p)-1]
		}
	}
	return strings.Join(paths, "/")
}

func buildPathVars(_ *protogen.Method, path string) (res []string) {
	for _, v := range strings.Split(path, "/") {
		if strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}") {
			res = append(res, strings.TrimSuffix(strings.TrimPrefix(v, "{"), "}"))
		}
	}
	return
}

func camelCaseVars(s string) string {
	vars := make([]string, 0)
	subs := strings.Split(s, ".")
	for _, sub := range subs {
		vars = append(vars, infra.CamelCase(sub))
	}
	return strings.Join(vars, ".")
}
