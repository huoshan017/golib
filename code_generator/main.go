package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/huoshan017/mysql-go/generate"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "args num not enough\n")
		return
	}

	arg_config_file := flag.String("c", "", "config file path")
	arg_dest_path := flag.String("d", "", "dest source path")
	arg_protoc_path := flag.String("p", "", "protoc file path")
	flag.Parse()

	var config_path string
	if nil != arg_config_file {
		config_path = *arg_config_file
		fmt.Fprintf(os.Stdout, "config file path %v\n", config_path)
	} else {
		fmt.Fprintf(os.Stderr, "not found config file arg\n")
		return
	}

	var dest_path string
	if nil != arg_dest_path {
		dest_path = *arg_dest_path
		fmt.Fprintf(os.Stdout, "dest path %v\n", dest_path)
	} else {
		fmt.Fprintf(os.Stderr, "not found dest path arg\n")
		return
	}

	var protoc_path string
	if nil != arg_protoc_path {
		protoc_path = *arg_protoc_path
		fmt.Fprintf(os.Stdout, "protoc path %v\n", protoc_path)
	} else {
		fmt.Fprintf(os.Stderr, "not found dest proto arg\n")
		return
	}

	var config_loader mysql_generate.ConfigLoader
	if !config_loader.Load(config_path) {
		return
	}

	if !config_loader.Generate(dest_path) {
		return
	}

	fmt.Fprintf(os.Stdout, "generated source\n")

	proto_dest_path, config_file := path.Split(config_path)
	proto_file := strings.Replace(config_file, "json", "proto", -1)
	fmt.Fprintf(os.Stdout, "proto_dest_path: %v    proto_file: %v\n", proto_dest_path, proto_file)
	if !config_loader.GenerateFieldStructsProto(proto_dest_path + proto_file) {
		return
	}

	fmt.Fprintf(os.Stdout, "generated proto\n")

	cmd := exec.Command(protoc_path, "--go_out", dest_path+"/"+config_loader.DBPkg, "--proto_path", proto_dest_path, proto_file)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cmd run err: %v\n", err.Error())
		return
	}
	fmt.Printf("%s", out.String())

	if !config_loader.GenerateInitFunc(dest_path) {
		return
	}

	fmt.Fprintf(os.Stdout, "generated init funcs\ngenerated all\n")
}
