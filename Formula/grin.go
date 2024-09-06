package main

// Grin - Minimal implementation of the Mimblewimble protocol
// Homepage: https://grin.mw/

import (
	"fmt"
	
	"os/exec"
)

func installGrin() {
	// Método 1: Descargar y extraer .tar.gz
	grin_tar_url := "https://github.com/mimblewimble/grin/archive/refs/tags/v5.3.2.tar.gz"
	grin_cmd_tar := exec.Command("curl", "-L", grin_tar_url, "-o", "package.tar.gz")
	err := grin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grin_zip_url := "https://github.com/mimblewimble/grin/archive/refs/tags/v5.3.2.zip"
	grin_cmd_zip := exec.Command("curl", "-L", grin_zip_url, "-o", "package.zip")
	err = grin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grin_bin_url := "https://github.com/mimblewimble/grin/archive/refs/tags/v5.3.2.bin"
	grin_cmd_bin := exec.Command("curl", "-L", grin_bin_url, "-o", "binary.bin")
	err = grin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grin_src_url := "https://github.com/mimblewimble/grin/archive/refs/tags/v5.3.2.src.tar.gz"
	grin_cmd_src := exec.Command("curl", "-L", grin_src_url, "-o", "source.tar.gz")
	err = grin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grin_cmd_direct := exec.Command("./binary")
	err = grin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
