package main

// Dc3dd - Patched GNU dd that is intended for forensic acquisition of data
// Homepage: https://sourceforge.net/projects/dc3dd/

import (
	"fmt"
	
	"os/exec"
)

func installDc3dd() {
	// Método 1: Descargar y extraer .tar.gz
	dc3dd_tar_url := "https://downloads.sourceforge.net/project/dc3dd/dc3dd/7.3.1/dc3dd-7.3.1.zip"
	dc3dd_cmd_tar := exec.Command("curl", "-L", dc3dd_tar_url, "-o", "package.tar.gz")
	err := dc3dd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dc3dd_zip_url := "https://downloads.sourceforge.net/project/dc3dd/dc3dd/7.3.1/dc3dd-7.3.1.zip"
	dc3dd_cmd_zip := exec.Command("curl", "-L", dc3dd_zip_url, "-o", "package.zip")
	err = dc3dd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dc3dd_bin_url := "https://downloads.sourceforge.net/project/dc3dd/dc3dd/7.3.1/dc3dd-7.3.1.zip"
	dc3dd_cmd_bin := exec.Command("curl", "-L", dc3dd_bin_url, "-o", "binary.bin")
	err = dc3dd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dc3dd_src_url := "https://downloads.sourceforge.net/project/dc3dd/dc3dd/7.3.1/dc3dd-7.3.1.zip"
	dc3dd_cmd_src := exec.Command("curl", "-L", dc3dd_src_url, "-o", "source.tar.gz")
	err = dc3dd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dc3dd_cmd_direct := exec.Command("./binary")
	err = dc3dd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
