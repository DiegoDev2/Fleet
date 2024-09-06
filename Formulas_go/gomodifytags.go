package main

// Gomodifytags - Go tool to modify struct field tags
// Homepage: https://github.com/fatih/gomodifytags

import (
	"fmt"
	
	"os/exec"
)

func installGomodifytags() {
	// Método 1: Descargar y extraer .tar.gz
	gomodifytags_tar_url := "https://github.com/fatih/gomodifytags/archive/refs/tags/v1.17.0.tar.gz"
	gomodifytags_cmd_tar := exec.Command("curl", "-L", gomodifytags_tar_url, "-o", "package.tar.gz")
	err := gomodifytags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gomodifytags_zip_url := "https://github.com/fatih/gomodifytags/archive/refs/tags/v1.17.0.zip"
	gomodifytags_cmd_zip := exec.Command("curl", "-L", gomodifytags_zip_url, "-o", "package.zip")
	err = gomodifytags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gomodifytags_bin_url := "https://github.com/fatih/gomodifytags/archive/refs/tags/v1.17.0.bin"
	gomodifytags_cmd_bin := exec.Command("curl", "-L", gomodifytags_bin_url, "-o", "binary.bin")
	err = gomodifytags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gomodifytags_src_url := "https://github.com/fatih/gomodifytags/archive/refs/tags/v1.17.0.src.tar.gz"
	gomodifytags_cmd_src := exec.Command("curl", "-L", gomodifytags_src_url, "-o", "source.tar.gz")
	err = gomodifytags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gomodifytags_cmd_direct := exec.Command("./binary")
	err = gomodifytags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
