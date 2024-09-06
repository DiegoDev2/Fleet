package main

// Cgoban - Go-related services
// Homepage: https://cgoban1.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCgoban() {
	// Método 1: Descargar y extraer .tar.gz
	cgoban_tar_url := "https://downloads.sourceforge.net/project/cgoban1/cgoban1/1.9.14/cgoban-1.9.14.tar.gz"
	cgoban_cmd_tar := exec.Command("curl", "-L", cgoban_tar_url, "-o", "package.tar.gz")
	err := cgoban_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cgoban_zip_url := "https://downloads.sourceforge.net/project/cgoban1/cgoban1/1.9.14/cgoban-1.9.14.zip"
	cgoban_cmd_zip := exec.Command("curl", "-L", cgoban_zip_url, "-o", "package.zip")
	err = cgoban_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cgoban_bin_url := "https://downloads.sourceforge.net/project/cgoban1/cgoban1/1.9.14/cgoban-1.9.14.bin"
	cgoban_cmd_bin := exec.Command("curl", "-L", cgoban_bin_url, "-o", "binary.bin")
	err = cgoban_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cgoban_src_url := "https://downloads.sourceforge.net/project/cgoban1/cgoban1/1.9.14/cgoban-1.9.14.src.tar.gz"
	cgoban_cmd_src := exec.Command("curl", "-L", cgoban_src_url, "-o", "source.tar.gz")
	err = cgoban_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cgoban_cmd_direct := exec.Command("./binary")
	err = cgoban_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
}
