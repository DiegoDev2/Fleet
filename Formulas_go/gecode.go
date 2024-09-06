package main

// Gecode - Toolkit for developing constraint-based systems and applications
// Homepage: https://www.gecode.org/

import (
	"fmt"
	
	"os/exec"
)

func installGecode() {
	// Método 1: Descargar y extraer .tar.gz
	gecode_tar_url := "https://github.com/Gecode/gecode/archive/refs/tags/release-6.2.0.tar.gz"
	gecode_cmd_tar := exec.Command("curl", "-L", gecode_tar_url, "-o", "package.tar.gz")
	err := gecode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gecode_zip_url := "https://github.com/Gecode/gecode/archive/refs/tags/release-6.2.0.zip"
	gecode_cmd_zip := exec.Command("curl", "-L", gecode_zip_url, "-o", "package.zip")
	err = gecode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gecode_bin_url := "https://github.com/Gecode/gecode/archive/refs/tags/release-6.2.0.bin"
	gecode_cmd_bin := exec.Command("curl", "-L", gecode_bin_url, "-o", "binary.bin")
	err = gecode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gecode_src_url := "https://github.com/Gecode/gecode/archive/refs/tags/release-6.2.0.src.tar.gz"
	gecode_cmd_src := exec.Command("curl", "-L", gecode_src_url, "-o", "source.tar.gz")
	err = gecode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gecode_cmd_direct := exec.Command("./binary")
	err = gecode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
}
