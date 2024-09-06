package main

// Cafeobj - New generation algebraic specification and programming language
// Homepage: https://cafeobj.org/

import (
	"fmt"
	
	"os/exec"
)

func installCafeobj() {
	// Método 1: Descargar y extraer .tar.gz
	cafeobj_tar_url := "https://github.com/CafeOBJ/cafeobj/archive/refs/tags/v1.6.1.tar.gz"
	cafeobj_cmd_tar := exec.Command("curl", "-L", cafeobj_tar_url, "-o", "package.tar.gz")
	err := cafeobj_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cafeobj_zip_url := "https://github.com/CafeOBJ/cafeobj/archive/refs/tags/v1.6.1.zip"
	cafeobj_cmd_zip := exec.Command("curl", "-L", cafeobj_zip_url, "-o", "package.zip")
	err = cafeobj_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cafeobj_bin_url := "https://github.com/CafeOBJ/cafeobj/archive/refs/tags/v1.6.1.bin"
	cafeobj_cmd_bin := exec.Command("curl", "-L", cafeobj_bin_url, "-o", "binary.bin")
	err = cafeobj_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cafeobj_src_url := "https://github.com/CafeOBJ/cafeobj/archive/refs/tags/v1.6.1.src.tar.gz"
	cafeobj_cmd_src := exec.Command("curl", "-L", cafeobj_src_url, "-o", "source.tar.gz")
	err = cafeobj_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cafeobj_cmd_direct := exec.Command("./binary")
	err = cafeobj_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sbcl")
exec.Command("latte", "install", "sbcl")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
