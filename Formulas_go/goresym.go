package main

// Goresym - Go symbol recovery tool
// Homepage: https://github.com/mandiant/GoReSym

import (
	"fmt"
	
	"os/exec"
)

func installGoresym() {
	// Método 1: Descargar y extraer .tar.gz
	goresym_tar_url := "https://github.com/mandiant/GoReSym/archive/refs/tags/v2.7.4.tar.gz"
	goresym_cmd_tar := exec.Command("curl", "-L", goresym_tar_url, "-o", "package.tar.gz")
	err := goresym_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goresym_zip_url := "https://github.com/mandiant/GoReSym/archive/refs/tags/v2.7.4.zip"
	goresym_cmd_zip := exec.Command("curl", "-L", goresym_zip_url, "-o", "package.zip")
	err = goresym_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goresym_bin_url := "https://github.com/mandiant/GoReSym/archive/refs/tags/v2.7.4.bin"
	goresym_cmd_bin := exec.Command("curl", "-L", goresym_bin_url, "-o", "binary.bin")
	err = goresym_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goresym_src_url := "https://github.com/mandiant/GoReSym/archive/refs/tags/v2.7.4.src.tar.gz"
	goresym_cmd_src := exec.Command("curl", "-L", goresym_src_url, "-o", "source.tar.gz")
	err = goresym_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goresym_cmd_direct := exec.Command("./binary")
	err = goresym_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
