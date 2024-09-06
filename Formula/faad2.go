package main

// Faad2 - ISO AAC audio decoder
// Homepage: https://sourceforge.net/projects/faac/

import (
	"fmt"
	
	"os/exec"
)

func installFaad2() {
	// Método 1: Descargar y extraer .tar.gz
	faad2_tar_url := "https://github.com/knik0/faad2/archive/refs/tags/2.11.1.tar.gz"
	faad2_cmd_tar := exec.Command("curl", "-L", faad2_tar_url, "-o", "package.tar.gz")
	err := faad2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	faad2_zip_url := "https://github.com/knik0/faad2/archive/refs/tags/2.11.1.zip"
	faad2_cmd_zip := exec.Command("curl", "-L", faad2_zip_url, "-o", "package.zip")
	err = faad2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	faad2_bin_url := "https://github.com/knik0/faad2/archive/refs/tags/2.11.1.bin"
	faad2_cmd_bin := exec.Command("curl", "-L", faad2_bin_url, "-o", "binary.bin")
	err = faad2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	faad2_src_url := "https://github.com/knik0/faad2/archive/refs/tags/2.11.1.src.tar.gz"
	faad2_cmd_src := exec.Command("curl", "-L", faad2_src_url, "-o", "source.tar.gz")
	err = faad2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	faad2_cmd_direct := exec.Command("./binary")
	err = faad2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
