package main

// Vgrep - User-friendly pager for grep
// Homepage: https://github.com/vrothberg/vgrep

import (
	"fmt"
	
	"os/exec"
)

func installVgrep() {
	// Método 1: Descargar y extraer .tar.gz
	vgrep_tar_url := "https://github.com/vrothberg/vgrep/archive/refs/tags/v2.8.0.tar.gz"
	vgrep_cmd_tar := exec.Command("curl", "-L", vgrep_tar_url, "-o", "package.tar.gz")
	err := vgrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vgrep_zip_url := "https://github.com/vrothberg/vgrep/archive/refs/tags/v2.8.0.zip"
	vgrep_cmd_zip := exec.Command("curl", "-L", vgrep_zip_url, "-o", "package.zip")
	err = vgrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vgrep_bin_url := "https://github.com/vrothberg/vgrep/archive/refs/tags/v2.8.0.bin"
	vgrep_cmd_bin := exec.Command("curl", "-L", vgrep_bin_url, "-o", "binary.bin")
	err = vgrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vgrep_src_url := "https://github.com/vrothberg/vgrep/archive/refs/tags/v2.8.0.src.tar.gz"
	vgrep_cmd_src := exec.Command("curl", "-L", vgrep_src_url, "-o", "source.tar.gz")
	err = vgrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vgrep_cmd_direct := exec.Command("./binary")
	err = vgrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: go-md2man")
exec.Command("latte", "install", "go-md2man")
}
