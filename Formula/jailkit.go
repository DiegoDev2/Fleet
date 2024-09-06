package main

// Jailkit - Utilities to create limited user accounts in a chroot jail
// Homepage: https://olivier.sessink.nl/jailkit/

import (
	"fmt"
	
	"os/exec"
)

func installJailkit() {
	// Método 1: Descargar y extraer .tar.gz
	jailkit_tar_url := "https://olivier.sessink.nl/jailkit/jailkit-2.23.tar.bz2"
	jailkit_cmd_tar := exec.Command("curl", "-L", jailkit_tar_url, "-o", "package.tar.gz")
	err := jailkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jailkit_zip_url := "https://olivier.sessink.nl/jailkit/jailkit-2.23.tar.bz2"
	jailkit_cmd_zip := exec.Command("curl", "-L", jailkit_zip_url, "-o", "package.zip")
	err = jailkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jailkit_bin_url := "https://olivier.sessink.nl/jailkit/jailkit-2.23.tar.bz2"
	jailkit_cmd_bin := exec.Command("curl", "-L", jailkit_bin_url, "-o", "binary.bin")
	err = jailkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jailkit_src_url := "https://olivier.sessink.nl/jailkit/jailkit-2.23.tar.bz2"
	jailkit_cmd_src := exec.Command("curl", "-L", jailkit_src_url, "-o", "source.tar.gz")
	err = jailkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jailkit_cmd_direct := exec.Command("./binary")
	err = jailkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
