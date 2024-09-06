package main

// FaunaShell - Interactive shell for FaunaDB
// Homepage: https://fauna.com/

import (
	"fmt"
	
	"os/exec"
)

func installFaunaShell() {
	// Método 1: Descargar y extraer .tar.gz
	faunashell_tar_url := "https://registry.npmjs.org/fauna-shell/-/fauna-shell-2.0.2.tgz"
	faunashell_cmd_tar := exec.Command("curl", "-L", faunashell_tar_url, "-o", "package.tar.gz")
	err := faunashell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	faunashell_zip_url := "https://registry.npmjs.org/fauna-shell/-/fauna-shell-2.0.2.tgz"
	faunashell_cmd_zip := exec.Command("curl", "-L", faunashell_zip_url, "-o", "package.zip")
	err = faunashell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	faunashell_bin_url := "https://registry.npmjs.org/fauna-shell/-/fauna-shell-2.0.2.tgz"
	faunashell_cmd_bin := exec.Command("curl", "-L", faunashell_bin_url, "-o", "binary.bin")
	err = faunashell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	faunashell_src_url := "https://registry.npmjs.org/fauna-shell/-/fauna-shell-2.0.2.tgz"
	faunashell_cmd_src := exec.Command("curl", "-L", faunashell_src_url, "-o", "source.tar.gz")
	err = faunashell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	faunashell_cmd_direct := exec.Command("./binary")
	err = faunashell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
