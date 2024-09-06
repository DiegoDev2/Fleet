package main

// Libmagic - Implementation of the file(1) command
// Homepage: https://www.darwinsys.com/file/

import (
	"fmt"
	
	"os/exec"
)

func installLibmagic() {
	// Método 1: Descargar y extraer .tar.gz
	libmagic_tar_url := "https://astron.com/pub/file/file-5.45.tar.gz"
	libmagic_cmd_tar := exec.Command("curl", "-L", libmagic_tar_url, "-o", "package.tar.gz")
	err := libmagic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmagic_zip_url := "https://astron.com/pub/file/file-5.45.zip"
	libmagic_cmd_zip := exec.Command("curl", "-L", libmagic_zip_url, "-o", "package.zip")
	err = libmagic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmagic_bin_url := "https://astron.com/pub/file/file-5.45.bin"
	libmagic_cmd_bin := exec.Command("curl", "-L", libmagic_bin_url, "-o", "binary.bin")
	err = libmagic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmagic_src_url := "https://astron.com/pub/file/file-5.45.src.tar.gz"
	libmagic_cmd_src := exec.Command("curl", "-L", libmagic_src_url, "-o", "source.tar.gz")
	err = libmagic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmagic_cmd_direct := exec.Command("./binary")
	err = libmagic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
