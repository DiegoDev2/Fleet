package main

// Faac - ISO AAC audio encoder
// Homepage: https://sourceforge.net/projects/faac/

import (
	"fmt"
	
	"os/exec"
)

func installFaac() {
	// Método 1: Descargar y extraer .tar.gz
	faac_tar_url := "https://github.com/knik0/faac/archive/refs/tags/1_30.tar.gz"
	faac_cmd_tar := exec.Command("curl", "-L", faac_tar_url, "-o", "package.tar.gz")
	err := faac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	faac_zip_url := "https://github.com/knik0/faac/archive/refs/tags/1_30.zip"
	faac_cmd_zip := exec.Command("curl", "-L", faac_zip_url, "-o", "package.zip")
	err = faac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	faac_bin_url := "https://github.com/knik0/faac/archive/refs/tags/1_30.bin"
	faac_cmd_bin := exec.Command("curl", "-L", faac_bin_url, "-o", "binary.bin")
	err = faac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	faac_src_url := "https://github.com/knik0/faac/archive/refs/tags/1_30.src.tar.gz"
	faac_cmd_src := exec.Command("curl", "-L", faac_src_url, "-o", "source.tar.gz")
	err = faac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	faac_cmd_direct := exec.Command("./binary")
	err = faac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
