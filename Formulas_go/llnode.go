package main

// Llnode - LLDB plugin for live/post-mortem debugging of node.js apps
// Homepage: https://github.com/nodejs/llnode

import (
	"fmt"
	
	"os/exec"
)

func installLlnode() {
	// Método 1: Descargar y extraer .tar.gz
	llnode_tar_url := "https://github.com/nodejs/llnode/archive/refs/tags/v4.0.0.tar.gz"
	llnode_cmd_tar := exec.Command("curl", "-L", llnode_tar_url, "-o", "package.tar.gz")
	err := llnode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	llnode_zip_url := "https://github.com/nodejs/llnode/archive/refs/tags/v4.0.0.zip"
	llnode_cmd_zip := exec.Command("curl", "-L", llnode_zip_url, "-o", "package.zip")
	err = llnode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	llnode_bin_url := "https://github.com/nodejs/llnode/archive/refs/tags/v4.0.0.bin"
	llnode_cmd_bin := exec.Command("curl", "-L", llnode_bin_url, "-o", "binary.bin")
	err = llnode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	llnode_src_url := "https://github.com/nodejs/llnode/archive/refs/tags/v4.0.0.src.tar.gz"
	llnode_cmd_src := exec.Command("curl", "-L", llnode_src_url, "-o", "source.tar.gz")
	err = llnode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	llnode_cmd_direct := exec.Command("./binary")
	err = llnode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
