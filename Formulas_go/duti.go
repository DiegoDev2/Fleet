package main

// Duti - Select default apps for documents and URL schemes on macOS
// Homepage: https://github.com/moretension/duti/

import (
	"fmt"
	
	"os/exec"
)

func installDuti() {
	// Método 1: Descargar y extraer .tar.gz
	duti_tar_url := "https://github.com/moretension/duti/archive/refs/tags/duti-1.5.4.tar.gz"
	duti_cmd_tar := exec.Command("curl", "-L", duti_tar_url, "-o", "package.tar.gz")
	err := duti_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	duti_zip_url := "https://github.com/moretension/duti/archive/refs/tags/duti-1.5.4.zip"
	duti_cmd_zip := exec.Command("curl", "-L", duti_zip_url, "-o", "package.zip")
	err = duti_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	duti_bin_url := "https://github.com/moretension/duti/archive/refs/tags/duti-1.5.4.bin"
	duti_cmd_bin := exec.Command("curl", "-L", duti_bin_url, "-o", "binary.bin")
	err = duti_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	duti_src_url := "https://github.com/moretension/duti/archive/refs/tags/duti-1.5.4.src.tar.gz"
	duti_cmd_src := exec.Command("curl", "-L", duti_src_url, "-o", "source.tar.gz")
	err = duti_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	duti_cmd_direct := exec.Command("./binary")
	err = duti_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
}
