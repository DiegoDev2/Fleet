package main

// SLang - Library for creating multi-platform software
// Homepage: https://www.jedsoft.org/slang/

import (
	"fmt"
	
	"os/exec"
)

func installSLang() {
	// Método 1: Descargar y extraer .tar.gz
	slang_tar_url := "https://www.jedsoft.org/releases/slang/slang-2.3.3.tar.bz2"
	slang_cmd_tar := exec.Command("curl", "-L", slang_tar_url, "-o", "package.tar.gz")
	err := slang_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slang_zip_url := "https://www.jedsoft.org/releases/slang/slang-2.3.3.tar.bz2"
	slang_cmd_zip := exec.Command("curl", "-L", slang_zip_url, "-o", "package.zip")
	err = slang_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slang_bin_url := "https://www.jedsoft.org/releases/slang/slang-2.3.3.tar.bz2"
	slang_cmd_bin := exec.Command("curl", "-L", slang_bin_url, "-o", "binary.bin")
	err = slang_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slang_src_url := "https://www.jedsoft.org/releases/slang/slang-2.3.3.tar.bz2"
	slang_cmd_src := exec.Command("curl", "-L", slang_src_url, "-o", "source.tar.gz")
	err = slang_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slang_cmd_direct := exec.Command("./binary")
	err = slang_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
