package main

// X265 - H.265/HEVC encoder
// Homepage: https://bitbucket.org/multicoreware/x265_git

import (
	"fmt"
	
	"os/exec"
)

func installX265() {
	// Método 1: Descargar y extraer .tar.gz
	x265_tar_url := "https://bitbucket.org/multicoreware/x265_git/get/3.6.tar.gz"
	x265_cmd_tar := exec.Command("curl", "-L", x265_tar_url, "-o", "package.tar.gz")
	err := x265_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	x265_zip_url := "https://bitbucket.org/multicoreware/x265_git/get/3.6.zip"
	x265_cmd_zip := exec.Command("curl", "-L", x265_zip_url, "-o", "package.zip")
	err = x265_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	x265_bin_url := "https://bitbucket.org/multicoreware/x265_git/get/3.6.bin"
	x265_cmd_bin := exec.Command("curl", "-L", x265_bin_url, "-o", "binary.bin")
	err = x265_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	x265_src_url := "https://bitbucket.org/multicoreware/x265_git/get/3.6.src.tar.gz"
	x265_cmd_src := exec.Command("curl", "-L", x265_src_url, "-o", "source.tar.gz")
	err = x265_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	x265_cmd_direct := exec.Command("./binary")
	err = x265_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
}
