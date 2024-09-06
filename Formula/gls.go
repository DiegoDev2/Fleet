package main

// GLs - Powerful and cross-platform ls
// Homepage: https://g.equationzhao.space

import (
	"fmt"
	
	"os/exec"
)

func installGLs() {
	// Método 1: Descargar y extraer .tar.gz
	gls_tar_url := "https://github.com/Equationzhao/g/archive/refs/tags/v0.29.0.tar.gz"
	gls_cmd_tar := exec.Command("curl", "-L", gls_tar_url, "-o", "package.tar.gz")
	err := gls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gls_zip_url := "https://github.com/Equationzhao/g/archive/refs/tags/v0.29.0.zip"
	gls_cmd_zip := exec.Command("curl", "-L", gls_zip_url, "-o", "package.zip")
	err = gls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gls_bin_url := "https://github.com/Equationzhao/g/archive/refs/tags/v0.29.0.bin"
	gls_cmd_bin := exec.Command("curl", "-L", gls_bin_url, "-o", "binary.bin")
	err = gls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gls_src_url := "https://github.com/Equationzhao/g/archive/refs/tags/v0.29.0.src.tar.gz"
	gls_cmd_src := exec.Command("curl", "-L", gls_src_url, "-o", "source.tar.gz")
	err = gls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gls_cmd_direct := exec.Command("./binary")
	err = gls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
