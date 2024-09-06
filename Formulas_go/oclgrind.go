package main

// Oclgrind - OpenCL device simulator and debugger
// Homepage: https://github.com/jrprice/Oclgrind

import (
	"fmt"
	
	"os/exec"
)

func installOclgrind() {
	// Método 1: Descargar y extraer .tar.gz
	oclgrind_tar_url := "https://github.com/jrprice/Oclgrind/archive/refs/tags/v21.10.tar.gz"
	oclgrind_cmd_tar := exec.Command("curl", "-L", oclgrind_tar_url, "-o", "package.tar.gz")
	err := oclgrind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oclgrind_zip_url := "https://github.com/jrprice/Oclgrind/archive/refs/tags/v21.10.zip"
	oclgrind_cmd_zip := exec.Command("curl", "-L", oclgrind_zip_url, "-o", "package.zip")
	err = oclgrind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oclgrind_bin_url := "https://github.com/jrprice/Oclgrind/archive/refs/tags/v21.10.bin"
	oclgrind_cmd_bin := exec.Command("curl", "-L", oclgrind_bin_url, "-o", "binary.bin")
	err = oclgrind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oclgrind_src_url := "https://github.com/jrprice/Oclgrind/archive/refs/tags/v21.10.src.tar.gz"
	oclgrind_cmd_src := exec.Command("curl", "-L", oclgrind_src_url, "-o", "source.tar.gz")
	err = oclgrind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oclgrind_cmd_direct := exec.Command("./binary")
	err = oclgrind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm@14")
exec.Command("latte", "install", "llvm@14")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
}
