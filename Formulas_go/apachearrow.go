package main

// ApacheArrow - Columnar in-memory analytics layer designed to accelerate big data
// Homepage: https://arrow.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installApacheArrow() {
	// Método 1: Descargar y extraer .tar.gz
	apachearrow_tar_url := "https://www.apache.org/dyn/closer.lua?path=arrow/arrow-17.0.0/apache-arrow-17.0.0.tar.gz"
	apachearrow_cmd_tar := exec.Command("curl", "-L", apachearrow_tar_url, "-o", "package.tar.gz")
	err := apachearrow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apachearrow_zip_url := "https://www.apache.org/dyn/closer.lua?path=arrow/arrow-17.0.0/apache-arrow-17.0.0.zip"
	apachearrow_cmd_zip := exec.Command("curl", "-L", apachearrow_zip_url, "-o", "package.zip")
	err = apachearrow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apachearrow_bin_url := "https://www.apache.org/dyn/closer.lua?path=arrow/arrow-17.0.0/apache-arrow-17.0.0.bin"
	apachearrow_cmd_bin := exec.Command("curl", "-L", apachearrow_bin_url, "-o", "binary.bin")
	err = apachearrow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apachearrow_src_url := "https://www.apache.org/dyn/closer.lua?path=arrow/arrow-17.0.0/apache-arrow-17.0.0.src.tar.gz"
	apachearrow_cmd_src := exec.Command("curl", "-L", apachearrow_src_url, "-o", "source.tar.gz")
	err = apachearrow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apachearrow_cmd_direct := exec.Command("./binary")
	err = apachearrow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: aws-sdk-cpp")
exec.Command("latte", "install", "aws-sdk-cpp")
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
	fmt.Println("Instalando dependencia: bzip2")
exec.Command("latte", "install", "bzip2")
	fmt.Println("Instalando dependencia: c-ares")
exec.Command("latte", "install", "c-ares")
	fmt.Println("Instalando dependencia: glog")
exec.Command("latte", "install", "glog")
	fmt.Println("Instalando dependencia: grpc")
exec.Command("latte", "install", "grpc")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: rapidjson")
exec.Command("latte", "install", "rapidjson")
	fmt.Println("Instalando dependencia: re2")
exec.Command("latte", "install", "re2")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
	fmt.Println("Instalando dependencia: thrift")
exec.Command("latte", "install", "thrift")
	fmt.Println("Instalando dependencia: utf8proc")
exec.Command("latte", "install", "utf8proc")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
