package main

// Hck - Sharp cut(1) clone
// Homepage: https://github.com/sstadick/hck

import (
	"fmt"
	
	"os/exec"
)

func installHck() {
	// Método 1: Descargar y extraer .tar.gz
	hck_tar_url := "https://github.com/sstadick/hck/archive/refs/tags/v0.10.1.tar.gz"
	hck_cmd_tar := exec.Command("curl", "-L", hck_tar_url, "-o", "package.tar.gz")
	err := hck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hck_zip_url := "https://github.com/sstadick/hck/archive/refs/tags/v0.10.1.zip"
	hck_cmd_zip := exec.Command("curl", "-L", hck_zip_url, "-o", "package.zip")
	err = hck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hck_bin_url := "https://github.com/sstadick/hck/archive/refs/tags/v0.10.1.bin"
	hck_cmd_bin := exec.Command("curl", "-L", hck_bin_url, "-o", "binary.bin")
	err = hck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hck_src_url := "https://github.com/sstadick/hck/archive/refs/tags/v0.10.1.src.tar.gz"
	hck_cmd_src := exec.Command("curl", "-L", hck_src_url, "-o", "source.tar.gz")
	err = hck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hck_cmd_direct := exec.Command("./binary")
	err = hck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
