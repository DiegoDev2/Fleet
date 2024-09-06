package main

// Opendbx - Lightweight but extensible database access library in C
// Homepage: https://linuxnetworks.de/doc/index.php/OpenDBX

import (
	"fmt"
	
	"os/exec"
)

func installOpendbx() {
	// Método 1: Descargar y extraer .tar.gz
	opendbx_tar_url := "https://linuxnetworks.de/opendbx/download/opendbx-1.4.6.tar.gz"
	opendbx_cmd_tar := exec.Command("curl", "-L", opendbx_tar_url, "-o", "package.tar.gz")
	err := opendbx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opendbx_zip_url := "https://linuxnetworks.de/opendbx/download/opendbx-1.4.6.zip"
	opendbx_cmd_zip := exec.Command("curl", "-L", opendbx_zip_url, "-o", "package.zip")
	err = opendbx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opendbx_bin_url := "https://linuxnetworks.de/opendbx/download/opendbx-1.4.6.bin"
	opendbx_cmd_bin := exec.Command("curl", "-L", opendbx_bin_url, "-o", "binary.bin")
	err = opendbx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opendbx_src_url := "https://linuxnetworks.de/opendbx/download/opendbx-1.4.6.src.tar.gz"
	opendbx_cmd_src := exec.Command("curl", "-L", opendbx_src_url, "-o", "source.tar.gz")
	err = opendbx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opendbx_cmd_direct := exec.Command("./binary")
	err = opendbx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
