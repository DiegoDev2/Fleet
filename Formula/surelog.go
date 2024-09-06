package main

// Surelog - SystemVerilog Pre-processor, parser, elaborator, UHDM compiler
// Homepage: https://github.com/chipsalliance/Surelog

import (
	"fmt"
	
	"os/exec"
)

func installSurelog() {
	// Método 1: Descargar y extraer .tar.gz
	surelog_tar_url := "https://github.com/chipsalliance/Surelog/archive/refs/tags/v1.84.tar.gz"
	surelog_cmd_tar := exec.Command("curl", "-L", surelog_tar_url, "-o", "package.tar.gz")
	err := surelog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	surelog_zip_url := "https://github.com/chipsalliance/Surelog/archive/refs/tags/v1.84.zip"
	surelog_cmd_zip := exec.Command("curl", "-L", surelog_zip_url, "-o", "package.zip")
	err = surelog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	surelog_bin_url := "https://github.com/chipsalliance/Surelog/archive/refs/tags/v1.84.bin"
	surelog_cmd_bin := exec.Command("curl", "-L", surelog_bin_url, "-o", "binary.bin")
	err = surelog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	surelog_src_url := "https://github.com/chipsalliance/Surelog/archive/refs/tags/v1.84.src.tar.gz"
	surelog_cmd_src := exec.Command("curl", "-L", surelog_src_url, "-o", "source.tar.gz")
	err = surelog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	surelog_cmd_direct := exec.Command("./binary")
	err = surelog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: antlr")
	exec.Command("latte", "install", "antlr").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: nlohmann-json")
	exec.Command("latte", "install", "nlohmann-json").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: antlr4-cpp-runtime")
	exec.Command("latte", "install", "antlr4-cpp-runtime").Run()
	fmt.Println("Instalando dependencia: capnp")
	exec.Command("latte", "install", "capnp").Run()
	fmt.Println("Instalando dependencia: uhdm")
	exec.Command("latte", "install", "uhdm").Run()
}
