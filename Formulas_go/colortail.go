package main

// Colortail - Like tail(1), but with various colors for specified output
// Homepage: https://github.com/joakim666/colortail

import (
	"fmt"
	
	"os/exec"
)

func installColortail() {
	// Método 1: Descargar y extraer .tar.gz
	colortail_tar_url := "https://github.com/joakim666/colortail.git"
	colortail_cmd_tar := exec.Command("curl", "-L", colortail_tar_url, "-o", "package.tar.gz")
	err := colortail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	colortail_zip_url := "https://github.com/joakim666/colortail.git"
	colortail_cmd_zip := exec.Command("curl", "-L", colortail_zip_url, "-o", "package.zip")
	err = colortail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	colortail_bin_url := "https://github.com/joakim666/colortail.git"
	colortail_cmd_bin := exec.Command("curl", "-L", colortail_bin_url, "-o", "binary.bin")
	err = colortail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	colortail_src_url := "https://github.com/joakim666/colortail.git"
	colortail_cmd_src := exec.Command("curl", "-L", colortail_src_url, "-o", "source.tar.gz")
	err = colortail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	colortail_cmd_direct := exec.Command("./binary")
	err = colortail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
