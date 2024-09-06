package main

// Redict - Distributed key/value database
// Homepage: https://redict.io/

import (
	"fmt"
	
	"os/exec"
)

func installRedict() {
	// Método 1: Descargar y extraer .tar.gz
	redict_tar_url := "https://codeberg.org/redict/redict/archive/7.3.0.tar.gz"
	redict_cmd_tar := exec.Command("curl", "-L", redict_tar_url, "-o", "package.tar.gz")
	err := redict_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redict_zip_url := "https://codeberg.org/redict/redict/archive/7.3.0.zip"
	redict_cmd_zip := exec.Command("curl", "-L", redict_zip_url, "-o", "package.zip")
	err = redict_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redict_bin_url := "https://codeberg.org/redict/redict/archive/7.3.0.bin"
	redict_cmd_bin := exec.Command("curl", "-L", redict_bin_url, "-o", "binary.bin")
	err = redict_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redict_src_url := "https://codeberg.org/redict/redict/archive/7.3.0.src.tar.gz"
	redict_cmd_src := exec.Command("curl", "-L", redict_src_url, "-o", "source.tar.gz")
	err = redict_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redict_cmd_direct := exec.Command("./binary")
	err = redict_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
