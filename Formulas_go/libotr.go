package main

// Libotr - Off-The-Record (OTR) messaging library
// Homepage: https://otr.cypherpunks.ca/

import (
	"fmt"
	
	"os/exec"
)

func installLibotr() {
	// Método 1: Descargar y extraer .tar.gz
	libotr_tar_url := "https://otr.cypherpunks.ca/libotr-4.1.1.tar.gz"
	libotr_cmd_tar := exec.Command("curl", "-L", libotr_tar_url, "-o", "package.tar.gz")
	err := libotr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libotr_zip_url := "https://otr.cypherpunks.ca/libotr-4.1.1.zip"
	libotr_cmd_zip := exec.Command("curl", "-L", libotr_zip_url, "-o", "package.zip")
	err = libotr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libotr_bin_url := "https://otr.cypherpunks.ca/libotr-4.1.1.bin"
	libotr_cmd_bin := exec.Command("curl", "-L", libotr_bin_url, "-o", "binary.bin")
	err = libotr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libotr_src_url := "https://otr.cypherpunks.ca/libotr-4.1.1.src.tar.gz"
	libotr_cmd_src := exec.Command("curl", "-L", libotr_src_url, "-o", "source.tar.gz")
	err = libotr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libotr_cmd_direct := exec.Command("./binary")
	err = libotr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
}
