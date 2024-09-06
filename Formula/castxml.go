package main

// Castxml - C-family Abstract Syntax Tree XML Output
// Homepage: https://github.com/CastXML/CastXML

import (
	"fmt"
	
	"os/exec"
)

func installCastxml() {
	// Método 1: Descargar y extraer .tar.gz
	castxml_tar_url := "https://github.com/CastXML/CastXML/archive/refs/tags/v0.6.8.tar.gz"
	castxml_cmd_tar := exec.Command("curl", "-L", castxml_tar_url, "-o", "package.tar.gz")
	err := castxml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	castxml_zip_url := "https://github.com/CastXML/CastXML/archive/refs/tags/v0.6.8.zip"
	castxml_cmd_zip := exec.Command("curl", "-L", castxml_zip_url, "-o", "package.zip")
	err = castxml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	castxml_bin_url := "https://github.com/CastXML/CastXML/archive/refs/tags/v0.6.8.bin"
	castxml_cmd_bin := exec.Command("curl", "-L", castxml_bin_url, "-o", "binary.bin")
	err = castxml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	castxml_src_url := "https://github.com/CastXML/CastXML/archive/refs/tags/v0.6.8.src.tar.gz"
	castxml_cmd_src := exec.Command("curl", "-L", castxml_src_url, "-o", "source.tar.gz")
	err = castxml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	castxml_cmd_direct := exec.Command("./binary")
	err = castxml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
