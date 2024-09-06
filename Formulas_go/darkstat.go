package main

// Darkstat - Network traffic analyzer
// Homepage: https://unix4lyfe.org/darkstat/

import (
	"fmt"
	
	"os/exec"
)

func installDarkstat() {
	// Método 1: Descargar y extraer .tar.gz
	darkstat_tar_url := "https://github.com/emikulic/darkstat/archive/refs/tags/3.0.721.tar.gz"
	darkstat_cmd_tar := exec.Command("curl", "-L", darkstat_tar_url, "-o", "package.tar.gz")
	err := darkstat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	darkstat_zip_url := "https://github.com/emikulic/darkstat/archive/refs/tags/3.0.721.zip"
	darkstat_cmd_zip := exec.Command("curl", "-L", darkstat_zip_url, "-o", "package.zip")
	err = darkstat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	darkstat_bin_url := "https://github.com/emikulic/darkstat/archive/refs/tags/3.0.721.bin"
	darkstat_cmd_bin := exec.Command("curl", "-L", darkstat_bin_url, "-o", "binary.bin")
	err = darkstat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	darkstat_src_url := "https://github.com/emikulic/darkstat/archive/refs/tags/3.0.721.src.tar.gz"
	darkstat_cmd_src := exec.Command("curl", "-L", darkstat_src_url, "-o", "source.tar.gz")
	err = darkstat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	darkstat_cmd_direct := exec.Command("./binary")
	err = darkstat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
