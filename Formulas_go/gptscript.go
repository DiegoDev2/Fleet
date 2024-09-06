package main

// Gptscript - Develop LLM Apps in Natural Language
// Homepage: https://gptscript.ai

import (
	"fmt"
	
	"os/exec"
)

func installGptscript() {
	// Método 1: Descargar y extraer .tar.gz
	gptscript_tar_url := "https://github.com/gptscript-ai/gptscript/archive/refs/tags/v0.9.4.tar.gz"
	gptscript_cmd_tar := exec.Command("curl", "-L", gptscript_tar_url, "-o", "package.tar.gz")
	err := gptscript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gptscript_zip_url := "https://github.com/gptscript-ai/gptscript/archive/refs/tags/v0.9.4.zip"
	gptscript_cmd_zip := exec.Command("curl", "-L", gptscript_zip_url, "-o", "package.zip")
	err = gptscript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gptscript_bin_url := "https://github.com/gptscript-ai/gptscript/archive/refs/tags/v0.9.4.bin"
	gptscript_cmd_bin := exec.Command("curl", "-L", gptscript_bin_url, "-o", "binary.bin")
	err = gptscript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gptscript_src_url := "https://github.com/gptscript-ai/gptscript/archive/refs/tags/v0.9.4.src.tar.gz"
	gptscript_cmd_src := exec.Command("curl", "-L", gptscript_src_url, "-o", "source.tar.gz")
	err = gptscript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gptscript_cmd_direct := exec.Command("./binary")
	err = gptscript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
