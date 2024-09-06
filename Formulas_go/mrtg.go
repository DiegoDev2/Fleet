package main

// Mrtg - Multi router traffic grapher
// Homepage: https://oss.oetiker.ch/mrtg/

import (
	"fmt"
	
	"os/exec"
)

func installMrtg() {
	// Método 1: Descargar y extraer .tar.gz
	mrtg_tar_url := "https://oss.oetiker.ch/mrtg/pub/mrtg-2.17.10.tar.gz"
	mrtg_cmd_tar := exec.Command("curl", "-L", mrtg_tar_url, "-o", "package.tar.gz")
	err := mrtg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mrtg_zip_url := "https://oss.oetiker.ch/mrtg/pub/mrtg-2.17.10.zip"
	mrtg_cmd_zip := exec.Command("curl", "-L", mrtg_zip_url, "-o", "package.zip")
	err = mrtg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mrtg_bin_url := "https://oss.oetiker.ch/mrtg/pub/mrtg-2.17.10.bin"
	mrtg_cmd_bin := exec.Command("curl", "-L", mrtg_bin_url, "-o", "binary.bin")
	err = mrtg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mrtg_src_url := "https://oss.oetiker.ch/mrtg/pub/mrtg-2.17.10.src.tar.gz"
	mrtg_cmd_src := exec.Command("curl", "-L", mrtg_src_url, "-o", "source.tar.gz")
	err = mrtg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mrtg_cmd_direct := exec.Command("./binary")
	err = mrtg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
