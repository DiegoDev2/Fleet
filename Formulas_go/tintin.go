package main

// Tintin - MUD client
// Homepage: https://tintin.mudhalla.net/

import (
	"fmt"
	
	"os/exec"
)

func installTintin() {
	// Método 1: Descargar y extraer .tar.gz
	tintin_tar_url := "https://github.com/scandum/tintin/releases/download/2.02.41/tintin-2.02.41.tar.gz"
	tintin_cmd_tar := exec.Command("curl", "-L", tintin_tar_url, "-o", "package.tar.gz")
	err := tintin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tintin_zip_url := "https://github.com/scandum/tintin/releases/download/2.02.41/tintin-2.02.41.zip"
	tintin_cmd_zip := exec.Command("curl", "-L", tintin_zip_url, "-o", "package.zip")
	err = tintin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tintin_bin_url := "https://github.com/scandum/tintin/releases/download/2.02.41/tintin-2.02.41.bin"
	tintin_cmd_bin := exec.Command("curl", "-L", tintin_bin_url, "-o", "binary.bin")
	err = tintin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tintin_src_url := "https://github.com/scandum/tintin/releases/download/2.02.41/tintin-2.02.41.src.tar.gz"
	tintin_cmd_src := exec.Command("curl", "-L", tintin_src_url, "-o", "source.tar.gz")
	err = tintin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tintin_cmd_direct := exec.Command("./binary")
	err = tintin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
}
