package main

// Goawk - POSIX-compliant AWK interpreter written in Go
// Homepage: https://benhoyt.com/writings/goawk/

import (
	"fmt"
	
	"os/exec"
)

func installGoawk() {
	// Método 1: Descargar y extraer .tar.gz
	goawk_tar_url := "https://github.com/benhoyt/goawk/archive/refs/tags/v1.27.0.tar.gz"
	goawk_cmd_tar := exec.Command("curl", "-L", goawk_tar_url, "-o", "package.tar.gz")
	err := goawk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goawk_zip_url := "https://github.com/benhoyt/goawk/archive/refs/tags/v1.27.0.zip"
	goawk_cmd_zip := exec.Command("curl", "-L", goawk_zip_url, "-o", "package.zip")
	err = goawk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goawk_bin_url := "https://github.com/benhoyt/goawk/archive/refs/tags/v1.27.0.bin"
	goawk_cmd_bin := exec.Command("curl", "-L", goawk_bin_url, "-o", "binary.bin")
	err = goawk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goawk_src_url := "https://github.com/benhoyt/goawk/archive/refs/tags/v1.27.0.src.tar.gz"
	goawk_cmd_src := exec.Command("curl", "-L", goawk_src_url, "-o", "source.tar.gz")
	err = goawk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goawk_cmd_direct := exec.Command("./binary")
	err = goawk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
