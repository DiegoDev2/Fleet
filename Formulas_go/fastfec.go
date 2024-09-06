package main

// Fastfec - Extremely fast FEC filing parser written in C
// Homepage: https://www.washingtonpost.com/fastfec/

import (
	"fmt"
	
	"os/exec"
)

func installFastfec() {
	// Método 1: Descargar y extraer .tar.gz
	fastfec_tar_url := "https://github.com/washingtonpost/FastFEC/archive/refs/tags/0.2.0.tar.gz"
	fastfec_cmd_tar := exec.Command("curl", "-L", fastfec_tar_url, "-o", "package.tar.gz")
	err := fastfec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastfec_zip_url := "https://github.com/washingtonpost/FastFEC/archive/refs/tags/0.2.0.zip"
	fastfec_cmd_zip := exec.Command("curl", "-L", fastfec_zip_url, "-o", "package.zip")
	err = fastfec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastfec_bin_url := "https://github.com/washingtonpost/FastFEC/archive/refs/tags/0.2.0.bin"
	fastfec_cmd_bin := exec.Command("curl", "-L", fastfec_bin_url, "-o", "binary.bin")
	err = fastfec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastfec_src_url := "https://github.com/washingtonpost/FastFEC/archive/refs/tags/0.2.0.src.tar.gz"
	fastfec_cmd_src := exec.Command("curl", "-L", fastfec_src_url, "-o", "source.tar.gz")
	err = fastfec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastfec_cmd_direct := exec.Command("./binary")
	err = fastfec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: zig")
exec.Command("latte", "install", "zig")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
}
