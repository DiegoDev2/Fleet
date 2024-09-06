package main

// Makensis - System to create Windows installers
// Homepage: https://nsis.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMakensis() {
	// Método 1: Descargar y extraer .tar.gz
	makensis_tar_url := "https://downloads.sourceforge.net/project/nsis/NSIS%203/3.10/nsis-3.10-src.tar.bz2"
	makensis_cmd_tar := exec.Command("curl", "-L", makensis_tar_url, "-o", "package.tar.gz")
	err := makensis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	makensis_zip_url := "https://downloads.sourceforge.net/project/nsis/NSIS%203/3.10/nsis-3.10-src.tar.bz2"
	makensis_cmd_zip := exec.Command("curl", "-L", makensis_zip_url, "-o", "package.zip")
	err = makensis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	makensis_bin_url := "https://downloads.sourceforge.net/project/nsis/NSIS%203/3.10/nsis-3.10-src.tar.bz2"
	makensis_cmd_bin := exec.Command("curl", "-L", makensis_bin_url, "-o", "binary.bin")
	err = makensis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	makensis_src_url := "https://downloads.sourceforge.net/project/nsis/NSIS%203/3.10/nsis-3.10-src.tar.bz2"
	makensis_cmd_src := exec.Command("curl", "-L", makensis_src_url, "-o", "source.tar.gz")
	err = makensis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	makensis_cmd_direct := exec.Command("./binary")
	err = makensis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mingw-w64")
	exec.Command("latte", "install", "mingw-w64").Run()
	fmt.Println("Instalando dependencia: scons")
	exec.Command("latte", "install", "scons").Run()
}
