package main

// Libzzip - Library providing read access on ZIP-archives
// Homepage: https://github.com/gdraheim/zziplib

import (
	"fmt"
	
	"os/exec"
)

func installLibzzip() {
	// Método 1: Descargar y extraer .tar.gz
	libzzip_tar_url := "https://github.com/gdraheim/zziplib/archive/refs/tags/v0.13.78.tar.gz"
	libzzip_cmd_tar := exec.Command("curl", "-L", libzzip_tar_url, "-o", "package.tar.gz")
	err := libzzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libzzip_zip_url := "https://github.com/gdraheim/zziplib/archive/refs/tags/v0.13.78.zip"
	libzzip_cmd_zip := exec.Command("curl", "-L", libzzip_zip_url, "-o", "package.zip")
	err = libzzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libzzip_bin_url := "https://github.com/gdraheim/zziplib/archive/refs/tags/v0.13.78.bin"
	libzzip_cmd_bin := exec.Command("curl", "-L", libzzip_bin_url, "-o", "binary.bin")
	err = libzzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libzzip_src_url := "https://github.com/gdraheim/zziplib/archive/refs/tags/v0.13.78.src.tar.gz"
	libzzip_cmd_src := exec.Command("curl", "-L", libzzip_src_url, "-o", "source.tar.gz")
	err = libzzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libzzip_cmd_direct := exec.Command("./binary")
	err = libzzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
