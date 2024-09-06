package main

// Thriftgo - Implementation of thrift compiler in go language with plugin mechanism
// Homepage: https://github.com/cloudwego/thriftgo

import (
	"fmt"
	
	"os/exec"
)

func installThriftgo() {
	// Método 1: Descargar y extraer .tar.gz
	thriftgo_tar_url := "https://github.com/cloudwego/thriftgo/archive/refs/tags/v0.3.17.tar.gz"
	thriftgo_cmd_tar := exec.Command("curl", "-L", thriftgo_tar_url, "-o", "package.tar.gz")
	err := thriftgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	thriftgo_zip_url := "https://github.com/cloudwego/thriftgo/archive/refs/tags/v0.3.17.zip"
	thriftgo_cmd_zip := exec.Command("curl", "-L", thriftgo_zip_url, "-o", "package.zip")
	err = thriftgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	thriftgo_bin_url := "https://github.com/cloudwego/thriftgo/archive/refs/tags/v0.3.17.bin"
	thriftgo_cmd_bin := exec.Command("curl", "-L", thriftgo_bin_url, "-o", "binary.bin")
	err = thriftgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	thriftgo_src_url := "https://github.com/cloudwego/thriftgo/archive/refs/tags/v0.3.17.src.tar.gz"
	thriftgo_cmd_src := exec.Command("curl", "-L", thriftgo_src_url, "-o", "source.tar.gz")
	err = thriftgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	thriftgo_cmd_direct := exec.Command("./binary")
	err = thriftgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
