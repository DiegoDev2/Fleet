package main

// Cloog - Generate code for scanning Z-polyhedra
// Homepage: https://github.com/periscop/cloog

import (
	"fmt"
	
	"os/exec"
)

func installCloog() {
	// Método 1: Descargar y extraer .tar.gz
	cloog_tar_url := "https://github.com/periscop/cloog/releases/download/cloog-0.21.1/cloog-0.21.1.tar.gz"
	cloog_cmd_tar := exec.Command("curl", "-L", cloog_tar_url, "-o", "package.tar.gz")
	err := cloog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloog_zip_url := "https://github.com/periscop/cloog/releases/download/cloog-0.21.1/cloog-0.21.1.zip"
	cloog_cmd_zip := exec.Command("curl", "-L", cloog_zip_url, "-o", "package.zip")
	err = cloog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloog_bin_url := "https://github.com/periscop/cloog/releases/download/cloog-0.21.1/cloog-0.21.1.bin"
	cloog_cmd_bin := exec.Command("curl", "-L", cloog_bin_url, "-o", "binary.bin")
	err = cloog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloog_src_url := "https://github.com/periscop/cloog/releases/download/cloog-0.21.1/cloog-0.21.1.src.tar.gz"
	cloog_cmd_src := exec.Command("curl", "-L", cloog_src_url, "-o", "source.tar.gz")
	err = cloog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloog_cmd_direct := exec.Command("./binary")
	err = cloog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: isl")
	exec.Command("latte", "install", "isl").Run()
}
