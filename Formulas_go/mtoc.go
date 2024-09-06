package main

// Mtoc - Mach-O to PE/COFF binary converter
// Homepage: https://opensource.apple.com/

import (
	"fmt"
	
	"os/exec"
)

func installMtoc() {
	// Método 1: Descargar y extraer .tar.gz
	mtoc_tar_url := "https://github.com/apple-oss-distributions/cctools/archive/refs/tags/cctools-1010.6.tar.gz"
	mtoc_cmd_tar := exec.Command("curl", "-L", mtoc_tar_url, "-o", "package.tar.gz")
	err := mtoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mtoc_zip_url := "https://github.com/apple-oss-distributions/cctools/archive/refs/tags/cctools-1010.6.zip"
	mtoc_cmd_zip := exec.Command("curl", "-L", mtoc_zip_url, "-o", "package.zip")
	err = mtoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mtoc_bin_url := "https://github.com/apple-oss-distributions/cctools/archive/refs/tags/cctools-1010.6.bin"
	mtoc_cmd_bin := exec.Command("curl", "-L", mtoc_bin_url, "-o", "binary.bin")
	err = mtoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mtoc_src_url := "https://github.com/apple-oss-distributions/cctools/archive/refs/tags/cctools-1010.6.src.tar.gz"
	mtoc_cmd_src := exec.Command("curl", "-L", mtoc_src_url, "-o", "source.tar.gz")
	err = mtoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mtoc_cmd_direct := exec.Command("./binary")
	err = mtoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
