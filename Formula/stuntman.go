package main

// Stuntman - Implementation of the STUN protocol
// Homepage: https://www.stunprotocol.org

import (
	"fmt"
	
	"os/exec"
)

func installStuntman() {
	// Método 1: Descargar y extraer .tar.gz
	stuntman_tar_url := "https://www.stunprotocol.org/stunserver-1.2.16.tgz"
	stuntman_cmd_tar := exec.Command("curl", "-L", stuntman_tar_url, "-o", "package.tar.gz")
	err := stuntman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stuntman_zip_url := "https://www.stunprotocol.org/stunserver-1.2.16.tgz"
	stuntman_cmd_zip := exec.Command("curl", "-L", stuntman_zip_url, "-o", "package.zip")
	err = stuntman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stuntman_bin_url := "https://www.stunprotocol.org/stunserver-1.2.16.tgz"
	stuntman_cmd_bin := exec.Command("curl", "-L", stuntman_bin_url, "-o", "binary.bin")
	err = stuntman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stuntman_src_url := "https://www.stunprotocol.org/stunserver-1.2.16.tgz"
	stuntman_cmd_src := exec.Command("curl", "-L", stuntman_src_url, "-o", "source.tar.gz")
	err = stuntman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stuntman_cmd_direct := exec.Command("./binary")
	err = stuntman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
