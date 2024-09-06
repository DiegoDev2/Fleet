package main

// Ne - Text editor based on the POSIX standard
// Homepage: https://github.com/vigna/ne

import (
	"fmt"
	
	"os/exec"
)

func installNe() {
	// Método 1: Descargar y extraer .tar.gz
	ne_tar_url := "https://github.com/vigna/ne/archive/refs/tags/3.3.3.tar.gz"
	ne_cmd_tar := exec.Command("curl", "-L", ne_tar_url, "-o", "package.tar.gz")
	err := ne_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ne_zip_url := "https://github.com/vigna/ne/archive/refs/tags/3.3.3.zip"
	ne_cmd_zip := exec.Command("curl", "-L", ne_zip_url, "-o", "package.zip")
	err = ne_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ne_bin_url := "https://github.com/vigna/ne/archive/refs/tags/3.3.3.bin"
	ne_cmd_bin := exec.Command("curl", "-L", ne_bin_url, "-o", "binary.bin")
	err = ne_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ne_src_url := "https://github.com/vigna/ne/archive/refs/tags/3.3.3.src.tar.gz"
	ne_cmd_src := exec.Command("curl", "-L", ne_src_url, "-o", "source.tar.gz")
	err = ne_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ne_cmd_direct := exec.Command("./binary")
	err = ne_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
}
