package main

// Base16384 - Encode binary files to printable utf16be
// Homepage: https://github.com/fumiama/base16384

import (
	"fmt"
	
	"os/exec"
)

func installBase16384() {
	// Método 1: Descargar y extraer .tar.gz
	base16384_tar_url := "https://github.com/fumiama/base16384/archive/refs/tags/v2.3.1.tar.gz"
	base16384_cmd_tar := exec.Command("curl", "-L", base16384_tar_url, "-o", "package.tar.gz")
	err := base16384_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	base16384_zip_url := "https://github.com/fumiama/base16384/archive/refs/tags/v2.3.1.zip"
	base16384_cmd_zip := exec.Command("curl", "-L", base16384_zip_url, "-o", "package.zip")
	err = base16384_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	base16384_bin_url := "https://github.com/fumiama/base16384/archive/refs/tags/v2.3.1.bin"
	base16384_cmd_bin := exec.Command("curl", "-L", base16384_bin_url, "-o", "binary.bin")
	err = base16384_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	base16384_src_url := "https://github.com/fumiama/base16384/archive/refs/tags/v2.3.1.src.tar.gz"
	base16384_cmd_src := exec.Command("curl", "-L", base16384_src_url, "-o", "source.tar.gz")
	err = base16384_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	base16384_cmd_direct := exec.Command("./binary")
	err = base16384_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
