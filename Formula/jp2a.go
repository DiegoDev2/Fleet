package main

// Jp2a - Convert JPG images to ASCII
// Homepage: https://github.com/Talinx/jp2a

import (
	"fmt"
	
	"os/exec"
)

func installJp2a() {
	// Método 1: Descargar y extraer .tar.gz
	jp2a_tar_url := "https://github.com/Talinx/jp2a/releases/download/v1.2.0/jp2a-1.2.0.tar.bz2"
	jp2a_cmd_tar := exec.Command("curl", "-L", jp2a_tar_url, "-o", "package.tar.gz")
	err := jp2a_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jp2a_zip_url := "https://github.com/Talinx/jp2a/releases/download/v1.2.0/jp2a-1.2.0.tar.bz2"
	jp2a_cmd_zip := exec.Command("curl", "-L", jp2a_zip_url, "-o", "package.zip")
	err = jp2a_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jp2a_bin_url := "https://github.com/Talinx/jp2a/releases/download/v1.2.0/jp2a-1.2.0.tar.bz2"
	jp2a_cmd_bin := exec.Command("curl", "-L", jp2a_bin_url, "-o", "binary.bin")
	err = jp2a_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jp2a_src_url := "https://github.com/Talinx/jp2a/releases/download/v1.2.0/jp2a-1.2.0.tar.bz2"
	jp2a_cmd_src := exec.Command("curl", "-L", jp2a_src_url, "-o", "source.tar.gz")
	err = jp2a_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jp2a_cmd_direct := exec.Command("./binary")
	err = jp2a_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
