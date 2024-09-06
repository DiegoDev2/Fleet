package main

// Libopennet - Provides open_net() (similar to open())
// Homepage: https://www.rkeene.org/oss/libopennet

import (
	"fmt"
	
	"os/exec"
)

func installLibopennet() {
	// Método 1: Descargar y extraer .tar.gz
	libopennet_tar_url := "https://www.rkeene.org/files/oss/libopennet/libopennet-0.9.9.tar.gz"
	libopennet_cmd_tar := exec.Command("curl", "-L", libopennet_tar_url, "-o", "package.tar.gz")
	err := libopennet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libopennet_zip_url := "https://www.rkeene.org/files/oss/libopennet/libopennet-0.9.9.zip"
	libopennet_cmd_zip := exec.Command("curl", "-L", libopennet_zip_url, "-o", "package.zip")
	err = libopennet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libopennet_bin_url := "https://www.rkeene.org/files/oss/libopennet/libopennet-0.9.9.bin"
	libopennet_cmd_bin := exec.Command("curl", "-L", libopennet_bin_url, "-o", "binary.bin")
	err = libopennet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libopennet_src_url := "https://www.rkeene.org/files/oss/libopennet/libopennet-0.9.9.src.tar.gz"
	libopennet_cmd_src := exec.Command("curl", "-L", libopennet_src_url, "-o", "source.tar.gz")
	err = libopennet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libopennet_cmd_direct := exec.Command("./binary")
	err = libopennet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
