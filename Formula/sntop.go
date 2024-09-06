package main

// Sntop - Curses-based utility that polls hosts to determine connectivity
// Homepage: https://sntop.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSntop() {
	// Método 1: Descargar y extraer .tar.gz
	sntop_tar_url := "https://downloads.sourceforge.net/project/sntop/sntop/1.4.3/sntop-1.4.3.tar.gz"
	sntop_cmd_tar := exec.Command("curl", "-L", sntop_tar_url, "-o", "package.tar.gz")
	err := sntop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sntop_zip_url := "https://downloads.sourceforge.net/project/sntop/sntop/1.4.3/sntop-1.4.3.zip"
	sntop_cmd_zip := exec.Command("curl", "-L", sntop_zip_url, "-o", "package.zip")
	err = sntop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sntop_bin_url := "https://downloads.sourceforge.net/project/sntop/sntop/1.4.3/sntop-1.4.3.bin"
	sntop_cmd_bin := exec.Command("curl", "-L", sntop_bin_url, "-o", "binary.bin")
	err = sntop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sntop_src_url := "https://downloads.sourceforge.net/project/sntop/sntop/1.4.3/sntop-1.4.3.src.tar.gz"
	sntop_cmd_src := exec.Command("curl", "-L", sntop_src_url, "-o", "source.tar.gz")
	err = sntop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sntop_cmd_direct := exec.Command("./binary")
	err = sntop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fping")
	exec.Command("latte", "install", "fping").Run()
}
