package main

// Chroma - General purpose syntax highlighter in pure Go
// Homepage: https://github.com/alecthomas/chroma

import (
	"fmt"
	
	"os/exec"
)

func installChroma() {
	// Método 1: Descargar y extraer .tar.gz
	chroma_tar_url := "https://github.com/alecthomas/chroma/archive/refs/tags/v2.14.0.tar.gz"
	chroma_cmd_tar := exec.Command("curl", "-L", chroma_tar_url, "-o", "package.tar.gz")
	err := chroma_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chroma_zip_url := "https://github.com/alecthomas/chroma/archive/refs/tags/v2.14.0.zip"
	chroma_cmd_zip := exec.Command("curl", "-L", chroma_zip_url, "-o", "package.zip")
	err = chroma_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chroma_bin_url := "https://github.com/alecthomas/chroma/archive/refs/tags/v2.14.0.bin"
	chroma_cmd_bin := exec.Command("curl", "-L", chroma_bin_url, "-o", "binary.bin")
	err = chroma_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chroma_src_url := "https://github.com/alecthomas/chroma/archive/refs/tags/v2.14.0.src.tar.gz"
	chroma_cmd_src := exec.Command("curl", "-L", chroma_src_url, "-o", "source.tar.gz")
	err = chroma_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chroma_cmd_direct := exec.Command("./binary")
	err = chroma_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
