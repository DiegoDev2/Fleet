package main

// Bossa - Flash utility for Atmel SAM microcontrollers
// Homepage: https://github.com/shumatech/BOSSA

import (
	"fmt"
	
	"os/exec"
)

func installBossa() {
	// Método 1: Descargar y extraer .tar.gz
	bossa_tar_url := "https://github.com/shumatech/BOSSA/archive/refs/tags/1.9.1.tar.gz"
	bossa_cmd_tar := exec.Command("curl", "-L", bossa_tar_url, "-o", "package.tar.gz")
	err := bossa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bossa_zip_url := "https://github.com/shumatech/BOSSA/archive/refs/tags/1.9.1.zip"
	bossa_cmd_zip := exec.Command("curl", "-L", bossa_zip_url, "-o", "package.zip")
	err = bossa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bossa_bin_url := "https://github.com/shumatech/BOSSA/archive/refs/tags/1.9.1.bin"
	bossa_cmd_bin := exec.Command("curl", "-L", bossa_bin_url, "-o", "binary.bin")
	err = bossa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bossa_src_url := "https://github.com/shumatech/BOSSA/archive/refs/tags/1.9.1.src.tar.gz"
	bossa_cmd_src := exec.Command("curl", "-L", bossa_src_url, "-o", "source.tar.gz")
	err = bossa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bossa_cmd_direct := exec.Command("./binary")
	err = bossa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
