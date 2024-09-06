package main

// Cntb - Contabo Command-Line Interface (CLI)
// Homepage: https://github.com/contabo/cntb

import (
	"fmt"
	
	"os/exec"
)

func installCntb() {
	// Método 1: Descargar y extraer .tar.gz
	cntb_tar_url := "https://github.com/contabo/cntb/archive/refs/tags/v1.4.12.tar.gz"
	cntb_cmd_tar := exec.Command("curl", "-L", cntb_tar_url, "-o", "package.tar.gz")
	err := cntb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cntb_zip_url := "https://github.com/contabo/cntb/archive/refs/tags/v1.4.12.zip"
	cntb_cmd_zip := exec.Command("curl", "-L", cntb_zip_url, "-o", "package.zip")
	err = cntb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cntb_bin_url := "https://github.com/contabo/cntb/archive/refs/tags/v1.4.12.bin"
	cntb_cmd_bin := exec.Command("curl", "-L", cntb_bin_url, "-o", "binary.bin")
	err = cntb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cntb_src_url := "https://github.com/contabo/cntb/archive/refs/tags/v1.4.12.src.tar.gz"
	cntb_cmd_src := exec.Command("curl", "-L", cntb_src_url, "-o", "source.tar.gz")
	err = cntb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cntb_cmd_direct := exec.Command("./binary")
	err = cntb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
