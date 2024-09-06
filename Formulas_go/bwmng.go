package main

// BwmNg - Console-based live network and disk I/O bandwidth monitor
// Homepage: https://www.gropp.org/?id=projects&sub=bwm-ng

import (
	"fmt"
	
	"os/exec"
)

func installBwmNg() {
	// Método 1: Descargar y extraer .tar.gz
	bwmng_tar_url := "https://github.com/vgropp/bwm-ng/archive/refs/tags/v0.6.3.tar.gz"
	bwmng_cmd_tar := exec.Command("curl", "-L", bwmng_tar_url, "-o", "package.tar.gz")
	err := bwmng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bwmng_zip_url := "https://github.com/vgropp/bwm-ng/archive/refs/tags/v0.6.3.zip"
	bwmng_cmd_zip := exec.Command("curl", "-L", bwmng_zip_url, "-o", "package.zip")
	err = bwmng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bwmng_bin_url := "https://github.com/vgropp/bwm-ng/archive/refs/tags/v0.6.3.bin"
	bwmng_cmd_bin := exec.Command("curl", "-L", bwmng_bin_url, "-o", "binary.bin")
	err = bwmng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bwmng_src_url := "https://github.com/vgropp/bwm-ng/archive/refs/tags/v0.6.3.src.tar.gz"
	bwmng_cmd_src := exec.Command("curl", "-L", bwmng_src_url, "-o", "source.tar.gz")
	err = bwmng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bwmng_cmd_direct := exec.Command("./binary")
	err = bwmng_cmd_direct.Run()
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
