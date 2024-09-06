package main

// Uvg266 - Open-source VVC/H.266 encoder
// Homepage: https://github.com/ultravideo/uvg266

import (
	"fmt"
	
	"os/exec"
)

func installUvg266() {
	// Método 1: Descargar y extraer .tar.gz
	uvg266_tar_url := "https://github.com/ultravideo/uvg266/archive/refs/tags/v0.8.1.tar.gz"
	uvg266_cmd_tar := exec.Command("curl", "-L", uvg266_tar_url, "-o", "package.tar.gz")
	err := uvg266_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uvg266_zip_url := "https://github.com/ultravideo/uvg266/archive/refs/tags/v0.8.1.zip"
	uvg266_cmd_zip := exec.Command("curl", "-L", uvg266_zip_url, "-o", "package.zip")
	err = uvg266_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uvg266_bin_url := "https://github.com/ultravideo/uvg266/archive/refs/tags/v0.8.1.bin"
	uvg266_cmd_bin := exec.Command("curl", "-L", uvg266_bin_url, "-o", "binary.bin")
	err = uvg266_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uvg266_src_url := "https://github.com/ultravideo/uvg266/archive/refs/tags/v0.8.1.src.tar.gz"
	uvg266_cmd_src := exec.Command("curl", "-L", uvg266_src_url, "-o", "source.tar.gz")
	err = uvg266_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uvg266_cmd_direct := exec.Command("./binary")
	err = uvg266_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
