package main

// Kuzu - Embeddable graph database management system built for query speed & scalability
// Homepage: https://kuzudb.com/

import (
	"fmt"
	
	"os/exec"
)

func installKuzu() {
	// Método 1: Descargar y extraer .tar.gz
	kuzu_tar_url := "https://github.com/kuzudb/kuzu/archive/refs/tags/v0.6.0.tar.gz"
	kuzu_cmd_tar := exec.Command("curl", "-L", kuzu_tar_url, "-o", "package.tar.gz")
	err := kuzu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kuzu_zip_url := "https://github.com/kuzudb/kuzu/archive/refs/tags/v0.6.0.zip"
	kuzu_cmd_zip := exec.Command("curl", "-L", kuzu_zip_url, "-o", "package.zip")
	err = kuzu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kuzu_bin_url := "https://github.com/kuzudb/kuzu/archive/refs/tags/v0.6.0.bin"
	kuzu_cmd_bin := exec.Command("curl", "-L", kuzu_bin_url, "-o", "binary.bin")
	err = kuzu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kuzu_src_url := "https://github.com/kuzudb/kuzu/archive/refs/tags/v0.6.0.src.tar.gz"
	kuzu_cmd_src := exec.Command("curl", "-L", kuzu_src_url, "-o", "source.tar.gz")
	err = kuzu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kuzu_cmd_direct := exec.Command("./binary")
	err = kuzu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
