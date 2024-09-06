package main

// Nsd - Name server daemon
// Homepage: https://www.nlnetlabs.nl/projects/nsd/

import (
	"fmt"
	
	"os/exec"
)

func installNsd() {
	// Método 1: Descargar y extraer .tar.gz
	nsd_tar_url := "https://www.nlnetlabs.nl/downloads/nsd/nsd-4.10.1.tar.gz"
	nsd_cmd_tar := exec.Command("curl", "-L", nsd_tar_url, "-o", "package.tar.gz")
	err := nsd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nsd_zip_url := "https://www.nlnetlabs.nl/downloads/nsd/nsd-4.10.1.zip"
	nsd_cmd_zip := exec.Command("curl", "-L", nsd_zip_url, "-o", "package.zip")
	err = nsd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nsd_bin_url := "https://www.nlnetlabs.nl/downloads/nsd/nsd-4.10.1.bin"
	nsd_cmd_bin := exec.Command("curl", "-L", nsd_bin_url, "-o", "binary.bin")
	err = nsd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nsd_src_url := "https://www.nlnetlabs.nl/downloads/nsd/nsd-4.10.1.src.tar.gz"
	nsd_cmd_src := exec.Command("curl", "-L", nsd_src_url, "-o", "source.tar.gz")
	err = nsd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nsd_cmd_direct := exec.Command("./binary")
	err = nsd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
