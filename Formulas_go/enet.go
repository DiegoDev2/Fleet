package main

// Enet - Provides a network communication layer on top of UDP
// Homepage: http://enet.bespin.org

import (
	"fmt"
	
	"os/exec"
)

func installEnet() {
	// Método 1: Descargar y extraer .tar.gz
	enet_tar_url := "http://enet.bespin.org/download/enet-1.3.18.tar.gz"
	enet_cmd_tar := exec.Command("curl", "-L", enet_tar_url, "-o", "package.tar.gz")
	err := enet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	enet_zip_url := "http://enet.bespin.org/download/enet-1.3.18.zip"
	enet_cmd_zip := exec.Command("curl", "-L", enet_zip_url, "-o", "package.zip")
	err = enet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	enet_bin_url := "http://enet.bespin.org/download/enet-1.3.18.bin"
	enet_cmd_bin := exec.Command("curl", "-L", enet_bin_url, "-o", "binary.bin")
	err = enet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	enet_src_url := "http://enet.bespin.org/download/enet-1.3.18.src.tar.gz"
	enet_cmd_src := exec.Command("curl", "-L", enet_src_url, "-o", "source.tar.gz")
	err = enet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	enet_cmd_direct := exec.Command("./binary")
	err = enet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
