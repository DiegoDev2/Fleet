package main

// OpentracingCpp - OpenTracing API for C++
// Homepage: https://opentracing.io/

import (
	"fmt"
	
	"os/exec"
)

func installOpentracingCpp() {
	// Método 1: Descargar y extraer .tar.gz
	opentracingcpp_tar_url := "https://github.com/opentracing/opentracing-cpp/archive/refs/tags/v1.6.0.tar.gz"
	opentracingcpp_cmd_tar := exec.Command("curl", "-L", opentracingcpp_tar_url, "-o", "package.tar.gz")
	err := opentracingcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opentracingcpp_zip_url := "https://github.com/opentracing/opentracing-cpp/archive/refs/tags/v1.6.0.zip"
	opentracingcpp_cmd_zip := exec.Command("curl", "-L", opentracingcpp_zip_url, "-o", "package.zip")
	err = opentracingcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opentracingcpp_bin_url := "https://github.com/opentracing/opentracing-cpp/archive/refs/tags/v1.6.0.bin"
	opentracingcpp_cmd_bin := exec.Command("curl", "-L", opentracingcpp_bin_url, "-o", "binary.bin")
	err = opentracingcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opentracingcpp_src_url := "https://github.com/opentracing/opentracing-cpp/archive/refs/tags/v1.6.0.src.tar.gz"
	opentracingcpp_cmd_src := exec.Command("curl", "-L", opentracingcpp_src_url, "-o", "source.tar.gz")
	err = opentracingcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opentracingcpp_cmd_direct := exec.Command("./binary")
	err = opentracingcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
