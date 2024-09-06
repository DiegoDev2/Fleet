package main

// Pegtl - Parsing Expression Grammar Template Library
// Homepage: https://github.com/taocpp/PEGTL

import (
	"fmt"
	
	"os/exec"
)

func installPegtl() {
	// Método 1: Descargar y extraer .tar.gz
	pegtl_tar_url := "https://github.com/taocpp/PEGTL/archive/refs/tags/3.2.7.tar.gz"
	pegtl_cmd_tar := exec.Command("curl", "-L", pegtl_tar_url, "-o", "package.tar.gz")
	err := pegtl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pegtl_zip_url := "https://github.com/taocpp/PEGTL/archive/refs/tags/3.2.7.zip"
	pegtl_cmd_zip := exec.Command("curl", "-L", pegtl_zip_url, "-o", "package.zip")
	err = pegtl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pegtl_bin_url := "https://github.com/taocpp/PEGTL/archive/refs/tags/3.2.7.bin"
	pegtl_cmd_bin := exec.Command("curl", "-L", pegtl_bin_url, "-o", "binary.bin")
	err = pegtl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pegtl_src_url := "https://github.com/taocpp/PEGTL/archive/refs/tags/3.2.7.src.tar.gz"
	pegtl_cmd_src := exec.Command("curl", "-L", pegtl_src_url, "-o", "source.tar.gz")
	err = pegtl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pegtl_cmd_direct := exec.Command("./binary")
	err = pegtl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
