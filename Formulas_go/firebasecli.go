package main

// FirebaseCli - Firebase command-line tools
// Homepage: https://firebase.google.com/docs/cli/

import (
	"fmt"
	
	"os/exec"
)

func installFirebaseCli() {
	// Método 1: Descargar y extraer .tar.gz
	firebasecli_tar_url := "https://registry.npmjs.org/firebase-tools/-/firebase-tools-13.16.0.tgz"
	firebasecli_cmd_tar := exec.Command("curl", "-L", firebasecli_tar_url, "-o", "package.tar.gz")
	err := firebasecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	firebasecli_zip_url := "https://registry.npmjs.org/firebase-tools/-/firebase-tools-13.16.0.tgz"
	firebasecli_cmd_zip := exec.Command("curl", "-L", firebasecli_zip_url, "-o", "package.zip")
	err = firebasecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	firebasecli_bin_url := "https://registry.npmjs.org/firebase-tools/-/firebase-tools-13.16.0.tgz"
	firebasecli_cmd_bin := exec.Command("curl", "-L", firebasecli_bin_url, "-o", "binary.bin")
	err = firebasecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	firebasecli_src_url := "https://registry.npmjs.org/firebase-tools/-/firebase-tools-13.16.0.tgz"
	firebasecli_cmd_src := exec.Command("curl", "-L", firebasecli_src_url, "-o", "source.tar.gz")
	err = firebasecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	firebasecli_cmd_direct := exec.Command("./binary")
	err = firebasecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
