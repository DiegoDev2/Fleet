package main

// Chapel - Programming language for productive parallel computing at scale
// Homepage: https://chapel-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installChapel() {
	// Método 1: Descargar y extraer .tar.gz
	chapel_tar_url := "https://github.com/chapel-lang/chapel/releases/download/2.1.0/chapel-2.1.0.tar.gz"
	chapel_cmd_tar := exec.Command("curl", "-L", chapel_tar_url, "-o", "package.tar.gz")
	err := chapel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chapel_zip_url := "https://github.com/chapel-lang/chapel/releases/download/2.1.0/chapel-2.1.0.zip"
	chapel_cmd_zip := exec.Command("curl", "-L", chapel_zip_url, "-o", "package.zip")
	err = chapel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chapel_bin_url := "https://github.com/chapel-lang/chapel/releases/download/2.1.0/chapel-2.1.0.bin"
	chapel_cmd_bin := exec.Command("curl", "-L", chapel_bin_url, "-o", "binary.bin")
	err = chapel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chapel_src_url := "https://github.com/chapel-lang/chapel/releases/download/2.1.0/chapel-2.1.0.src.tar.gz"
	chapel_cmd_src := exec.Command("curl", "-L", chapel_src_url, "-o", "source.tar.gz")
	err = chapel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chapel_cmd_direct := exec.Command("./binary")
	err = chapel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: hwloc")
	exec.Command("latte", "install", "hwloc").Run()
	fmt.Println("Instalando dependencia: jemalloc")
	exec.Command("latte", "install", "jemalloc").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
