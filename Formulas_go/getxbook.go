package main

// Getxbook - Tools to download ebooks from various sources
// Homepage: https://njw.name/getxbook/

import (
	"fmt"
	
	"os/exec"
)

func installGetxbook() {
	// Método 1: Descargar y extraer .tar.gz
	getxbook_tar_url := "https://njw.name/getxbook/getxbook-1.2.tar.xz"
	getxbook_cmd_tar := exec.Command("curl", "-L", getxbook_tar_url, "-o", "package.tar.gz")
	err := getxbook_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	getxbook_zip_url := "https://njw.name/getxbook/getxbook-1.2.tar.xz"
	getxbook_cmd_zip := exec.Command("curl", "-L", getxbook_zip_url, "-o", "package.zip")
	err = getxbook_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	getxbook_bin_url := "https://njw.name/getxbook/getxbook-1.2.tar.xz"
	getxbook_cmd_bin := exec.Command("curl", "-L", getxbook_bin_url, "-o", "binary.bin")
	err = getxbook_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	getxbook_src_url := "https://njw.name/getxbook/getxbook-1.2.tar.xz"
	getxbook_cmd_src := exec.Command("curl", "-L", getxbook_src_url, "-o", "source.tar.gz")
	err = getxbook_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	getxbook_cmd_direct := exec.Command("./binary")
	err = getxbook_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
