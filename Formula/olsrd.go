package main

// Olsrd - Implementation of the optimized link state routing protocol
// Homepage: https://github.com/OLSR/olsrd

import (
	"fmt"
	
	"os/exec"
)

func installOlsrd() {
	// Método 1: Descargar y extraer .tar.gz
	olsrd_tar_url := "https://github.com/OLSR/olsrd/archive/refs/tags/v0.9.8.tar.gz"
	olsrd_cmd_tar := exec.Command("curl", "-L", olsrd_tar_url, "-o", "package.tar.gz")
	err := olsrd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	olsrd_zip_url := "https://github.com/OLSR/olsrd/archive/refs/tags/v0.9.8.zip"
	olsrd_cmd_zip := exec.Command("curl", "-L", olsrd_zip_url, "-o", "package.zip")
	err = olsrd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	olsrd_bin_url := "https://github.com/OLSR/olsrd/archive/refs/tags/v0.9.8.bin"
	olsrd_cmd_bin := exec.Command("curl", "-L", olsrd_bin_url, "-o", "binary.bin")
	err = olsrd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	olsrd_src_url := "https://github.com/OLSR/olsrd/archive/refs/tags/v0.9.8.src.tar.gz"
	olsrd_cmd_src := exec.Command("curl", "-L", olsrd_src_url, "-o", "source.tar.gz")
	err = olsrd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	olsrd_cmd_direct := exec.Command("./binary")
	err = olsrd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: gpsd")
	exec.Command("latte", "install", "gpsd").Run()
}
