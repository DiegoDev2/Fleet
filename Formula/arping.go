package main

// Arping - Utility to check whether MAC addresses are already taken on a LAN
// Homepage: https://github.com/ThomasHabets/arping

import (
	"fmt"
	
	"os/exec"
)

func installArping() {
	// Método 1: Descargar y extraer .tar.gz
	arping_tar_url := "https://github.com/ThomasHabets/arping/archive/refs/tags/arping-2.25.tar.gz"
	arping_cmd_tar := exec.Command("curl", "-L", arping_tar_url, "-o", "package.tar.gz")
	err := arping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arping_zip_url := "https://github.com/ThomasHabets/arping/archive/refs/tags/arping-2.25.zip"
	arping_cmd_zip := exec.Command("curl", "-L", arping_zip_url, "-o", "package.zip")
	err = arping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arping_bin_url := "https://github.com/ThomasHabets/arping/archive/refs/tags/arping-2.25.bin"
	arping_cmd_bin := exec.Command("curl", "-L", arping_bin_url, "-o", "binary.bin")
	err = arping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arping_src_url := "https://github.com/ThomasHabets/arping/archive/refs/tags/arping-2.25.src.tar.gz"
	arping_cmd_src := exec.Command("curl", "-L", arping_src_url, "-o", "source.tar.gz")
	err = arping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arping_cmd_direct := exec.Command("./binary")
	err = arping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libnet")
	exec.Command("latte", "install", "libnet").Run()
}
