package main

// SNail - Fork of Heirloom mailx
// Homepage: https://www.sdaoden.eu/code.html

import (
	"fmt"
	
	"os/exec"
)

func installSNail() {
	// Método 1: Descargar y extraer .tar.gz
	snail_tar_url := "https://www.sdaoden.eu/downloads/s-nail-14.9.25.tar.xz"
	snail_cmd_tar := exec.Command("curl", "-L", snail_tar_url, "-o", "package.tar.gz")
	err := snail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snail_zip_url := "https://www.sdaoden.eu/downloads/s-nail-14.9.25.tar.xz"
	snail_cmd_zip := exec.Command("curl", "-L", snail_zip_url, "-o", "package.zip")
	err = snail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snail_bin_url := "https://www.sdaoden.eu/downloads/s-nail-14.9.25.tar.xz"
	snail_cmd_bin := exec.Command("curl", "-L", snail_bin_url, "-o", "binary.bin")
	err = snail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snail_src_url := "https://www.sdaoden.eu/downloads/s-nail-14.9.25.tar.xz"
	snail_cmd_src := exec.Command("curl", "-L", snail_src_url, "-o", "source.tar.gz")
	err = snail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snail_cmd_direct := exec.Command("./binary")
	err = snail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
