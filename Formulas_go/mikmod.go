package main

// Mikmod - Portable tracked music player
// Homepage: https://mikmod.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMikmod() {
	// Método 1: Descargar y extraer .tar.gz
	mikmod_tar_url := "https://downloads.sourceforge.net/project/mikmod/mikmod/3.2.8/mikmod-3.2.8.tar.gz"
	mikmod_cmd_tar := exec.Command("curl", "-L", mikmod_tar_url, "-o", "package.tar.gz")
	err := mikmod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mikmod_zip_url := "https://downloads.sourceforge.net/project/mikmod/mikmod/3.2.8/mikmod-3.2.8.zip"
	mikmod_cmd_zip := exec.Command("curl", "-L", mikmod_zip_url, "-o", "package.zip")
	err = mikmod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mikmod_bin_url := "https://downloads.sourceforge.net/project/mikmod/mikmod/3.2.8/mikmod-3.2.8.bin"
	mikmod_cmd_bin := exec.Command("curl", "-L", mikmod_bin_url, "-o", "binary.bin")
	err = mikmod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mikmod_src_url := "https://downloads.sourceforge.net/project/mikmod/mikmod/3.2.8/mikmod-3.2.8.src.tar.gz"
	mikmod_cmd_src := exec.Command("curl", "-L", mikmod_src_url, "-o", "source.tar.gz")
	err = mikmod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mikmod_cmd_direct := exec.Command("./binary")
	err = mikmod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libmikmod")
exec.Command("latte", "install", "libmikmod")
}
