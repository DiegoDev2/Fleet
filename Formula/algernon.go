package main

// Algernon - Pure Go web server with Lua, Markdown, HTTP/2 and template support
// Homepage: https://github.com/xyproto/algernon

import (
	"fmt"
	
	"os/exec"
)

func installAlgernon() {
	// Método 1: Descargar y extraer .tar.gz
	algernon_tar_url := "https://github.com/xyproto/algernon/archive/refs/tags/v1.17.1.tar.gz"
	algernon_cmd_tar := exec.Command("curl", "-L", algernon_tar_url, "-o", "package.tar.gz")
	err := algernon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	algernon_zip_url := "https://github.com/xyproto/algernon/archive/refs/tags/v1.17.1.zip"
	algernon_cmd_zip := exec.Command("curl", "-L", algernon_zip_url, "-o", "package.zip")
	err = algernon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	algernon_bin_url := "https://github.com/xyproto/algernon/archive/refs/tags/v1.17.1.bin"
	algernon_cmd_bin := exec.Command("curl", "-L", algernon_bin_url, "-o", "binary.bin")
	err = algernon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	algernon_src_url := "https://github.com/xyproto/algernon/archive/refs/tags/v1.17.1.src.tar.gz"
	algernon_cmd_src := exec.Command("curl", "-L", algernon_src_url, "-o", "source.tar.gz")
	err = algernon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	algernon_cmd_direct := exec.Command("./binary")
	err = algernon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
