package main

// Elfx86exts - Decodes x86 binaries (ELF and Mach-O) and prints out ISA extensions in use
// Homepage: https://github.com/pkgw/elfx86exts

import (
	"fmt"
	
	"os/exec"
)

func installElfx86exts() {
	// Método 1: Descargar y extraer .tar.gz
	elfx86exts_tar_url := "https://github.com/pkgw/elfx86exts/archive/refs/tags/elfx86exts@0.6.2.tar.gz"
	elfx86exts_cmd_tar := exec.Command("curl", "-L", elfx86exts_tar_url, "-o", "package.tar.gz")
	err := elfx86exts_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	elfx86exts_zip_url := "https://github.com/pkgw/elfx86exts/archive/refs/tags/elfx86exts@0.6.2.zip"
	elfx86exts_cmd_zip := exec.Command("curl", "-L", elfx86exts_zip_url, "-o", "package.zip")
	err = elfx86exts_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	elfx86exts_bin_url := "https://github.com/pkgw/elfx86exts/archive/refs/tags/elfx86exts@0.6.2.bin"
	elfx86exts_cmd_bin := exec.Command("curl", "-L", elfx86exts_bin_url, "-o", "binary.bin")
	err = elfx86exts_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	elfx86exts_src_url := "https://github.com/pkgw/elfx86exts/archive/refs/tags/elfx86exts@0.6.2.src.tar.gz"
	elfx86exts_cmd_src := exec.Command("curl", "-L", elfx86exts_src_url, "-o", "source.tar.gz")
	err = elfx86exts_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	elfx86exts_cmd_direct := exec.Command("./binary")
	err = elfx86exts_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: capstone")
	exec.Command("latte", "install", "capstone").Run()
}
