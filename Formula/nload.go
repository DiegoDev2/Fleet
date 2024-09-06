package main

// Nload - Realtime console network usage monitor
// Homepage: https://www.roland-riegel.de/nload/

import (
	"fmt"
	
	"os/exec"
)

func installNload() {
	// Método 1: Descargar y extraer .tar.gz
	nload_tar_url := "https://www.roland-riegel.de/nload/nload-0.7.4.tar.gz"
	nload_cmd_tar := exec.Command("curl", "-L", nload_tar_url, "-o", "package.tar.gz")
	err := nload_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nload_zip_url := "https://www.roland-riegel.de/nload/nload-0.7.4.zip"
	nload_cmd_zip := exec.Command("curl", "-L", nload_zip_url, "-o", "package.zip")
	err = nload_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nload_bin_url := "https://www.roland-riegel.de/nload/nload-0.7.4.bin"
	nload_cmd_bin := exec.Command("curl", "-L", nload_bin_url, "-o", "binary.bin")
	err = nload_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nload_src_url := "https://www.roland-riegel.de/nload/nload-0.7.4.src.tar.gz"
	nload_cmd_src := exec.Command("curl", "-L", nload_src_url, "-o", "source.tar.gz")
	err = nload_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nload_cmd_direct := exec.Command("./binary")
	err = nload_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
